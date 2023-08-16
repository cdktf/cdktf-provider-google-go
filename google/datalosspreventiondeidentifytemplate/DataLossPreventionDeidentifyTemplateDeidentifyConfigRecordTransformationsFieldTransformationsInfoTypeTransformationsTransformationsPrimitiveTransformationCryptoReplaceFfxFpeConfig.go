package datalosspreventiondeidentifytemplate


type DataLossPreventionDeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoReplaceFfxFpeConfig struct {
	// crypto_key block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/data_loss_prevention_deidentify_template#crypto_key DataLossPreventionDeidentifyTemplate#crypto_key}
	CryptoKey *DataLossPreventionDeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoReplaceFfxFpeConfigCryptoKey `field:"required" json:"cryptoKey" yaml:"cryptoKey"`
	// Common alphabets. Only one of this, 'custom_alphabet' or 'radix' must be specified. Possible values: ["NUMERIC", "HEXADECIMAL", "UPPER_CASE_ALPHA_NUMERIC", "ALPHA_NUMERIC"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/data_loss_prevention_deidentify_template#common_alphabet DataLossPreventionDeidentifyTemplate#common_alphabet}
	CommonAlphabet *string `field:"optional" json:"commonAlphabet" yaml:"commonAlphabet"`
	// context block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/data_loss_prevention_deidentify_template#context DataLossPreventionDeidentifyTemplate#context}
	Context *DataLossPreventionDeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoReplaceFfxFpeConfigContext `field:"optional" json:"context" yaml:"context"`
	// This is supported by mapping these to the alphanumeric characters that the FFX mode natively supports.
	//
	// This happens before/after encryption/decryption. Each character listed must appear only once. Number of characters must be in the range \[2, 95\]. This must be encoded as ASCII. The order of characters does not matter. The full list of allowed characters is:
	//
	// ''0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz ~'!@#$%^&*()_-+={[}]|:;"'<,>.?/''. Only one of this, 'common_alphabet' or 'radix' must be specified.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/data_loss_prevention_deidentify_template#custom_alphabet DataLossPreventionDeidentifyTemplate#custom_alphabet}
	CustomAlphabet *string `field:"optional" json:"customAlphabet" yaml:"customAlphabet"`
	// The native way to select the alphabet.
	//
	// Must be in the range \[2, 95\]. Only one of this, 'custom_alphabet' or 'common_alphabet' must be specified.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/data_loss_prevention_deidentify_template#radix DataLossPreventionDeidentifyTemplate#radix}
	Radix *float64 `field:"optional" json:"radix" yaml:"radix"`
	// surrogate_info_type block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/data_loss_prevention_deidentify_template#surrogate_info_type DataLossPreventionDeidentifyTemplate#surrogate_info_type}
	SurrogateInfoType *DataLossPreventionDeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoReplaceFfxFpeConfigSurrogateInfoType `field:"optional" json:"surrogateInfoType" yaml:"surrogateInfoType"`
}

