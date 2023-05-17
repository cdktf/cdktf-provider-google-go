package datalosspreventionjobtrigger


type DataLossPreventionJobTriggerInspectJobInspectConfigRuleSetRulesHotwordRuleHotwordRegex struct {
	// The index of the submatch to extract as findings.
	//
	// When not specified,
	// the entire match is returned. No more than 3 may be included.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/data_loss_prevention_job_trigger#group_indexes DataLossPreventionJobTrigger#group_indexes}
	GroupIndexes *[]*float64 `field:"optional" json:"groupIndexes" yaml:"groupIndexes"`
	// Pattern defining the regular expression. Its syntax (https://github.com/google/re2/wiki/Syntax) can be found under the google/re2 repository on GitHub.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/data_loss_prevention_job_trigger#pattern DataLossPreventionJobTrigger#pattern}
	Pattern *string `field:"optional" json:"pattern" yaml:"pattern"`
}

