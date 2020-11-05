/**
* @Author : henry
* @Data: 2020-10-12 14:32
* @Note: 反射的基本用法
**/

package main

import (
	"fmt"
	"reflect"
)

type MyString string

func main() {
	//var Str MyString

	Str := "hello world"

	var rType reflect.Type
	rType = reflect.TypeOf(Str) // 返回类型
	fmt.Printf("reflect.TypeOf(Str): %s\n", rType)
	fmt.Printf("str type:%T\n", Str)

	var rValue reflect.Value
	rValue = reflect.ValueOf(Str)
	fmt.Printf("reflect.ValueOf(Str): %#v\n", rValue)
	t := rValue.Type()
	fmt.Println(t.String())

	var Str1 MyString
	Str1 = "hello world"
	rValue1 := reflect.ValueOf(Str1)
	var rKind reflect.Kind

	rKind = rValue1.Kind() // 返回底层数据类型
	fmt.Printf("reflect.ValueOf(Str1).Type(): %s\n", rValue1.Type())
	fmt.Printf("reflect.ValueOf(Str1).Kind(): %s\n", rKind)

	var rInterface interface{}
	rInterface = reflect.ValueOf(Str).Interface() // 把数据转换为interface{}结构
	fmt.Printf("reflect.ValueOf(Str).Interface: %#v\n", rInterface)

	var rString string
	rString = reflect.ValueOf(Str).String()
	fmt.Printf("reflect.Value(Str).String(): %#v\n", rString)

	reflect.ValueOf(&Str).Elem().SetString(Str)
	fmt.Printf("reflect.ValueOf(Str).SetString(Str): %#v\n", Str)
}
