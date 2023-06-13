package datalosspreventionjobtrigger


type DataLossPreventionJobTriggerInspectJobInspectConfigRuleSet struct {
	// rules block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/resources/data_loss_prevention_job_trigger#rules DataLossPreventionJobTrigger#rules}
	Rules interface{} `field:"required" json:"rules" yaml:"rules"`
	// info_types block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/resources/data_loss_prevention_job_trigger#info_types DataLossPreventionJobTrigger#info_types}
	InfoTypes interface{} `field:"optional" json:"infoTypes" yaml:"infoTypes"`
}

