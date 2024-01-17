// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datalosspreventionstoredinfotype


type DataLossPreventionStoredInfoTypeLargeCustomDictionaryOutputPath struct {
	// A url representing a file or path (no wildcards) in Cloud Storage. Example: 'gs://[BUCKET_NAME]/dictionary.txt'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.12.0/docs/resources/data_loss_prevention_stored_info_type#path DataLossPreventionStoredInfoType#path}
	Path *string `field:"required" json:"path" yaml:"path"`
}

