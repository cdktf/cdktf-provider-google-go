// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package integrationsauthconfig


type IntegrationsAuthConfigClientCertificate struct {
	// The ssl certificate encoded in PEM format. This string must include the begin header and end footer lines.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.40.0/docs/resources/integrations_auth_config#encrypted_private_key IntegrationsAuthConfig#encrypted_private_key}
	EncryptedPrivateKey *string `field:"required" json:"encryptedPrivateKey" yaml:"encryptedPrivateKey"`
	// The ssl certificate encoded in PEM format. This string must include the begin header and end footer lines.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.40.0/docs/resources/integrations_auth_config#ssl_certificate IntegrationsAuthConfig#ssl_certificate}
	SslCertificate *string `field:"required" json:"sslCertificate" yaml:"sslCertificate"`
	// 'passphrase' should be left unset if private key is not encrypted.
	//
	// Note that 'passphrase' is not the password for web server, but an extra layer of security to protected private key.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.40.0/docs/resources/integrations_auth_config#passphrase IntegrationsAuthConfig#passphrase}
	Passphrase *string `field:"optional" json:"passphrase" yaml:"passphrase"`
}

