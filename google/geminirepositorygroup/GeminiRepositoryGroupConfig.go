// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package geminirepositorygroup

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type GeminiRepositoryGroupConfig struct {
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
	// Required. Id of the Code Repository Index.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.46.0/docs/resources/gemini_repository_group#code_repository_index GeminiRepositoryGroup#code_repository_index}
	CodeRepositoryIndex *string `field:"required" json:"codeRepositoryIndex" yaml:"codeRepositoryIndex"`
	// The location of the Code Repository Index, for example 'us-central1'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.46.0/docs/resources/gemini_repository_group#location GeminiRepositoryGroup#location}
	Location *string `field:"required" json:"location" yaml:"location"`
	// repositories block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.46.0/docs/resources/gemini_repository_group#repositories GeminiRepositoryGroup#repositories}
	Repositories interface{} `field:"required" json:"repositories" yaml:"repositories"`
	// Required. Id of the Repository Group.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.46.0/docs/resources/gemini_repository_group#repository_group_id GeminiRepositoryGroup#repository_group_id}
	RepositoryGroupId *string `field:"required" json:"repositoryGroupId" yaml:"repositoryGroupId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.46.0/docs/resources/gemini_repository_group#id GeminiRepositoryGroup#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Optional. Labels as key value pairs.
	//
	// **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
	// Please refer to the field 'effective_labels' for all of the labels present on the resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.46.0/docs/resources/gemini_repository_group#labels GeminiRepositoryGroup#labels}
	Labels *map[string]*string `field:"optional" json:"labels" yaml:"labels"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.46.0/docs/resources/gemini_repository_group#project GeminiRepositoryGroup#project}.
	Project *string `field:"optional" json:"project" yaml:"project"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.46.0/docs/resources/gemini_repository_group#timeouts GeminiRepositoryGroup#timeouts}
	Timeouts *GeminiRepositoryGroupTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

