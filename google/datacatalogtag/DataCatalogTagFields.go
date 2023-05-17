package datacatalogtag


type DataCatalogTagFields struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/data_catalog_tag#field_name DataCatalogTag#field_name}.
	FieldName *string `field:"required" json:"fieldName" yaml:"fieldName"`
	// Holds the value for a tag field with boolean type.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/data_catalog_tag#bool_value DataCatalogTag#bool_value}
	BoolValue interface{} `field:"optional" json:"boolValue" yaml:"boolValue"`
	// Holds the value for a tag field with double type.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/data_catalog_tag#double_value DataCatalogTag#double_value}
	DoubleValue *float64 `field:"optional" json:"doubleValue" yaml:"doubleValue"`
	// The display name of the enum value.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/data_catalog_tag#enum_value DataCatalogTag#enum_value}
	EnumValue *string `field:"optional" json:"enumValue" yaml:"enumValue"`
	// Holds the value for a tag field with string type.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/data_catalog_tag#string_value DataCatalogTag#string_value}
	StringValue *string `field:"optional" json:"stringValue" yaml:"stringValue"`
	// Holds the value for a tag field with timestamp type.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/data_catalog_tag#timestamp_value DataCatalogTag#timestamp_value}
	TimestampValue *string `field:"optional" json:"timestampValue" yaml:"timestampValue"`
}

