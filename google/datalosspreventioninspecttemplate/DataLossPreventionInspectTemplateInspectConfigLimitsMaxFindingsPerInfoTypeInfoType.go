// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datalosspreventioninspecttemplate


type DataLossPreventionInspectTemplateInspectConfigLimitsMaxFindingsPerInfoTypeInfoType struct {
	// Name of the information type.
	//
	// Either a name of your choosing when creating a CustomInfoType, or one of the names listed
	// at https://cloud.google.com/dlp/docs/infotypes-reference when specifying a built-in type.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.14.1/docs/resources/data_loss_prevention_inspect_template#name DataLossPreventionInspectTemplate#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// sensitivity_score block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.14.1/docs/resources/data_loss_prevention_inspect_template#sensitivity_score DataLossPreventionInspectTemplate#sensitivity_score}
	SensitivityScore *DataLossPreventionInspectTemplateInspectConfigLimitsMaxFindingsPerInfoTypeInfoTypeSensitivityScore `field:"optional" json:"sensitivityScore" yaml:"sensitivityScore"`
	// Version name for this InfoType.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.14.1/docs/resources/data_loss_prevention_inspect_template#version DataLossPreventionInspectTemplate#version}
	Version *string `field:"optional" json:"version" yaml:"version"`
}

