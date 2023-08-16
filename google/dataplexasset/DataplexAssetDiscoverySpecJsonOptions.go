package dataplexasset


type DataplexAssetDiscoverySpecJsonOptions struct {
	// Optional.
	//
	// Whether to disable the inference of data type for Json data. If true, all columns will be registered as their primitive types (strings, number or boolean).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/dataplex_asset#disable_type_inference DataplexAsset#disable_type_inference}
	DisableTypeInference interface{} `field:"optional" json:"disableTypeInference" yaml:"disableTypeInference"`
	// Optional. The character encoding of the data. The default is UTF-8.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/dataplex_asset#encoding DataplexAsset#encoding}
	Encoding *string `field:"optional" json:"encoding" yaml:"encoding"`
}

