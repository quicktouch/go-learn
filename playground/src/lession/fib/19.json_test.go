package fib

import (
	"encoding/json"
	"fmt"
	"testing"
)

/**
常见的json解析

内置的json解析
- 利用反射实现，通过FieldTag来标识对应的json值

因为利用了反射，性能不行，一般可用于配置文件的解析
*/

type BasicInfo struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

type JobInfo struct {
	Skills []string `json:"skills"`
}

type EmployeeStruct struct {
	BasicInfo BasicInfo `json:"basic_info"`
	JobInfo   JobInfo   `json:"job_info"`
}

var jsonStr = `{
	"basic_info": {
		"name":"Mike",
		"age":30
	},
	"job_info": {
		"skills":["Java","Go","C"]
	}
}`

func TestEmbeddedJson(t *testing.T) {
	e := new(EmployeeStruct)
	// json 赋值给struct
	err := json.Unmarshal([]byte(jsonStr), e)
	if err != nil {
		t.Error(err) // {{Mike 30} {[Java Go C]}}
	}
	fmt.Println(*e)
	// struct反序列化
	if v, err := json.Marshal(e); err == nil {
		fmt.Println(string(v))
		// {"basic_info":{"name":"Mike","age":30},"job_info":{"skills":["Java","Go","C"]}}
	} else {
		t.Error(err)
	}
}

/*
EasyJson
采用代码生成而非反射
安装： ` go get -u github.com/mailru/easyjson/...`
*/
