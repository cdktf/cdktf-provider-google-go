// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package geminicoderepositoryindex

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type GeminiCodeRepositoryIndexConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.47.0/docs/resources/gemini_code_repository_index#code_repository_index_id GeminiCodeRepositoryIndex#code_repository_index_id}
	CodeRepositoryIndexId *string `field:"required" json:"codeRepositoryIndexId" yaml:"codeRepositoryIndexId"`
	// The location of the Code Repository Index, for example 'us-central1'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.47.0/docs/resources/gemini_code_repository_index#location GeminiCodeRepositoryIndex#location}
	Location *string `field:"required" json:"location" yaml:"location"`
	// If set to true, will allow deletion of the CodeRepositoryIndex even if there are existing RepositoryGroups for the resource.
	//
	// These RepositoryGroups will also be deleted.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.47.0/docs/resources/gemini_code_repository_index#force_destroy GeminiCodeRepositoryIndex#force_destroy}
	ForceDestroy interface{} `field:"optional" json:"forceDestroy" yaml:"forceDestroy"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.47.0/docs/resources/gemini_code_repository_index#id GeminiCodeRepositoryIndex#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Optional. Immutable. Customer-managed encryption key name, in the format 'projects/* /locations/* /keyRings/* /cryptoKeys/*'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.47.0/docs/resources/gemini_code_repository_index#kms_key GeminiCodeRepositoryIndex#kms_key}
	//
	// Note: The above comment contained a comment block ending sequence (* followed by /). We have introduced a space between to prevent syntax errors. Please ignore the space.
	KmsKey *string `field:"optional" json:"kmsKey" yaml:"kmsKey"`
	// Optional. Labels as key value pairs.
	//
	// **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
	// Please refer to the field 'effective_labels' for all of the labels present on the resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.47.0/docs/resources/gemini_code_repository_index#labels GeminiCodeRepositoryIndex#labels}
	Labels *map[string]*string `field:"optional" json:"labels" yaml:"labels"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.47.0/docs/resources/gemini_code_repository_index#project GeminiCodeRepositoryIndex#project}.
	Project *string `field:"optional" json:"project" yaml:"project"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.47.0/docs/resources/gemini_code_repository_index#timeouts GeminiCodeRepositoryIndex#timeouts}
	Timeouts *GeminiCodeRepositoryIndexTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

