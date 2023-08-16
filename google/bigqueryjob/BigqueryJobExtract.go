package bigqueryjob


type BigqueryJobExtract struct {
	// A list of fully-qualified Google Cloud Storage URIs where the extracted table should be written.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/bigquery_job#destination_uris BigqueryJob#destination_uris}
	DestinationUris *[]*string `field:"required" json:"destinationUris" yaml:"destinationUris"`
	// The compression type to use for exported files.
	//
	// Possible values include GZIP, DEFLATE, SNAPPY, and NONE.
	// The default value is NONE. DEFLATE and SNAPPY are only supported for Avro.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/bigquery_job#compression BigqueryJob#compression}
	Compression *string `field:"optional" json:"compression" yaml:"compression"`
	// The exported file format.
	//
	// Possible values include CSV, NEWLINE_DELIMITED_JSON and AVRO for tables and SAVED_MODEL for models.
	// The default value for tables is CSV. Tables with nested or repeated fields cannot be exported as CSV.
	// The default value for models is SAVED_MODEL.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/bigquery_job#destination_format BigqueryJob#destination_format}
	DestinationFormat *string `field:"optional" json:"destinationFormat" yaml:"destinationFormat"`
	// When extracting data in CSV format, this defines the delimiter to use between fields in the exported data.
	//
	// Default is ','
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/bigquery_job#field_delimiter BigqueryJob#field_delimiter}
	FieldDelimiter *string `field:"optional" json:"fieldDelimiter" yaml:"fieldDelimiter"`
	// Whether to print out a header row in the results. Default is true.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/bigquery_job#print_header BigqueryJob#print_header}
	PrintHeader interface{} `field:"optional" json:"printHeader" yaml:"printHeader"`
	// source_model block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/bigquery_job#source_model BigqueryJob#source_model}
	SourceModel *BigqueryJobExtractSourceModel `field:"optional" json:"sourceModel" yaml:"sourceModel"`
	// source_table block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/bigquery_job#source_table BigqueryJob#source_table}
	SourceTable *BigqueryJobExtractSourceTable `field:"optional" json:"sourceTable" yaml:"sourceTable"`
	// Whether to use logical types when extracting to AVRO format.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/bigquery_job#use_avro_logical_types BigqueryJob#use_avro_logical_types}
	UseAvroLogicalTypes interface{} `field:"optional" json:"useAvroLogicalTypes" yaml:"useAvroLogicalTypes"`
}

