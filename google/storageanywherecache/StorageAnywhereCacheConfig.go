// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package storageanywherecache

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type StorageAnywhereCacheConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktf.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// A reference to Bucket resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.0/docs/resources/storage_anywhere_cache#bucket StorageAnywhereCache#bucket}
	Bucket *string `field:"required" json:"bucket" yaml:"bucket"`
	// The zone in which the cache instance needs to be created. For example, 'us-central1-a.'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.0/docs/resources/storage_anywhere_cache#zone StorageAnywhereCache#zone}
	Zone *string `field:"required" json:"zone" yaml:"zone"`
	// The cache admission policy dictates whether a block should be inserted upon a cache miss.
	//
	// Default value: "admit-on-first-miss" Possible values: ["admit-on-first-miss", "admit-on-second-miss"]
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.0/docs/resources/storage_anywhere_cache#admission_policy StorageAnywhereCache#admission_policy}
	AdmissionPolicy *string `field:"optional" json:"admissionPolicy" yaml:"admissionPolicy"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.0/docs/resources/storage_anywhere_cache#id StorageAnywhereCache#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.0/docs/resources/storage_anywhere_cache#timeouts StorageAnywhereCache#timeouts}
	Timeouts *StorageAnywhereCacheTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
	// The TTL of all cache entries in whole seconds. e.g., "7200s". It defaults to '86400s'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.0/docs/resources/storage_anywhere_cache#ttl StorageAnywhereCache#ttl}
	Ttl *string `field:"optional" json:"ttl" yaml:"ttl"`
}

