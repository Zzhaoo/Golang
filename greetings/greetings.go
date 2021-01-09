package greetings

import (
	"fmt"
)

// Hello returns a greeting for the named person.
func Hello(name string) string {
	message := fmt.Sprintf("Hi, %v. welcome!", name) //%v是指value，就是输出本体，还有%f，%d，%p(指针)等
	return message
}
