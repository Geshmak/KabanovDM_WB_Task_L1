package main

import "fmt"

//foo не удовлетворяет интерфейсу, а адаптер да
//  не всегда будет возможность менять foo, тогда такой подход решит проблему

type foobarer interface {
	foobar()
}

type bar struct {
	barflo float32
	barstr string
}

type foo struct {
	fooint int
	foostr string
}

type adapter struct {
	adapt foo
}

func (a adapter) foobar() {
	a.adapt.foostring()
}
func (b bar) foobar() {
	fmt.Printf("bar: %f, %s\n", b.barflo, b.barstr)
}
func (f foo) foostring() {
	fmt.Printf("foo: %d, %s\n", f.fooint, f.foostr)
}
func ShowYourSelf(x foobarer) {
	x.foobar()
}

func main() {
	b := bar{23.4, "dkab_one_check"}
	f := foo{57, "dkab_two_check"}

	adaptf := adapter{f}

	ShowYourSelf(b)
	ShowYourSelf(adaptf)
}
