package main

import (
	"bytes"
	"fmt"
	"io"
	"net"
	"os"
	"unsafe"
)

/*
What Makes Go Interfaces Different?

In Go, interfaces are implicit. It means you don’t need to declare that a type
implements an interface explicitly, it just needs the right methods.
This simple idea leads to incredibly flexible and maintainable code.
*/

// ====================== First======================================================
type User struct {
	Name string
	Age  int
}

// String() mthod here satisfies fmt.Stringer implicitly
func (u User) String() string {
	return fmt.Sprintf("%s (%d years old)", u.Name, u.Age)
}

func printUser() {
	user := User{Name: "Alice", Age: 28}
	fmt.Println(user) // Uses fmt.Stringer's String() method
}

// -------------Q&A------------
/*
What is the data flow here? Clarify data flow and what happend behind the scenes when `fmt.Println(user)`
executes?
sequence:
1.fmt.Println() this receives as an interface.
	- fmt.Println(user) this received user(Concrete User)
2.func Println(a ...interface{})
	├─  user gets stored in interface{} (empty interface)
    │  └─ Contains: (type: User, value: {Alice, 28})
3.fmt.Fprintln() / internal printValue()
	- Extract concrete value: value.Interface()
	- Type assertion: value.Interface().(fmt.Stringer)
		- Does User have String()? ✓ YES
		- Create temporary Stringer interface
			- Contains: (type: User, value: {Alice, 28})
4.stringer.String()  // Calls User.String() method
*/

//*******************************
// KEY: How a concrete type gets "wrapped" into an interface, or how Go stores values in interfaces
// internally. An interface value actually a two-word pair value. (On 64bit Systems : 16 bytes total)

// Internal representation of an interface (simplified)
type interface struct {
	type *itab          // pointer to type information + method table
	value unsafe.Pointer // pointer to the actual data
}

/*
"Wrapped into an interface" means:
Go allocates memory for an interface structure (2 words)
It stores a pointer to the type information (including method table)
It copies your concrete value into another memory location
It stores a pointer to that copied value in the interface's data field
The original user variable and the wrapped value inside the interface are `separate memory locations`(unless you wrap a pointer).
This is why interfaces with large structs can have overhead - they copy the entire value!
That's why sometimes people pass pointers to interfaces instead of values.
*/

//**************************************

// ====================== Second======================================================

// function works with any type that has a Read() method (io.Reader)
func processData(r io.Reader) error {
	data := make([]byte, 1024)
	_, err := r.Read(data)
	return err
}

func implmentIoReader() {
	// below types have a Read method, so they implement io.Reader

	// files
	file, _ := os.Open("data.txt")
	processData(file)

	// network connections
	conn, _ := net.Dial("tcp", "example.com:80")
	processData(conn)

	// in-memory buffers
	buf := bytes.NewBuffer([]byte("Hello"))
	processData(buf)
}

// ====================== Thrid======================================================

type DataProcessor interface {
	Process(data []byte) error
}

type DataValidator interface {
	Validate(data []byte) error
}

type DataHandler interface {
	DataProcessor
	DataValidator
}

type JSONHandler struct{}

func (j JSONHandler) Process(data []byte) error {
	fmt.Println("Processing JSON data...")
	return nil
}

func (j JSONHandler) Validate(data []byte) error {
	fmt.Println("Validating JSON format...")
	return nil
}

// ====================== Four======================================================

func printAny(v any) {
	switch t := v.(type) {
	case string:
		fmt.Printf("String %s\n", t)
	case int:
		fmt.Printf("Integer %d\n", t)
	case bool:
		fmt.Printf("Boolean %v\n", t)
	default:
		fmt.Printf("Unknown type %T %v\n", t, t)
	}

	if str, ok := v.(string); ok {
		fmt.Printf("Got a string %s\n", str)
	}
}

func testPrintAny() {
	printAny("Hello")
	printAny(23)
	printAny(true)
	printAny('a')
}

func main() {
	
}
