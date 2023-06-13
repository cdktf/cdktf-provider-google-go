package datalosspreventionjobtrigger


type DataLossPreventionJobTriggerInspectJobInspectConfigLimits struct {
	// max_findings_per_info_type block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/resources/data_loss_prevention_job_trigger#max_findings_per_info_type DataLossPreventionJobTrigger#max_findings_per_info_type}
	MaxFindingsPerInfoType interface{} `field:"optional" json:"maxFindingsPerInfoType" yaml:"maxFindingsPerInfoType"`
	// Max number of findings that will be returned for each item scanned. The maximum returned is 2000.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/resources/data_loss_prevention_job_trigger#max_findings_per_item DataLossPreventionJobTrigger#max_findings_per_item}
	MaxFindingsPerItem *float64 `field:"optional" json:"maxFindingsPerItem" yaml:"maxFindingsPerItem"`
	// Max number of findings that will be returned per request/job. The maximum returned is 2000.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/resources/data_loss_prevention_job_trigger#max_findings_per_request DataLossPreventionJobTrigger#max_findings_per_request}
	MaxFindingsPerRequest *float64 `field:"optional" json:"maxFindingsPerRequest" yaml:"maxFindingsPerRequest"`
}

