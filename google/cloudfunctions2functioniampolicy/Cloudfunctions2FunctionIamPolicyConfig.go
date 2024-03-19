// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package cloudfunctions2functioniampolicy

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type Cloudfunctions2FunctionIamPolicyConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.21.0/docs/resources/cloudfunctions2_function_iam_policy#cloud_function Cloudfunctions2FunctionIamPolicy#cloud_function}.
	CloudFunction *string `field:"required" json:"cloudFunction" yaml:"cloudFunction"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.21.0/docs/resources/cloudfunctions2_function_iam_policy#policy_data Cloudfunctions2FunctionIamPolicy#policy_data}.
	PolicyData *string `field:"required" json:"policyData" yaml:"policyData"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.21.0/docs/resources/cloudfunctions2_function_iam_policy#id Cloudfunctions2FunctionIamPolicy#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.21.0/docs/resources/cloudfunctions2_function_iam_policy#location Cloudfunctions2FunctionIamPolicy#location}.
	Location *string `field:"optional" json:"location" yaml:"location"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.21.0/docs/resources/cloudfunctions2_function_iam_policy#project Cloudfunctions2FunctionIamPolicy#project}.
	Project *string `field:"optional" json:"project" yaml:"project"`
}

