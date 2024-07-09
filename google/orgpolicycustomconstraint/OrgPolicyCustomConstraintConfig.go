// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package orgpolicycustomconstraint

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type OrgPolicyCustomConstraintConfig struct {
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
	// The action to take if the condition is met. Possible values: ["ALLOW", "DENY"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.37.0/docs/resources/org_policy_custom_constraint#action_type OrgPolicyCustomConstraint#action_type}
	ActionType *string `field:"required" json:"actionType" yaml:"actionType"`
	// A CEL condition that refers to a supported service resource, for example 'resource.management.autoUpgrade == false'. For details about CEL usage, see [Common Expression Language](https://cloud.google.com/resource-manager/docs/organization-policy/creating-managing-custom-constraints#common_expression_language).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.37.0/docs/resources/org_policy_custom_constraint#condition OrgPolicyCustomConstraint#condition}
	Condition *string `field:"required" json:"condition" yaml:"condition"`
	// A list of RESTful methods for which to enforce the constraint.
	//
	// Can be 'CREATE', 'UPDATE', or both. Not all Google Cloud services support both methods. To see supported methods for each service, find the service in [Supported services](https://cloud.google.com/resource-manager/docs/organization-policy/custom-constraint-supported-services).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.37.0/docs/resources/org_policy_custom_constraint#method_types OrgPolicyCustomConstraint#method_types}
	MethodTypes *[]*string `field:"required" json:"methodTypes" yaml:"methodTypes"`
	// Immutable. The name of the custom constraint. This is unique within the organization.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.37.0/docs/resources/org_policy_custom_constraint#name OrgPolicyCustomConstraint#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The parent of the resource, an organization. Format should be 'organizations/{organization_id}'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.37.0/docs/resources/org_policy_custom_constraint#parent OrgPolicyCustomConstraint#parent}
	Parent *string `field:"required" json:"parent" yaml:"parent"`
	// Immutable.
	//
	// The fully qualified name of the Google Cloud REST resource containing the object and field you want to restrict. For example, 'container.googleapis.com/NodePool'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.37.0/docs/resources/org_policy_custom_constraint#resource_types OrgPolicyCustomConstraint#resource_types}
	ResourceTypes *[]*string `field:"required" json:"resourceTypes" yaml:"resourceTypes"`
	// A human-friendly description of the constraint to display as an error message when the policy is violated.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.37.0/docs/resources/org_policy_custom_constraint#description OrgPolicyCustomConstraint#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// A human-friendly name for the constraint.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.37.0/docs/resources/org_policy_custom_constraint#display_name OrgPolicyCustomConstraint#display_name}
	DisplayName *string `field:"optional" json:"displayName" yaml:"displayName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.37.0/docs/resources/org_policy_custom_constraint#id OrgPolicyCustomConstraint#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.37.0/docs/resources/org_policy_custom_constraint#timeouts OrgPolicyCustomConstraint#timeouts}
	Timeouts *OrgPolicyCustomConstraintTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

