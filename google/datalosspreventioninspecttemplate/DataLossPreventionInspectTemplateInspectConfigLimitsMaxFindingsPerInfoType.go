// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datalosspreventioninspecttemplate


type DataLossPreventionInspectTemplateInspectConfigLimitsMaxFindingsPerInfoType struct {
	// info_type block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.83.0/docs/resources/data_loss_prevention_inspect_template#info_type DataLossPreventionInspectTemplate#info_type}
	InfoType *DataLossPreventionInspectTemplateInspectConfigLimitsMaxFindingsPerInfoTypeInfoType `field:"required" json:"infoType" yaml:"infoType"`
	// Max findings limit for the given infoType.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.83.0/docs/resources/data_loss_prevention_inspect_template#max_findings DataLossPreventionInspectTemplate#max_findings}
	MaxFindings *float64 `field:"required" json:"maxFindings" yaml:"maxFindings"`
}

