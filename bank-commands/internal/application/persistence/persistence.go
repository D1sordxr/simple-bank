package persistence

type UoWManager interface {
	GetUoW() UnitOfWork
}

type UnitOfWork interface {
	Begin() (interface{}, error)
	Commit() error
	Rollback() error
	BeginSerializableTx() (interface{}, error)
	BeginSerializableTxWithRetry() (interface{}, error)
}
