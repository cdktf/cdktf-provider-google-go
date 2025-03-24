// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package vmwareengineprivatecloud


type VmwareenginePrivateCloudManagementClusterStretchedClusterConfig struct {
	// Zone that will remain operational when connection between the two zones is lost.
	//
	// Specify the zone in the following format: projects/{project}/locations/{location}.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.26.0/docs/resources/vmwareengine_private_cloud#preferred_location VmwareenginePrivateCloud#preferred_location}
	PreferredLocation *string `field:"optional" json:"preferredLocation" yaml:"preferredLocation"`
	// Additional zone for a higher level of availability and load balancing. Specify the zone in the following format: projects/{project}/locations/{location}.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.26.0/docs/resources/vmwareengine_private_cloud#secondary_location VmwareenginePrivateCloud#secondary_location}
	SecondaryLocation *string `field:"optional" json:"secondaryLocation" yaml:"secondaryLocation"`
}

