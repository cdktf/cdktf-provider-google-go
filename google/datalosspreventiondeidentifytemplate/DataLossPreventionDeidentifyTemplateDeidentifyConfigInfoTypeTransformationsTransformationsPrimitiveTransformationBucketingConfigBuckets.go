// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datalosspreventiondeidentifytemplate


type DataLossPreventionDeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBuckets struct {
	// replacement_value block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.47.0/docs/resources/data_loss_prevention_deidentify_template#replacement_value DataLossPreventionDeidentifyTemplate#replacement_value}
	ReplacementValue *DataLossPreventionDeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBucketsReplacementValue `field:"required" json:"replacementValue" yaml:"replacementValue"`
	// max block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.47.0/docs/resources/data_loss_prevention_deidentify_template#max DataLossPreventionDeidentifyTemplate#max}
	Max *DataLossPreventionDeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBucketsMax `field:"optional" json:"max" yaml:"max"`
	// min block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.47.0/docs/resources/data_loss_prevention_deidentify_template#min DataLossPreventionDeidentifyTemplate#min}
	Min *DataLossPreventionDeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationsPrimitiveTransformationBucketingConfigBucketsMin `field:"optional" json:"min" yaml:"min"`
}

