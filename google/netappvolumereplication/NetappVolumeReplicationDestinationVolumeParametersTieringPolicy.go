// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package netappvolumereplication


type NetappVolumeReplicationDestinationVolumeParametersTieringPolicy struct {
	// Optional.
	//
	// Time in days to mark the volume's data block as cold and make it eligible for tiering, can be range from 2-183.
	// Default is 31.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.2/docs/resources/netapp_volume_replication#cooling_threshold_days NetappVolumeReplication#cooling_threshold_days}
	CoolingThresholdDays *float64 `field:"optional" json:"coolingThresholdDays" yaml:"coolingThresholdDays"`
	// Optional.
	//
	// Flag indicating if the volume has tiering policy enable/pause. Default is PAUSED. Default value: "PAUSED" Possible values: ["ENABLED", "PAUSED"]
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.2/docs/resources/netapp_volume_replication#tier_action NetappVolumeReplication#tier_action}
	TierAction *string `field:"optional" json:"tierAction" yaml:"tierAction"`
}

