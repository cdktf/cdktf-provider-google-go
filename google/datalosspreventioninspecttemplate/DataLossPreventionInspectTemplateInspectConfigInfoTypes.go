package datalosspreventioninspecttemplate


type DataLossPreventionInspectTemplateInspectConfigInfoTypes struct {
	// Name of the information type.
	//
	// Either a name of your choosing when creating a CustomInfoType, or one of the names listed
	// at https://cloud.google.com/dlp/docs/infotypes-reference when specifying a built-in type.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/data_loss_prevention_inspect_template#name DataLossPreventionInspectTemplate#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// sensitivity_score block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/data_loss_prevention_inspect_template#sensitivity_score DataLossPreventionInspectTemplate#sensitivity_score}
	SensitivityScore *DataLossPreventionInspectTemplateInspectConfigInfoTypesSensitivityScore `field:"optional" json:"sensitivityScore" yaml:"sensitivityScore"`
	// Version of the information type to use. By default, the version is set to stable.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/data_loss_prevention_inspect_template#version DataLossPreventionInspectTemplate#version}
	Version *string `field:"optional" json:"version" yaml:"version"`
}

