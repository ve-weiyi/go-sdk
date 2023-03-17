package jsonconv

import (
	"unicode"
)

const (
	Camel = 0 //驼峰
	Case  = 1 //下划线
)

/**
 * 驼峰式写法转为下划线写法
 * @description xx_yy to XxYx  xx_y_y to XxYY  XxYY to XxYY
 **/
func Camel2Case(XxYY string) string {
	xx_y_y := make([]byte, 0)

	for i, w := range XxYY {
		//如果是大写
		if unicode.IsUpper(w) {

			//非首个字符
			if i != 0 {
				xx_y_y = append(xx_y_y, '_')
			}
			xx_y_y = append(xx_y_y, byte(unicode.ToLower(w)))
		} else {
			xx_y_y = append(xx_y_y, byte(w))
		}
	}
	return string(xx_y_y[:])
}

/**
 * 下划线转驼峰
 * @description xx_yy to XxYx  xx_y_y to XxYY  XxYY to XxYY
 * @date 2023/2/15
 * @param xx_y_y
 * @return xxYY
 **/
func Case2Camel(xx_y_y string) string {
	XxYY := make([]byte, 0, len(xx_y_y))
	//是否遇到下划线,初始化值为true则转换第一个字母
	line := true
	for _, w := range xx_y_y {
		//遇到 _
		if w == '_' {
			line = true
			continue
		}
		//遇到小写
		if w >= 'a' && w <= 'z' {
			if line {
				w = w - 32
			}
		}
		//遇到大写，跳过
		if w >= 'A' && w <= 'Z' {

		}
		//只对 _ 后一个字母生效
		if line {
			line = false
		}
		XxYY = append(XxYY, byte(w))
	}
	return string(XxYY[:])
}

// 首字母大写
func Ucfirst(str string) string {
	for i, v := range str {
		return string(unicode.ToUpper(v)) + str[i+1:]
	}
	return ""
}

// 首字母小写
func Lcfirst(str string) string {
	for i, v := range str {
		return string(unicode.ToLower(v)) + str[i+1:]
	}
	return ""
}
