package datalosspreventionjobtrigger


type DataLossPreventionJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleExcludeByHotword struct {
	// hotword_regex block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/data_loss_prevention_job_trigger#hotword_regex DataLossPreventionJobTrigger#hotword_regex}
	HotwordRegex *DataLossPreventionJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleExcludeByHotwordHotwordRegex `field:"optional" json:"hotwordRegex" yaml:"hotwordRegex"`
	// proximity block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/data_loss_prevention_job_trigger#proximity DataLossPreventionJobTrigger#proximity}
	Proximity *DataLossPreventionJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRuleExcludeByHotwordProximity `field:"optional" json:"proximity" yaml:"proximity"`
}

