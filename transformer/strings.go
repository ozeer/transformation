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

// 构造函数，初始化ToUpper对象
func NewToUpper(input string) *ToUpper {
	return &ToUpper{Input: input}
}

// 构造函数，初始化JoinStrings对象
func NewJoinStrings(inputs ...string) *JoinStrings {
	return &JoinStrings{Input: inputs}
}

// 给对象ToUpper增加方法Transform，作用是转换ToUpper对象中的属性input字符串（输入）转换为大写output（输出）
func (t *ToUpper) Transform() {
	t.Output = strings.ToUpper(t.Input)
}

// 给对象JoinStrings增加方法Transform，作用是转换JoinStrings对象中的属性input字符串数组（输入）采用空格连接为output（输出）
func (t *JoinStrings) Transform() {
	t.Output = strings.Join(t.Input, " ")
}
