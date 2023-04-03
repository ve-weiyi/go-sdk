package convert

import (
	"fmt"
	"github.com/ve-weiyi/go-sdk/utils/jsonconv"
	"log"
	"path"
	"strings"
	"testing"
)

func TestString(t *testing.T) {
	addImportCode := "hello  github.com/ve-weiyi/ve-admin-store/server/api/v1/test"
	var (
		importAlias   string
		importPackage string
		packageName   string
	)
	// 删除多余的空格
	addImportCode = strings.Join(strings.Fields(addImportCode), " ")
	// 以空格划分
	importArr := strings.Split(addImportCode, " ")
	log.Println("-->", jsonconv.ObjectToJsonIndent(importArr))
	switch len(importArr) {
	case 1:
		importAlias = ""
		importPackage = importArr[0]
		packageName = path.Base(importPackage)
		break
	case 2:
		importAlias = importArr[0]
		importPackage = importArr[1]
		packageName = importAlias
		break
	default:
		break
	}

	log.Println(importAlias, importPackage, packageName)
}

func TestType(t *testing.T) {

	strings := []string{"11.0", "\"22\"", "11"}

	for _, str := range strings {
		result, err := InferType(str)
		if err != nil {
			fmt.Printf("Error: %v\n", err)
		} else {
			fmt.Printf("Result: %v (type: %T)\n", result, result)
		}
	}

}
