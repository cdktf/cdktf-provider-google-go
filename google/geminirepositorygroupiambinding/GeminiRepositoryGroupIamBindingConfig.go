// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package geminirepositorygroupiambinding

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type GeminiRepositoryGroupIamBindingConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.34.0/docs/resources/gemini_repository_group_iam_binding#code_repository_index GeminiRepositoryGroupIamBinding#code_repository_index}.
	CodeRepositoryIndex *string `field:"required" json:"codeRepositoryIndex" yaml:"codeRepositoryIndex"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.34.0/docs/resources/gemini_repository_group_iam_binding#members GeminiRepositoryGroupIamBinding#members}.
	Members *[]*string `field:"required" json:"members" yaml:"members"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.34.0/docs/resources/gemini_repository_group_iam_binding#repository_group_id GeminiRepositoryGroupIamBinding#repository_group_id}.
	RepositoryGroupId *string `field:"required" json:"repositoryGroupId" yaml:"repositoryGroupId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.34.0/docs/resources/gemini_repository_group_iam_binding#role GeminiRepositoryGroupIamBinding#role}.
	Role *string `field:"required" json:"role" yaml:"role"`
	// condition block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.34.0/docs/resources/gemini_repository_group_iam_binding#condition GeminiRepositoryGroupIamBinding#condition}
	Condition *GeminiRepositoryGroupIamBindingCondition `field:"optional" json:"condition" yaml:"condition"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.34.0/docs/resources/gemini_repository_group_iam_binding#id GeminiRepositoryGroupIamBinding#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.34.0/docs/resources/gemini_repository_group_iam_binding#location GeminiRepositoryGroupIamBinding#location}.
	Location *string `field:"optional" json:"location" yaml:"location"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.34.0/docs/resources/gemini_repository_group_iam_binding#project GeminiRepositoryGroupIamBinding#project}.
	Project *string `field:"optional" json:"project" yaml:"project"`
}

