package bigqueryjob


type BigqueryJobLoad struct {
	// destination_table block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/resources/bigquery_job#destination_table BigqueryJob#destination_table}
	DestinationTable *BigqueryJobLoadDestinationTable `field:"required" json:"destinationTable" yaml:"destinationTable"`
	// The fully-qualified URIs that point to your data in Google Cloud.
	//
	// For Google Cloud Storage URIs: Each URI can contain one '\*' wildcard character
	// and it must come after the 'bucket' name. Size limits related to load jobs apply
	// to external data sources. For Google Cloud Bigtable URIs: Exactly one URI can be
	// specified and it has be a fully specified and valid HTTPS URL for a Google Cloud Bigtable table.
	// For Google Cloud Datastore backups: Exactly one URI can be specified. Also, the '\*' wildcard character is not allowed.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/resources/bigquery_job#source_uris BigqueryJob#source_uris}
	SourceUris *[]*string `field:"required" json:"sourceUris" yaml:"sourceUris"`
	// Accept rows that are missing trailing optional columns.
	//
	// The missing values are treated as nulls.
	// If false, records with missing trailing columns are treated as bad records, and if there are too many bad records,
	// an invalid error is returned in the job result. The default value is false. Only applicable to CSV, ignored for other formats.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/resources/bigquery_job#allow_jagged_rows BigqueryJob#allow_jagged_rows}
	AllowJaggedRows interface{} `field:"optional" json:"allowJaggedRows" yaml:"allowJaggedRows"`
	// Indicates if BigQuery should allow quoted data sections that contain newline characters in a CSV file.
	//
	// The default value is false.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/resources/bigquery_job#allow_quoted_newlines BigqueryJob#allow_quoted_newlines}
	AllowQuotedNewlines interface{} `field:"optional" json:"allowQuotedNewlines" yaml:"allowQuotedNewlines"`
	// Indicates if we should automatically infer the options and schema for CSV and JSON sources.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/resources/bigquery_job#autodetect BigqueryJob#autodetect}
	Autodetect interface{} `field:"optional" json:"autodetect" yaml:"autodetect"`
	// Specifies whether the job is allowed to create new tables.
	//
	// The following values are supported:
	// CREATE_IF_NEEDED: If the table does not exist, BigQuery creates the table.
	// CREATE_NEVER: The table must already exist. If it does not, a 'notFound' error is returned in the job result.
	// Creation, truncation and append actions occur as one atomic update upon job completion Default value: "CREATE_IF_NEEDED" Possible values: ["CREATE_IF_NEEDED", "CREATE_NEVER"]
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/resources/bigquery_job#create_disposition BigqueryJob#create_disposition}
	CreateDisposition *string `field:"optional" json:"createDisposition" yaml:"createDisposition"`
	// destination_encryption_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/resources/bigquery_job#destination_encryption_configuration BigqueryJob#destination_encryption_configuration}
	DestinationEncryptionConfiguration *BigqueryJobLoadDestinationEncryptionConfiguration `field:"optional" json:"destinationEncryptionConfiguration" yaml:"destinationEncryptionConfiguration"`
	// The character encoding of the data.
	//
	// The supported values are UTF-8 or ISO-8859-1.
	// The default value is UTF-8. BigQuery decodes the data after the raw, binary data
	// has been split using the values of the quote and fieldDelimiter properties.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/resources/bigquery_job#encoding BigqueryJob#encoding}
	Encoding *string `field:"optional" json:"encoding" yaml:"encoding"`
	// The separator for fields in a CSV file.
	//
	// The separator can be any ISO-8859-1 single-byte character.
	// To use a character in the range 128-255, you must encode the character as UTF8. BigQuery converts
	// the string to ISO-8859-1 encoding, and then uses the first byte of the encoded string to split the
	// data in its raw, binary state. BigQuery also supports the escape sequence "\t" to specify a tab separator.
	// The default value is a comma (',').
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/resources/bigquery_job#field_delimiter BigqueryJob#field_delimiter}
	FieldDelimiter *string `field:"optional" json:"fieldDelimiter" yaml:"fieldDelimiter"`
	// Indicates if BigQuery should allow extra values that are not represented in the table schema.
	//
	// If true, the extra values are ignored. If false, records with extra columns are treated as bad records,
	// and if there are too many bad records, an invalid error is returned in the job result.
	// The default value is false. The sourceFormat property determines what BigQuery treats as an extra value:
	// CSV: Trailing columns
	// JSON: Named values that don't match any column names
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/resources/bigquery_job#ignore_unknown_values BigqueryJob#ignore_unknown_values}
	IgnoreUnknownValues interface{} `field:"optional" json:"ignoreUnknownValues" yaml:"ignoreUnknownValues"`
	// If sourceFormat is set to newline-delimited JSON, indicates whether it should be processed as a JSON variant such as GeoJSON.
	//
	// For a sourceFormat other than JSON, omit this field. If the sourceFormat is newline-delimited JSON: - for newline-delimited
	// GeoJSON: set to GEOJSON.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/resources/bigquery_job#json_extension BigqueryJob#json_extension}
	JsonExtension *string `field:"optional" json:"jsonExtension" yaml:"jsonExtension"`
	// The maximum number of bad records that BigQuery can ignore when running the job.
	//
	// If the number of bad records exceeds this value,
	// an invalid error is returned in the job result. The default value is 0, which requires that all records are valid.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/resources/bigquery_job#max_bad_records BigqueryJob#max_bad_records}
	MaxBadRecords *float64 `field:"optional" json:"maxBadRecords" yaml:"maxBadRecords"`
	// Specifies a string that represents a null value in a CSV file.
	//
	// For example, if you specify "\N", BigQuery interprets "\N" as a null value
	// when loading a CSV file. The default value is the empty string. If you set this property to a custom value, BigQuery throws an error if an
	// empty string is present for all data types except for STRING and BYTE. For STRING and BYTE columns, BigQuery interprets the empty string as
	// an empty value.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/resources/bigquery_job#null_marker BigqueryJob#null_marker}
	NullMarker *string `field:"optional" json:"nullMarker" yaml:"nullMarker"`
	// parquet_options block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/resources/bigquery_job#parquet_options BigqueryJob#parquet_options}
	ParquetOptions *BigqueryJobLoadParquetOptions `field:"optional" json:"parquetOptions" yaml:"parquetOptions"`
	// If sourceFormat is set to "DATASTORE_BACKUP", indicates which entity properties to load into BigQuery from a Cloud Datastore backup.
	//
	// Property names are case sensitive and must be top-level properties. If no properties are specified, BigQuery loads all properties.
	// If any named property isn't found in the Cloud Datastore backup, an invalid error is returned in the job result.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/resources/bigquery_job#projection_fields BigqueryJob#projection_fields}
	ProjectionFields *[]*string `field:"optional" json:"projectionFields" yaml:"projectionFields"`
	// The value that is used to quote data sections in a CSV file.
	//
	// BigQuery converts the string to ISO-8859-1 encoding,
	// and then uses the first byte of the encoded string to split the data in its raw, binary state.
	// The default value is a double-quote ('"'). If your data does not contain quoted sections, set the property value to an empty string.
	// If your data contains quoted newline characters, you must also set the allowQuotedNewlines property to true.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/resources/bigquery_job#quote BigqueryJob#quote}
	Quote *string `field:"optional" json:"quote" yaml:"quote"`
	// Allows the schema of the destination table to be updated as a side effect of the load job if a schema is autodetected or supplied in the job configuration.
	//
	// Schema update options are supported in two cases: when writeDisposition is WRITE_APPEND;
	// when writeDisposition is WRITE_TRUNCATE and the destination table is a partition of a table, specified by partition decorators.
	// For normal tables, WRITE_TRUNCATE will always overwrite the schema. One or more of the following values are specified:
	// ALLOW_FIELD_ADDITION: allow adding a nullable field to the schema.
	// ALLOW_FIELD_RELAXATION: allow relaxing a required field in the original schema to nullable.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/resources/bigquery_job#schema_update_options BigqueryJob#schema_update_options}
	SchemaUpdateOptions *[]*string `field:"optional" json:"schemaUpdateOptions" yaml:"schemaUpdateOptions"`
	// The number of rows at the top of a CSV file that BigQuery will skip when loading the data.
	//
	// The default value is 0. This property is useful if you have header rows in the file that should be skipped.
	// When autodetect is on, the behavior is the following:
	// skipLeadingRows unspecified - Autodetect tries to detect headers in the first row. If they are not detected,
	// the row is read as data. Otherwise data is read starting from the second row.
	// skipLeadingRows is 0 - Instructs autodetect that there are no headers and data should be read starting from the first row.
	// skipLeadingRows = N > 0 - Autodetect skips N-1 rows and tries to detect headers in row N. If headers are not detected,
	// row N is just skipped. Otherwise row N is used to extract column names for the detected schema.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/resources/bigquery_job#skip_leading_rows BigqueryJob#skip_leading_rows}
	SkipLeadingRows *float64 `field:"optional" json:"skipLeadingRows" yaml:"skipLeadingRows"`
	// The format of the data files.
	//
	// For CSV files, specify "CSV". For datastore backups, specify "DATASTORE_BACKUP".
	// For newline-delimited JSON, specify "NEWLINE_DELIMITED_JSON". For Avro, specify "AVRO". For parquet, specify "PARQUET".
	// For orc, specify "ORC". [Beta] For Bigtable, specify "BIGTABLE".
	// The default value is CSV.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/resources/bigquery_job#source_format BigqueryJob#source_format}
	SourceFormat *string `field:"optional" json:"sourceFormat" yaml:"sourceFormat"`
	// time_partitioning block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/resources/bigquery_job#time_partitioning BigqueryJob#time_partitioning}
	TimePartitioning *BigqueryJobLoadTimePartitioning `field:"optional" json:"timePartitioning" yaml:"timePartitioning"`
	// Specifies the action that occurs if the destination table already exists.
	//
	// The following values are supported:
	// WRITE_TRUNCATE: If the table already exists, BigQuery overwrites the table data and uses the schema from the query result.
	// WRITE_APPEND: If the table already exists, BigQuery appends the data to the table.
	// WRITE_EMPTY: If the table already exists and contains data, a 'duplicate' error is returned in the job result.
	// Each action is atomic and only occurs if BigQuery is able to complete the job successfully.
	// Creation, truncation and append actions occur as one atomic update upon job completion. Default value: "WRITE_EMPTY" Possible values: ["WRITE_TRUNCATE", "WRITE_APPEND", "WRITE_EMPTY"]
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/resources/bigquery_job#write_disposition BigqueryJob#write_disposition}
	WriteDisposition *string `field:"optional" json:"writeDisposition" yaml:"writeDisposition"`
}

