package common

type BaseModel interface {
	FindByWhere()
	GetManyByWhere()
	CreateOne()
	CreateAll()
	UpdateByWhere()
	deleteOne()
	DeleteAll()
}
