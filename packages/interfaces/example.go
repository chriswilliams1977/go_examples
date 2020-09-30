package interfaces

import "fmt"

type MyInterface interface{
	MethodWithoutParams()
	MethodWithParams(float64)
	MethodWithReturnValue() string
}

type MyType int

func (m MyType) MethodWithoutParams(){
	fmt.Println("MethodWithoutParams Called")
}

func (m MyType) MethodWithParams(f float64){
	fmt.Println("MethodWithParams Called")
}

func (m MyType) MethodWithReturnValue() string{
	return "MethodWithReturnValue Called"
}

func CallInterfaceMethod(m MyInterface) string{
	return m.MethodWithReturnValue()
}