package bigqueryjob


type BigqueryJobLoadParquetOptions struct {
	// If sourceFormat is set to PARQUET, indicates whether to use schema inference specifically for Parquet LIST logical type.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/bigquery_job#enable_list_inference BigqueryJob#enable_list_inference}
	EnableListInference interface{} `field:"optional" json:"enableListInference" yaml:"enableListInference"`
	// If sourceFormat is set to PARQUET, indicates whether to infer Parquet ENUM logical type as STRING instead of BYTES by default.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/bigquery_job#enum_as_string BigqueryJob#enum_as_string}
	EnumAsString interface{} `field:"optional" json:"enumAsString" yaml:"enumAsString"`
}

