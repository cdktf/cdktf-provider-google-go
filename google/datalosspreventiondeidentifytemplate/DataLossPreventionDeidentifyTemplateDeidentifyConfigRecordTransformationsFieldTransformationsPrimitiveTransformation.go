package datalosspreventiondeidentifytemplate


type DataLossPreventionDeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformation struct {
	// bucketing_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/data_loss_prevention_deidentify_template#bucketing_config DataLossPreventionDeidentifyTemplate#bucketing_config}
	BucketingConfig *DataLossPreventionDeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationBucketingConfig `field:"optional" json:"bucketingConfig" yaml:"bucketingConfig"`
	// character_mask_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/data_loss_prevention_deidentify_template#character_mask_config DataLossPreventionDeidentifyTemplate#character_mask_config}
	CharacterMaskConfig *DataLossPreventionDeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationCharacterMaskConfig `field:"optional" json:"characterMaskConfig" yaml:"characterMaskConfig"`
	// crypto_deterministic_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/data_loss_prevention_deidentify_template#crypto_deterministic_config DataLossPreventionDeidentifyTemplate#crypto_deterministic_config}
	CryptoDeterministicConfig *DataLossPreventionDeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationCryptoDeterministicConfig `field:"optional" json:"cryptoDeterministicConfig" yaml:"cryptoDeterministicConfig"`
	// crypto_hash_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/data_loss_prevention_deidentify_template#crypto_hash_config DataLossPreventionDeidentifyTemplate#crypto_hash_config}
	CryptoHashConfig *DataLossPreventionDeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationCryptoHashConfig `field:"optional" json:"cryptoHashConfig" yaml:"cryptoHashConfig"`
	// crypto_replace_ffx_fpe_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/data_loss_prevention_deidentify_template#crypto_replace_ffx_fpe_config DataLossPreventionDeidentifyTemplate#crypto_replace_ffx_fpe_config}
	CryptoReplaceFfxFpeConfig *DataLossPreventionDeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationCryptoReplaceFfxFpeConfig `field:"optional" json:"cryptoReplaceFfxFpeConfig" yaml:"cryptoReplaceFfxFpeConfig"`
	// date_shift_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/data_loss_prevention_deidentify_template#date_shift_config DataLossPreventionDeidentifyTemplate#date_shift_config}
	DateShiftConfig *DataLossPreventionDeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationDateShiftConfig `field:"optional" json:"dateShiftConfig" yaml:"dateShiftConfig"`
	// fixed_size_bucketing_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/data_loss_prevention_deidentify_template#fixed_size_bucketing_config DataLossPreventionDeidentifyTemplate#fixed_size_bucketing_config}
	FixedSizeBucketingConfig *DataLossPreventionDeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationFixedSizeBucketingConfig `field:"optional" json:"fixedSizeBucketingConfig" yaml:"fixedSizeBucketingConfig"`
	// redact_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/data_loss_prevention_deidentify_template#redact_config DataLossPreventionDeidentifyTemplate#redact_config}
	RedactConfig *DataLossPreventionDeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationRedactConfig `field:"optional" json:"redactConfig" yaml:"redactConfig"`
	// replace_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/data_loss_prevention_deidentify_template#replace_config DataLossPreventionDeidentifyTemplate#replace_config}
	ReplaceConfig *DataLossPreventionDeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationReplaceConfig `field:"optional" json:"replaceConfig" yaml:"replaceConfig"`
	// replace_dictionary_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/data_loss_prevention_deidentify_template#replace_dictionary_config DataLossPreventionDeidentifyTemplate#replace_dictionary_config}
	ReplaceDictionaryConfig *DataLossPreventionDeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationReplaceDictionaryConfig `field:"optional" json:"replaceDictionaryConfig" yaml:"replaceDictionaryConfig"`
	// time_part_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/data_loss_prevention_deidentify_template#time_part_config DataLossPreventionDeidentifyTemplate#time_part_config}
	TimePartConfig *DataLossPreventionDeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationTimePartConfig `field:"optional" json:"timePartConfig" yaml:"timePartConfig"`
}

