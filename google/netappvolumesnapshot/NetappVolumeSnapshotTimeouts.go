// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package netappvolumesnapshot


type NetappVolumeSnapshotTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.18.0/docs/resources/netapp_volume_snapshot#create NetappVolumeSnapshot#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.18.0/docs/resources/netapp_volume_snapshot#delete NetappVolumeSnapshot#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.18.0/docs/resources/netapp_volume_snapshot#update NetappVolumeSnapshot#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

