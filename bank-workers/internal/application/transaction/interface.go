package transaction

type ProcessTransaction interface {
	Handle(dto ProcessDTO) error
}
