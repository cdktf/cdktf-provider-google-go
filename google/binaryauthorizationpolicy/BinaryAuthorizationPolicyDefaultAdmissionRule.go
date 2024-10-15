// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package binaryauthorizationpolicy


type BinaryAuthorizationPolicyDefaultAdmissionRule struct {
	// The action when a pod creation is denied by the admission rule. Possible values: ["ENFORCED_BLOCK_AND_AUDIT_LOG", "DRYRUN_AUDIT_LOG_ONLY"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.7.0/docs/resources/binary_authorization_policy#enforcement_mode BinaryAuthorizationPolicy#enforcement_mode}
	EnforcementMode *string `field:"required" json:"enforcementMode" yaml:"enforcementMode"`
	// How this admission rule will be evaluated. Possible values: ["ALWAYS_ALLOW", "REQUIRE_ATTESTATION", "ALWAYS_DENY"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.7.0/docs/resources/binary_authorization_policy#evaluation_mode BinaryAuthorizationPolicy#evaluation_mode}
	EvaluationMode *string `field:"required" json:"evaluationMode" yaml:"evaluationMode"`
	// The resource names of the attestors that must attest to a container image.
	//
	// If the attestor is in a different project from the
	// policy, it should be specified in the format 'projects/* /attestors/*'.
	// Each attestor must exist before a policy can reference it. To add an
	// attestor to a policy the principal issuing the policy change
	// request must be able to read the attestor resource.
	//
	// Note: this field must be non-empty when the evaluation_mode field
	// specifies REQUIRE_ATTESTATION, otherwise it must be empty.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.7.0/docs/resources/binary_authorization_policy#require_attestations_by BinaryAuthorizationPolicy#require_attestations_by}
	//
	// Note: The above comment contained a comment block ending sequence (* followed by /). We have introduced a space between to prevent syntax errors. Please ignore the space.
	RequireAttestationsBy *[]*string `field:"optional" json:"requireAttestationsBy" yaml:"requireAttestationsBy"`
}

