// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package computesecuritypolicyrule

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ComputeSecurityPolicyRuleAConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktf.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// The Action to perform when the rule is matched. The following are the valid actions:.
	//
	// * allow: allow access to target.
	//
	// * deny(STATUS): deny access to target, returns the HTTP response code specified. Valid values for STATUS are 403, 404, and 502.
	//
	// * rate_based_ban: limit client traffic to the configured threshold and ban the client if the traffic exceeds the threshold. Configure parameters for this action in RateLimitOptions. Requires rateLimitOptions to be set.
	//
	// * redirect: redirect to a different target. This can either be an internal reCAPTCHA redirect, or an external URL-based redirect via a 302 response. Parameters for this action can be configured via redirectOptions. This action is only supported in Global Security Policies of type CLOUD_ARMOR.
	//
	// * throttle: limit client traffic to the configured threshold. Configure parameters for this action in rateLimitOptions. Requires rateLimitOptions to be set for this.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.5.0/docs/resources/compute_security_policy_rule#action ComputeSecurityPolicyRuleA#action}
	Action *string `field:"required" json:"action" yaml:"action"`
	// An integer indicating the priority of a rule in the list.
	//
	// The priority must be a positive value between 0 and 2147483647.
	// Rules are evaluated from highest to lowest priority where 0 is the highest priority and 2147483647 is the lowest priority.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.5.0/docs/resources/compute_security_policy_rule#priority ComputeSecurityPolicyRuleA#priority}
	Priority *float64 `field:"required" json:"priority" yaml:"priority"`
	// The name of the security policy this rule belongs to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.5.0/docs/resources/compute_security_policy_rule#security_policy ComputeSecurityPolicyRuleA#security_policy}
	SecurityPolicy *string `field:"required" json:"securityPolicy" yaml:"securityPolicy"`
	// An optional description of this resource. Provide this property when you create the resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.5.0/docs/resources/compute_security_policy_rule#description ComputeSecurityPolicyRuleA#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.5.0/docs/resources/compute_security_policy_rule#id ComputeSecurityPolicyRuleA#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// match block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.5.0/docs/resources/compute_security_policy_rule#match ComputeSecurityPolicyRuleA#match}
	Match *ComputeSecurityPolicyRuleMatchA `field:"optional" json:"match" yaml:"match"`
	// preconfigured_waf_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.5.0/docs/resources/compute_security_policy_rule#preconfigured_waf_config ComputeSecurityPolicyRuleA#preconfigured_waf_config}
	PreconfiguredWafConfig *ComputeSecurityPolicyRulePreconfiguredWafConfig `field:"optional" json:"preconfiguredWafConfig" yaml:"preconfiguredWafConfig"`
	// If set to true, the specified action is not enforced.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.5.0/docs/resources/compute_security_policy_rule#preview ComputeSecurityPolicyRuleA#preview}
	Preview interface{} `field:"optional" json:"preview" yaml:"preview"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.5.0/docs/resources/compute_security_policy_rule#project ComputeSecurityPolicyRuleA#project}.
	Project *string `field:"optional" json:"project" yaml:"project"`
	// rate_limit_options block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.5.0/docs/resources/compute_security_policy_rule#rate_limit_options ComputeSecurityPolicyRuleA#rate_limit_options}
	RateLimitOptions *ComputeSecurityPolicyRuleRateLimitOptionsA `field:"optional" json:"rateLimitOptions" yaml:"rateLimitOptions"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.5.0/docs/resources/compute_security_policy_rule#timeouts ComputeSecurityPolicyRuleA#timeouts}
	Timeouts *ComputeSecurityPolicyRuleTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

