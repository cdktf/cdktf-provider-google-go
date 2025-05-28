// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package netappvolume


type NetappVolumeHybridReplicationParameters struct {
	// Optional.
	//
	// Name of source cluster location associated with the Hybrid replication. This is a free-form field for the display purpose only.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.37.0/docs/resources/netapp_volume#cluster_location NetappVolume#cluster_location}
	ClusterLocation *string `field:"optional" json:"clusterLocation" yaml:"clusterLocation"`
	// Optional. Description of the replication.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.37.0/docs/resources/netapp_volume#description NetappVolume#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Optional.
	//
	// Labels to be added to the replication as the key value pairs.
	// An object containing a list of "key": value pairs. Example: { "name": "wrench", "mass": "1.3kg", "count": "3" }.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.37.0/docs/resources/netapp_volume#labels NetappVolume#labels}
	Labels *map[string]*string `field:"optional" json:"labels" yaml:"labels"`
	// Required. Name of the user's local source cluster to be peered with the destination cluster.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.37.0/docs/resources/netapp_volume#peer_cluster_name NetappVolume#peer_cluster_name}
	PeerClusterName *string `field:"optional" json:"peerClusterName" yaml:"peerClusterName"`
	// Required. List of node ip addresses to be peered with.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.37.0/docs/resources/netapp_volume#peer_ip_addresses NetappVolume#peer_ip_addresses}
	PeerIpAddresses *string `field:"optional" json:"peerIpAddresses" yaml:"peerIpAddresses"`
	// Required. Name of the user's local source vserver svm to be peered with the destination vserver svm.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.37.0/docs/resources/netapp_volume#peer_svm_name NetappVolume#peer_svm_name}
	PeerSvmName *string `field:"optional" json:"peerSvmName" yaml:"peerSvmName"`
	// Required. Name of the user's local source volume to be peered with the destination volume.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.37.0/docs/resources/netapp_volume#peer_volume_name NetappVolume#peer_volume_name}
	PeerVolumeName *string `field:"optional" json:"peerVolumeName" yaml:"peerVolumeName"`
	// Required. Desired name for the replication of this volume.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.37.0/docs/resources/netapp_volume#replication NetappVolume#replication}
	Replication *string `field:"optional" json:"replication" yaml:"replication"`
}

