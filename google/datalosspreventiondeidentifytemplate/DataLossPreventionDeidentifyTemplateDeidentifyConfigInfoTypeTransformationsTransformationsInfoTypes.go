package datalosspreventiondeidentifytemplate


type DataLossPreventionDeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsInfoTypes struct {
	// Name of the information type.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.72.1/docs/resources/data_loss_prevention_deidentify_template#name DataLossPreventionDeidentifyTemplate#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// sensitivity_score block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.72.1/docs/resources/data_loss_prevention_deidentify_template#sensitivity_score DataLossPreventionDeidentifyTemplate#sensitivity_score}
	SensitivityScore *DataLossPreventionDeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsInfoTypesSensitivityScore `field:"optional" json:"sensitivityScore" yaml:"sensitivityScore"`
	// Version name for this InfoType.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.72.1/docs/resources/data_loss_prevention_deidentify_template#version DataLossPreventionDeidentifyTemplate#version}
	Version *string `field:"optional" json:"version" yaml:"version"`
}

