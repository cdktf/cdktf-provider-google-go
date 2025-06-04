// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package storagecontrolprojectintelligenceconfig

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type StorageControlProjectIntelligenceConfigConfig struct {
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
	// Identifier of the GCP project. For GCP project, this field can be project name or project number.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.38.0/docs/resources/storage_control_project_intelligence_config#name StorageControlProjectIntelligenceConfig#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Edition configuration of the Storage Intelligence resource. Valid values are INHERIT, TRIAL, DISABLED and STANDARD.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.38.0/docs/resources/storage_control_project_intelligence_config#edition_config StorageControlProjectIntelligenceConfig#edition_config}
	EditionConfig *string `field:"optional" json:"editionConfig" yaml:"editionConfig"`
	// filter block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.38.0/docs/resources/storage_control_project_intelligence_config#filter StorageControlProjectIntelligenceConfig#filter}
	Filter *StorageControlProjectIntelligenceConfigFilter `field:"optional" json:"filter" yaml:"filter"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.38.0/docs/resources/storage_control_project_intelligence_config#id StorageControlProjectIntelligenceConfig#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.38.0/docs/resources/storage_control_project_intelligence_config#timeouts StorageControlProjectIntelligenceConfig#timeouts}
	Timeouts *StorageControlProjectIntelligenceConfigTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

