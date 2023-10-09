// 数据文件夹是我们从源中提取数据后数据的表示形式。
// 例如，它可以是由生产者推送到 rabbitMQ 上的队列并由我们的代码使用的记录

package data

import "time"

// 我们定义了两个结构，Person 和 PersonTreated。前者代表一个人的基本信息，
// 后者在经过我们的管道处理和丰富后，会保存一个人的信息。
type Person struct {
	Name        string // 名
	Surname     string // 姓
	City        string
	DateOfBirth time.Time
	Height      float64
	Weight      float64
}

type PersonTreated struct {
	completeName string
	country      string
	age          int
	bmi          float64
}

// 定义了两个构造函数来初始化我们的结构
func NewPerson(name, surname, city string, dateOfBirth time.Time, height, weight float64) Person {
	return Person{
		Name:        name,
		Surname:     surname,
		City:        city,
		DateOfBirth: dateOfBirth,
		Height:      height,
		Weight:      weight,
	}
}

func NewPersonTreated(completeName, country string, age int, bmi float64) PersonTreated {
	return PersonTreated{
		completeName: completeName,
		country:      country,
		age:          age,
		bmi:          bmi,
	}
}
