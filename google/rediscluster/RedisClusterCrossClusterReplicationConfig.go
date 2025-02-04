// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package rediscluster


type RedisClusterCrossClusterReplicationConfig struct {
	// The role of the cluster in cross cluster replication. Supported values are:.
	//
	// 1. 'CLUSTER_ROLE_UNSPECIFIED': This is an independent cluster that has never participated in cross cluster replication. It allows both reads and writes.
	//
	// 1. 'NONE': This is an independent cluster that previously participated in cross cluster replication(either as a 'PRIMARY' or 'SECONDARY' cluster). It allows both reads and writes.
	//
	// 1. 'PRIMARY': This cluster serves as the replication source for secondary clusters that are replicating from it. Any data written to it is automatically replicated to its secondary clusters. It allows both reads and writes.
	//
	// 1. 'SECONDARY': This cluster replicates data from the primary cluster. It allows only reads. Possible values: ["CLUSTER_ROLE_UNSPECIFIED", "NONE", "PRIMARY", "SECONDARY"]
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.19.0/docs/resources/redis_cluster#cluster_role RedisCluster#cluster_role}
	ClusterRole *string `field:"optional" json:"clusterRole" yaml:"clusterRole"`
	// primary_cluster block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.19.0/docs/resources/redis_cluster#primary_cluster RedisCluster#primary_cluster}
	PrimaryCluster *RedisClusterCrossClusterReplicationConfigPrimaryCluster `field:"optional" json:"primaryCluster" yaml:"primaryCluster"`
	// secondary_clusters block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.19.0/docs/resources/redis_cluster#secondary_clusters RedisCluster#secondary_clusters}
	SecondaryClusters interface{} `field:"optional" json:"secondaryClusters" yaml:"secondaryClusters"`
}

