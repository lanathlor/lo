package lo

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPartial(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	add := func(x float64, y int) string {
		return strconv.Itoa(int(x) + y)
	}
	f := Partial(add, 5)
	is.Equal("15", f(10))
	is.Equal("0", f(-5))
}

func TestPartial1(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	add := func(x float64, y int) string {
		return strconv.Itoa(int(x) + y)
	}
	f := Partial1(add, 5)
	is.Equal("15", f(10))
	is.Equal("0", f(-5))
}

func TestPartial2(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	add := func(x float64, y int, z int) string {
		return strconv.Itoa(int(x) + y + z)
	}
	f := Partial2(add, 5)
	is.Equal("24", f(10, 9))
	is.Equal("8", f(-5, 8))
}

func TestPartial3(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	add := func(x float64, y int, z int, a float32) string {
		return strconv.Itoa(int(x) + y + z + int(a))
	}
	f := Partial3(add, 5)
	is.Equal("21", f(10, 9, -3))
	is.Equal("15", f(-5, 8, 7))
}

func TestPartial4(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	add := func(x float64, y int, z int, a float32, b int32) string {
		return strconv.Itoa(int(x) + y + z + int(a) + int(b))
	}
	f := Partial4(add, 5)
	is.Equal("21", f(10, 9, -3, 0))
	is.Equal("14", f(-5, 8, 7, -1))
}

func TestPartial5(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	add := func(x float64, y int, z int, a float32, b int32, c int) string {
		return strconv.Itoa(int(x) + y + z + int(a) + int(b) + c)
	}
	f := Partial5(add, 5)
	is.Equal("26", f(10, 9, -3, 0, 5))
	is.Equal("21", f(-5, 8, 7, -1, 7))
}

func sumBy2(x int) int { return x + 2 }
func mulBy3(x int) int { return x * 3 }

func TestCompose(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	sumBy2AndMulBy3 := Compose(mulBy3, sumBy2)
	mulBy3AndSumBy2 := Compose(sumBy2, mulBy3)

	val := 1
	is.Equal(9, sumBy2AndMulBy3(val))
	is.Equal(5, mulBy3AndSumBy2(val))
}

func TestCompose3(t *testing.T) {
	t.Parallel()

	sumBy2MulBy3AndSumBy2 := Compose3(sumBy2, mulBy3, sumBy2)
	mulBy3SumBy2AndMulBy3 := Compose3(mulBy3, sumBy2, mulBy3)

	is := assert.New(t)

	val := 1
	is.Equal(11, sumBy2MulBy3AndSumBy2(val))
	is.Equal(15, mulBy3SumBy2AndMulBy3(val))
}

func TestCompose4(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	// Compose4: (((sumBy2 -> mulBy3) -> sumBy2) -> mulBy3)
	composed := Compose4(mulBy3, sumBy2, mulBy3, sumBy2)
	val := 1
	// (((1+2)=3) *3=9) +2=11, *3=33
	is.Equal(33, composed(val))

	// Compose4: (((mulBy3 -> sumBy2) -> mulBy3) -> sumBy2)
	composed2 := Compose4(sumBy2, mulBy3, sumBy2, mulBy3)
	// (((1*3)=3)+2=5)*3=15, +2=17
	is.Equal(17, composed2(val))
}

func TestCompose5(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	composed := Compose5(sumBy2, mulBy3, sumBy2, mulBy3, sumBy2)
	val := 1
	// ((((1+2)=3)*3=9)+2=11)*3=33, +2=35
	is.Equal(35, composed(val))

	composed2 := Compose5(mulBy3, sumBy2, mulBy3, sumBy2, mulBy3)
	// ((((1*3)=3)+2=5)*3=15)+2=17, *3=51
	is.Equal(51, composed2(val))
}

func TestComposeN(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	val := 1
	// ComposeN applies right-to-left
	composed := ComposeN(sumBy2, mulBy3, sumBy2, mulBy3, sumBy2)
	// (((((1+2)=3)*3=9)+2=11)*3=33)+2=35
	is.Equal(35, composed(val))

	composed2 := ComposeN(mulBy3, sumBy2, mulBy3, sumBy2, mulBy3)
	// (((((1*3)=3)+2=5)*3=15)+2=17)*3=51
	is.Equal(51, composed2(val))

	// ComposeN with no functions should return input
	identity := ComposeN[int]()
	is.Equal(42, identity(42))
}

func TestComposeUnsafe(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	add2 := func(x interface{}) interface{} {
		return x.(int) + 2
	}
	mul3 := func(x interface{}) interface{} {
		return x.(int) * 3
	}
	toString := func(x interface{}) interface{} {
		return strconv.Itoa(x.(int))
	}

	compose := ComposeUnsafe(toString, mul3, add2)
	result := compose(4)
	is.Equal("18", result) // (4 + 2) * 3 = 18, then to string

	compose2 := ComposeUnsafe(toString)
	is.Equal("7", compose2(7))

	compose3 := ComposeUnsafe()
	is.Equal(5, compose3(5))
}

func TestPipe(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	sumBy2AndMulBy3 := Pipe(sumBy2, mulBy3)
	mulBy3AndSumBy2 := Pipe(mulBy3, sumBy2)

	val := 1
	is.Equal(9, sumBy2AndMulBy3(val))
	is.Equal(5, mulBy3AndSumBy2(val))
}

func TestPipe3(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	sumBy2MulBy3AndSumBy2 := Pipe3(sumBy2, mulBy3, sumBy2)
	mulBy3SumBy2AndMulBy3 := Pipe3(mulBy3, sumBy2, mulBy3)

	val := 1
	is.Equal(11, sumBy2MulBy3AndSumBy2(val))
	is.Equal(15, mulBy3SumBy2AndMulBy3(val))
}

func TestPipe4(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	sumBy2MulBy3AndSumBy2 := Pipe4(sumBy2, mulBy3, sumBy2, mulBy3)
	mulBy3SumBy2AndMulBy3 := Pipe4(mulBy3, sumBy2, mulBy3, sumBy2)

	val := 1
	is.Equal(33, sumBy2MulBy3AndSumBy2(val))
	is.Equal(17, mulBy3SumBy2AndMulBy3(val))
}

func TestPipe5(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	sumBy2MulBy3AndSumBy2 := Pipe5(sumBy2, mulBy3, sumBy2, mulBy3, sumBy2)
	mulBy3SumBy2AndMulBy3 := Pipe5(mulBy3, sumBy2, mulBy3, sumBy2, mulBy3)

	val := 1
	is.Equal(35, sumBy2MulBy3AndSumBy2(val))
	is.Equal(51, mulBy3SumBy2AndMulBy3(val))
}

func TestPipeN(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	sumBy2MulBy3AndSumBy2 := PipeN(sumBy2, mulBy3, sumBy2, mulBy3, sumBy2)
	mulBy3SumBy2AndMulBy3 := PipeN(mulBy3, sumBy2, mulBy3, sumBy2, mulBy3)

	val := 1
	is.Equal(35, sumBy2MulBy3AndSumBy2(val))
	is.Equal(51, mulBy3SumBy2AndMulBy3(val))
}

func TestPipeUnsafe(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	// Define AnyFunc functions
	add2 := func(x interface{}) interface{} {
		return x.(int) + 2
	}
	mul3 := func(x interface{}) interface{} {
		return x.(int) * 3
	}
	toString := func(x interface{}) interface{} {
		return strconv.Itoa(x.(int))
	}

	// Compose them using PipeUnsafe
	pipe := PipeUnsafe(add2, mul3, toString)

	result := pipe(4)
	is.Equal("18", result) // (4 + 2) * 3 = 18, then to string

	// Test with only one function
	pipe2 := PipeUnsafe(toString)
	is.Equal("7", pipe2(7))

	// Test with no functions (should return input)
	pipe3 := PipeUnsafe()
	is.Equal(5, pipe3(5))
}
