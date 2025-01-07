// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package rediscluster


type RedisClusterMaintenancePolicy struct {
	// weekly_maintenance_window block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.15.0/docs/resources/redis_cluster#weekly_maintenance_window RedisCluster#weekly_maintenance_window}
	WeeklyMaintenanceWindow interface{} `field:"optional" json:"weeklyMaintenanceWindow" yaml:"weeklyMaintenanceWindow"`
}

