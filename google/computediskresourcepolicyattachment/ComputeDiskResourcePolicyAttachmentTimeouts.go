// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package computediskresourcepolicyattachment


type ComputeDiskResourcePolicyAttachmentTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.30.0/docs/resources/compute_disk_resource_policy_attachment#create ComputeDiskResourcePolicyAttachment#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.30.0/docs/resources/compute_disk_resource_policy_attachment#delete ComputeDiskResourcePolicyAttachment#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
}

