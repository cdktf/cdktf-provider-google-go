package cloudbuildbitbucketserverconfig

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type CloudbuildBitbucketServerConfigConfig struct {
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
	// Immutable.
	//
	// API Key that will be attached to webhook. Once this field has been set, it cannot be changed.
	// Changing this field will result in deleting/ recreating the resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/cloudbuild_bitbucket_server_config#api_key CloudbuildBitbucketServerConfig#api_key}
	ApiKey *string `field:"required" json:"apiKey" yaml:"apiKey"`
	// The ID to use for the BitbucketServerConfig, which will become the final component of the BitbucketServerConfig's resource name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/cloudbuild_bitbucket_server_config#config_id CloudbuildBitbucketServerConfig#config_id}
	ConfigId *string `field:"required" json:"configId" yaml:"configId"`
	// Immutable.
	//
	// The URI of the Bitbucket Server host. Once this field has been set, it cannot be changed.
	// If you need to change it, please create another BitbucketServerConfig.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/cloudbuild_bitbucket_server_config#host_uri CloudbuildBitbucketServerConfig#host_uri}
	HostUri *string `field:"required" json:"hostUri" yaml:"hostUri"`
	// The location of this bitbucket server config.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/cloudbuild_bitbucket_server_config#location CloudbuildBitbucketServerConfig#location}
	Location *string `field:"required" json:"location" yaml:"location"`
	// secrets block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/cloudbuild_bitbucket_server_config#secrets CloudbuildBitbucketServerConfig#secrets}
	Secrets *CloudbuildBitbucketServerConfigSecrets `field:"required" json:"secrets" yaml:"secrets"`
	// Username of the account Cloud Build will use on Bitbucket Server.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/cloudbuild_bitbucket_server_config#username CloudbuildBitbucketServerConfig#username}
	Username *string `field:"required" json:"username" yaml:"username"`
	// connected_repositories block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/cloudbuild_bitbucket_server_config#connected_repositories CloudbuildBitbucketServerConfig#connected_repositories}
	ConnectedRepositories interface{} `field:"optional" json:"connectedRepositories" yaml:"connectedRepositories"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/cloudbuild_bitbucket_server_config#id CloudbuildBitbucketServerConfig#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// The network to be used when reaching out to the Bitbucket Server instance.
	//
	// The VPC network must be enabled for private service connection.
	// This should be set if the Bitbucket Server instance is hosted on-premises and not reachable by public internet. If this field is left empty,
	// no network peering will occur and calls to the Bitbucket Server instance will be made over the public internet. Must be in the format
	// projects/{project}/global/networks/{network}, where {project} is a project number or id and {network} is the name of a VPC network in the project.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/cloudbuild_bitbucket_server_config#peered_network CloudbuildBitbucketServerConfig#peered_network}
	PeeredNetwork *string `field:"optional" json:"peeredNetwork" yaml:"peeredNetwork"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/cloudbuild_bitbucket_server_config#project CloudbuildBitbucketServerConfig#project}.
	Project *string `field:"optional" json:"project" yaml:"project"`
	// SSL certificate to use for requests to Bitbucket Server.
	//
	// The format should be PEM format but the extension can be one of .pem, .cer, or .crt.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/cloudbuild_bitbucket_server_config#ssl_ca CloudbuildBitbucketServerConfig#ssl_ca}
	SslCa *string `field:"optional" json:"sslCa" yaml:"sslCa"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/cloudbuild_bitbucket_server_config#timeouts CloudbuildBitbucketServerConfig#timeouts}
	Timeouts *CloudbuildBitbucketServerConfigTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

