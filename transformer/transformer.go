package transformer

// 我们定义了一个实现方法 Transform（） 的接口。这是负责将转换应用于数据的方法。
type Transformer interface {
	Transform()
}
