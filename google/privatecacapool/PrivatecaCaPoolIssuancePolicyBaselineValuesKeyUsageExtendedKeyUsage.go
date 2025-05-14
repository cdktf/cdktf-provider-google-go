// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package privatecacapool


type PrivatecaCaPoolIssuancePolicyBaselineValuesKeyUsageExtendedKeyUsage struct {
	// Corresponds to OID 1.3.6.1.5.5.7.3.2. Officially described as "TLS WWW client authentication", though regularly used for non-WWW TLS.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.35.0/docs/resources/privateca_ca_pool#client_auth PrivatecaCaPool#client_auth}
	ClientAuth interface{} `field:"optional" json:"clientAuth" yaml:"clientAuth"`
	// Corresponds to OID 1.3.6.1.5.5.7.3.3. Officially described as "Signing of downloadable executable code client authentication".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.35.0/docs/resources/privateca_ca_pool#code_signing PrivatecaCaPool#code_signing}
	CodeSigning interface{} `field:"optional" json:"codeSigning" yaml:"codeSigning"`
	// Corresponds to OID 1.3.6.1.5.5.7.3.4. Officially described as "Email protection".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.35.0/docs/resources/privateca_ca_pool#email_protection PrivatecaCaPool#email_protection}
	EmailProtection interface{} `field:"optional" json:"emailProtection" yaml:"emailProtection"`
	// Corresponds to OID 1.3.6.1.5.5.7.3.9. Officially described as "Signing OCSP responses".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.35.0/docs/resources/privateca_ca_pool#ocsp_signing PrivatecaCaPool#ocsp_signing}
	OcspSigning interface{} `field:"optional" json:"ocspSigning" yaml:"ocspSigning"`
	// Corresponds to OID 1.3.6.1.5.5.7.3.1. Officially described as "TLS WWW server authentication", though regularly used for non-WWW TLS.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.35.0/docs/resources/privateca_ca_pool#server_auth PrivatecaCaPool#server_auth}
	ServerAuth interface{} `field:"optional" json:"serverAuth" yaml:"serverAuth"`
	// Corresponds to OID 1.3.6.1.5.5.7.3.8. Officially described as "Binding the hash of an object to a time".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.35.0/docs/resources/privateca_ca_pool#time_stamping PrivatecaCaPool#time_stamping}
	TimeStamping interface{} `field:"optional" json:"timeStamping" yaml:"timeStamping"`
}

