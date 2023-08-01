package datalosspreventiondeidentifytemplate


type DataLossPreventionDeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationCharacterMaskConfig struct {
	// characters_to_ignore block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/data_loss_prevention_deidentify_template#characters_to_ignore DataLossPreventionDeidentifyTemplate#characters_to_ignore}
	CharactersToIgnore interface{} `field:"optional" json:"charactersToIgnore" yaml:"charactersToIgnore"`
	// Character to use to mask the sensitive values—for example, * for an alphabetic string such as a name, or 0 for a numeric string such as ZIP code or credit card number.
	//
	// This string must have a length of 1. If not supplied, this value defaults to * for
	// strings, and 0 for digits.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/data_loss_prevention_deidentify_template#masking_character DataLossPreventionDeidentifyTemplate#masking_character}
	MaskingCharacter *string `field:"optional" json:"maskingCharacter" yaml:"maskingCharacter"`
	// Number of characters to mask.
	//
	// If not set, all matching chars will be masked. Skipped characters do not count towards this tally.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/data_loss_prevention_deidentify_template#number_to_mask DataLossPreventionDeidentifyTemplate#number_to_mask}
	NumberToMask *float64 `field:"optional" json:"numberToMask" yaml:"numberToMask"`
	// Mask characters in reverse order.
	//
	// For example, if masking_character is 0, number_to_mask is 14, and reverse_order is 'false', then the
	// input string '1234-5678-9012-3456' is masked as '00000000000000-3456'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/data_loss_prevention_deidentify_template#reverse_order DataLossPreventionDeidentifyTemplate#reverse_order}
	ReverseOrder interface{} `field:"optional" json:"reverseOrder" yaml:"reverseOrder"`
}

