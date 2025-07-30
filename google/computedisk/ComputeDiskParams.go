// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package computedisk


type ComputeDiskParams struct {
	// Resource manager tags to be bound to the disk.
	//
	// Tag keys and values have the
	// same definition as resource manager tags. Keys must be in the format tagKeys/{tag_key_id},
	// and values are in the format tagValues/456.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.46.0/docs/resources/compute_disk#resource_manager_tags ComputeDisk#resource_manager_tags}
	ResourceManagerTags *map[string]*string `field:"optional" json:"resourceManagerTags" yaml:"resourceManagerTags"`
}

