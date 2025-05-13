// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package computeresourcepolicy


type ComputeResourcePolicyInstanceSchedulePolicyVmStartSchedule struct {
	// Specifies the frequency for the operation, using the unix-cron format.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.34.1/docs/resources/compute_resource_policy#schedule ComputeResourcePolicy#schedule}
	Schedule *string `field:"required" json:"schedule" yaml:"schedule"`
}

