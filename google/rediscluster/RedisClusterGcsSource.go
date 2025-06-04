// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package rediscluster


type RedisClusterGcsSource struct {
	// URIs of the GCS objects to import. Example: gs://bucket1/object1, gs://bucket2/folder2/object2.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.38.0/docs/resources/redis_cluster#uris RedisCluster#uris}
	Uris *[]*string `field:"required" json:"uris" yaml:"uris"`
}

