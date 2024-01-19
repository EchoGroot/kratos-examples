package stringx

import (
	"sort"
	"strings"

	"github.com/google/uuid"
)

// GeneratePassword24 生成一个24位密码
func GeneratePassword24() string {
	password := uuid.New().String()
	password = strings.ReplaceAll(password, "-", "")
	return password[4:28]
}

// FirstUpper 首字母大写
func FirstUpper(str string) string {
	if len(str) == 0 {
		return ""
	}
	return strings.ToUpper(str[0:1]) + str[1:]
}

// FirstLower 首字母小写
func FirstLower(str string) string {
	if len(str) == 0 {
		return ""
	}
	return strings.ToLower(str[0:1]) + str[1:]
}

// Snake2BigCamel 蛇形转大驼峰
// xx_yy to XxYy
// xx_y_y to XxYY
func Snake2BigCamel(s string) string {
	data := make([]byte, 0, len(s))
	j := false
	k := false
	num := len(s) - 1
	for i := 0; i <= num; i++ {
		d := s[i]
		if !k && d >= 'A' && d <= 'Z' {
			k = true
		}
		if d >= 'a' && d <= 'z' && (j || !k) {
			d = d - 32
			j = false
			k = true
		}
		if k && d == '_' && num > i && s[i+1] >= 'a' && s[i+1] <= 'z' {
			j = true
			continue
		}
		data = append(data, d)
	}
	return string(data[:])
}

// Camel2Snake 驼峰转蛇形 snake string
// XxYy to xx_yy
// XxYY to xx_y_y
// xxYy to xx_yy
func Camel2Snake(s string) string {
	data := make([]byte, 0, len(s)*2)
	j := false
	num := len(s)
	for i := 0; i < num; i++ {
		d := s[i]
		// or通过ASCII码进行大小写的转化
		// 65-90（A-Z），97-122（a-z）
		// 判断如果字母为大写的A-Z就在前面拼接一个_
		if i > 0 && d >= 'A' && d <= 'Z' && j {
			data = append(data, '_')
		}
		if d != '_' {
			j = true
		}
		data = append(data, d)
	}
	// ToLower把大写字母统一转小写
	return strings.ToLower(string(data[:]))
}

// SlicesEqual 判断两个字符串切片是否相等, 顺序可以不一致
// 如下所示：返回true
// slice1 := []string{"apple", "banana", "cherry"}
// slice2 := []string{"cherry", "banana", "apple"}
func SlicesEqual(slice1, slice2 []string) bool {
	if len(slice1) != len(slice2) {
		return false
	}

	// 对切片进行排序
	sort.Strings(slice1)
	sort.Strings(slice2)

	for i := range slice1 {
		if slice1[i] != slice2[i] {
			return false
		}
	}

	return true
}
