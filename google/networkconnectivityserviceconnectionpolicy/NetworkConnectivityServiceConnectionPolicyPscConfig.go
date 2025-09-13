// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package networkconnectivityserviceconnectionpolicy


type NetworkConnectivityServiceConnectionPolicyPscConfig struct {
	// IDs of the subnetworks or fully qualified identifiers for the subnetworks.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.3/docs/resources/network_connectivity_service_connection_policy#subnetworks NetworkConnectivityServiceConnectionPolicy#subnetworks}
	Subnetworks *[]*string `field:"required" json:"subnetworks" yaml:"subnetworks"`
	// List of Projects, Folders, or Organizations from where the Producer instance can be within.
	//
	// For example,
	// a network administrator can provide both 'organizations/foo' and 'projects/bar' as
	// allowed_google_producers_resource_hierarchy_levels. This allowlists this network to connect with any Producer
	// instance within the 'foo' organization or the 'bar' project. By default,
	// allowedGoogleProducersResourceHierarchyLevel is empty. The format for each
	// allowedGoogleProducersResourceHierarchyLevel is / where is one of 'projects', 'folders', or 'organizations'
	// and is either the ID or the number of the resource type. Format for each
	// allowedGoogleProducersResourceHierarchyLevel value: 'projects/' or 'folders/' or 'organizations/' Eg.
	// [projects/my-project-id, projects/567, folders/891, organizations/123]
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.3/docs/resources/network_connectivity_service_connection_policy#allowed_google_producers_resource_hierarchy_level NetworkConnectivityServiceConnectionPolicy#allowed_google_producers_resource_hierarchy_level}
	AllowedGoogleProducersResourceHierarchyLevel *[]*string `field:"optional" json:"allowedGoogleProducersResourceHierarchyLevel" yaml:"allowedGoogleProducersResourceHierarchyLevel"`
	// Max number of PSC connections for this policy.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.3/docs/resources/network_connectivity_service_connection_policy#limit NetworkConnectivityServiceConnectionPolicy#limit}
	Limit *string `field:"optional" json:"limit" yaml:"limit"`
	// ProducerInstanceLocation is used to specify which authorization mechanism to use to determine which projects the Producer instance can be within.
	//
	// Possible values: ["PRODUCER_INSTANCE_LOCATION_UNSPECIFIED", "CUSTOM_RESOURCE_HIERARCHY_LEVELS"]
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.3/docs/resources/network_connectivity_service_connection_policy#producer_instance_location NetworkConnectivityServiceConnectionPolicy#producer_instance_location}
	ProducerInstanceLocation *string `field:"optional" json:"producerInstanceLocation" yaml:"producerInstanceLocation"`
}

