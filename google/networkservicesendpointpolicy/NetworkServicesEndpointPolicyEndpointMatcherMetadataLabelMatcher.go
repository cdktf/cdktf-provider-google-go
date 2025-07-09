// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package networkservicesendpointpolicy


type NetworkServicesEndpointPolicyEndpointMatcherMetadataLabelMatcher struct {
	// Specifies how matching should be done. Possible values: ["MATCH_ANY", "MATCH_ALL"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.43.0/docs/resources/network_services_endpoint_policy#metadata_label_match_criteria NetworkServicesEndpointPolicy#metadata_label_match_criteria}
	MetadataLabelMatchCriteria *string `field:"required" json:"metadataLabelMatchCriteria" yaml:"metadataLabelMatchCriteria"`
	// metadata_labels block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.43.0/docs/resources/network_services_endpoint_policy#metadata_labels NetworkServicesEndpointPolicy#metadata_labels}
	MetadataLabels interface{} `field:"optional" json:"metadataLabels" yaml:"metadataLabels"`
}

