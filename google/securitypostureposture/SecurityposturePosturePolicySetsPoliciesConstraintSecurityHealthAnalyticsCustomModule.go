// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package securitypostureposture


type SecurityposturePosturePolicySetsPoliciesConstraintSecurityHealthAnalyticsCustomModule struct {
	// config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.39.0/docs/resources/securityposture_posture#config SecurityposturePosture#config}
	Config *SecurityposturePosturePolicySetsPoliciesConstraintSecurityHealthAnalyticsCustomModuleConfig `field:"required" json:"config" yaml:"config"`
	// The display name of the Security Health Analytics custom module.
	//
	// This
	// display name becomes the finding category for all findings that are
	// returned by this custom module.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.39.0/docs/resources/securityposture_posture#display_name SecurityposturePosture#display_name}
	DisplayName *string `field:"optional" json:"displayName" yaml:"displayName"`
	// The state of enablement for the module at its level of the resource hierarchy. Possible values: ["ENABLEMENT_STATE_UNSPECIFIED", "ENABLED", "DISABLED"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.39.0/docs/resources/securityposture_posture#module_enablement_state SecurityposturePosture#module_enablement_state}
	ModuleEnablementState *string `field:"optional" json:"moduleEnablementState" yaml:"moduleEnablementState"`
}

