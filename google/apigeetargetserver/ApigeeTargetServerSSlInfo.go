// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package apigeetargetserver


type ApigeeTargetServerSSlInfo struct {
	// Enables TLS. If false, neither one-way nor two-way TLS will be enabled.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.28.0/docs/resources/apigee_target_server#enabled ApigeeTargetServer#enabled}
	Enabled interface{} `field:"required" json:"enabled" yaml:"enabled"`
	// The SSL/TLS cipher suites to be used.
	//
	// For programmable proxies, it must be one of the cipher suite names listed in: http://docs.oracle.com/javase/8/docs/technotes/guides/security/StandardNames.html#ciphersuites. For configurable proxies, it must follow the configuration specified in: https://commondatastorage.googleapis.com/chromium-boringssl-docs/ssl.h.html#Cipher-suite-configuration. This setting has no effect for configurable proxies when negotiating TLS 1.3.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.28.0/docs/resources/apigee_target_server#ciphers ApigeeTargetServer#ciphers}
	Ciphers *[]*string `field:"optional" json:"ciphers" yaml:"ciphers"`
	// Enables two-way TLS.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.28.0/docs/resources/apigee_target_server#client_auth_enabled ApigeeTargetServer#client_auth_enabled}
	ClientAuthEnabled interface{} `field:"optional" json:"clientAuthEnabled" yaml:"clientAuthEnabled"`
	// common_name block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.28.0/docs/resources/apigee_target_server#common_name ApigeeTargetServer#common_name}
	CommonName *ApigeeTargetServerSSlInfoCommonName `field:"optional" json:"commonName" yaml:"commonName"`
	// If true, Edge ignores TLS certificate errors.
	//
	// Valid when configuring TLS for target servers and target endpoints, and when configuring virtual hosts that use 2-way TLS. When used with a target endpoint/target server, if the backend system uses SNI and returns a cert with a subject Distinguished Name (DN) that does not match the hostname, there is no way to ignore the error and the connection fails.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.28.0/docs/resources/apigee_target_server#ignore_validation_errors ApigeeTargetServer#ignore_validation_errors}
	IgnoreValidationErrors interface{} `field:"optional" json:"ignoreValidationErrors" yaml:"ignoreValidationErrors"`
	// Required if clientAuthEnabled is true. The resource ID for the alias containing the private key and cert.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.28.0/docs/resources/apigee_target_server#key_alias ApigeeTargetServer#key_alias}
	KeyAlias *string `field:"optional" json:"keyAlias" yaml:"keyAlias"`
	// Required if clientAuthEnabled is true. The resource ID of the keystore.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.28.0/docs/resources/apigee_target_server#key_store ApigeeTargetServer#key_store}
	KeyStore *string `field:"optional" json:"keyStore" yaml:"keyStore"`
	// The TLS versioins to be used.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.28.0/docs/resources/apigee_target_server#protocols ApigeeTargetServer#protocols}
	Protocols *[]*string `field:"optional" json:"protocols" yaml:"protocols"`
	// The resource ID of the truststore.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.28.0/docs/resources/apigee_target_server#trust_store ApigeeTargetServer#trust_store}
	TrustStore *string `field:"optional" json:"trustStore" yaml:"trustStore"`
}

