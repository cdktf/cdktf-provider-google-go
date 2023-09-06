// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package computeresourcepolicy


type ComputeResourcePolicySnapshotSchedulePolicy struct {
	// schedule block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.81.0/docs/resources/compute_resource_policy#schedule ComputeResourcePolicy#schedule}
	Schedule *ComputeResourcePolicySnapshotSchedulePolicySchedule `field:"required" json:"schedule" yaml:"schedule"`
	// retention_policy block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.81.0/docs/resources/compute_resource_policy#retention_policy ComputeResourcePolicy#retention_policy}
	RetentionPolicy *ComputeResourcePolicySnapshotSchedulePolicyRetentionPolicy `field:"optional" json:"retentionPolicy" yaml:"retentionPolicy"`
	// snapshot_properties block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.81.0/docs/resources/compute_resource_policy#snapshot_properties ComputeResourcePolicy#snapshot_properties}
	SnapshotProperties *ComputeResourcePolicySnapshotSchedulePolicySnapshotProperties `field:"optional" json:"snapshotProperties" yaml:"snapshotProperties"`
}

