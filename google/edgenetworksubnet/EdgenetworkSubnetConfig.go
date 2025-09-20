// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package edgenetworksubnet

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type EdgenetworkSubnetConfig struct {
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
	// The Google Cloud region to which the target Distributed Cloud Edge zone belongs.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.50.0/docs/resources/edgenetwork_subnet#location EdgenetworkSubnet#location}
	Location *string `field:"required" json:"location" yaml:"location"`
	// The ID of the network to which this router belongs. Must be of the form: 'projects/{{project}}/locations/{{location}}/zones/{{zone}}/networks/{{network_id}}'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.50.0/docs/resources/edgenetwork_subnet#network EdgenetworkSubnet#network}
	Network *string `field:"required" json:"network" yaml:"network"`
	// A unique ID that identifies this subnet.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.50.0/docs/resources/edgenetwork_subnet#subnet_id EdgenetworkSubnet#subnet_id}
	SubnetId *string `field:"required" json:"subnetId" yaml:"subnetId"`
	// The name of the target Distributed Cloud Edge zone.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.50.0/docs/resources/edgenetwork_subnet#zone EdgenetworkSubnet#zone}
	Zone *string `field:"required" json:"zone" yaml:"zone"`
	// A free-text description of the resource. Max length 1024 characters.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.50.0/docs/resources/edgenetwork_subnet#description EdgenetworkSubnet#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.50.0/docs/resources/edgenetwork_subnet#id EdgenetworkSubnet#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// The ranges of ipv4 addresses that are owned by this subnetwork, in CIDR format.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.50.0/docs/resources/edgenetwork_subnet#ipv4_cidr EdgenetworkSubnet#ipv4_cidr}
	Ipv4Cidr *[]*string `field:"optional" json:"ipv4Cidr" yaml:"ipv4Cidr"`
	// The ranges of ipv6 addresses that are owned by this subnetwork, in CIDR format.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.50.0/docs/resources/edgenetwork_subnet#ipv6_cidr EdgenetworkSubnet#ipv6_cidr}
	Ipv6Cidr *[]*string `field:"optional" json:"ipv6Cidr" yaml:"ipv6Cidr"`
	// Labels associated with this resource.
	//
	// **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
	// Please refer to the field 'effective_labels' for all of the labels present on the resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.50.0/docs/resources/edgenetwork_subnet#labels EdgenetworkSubnet#labels}
	Labels *map[string]*string `field:"optional" json:"labels" yaml:"labels"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.50.0/docs/resources/edgenetwork_subnet#project EdgenetworkSubnet#project}.
	Project *string `field:"optional" json:"project" yaml:"project"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.50.0/docs/resources/edgenetwork_subnet#timeouts EdgenetworkSubnet#timeouts}
	Timeouts *EdgenetworkSubnetTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
	// VLAN ID for this subnetwork. If not specified, one is assigned automatically.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.50.0/docs/resources/edgenetwork_subnet#vlan_id EdgenetworkSubnet#vlan_id}
	VlanId *float64 `field:"optional" json:"vlanId" yaml:"vlanId"`
}

