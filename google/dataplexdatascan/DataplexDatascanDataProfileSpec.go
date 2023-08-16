package dataplexdatascan


type DataplexDatascanDataProfileSpec struct {
	// A filter applied to all rows in a single DataScan job.
	//
	// The filter needs to be a valid SQL expression for a WHERE clause in BigQuery standard SQL syntax. Example: col1 >= 0 AND col2 < 10
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/dataplex_datascan#row_filter DataplexDatascan#row_filter}
	RowFilter *string `field:"optional" json:"rowFilter" yaml:"rowFilter"`
	// The percentage of the records to be selected from the dataset for DataScan.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/dataplex_datascan#sampling_percent DataplexDatascan#sampling_percent}
	SamplingPercent *float64 `field:"optional" json:"samplingPercent" yaml:"samplingPercent"`
}

