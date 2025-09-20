// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package beyondcorpapplicationiambinding

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type BeyondcorpApplicationIamBindingConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.50.0/docs/resources/beyondcorp_application_iam_binding#application_id BeyondcorpApplicationIamBinding#application_id}.
	ApplicationId *string `field:"required" json:"applicationId" yaml:"applicationId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.50.0/docs/resources/beyondcorp_application_iam_binding#members BeyondcorpApplicationIamBinding#members}.
	Members *[]*string `field:"required" json:"members" yaml:"members"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.50.0/docs/resources/beyondcorp_application_iam_binding#role BeyondcorpApplicationIamBinding#role}.
	Role *string `field:"required" json:"role" yaml:"role"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.50.0/docs/resources/beyondcorp_application_iam_binding#security_gateways_id BeyondcorpApplicationIamBinding#security_gateways_id}.
	SecurityGatewaysId *string `field:"required" json:"securityGatewaysId" yaml:"securityGatewaysId"`
	// condition block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.50.0/docs/resources/beyondcorp_application_iam_binding#condition BeyondcorpApplicationIamBinding#condition}
	Condition *BeyondcorpApplicationIamBindingCondition `field:"optional" json:"condition" yaml:"condition"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.50.0/docs/resources/beyondcorp_application_iam_binding#id BeyondcorpApplicationIamBinding#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.50.0/docs/resources/beyondcorp_application_iam_binding#project BeyondcorpApplicationIamBinding#project}.
	Project *string `field:"optional" json:"project" yaml:"project"`
}

