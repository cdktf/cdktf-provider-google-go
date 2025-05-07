// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package apigeednszone

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ApigeeDnsZoneConfig struct {
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
	// Description for the zone.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.34.0/docs/resources/apigee_dns_zone#description ApigeeDnsZone#description}
	Description *string `field:"required" json:"description" yaml:"description"`
	// ID of the dns zone.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.34.0/docs/resources/apigee_dns_zone#dns_zone_id ApigeeDnsZone#dns_zone_id}
	DnsZoneId *string `field:"required" json:"dnsZoneId" yaml:"dnsZoneId"`
	// Doamin for the zone.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.34.0/docs/resources/apigee_dns_zone#domain ApigeeDnsZone#domain}
	Domain *string `field:"required" json:"domain" yaml:"domain"`
	// The Apigee Organization associated with the Apigee instance, in the format 'organizations/{{org_name}}'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.34.0/docs/resources/apigee_dns_zone#org_id ApigeeDnsZone#org_id}
	OrgId *string `field:"required" json:"orgId" yaml:"orgId"`
	// peering_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.34.0/docs/resources/apigee_dns_zone#peering_config ApigeeDnsZone#peering_config}
	PeeringConfig *ApigeeDnsZonePeeringConfig `field:"required" json:"peeringConfig" yaml:"peeringConfig"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.34.0/docs/resources/apigee_dns_zone#id ApigeeDnsZone#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.34.0/docs/resources/apigee_dns_zone#timeouts ApigeeDnsZone#timeouts}
	Timeouts *ApigeeDnsZoneTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

