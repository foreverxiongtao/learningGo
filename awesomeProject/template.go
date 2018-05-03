package main

import (
	"html/template"
	"os"
)

type Dog struct {
	DogName string
}

func main()  {

	mytempate:=template.New("mytemplate")
	mytempate.Parse("hello {{.DogName}}")
	var dog=Dog{DogName:"yellow"}
	mytempate.Execute(os.Stdout,dog)
}
