// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package managedkafkacluster


type ManagedKafkaClusterTlsConfigTrustConfig struct {
	// cas_configs block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.50.0/docs/resources/managed_kafka_cluster#cas_configs ManagedKafkaCluster#cas_configs}
	CasConfigs interface{} `field:"optional" json:"casConfigs" yaml:"casConfigs"`
}

