// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package computeimage


type ComputeImageShieldedInstanceInitialStateKeks struct {
	// The raw content in the secure keys file.
	//
	// A base64-encoded string.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.31.0/docs/resources/compute_image#content ComputeImage#content}
	Content *string `field:"required" json:"content" yaml:"content"`
	// The file type of source file.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.31.0/docs/resources/compute_image#file_type ComputeImage#file_type}
	FileType *string `field:"optional" json:"fileType" yaml:"fileType"`
}

