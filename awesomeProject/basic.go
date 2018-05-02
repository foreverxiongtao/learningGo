package main

import (
	"fmt"
)

func main() {

	//testMap()
	//i, s := testFunc(10, 20)
	//fmt.Println("result is ",i,"status is ",s)
	//testStruct()
	testMethod()
}

func slice() {
	var slice []int //创建了一个空的切片  也可以通过make函数去解决这个问题
	for i := 0; i < 5; i++ {
		slice = append(slice, i)
	}

	fmt.Println("slice length is ", len(slice))
}

func testMap() {
	var cityMap map[string]string     //申明了变量，但是为初始化
	cityMap = make(map[string]string) //使用make初始化
	cityMap["city1"] = "beijing"
	for key, value := range cityMap {
		fmt.Println("key :", key, "value:", value)
	}
}

func testFunc(param1 int, param2 int) (int, string) {
	return param1 + param2, "success"
}

type person struct {
	name string
	age  int
}

type Rectant struct {
	width  float32
	height float32
}

//带有接受者的方法，也就是对象方法
func (rectant Rectant) area() float32 {
	return rectant.width * rectant.height
}

func testStruct() {
	var p person //声明了一个struct类型的变量
	p.age = 10;
	p.name = "xt"
	fmt.Println("name is ", p.name, "age is ", p.age)
}

func testMethod() {
	var rec Rectant
	rec.width = 100
	rec.height = 2
	fmt.Println("the Rectant area is ", rec.area())
}

type Human struct {
	name string
	age  int
}

type Person struct {
	humen  Human
	height float32
}

func (human Human) sayHello() {
	fmt.Println("name :", human.name)
}

func (person Person) sayHello() {
	fmt.Println("height:", person.height)
}

func testSayHello()  {
	var person Person
}