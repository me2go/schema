/*
 scheam is a common library, which is designed to reflect a struct type easily
*/
package schema

import "reflect"

type FieldVisitor interface {
	Visit([]*reflect.StructField, *reflect.StructField, *reflect.Value)
}
type MethodVisitor interface {
	Visit(m *reflect.Method, in []reflect.Type, out []reflect.Type)
}

/*
 Fields will visit all the struct fields
*/
func Fields(i interface{}, v FieldVisitor, recursive bool) {

}

/*
 Methods will visit all the struct methods
*/
func Methods(i interface{}, v MethodVisitor) {

}
