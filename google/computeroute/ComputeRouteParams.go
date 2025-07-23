// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package computeroute


type ComputeRouteParams struct {
	// Resource manager tags to be bound to the route.
	//
	// Tag keys and values have the
	// same definition as resource manager tags. Keys must be in the format tagKeys/{tag_key_id},
	// and values are in the format tagValues/456. The field is ignored when empty.
	// The field is immutable and causes resource replacement when mutated. This field is only
	// set at create time and modifying this field after creation will trigger recreation.
	// To apply tags to an existing resource, see the google_tags_tag_binding resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/compute_route#resource_manager_tags ComputeRoute#resource_manager_tags}
	ResourceManagerTags *map[string]*string `field:"optional" json:"resourceManagerTags" yaml:"resourceManagerTags"`
}

