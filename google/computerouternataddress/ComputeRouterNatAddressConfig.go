// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package computerouternataddress

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ComputeRouterNatAddressConfig struct {
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
	// Self-links of NAT IPs to be used in a Nat service.
	//
	// Only valid if the referenced RouterNat
	// natIpAllocateOption is set to MANUAL_ONLY.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.22.0/docs/resources/compute_router_nat_address#nat_ips ComputeRouterNatAddress#nat_ips}
	NatIps *[]*string `field:"required" json:"natIps" yaml:"natIps"`
	// The name of the Cloud Router in which the referenced NAT service is configured.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.22.0/docs/resources/compute_router_nat_address#router ComputeRouterNatAddress#router}
	Router *string `field:"required" json:"router" yaml:"router"`
	// The name of the Nat service in which this address will be configured.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.22.0/docs/resources/compute_router_nat_address#router_nat ComputeRouterNatAddress#router_nat}
	RouterNat *string `field:"required" json:"routerNat" yaml:"routerNat"`
	// A list of URLs of the IP resources to be drained.
	//
	// These IPs must be
	// valid static external IPs that have been assigned to the NAT.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.22.0/docs/resources/compute_router_nat_address#drain_nat_ips ComputeRouterNatAddress#drain_nat_ips}
	DrainNatIps *[]*string `field:"optional" json:"drainNatIps" yaml:"drainNatIps"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.22.0/docs/resources/compute_router_nat_address#id ComputeRouterNatAddress#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.22.0/docs/resources/compute_router_nat_address#project ComputeRouterNatAddress#project}.
	Project *string `field:"optional" json:"project" yaml:"project"`
	// Region where the NAT service reside.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.22.0/docs/resources/compute_router_nat_address#region ComputeRouterNatAddress#region}
	Region *string `field:"optional" json:"region" yaml:"region"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.22.0/docs/resources/compute_router_nat_address#timeouts ComputeRouterNatAddress#timeouts}
	Timeouts *ComputeRouterNatAddressTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

