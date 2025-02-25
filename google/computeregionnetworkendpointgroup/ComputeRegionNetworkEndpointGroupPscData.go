// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package computeregionnetworkendpointgroup


type ComputeRegionNetworkEndpointGroupPscData struct {
	// The PSC producer port to use when consumer PSC NEG connects to a producer.
	//
	// If
	// this flag isn't specified for a PSC NEG with endpoint type
	// private-service-connect, then PSC NEG will be connected to a first port in the
	// available PSC producer port range.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.22.0/docs/resources/compute_region_network_endpoint_group#producer_port ComputeRegionNetworkEndpointGroup#producer_port}
	ProducerPort *string `field:"optional" json:"producerPort" yaml:"producerPort"`
}

