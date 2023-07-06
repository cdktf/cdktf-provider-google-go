package datalosspreventionjobtrigger


type DataLossPreventionJobTriggerInspectJobInspectConfigCustomInfoTypesStoredType struct {
	// Resource name of the requested StoredInfoType, for example 'organizations/433245324/storedInfoTypes/432452342' or 'projects/project-id/storedInfoTypes/432452342'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.72.1/docs/resources/data_loss_prevention_job_trigger#name DataLossPreventionJobTrigger#name}
	Name *string `field:"required" json:"name" yaml:"name"`
}

