// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package rediscluster


type RedisClusterPersistenceConfigRdbConfig struct {
	// Optional. Available snapshot periods for scheduling.
	//
	// - ONE_HOUR:	Snapshot every 1 hour.
	// - SIX_HOURS:	Snapshot every 6 hours.
	// - TWELVE_HOURS:	Snapshot every 12 hours.
	// - TWENTY_FOUR_HOURS:	Snapshot every 24 hours. Possible values: ["SNAPSHOT_PERIOD_UNSPECIFIED", "ONE_HOUR", "SIX_HOURS", "TWELVE_HOURS", "TWENTY_FOUR_HOURS"]
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.36.0/docs/resources/redis_cluster#rdb_snapshot_period RedisCluster#rdb_snapshot_period}
	RdbSnapshotPeriod *string `field:"optional" json:"rdbSnapshotPeriod" yaml:"rdbSnapshotPeriod"`
	// The time that the first snapshot was/will be attempted, and to which future snapshots will be aligned.
	//
	// If not provided, the current time will be used.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.36.0/docs/resources/redis_cluster#rdb_snapshot_start_time RedisCluster#rdb_snapshot_start_time}
	RdbSnapshotStartTime *string `field:"optional" json:"rdbSnapshotStartTime" yaml:"rdbSnapshotStartTime"`
}

