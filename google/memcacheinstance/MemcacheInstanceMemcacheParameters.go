// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package memcacheinstance


type MemcacheInstanceMemcacheParameters struct {
	// User-defined set of parameters to use in the memcache process.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.43.0/docs/resources/memcache_instance#params MemcacheInstance#params}
	Params *map[string]*string `field:"optional" json:"params" yaml:"params"`
}

