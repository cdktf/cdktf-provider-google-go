// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package computebackendservice


type ComputeBackendServiceTlsSettings struct {
	// Reference to the BackendAuthenticationConfig resource from the networksecurity.googleapis.com namespace. Can be used in authenticating TLS connections to the backend, as specified by the authenticationMode field. Can only be specified if authenticationMode is not NONE.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.1/docs/resources/compute_backend_service#authentication_config ComputeBackendService#authentication_config}
	AuthenticationConfig *string `field:"optional" json:"authenticationConfig" yaml:"authenticationConfig"`
	// Server Name Indication - see RFC3546 section 3.1. If set, the load balancer sends this string as the SNI hostname in the TLS connection to the backend, and requires that this string match a Subject Alternative Name (SAN) in the backend's server certificate. With a Regional Internet NEG backend, if the SNI is specified here, the load balancer uses it regardless of whether the Regional Internet NEG is specified with FQDN or IP address and port.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.1/docs/resources/compute_backend_service#sni ComputeBackendService#sni}
	Sni *string `field:"optional" json:"sni" yaml:"sni"`
	// subject_alt_names block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.1/docs/resources/compute_backend_service#subject_alt_names ComputeBackendService#subject_alt_names}
	SubjectAltNames interface{} `field:"optional" json:"subjectAltNames" yaml:"subjectAltNames"`
}

