package tmpl

const NotEditMark = `
// Code generated by github.com/ve-weiyi/ve-admin-store/server
// Code generated by github.com/ve-weiyi/ve-admin-store/server
// Code generated by github.com/ve-weiyi/ve-admin-store/server
`

const Header = NotEditMark + `
package {{.Package}}

import(	
	{{range .ImportPkgPaths}}{{.}}` + "\n" + `{{end}}
)
`