package datalosspreventioninspecttemplate


type DataLossPreventionInspectTemplateInspectConfigRuleSetRulesHotwordRule struct {
	// hotword_regex block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/data_loss_prevention_inspect_template#hotword_regex DataLossPreventionInspectTemplate#hotword_regex}
	HotwordRegex *DataLossPreventionInspectTemplateInspectConfigRuleSetRulesHotwordRuleHotwordRegex `field:"required" json:"hotwordRegex" yaml:"hotwordRegex"`
	// likelihood_adjustment block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/data_loss_prevention_inspect_template#likelihood_adjustment DataLossPreventionInspectTemplate#likelihood_adjustment}
	LikelihoodAdjustment *DataLossPreventionInspectTemplateInspectConfigRuleSetRulesHotwordRuleLikelihoodAdjustment `field:"required" json:"likelihoodAdjustment" yaml:"likelihoodAdjustment"`
	// proximity block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/data_loss_prevention_inspect_template#proximity DataLossPreventionInspectTemplate#proximity}
	Proximity *DataLossPreventionInspectTemplateInspectConfigRuleSetRulesHotwordRuleProximity `field:"required" json:"proximity" yaml:"proximity"`
}

