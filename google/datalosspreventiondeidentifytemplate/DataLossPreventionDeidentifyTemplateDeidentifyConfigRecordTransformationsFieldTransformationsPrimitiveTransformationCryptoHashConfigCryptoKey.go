// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datalosspreventiondeidentifytemplate


type DataLossPreventionDeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationCryptoHashConfigCryptoKey struct {
	// kms_wrapped block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.14.0/docs/resources/data_loss_prevention_deidentify_template#kms_wrapped DataLossPreventionDeidentifyTemplate#kms_wrapped}
	KmsWrapped *DataLossPreventionDeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationCryptoHashConfigCryptoKeyKmsWrapped `field:"optional" json:"kmsWrapped" yaml:"kmsWrapped"`
	// transient block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.14.0/docs/resources/data_loss_prevention_deidentify_template#transient DataLossPreventionDeidentifyTemplate#transient}
	Transient *DataLossPreventionDeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationCryptoHashConfigCryptoKeyTransient `field:"optional" json:"transient" yaml:"transient"`
	// unwrapped block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.14.0/docs/resources/data_loss_prevention_deidentify_template#unwrapped DataLossPreventionDeidentifyTemplate#unwrapped}
	Unwrapped *DataLossPreventionDeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationCryptoHashConfigCryptoKeyUnwrapped `field:"optional" json:"unwrapped" yaml:"unwrapped"`
}

