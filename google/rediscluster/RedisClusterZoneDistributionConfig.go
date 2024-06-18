// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package rediscluster


type RedisClusterZoneDistributionConfig struct {
	// Immutable.
	//
	// The mode for zone distribution for Memorystore Redis cluster.
	// If not provided, MULTI_ZONE will be used as default Possible values: ["MULTI_ZONE", "SINGLE_ZONE"]
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.34.0/docs/resources/redis_cluster#mode RedisCluster#mode}
	Mode *string `field:"optional" json:"mode" yaml:"mode"`
	// Immutable. The zone for single zone Memorystore Redis cluster.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.34.0/docs/resources/redis_cluster#zone RedisCluster#zone}
	Zone *string `field:"optional" json:"zone" yaml:"zone"`
}

