package transformer

import (
	"math"
	"time"
)

// 定义了用于初始化字符串转换器的构造函数。我们定义了三个实现转换器接口的结构。

// 此结构表示将数据扩充操作，用于将城市映射到其相应的国家/地区。
type EnrichCountryByCity struct {
	City    string
	Country string
}

// 此结构表示数据扩充操作，用于根据出生日期计算人员的年龄。
type EnrichAgeByDateOfBirth struct {
	DateOfBirth time.Time
	Age         int
}

// 此结构表示数据扩充操作，用于根据一个人的身高和体重计算一个人的 BMI（身体质量指数）。
type EnrichBMIByHeightAndWeight struct {
	Height float64
	Weight float64
	BMI    float64
}

func NewEnrichCountryByCity(city string) *EnrichCountryByCity {
	return &EnrichCountryByCity{City: city}
}

func NewEnrichAgeByDateOfBirth(dateOfBirth time.Time) *EnrichAgeByDateOfBirth {
	return &EnrichAgeByDateOfBirth{DateOfBirth: dateOfBirth}
}

func NewEnrichBMIByHeightAndWeight(height, weight float64) *EnrichBMIByHeightAndWeight {
	return &EnrichBMIByHeightAndWeight{Height: height, Weight: weight}
}

func (t *EnrichCountryByCity) Transform() {
	var countries map[string]string
	countries = make(map[string]string)

	countries["LONDON"] = "UK"

	t.Country = countries[t.City]
}

func (t *EnrichAgeByDateOfBirth) Transform() {
	currentDate := time.Now()
	birthDate := t.DateOfBirth

	// Calculate age based on the difference in years
	age := currentDate.Year() - birthDate.Year()

	// Check if the birth date for this year has occurred or not
	if birthDate.YearDay() > currentDate.YearDay() {
		age--
	}

	t.Age = age
}

func (t *EnrichBMIByHeightAndWeight) Transform() {
	t.BMI = math.Round(t.Weight / (t.Height * t.Height))
}
