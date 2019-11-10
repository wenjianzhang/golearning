package interface__test

import "testing"

type Progarmmer interface {
	WriteHelloWorld() string
}

type GoProgarmmer struct {
}

func (goo *GoProgarmmer) WriteHelloWorld() string {
	return "fmt.Println(\"Hello World\")"
}

func TestClient(t *testing.T) {
	var p Progarmmer
	p = new(GoProgarmmer)
	t.Log(p.WriteHelloWorld())
}
