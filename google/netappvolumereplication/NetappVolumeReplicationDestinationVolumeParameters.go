// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package netappvolumereplication


type NetappVolumeReplicationDestinationVolumeParameters struct {
	// Name of an existing storage pool for the destination volume with format: 'projects/{{project}}/locations/{{location}}/storagePools/{{poolId}}'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.11.1/docs/resources/netapp_volume_replication#storage_pool NetappVolumeReplication#storage_pool}
	StoragePool *string `field:"required" json:"storagePool" yaml:"storagePool"`
	// Description for the destination volume.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.11.1/docs/resources/netapp_volume_replication#description NetappVolumeReplication#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Share name for destination volume. If not specified, name of source volume's share name will be used.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.11.1/docs/resources/netapp_volume_replication#share_name NetappVolumeReplication#share_name}
	ShareName *string `field:"optional" json:"shareName" yaml:"shareName"`
	// Name for the destination volume to be created.
	//
	// If not specified, the name of the source volume will be used.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.11.1/docs/resources/netapp_volume_replication#volume_id NetappVolumeReplication#volume_id}
	VolumeId *string `field:"optional" json:"volumeId" yaml:"volumeId"`
}

