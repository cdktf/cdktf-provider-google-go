// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package computeinterconnectattachmentgroup


type ComputeInterconnectAttachmentGroupIntent struct {
	// Which SLA the user intends this group to support. Possible values: ["PRODUCTION_NON_CRITICAL", "PRODUCTION_CRITICAL", "NO_SLA", "AVAILABILITY_SLA_UNSPECIFIED"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.46.0/docs/resources/compute_interconnect_attachment_group#availability_sla ComputeInterconnectAttachmentGroup#availability_sla}
	AvailabilitySla *string `field:"optional" json:"availabilitySla" yaml:"availabilitySla"`
}

