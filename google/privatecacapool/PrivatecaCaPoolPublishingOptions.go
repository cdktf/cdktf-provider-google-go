// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package privatecacapool


type PrivatecaCaPoolPublishingOptions struct {
	// When true, publishes each CertificateAuthority's CA certificate and includes its URL in the "Authority Information Access" X.509 extension in all issued Certificates. If this is false, the CA certificate will not be published and the corresponding X.509 extension will not be written in issued certificates.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.29.0/docs/resources/privateca_ca_pool#publish_ca_cert PrivatecaCaPool#publish_ca_cert}
	PublishCaCert interface{} `field:"required" json:"publishCaCert" yaml:"publishCaCert"`
	// When true, publishes each CertificateAuthority's CRL and includes its URL in the "CRL Distribution Points" X.509 extension in all issued Certificates. If this is false, CRLs will not be published and the corresponding X.509 extension will not be written in issued certificates. CRLs will expire 7 days from their creation. However, we will rebuild daily. CRLs are also rebuilt shortly after a certificate is revoked.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.29.0/docs/resources/privateca_ca_pool#publish_crl PrivatecaCaPool#publish_crl}
	PublishCrl interface{} `field:"required" json:"publishCrl" yaml:"publishCrl"`
	// Specifies the encoding format of each CertificateAuthority's CA certificate and CRLs.
	//
	// If this is omitted, CA certificates and CRLs
	// will be published in PEM. Possible values: ["PEM", "DER"]
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.29.0/docs/resources/privateca_ca_pool#encoding_format PrivatecaCaPool#encoding_format}
	EncodingFormat *string `field:"optional" json:"encodingFormat" yaml:"encodingFormat"`
}

