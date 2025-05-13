// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package networksecurityauthzpolicy

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type NetworkSecurityAuthzPolicyConfig struct {
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
	// When the action is CUSTOM, customProvider must be specified.
	//
	// When the action is ALLOW, only requests matching the policy will be allowed.
	// When the action is DENY, only requests matching the policy will be denied.
	//
	// When a request arrives, the policies are evaluated in the following order:
	// 1. If there is a CUSTOM policy that matches the request, the CUSTOM policy is evaluated using the custom authorization providers and the request is denied if the provider rejects the request.
	// 2. If there are any DENY policies that match the request, the request is denied.
	// 3. If there are no ALLOW policies for the resource or if any of the ALLOW policies match the request, the request is allowed.
	// 4. Else the request is denied by default if none of the configured AuthzPolicies with ALLOW action match the request. Possible values: ["ALLOW", "DENY", "CUSTOM"]
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.34.1/docs/resources/network_security_authz_policy#action NetworkSecurityAuthzPolicy#action}
	Action *string `field:"required" json:"action" yaml:"action"`
	// The location of the resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.34.1/docs/resources/network_security_authz_policy#location NetworkSecurityAuthzPolicy#location}
	Location *string `field:"required" json:"location" yaml:"location"`
	// Identifier. Name of the AuthzPolicy resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.34.1/docs/resources/network_security_authz_policy#name NetworkSecurityAuthzPolicy#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// target block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.34.1/docs/resources/network_security_authz_policy#target NetworkSecurityAuthzPolicy#target}
	Target *NetworkSecurityAuthzPolicyTarget `field:"required" json:"target" yaml:"target"`
	// custom_provider block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.34.1/docs/resources/network_security_authz_policy#custom_provider NetworkSecurityAuthzPolicy#custom_provider}
	CustomProvider *NetworkSecurityAuthzPolicyCustomProvider `field:"optional" json:"customProvider" yaml:"customProvider"`
	// A human-readable description of the resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.34.1/docs/resources/network_security_authz_policy#description NetworkSecurityAuthzPolicy#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// http_rules block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.34.1/docs/resources/network_security_authz_policy#http_rules NetworkSecurityAuthzPolicy#http_rules}
	HttpRules interface{} `field:"optional" json:"httpRules" yaml:"httpRules"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.34.1/docs/resources/network_security_authz_policy#id NetworkSecurityAuthzPolicy#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Set of labels associated with the AuthzExtension resource.
	//
	// **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
	// Please refer to the field 'effective_labels' for all of the labels present on the resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.34.1/docs/resources/network_security_authz_policy#labels NetworkSecurityAuthzPolicy#labels}
	Labels *map[string]*string `field:"optional" json:"labels" yaml:"labels"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.34.1/docs/resources/network_security_authz_policy#project NetworkSecurityAuthzPolicy#project}.
	Project *string `field:"optional" json:"project" yaml:"project"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.34.1/docs/resources/network_security_authz_policy#timeouts NetworkSecurityAuthzPolicy#timeouts}
	Timeouts *NetworkSecurityAuthzPolicyTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

