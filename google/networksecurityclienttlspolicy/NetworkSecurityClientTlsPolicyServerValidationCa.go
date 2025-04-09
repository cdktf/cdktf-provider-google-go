// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package networksecurityclienttlspolicy


type NetworkSecurityClientTlsPolicyServerValidationCa struct {
	// certificate_provider_instance block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.29.0/docs/resources/network_security_client_tls_policy#certificate_provider_instance NetworkSecurityClientTlsPolicy#certificate_provider_instance}
	CertificateProviderInstance *NetworkSecurityClientTlsPolicyServerValidationCaCertificateProviderInstance `field:"optional" json:"certificateProviderInstance" yaml:"certificateProviderInstance"`
	// grpc_endpoint block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.29.0/docs/resources/network_security_client_tls_policy#grpc_endpoint NetworkSecurityClientTlsPolicy#grpc_endpoint}
	GrpcEndpoint *NetworkSecurityClientTlsPolicyServerValidationCaGrpcEndpoint `field:"optional" json:"grpcEndpoint" yaml:"grpcEndpoint"`
}

