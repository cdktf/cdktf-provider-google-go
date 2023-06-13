package datalosspreventiondeidentifytemplate


type DataLossPreventionDeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformation struct {
	// bucketing_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/resources/data_loss_prevention_deidentify_template#bucketing_config DataLossPreventionDeidentifyTemplate#bucketing_config}
	BucketingConfig *DataLossPreventionDeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfig `field:"optional" json:"bucketingConfig" yaml:"bucketingConfig"`
	// character_mask_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/resources/data_loss_prevention_deidentify_template#character_mask_config DataLossPreventionDeidentifyTemplate#character_mask_config}
	CharacterMaskConfig *DataLossPreventionDeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationCharacterMaskConfig `field:"optional" json:"characterMaskConfig" yaml:"characterMaskConfig"`
	// crypto_deterministic_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/resources/data_loss_prevention_deidentify_template#crypto_deterministic_config DataLossPreventionDeidentifyTemplate#crypto_deterministic_config}
	CryptoDeterministicConfig *DataLossPreventionDeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoDeterministicConfig `field:"optional" json:"cryptoDeterministicConfig" yaml:"cryptoDeterministicConfig"`
	// crypto_replace_ffx_fpe_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/resources/data_loss_prevention_deidentify_template#crypto_replace_ffx_fpe_config DataLossPreventionDeidentifyTemplate#crypto_replace_ffx_fpe_config}
	CryptoReplaceFfxFpeConfig *DataLossPreventionDeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoReplaceFfxFpeConfig `field:"optional" json:"cryptoReplaceFfxFpeConfig" yaml:"cryptoReplaceFfxFpeConfig"`
	// date_shift_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/resources/data_loss_prevention_deidentify_template#date_shift_config DataLossPreventionDeidentifyTemplate#date_shift_config}
	DateShiftConfig *DataLossPreventionDeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationDateShiftConfig `field:"optional" json:"dateShiftConfig" yaml:"dateShiftConfig"`
	// fixed_size_bucketing_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/resources/data_loss_prevention_deidentify_template#fixed_size_bucketing_config DataLossPreventionDeidentifyTemplate#fixed_size_bucketing_config}
	FixedSizeBucketingConfig *DataLossPreventionDeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationFixedSizeBucketingConfig `field:"optional" json:"fixedSizeBucketingConfig" yaml:"fixedSizeBucketingConfig"`
	// redact_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/resources/data_loss_prevention_deidentify_template#redact_config DataLossPreventionDeidentifyTemplate#redact_config}
	RedactConfig *DataLossPreventionDeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationRedactConfig `field:"optional" json:"redactConfig" yaml:"redactConfig"`
	// replace_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/resources/data_loss_prevention_deidentify_template#replace_config DataLossPreventionDeidentifyTemplate#replace_config}
	ReplaceConfig *DataLossPreventionDeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationReplaceConfig `field:"optional" json:"replaceConfig" yaml:"replaceConfig"`
	// replace_dictionary_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/resources/data_loss_prevention_deidentify_template#replace_dictionary_config DataLossPreventionDeidentifyTemplate#replace_dictionary_config}
	ReplaceDictionaryConfig *DataLossPreventionDeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationReplaceDictionaryConfig `field:"optional" json:"replaceDictionaryConfig" yaml:"replaceDictionaryConfig"`
	// Replace each matching finding with the name of the info type.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/resources/data_loss_prevention_deidentify_template#replace_with_info_type_config DataLossPreventionDeidentifyTemplate#replace_with_info_type_config}
	ReplaceWithInfoTypeConfig interface{} `field:"optional" json:"replaceWithInfoTypeConfig" yaml:"replaceWithInfoTypeConfig"`
	// time_part_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/resources/data_loss_prevention_deidentify_template#time_part_config DataLossPreventionDeidentifyTemplate#time_part_config}
	TimePartConfig *DataLossPreventionDeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationTimePartConfig `field:"optional" json:"timePartConfig" yaml:"timePartConfig"`
}

