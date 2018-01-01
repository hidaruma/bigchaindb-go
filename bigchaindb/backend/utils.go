package backend

type Iterator struct {
	max int
	current int
	err error
}

func (i *Iterator) Next() bool {
	if i.err != nil {
		return false
	}
	i.current++
	return i.current <= i.max
}

func (i *Iterator) Value() int {
	if i.err != nil || i.current > i.max {
		panic()
	}
	return i.current
}
type Exception error

type ModuleDispatchRegistrationError Exception

func ModuleDispatchRegistrar() {

}