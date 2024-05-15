// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package iamdenypolicy


type IamDenyPolicyRulesDenyRule struct {
	// denial_condition block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.29.1/docs/resources/iam_deny_policy#denial_condition IamDenyPolicy#denial_condition}
	DenialCondition *IamDenyPolicyRulesDenyRuleDenialCondition `field:"optional" json:"denialCondition" yaml:"denialCondition"`
	// The permissions that are explicitly denied by this rule.
	//
	// Each permission uses the format '{service-fqdn}/{resource}.{verb}',
	// where '{service-fqdn}' is the fully qualified domain name for the service. For example, 'iam.googleapis.com/roles.list'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.29.1/docs/resources/iam_deny_policy#denied_permissions IamDenyPolicy#denied_permissions}
	DeniedPermissions *[]*string `field:"optional" json:"deniedPermissions" yaml:"deniedPermissions"`
	// The identities that are prevented from using one or more permissions on Google Cloud resources.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.29.1/docs/resources/iam_deny_policy#denied_principals IamDenyPolicy#denied_principals}
	DeniedPrincipals *[]*string `field:"optional" json:"deniedPrincipals" yaml:"deniedPrincipals"`
	// Specifies the permissions that this rule excludes from the set of denied permissions given by deniedPermissions.
	//
	// If a permission appears in deniedPermissions and in exceptionPermissions then it will not be denied.
	// The excluded permissions can be specified using the same syntax as deniedPermissions.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.29.1/docs/resources/iam_deny_policy#exception_permissions IamDenyPolicy#exception_permissions}
	ExceptionPermissions *[]*string `field:"optional" json:"exceptionPermissions" yaml:"exceptionPermissions"`
	// The identities that are excluded from the deny rule, even if they are listed in the deniedPrincipals.
	//
	// For example, you could add a Google group to the deniedPrincipals, then exclude specific users who belong to that group.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.29.1/docs/resources/iam_deny_policy#exception_principals IamDenyPolicy#exception_principals}
	ExceptionPrincipals *[]*string `field:"optional" json:"exceptionPrincipals" yaml:"exceptionPrincipals"`
}

