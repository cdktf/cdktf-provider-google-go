// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datalosspreventiondeidentifytemplate


type DataLossPreventionDeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoDeterministicConfig struct {
	// crypto_key block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.43.1/docs/resources/data_loss_prevention_deidentify_template#crypto_key DataLossPreventionDeidentifyTemplate#crypto_key}
	CryptoKey *DataLossPreventionDeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoDeterministicConfigCryptoKey `field:"required" json:"cryptoKey" yaml:"cryptoKey"`
	// surrogate_info_type block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.43.1/docs/resources/data_loss_prevention_deidentify_template#surrogate_info_type DataLossPreventionDeidentifyTemplate#surrogate_info_type}
	SurrogateInfoType *DataLossPreventionDeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoDeterministicConfigSurrogateInfoType `field:"required" json:"surrogateInfoType" yaml:"surrogateInfoType"`
	// context block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.43.1/docs/resources/data_loss_prevention_deidentify_template#context DataLossPreventionDeidentifyTemplate#context}
	Context *DataLossPreventionDeidentifyTemplateDeidentifyConfigRecordTransformationsFieldTransformationsInfoTypeTransformationsTransformationsPrimitiveTransformationCryptoDeterministicConfigContext `field:"optional" json:"context" yaml:"context"`
}

