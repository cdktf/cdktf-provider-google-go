package dataplexdatascan


type DataplexDatascanDataQualitySpecRulesRegexExpectation struct {
	// A regular expression the column value is expected to match.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/dataplex_datascan#regex DataplexDatascan#regex}
	Regex *string `field:"required" json:"regex" yaml:"regex"`
}

