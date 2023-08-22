package resourcemanagerlien

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ResourceManagerLienConfig struct {
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
	// A stable, user-visible/meaningful string identifying the origin of the Lien, intended to be inspected programmatically. Maximum length of 200 characters.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/resource_manager_lien#origin ResourceManagerLien#origin}
	Origin *string `field:"required" json:"origin" yaml:"origin"`
	// A reference to the resource this Lien is attached to.
	//
	// The server will validate the parent against those for which Liens are supported.
	// Since a variety of objects can have Liens against them, you must provide the type
	// prefix (e.g. "projects/my-project-name").
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/resource_manager_lien#parent ResourceManagerLien#parent}
	Parent *string `field:"required" json:"parent" yaml:"parent"`
	// Concise user-visible strings indicating why an action cannot be performed on a resource. Maximum length of 200 characters.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/resource_manager_lien#reason ResourceManagerLien#reason}
	Reason *string `field:"required" json:"reason" yaml:"reason"`
	// The types of operations which should be blocked as a result of this Lien.
	//
	// Each value should correspond to an IAM permission. The server will validate
	// the permissions against those for which Liens are supported.  An empty
	// list is meaningless and will be rejected.
	// e.g. ['resourcemanager.projects.delete']
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/resource_manager_lien#restrictions ResourceManagerLien#restrictions}
	Restrictions *[]*string `field:"required" json:"restrictions" yaml:"restrictions"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/resource_manager_lien#id ResourceManagerLien#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/resource_manager_lien#timeouts ResourceManagerLien#timeouts}
	Timeouts *ResourceManagerLienTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

