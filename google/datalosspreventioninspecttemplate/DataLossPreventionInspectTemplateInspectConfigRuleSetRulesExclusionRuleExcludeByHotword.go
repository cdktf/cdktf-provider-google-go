package datalosspreventioninspecttemplate


type DataLossPreventionInspectTemplateInspectConfigRuleSetRulesExclusionRuleExcludeByHotword struct {
	// hotword_regex block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/data_loss_prevention_inspect_template#hotword_regex DataLossPreventionInspectTemplate#hotword_regex}
	HotwordRegex *DataLossPreventionInspectTemplateInspectConfigRuleSetRulesExclusionRuleExcludeByHotwordHotwordRegex `field:"required" json:"hotwordRegex" yaml:"hotwordRegex"`
	// proximity block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/data_loss_prevention_inspect_template#proximity DataLossPreventionInspectTemplate#proximity}
	Proximity *DataLossPreventionInspectTemplateInspectConfigRuleSetRulesExclusionRuleExcludeByHotwordProximity `field:"required" json:"proximity" yaml:"proximity"`
}

