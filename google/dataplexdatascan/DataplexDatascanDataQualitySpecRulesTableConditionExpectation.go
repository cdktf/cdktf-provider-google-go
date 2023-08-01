package dataplexdatascan


type DataplexDatascanDataQualitySpecRulesTableConditionExpectation struct {
	// The SQL expression.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/dataplex_datascan#sql_expression DataplexDatascan#sql_expression}
	SqlExpression *string `field:"required" json:"sqlExpression" yaml:"sqlExpression"`
}

