package cloudbuildv2repository

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type Cloudbuildv2RepositoryConfig struct {
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
	// Name of the repository.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/cloudbuildv2_repository#name Cloudbuildv2Repository#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The connection for the resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/cloudbuildv2_repository#parent_connection Cloudbuildv2Repository#parent_connection}
	ParentConnection *string `field:"required" json:"parentConnection" yaml:"parentConnection"`
	// Required. Git Clone HTTPS URI.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/cloudbuildv2_repository#remote_uri Cloudbuildv2Repository#remote_uri}
	RemoteUri *string `field:"required" json:"remoteUri" yaml:"remoteUri"`
	// Allows clients to store small amounts of arbitrary data.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/cloudbuildv2_repository#annotations Cloudbuildv2Repository#annotations}
	Annotations *map[string]*string `field:"optional" json:"annotations" yaml:"annotations"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/cloudbuildv2_repository#id Cloudbuildv2Repository#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// The location for the resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/cloudbuildv2_repository#location Cloudbuildv2Repository#location}
	Location *string `field:"optional" json:"location" yaml:"location"`
	// The project for the resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/cloudbuildv2_repository#project Cloudbuildv2Repository#project}
	Project *string `field:"optional" json:"project" yaml:"project"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/cloudbuildv2_repository#timeouts Cloudbuildv2Repository#timeouts}
	Timeouts *Cloudbuildv2RepositoryTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

