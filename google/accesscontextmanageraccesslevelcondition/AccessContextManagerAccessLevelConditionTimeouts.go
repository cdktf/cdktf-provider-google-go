// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package accesscontextmanageraccesslevelcondition


type AccessContextManagerAccessLevelConditionTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.26.0/docs/resources/access_context_manager_access_level_condition#create AccessContextManagerAccessLevelCondition#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.26.0/docs/resources/access_context_manager_access_level_condition#delete AccessContextManagerAccessLevelCondition#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
}

