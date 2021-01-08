package dependency_injection

import (
	"fmt"
	"io"
)

func Greet(writer io.Writer, name string) {
	//fmt.Printf("Hello, %s", name) //going to stdout
	fmt.Fprintf(writer, "Hello, %s", name)
}
