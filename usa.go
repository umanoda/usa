package usa

import (
	"reflect"
)

// Filter lile process on golang
// v0.4.0

type Ch chan interface{}
type function interface{}

// funcs[0] -> pipes[0] -> funcs[1] -> pipes[1] -..-> pipes[n]
type Usa struct {
	funcs  [](func(Ch, Ch))
	pipes  []Ch
	chSize int
}

func Filter(f func(Ch, Ch)) *Usa {
	u := new(Usa)
	return u.Filter(f)
}

func Pipe(f function) *Usa {
	u := new(Usa)
	return u.Pipe(f)
}

func (u *Usa) Pipe(f function) *Usa {
	fv := reflect.ValueOf(f)
	if fv.Kind() != reflect.Func {
		panic("get none functional argument.")
	}

	u.Filter(
		func(in, out Ch) {
			for val := range in {
				arg := reflect.ValueOf(val)
				ret := fv.Call([]reflect.Value{arg})[0]
				out <- ret.Interface()
			}
			close(out)
		})
	return u
}

func (u *Usa) Filter(f func(Ch, Ch)) *Usa {
	u.funcs = append(u.funcs, f)
	return u
}

func (u *Usa) Run() Ch {
	var in, out Ch
	for i, f := range u.funcs {
		pipe := make(Ch, u.chSize)
		u.pipes = append(u.pipes, pipe)
		if i > 0 {
			in = u.pipes[i-1]
		}
		out = pipe

		go f(in, out)
	}
	return out
}

func (c Ch) Wait(){
  for _ = range c{}
}

func (c Ch) ToArray() []interface{}{
  var arr []interface{}
  for v := range c{
    arr = append(arr, v)
  }
  return arr
}
