package easyjsonttest

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

// 生成文件 (struct文件需要在gopath下)

// ~/go/bin/easyjson -all struct_def.go
