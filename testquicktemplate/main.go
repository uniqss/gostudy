//package main
//
//import (
//	"fmt"
//
//	"testquicktemplate/templates"
//)
//
//func main() {
//	fmt.Printf("%s\n", templates.Hello("Foo"))
//	fmt.Printf("%s\n", templates.Hello("Bar"))
//}

package main

import (
	"bytes"
	"fmt"

	"testquicktemplate/templates"
)

func main() {
	names := []string{"Kate", "Go", "John", "Brad"}

	// qtc creates Write* function for each template function.
	// Such functions accept io.Writer as first parameter:
	var buf bytes.Buffer
	var ret []string
	templates.WriteGreetings(&buf, names, &ret)

	fmt.Println(ret)

	//fmt.Printf("buf=\n%s", buf.Bytes())
}