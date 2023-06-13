package datalosspreventionjobtrigger


type DataLossPreventionJobTriggerInspectJobInspectConfigLimitsMaxFindingsPerInfoType struct {
	// info_type block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/resources/data_loss_prevention_job_trigger#info_type DataLossPreventionJobTrigger#info_type}
	InfoType *DataLossPreventionJobTriggerInspectJobInspectConfigLimitsMaxFindingsPerInfoTypeInfoType `field:"optional" json:"infoType" yaml:"infoType"`
	// Max findings limit for the given infoType.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/resources/data_loss_prevention_job_trigger#max_findings DataLossPreventionJobTrigger#max_findings}
	MaxFindings *float64 `field:"optional" json:"maxFindings" yaml:"maxFindings"`
}

