package usa

import (
	"fmt"
	"testing"
)

func testAssert(t *testing.T, b bool, objs ...interface{}) {
	if !b {
		t.Fatal(objs...)
	}
}

func TestFilter(t *testing.T) {
	seed := func(_, out Ch) {
		for i := 0; i < 3; i++ {
			out <- i * 10
		}
		close(out)
	}
	u := Filter(seed).Run()
	testAssert(t, (<-u) == 0, "Could not assert that 1st value has not return.")
	testAssert(t, (<-u) == 10)
	testAssert(t, (<-u) == 20, "Could not assert that last value has not return.")
}

func TestToArray(t *testing.T) {
	seed := func(_, out Ch) {
		for i := 0; i < 3; i++ {
			out <- i * 10
		}
		close(out)
	}
	arr := Filter(seed).Run().ToArray()
	testAssert(t, arr[0] == 0, "Could not assert that Usa get by array.")
	testAssert(t, arr[1] == 10)
	testAssert(t, arr[2] == 20)
}

func TestFilterPipe(t *testing.T) {
	seed := func(_, out Ch) {
		for i := 1; i < 5; i++ {
			out <- i
		}
		close(out)
	}
	isEven := func(i int) bool { return i%2 == 0 }

	arr := Filter(seed).Pipe(isEven).Run().ToArray()
	testAssert(t, arr[0] == false, "Could not assert that Usa get by array.")
	testAssert(t, arr[1] == true)
	testAssert(t, arr[2] == false)
	testAssert(t, arr[3] == true)
}

func TestFilteirFilter(t *testing.T) {
	seed := func(_, out Ch) {
		for i := 1; i < 5; i++ {
			out <- i
		}
		close(out)
	}
	isEven := func(in, out Ch) {
		for i := range in {
			if i.(int)%2 == 0 {
				out <- i
			}
		}
		close(out)
	}

	arr := Filter(seed).Filter(isEven).Run().ToArray()
	testAssert(t, arr[0] == 2, "Could not pass data to filter.")
	testAssert(t, arr[1] == 4)
}

func TestBrakeStream(t *testing.T) {
	seed := func(_, out Ch) {
		for i := 1; i < 100; i++ {
			out <- i
		}
		close(out)
	}
	getHead3 := func(in, out Ch) {
    cnt := 0
    for i := range in{
      out <- i
      cnt += 1
      if cnt > 2{
        break
      }
		}
		close(out)
	}

	arr := Filter(seed).Filter(getHead3).Run().ToArray()
  fmt.Println(arr)
	testAssert(t, len(arr) == 3, "Could not break pipe stream  when upstream is living.")
}

func Example() {
	seq5 := func(_, out Ch) {
		for n := 0; n < 5; n++ {
			out <- n + 1
		}
		close(out)
	}
	add10 := func(i int) int { return i + 10 }
	puts := func(n int) bool {
		fmt.Println(n)
		return true
	}

	Filter(seq5).Pipe(add10).Pipe(puts).Run().Wait()
	// Output:
	// 11
	// 12
	// 13
	// 14
	// 15
}
