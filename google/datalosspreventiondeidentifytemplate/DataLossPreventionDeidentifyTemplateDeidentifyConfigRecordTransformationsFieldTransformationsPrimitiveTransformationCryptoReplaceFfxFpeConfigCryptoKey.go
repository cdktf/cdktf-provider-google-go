package datalosspreventiondeidentifytemplate


type DataLossPreventionDeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationCryptoReplaceFfxFpeConfigCryptoKey struct {
	// kms_wrapped block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/data_loss_prevention_deidentify_template#kms_wrapped DataLossPreventionDeidentifyTemplate#kms_wrapped}
	KmsWrapped *DataLossPreventionDeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationCryptoReplaceFfxFpeConfigCryptoKeyKmsWrapped `field:"optional" json:"kmsWrapped" yaml:"kmsWrapped"`
	// transient block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/data_loss_prevention_deidentify_template#transient DataLossPreventionDeidentifyTemplate#transient}
	Transient *DataLossPreventionDeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationCryptoReplaceFfxFpeConfigCryptoKeyTransient `field:"optional" json:"transient" yaml:"transient"`
	// unwrapped block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/data_loss_prevention_deidentify_template#unwrapped DataLossPreventionDeidentifyTemplate#unwrapped}
	Unwrapped *DataLossPreventionDeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationCryptoReplaceFfxFpeConfigCryptoKeyUnwrapped `field:"optional" json:"unwrapped" yaml:"unwrapped"`
}
