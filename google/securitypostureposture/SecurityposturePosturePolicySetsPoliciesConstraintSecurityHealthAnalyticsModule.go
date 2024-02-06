// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package securitypostureposture


type SecurityposturePosturePolicySetsPoliciesConstraintSecurityHealthAnalyticsModule struct {
	// The name of the module eg: BIGQUERY_TABLE_CMEK_DISABLED.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.15.0/docs/resources/securityposture_posture#module_name SecurityposturePosture#module_name}
	ModuleName *string `field:"required" json:"moduleName" yaml:"moduleName"`
	// The state of enablement for the module at its level of the resource hierarchy. Possible values: ["ENABLEMENT_STATE_UNSPECIFIED", "ENABLED", "DISABLED"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.15.0/docs/resources/securityposture_posture#module_enablement_state SecurityposturePosture#module_enablement_state}
	ModuleEnablementState *string `field:"optional" json:"moduleEnablementState" yaml:"moduleEnablementState"`
}

