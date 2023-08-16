package documentaiwarehousedocumentschema


type DocumentAiWarehouseDocumentSchemaPropertyDefinitions struct {
	// The name of the metadata property.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/document_ai_warehouse_document_schema#name DocumentAiWarehouseDocumentSchema#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// date_time_type_options block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/document_ai_warehouse_document_schema#date_time_type_options DocumentAiWarehouseDocumentSchema#date_time_type_options}
	DateTimeTypeOptions *DocumentAiWarehouseDocumentSchemaPropertyDefinitionsDateTimeTypeOptions `field:"optional" json:"dateTimeTypeOptions" yaml:"dateTimeTypeOptions"`
	// The display-name for the property, used for front-end.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/document_ai_warehouse_document_schema#display_name DocumentAiWarehouseDocumentSchema#display_name}
	DisplayName *string `field:"optional" json:"displayName" yaml:"displayName"`
	// enum_type_options block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/document_ai_warehouse_document_schema#enum_type_options DocumentAiWarehouseDocumentSchema#enum_type_options}
	EnumTypeOptions *DocumentAiWarehouseDocumentSchemaPropertyDefinitionsEnumTypeOptions `field:"optional" json:"enumTypeOptions" yaml:"enumTypeOptions"`
	// float_type_options block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/document_ai_warehouse_document_schema#float_type_options DocumentAiWarehouseDocumentSchema#float_type_options}
	FloatTypeOptions *DocumentAiWarehouseDocumentSchemaPropertyDefinitionsFloatTypeOptions `field:"optional" json:"floatTypeOptions" yaml:"floatTypeOptions"`
	// integer_type_options block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/document_ai_warehouse_document_schema#integer_type_options DocumentAiWarehouseDocumentSchema#integer_type_options}
	IntegerTypeOptions *DocumentAiWarehouseDocumentSchemaPropertyDefinitionsIntegerTypeOptions `field:"optional" json:"integerTypeOptions" yaml:"integerTypeOptions"`
	// Whether the property can be filtered. If this is a sub-property, all the parent properties must be marked filterable.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/document_ai_warehouse_document_schema#is_filterable DocumentAiWarehouseDocumentSchema#is_filterable}
	IsFilterable interface{} `field:"optional" json:"isFilterable" yaml:"isFilterable"`
	// Whether the property is user supplied metadata.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/document_ai_warehouse_document_schema#is_metadata DocumentAiWarehouseDocumentSchema#is_metadata}
	IsMetadata interface{} `field:"optional" json:"isMetadata" yaml:"isMetadata"`
	// Whether the property can have multiple values.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/document_ai_warehouse_document_schema#is_repeatable DocumentAiWarehouseDocumentSchema#is_repeatable}
	IsRepeatable interface{} `field:"optional" json:"isRepeatable" yaml:"isRepeatable"`
	// Whether the property is mandatory.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/document_ai_warehouse_document_schema#is_required DocumentAiWarehouseDocumentSchema#is_required}
	IsRequired interface{} `field:"optional" json:"isRequired" yaml:"isRequired"`
	// Indicates that the property should be included in a global search.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/document_ai_warehouse_document_schema#is_searchable DocumentAiWarehouseDocumentSchema#is_searchable}
	IsSearchable interface{} `field:"optional" json:"isSearchable" yaml:"isSearchable"`
	// map_type_options block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/document_ai_warehouse_document_schema#map_type_options DocumentAiWarehouseDocumentSchema#map_type_options}
	MapTypeOptions *DocumentAiWarehouseDocumentSchemaPropertyDefinitionsMapTypeOptions `field:"optional" json:"mapTypeOptions" yaml:"mapTypeOptions"`
	// property_type_options block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/document_ai_warehouse_document_schema#property_type_options DocumentAiWarehouseDocumentSchema#property_type_options}
	PropertyTypeOptions *DocumentAiWarehouseDocumentSchemaPropertyDefinitionsPropertyTypeOptions `field:"optional" json:"propertyTypeOptions" yaml:"propertyTypeOptions"`
	// Stores the retrieval importance. Possible values: ["HIGHEST", "HIGHER", "HIGH", "MEDIUM", "LOW", "LOWEST"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/document_ai_warehouse_document_schema#retrieval_importance DocumentAiWarehouseDocumentSchema#retrieval_importance}
	RetrievalImportance *string `field:"optional" json:"retrievalImportance" yaml:"retrievalImportance"`
	// schema_sources block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/document_ai_warehouse_document_schema#schema_sources DocumentAiWarehouseDocumentSchema#schema_sources}
	SchemaSources interface{} `field:"optional" json:"schemaSources" yaml:"schemaSources"`
	// text_type_options block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/document_ai_warehouse_document_schema#text_type_options DocumentAiWarehouseDocumentSchema#text_type_options}
	TextTypeOptions *DocumentAiWarehouseDocumentSchemaPropertyDefinitionsTextTypeOptions `field:"optional" json:"textTypeOptions" yaml:"textTypeOptions"`
	// timestamp_type_options block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/document_ai_warehouse_document_schema#timestamp_type_options DocumentAiWarehouseDocumentSchema#timestamp_type_options}
	TimestampTypeOptions *DocumentAiWarehouseDocumentSchemaPropertyDefinitionsTimestampTypeOptions `field:"optional" json:"timestampTypeOptions" yaml:"timestampTypeOptions"`
}

