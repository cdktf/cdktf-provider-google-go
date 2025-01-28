// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package managedkafkacluster


type ManagedKafkaClusterRebalanceConfig struct {
	// The rebalance behavior for the cluster. When not specified, defaults to 'NO_REBALANCE'. Possible values: 'MODE_UNSPECIFIED', 'NO_REBALANCE', 'AUTO_REBALANCE_ON_SCALE_UP'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.18.0/docs/resources/managed_kafka_cluster#mode ManagedKafkaCluster#mode}
	Mode *string `field:"optional" json:"mode" yaml:"mode"`
}

