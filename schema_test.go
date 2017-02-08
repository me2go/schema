package schema

import (
	"fmt"
	"reflect"
	"testing"
)

type MySubStruct struct {
	A int
	B string
}

type myStruct struct {
	A int
	B string
	MySubStruct
}

var ss *StructSchema = &StructSchema{
	C: NewCache(),
	F: &MyFieldVisitor{},
}

type MyFieldVisitor struct {
}

func (mfv *MyFieldVisitor) Skip(f *reflect.StructField) bool {
	return false
}
func (mfv *MyFieldVisitor) Visit(f *reflect.StructField, v *reflect.Value) {
	fmt.Printf("Field: %s\n", f.Name)
	switch v.Kind() {
	case reflect.Int:
		v.SetInt(10)
	case reflect.String:
		v.SetString("hello")
	case reflect.Struct:
		ss.Fields(v.Addr().Interface())
	}
}
func (m *myStruct) Print() string {
	return fmt.Sprintf("%v", reflect.TypeOf(m))
}

func TestFields(t *testing.T) {
	i := &myStruct{}
	ss.Fields(i)
	fmt.Printf("%+v", i)
}
