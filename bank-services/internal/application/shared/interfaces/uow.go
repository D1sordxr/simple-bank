package interfaces

type UnitOfWork interface {
	Begin() (interface{}, error)
	Commit() error
	Rollback() error
	BeginSerializableTx() (interface{}, error)
	BeginSerializableTxWithRetry() (interface{}, error)
}
