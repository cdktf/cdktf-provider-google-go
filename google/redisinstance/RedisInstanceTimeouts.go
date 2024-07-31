// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package redisinstance


type RedisInstanceTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.39.1/docs/resources/redis_instance#create RedisInstance#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.39.1/docs/resources/redis_instance#delete RedisInstance#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.39.1/docs/resources/redis_instance#update RedisInstance#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

