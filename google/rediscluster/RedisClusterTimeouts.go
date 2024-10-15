// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package rediscluster


type RedisClusterTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.7.0/docs/resources/redis_cluster#create RedisCluster#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.7.0/docs/resources/redis_cluster#delete RedisCluster#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.7.0/docs/resources/redis_cluster#update RedisCluster#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

