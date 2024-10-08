// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package edgenetworknetwork

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type EdgenetworkNetworkConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.6.0/docs/resources/edgenetwork_network#location EdgenetworkNetwork#location}
	Location *string `field:"required" json:"location" yaml:"location"`
	// A unique ID that identifies this network.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.6.0/docs/resources/edgenetwork_network#network_id EdgenetworkNetwork#network_id}
	NetworkId *string `field:"required" json:"networkId" yaml:"networkId"`
	// The name of the target Distributed Cloud Edge zone.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.6.0/docs/resources/edgenetwork_network#zone EdgenetworkNetwork#zone}
	Zone *string `field:"required" json:"zone" yaml:"zone"`
	// A free-text description of the resource. Max length 1024 characters.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.6.0/docs/resources/edgenetwork_network#description EdgenetworkNetwork#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.6.0/docs/resources/edgenetwork_network#id EdgenetworkNetwork#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Labels associated with this resource.
	//
	// **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
	// Please refer to the field 'effective_labels' for all of the labels present on the resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.6.0/docs/resources/edgenetwork_network#labels EdgenetworkNetwork#labels}
	Labels *map[string]*string `field:"optional" json:"labels" yaml:"labels"`
	// IP (L3) MTU value of the network. Default value is '1500'. Possible values are: '1500', '9000'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.6.0/docs/resources/edgenetwork_network#mtu EdgenetworkNetwork#mtu}
	Mtu *float64 `field:"optional" json:"mtu" yaml:"mtu"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.6.0/docs/resources/edgenetwork_network#project EdgenetworkNetwork#project}.
	Project *string `field:"optional" json:"project" yaml:"project"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.6.0/docs/resources/edgenetwork_network#timeouts EdgenetworkNetwork#timeouts}
	Timeouts *EdgenetworkNetworkTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

