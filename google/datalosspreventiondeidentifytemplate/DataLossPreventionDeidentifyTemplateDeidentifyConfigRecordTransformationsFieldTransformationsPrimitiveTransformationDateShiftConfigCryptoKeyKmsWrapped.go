package datalosspreventiondeidentifytemplate


type DataLossPreventionDeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationDateShiftConfigCryptoKeyKmsWrapped struct {
	// The resource name of the KMS CryptoKey to use for unwrapping.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/data_loss_prevention_deidentify_template#crypto_key_name DataLossPreventionDeidentifyTemplate#crypto_key_name}
	CryptoKeyName *string `field:"required" json:"cryptoKeyName" yaml:"cryptoKeyName"`
	// The wrapped data crypto key.
	//
	// A base64-encoded string.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/data_loss_prevention_deidentify_template#wrapped_key DataLossPreventionDeidentifyTemplate#wrapped_key}
	WrappedKey *string `field:"required" json:"wrappedKey" yaml:"wrappedKey"`
}
