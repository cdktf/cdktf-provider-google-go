// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package computesecuritypolicyrule


type ComputeSecurityPolicyRulePreconfiguredWafConfigExclusionRequestCookie struct {
	// You can specify an exact match or a partial match by using a field operator and a field value.
	//
	// Available options:
	// EQUALS: The operator matches if the field value equals the specified value.
	// STARTS_WITH: The operator matches if the field value starts with the specified value.
	// ENDS_WITH: The operator matches if the field value ends with the specified value.
	// CONTAINS: The operator matches if the field value contains the specified value.
	// EQUALS_ANY: The operator matches if the field value is any value.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.40.0/docs/resources/compute_security_policy_rule#operator ComputeSecurityPolicyRuleA#operator}
	Operator *string `field:"required" json:"operator" yaml:"operator"`
	// A request field matching the specified value will be excluded from inspection during preconfigured WAF evaluation.
	//
	// The field value must be given if the field operator is not EQUALS_ANY, and cannot be given if the field operator is EQUALS_ANY.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.40.0/docs/resources/compute_security_policy_rule#value ComputeSecurityPolicyRuleA#value}
	Value *string `field:"optional" json:"value" yaml:"value"`
}

