// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datagoogleartifactregistrydockerimage

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataGoogleArtifactRegistryDockerImageConfig struct {
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
	// The image name to fetch.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.23.0/docs/data-sources/artifact_registry_docker_image#image_name DataGoogleArtifactRegistryDockerImage#image_name}
	ImageName *string `field:"required" json:"imageName" yaml:"imageName"`
	// The region of the artifact registry repository. For example, "us-west1".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.23.0/docs/data-sources/artifact_registry_docker_image#location DataGoogleArtifactRegistryDockerImage#location}
	Location *string `field:"required" json:"location" yaml:"location"`
	// The last part of the repository name to fetch from.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.23.0/docs/data-sources/artifact_registry_docker_image#repository_id DataGoogleArtifactRegistryDockerImage#repository_id}
	RepositoryId *string `field:"required" json:"repositoryId" yaml:"repositoryId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.23.0/docs/data-sources/artifact_registry_docker_image#id DataGoogleArtifactRegistryDockerImage#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Project ID of the project.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.23.0/docs/data-sources/artifact_registry_docker_image#project DataGoogleArtifactRegistryDockerImage#project}
	Project *string `field:"optional" json:"project" yaml:"project"`
}

