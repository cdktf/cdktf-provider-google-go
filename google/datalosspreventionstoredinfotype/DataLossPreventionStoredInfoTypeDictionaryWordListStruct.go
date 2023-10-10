// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datalosspreventionstoredinfotype


type DataLossPreventionStoredInfoTypeDictionaryWordListStruct struct {
	// Words or phrases defining the dictionary.
	//
	// The dictionary must contain at least one
	// phrase and every phrase must contain at least 2 characters that are letters or digits.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.1.0/docs/resources/data_loss_prevention_stored_info_type#words DataLossPreventionStoredInfoType#words}
	Words *[]*string `field:"required" json:"words" yaml:"words"`
}

