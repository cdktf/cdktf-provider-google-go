// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package rediscluster


type RedisClusterAutomatedBackupConfigFixedFrequencySchedule struct {
	// start_time block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.42.0/docs/resources/redis_cluster#start_time RedisCluster#start_time}
	StartTime *RedisClusterAutomatedBackupConfigFixedFrequencyScheduleStartTime `field:"required" json:"startTime" yaml:"startTime"`
}

