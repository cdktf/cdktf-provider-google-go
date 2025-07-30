// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package networksecurityclienttlspolicy


type NetworkSecurityClientTlsPolicyClientCertificateCertificateProviderInstance struct {
	// Plugin instance name, used to locate and load CertificateProvider instance configuration.
	//
	// Set to "google_cloud_private_spiffe" to use Certificate Authority Service certificate provider instance.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.46.0/docs/resources/network_security_client_tls_policy#plugin_instance NetworkSecurityClientTlsPolicy#plugin_instance}
	PluginInstance *string `field:"required" json:"pluginInstance" yaml:"pluginInstance"`
}

