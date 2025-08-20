// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package kmskeyhandle


type KmsKeyHandleTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.0/docs/resources/kms_key_handle#create KmsKeyHandle#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.0/docs/resources/kms_key_handle#delete KmsKeyHandle#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
}

