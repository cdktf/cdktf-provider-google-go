// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package containercluster


type ContainerClusterNodePoolNodeConfigContainerdConfigPrivateRegistryAccessConfig struct {
	// Whether or not private registries are configured.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.18.0/docs/resources/container_cluster#enabled ContainerCluster#enabled}
	Enabled interface{} `field:"required" json:"enabled" yaml:"enabled"`
	// certificate_authority_domain_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.18.0/docs/resources/container_cluster#certificate_authority_domain_config ContainerCluster#certificate_authority_domain_config}
	CertificateAuthorityDomainConfig interface{} `field:"optional" json:"certificateAuthorityDomainConfig" yaml:"certificateAuthorityDomainConfig"`
}

