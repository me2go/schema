package schema

import (
	"fmt"
	"log"
	"testing"
)

type testStruct struct {
	a int
	b int
}

func (ts *testStruct) FuncA() {
	fmt.Printf("hello")
}
func (ts testStruct) FuncB() {
	fmt.Printf("hello")
}
func TestFrom(t *testing.T) {
	i := testStruct{}
	sm := From(&i)
	log.Printf("%v", sm)
	if sm.Name != "schema.testStruct" {
		t.Fail()
	}
	if len(sm.Methods) != 2 {
		t.Fail()
	}
	if len(sm.Fields) != 2 {
		t.Fail()
	}
}
