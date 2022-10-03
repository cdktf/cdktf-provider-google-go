package dataplexzone


type DataplexZoneDiscoverySpecJsonOptions struct {
	// Optional.
	//
	// Whether to disable the inference of data type for Json data. If true, all columns will be registered as their primitive types (strings, number or boolean).
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/dataplex_zone#disable_type_inference DataplexZone#disable_type_inference}
	DisableTypeInference interface{} `field:"optional" json:"disableTypeInference" yaml:"disableTypeInference"`
	// Optional. The character encoding of the data. The default is UTF-8.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/dataplex_zone#encoding DataplexZone#encoding}
	Encoding *string `field:"optional" json:"encoding" yaml:"encoding"`
}

