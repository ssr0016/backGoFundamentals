package utils

import (
	"fmt"
	"reflect"
)

// What is reflection?
// Reflection is the ability of a program to inspect its variables and values at run time and find their type.

type order struct {
	ordId      int
	customerId int
}

type employees struct {
	name    string
	id      int
	address string
	salary  int
	country string
}

func createQuery(o order) string {
	i := fmt.Sprintf("insert into order values(%d,%d)", o.ordId, o.customerId)
	return i
}

func createQuery1(q interface{}) {
	t := reflect.TypeOf(q)
	v := reflect.ValueOf(q)
	fmt.Println("Type", t)
	fmt.Println("Value", v)
}

func ReflectionExample() {
	o := order{
		ordId:      1234,
		customerId: 567,
	}
	fmt.Println(createQuery(o))
}

func ReflectionExample1() {
	o := order{
		ordId:      1234,
		customerId: 567,
	}
	createQuery1(o)
}

type order1 struct {
	ordId      int
	customerId int
}

func createQuery2(q interface{}) {
	if reflect.ValueOf(q).Kind() == reflect.Struct {
		v := reflect.ValueOf(q)
		fmt.Println("Number of fields", v.NumField())
		for i := 0; i < v.NumField(); i++ {
			fmt.Printf("Field: %d type: %T value: %v\n", i, v.Field(i), v.Field(i))
		}

	}
}

func NumFieldAndFieldMethods() {
	o := order1{
		ordId:      1234,
		customerId: 567,
	}
	createQuery2(o)
}

// Int() and String() methods
func IntAndStringMethods() {
	a := 56
	x := reflect.ValueOf(a).Int()
	fmt.Printf("type:%T value:%v\n", x, x)
	b := "JhayEll"
	y := reflect.ValueOf(b).String()
	fmt.Printf("type:%T value:%v\n", y, y)
}

type orderFinal struct {
	ordId      int
	customerId int
}

type employeeFinal struct {
	name    string
	id      int
	address string
	salary  int
	country string
}

func CompleteProgramReflection(q interface{}) {
	if reflect.ValueOf(q).Kind() == reflect.Struct {
		t := reflect.TypeOf(q).Name()
		query := fmt.Sprintf("insert into %s values(", t)
		v := reflect.ValueOf(q)
		for i := 0; i < v.NumField(); i++ {
			switch v.Field(i).Kind() {
			case reflect.Int:
				if i == 0 {
					query = fmt.Sprintf("%s%d", query, v.Field(i).Int())
				} else {
					query = fmt.Sprintf("%s,%d", query, v.Field(i).Int())
				}
			case reflect.String:
				if i == 0 {
					query = fmt.Sprintf("%s\"%s\"", query, v.Field(i).String())
				} else {
					query = fmt.Sprintf("%s,\"%s\"", query, v.Field(i).String())
				}
			default:
				fmt.Printf("Unsupported type\n")
				return
			}
		}
		query = fmt.Sprintf("%s)", query)
		fmt.Println(query)
		return
	}
	fmt.Println("Unsupported type\n")

}

func ReflectionFinalProgram() {
	o := orderFinal{
		ordId:      1234,
		customerId: 567,
	}
	CompleteProgramReflection(o)

	e := employeeFinal{
		name:    "John",
		id:      123,
		address: "123 Main Street",
		salary:  50000,
		country: "USA",
	}

	CompleteProgramReflection(e)
	i := 90
	CompleteProgramReflection(i)
}
