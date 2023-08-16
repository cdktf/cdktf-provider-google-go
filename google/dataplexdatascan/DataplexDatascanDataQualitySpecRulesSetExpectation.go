package dataplexdatascan


type DataplexDatascanDataQualitySpecRulesSetExpectation struct {
	// Expected values for the column value.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/dataplex_datascan#values DataplexDatascan#values}
	Values *[]*string `field:"required" json:"values" yaml:"values"`
}

