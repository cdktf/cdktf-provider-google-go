// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package networksecurityfirewallendpoint

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type NetworkSecurityFirewallEndpointConfig struct {
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
	// Project to bill on endpoint uptime usage.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.34.0/docs/resources/network_security_firewall_endpoint#billing_project_id NetworkSecurityFirewallEndpoint#billing_project_id}
	BillingProjectId *string `field:"required" json:"billingProjectId" yaml:"billingProjectId"`
	// The location (zone) of the firewall endpoint.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.34.0/docs/resources/network_security_firewall_endpoint#location NetworkSecurityFirewallEndpoint#location}
	Location *string `field:"required" json:"location" yaml:"location"`
	// The name of the firewall endpoint resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.34.0/docs/resources/network_security_firewall_endpoint#name NetworkSecurityFirewallEndpoint#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The name of the parent this firewall endpoint belongs to. Format: organizations/{organization_id}.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.34.0/docs/resources/network_security_firewall_endpoint#parent NetworkSecurityFirewallEndpoint#parent}
	Parent *string `field:"required" json:"parent" yaml:"parent"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.34.0/docs/resources/network_security_firewall_endpoint#id NetworkSecurityFirewallEndpoint#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// A map of key/value label pairs to assign to the resource.
	//
	// **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
	// Please refer to the field 'effective_labels' for all of the labels present on the resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.34.0/docs/resources/network_security_firewall_endpoint#labels NetworkSecurityFirewallEndpoint#labels}
	Labels *map[string]*string `field:"optional" json:"labels" yaml:"labels"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.34.0/docs/resources/network_security_firewall_endpoint#timeouts NetworkSecurityFirewallEndpoint#timeouts}
	Timeouts *NetworkSecurityFirewallEndpointTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

