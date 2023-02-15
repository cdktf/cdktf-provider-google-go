package storagebucketacl

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type StorageBucketAclConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count *float64 `field:"optional" json:"count" yaml:"count"`
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
	// The name of the bucket it applies to.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/storage_bucket_acl#bucket StorageBucketAcl#bucket}
	Bucket *string `field:"required" json:"bucket" yaml:"bucket"`
	// Configure this ACL to be the default ACL.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/storage_bucket_acl#default_acl StorageBucketAcl#default_acl}
	DefaultAcl *string `field:"optional" json:"defaultAcl" yaml:"defaultAcl"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/storage_bucket_acl#id StorageBucketAcl#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// The canned GCS ACL to apply. Must be set if role_entity is not.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/storage_bucket_acl#predefined_acl StorageBucketAcl#predefined_acl}
	PredefinedAcl *string `field:"optional" json:"predefinedAcl" yaml:"predefinedAcl"`
	// List of role/entity pairs in the form ROLE:entity.
	//
	// See GCS Bucket ACL documentation  for more details. Must be set if predefined_acl is not.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/storage_bucket_acl#role_entity StorageBucketAcl#role_entity}
	RoleEntity *[]*string `field:"optional" json:"roleEntity" yaml:"roleEntity"`
}

