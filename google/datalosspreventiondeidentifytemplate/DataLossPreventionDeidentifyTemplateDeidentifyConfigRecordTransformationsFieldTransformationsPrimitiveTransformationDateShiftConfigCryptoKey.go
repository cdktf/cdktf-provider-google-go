// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datalosspreventiondeidentifytemplate


type DataLossPreventionDeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationDateShiftConfigCryptoKey struct {
	// kms_wrapped block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.31.1/docs/resources/data_loss_prevention_deidentify_template#kms_wrapped DataLossPreventionDeidentifyTemplate#kms_wrapped}
	KmsWrapped *DataLossPreventionDeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationDateShiftConfigCryptoKeyKmsWrapped `field:"optional" json:"kmsWrapped" yaml:"kmsWrapped"`
	// transient block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.31.1/docs/resources/data_loss_prevention_deidentify_template#transient DataLossPreventionDeidentifyTemplate#transient}
	Transient *DataLossPreventionDeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationDateShiftConfigCryptoKeyTransient `field:"optional" json:"transient" yaml:"transient"`
	// unwrapped block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.31.1/docs/resources/data_loss_prevention_deidentify_template#unwrapped DataLossPreventionDeidentifyTemplate#unwrapped}
	Unwrapped *DataLossPreventionDeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsPrimitiveTransformationDateShiftConfigCryptoKeyUnwrapped `field:"optional" json:"unwrapped" yaml:"unwrapped"`
}

