package datastreamstream


type DatastreamStreamDestinationConfigGcsDestinationConfigJsonFileFormat struct {
	// Compression of the loaded JSON file. Possible values: ["NO_COMPRESSION", "GZIP"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/datastream_stream#compression DatastreamStream#compression}
	Compression *string `field:"optional" json:"compression" yaml:"compression"`
	// The schema file format along JSON data files. Possible values: ["NO_SCHEMA_FILE", "AVRO_SCHEMA_FILE"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/datastream_stream#schema_file_format DatastreamStream#schema_file_format}
	SchemaFileFormat *string `field:"optional" json:"schemaFileFormat" yaml:"schemaFileFormat"`
}

