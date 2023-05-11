package plate

import (
	"bytes"
	"errors"
	"github.com/ve-weiyi/go-sdk/utils/file"
	"os"
	"path/filepath"
	"text/template"
)

// 生成目录，文件名，模板内容，填充数据 data
type PlateMeta struct {
	Key          string
	AutoCodePath string //生成的代码路径  blog/api/article.go

	TemplateString string //模版文件内容
	//TemplatePath   string      //模版文件路径   tpl/api.go.tpl
	Data interface{} //填充内容
}

func (tempMeta *PlateMeta) CreateTempFile() error {
	//创建文件夹
	err := os.MkdirAll(filepath.Dir(tempMeta.AutoCodePath), 0755)
	if err != nil {
		return err
	}
	//创建.go文件
	f, err := os.Create(tempMeta.AutoCodePath)
	if err != nil {
		return err
	}
	defer f.Close()

	//解析模板
	temp, err := tempMeta.getTemplate()
	if err != nil {
		return err
	}

	var buf bytes.Buffer
	err = temp.Execute(&buf, tempMeta.Data)
	if err != nil {
		return err
	}

	err = output(tempMeta.AutoCodePath, buf.Bytes())
	if err != nil {
		return err
	}
	return nil
}

func (tempMeta *PlateMeta) RollBack() error {
	err := file.DeLFile(tempMeta.AutoCodePath)
	if err != nil {
		return err
	}
	return nil
}

func (tempMeta *PlateMeta) getTemplate() (*template.Template, error) {
	//if tempMeta.TemplatePath != "" {
	//	//解析模板
	//	temp, err := template.ParseFiles(tempMeta.TemplatePath)
	//	if err != nil {
	//		return nil, err
	//	}
	//	return temp, nil
	//}

	if tempMeta.TemplateString != "" {
		//解析模板
		temp, err := template.New("temp").Parse(tempMeta.TemplateString)
		if err != nil {
			return nil, err
		}
		return temp, nil
	}

	return nil, errors.New("TemplateString or TemplatePath all null ")
}
