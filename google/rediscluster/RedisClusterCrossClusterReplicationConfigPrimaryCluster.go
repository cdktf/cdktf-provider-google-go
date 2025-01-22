// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package rediscluster


type RedisClusterCrossClusterReplicationConfigPrimaryCluster struct {
	// The full resource path of the primary cluster in the format: projects/{project}/locations/{region}/clusters/{cluster-id}.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.17.0/docs/resources/redis_cluster#cluster RedisCluster#cluster}
	Cluster *string `field:"optional" json:"cluster" yaml:"cluster"`
}

