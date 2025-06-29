package lo

// Partial returns new function that, when called, has its first argument set to the provided value.
func Partial[T1, T2, R any](f func(a T1, b T2) R, arg1 T1) func(T2) R {
	return func(t2 T2) R {
		return f(arg1, t2)
	}
}

// Partial1 returns new function that, when called, has its first argument set to the provided value.
func Partial1[T1, T2, R any](f func(T1, T2) R, arg1 T1) func(T2) R {
	return Partial(f, arg1)
}

// Partial2 returns new function that, when called, has its first argument set to the provided value.
func Partial2[T1, T2, T3, R any](f func(T1, T2, T3) R, arg1 T1) func(T2, T3) R {
	return func(t2 T2, t3 T3) R {
		return f(arg1, t2, t3)
	}
}

// Partial3 returns new function that, when called, has its first argument set to the provided value.
func Partial3[T1, T2, T3, T4, R any](f func(T1, T2, T3, T4) R, arg1 T1) func(T2, T3, T4) R {
	return func(t2 T2, t3 T3, t4 T4) R {
		return f(arg1, t2, t3, t4)
	}
}

// Partial4 returns new function that, when called, has its first argument set to the provided value.
func Partial4[T1, T2, T3, T4, T5, R any](f func(T1, T2, T3, T4, T5) R, arg1 T1) func(T2, T3, T4, T5) R {
	return func(t2 T2, t3 T3, t4 T4, t5 T5) R {
		return f(arg1, t2, t3, t4, t5)
	}
}

// Partial5 returns new function that, when called, has its first argument set to the provided value
func Partial5[T1, T2, T3, T4, T5, T6, R any](f func(T1, T2, T3, T4, T5, T6) R, arg1 T1) func(T2, T3, T4, T5, T6) R {
	return func(t2 T2, t3 T3, t4 T4, t5 T5, t6 T6) R {
		return f(arg1, t2, t3, t4, t5, t6)
	}
}

// Curry2 is a function that takes a function with one parameter and returns a curried version of it.
func Curry2[A, B, R any](f func(A, B) R) func(A) func(B) R {
	return func(a A) func(B) R {
		return func(b B) R {
			return f(a, b)
		}
	}
}

// Curry3 is a function that takes a function with three parameters and returns a curried version of it.
func Curry3[A, B, C, R any](f func(A, B, C) R) func(A) func(B) func(C) R {
	return func(a A) func(B) func(C) R {
		return func(b B) func(C) R {
			return func(c C) R {
				return f(a, b, c)
			}
		}
	}
}

// Curry4 is a function that takes a function with four parameters and returns a curried version of it.
func Curry4[A, B, C, D, R any](f func(A, B, C, D) R) func(A) func(B) func(C) func(D) R {
	return func(a A) func(B) func(C) func(D) R {
		return func(b B) func(C) func(D) R {
			return func(c C) func(D) R {
				return func(d D) R {
					return f(a, b, c, d)
				}
			}
		}
	}
}

// Curry5 is a function that takes a function with five parameters and returns a curried version of it.
func Curry5[A, B, C, D, E, R any](f func(A, B, C, D, E) R) func(A) func(B) func(C) func(D) func(E) R {
	return func(a A) func(B) func(C) func(D) func(E) R {
		return func(b B) func(C) func(D) func(E) R {
			return func(c C) func(D) func(E) R {
				return func(d D) func(E) R {
					return func(e E) R {
						return f(a, b, c, d, e)
					}
				}
			}
		}
	}
}

// Compose returns new function that, when called, returns the result of calling g and then f with the result from g
func Compose[T, U, V any](f func(U) V, g func(T) U) func(T) V {
	return func(t T) V {
		return f(g(t))
	}
}

func Compose3[T1, T2, T3, R any](f func(T3) R, g func(T2) T3, h func(T1) T2) func(T1) R {
	return Compose(f, Compose(g, h))
}

func Compose4[T1, T2, T3, T4, R any](f func(T4) R, g func(T3) T4, h func(T2) T3, i func(T1) T2) func(T1) R {
	return Compose(f, Compose(g, Compose(h, i)))
}

func Compose5[T1, T2, T3, T4, T5, R any](f func(T5) R, g func(T4) T5, h func(T3) T4, i func(T2) T3, j func(T1) T2) func(T1) R {
	return Compose(f, Compose(g, Compose(h, Compose(i, j))))
}

func ComposeN[T any](funcs ...func(T) T) func(T) T {
	return func(value T) T {
		for i := len(funcs) - 1; i >= 0; i-- {
			value = funcs[i](value)
		}
		return value
	}
}

func ComposeUnsafe(funcs ...AnyFunc) func(interface{}) interface{} {
	return func(value interface{}) interface{} {
		for i := len(funcs) - 1; i >= 0; i-- {
			value = funcs[i](value)
		}
		return value
	}
}

// Pipe returns new function that, when called, returns the result of calling f and then g with the result from f
func Pipe[T, U, V any](f func(T) U, g func(U) V) func(T) V {
	return Compose(g, f)
}

func Pipe3[T1, T2, T3, R any](f func(T1) T2, g func(T2) T3, h func(T3) R) func(T1) R {
	return Pipe(f, Pipe(g, h))
}

func Pipe4[T1, T2, T3, T4, R any](f func(T1) T2, g func(T2) T3, h func(T3) T4, i func(T4) R) func(T1) R {
	return Pipe(f, Pipe(g, Pipe(h, i)))
}

func Pipe5[T1, T2, T3, T4, T5, R any](f func(T1) T2, g func(T2) T3, h func(T3) T4, i func(T4) T5, j func(T5) R) func(T1) R {
	return Pipe(f, Pipe(g, Pipe(h, Pipe(i, j))))
}

func PipeN[T any](funcs ...func(T) T) func(T) T {
	return func(value T) T {
		for _, fn := range funcs {
			value = fn(value)
		}
		return value
	}
}

type AnyFunc func(interface{}) interface{}

// PipeUnsafe is a version of Pipe that accepts AnyFunc, allowing for more flexible function signatures.
func PipeUnsafe(funcs ...AnyFunc) func(interface{}) interface{} {
	return func(value interface{}) interface{} {
		for _, fn := range funcs {
			value = fn(value)
		}
		return value
	}
}
