package main

import (
	"fmt"
	"os"
)

func Println(str any) {
	fmt.Fprint(os.Stdout, str)
}

func print(str ...interface{}) {
	fmt.Fprintf(os.Stdout, "%v\n", str...)
}

func main() {
	Println("gg")
	Println(23)

	print("")
	print('s')
	print("haha")
	print(42)

}
