// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package servicenetworkingvpcservicecontrols

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ServiceNetworkingVpcServiceControlsConfig struct {
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
	// Desired VPC Service Controls state service producer VPC network, as described at the top of this page.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.48.0/docs/resources/service_networking_vpc_service_controls#enabled ServiceNetworkingVpcServiceControls#enabled}
	Enabled interface{} `field:"required" json:"enabled" yaml:"enabled"`
	// The network that the consumer is using to connect with services.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.48.0/docs/resources/service_networking_vpc_service_controls#network ServiceNetworkingVpcServiceControls#network}
	Network *string `field:"required" json:"network" yaml:"network"`
	// The service that is managing peering connectivity for a service producer's organization.
	//
	// For Google services that support this
	// functionality, this value is 'servicenetworking.googleapis.com'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.48.0/docs/resources/service_networking_vpc_service_controls#service ServiceNetworkingVpcServiceControls#service}
	Service *string `field:"required" json:"service" yaml:"service"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.48.0/docs/resources/service_networking_vpc_service_controls#id ServiceNetworkingVpcServiceControls#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// The id of the Google Cloud project containing the consumer network.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.48.0/docs/resources/service_networking_vpc_service_controls#project ServiceNetworkingVpcServiceControls#project}
	Project *string `field:"optional" json:"project" yaml:"project"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.48.0/docs/resources/service_networking_vpc_service_controls#timeouts ServiceNetworkingVpcServiceControls#timeouts}
	Timeouts *ServiceNetworkingVpcServiceControlsTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

