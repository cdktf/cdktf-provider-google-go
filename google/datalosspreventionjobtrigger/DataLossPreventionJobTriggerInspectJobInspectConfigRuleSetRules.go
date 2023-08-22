package datalosspreventionjobtrigger


type DataLossPreventionJobTriggerInspectJobInspectConfigRuleSetRules struct {
	// exclusion_rule block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/data_loss_prevention_job_trigger#exclusion_rule DataLossPreventionJobTrigger#exclusion_rule}
	ExclusionRule *DataLossPreventionJobTriggerInspectJobInspectConfigRuleSetRulesExclusionRule `field:"optional" json:"exclusionRule" yaml:"exclusionRule"`
	// hotword_rule block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/data_loss_prevention_job_trigger#hotword_rule DataLossPreventionJobTrigger#hotword_rule}
	HotwordRule *DataLossPreventionJobTriggerInspectJobInspectConfigRuleSetRulesHotwordRule `field:"optional" json:"hotwordRule" yaml:"hotwordRule"`
}

