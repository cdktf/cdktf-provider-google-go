// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package containernodepool


type ContainerNodePoolNodeConfigContainerdConfigPrivateRegistryAccessConfigCertificateAuthorityDomainConfig struct {
	// List of fully-qualified-domain-names. IPv4s and port specification are supported.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.22.0/docs/resources/container_node_pool#fqdns ContainerNodePool#fqdns}
	Fqdns *[]*string `field:"required" json:"fqdns" yaml:"fqdns"`
	// gcp_secret_manager_certificate_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.22.0/docs/resources/container_node_pool#gcp_secret_manager_certificate_config ContainerNodePool#gcp_secret_manager_certificate_config}
	GcpSecretManagerCertificateConfig *ContainerNodePoolNodeConfigContainerdConfigPrivateRegistryAccessConfigCertificateAuthorityDomainConfigGcpSecretManagerCertificateConfig `field:"required" json:"gcpSecretManagerCertificateConfig" yaml:"gcpSecretManagerCertificateConfig"`
}

