// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package beyondcorpapplication

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type BeyondcorpApplicationConfig struct {
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
	// Optional.
	//
	// User-settable Application resource ID.
	// * Must start with a letter.
	// * Must contain between 4-63 characters from '/a-z-/'.
	// * Must end with a number or letter.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.20.0/docs/resources/beyondcorp_application#application_id BeyondcorpApplication#application_id}
	ApplicationId *string `field:"required" json:"applicationId" yaml:"applicationId"`
	// endpoint_matchers block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.20.0/docs/resources/beyondcorp_application#endpoint_matchers BeyondcorpApplication#endpoint_matchers}
	EndpointMatchers interface{} `field:"required" json:"endpointMatchers" yaml:"endpointMatchers"`
	// Part of 'parent'. See documentation of 'projectsId'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.20.0/docs/resources/beyondcorp_application#security_gateways_id BeyondcorpApplication#security_gateways_id}
	SecurityGatewaysId *string `field:"required" json:"securityGatewaysId" yaml:"securityGatewaysId"`
	// Optional. An arbitrary user-provided name for the Application resource. Cannot exceed 64 characters.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.20.0/docs/resources/beyondcorp_application#display_name BeyondcorpApplication#display_name}
	DisplayName *string `field:"optional" json:"displayName" yaml:"displayName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.20.0/docs/resources/beyondcorp_application#id BeyondcorpApplication#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.20.0/docs/resources/beyondcorp_application#project BeyondcorpApplication#project}.
	Project *string `field:"optional" json:"project" yaml:"project"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.20.0/docs/resources/beyondcorp_application#timeouts BeyondcorpApplication#timeouts}
	Timeouts *BeyondcorpApplicationTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

