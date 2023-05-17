package datalosspreventiondeidentifytemplate


type DataLossPreventionDeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsInfoTypes struct {
	// Name of the information type.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/data_loss_prevention_deidentify_template#name DataLossPreventionDeidentifyTemplate#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Version name for this InfoType.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/data_loss_prevention_deidentify_template#version DataLossPreventionDeidentifyTemplate#version}
	Version *string `field:"optional" json:"version" yaml:"version"`
}

