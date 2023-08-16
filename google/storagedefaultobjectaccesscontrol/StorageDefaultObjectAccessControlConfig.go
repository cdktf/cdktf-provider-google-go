package storagedefaultobjectaccesscontrol

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type StorageDefaultObjectAccessControlConfig struct {
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
	// The name of the bucket.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/storage_default_object_access_control#bucket StorageDefaultObjectAccessControl#bucket}
	Bucket *string `field:"required" json:"bucket" yaml:"bucket"`
	// The entity holding the permission, in one of the following forms: user-{{userId}} user-{{email}} (such as "user-liz@example.com") group-{{groupId}} group-{{email}} (such as "group-example@googlegroups.com") domain-{{domain}} (such as "domain-example.com") project-team-{{projectId}} allUsers allAuthenticatedUsers.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/storage_default_object_access_control#entity StorageDefaultObjectAccessControl#entity}
	Entity *string `field:"required" json:"entity" yaml:"entity"`
	// The access permission for the entity. Possible values: ["OWNER", "READER"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/storage_default_object_access_control#role StorageDefaultObjectAccessControl#role}
	Role *string `field:"required" json:"role" yaml:"role"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/storage_default_object_access_control#id StorageDefaultObjectAccessControl#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// The name of the object, if applied to an object.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/storage_default_object_access_control#object StorageDefaultObjectAccessControl#object}
	Object *string `field:"optional" json:"object" yaml:"object"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/storage_default_object_access_control#timeouts StorageDefaultObjectAccessControl#timeouts}
	Timeouts *StorageDefaultObjectAccessControlTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

