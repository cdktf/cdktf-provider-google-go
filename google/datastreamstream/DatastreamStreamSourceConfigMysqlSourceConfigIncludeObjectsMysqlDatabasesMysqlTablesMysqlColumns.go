package datastreamstream


type DatastreamStreamSourceConfigMysqlSourceConfigIncludeObjectsMysqlDatabasesMysqlTablesMysqlColumns struct {
	// Column collation.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/datastream_stream#collation DatastreamStream#collation}
	Collation *string `field:"optional" json:"collation" yaml:"collation"`
	// Column name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/datastream_stream#column DatastreamStream#column}
	Column *string `field:"optional" json:"column" yaml:"column"`
	// The MySQL data type. Full data types list can be found here: https://dev.mysql.com/doc/refman/8.0/en/data-types.html.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/datastream_stream#data_type DatastreamStream#data_type}
	DataType *string `field:"optional" json:"dataType" yaml:"dataType"`
	// Whether or not the column can accept a null value.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/datastream_stream#nullable DatastreamStream#nullable}
	Nullable interface{} `field:"optional" json:"nullable" yaml:"nullable"`
	// The ordinal position of the column in the table.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/datastream_stream#ordinal_position DatastreamStream#ordinal_position}
	OrdinalPosition *float64 `field:"optional" json:"ordinalPosition" yaml:"ordinalPosition"`
	// Whether or not the column represents a primary key.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/datastream_stream#primary_key DatastreamStream#primary_key}
	PrimaryKey interface{} `field:"optional" json:"primaryKey" yaml:"primaryKey"`
}

