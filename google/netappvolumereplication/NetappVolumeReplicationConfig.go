// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package netappvolumereplication

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type NetappVolumeReplicationConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktf.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// Name of region for this resource. The resource needs to be created in the region of the destination volume.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.28.0/docs/resources/netapp_volume_replication#location NetappVolumeReplication#location}
	Location *string `field:"required" json:"location" yaml:"location"`
	// The name of the replication. Needs to be unique per location.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.28.0/docs/resources/netapp_volume_replication#name NetappVolumeReplication#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Specifies the replication interval. Possible values: ["EVERY_10_MINUTES", "HOURLY", "DAILY"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.28.0/docs/resources/netapp_volume_replication#replication_schedule NetappVolumeReplication#replication_schedule}
	ReplicationSchedule *string `field:"required" json:"replicationSchedule" yaml:"replicationSchedule"`
	// The name of the existing source volume.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.28.0/docs/resources/netapp_volume_replication#volume_name NetappVolumeReplication#volume_name}
	VolumeName *string `field:"required" json:"volumeName" yaml:"volumeName"`
	// A destination volume is created as part of replication creation.
	//
	// The destination volume will not became
	// under Terraform management unless you import it manually. If you delete the replication, this volume
	// will remain.
	// Setting this parameter to true will delete the *current* destination volume when destroying the
	// replication. If you reversed the replication direction, this will be your former source volume!
	// For production use, it is recommended to keep this parameter false to avoid accidental volume
	// deletion. Handle with care. Default is false.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.28.0/docs/resources/netapp_volume_replication#delete_destination_volume NetappVolumeReplication#delete_destination_volume}
	DeleteDestinationVolume interface{} `field:"optional" json:"deleteDestinationVolume" yaml:"deleteDestinationVolume"`
	// An description of this resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.28.0/docs/resources/netapp_volume_replication#description NetappVolumeReplication#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// destination_volume_parameters block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.28.0/docs/resources/netapp_volume_replication#destination_volume_parameters NetappVolumeReplication#destination_volume_parameters}
	DestinationVolumeParameters *NetappVolumeReplicationDestinationVolumeParameters `field:"optional" json:"destinationVolumeParameters" yaml:"destinationVolumeParameters"`
	// Only replications with mirror_state=MIRRORED can be stopped.
	//
	// A replication in mirror_state=TRANSFERRING
	// currently receives an update and stopping the update might be undesirable. Set this parameter to true
	// to stop anyway. All data transferred to the destination will be discarded and content of destination
	// volume will remain at the state of the last successful update. Default is false.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.28.0/docs/resources/netapp_volume_replication#force_stopping NetappVolumeReplication#force_stopping}
	ForceStopping interface{} `field:"optional" json:"forceStopping" yaml:"forceStopping"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.28.0/docs/resources/netapp_volume_replication#id NetappVolumeReplication#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Labels as key value pairs. Example: '{ "owner": "Bob", "department": "finance", "purpose": "testing" }'.
	//
	// **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
	// Please refer to the field 'effective_labels' for all of the labels present on the resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.28.0/docs/resources/netapp_volume_replication#labels NetappVolumeReplication#labels}
	Labels *map[string]*string `field:"optional" json:"labels" yaml:"labels"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.28.0/docs/resources/netapp_volume_replication#project NetappVolumeReplication#project}.
	Project *string `field:"optional" json:"project" yaml:"project"`
	// Set to false to stop/break the mirror.
	//
	// Stopping the mirror makes the destination volume read-write
	// and act independently from the source volume.
	// Set to true to enable/resume the mirror. WARNING: Resuming a mirror overwrites any changes
	// done to the destination volume with the content of the source volume.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.28.0/docs/resources/netapp_volume_replication#replication_enabled NetappVolumeReplication#replication_enabled}
	ReplicationEnabled interface{} `field:"optional" json:"replicationEnabled" yaml:"replicationEnabled"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.28.0/docs/resources/netapp_volume_replication#timeouts NetappVolumeReplication#timeouts}
	Timeouts *NetappVolumeReplicationTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
	// Replication resource state is independent of mirror_state.
	//
	// With enough data, it can take many hours
	// for mirror_state to reach MIRRORED. If you want Terraform to wait for the mirror to finish on
	// create/stop/resume operations, set this parameter to true. Default is false.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.28.0/docs/resources/netapp_volume_replication#wait_for_mirror NetappVolumeReplication#wait_for_mirror}
	WaitForMirror interface{} `field:"optional" json:"waitForMirror" yaml:"waitForMirror"`
}

