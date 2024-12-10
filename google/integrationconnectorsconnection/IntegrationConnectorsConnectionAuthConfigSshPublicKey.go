// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package integrationconnectorsconnection


type IntegrationConnectorsConnectionAuthConfigSshPublicKey struct {
	// The user account used to authenticate.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.13.0/docs/resources/integration_connectors_connection#username IntegrationConnectorsConnection#username}
	Username *string `field:"required" json:"username" yaml:"username"`
	// Format of SSH Client cert.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.13.0/docs/resources/integration_connectors_connection#cert_type IntegrationConnectorsConnection#cert_type}
	CertType *string `field:"optional" json:"certType" yaml:"certType"`
	// ssh_client_cert block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.13.0/docs/resources/integration_connectors_connection#ssh_client_cert IntegrationConnectorsConnection#ssh_client_cert}
	SshClientCert *IntegrationConnectorsConnectionAuthConfigSshPublicKeySshClientCert `field:"optional" json:"sshClientCert" yaml:"sshClientCert"`
	// ssh_client_cert_pass block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.13.0/docs/resources/integration_connectors_connection#ssh_client_cert_pass IntegrationConnectorsConnection#ssh_client_cert_pass}
	SshClientCertPass *IntegrationConnectorsConnectionAuthConfigSshPublicKeySshClientCertPass `field:"optional" json:"sshClientCertPass" yaml:"sshClientCertPass"`
}

