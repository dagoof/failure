package failure

type Failure error

func Recover(err *error) {
	v := recover()
	if v == nil {
		return
	}

	switch vt := v.(type) {
	case Failure:
		*err = vt
		return
	}

	panic(v)
}

func Fail(err error) {
	if err != nil {
		panic(Failure(err))
	}
}

func FailFunc(fns ...func()) func(error) {
	return func(err error) {
		if err != nil {

			for _, fn := range fns {
				fn()
			}

			panic(Failure(err))
		}
	}
}

func FailErrorFunc(fns ...func() error) func(error) {
	return func(err error) {
		if err != nil {

			for _, fn := range fns {
				fn()
			}

			panic(Failure(err))
		}
	}
}
