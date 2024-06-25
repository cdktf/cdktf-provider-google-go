// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package apigeetargetserver


type ApigeeTargetServerSSlInfoCommonName struct {
	// The TLS Common Name string of the certificate.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.35.0/docs/resources/apigee_target_server#value ApigeeTargetServer#value}
	Value *string `field:"optional" json:"value" yaml:"value"`
	// Indicates whether the cert should be matched against as a wildcard cert.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.35.0/docs/resources/apigee_target_server#wildcard_match ApigeeTargetServer#wildcard_match}
	WildcardMatch interface{} `field:"optional" json:"wildcardMatch" yaml:"wildcardMatch"`
}

