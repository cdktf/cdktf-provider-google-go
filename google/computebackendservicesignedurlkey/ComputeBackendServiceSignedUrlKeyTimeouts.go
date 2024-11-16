// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package computebackendservicesignedurlkey


type ComputeBackendServiceSignedUrlKeyTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.11.2/docs/resources/compute_backend_service_signed_url_key#create ComputeBackendServiceSignedUrlKey#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.11.2/docs/resources/compute_backend_service_signed_url_key#delete ComputeBackendServiceSignedUrlKey#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
}

