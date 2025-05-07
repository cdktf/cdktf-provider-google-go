// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package computebackendservice


type ComputeBackendServiceConsistentHashHttpCookie struct {
	// Name of the cookie.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.34.0/docs/resources/compute_backend_service#name ComputeBackendService#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Path to set for the cookie.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.34.0/docs/resources/compute_backend_service#path ComputeBackendService#path}
	Path *string `field:"optional" json:"path" yaml:"path"`
	// ttl block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.34.0/docs/resources/compute_backend_service#ttl ComputeBackendService#ttl}
	Ttl *ComputeBackendServiceConsistentHashHttpCookieTtl `field:"optional" json:"ttl" yaml:"ttl"`
}

