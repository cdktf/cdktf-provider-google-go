// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package rediscluster


type RedisClusterManagedBackupSource struct {
	// Example: //redis.googleapis.com/projects/{project}/locations/{location}/backupCollections/{collection}/backups/{backup} A shorter version (without the prefix) of the backup name is also supported, like projects/{project}/locations/{location}/backupCollections/{collection}/backups/{backupId}. In this case, it assumes the backup is under redis.googleapis.com.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.34.1/docs/resources/redis_cluster#backup RedisCluster#backup}
	Backup *string `field:"required" json:"backup" yaml:"backup"`
}

