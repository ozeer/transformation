package transformer

import "strings"

// 我们定义了两个实现转换器接口的结构：ToUpper、JoinStrings
type ToUpper struct { // 此结构表示将字符串转换为大写的转换操作。
	Input  string
	Output string
}

type JoinStrings struct { // 此结构表示将字符串切片联接到单个字符串中且它们之间有空格的转换操作。
	Input  []string
	Output string
}

func NewToUpper(input string) *ToUpper {
	return &ToUpper{Input: input}
}

func NewJoinStrings(inputs ...string) *JoinStrings {
	return &JoinStrings{Input: inputs}
}

func (t *ToUpper) Transform() {
	t.Output = strings.ToUpper(t.Input)
}

func (t *JoinStrings) Transform() {
	t.Output = strings.Join(t.Input, " ")
}
