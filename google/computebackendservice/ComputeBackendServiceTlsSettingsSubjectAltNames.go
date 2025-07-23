// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package computebackendservice


type ComputeBackendServiceTlsSettingsSubjectAltNames struct {
	// The SAN specified as a DNS Name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/compute_backend_service#dns_name ComputeBackendService#dns_name}
	DnsName *string `field:"optional" json:"dnsName" yaml:"dnsName"`
	// The SAN specified as a URI.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/compute_backend_service#uniform_resource_identifier ComputeBackendService#uniform_resource_identifier}
	UniformResourceIdentifier *string `field:"optional" json:"uniformResourceIdentifier" yaml:"uniformResourceIdentifier"`
}

