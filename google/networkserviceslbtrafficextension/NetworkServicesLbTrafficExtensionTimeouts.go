// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package networkserviceslbtrafficextension


type NetworkServicesLbTrafficExtensionTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.31.1/docs/resources/network_services_lb_traffic_extension#create NetworkServicesLbTrafficExtension#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.31.1/docs/resources/network_services_lb_traffic_extension#delete NetworkServicesLbTrafficExtension#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.31.1/docs/resources/network_services_lb_traffic_extension#update NetworkServicesLbTrafficExtension#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

