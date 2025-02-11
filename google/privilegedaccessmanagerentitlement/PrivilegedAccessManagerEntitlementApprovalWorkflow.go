// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package privilegedaccessmanagerentitlement


type PrivilegedAccessManagerEntitlementApprovalWorkflow struct {
	// manual_approvals block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.20.0/docs/resources/privileged_access_manager_entitlement#manual_approvals PrivilegedAccessManagerEntitlement#manual_approvals}
	ManualApprovals *PrivilegedAccessManagerEntitlementApprovalWorkflowManualApprovals `field:"required" json:"manualApprovals" yaml:"manualApprovals"`
}

