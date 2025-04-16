// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package privilegedaccessmanagerentitlement


type PrivilegedAccessManagerEntitlementApprovalWorkflowManualApprovalsStepsApprovers struct {
	// Users who are being allowed for the operation.
	//
	// Each entry should be a valid v1 IAM Principal Identifier. Format for these is documented at: https://cloud.google.com/iam/docs/principal-identifiers#v1
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.30.0/docs/resources/privileged_access_manager_entitlement#principals PrivilegedAccessManagerEntitlement#principals}
	Principals *[]*string `field:"required" json:"principals" yaml:"principals"`
}

