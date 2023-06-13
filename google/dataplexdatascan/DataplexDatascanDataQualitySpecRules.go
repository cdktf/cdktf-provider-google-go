package dataplexdatascan


type DataplexDatascanDataQualitySpecRules struct {
	// The dimension a rule belongs to.
	//
	// Results are also aggregated at the dimension level. Supported dimensions are ["COMPLETENESS", "ACCURACY", "CONSISTENCY", "VALIDITY", "UNIQUENESS", "INTEGRITY"]
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/resources/dataplex_datascan#dimension DataplexDatascan#dimension}
	Dimension *string `field:"required" json:"dimension" yaml:"dimension"`
	// The unnested column which this rule is evaluated against.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/resources/dataplex_datascan#column DataplexDatascan#column}
	Column *string `field:"optional" json:"column" yaml:"column"`
	// Rows with null values will automatically fail a rule, unless ignoreNull is true.
	//
	// In that case, such null rows are trivially considered passing. Only applicable to ColumnMap rules.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/resources/dataplex_datascan#ignore_null DataplexDatascan#ignore_null}
	IgnoreNull interface{} `field:"optional" json:"ignoreNull" yaml:"ignoreNull"`
	// non_null_expectation block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/resources/dataplex_datascan#non_null_expectation DataplexDatascan#non_null_expectation}
	NonNullExpectation *DataplexDatascanDataQualitySpecRulesNonNullExpectation `field:"optional" json:"nonNullExpectation" yaml:"nonNullExpectation"`
	// range_expectation block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/resources/dataplex_datascan#range_expectation DataplexDatascan#range_expectation}
	RangeExpectation *DataplexDatascanDataQualitySpecRulesRangeExpectation `field:"optional" json:"rangeExpectation" yaml:"rangeExpectation"`
	// regex_expectation block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/resources/dataplex_datascan#regex_expectation DataplexDatascan#regex_expectation}
	RegexExpectation *DataplexDatascanDataQualitySpecRulesRegexExpectation `field:"optional" json:"regexExpectation" yaml:"regexExpectation"`
	// row_condition_expectation block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/resources/dataplex_datascan#row_condition_expectation DataplexDatascan#row_condition_expectation}
	RowConditionExpectation *DataplexDatascanDataQualitySpecRulesRowConditionExpectation `field:"optional" json:"rowConditionExpectation" yaml:"rowConditionExpectation"`
	// set_expectation block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/resources/dataplex_datascan#set_expectation DataplexDatascan#set_expectation}
	SetExpectation *DataplexDatascanDataQualitySpecRulesSetExpectation `field:"optional" json:"setExpectation" yaml:"setExpectation"`
	// statistic_range_expectation block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/resources/dataplex_datascan#statistic_range_expectation DataplexDatascan#statistic_range_expectation}
	StatisticRangeExpectation *DataplexDatascanDataQualitySpecRulesStatisticRangeExpectation `field:"optional" json:"statisticRangeExpectation" yaml:"statisticRangeExpectation"`
	// table_condition_expectation block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/resources/dataplex_datascan#table_condition_expectation DataplexDatascan#table_condition_expectation}
	TableConditionExpectation *DataplexDatascanDataQualitySpecRulesTableConditionExpectation `field:"optional" json:"tableConditionExpectation" yaml:"tableConditionExpectation"`
	// The minimum ratio of passing_rows / total_rows required to pass this rule, with a range of [0.0, 1.0]. 0 indicates default value (i.e. 1.0).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/resources/dataplex_datascan#threshold DataplexDatascan#threshold}
	Threshold *float64 `field:"optional" json:"threshold" yaml:"threshold"`
	// uniqueness_expectation block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/resources/dataplex_datascan#uniqueness_expectation DataplexDatascan#uniqueness_expectation}
	UniquenessExpectation *DataplexDatascanDataQualitySpecRulesUniquenessExpectation `field:"optional" json:"uniquenessExpectation" yaml:"uniquenessExpectation"`
}

