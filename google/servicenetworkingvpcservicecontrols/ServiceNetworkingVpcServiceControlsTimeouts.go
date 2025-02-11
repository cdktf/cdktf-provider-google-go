// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package servicenetworkingvpcservicecontrols


type ServiceNetworkingVpcServiceControlsTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.20.0/docs/resources/service_networking_vpc_service_controls#create ServiceNetworkingVpcServiceControls#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.20.0/docs/resources/service_networking_vpc_service_controls#delete ServiceNetworkingVpcServiceControls#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.20.0/docs/resources/service_networking_vpc_service_controls#update ServiceNetworkingVpcServiceControls#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

