// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package containercluster


type ContainerClusterNodeConfigContainerdConfigPrivateRegistryAccessConfigCertificateAuthorityDomainConfig struct {
	// List of fully-qualified-domain-names. IPv4s and port specification are supported.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.37.0/docs/resources/container_cluster#fqdns ContainerCluster#fqdns}
	Fqdns *[]*string `field:"required" json:"fqdns" yaml:"fqdns"`
	// gcp_secret_manager_certificate_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.37.0/docs/resources/container_cluster#gcp_secret_manager_certificate_config ContainerCluster#gcp_secret_manager_certificate_config}
	GcpSecretManagerCertificateConfig *ContainerClusterNodeConfigContainerdConfigPrivateRegistryAccessConfigCertificateAuthorityDomainConfigGcpSecretManagerCertificateConfig `field:"required" json:"gcpSecretManagerCertificateConfig" yaml:"gcpSecretManagerCertificateConfig"`
}

