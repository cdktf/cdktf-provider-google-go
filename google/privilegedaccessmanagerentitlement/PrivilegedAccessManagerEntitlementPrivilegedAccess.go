// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package privilegedaccessmanagerentitlement


type PrivilegedAccessManagerEntitlementPrivilegedAccess struct {
	// gcp_iam_access block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.30.0/docs/resources/privileged_access_manager_entitlement#gcp_iam_access PrivilegedAccessManagerEntitlement#gcp_iam_access}
	GcpIamAccess *PrivilegedAccessManagerEntitlementPrivilegedAccessGcpIamAccess `field:"required" json:"gcpIamAccess" yaml:"gcpIamAccess"`
}

