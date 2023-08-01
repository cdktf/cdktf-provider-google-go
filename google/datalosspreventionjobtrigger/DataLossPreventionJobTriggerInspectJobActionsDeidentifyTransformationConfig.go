package datalosspreventionjobtrigger


type DataLossPreventionJobTriggerInspectJobActionsDeidentifyTransformationConfig struct {
	// If this template is specified, it will serve as the default de-identify template.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/data_loss_prevention_job_trigger#deidentify_template DataLossPreventionJobTrigger#deidentify_template}
	DeidentifyTemplate *string `field:"optional" json:"deidentifyTemplate" yaml:"deidentifyTemplate"`
	// If this template is specified, it will serve as the de-identify template for images.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/data_loss_prevention_job_trigger#image_redact_template DataLossPreventionJobTrigger#image_redact_template}
	ImageRedactTemplate *string `field:"optional" json:"imageRedactTemplate" yaml:"imageRedactTemplate"`
	// If this template is specified, it will serve as the de-identify template for structured content such as delimited files and tables.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/data_loss_prevention_job_trigger#structured_deidentify_template DataLossPreventionJobTrigger#structured_deidentify_template}
	StructuredDeidentifyTemplate *string `field:"optional" json:"structuredDeidentifyTemplate" yaml:"structuredDeidentifyTemplate"`
}

