// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package networkservicesendpointpolicy


type NetworkServicesEndpointPolicyEndpointMatcherMetadataLabelMatcherMetadataLabels struct {
	// Required. Label name presented as key in xDS Node Metadata.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.36.1/docs/resources/network_services_endpoint_policy#label_name NetworkServicesEndpointPolicy#label_name}
	LabelName *string `field:"required" json:"labelName" yaml:"labelName"`
	// Required. Label value presented as value corresponding to the above key, in xDS Node Metadata.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.36.1/docs/resources/network_services_endpoint_policy#label_value NetworkServicesEndpointPolicy#label_value}
	LabelValue *string `field:"required" json:"labelValue" yaml:"labelValue"`
}

