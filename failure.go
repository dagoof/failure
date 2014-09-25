// Package failure helps avoid tedious if err != nil chains; and facilitates
// resource cleanup.
package failure

// Failure is an internal error wrapping that is checked for when recovering.
type Failure error

// Recover from a possible failure and update the given error with whatever went
// wrong.
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

// Fail ends execution of the current goroutine if the error is not nil.
func Fail(err error) {
	if err != nil {
		panic(Failure(err))
	}
}

// FailFunc ends execution of the current goroutine if the error is not nil, but
// additionally calls any number of empty functions before exiting.
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

// FailErrorFunc is similar to FailFunc but accepts error-returning funcs. Their
// output is ignored, allowing for things like logging or database transaction
// rollbacks.
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
