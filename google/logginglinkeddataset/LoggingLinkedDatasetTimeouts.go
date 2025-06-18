// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package logginglinkeddataset


type LoggingLinkedDatasetTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.40.0/docs/resources/logging_linked_dataset#create LoggingLinkedDataset#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.40.0/docs/resources/logging_linked_dataset#delete LoggingLinkedDataset#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
}

