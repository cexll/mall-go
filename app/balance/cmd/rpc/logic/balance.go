package logic

import (
	"errors"
	"fmt"
	"github.com/mix-go/xsql"
	"mall-go/app/balance/cmd/pb"
	"mall-go/app/balance/cmd/rpc/model"
	redsync "mall-go/common/lock"
	"time"
)

type BalanceLogic struct {
	model model.MallBalance
}

// SubFrozenBalance 增加冻结
func (t *BalanceLogic) SubFrozenBalance(in *pb.SubFrozenBalanceRequest) (bool, error) {
	balance, err := t.model.FindById(in.Id)
	if err != nil {
		return false, err
	}
	if balance.Status == 2 {
		return false, errors.New("钱包已被冻结")
	}

	if in.Amount > 0 {
		balance.Frozen += in.Amount
	}

	balance.UpdatedAt = time.Now()

	rows, err := t.model.UpdateByWhere(balance, []string{
		"id = ?",
	}, []any{
		in.Id,
	})
	if err != nil {
		return false, err
	}
	if rows == 0 {
		return false, nil
	}

	return true, nil
}

// ReduceFrozenBalance 减少冻结
func (t *BalanceLogic) ReduceFrozenBalance(in *pb.ReduceFrozenBalanceRequest) (bool, error) {
	balance, err := t.model.FindById(in.Id)
	if err != nil {
		return false, err
	}
	if balance.Status == 2 {
		return false, errors.New("钱包已被冻结")
	}
	if balance.Frozen < in.Amount {
		return false, errors.New("可操作金额小于操作金额")
	}
	lock := redsync.NewLock()
	mutex := lock.NewMutex(fmt.Sprintf("ReduceFrozenBalance: %d", in.Id))
	if err = mutex.Lock(); err != nil {
		fmt.Println(err)
		return false, errors.New("业务忙,正在操作")
	}
	defer func() {
		if ok, err := mutex.Unlock(); !ok || err != nil {
			fmt.Println(err)
		}
	}()

	if in.Amount > 0 {
		balance.Frozen -= in.Amount
	}

	balance.UpdatedAt = time.Now()

	rows, err := t.model.UpdateByWhere(balance, []string{
		"id = ?",
	}, []any{
		in.Id,
	})
	if err != nil {
		return false, err
	}

	if rows == 0 {
		return false, nil
	}

	return true, nil
}

// SubBalance 增加余额
func (t *BalanceLogic) SubBalance(in *pb.SubBalanceRequest) (bool, error) {
	balance, err := t.model.FindById(in.Id)
	if err != nil {
		return false, err
	}
	if balance.Status == 2 {
		return false, errors.New("钱包已被冻结")
	}

	if in.Amount > 0 {
		balance.Available += in.Amount
	}
	balance.UpdatedAt = time.Now()

	rows, err := t.model.UpdateByWhere(balance, []string{
		"id = ?",
	}, []any{
		in.Id,
	})
	if err != nil {
		return false, err
	}

	if rows == 0 {
		return false, nil
	}

	return true, nil
}

// ReduceBalance 减少余额
func (t *BalanceLogic) ReduceBalance(in *pb.ReduceBalanceRequest) (bool, error) {
	balance, err := t.model.FindById(in.Id)
	if err != nil {
		return false, err
	}
	if balance.Status == 2 {
		return false, errors.New("钱包已被冻结")
	}
	if balance.Available < in.Amount {
		return false, errors.New("可操作金额小于操作金额")
	}
	lock := redsync.NewLock()
	mutex := lock.NewMutex(fmt.Sprintf("ReduceBalance: %d", in.Id))
	if err = mutex.Lock(); err != nil {
		return false, errors.New("业务忙,正在操作")
	}
	defer func() {
		if ok, err := mutex.Unlock(); !ok || err != nil {
			fmt.Println(err)
		}
	}()
	if in.Amount > 0 {
		balance.Available += in.Amount
	}
	balance.UpdatedAt = time.Now()
	rows, err := t.model.UpdateByWhere(balance, []string{
		"id = ?",
	}, []any{
		in.Id,
	})
	if err != nil {
		return false, err
	}

	if rows == 0 {
		return false, nil
	}

	return true, nil
}

// GetBalance 获取钱包
func (t *BalanceLogic) GetBalance(in *pb.GetBalanceRequest) (*model.MallBalance, error) {
	balance, err := t.model.FindByWhere([]string{
		"id", "user_id", "type", "available", "frozen", "status",
	}, []string{
		"user_id = ?",
		"type = ?",
	}, []any{
		in.UserId,
		in.Type,
	}, []string{})

	if err != nil {
		if err == xsql.ErrNoRows {
			// 不存在则创建一个钱包
			newBalance := &model.MallBalance{
				UserId:    in.UserId,
				Type:      int8(in.Type),
				Available: 0,
				Frozen:    0,
				Status:    1,
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
			}
			id, err := t.model.CreateOne(newBalance)
			if err != nil {
				return balance, err
			}
			newBalance.ID = id

			return newBalance, nil
		}
		return balance, err
	}
	if balance.Status == 2 {
		return balance, errors.New("钱包已被冻结")
	}
	return balance, nil
}
