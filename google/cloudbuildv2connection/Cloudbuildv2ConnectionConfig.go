package cloudbuildv2connection

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type Cloudbuildv2ConnectionConfig struct {
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
	// The location for the resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/cloudbuildv2_connection#location Cloudbuildv2Connection#location}
	Location *string `field:"required" json:"location" yaml:"location"`
	// Immutable. The resource name of the connection, in the format `projects/{project}/locations/{location}/connections/{connection_id}`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/cloudbuildv2_connection#name Cloudbuildv2Connection#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Allows clients to store small amounts of arbitrary data.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/cloudbuildv2_connection#annotations Cloudbuildv2Connection#annotations}
	Annotations *map[string]*string `field:"optional" json:"annotations" yaml:"annotations"`
	// If disabled is set to true, functionality is disabled for this connection.
	//
	// Repository based API methods and webhooks processing for repositories in this connection will be disabled.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/cloudbuildv2_connection#disabled Cloudbuildv2Connection#disabled}
	Disabled interface{} `field:"optional" json:"disabled" yaml:"disabled"`
	// github_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/cloudbuildv2_connection#github_config Cloudbuildv2Connection#github_config}
	GithubConfig *Cloudbuildv2ConnectionGithubConfig `field:"optional" json:"githubConfig" yaml:"githubConfig"`
	// github_enterprise_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/cloudbuildv2_connection#github_enterprise_config Cloudbuildv2Connection#github_enterprise_config}
	GithubEnterpriseConfig *Cloudbuildv2ConnectionGithubEnterpriseConfig `field:"optional" json:"githubEnterpriseConfig" yaml:"githubEnterpriseConfig"`
	// gitlab_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/cloudbuildv2_connection#gitlab_config Cloudbuildv2Connection#gitlab_config}
	GitlabConfig *Cloudbuildv2ConnectionGitlabConfig `field:"optional" json:"gitlabConfig" yaml:"gitlabConfig"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/cloudbuildv2_connection#id Cloudbuildv2Connection#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// The project for the resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/cloudbuildv2_connection#project Cloudbuildv2Connection#project}
	Project *string `field:"optional" json:"project" yaml:"project"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/cloudbuildv2_connection#timeouts Cloudbuildv2Connection#timeouts}
	Timeouts *Cloudbuildv2ConnectionTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

