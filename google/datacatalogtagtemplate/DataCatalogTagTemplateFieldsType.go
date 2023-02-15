package datacatalogtagtemplate


type DataCatalogTagTemplateFieldsType struct {
	// enum_type block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/data_catalog_tag_template#enum_type DataCatalogTagTemplate#enum_type}
	EnumType *DataCatalogTagTemplateFieldsTypeEnumType `field:"optional" json:"enumType" yaml:"enumType"`
	// Represents primitive types - string, bool etc.
	//
	// Exactly one of 'primitive_type' or 'enum_type' must be set Possible values: ["DOUBLE", "STRING", "BOOL", "TIMESTAMP"]
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/data_catalog_tag_template#primitive_type DataCatalogTagTemplate#primitive_type}
	PrimitiveType *string `field:"optional" json:"primitiveType" yaml:"primitiveType"`
}

