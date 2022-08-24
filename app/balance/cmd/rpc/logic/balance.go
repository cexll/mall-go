package logic

import (
	"database/sql"
	"errors"
	"fmt"
	"mall-go/app/balance/cmd/pb"
	"mall-go/app/balance/cmd/rpc/model"
	"mall-go/common/di"
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
		if err == sql.ErrNoRows {
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

type GetBalanceChangeListResult struct {
	List       []*pb.GetBalanceChangeListResponseList
	TotalPage  int64
	TotalCount int64
}

func (t *BalanceLogic) GetBalanceChangeList(in *pb.GetBalanceChangeListRequest) (*GetBalanceChangeListResult, error) {
	db := di.Gorm()
	var logs []model.MallBalanceChangeLog
	db.Where("balance_id = ?", in.Id)
	if in.Type != 0 {
		db.Where("type = ?", in.Type)
	}
	if in.TypeAmount != 0 {
		db.Where("type_amount = ?", in.TypeAmount)
	}
	db.Where("is_delete = ?", 0)

	db.Limit(int(in.PageSize)).Offset(int(in.Page - 1)).Find(&logs)
	if db.Error != nil {
		return nil, db.Error
	}
	var totalCount int64
	db.Count(&totalCount)
	if db.Error != nil {
		return nil, db.Error
	}
	var list []*pb.GetBalanceChangeListResponseList
	for _, v := range logs {
		list = append(list, &pb.GetBalanceChangeListResponseList{
			Id:           v.ID,
			Amount:       v.Amount,
			BeforeAmount: v.BeforeAmount,
			AfterAmount:  v.AfterAmount,
			Type:         int32(v.Type),
			TypeAmount:   int32(v.TypeAmount),
			CreatedAt:    v.CreatedAt.Local().Format("2006-01-02 15:01:05"),
		})
	}
	return &GetBalanceChangeListResult{
		TotalPage:  int64(int(totalCount / int64(in.PageSize))),
		TotalCount: totalCount,
		List:       list,
	}, nil
}
