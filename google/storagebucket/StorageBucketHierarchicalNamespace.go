// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package storagebucket


type StorageBucketHierarchicalNamespace struct {
	// Set this enabled flag to true when folders with logical files structure. Default value is false.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.9.0/docs/resources/storage_bucket#enabled StorageBucket#enabled}
	Enabled interface{} `field:"required" json:"enabled" yaml:"enabled"`
}

