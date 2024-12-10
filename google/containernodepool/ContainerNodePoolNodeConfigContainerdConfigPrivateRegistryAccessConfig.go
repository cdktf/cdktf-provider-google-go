// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package containernodepool


type ContainerNodePoolNodeConfigContainerdConfigPrivateRegistryAccessConfig struct {
	// Whether or not private registries are configured.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.13.0/docs/resources/container_node_pool#enabled ContainerNodePool#enabled}
	Enabled interface{} `field:"required" json:"enabled" yaml:"enabled"`
	// certificate_authority_domain_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.13.0/docs/resources/container_node_pool#certificate_authority_domain_config ContainerNodePool#certificate_authority_domain_config}
	CertificateAuthorityDomainConfig interface{} `field:"optional" json:"certificateAuthorityDomainConfig" yaml:"certificateAuthorityDomainConfig"`
}

