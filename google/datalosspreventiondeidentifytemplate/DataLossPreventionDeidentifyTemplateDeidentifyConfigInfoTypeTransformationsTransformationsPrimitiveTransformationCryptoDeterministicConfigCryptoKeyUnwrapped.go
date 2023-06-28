package datalosspreventiondeidentifytemplate


type DataLossPreventionDeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoDeterministicConfigCryptoKeyUnwrapped struct {
	// A 128/192/256 bit key.
	//
	// A base64-encoded string.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.71.0/docs/resources/data_loss_prevention_deidentify_template#key DataLossPreventionDeidentifyTemplate#key}
	Key *string `field:"required" json:"key" yaml:"key"`
}

