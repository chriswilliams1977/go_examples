#sometimes you dont care the type is
#you just want to know it can do certain things
#example you want to draw but dont care if its and pen or pencil
#interfaces you can use any type as long as they have certain methods

#Set of methods and value is expected to have
#define using interface keyword
#type myInterface interface{
#method()
#}
#Any type with all the same methods satisfies the interface
#that type can then be used anywhere the interface is called
#method name, params, return value all need to match
#the type can have methods in addition to those listed by interface but mustn't miss any
#a type can satisfy multiple interfaces
#an interface can have multiple types that satisfy it

#concrete type specifies what methods you can call and what they are (underlying type e.g int, struct, string..)
#interfaces dont describe what type they are they only describe what they can do

#cannot use type conversions with interfaces
#if you use pointers as reciever in  methods used by a type you must reference pointer if you assign interface to that type
#type MyType string
#func (m *MyType) MyMethod(){}
#
#type MyInterface interface(){
#MyMethod()
#}
#
#m := MyType("this is my type")
#var i MyInterface = &m

#type assertion allows you to get the concrete type back when assigned to a var with an interface
#bascially says i know var used interface but i know the underlying type is actually this concrete type
#then you can call methods on concrete type that are not in the interface
#var i MyInterface = MyType("this is my type")
#var myType MyType = i.(MyType)