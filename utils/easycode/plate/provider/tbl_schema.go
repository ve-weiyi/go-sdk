package provider

type Db struct {
	SchemaName string `json:"schemaName" gorm:"column:SCHEMA_NAME"`
}
