// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package memcacheinstance


type MemcacheInstanceTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.19.0/docs/resources/memcache_instance#create MemcacheInstance#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.19.0/docs/resources/memcache_instance#delete MemcacheInstance#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.19.0/docs/resources/memcache_instance#update MemcacheInstance#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

