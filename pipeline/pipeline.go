package pipeline

import (
	"fmt"

	"transformation/transformer"
)

// 管道文件夹包含定义新管道、向其添加阶段并最终运行它的逻辑。

// 此结构表示数据管道中的处理阶段。
type stage struct {
	name        string
	transformer transformer.Transformer // transformer ：类型的 transformer.Transformer 字段，指示此阶段将使用转换器应用转换。
}

// 此结构表示数据处理管道。
type pipeline struct {
	name   string
	stages []stage
}

// 是一个构造函数，用于创建 Pipeline 结构的新实例。
// 它需要一个 name 参数来设置管道的名称。
// 它采用一个 name 参数作为阶段的名称和一个参数，该 transformation 参数应是实现接口的 transformer.Transformer 转换器。
// 它将具有指定名称和转换器的新 stage 结构追加到管道的 stages 切片，并返回指向更新管道的指针。
func NewPipeline(name string) *pipeline {
	return &pipeline{
		name:   name,
		stages: []stage{},
	}
}

// AddStage 是 Pipeline 结构体的方法。
// 它允许您向管道添加处理阶段。
func (p *pipeline) AddStage(name string, transformation transformer.Transformer) *pipeline {
	p.stages = append(p.stages, stage{name: name, transformer: transformation})
	return p
}

// Run 是 Pipeline 结构体的方法。
// 它通过循环访问其阶段并应用转换来执行管道。
// 它打印消息，指示管道和每个阶段正在运行。
// 它在每个阶段的转换器上调用 Transform 该方法以执行实际的数据转换。
// 最后，它会打印一条消息，指示管道已完成。
func (p *pipeline) Run() {
	fmt.Println("Running pipeline:", p.name)
	for _, s := range p.stages {
		fmt.Println("Running stage:", s.name)
		s.transformer.Transform()
	}
	fmt.Print("Pipeline finished\n\n")
}
