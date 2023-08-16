package datalosspreventionjobtrigger


type DataLossPreventionJobTriggerInspectJobInspectConfigRuleSetRulesHotwordRule struct {
	// hotword_regex block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/data_loss_prevention_job_trigger#hotword_regex DataLossPreventionJobTrigger#hotword_regex}
	HotwordRegex *DataLossPreventionJobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleHotwordRegex `field:"optional" json:"hotwordRegex" yaml:"hotwordRegex"`
	// likelihood_adjustment block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/data_loss_prevention_job_trigger#likelihood_adjustment DataLossPreventionJobTrigger#likelihood_adjustment}
	LikelihoodAdjustment *DataLossPreventionJobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleLikelihoodAdjustment `field:"optional" json:"likelihoodAdjustment" yaml:"likelihoodAdjustment"`
	// proximity block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/data_loss_prevention_job_trigger#proximity DataLossPreventionJobTrigger#proximity}
	Proximity *DataLossPreventionJobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleProximity `field:"optional" json:"proximity" yaml:"proximity"`
}

