package plug

type Emitter interface {
	Emit(value any) error
}
