// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package computeinterconnect

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ComputeInterconnectConfig struct {
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
	// Type of interconnect.
	//
	// Note that a value IT_PRIVATE has been deprecated in favor of DEDICATED.
	// Can take one of the following values:
	//   - PARTNER: A partner-managed interconnection shared between customers though a partner.
	//   - DEDICATED: A dedicated physical interconnection with the customer. Possible values: ["DEDICATED", "PARTNER", "IT_PRIVATE"]
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.8.0/docs/resources/compute_interconnect#interconnect_type ComputeInterconnect#interconnect_type}
	InterconnectType *string `field:"required" json:"interconnectType" yaml:"interconnectType"`
	// Type of link requested.
	//
	// Note that this field indicates the speed of each of the links in the
	// bundle, not the speed of the entire bundle. Can take one of the following values:
	//   - LINK_TYPE_ETHERNET_10G_LR: A 10G Ethernet with LR optics.
	//   - LINK_TYPE_ETHERNET_100G_LR: A 100G Ethernet with LR optics. Possible values: ["LINK_TYPE_ETHERNET_10G_LR", "LINK_TYPE_ETHERNET_100G_LR"]
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.8.0/docs/resources/compute_interconnect#link_type ComputeInterconnect#link_type}
	LinkType *string `field:"required" json:"linkType" yaml:"linkType"`
	// Name of the resource.
	//
	// Provided by the client when the resource is created. The name must be
	// 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters
	// long and match the regular expression '[a-z]([-a-z0-9]*[a-z0-9])?' which means the first
	// character must be a lowercase letter, and all following characters must be a dash,
	// lowercase letter, or digit, except the last character, which cannot be a dash.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.8.0/docs/resources/compute_interconnect#name ComputeInterconnect#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Target number of physical links in the link bundle, as requested by the customer.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.8.0/docs/resources/compute_interconnect#requested_link_count ComputeInterconnect#requested_link_count}
	RequestedLinkCount *float64 `field:"required" json:"requestedLinkCount" yaml:"requestedLinkCount"`
	// Administrative status of the interconnect.
	//
	// When this is set to true, the Interconnect is
	// functional and can carry traffic. When set to false, no packets can be carried over the
	// interconnect and no BGP routes are exchanged over it. By default, the status is set to true.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.8.0/docs/resources/compute_interconnect#admin_enabled ComputeInterconnect#admin_enabled}
	AdminEnabled interface{} `field:"optional" json:"adminEnabled" yaml:"adminEnabled"`
	// Customer name, to put in the Letter of Authorization as the party authorized to request a crossconnect.
	//
	// This field is required for Dedicated and Partner Interconnect, should not be specified
	// for cross-cloud interconnect.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.8.0/docs/resources/compute_interconnect#customer_name ComputeInterconnect#customer_name}
	CustomerName *string `field:"optional" json:"customerName" yaml:"customerName"`
	// An optional description of this resource. Provide this property when you create the resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.8.0/docs/resources/compute_interconnect#description ComputeInterconnect#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.8.0/docs/resources/compute_interconnect#id ComputeInterconnect#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Labels for this resource.
	//
	// These can only be added or modified by the setLabels
	// method. Each label key/value pair must comply with RFC1035. Label values may be empty.
	//
	//
	// **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
	// Please refer to the field 'effective_labels' for all of the labels present on the resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.8.0/docs/resources/compute_interconnect#labels ComputeInterconnect#labels}
	Labels *map[string]*string `field:"optional" json:"labels" yaml:"labels"`
	// URL of the InterconnectLocation object that represents where this connection is to be provisioned.
	//
	// Specifies the location inside Google's Networks, should not be passed in case of cross-cloud interconnect.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.8.0/docs/resources/compute_interconnect#location ComputeInterconnect#location}
	Location *string `field:"optional" json:"location" yaml:"location"`
	// macsec block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.8.0/docs/resources/compute_interconnect#macsec ComputeInterconnect#macsec}
	Macsec *ComputeInterconnectMacsec `field:"optional" json:"macsec" yaml:"macsec"`
	// Enable or disable MACsec on this Interconnect connection. MACsec enablement fails if the MACsec object is not specified.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.8.0/docs/resources/compute_interconnect#macsec_enabled ComputeInterconnect#macsec_enabled}
	MacsecEnabled interface{} `field:"optional" json:"macsecEnabled" yaml:"macsecEnabled"`
	// Email address to contact the customer NOC for operations and maintenance notifications regarding this Interconnect.
	//
	// If specified, this will be used for notifications in addition to
	// all other forms described, such as Cloud Monitoring logs alerting and Cloud Notifications.
	// This field is required for users who sign up for Cloud Interconnect using workforce identity
	// federation.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.8.0/docs/resources/compute_interconnect#noc_contact_email ComputeInterconnect#noc_contact_email}
	NocContactEmail *string `field:"optional" json:"nocContactEmail" yaml:"nocContactEmail"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.8.0/docs/resources/compute_interconnect#project ComputeInterconnect#project}.
	Project *string `field:"optional" json:"project" yaml:"project"`
	// Indicates that this is a Cross-Cloud Interconnect.
	//
	// This field specifies the location outside
	// of Google's network that the interconnect is connected to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.8.0/docs/resources/compute_interconnect#remote_location ComputeInterconnect#remote_location}
	RemoteLocation *string `field:"optional" json:"remoteLocation" yaml:"remoteLocation"`
	// interconnects.list of features requested for this Interconnect connection. Options: IF_MACSEC ( If specified then the connection is created on MACsec capable hardware ports. If not specified, the default value is false, which allocates non-MACsec capable ports first if available). Note that MACSEC is still technically allowed for compatibility reasons, but it does not work with the API, and will be removed in an upcoming major version. Possible values: ["MACSEC", "IF_MACSEC"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.8.0/docs/resources/compute_interconnect#requested_features ComputeInterconnect#requested_features}
	RequestedFeatures *[]*string `field:"optional" json:"requestedFeatures" yaml:"requestedFeatures"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.8.0/docs/resources/compute_interconnect#timeouts ComputeInterconnect#timeouts}
	Timeouts *ComputeInterconnectTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

