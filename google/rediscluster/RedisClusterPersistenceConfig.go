// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package rediscluster


type RedisClusterPersistenceConfig struct {
	// aof_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.38.0/docs/resources/redis_cluster#aof_config RedisCluster#aof_config}
	AofConfig *RedisClusterPersistenceConfigAofConfig `field:"optional" json:"aofConfig" yaml:"aofConfig"`
	// Optional. Controls whether Persistence features are enabled. If not provided, the existing value will be used.
	//
	// - DISABLED: 	Persistence (both backup and restore) is disabled for the cluster.
	// - RDB: RDB based Persistence is enabled.
	// - AOF: AOF based Persistence is enabled. Possible values: ["PERSISTENCE_MODE_UNSPECIFIED", "DISABLED", "RDB", "AOF"]
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.38.0/docs/resources/redis_cluster#mode RedisCluster#mode}
	Mode *string `field:"optional" json:"mode" yaml:"mode"`
	// rdb_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.38.0/docs/resources/redis_cluster#rdb_config RedisCluster#rdb_config}
	RdbConfig *RedisClusterPersistenceConfigRdbConfig `field:"optional" json:"rdbConfig" yaml:"rdbConfig"`
}

