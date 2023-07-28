package datalosspreventionjobtrigger


type DataLossPreventionJobTriggerInspectJobInspectConfigInfoTypesSensitivityScore struct {
	// The sensitivity score applied to the resource. Possible values: ["SENSITIVITY_LOW", "SENSITIVITY_MODERATE", "SENSITIVITY_HIGH"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.75.1/docs/resources/data_loss_prevention_job_trigger#score DataLossPreventionJobTrigger#score}
	Score *string `field:"required" json:"score" yaml:"score"`
}

