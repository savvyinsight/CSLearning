package main

import (
	"fmt"
)

func AnonymousFunc() {
	//immediate invocation
	func() {
		fmt.Println("Inside anonymous function1111")
	}()

	// Invoking an Anonymous Function by Name
	v := func() {
		fmt.Println("Inside anonymous function2222")
	}

	t := func(n int) {
		fmt.Println(n)
	}
	v()
	t(10)

	// passing arguments
	func(v int) {
		fmt.Println(v)
	}(23)

	// passing anonymous function as argument
	d := func(val interface{}) {
		fmt.Println(val)
	}

	func(s string, k func(v interface{})) {
		k(s)
	}("good", d)
}
func Closure() {
	count := func() func() int {
		i := 0
		return func() int {
			i++
			return i
		}
	}()
	// count is a `function` that captures and maintains i , count()() first maintain i
	// Each call to count() access the same i variable
	fmt.Println(count())
	fmt.Println(count())
}

func add(a *int) int {
	(*a)++
	return *a
}

func isClosure() {
	count := func() int {
		i := 0
		return add(&i)
	}() // -->immediately called

	// count is an integer(value 1), not a function
	// No state is preserved between calls

	fmt.Println(count)
	fmt.Println(count)

	// IF count is a function
	count_1 := func() int {
		j := 0
		return add(&j)
	} // count is a FUNCTION, no parentheses.

	// but every time call this funcion, it still can't preserve state, because not in same
	// memory address , will create new stack every call.
	// in closure, count function will preserve state, not inner funcion, inner funcion only uses
	// outside preserved value and return
	// this happend in same heap
	fmt.Println(count_1())
	fmt.Println(count_1()) // still 1

	// Closure with Pointer
	k := 0
	count_2 := func() int {
		return add(&k)
	}
	fmt.Println(count_2())
	fmt.Println(count_2())
}

type Student struct {
	name string
	age  int
}

func updateAgeRef(s *Student) {
	(*s).age++ // or s.age++
	s.age++
	fmt.Println("Inside function, updated age =", s.age)
}

func updateAgeVal(s Student) {
	s.age++
	fmt.Println("Inside function, updated age =", s.age)
}

func CallByValue() {
	stu := Student{name: "Rahul", age: 20}
	updateAgeVal(stu)
	fmt.Println("Outside function, original age =", stu.age)
}

func CallByReference() {
	stu := Student{name: "Rahul", age: 20}
	updateAgeRef(&stu)
	fmt.Println("Outside function, original age =", stu.age)
}

func main() {
}
