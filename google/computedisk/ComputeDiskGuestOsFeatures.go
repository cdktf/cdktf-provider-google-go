// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package computedisk


type ComputeDiskGuestOsFeatures struct {
	// The type of supported feature. Read [Enabling guest operating system features](https://cloud.google.com/compute/docs/images/create-delete-deprecate-private-images#guest-os-features) to see a list of available options.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.18.0/docs/resources/compute_disk#type ComputeDisk#type}
	Type *string `field:"required" json:"type" yaml:"type"`
}

