package polymorphsim_test

import (
	"fmt"
	"testing"
)

type Code string
type Progarmmer interface {
	WriteHelloWorld() Code
}

type GoProgarmmer struct {
}

func (goo *GoProgarmmer) WriteHelloWorld() Code {
	return "fmt.Println(\"Hello World\")"
}

type JavaProgarmmer struct {
}

func (goo *JavaProgarmmer) WriteHelloWorld() Code {
	return "System.out.Println(\"Hello World\")"
}

func writeFirstProgram(p Progarmmer) {
	fmt.Printf("%T %v\n", p, p.WriteHelloWorld())
}

func TestClient(t *testing.T) {
	goProg := new(GoProgarmmer)
	javaProg := new(JavaProgarmmer)
	writeFirstProgram(goProg)
	writeFirstProgram(javaProg)
}
