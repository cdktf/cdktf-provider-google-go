// Prebuilt google Provider for Terraform CDK (cdktf)
package google


type DataLossPreventionInspectTemplateInspectConfigRuleSetRules struct {
	// exclusion_rule block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/data_loss_prevention_inspect_template#exclusion_rule DataLossPreventionInspectTemplate#exclusion_rule}
	ExclusionRule *DataLossPreventionInspectTemplateInspectConfigRuleSetRulesExclusionRule `field:"optional" json:"exclusionRule" yaml:"exclusionRule"`
	// hotword_rule block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/data_loss_prevention_inspect_template#hotword_rule DataLossPreventionInspectTemplate#hotword_rule}
	HotwordRule *DataLossPreventionInspectTemplateInspectConfigRuleSetRulesHotwordRule `field:"optional" json:"hotwordRule" yaml:"hotwordRule"`
}

