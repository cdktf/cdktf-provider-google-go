package datalosspreventionjobtrigger


type DataLossPreventionJobTriggerInspectJob struct {
	// actions block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/data_loss_prevention_job_trigger#actions DataLossPreventionJobTrigger#actions}
	Actions interface{} `field:"required" json:"actions" yaml:"actions"`
	// The name of the template to run when this job is triggered.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/data_loss_prevention_job_trigger#inspect_template_name DataLossPreventionJobTrigger#inspect_template_name}
	InspectTemplateName *string `field:"required" json:"inspectTemplateName" yaml:"inspectTemplateName"`
	// storage_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/data_loss_prevention_job_trigger#storage_config DataLossPreventionJobTrigger#storage_config}
	StorageConfig *DataLossPreventionJobTriggerInspectJobStorageConfig `field:"required" json:"storageConfig" yaml:"storageConfig"`
	// inspect_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/data_loss_prevention_job_trigger#inspect_config DataLossPreventionJobTrigger#inspect_config}
	InspectConfig *DataLossPreventionJobTriggerInspectJobInspectConfig `field:"optional" json:"inspectConfig" yaml:"inspectConfig"`
}

