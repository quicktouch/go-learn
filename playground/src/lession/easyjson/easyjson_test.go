package easyjsonttest

import (
	"encoding/json"
	"fmt"
	"testing"
)

var jsonStr = `{
	"basic_info": {
		"name":"Mike",
		"age":30
	},
	"job_info": {
		"skills":["Java","Go","C"]
	}
}`

func TestEasyJson(t *testing.T) {
	e := EmployeeStruct{}
	err := e.UnmarshalJSON([]byte(jsonStr))
	if err != nil {
		t.Error(err)
	} else {
		fmt.Println(e) // {{Mike 30} {[Java Go C]}}
	}

	if v, err := e.MarshalJSON(); err != nil {
		t.Error(err)
	} else {
		fmt.Println(string(v)) // {"basic_info":{"name":"Mike","age":30},"job_info":{"skills":["Java","Go","C"]}}
	}
}

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

// go test -v -cover
// go test -bench=.

func BenchmarkEasyJson(b *testing.B) {
	b.ResetTimer()
	e := EmployeeStruct{}
	for i := 0; i < b.N; i++ {
		err := e.UnmarshalJSON([]byte(jsonStr))
		if err != nil {
			b.Error(err)
		}
		if _, err := e.MarshalJSON(); err != nil {
			b.Error(err)
		}
	}
}

func BenchmarkEmbedJson(b *testing.B) {
	b.ResetTimer()
	e := new(EmployeeStruct)
	for i := 0; i < b.N; i++ {
		// json 赋值给struct
		err := json.Unmarshal([]byte(jsonStr), e)
		if err != nil {
			b.Error(err)
		}
		if _, err := json.Marshal(e); err != nil {
			b.Error(err)
		}
	}
}

// go test -bench=.

//goos: darwin
//goarch: amd64
//BenchmarkEasyJson-8      1262138               936 ns/op (相同任务运行时间少于1/3)
//BenchmarkEmbedJson-8      334260              3494 ns/op
