package pipfilter

// 输入(空接口别名-任意类型)
type Request interface{}

// 输出(空接口别名-任意类型)
type Response interface{}

// 处理默认实现
type Filter interface {
	Process(data Request) (Response, error)
}
