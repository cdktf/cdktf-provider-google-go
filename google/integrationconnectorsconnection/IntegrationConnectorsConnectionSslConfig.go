// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package integrationconnectorsconnection


type IntegrationConnectorsConnectionSslConfig struct {
	// Enum for controlling the SSL Type (TLS/MTLS) Possible values: ["TLS", "MTLS"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.11.1/docs/resources/integration_connectors_connection#type IntegrationConnectorsConnection#type}
	Type *string `field:"required" json:"type" yaml:"type"`
	// additional_variable block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.11.1/docs/resources/integration_connectors_connection#additional_variable IntegrationConnectorsConnection#additional_variable}
	AdditionalVariable interface{} `field:"optional" json:"additionalVariable" yaml:"additionalVariable"`
	// client_certificate block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.11.1/docs/resources/integration_connectors_connection#client_certificate IntegrationConnectorsConnection#client_certificate}
	ClientCertificate *IntegrationConnectorsConnectionSslConfigClientCertificate `field:"optional" json:"clientCertificate" yaml:"clientCertificate"`
	// Type of Client Cert (PEM/JKS/.. etc.) Possible values: ["PEM"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.11.1/docs/resources/integration_connectors_connection#client_cert_type IntegrationConnectorsConnection#client_cert_type}
	ClientCertType *string `field:"optional" json:"clientCertType" yaml:"clientCertType"`
	// client_private_key block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.11.1/docs/resources/integration_connectors_connection#client_private_key IntegrationConnectorsConnection#client_private_key}
	ClientPrivateKey *IntegrationConnectorsConnectionSslConfigClientPrivateKey `field:"optional" json:"clientPrivateKey" yaml:"clientPrivateKey"`
	// client_private_key_pass block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.11.1/docs/resources/integration_connectors_connection#client_private_key_pass IntegrationConnectorsConnection#client_private_key_pass}
	ClientPrivateKeyPass *IntegrationConnectorsConnectionSslConfigClientPrivateKeyPass `field:"optional" json:"clientPrivateKeyPass" yaml:"clientPrivateKeyPass"`
	// private_server_certificate block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.11.1/docs/resources/integration_connectors_connection#private_server_certificate IntegrationConnectorsConnection#private_server_certificate}
	PrivateServerCertificate *IntegrationConnectorsConnectionSslConfigPrivateServerCertificate `field:"optional" json:"privateServerCertificate" yaml:"privateServerCertificate"`
	// Type of Server Cert (PEM/JKS/.. etc.) Possible values: ["PEM"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.11.1/docs/resources/integration_connectors_connection#server_cert_type IntegrationConnectorsConnection#server_cert_type}
	ServerCertType *string `field:"optional" json:"serverCertType" yaml:"serverCertType"`
	// Enum for Trust Model Possible values: ["PUBLIC", "PRIVATE", "INSECURE"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.11.1/docs/resources/integration_connectors_connection#trust_model IntegrationConnectorsConnection#trust_model}
	TrustModel *string `field:"optional" json:"trustModel" yaml:"trustModel"`
	// Bool for enabling SSL.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.11.1/docs/resources/integration_connectors_connection#use_ssl IntegrationConnectorsConnection#use_ssl}
	UseSsl interface{} `field:"optional" json:"useSsl" yaml:"useSsl"`
}

