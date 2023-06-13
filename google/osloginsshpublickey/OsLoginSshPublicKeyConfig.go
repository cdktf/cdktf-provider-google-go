package osloginsshpublickey

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type OsLoginSshPublicKeyConfig struct {
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
	// Public key text in SSH format, defined by RFC4253 section 6.6.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/resources/os_login_ssh_public_key#key OsLoginSshPublicKey#key}
	Key *string `field:"required" json:"key" yaml:"key"`
	// The user email.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/resources/os_login_ssh_public_key#user OsLoginSshPublicKey#user}
	User *string `field:"required" json:"user" yaml:"user"`
	// An expiration time in microseconds since epoch.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/resources/os_login_ssh_public_key#expiration_time_usec OsLoginSshPublicKey#expiration_time_usec}
	ExpirationTimeUsec *string `field:"optional" json:"expirationTimeUsec" yaml:"expirationTimeUsec"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/resources/os_login_ssh_public_key#id OsLoginSshPublicKey#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// The project ID of the Google Cloud Platform project.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/resources/os_login_ssh_public_key#project OsLoginSshPublicKey#project}
	Project *string `field:"optional" json:"project" yaml:"project"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/resources/os_login_ssh_public_key#timeouts OsLoginSshPublicKey#timeouts}
	Timeouts *OsLoginSshPublicKeyTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

