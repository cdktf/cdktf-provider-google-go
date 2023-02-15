package datalosspreventioninspecttemplate


type DataLossPreventionInspectTemplateInspectConfigRuleSetRulesExclusionRule struct {
	// How the rule is applied. See the documentation for more information: https://cloud.google.com/dlp/docs/reference/rest/v2/InspectConfig#MatchingType Possible values: ["MATCHING_TYPE_FULL_MATCH", "MATCHING_TYPE_PARTIAL_MATCH", "MATCHING_TYPE_INVERSE_MATCH"].
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/data_loss_prevention_inspect_template#matching_type DataLossPreventionInspectTemplate#matching_type}
	MatchingType *string `field:"required" json:"matchingType" yaml:"matchingType"`
	// dictionary block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/data_loss_prevention_inspect_template#dictionary DataLossPreventionInspectTemplate#dictionary}
	Dictionary *DataLossPreventionInspectTemplateInspectConfigRuleSetRulesExclusionRuleDictionary `field:"optional" json:"dictionary" yaml:"dictionary"`
	// exclude_info_types block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/data_loss_prevention_inspect_template#exclude_info_types DataLossPreventionInspectTemplate#exclude_info_types}
	ExcludeInfoTypes *DataLossPreventionInspectTemplateInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypes `field:"optional" json:"excludeInfoTypes" yaml:"excludeInfoTypes"`
	// regex block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/data_loss_prevention_inspect_template#regex DataLossPreventionInspectTemplate#regex}
	Regex *DataLossPreventionInspectTemplateInspectConfigRuleSetRulesExclusionRuleRegex `field:"optional" json:"regex" yaml:"regex"`
}

