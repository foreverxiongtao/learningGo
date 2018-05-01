package main

import (
	"encoding/json"
	"fmt"
	"github.com/bitly/go-simplejson" // for json get

)

type Student struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
	No   string `json:"no"`
}

type Students struct {
	S []Student
}

func main() {
	var stu Students
	stu.S = append(stu.S, Student{Name: "xt", Age: 10, No: "0001"});
	stu.S = append(stu.S, Student{Name: "tx", Age: 12, No: "0002"});
	result, error := json.Marshal(stu);
	if error != nil {
		fmt.Print("error:", error);
	}
	fmt.Println("json:", string(result));
	parseJson();
}

func parseJson() {
	jsonStr := `{"Name":"xiongtao","Age":12,"No":"001"}`;
	js, error := simplejson.NewJson([]byte(jsonStr));
	if error != nil {
		fmt.Println(error);
	}
	Name := js.Get("Name").MustString();
	Age, _ := js.Get("Age").Int();
	No := js.Get("No").MustString();
	fmt.Println("Name:", Name, "Age:", Age, "No:", No)
}
