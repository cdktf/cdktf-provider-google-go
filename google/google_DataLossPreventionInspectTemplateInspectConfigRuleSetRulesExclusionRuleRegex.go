// Prebuilt google Provider for Terraform CDK (cdktf)
package google


type DataLossPreventionInspectTemplateInspectConfigRuleSetRulesExclusionRuleRegex struct {
	// Pattern defining the regular expression. Its syntax (https://github.com/google/re2/wiki/Syntax) can be found under the google/re2 repository on GitHub.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/data_loss_prevention_inspect_template#pattern DataLossPreventionInspectTemplate#pattern}
	Pattern *string `field:"required" json:"pattern" yaml:"pattern"`
	// The index of the submatch to extract as findings.
	//
	// When not specified, the entire match is returned. No more than 3 may be included.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/data_loss_prevention_inspect_template#group_indexes DataLossPreventionInspectTemplate#group_indexes}
	GroupIndexes *[]*float64 `field:"optional" json:"groupIndexes" yaml:"groupIndexes"`
}

