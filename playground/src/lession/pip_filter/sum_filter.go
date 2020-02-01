package pipfilter

import "errors"

var SumFilterWrongFormatError = errors.New("SumFilterWrongFormatError")

type SumFilter struct {
}

func NewSumFilter() *SumFilter {
	return &SumFilter{}
}

func (sf *SumFilter) Process(data Request) (Response, error) {
	items, ok := data.([]int)
	if !ok {
		return nil, SumFilterWrongFormatError
	}
	ret := 0
	for _, elem := range items {
		ret += elem
	}
	return ret, nil
}
