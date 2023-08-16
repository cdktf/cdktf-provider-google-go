package loggingfolderbucketconfig

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type LoggingFolderBucketConfigConfig struct {
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
	// The name of the logging bucket. Logging automatically creates two log buckets: _Required and _Default.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/logging_folder_bucket_config#bucket_id LoggingFolderBucketConfig#bucket_id}
	BucketId *string `field:"required" json:"bucketId" yaml:"bucketId"`
	// The parent resource that contains the logging bucket.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/logging_folder_bucket_config#folder LoggingFolderBucketConfig#folder}
	Folder *string `field:"required" json:"folder" yaml:"folder"`
	// The location of the bucket.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/logging_folder_bucket_config#location LoggingFolderBucketConfig#location}
	Location *string `field:"required" json:"location" yaml:"location"`
	// cmek_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/logging_folder_bucket_config#cmek_settings LoggingFolderBucketConfig#cmek_settings}
	CmekSettings *LoggingFolderBucketConfigCmekSettings `field:"optional" json:"cmekSettings" yaml:"cmekSettings"`
	// An optional description for this bucket.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/logging_folder_bucket_config#description LoggingFolderBucketConfig#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/logging_folder_bucket_config#id LoggingFolderBucketConfig#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Logs will be retained by default for this amount of time, after which they will automatically be deleted.
	//
	// The minimum retention period is 1 day. If this value is set to zero at bucket creation time, the default time of 30 days will be used.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/logging_folder_bucket_config#retention_days LoggingFolderBucketConfig#retention_days}
	RetentionDays *float64 `field:"optional" json:"retentionDays" yaml:"retentionDays"`
}

