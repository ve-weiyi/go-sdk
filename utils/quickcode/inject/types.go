package logic

import (
	"bytes"
	"errors"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"text/template"
)

type autoPackage struct {
	path       string
	tempString string
	name       string
}

type injectionMeta struct {
	path       string //文件路径
	funcName   string //插入的位置
	injectCode string //注入的代码
}

type astInjectionMeta struct {
	path           string
	packageAlias   string //导入别名
	packageImport  string
	findStructName string //在 struct findStructName{ //插入 elementAlias  packageAlias.elementName  }
	elementAlias   string
	elementName    string
}

type TempMeta struct {
	//easycode.TplMeta
	//Template         *template.Template
	//Key              string
	TemplateString   string //模版文件内容
	TemplatePath     string //模版文件路径   tpl/api.go.tpl
	AutoPackage      string //包名    blog
	AutoCodePath     string //生成的代码路径  blog/api/article.go
	AutoMoveFilePath string //需要移动 AutoCodePath ->AutoMoveFilePath
}

func (tempMeta *TempMeta) CreateTempFile(tempData interface{}) error {
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
	err = temp.Execute(&buf, tempData)
	if err != nil {
		return err
	}

	err = output(tempMeta.AutoCodePath, buf.Bytes())
	if err != nil {
		return err
	}
	return nil
}

func (tempMeta *TempMeta) MoveTempFile() error {
	//判断目标文件是否都可以移动
	if tempMeta.AutoMoveFilePath != "" {
		if FileExist(tempMeta.AutoMoveFilePath) {
			return errors.New(fmt.Sprintf("目标文件已存在:%s\n", tempMeta.AutoMoveFilePath))
		}

		if err := FileMove(tempMeta.AutoCodePath, tempMeta.AutoMoveFilePath); err != nil {
			return err
		}
		log.Println("file move success:", tempMeta.AutoMoveFilePath)
	}
	return nil
}

func (tempMeta *TempMeta) getTemplate() (*template.Template, error) {
	if tempMeta.TemplatePath != "" {
		//解析模板
		temp, err := template.ParseFiles(tempMeta.TemplatePath)
		if err != nil {
			return nil, err
		}
		return temp, nil
	}

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
