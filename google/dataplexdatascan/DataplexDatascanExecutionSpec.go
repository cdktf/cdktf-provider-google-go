package dataplexdatascan


type DataplexDatascanExecutionSpec struct {
	// trigger block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/dataplex_datascan#trigger DataplexDatascan#trigger}
	Trigger *DataplexDatascanExecutionSpecTrigger `field:"required" json:"trigger" yaml:"trigger"`
	// The unnested field (of type Date or Timestamp) that contains values which monotonically increase over time.
	//
	// If not specified, a data scan will run for all data in the table.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/dataplex_datascan#field DataplexDatascan#field}
	Field *string `field:"optional" json:"field" yaml:"field"`
}

