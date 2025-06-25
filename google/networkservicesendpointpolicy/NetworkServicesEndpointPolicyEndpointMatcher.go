// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package networkservicesendpointpolicy


type NetworkServicesEndpointPolicyEndpointMatcher struct {
	// metadata_label_matcher block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.41.0/docs/resources/network_services_endpoint_policy#metadata_label_matcher NetworkServicesEndpointPolicy#metadata_label_matcher}
	MetadataLabelMatcher *NetworkServicesEndpointPolicyEndpointMatcherMetadataLabelMatcher `field:"required" json:"metadataLabelMatcher" yaml:"metadataLabelMatcher"`
}

