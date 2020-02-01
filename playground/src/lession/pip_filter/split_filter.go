package pipfilter

import (
	"errors"
	"strings"
)

var SplitFilterWrongFormatError = errors.New("SplitFilterWrongFormatError")

// 自己组装的struct, 表示分隔符
type SplitFilter struct {
	delimiter string
}

func NewSplitFilter(delimiter string) *SplitFilter {
	return &SplitFilter{delimiter: delimiter}
}

// struct的方法扩展
func (sf *SplitFilter) Process(data Request) (Response, error) {
	str, ok := data.(string) //检查数据格式/类型,是否可以处理
	if !ok {
		return nil, SplitFilterWrongFormatError
	}
	parts := strings.Split(str, sf.delimiter)
	return parts, nil
}
