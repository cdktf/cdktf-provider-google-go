package datalosspreventiondeidentifytemplate


type DataLossPreventionDeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformation struct {
	// character_mask_config block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/data_loss_prevention_deidentify_template#character_mask_config DataLossPreventionDeidentifyTemplate#character_mask_config}
	CharacterMaskConfig *DataLossPreventionDeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationCharacterMaskConfig `field:"optional" json:"characterMaskConfig" yaml:"characterMaskConfig"`
	// redact_config block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/data_loss_prevention_deidentify_template#redact_config DataLossPreventionDeidentifyTemplate#redact_config}
	RedactConfig *DataLossPreventionDeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationRedactConfig `field:"optional" json:"redactConfig" yaml:"redactConfig"`
	// replace_config block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/data_loss_prevention_deidentify_template#replace_config DataLossPreventionDeidentifyTemplate#replace_config}
	ReplaceConfig *DataLossPreventionDeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationReplaceConfig `field:"optional" json:"replaceConfig" yaml:"replaceConfig"`
}

