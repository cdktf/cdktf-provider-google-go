// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package apigeesecurityaction

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ApigeeSecurityActionConfig struct {
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
	// condition_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.48.0/docs/resources/apigee_security_action#condition_config ApigeeSecurityAction#condition_config}
	ConditionConfig *ApigeeSecurityActionConditionConfig `field:"required" json:"conditionConfig" yaml:"conditionConfig"`
	// The Apigee environment that this security action applies to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.48.0/docs/resources/apigee_security_action#env_id ApigeeSecurityAction#env_id}
	EnvId *string `field:"required" json:"envId" yaml:"envId"`
	// The organization that this security action applies to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.48.0/docs/resources/apigee_security_action#org_id ApigeeSecurityAction#org_id}
	OrgId *string `field:"required" json:"orgId" yaml:"orgId"`
	// The ID to use for the SecurityAction, which will become the final component of the action's resource name.
	//
	// This value should be 0-61 characters, and valid format is (^a-z?$).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.48.0/docs/resources/apigee_security_action#security_action_id ApigeeSecurityAction#security_action_id}
	SecurityActionId *string `field:"required" json:"securityActionId" yaml:"securityActionId"`
	// Only an ENABLED SecurityAction is enforced.
	//
	// An ENABLED SecurityAction past its expiration time will not be enforced. Possible values: ["ENABLED", "DISABLED"]
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.48.0/docs/resources/apigee_security_action#state ApigeeSecurityAction#state}
	State *string `field:"required" json:"state" yaml:"state"`
	// allow block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.48.0/docs/resources/apigee_security_action#allow ApigeeSecurityAction#allow}
	Allow *ApigeeSecurityActionAllow `field:"optional" json:"allow" yaml:"allow"`
	// If unset, this would apply to all proxies in the environment.
	//
	// If set, this action is enforced only if at least one proxy in the repeated
	// list is deployed at the time of enforcement. If set, several restrictions are enforced on SecurityActions.
	// There can be at most 100 enabled actions with proxies set in an env.
	// Several other restrictions apply on conditions and are detailed later.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.48.0/docs/resources/apigee_security_action#api_proxies ApigeeSecurityAction#api_proxies}
	ApiProxies *[]*string `field:"optional" json:"apiProxies" yaml:"apiProxies"`
	// deny block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.48.0/docs/resources/apigee_security_action#deny ApigeeSecurityAction#deny}
	Deny *ApigeeSecurityActionDeny `field:"optional" json:"deny" yaml:"deny"`
	// An optional user provided description of the SecurityAction.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.48.0/docs/resources/apigee_security_action#description ApigeeSecurityAction#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The expiration for this SecurityAction.
	//
	// Uses RFC 3339, where generated output will always be Z-normalized and uses 0, 3, 6 or 9
	// fractional digits. Offsets other than "Z" are also accepted.
	// Examples: "2014-10-02T15:01:23Z", "2014-10-02T15:01:23.045123456Z" or "2014-10-02T15:01:23+05:30".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.48.0/docs/resources/apigee_security_action#expire_time ApigeeSecurityAction#expire_time}
	ExpireTime *string `field:"optional" json:"expireTime" yaml:"expireTime"`
	// flag block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.48.0/docs/resources/apigee_security_action#flag ApigeeSecurityAction#flag}
	Flag *ApigeeSecurityActionFlag `field:"optional" json:"flag" yaml:"flag"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.48.0/docs/resources/apigee_security_action#id ApigeeSecurityAction#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.48.0/docs/resources/apigee_security_action#timeouts ApigeeSecurityAction#timeouts}
	Timeouts *ApigeeSecurityActionTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
	// The TTL for this SecurityAction. A duration in seconds with up to nine fractional digits, ending with 's'. Example: "3.5s".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.48.0/docs/resources/apigee_security_action#ttl ApigeeSecurityAction#ttl}
	Ttl *string `field:"optional" json:"ttl" yaml:"ttl"`
}

