package filestoreinstance

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type FilestoreInstanceConfig struct {
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
	// file_shares block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/filestore_instance#file_shares FilestoreInstance#file_shares}
	FileShares *FilestoreInstanceFileShares `field:"required" json:"fileShares" yaml:"fileShares"`
	// The resource name of the instance.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/filestore_instance#name FilestoreInstance#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// networks block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/filestore_instance#networks FilestoreInstance#networks}
	Networks interface{} `field:"required" json:"networks" yaml:"networks"`
	// The service tier of the instance. Possible values include: STANDARD, PREMIUM, BASIC_HDD, BASIC_SSD, HIGH_SCALE_SSD and ENTERPRISE.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/filestore_instance#tier FilestoreInstance#tier}
	Tier *string `field:"required" json:"tier" yaml:"tier"`
	// A description of the instance.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/filestore_instance#description FilestoreInstance#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/filestore_instance#id FilestoreInstance#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// KMS key name used for data encryption.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/filestore_instance#kms_key_name FilestoreInstance#kms_key_name}
	KmsKeyName *string `field:"optional" json:"kmsKeyName" yaml:"kmsKeyName"`
	// Resource labels to represent user-provided metadata.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/filestore_instance#labels FilestoreInstance#labels}
	Labels *map[string]*string `field:"optional" json:"labels" yaml:"labels"`
	// The name of the location of the instance. This can be a region for ENTERPRISE tier instances.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/filestore_instance#location FilestoreInstance#location}
	Location *string `field:"optional" json:"location" yaml:"location"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/filestore_instance#project FilestoreInstance#project}.
	Project *string `field:"optional" json:"project" yaml:"project"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/filestore_instance#timeouts FilestoreInstance#timeouts}
	Timeouts *FilestoreInstanceTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
	// The name of the Filestore zone of the instance.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/filestore_instance#zone FilestoreInstance#zone}
	Zone *string `field:"optional" json:"zone" yaml:"zone"`
}

