package datagooglestoragebucketobject

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataGoogleStorageBucketObjectConfig struct {
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
	// The name of the containing bucket.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/d/storage_bucket_object#bucket DataGoogleStorageBucketObject#bucket}
	Bucket *string `field:"optional" json:"bucket" yaml:"bucket"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/d/storage_bucket_object#id DataGoogleStorageBucketObject#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// The name of the object. If you're interpolating the name of this object, see output_name instead.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/d/storage_bucket_object#name DataGoogleStorageBucketObject#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
}

