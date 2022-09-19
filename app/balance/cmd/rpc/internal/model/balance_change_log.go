package model

import "time"

// MallBalanceChangeLog 余额变动记录表
type MallBalanceChangeLog struct {
	ID           int64     `xorm:"id"`
	BalanceId    int64     `xorm:"balance_id"`    // 余额ID
	Amount       int64     `xorm:"amount"`        // 变动金额
	BeforeAmount int64     `xorm:"before_amount"` // 变动前余额
	AfterAmount  int64     `xorm:"after_amount"`  // 变动后余额
	Type         int8      `xorm:"type"`          // 变动类型 \r\n1 增加\r\n2 减少
	TypeAmount   int8      `xorm:"type_amount"`   // 余额类型\r\n1 可用余额\r\n2 冻结余额
	Desc         string    `xorm:"desc"`          // 描述
	IsDelete     int8      `xorm:"is_delete"`     // 是否删除
	CreatedAt    time.Time `xorm:"created_at"`
	UpdatedAt    time.Time `xorm:"updated_at"`
}

const (
	// TypeIncrease 增加余额
	TypeIncrease = 1
	// TypeReduce 减少余额
	TypeReduce = 2

	// TypeAvailable 可用余额
	TypeAvailable = 1
	// TypeAmountFreeze 冻结余额
	TypeAmountFreeze = 2
)

// TableName 表名称
func (*MallBalanceChangeLog) TableName() string {
	return "mall_balance_change_log"
}
