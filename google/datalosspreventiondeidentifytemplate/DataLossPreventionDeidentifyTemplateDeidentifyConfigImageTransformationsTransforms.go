package datalosspreventiondeidentifytemplate


type DataLossPreventionDeidentifyTemplateDeidentifyConfigImageTransformationsTransforms struct {
	// all_info_types block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/data_loss_prevention_deidentify_template#all_info_types DataLossPreventionDeidentifyTemplate#all_info_types}
	AllInfoTypes *DataLossPreventionDeidentifyTemplateDeidentifyConfigImageTransformationsTransformsAllInfoTypes `field:"optional" json:"allInfoTypes" yaml:"allInfoTypes"`
	// all_text block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/data_loss_prevention_deidentify_template#all_text DataLossPreventionDeidentifyTemplate#all_text}
	AllText *DataLossPreventionDeidentifyTemplateDeidentifyConfigImageTransformationsTransformsAllText `field:"optional" json:"allText" yaml:"allText"`
	// redaction_color block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/data_loss_prevention_deidentify_template#redaction_color DataLossPreventionDeidentifyTemplate#redaction_color}
	RedactionColor *DataLossPreventionDeidentifyTemplateDeidentifyConfigImageTransformationsTransformsRedactionColor `field:"optional" json:"redactionColor" yaml:"redactionColor"`
	// selected_info_types block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/data_loss_prevention_deidentify_template#selected_info_types DataLossPreventionDeidentifyTemplate#selected_info_types}
	SelectedInfoTypes *DataLossPreventionDeidentifyTemplateDeidentifyConfigImageTransformationsTransformsSelectedInfoTypes `field:"optional" json:"selectedInfoTypes" yaml:"selectedInfoTypes"`
}

