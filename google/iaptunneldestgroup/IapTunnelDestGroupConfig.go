// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package iaptunneldestgroup

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type IapTunnelDestGroupConfig struct {
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
	// Unique tunnel destination group name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.0/docs/resources/iap_tunnel_dest_group#group_name IapTunnelDestGroup#group_name}
	GroupName *string `field:"required" json:"groupName" yaml:"groupName"`
	// List of CIDRs that this group applies to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.0/docs/resources/iap_tunnel_dest_group#cidrs IapTunnelDestGroup#cidrs}
	Cidrs *[]*string `field:"optional" json:"cidrs" yaml:"cidrs"`
	// List of FQDNs that this group applies to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.0/docs/resources/iap_tunnel_dest_group#fqdns IapTunnelDestGroup#fqdns}
	Fqdns *[]*string `field:"optional" json:"fqdns" yaml:"fqdns"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.0/docs/resources/iap_tunnel_dest_group#id IapTunnelDestGroup#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.0/docs/resources/iap_tunnel_dest_group#project IapTunnelDestGroup#project}.
	Project *string `field:"optional" json:"project" yaml:"project"`
	// The region of the tunnel group. Must be the same as the network resources in the group.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.0/docs/resources/iap_tunnel_dest_group#region IapTunnelDestGroup#region}
	Region *string `field:"optional" json:"region" yaml:"region"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.0/docs/resources/iap_tunnel_dest_group#timeouts IapTunnelDestGroup#timeouts}
	Timeouts *IapTunnelDestGroupTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

