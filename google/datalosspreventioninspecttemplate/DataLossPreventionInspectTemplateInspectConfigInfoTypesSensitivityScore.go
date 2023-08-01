package datalosspreventioninspecttemplate


type DataLossPreventionInspectTemplateInspectConfigInfoTypesSensitivityScore struct {
	// The sensitivity score applied to the resource. Possible values: ["SENSITIVITY_LOW", "SENSITIVITY_MODERATE", "SENSITIVITY_HIGH"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/data_loss_prevention_inspect_template#score DataLossPreventionInspectTemplate#score}
	Score *string `field:"required" json:"score" yaml:"score"`
}

