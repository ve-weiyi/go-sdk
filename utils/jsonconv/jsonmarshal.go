package jsonconv

import (
	"bytes"
	"regexp"
)

// https://www.cnblogs.com/chenqionghe/p/13067596.html
/*************************************** 下划线json ***************************************/
type JsonSnakeCase struct {
	Value interface{}
}

func (c JsonSnakeCase) MarshalJSON() ([]byte, error) {
	// Regexp definitions
	var keyMatchRegex = regexp.MustCompile(`\"(\w+)\":`)
	var wordBarrierRegex = regexp.MustCompile(`(\w)([A-Z])`)
	marshalled, err := jjson.Marshal(c.Value)
	converted := keyMatchRegex.ReplaceAllFunc(
		marshalled,
		func(match []byte) []byte {
			return bytes.ToLower(wordBarrierRegex.ReplaceAll(
				match,
				[]byte(`${1}_${2}`),
			))
		},
	)
	return converted, err
}

/*************************************** 驼峰json ***************************************/
type JsonCamelCase struct {
	Value interface{}
}

func (c JsonCamelCase) MarshalJSON() ([]byte, error) {
	var keyMatchRegex = regexp.MustCompile(`\"(\w+)\":`)
	marshalled, err := jjson.Marshal(c.Value)
	converted := keyMatchRegex.ReplaceAllFunc(
		marshalled,
		func(match []byte) []byte {
			matchStr := string(match)
			key := matchStr[1 : len(matchStr)-2]
			resKey := Lcfirst(Case2Camel(key))
			return []byte(`"` + resKey + `":`)
		},
	)
	return converted, err
}

/*************************************** 其他方法 ***************************************/
