// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package storagebucketobject


type StorageBucketObjectRetention struct {
	// The object retention mode. Supported values include: "Unlocked", "Locked".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.9.0/docs/resources/storage_bucket_object#mode StorageBucketObject#mode}
	Mode *string `field:"required" json:"mode" yaml:"mode"`
	// Time in RFC 3339 (e.g. 2030-01-01T02:03:04Z) until which object retention protects this object.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.9.0/docs/resources/storage_bucket_object#retain_until_time StorageBucketObject#retain_until_time}
	RetainUntilTime *string `field:"required" json:"retainUntilTime" yaml:"retainUntilTime"`
}

