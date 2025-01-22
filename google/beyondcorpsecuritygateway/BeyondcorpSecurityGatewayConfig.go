// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package beyondcorpsecuritygateway

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type BeyondcorpSecurityGatewayConfig struct {
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
	// Resource ID segment making up resource 'name'. It identifies the resource within its parent collection as described in https://google.aip.dev/122.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.17.0/docs/resources/beyondcorp_security_gateway#location BeyondcorpSecurityGateway#location}
	Location *string `field:"required" json:"location" yaml:"location"`
	// Optional.
	//
	// User-settable SecurityGateway resource ID.
	// * Must start with a letter.
	// * Must contain between 4-63 characters from '/a-z-/'.
	// * Must end with a number or letter.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.17.0/docs/resources/beyondcorp_security_gateway#security_gateway_id BeyondcorpSecurityGateway#security_gateway_id}
	SecurityGatewayId *string `field:"required" json:"securityGatewayId" yaml:"securityGatewayId"`
	// Optional. An arbitrary user-provided name for the SecurityGateway. Cannot exceed 64 characters.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.17.0/docs/resources/beyondcorp_security_gateway#display_name BeyondcorpSecurityGateway#display_name}
	DisplayName *string `field:"optional" json:"displayName" yaml:"displayName"`
	// hubs block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.17.0/docs/resources/beyondcorp_security_gateway#hubs BeyondcorpSecurityGateway#hubs}
	Hubs interface{} `field:"optional" json:"hubs" yaml:"hubs"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.17.0/docs/resources/beyondcorp_security_gateway#id BeyondcorpSecurityGateway#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.17.0/docs/resources/beyondcorp_security_gateway#project BeyondcorpSecurityGateway#project}.
	Project *string `field:"optional" json:"project" yaml:"project"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.17.0/docs/resources/beyondcorp_security_gateway#timeouts BeyondcorpSecurityGateway#timeouts}
	Timeouts *BeyondcorpSecurityGatewayTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

