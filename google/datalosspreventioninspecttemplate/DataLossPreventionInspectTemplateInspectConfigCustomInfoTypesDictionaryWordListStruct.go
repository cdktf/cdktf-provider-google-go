// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datalosspreventioninspecttemplate


type DataLossPreventionInspectTemplateInspectConfigCustomInfoTypesDictionaryWordListStruct struct {
	// Words or phrases defining the dictionary.
	//
	// The dictionary must contain at least one
	// phrase and every phrase must contain at least 2 characters that are letters or digits.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.24.0/docs/resources/data_loss_prevention_inspect_template#words DataLossPreventionInspectTemplate#words}
	Words *[]*string `field:"required" json:"words" yaml:"words"`
}

