package datalosspreventiondeidentifytemplate


type DataLossPreventionDeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationCryptoReplaceFfxFpeConfig struct {
	// Common alphabets. Possible values: ["FFX_COMMON_NATIVE_ALPHABET_UNSPECIFIED", "NUMERIC", "HEXADECIMAL", "UPPER_CASE_ALPHA_NUMERIC", "ALPHA_NUMERIC"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/data_loss_prevention_deidentify_template#common_alphabet DataLossPreventionDeidentifyTemplate#common_alphabet}
	CommonAlphabet *string `field:"optional" json:"commonAlphabet" yaml:"commonAlphabet"`
	// context block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/data_loss_prevention_deidentify_template#context DataLossPreventionDeidentifyTemplate#context}
	Context *DataLossPreventionDeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationCryptoReplaceFfxFpeConfigContext `field:"optional" json:"context" yaml:"context"`
	// crypto_key block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/data_loss_prevention_deidentify_template#crypto_key DataLossPreventionDeidentifyTemplate#crypto_key}
	CryptoKey *DataLossPreventionDeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationCryptoReplaceFfxFpeConfigCryptoKey `field:"optional" json:"cryptoKey" yaml:"cryptoKey"`
	// This is supported by mapping these to the alphanumeric characters that the FFX mode natively supports.
	//
	// This happens before/after encryption/decryption. Each character listed must appear only once. Number of characters must be in the range \[2, 95\]. This must be encoded as ASCII. The order of characters does not matter. The full list of allowed characters is:
	//
	// ''0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz ~'!@#$%^&*()_-+={[}]|:;"'<,>.?/''
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/data_loss_prevention_deidentify_template#custom_alphabet DataLossPreventionDeidentifyTemplate#custom_alphabet}
	CustomAlphabet *string `field:"optional" json:"customAlphabet" yaml:"customAlphabet"`
	// The native way to select the alphabet. Must be in the range \[2, 95\].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/data_loss_prevention_deidentify_template#radix DataLossPreventionDeidentifyTemplate#radix}
	Radix *float64 `field:"optional" json:"radix" yaml:"radix"`
	// surrogate_info_type block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/data_loss_prevention_deidentify_template#surrogate_info_type DataLossPreventionDeidentifyTemplate#surrogate_info_type}
	SurrogateInfoType *DataLossPreventionDeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationCryptoReplaceFfxFpeConfigSurrogateInfoType `field:"optional" json:"surrogateInfoType" yaml:"surrogateInfoType"`
}

