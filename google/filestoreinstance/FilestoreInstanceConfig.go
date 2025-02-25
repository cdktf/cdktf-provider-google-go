// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package filestoreinstance

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type FilestoreInstanceConfig struct {
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
	// file_shares block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.22.0/docs/resources/filestore_instance#file_shares FilestoreInstance#file_shares}
	FileShares *FilestoreInstanceFileShares `field:"required" json:"fileShares" yaml:"fileShares"`
	// The resource name of the instance.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.22.0/docs/resources/filestore_instance#name FilestoreInstance#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// networks block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.22.0/docs/resources/filestore_instance#networks FilestoreInstance#networks}
	Networks interface{} `field:"required" json:"networks" yaml:"networks"`
	// The service tier of the instance. Possible values include: STANDARD, PREMIUM, BASIC_HDD, BASIC_SSD, HIGH_SCALE_SSD, ZONAL, REGIONAL and ENTERPRISE.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.22.0/docs/resources/filestore_instance#tier FilestoreInstance#tier}
	Tier *string `field:"required" json:"tier" yaml:"tier"`
	// Indicates whether the instance is protected against deletion.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.22.0/docs/resources/filestore_instance#deletion_protection_enabled FilestoreInstance#deletion_protection_enabled}
	DeletionProtectionEnabled interface{} `field:"optional" json:"deletionProtectionEnabled" yaml:"deletionProtectionEnabled"`
	// The reason for enabling deletion protection.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.22.0/docs/resources/filestore_instance#deletion_protection_reason FilestoreInstance#deletion_protection_reason}
	DeletionProtectionReason *string `field:"optional" json:"deletionProtectionReason" yaml:"deletionProtectionReason"`
	// A description of the instance.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.22.0/docs/resources/filestore_instance#description FilestoreInstance#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.22.0/docs/resources/filestore_instance#id FilestoreInstance#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// initial_replication block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.22.0/docs/resources/filestore_instance#initial_replication FilestoreInstance#initial_replication}
	InitialReplication *FilestoreInstanceInitialReplication `field:"optional" json:"initialReplication" yaml:"initialReplication"`
	// KMS key name used for data encryption.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.22.0/docs/resources/filestore_instance#kms_key_name FilestoreInstance#kms_key_name}
	KmsKeyName *string `field:"optional" json:"kmsKeyName" yaml:"kmsKeyName"`
	// Resource labels to represent user-provided metadata.
	//
	// **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
	// Please refer to the field 'effective_labels' for all of the labels present on the resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.22.0/docs/resources/filestore_instance#labels FilestoreInstance#labels}
	Labels *map[string]*string `field:"optional" json:"labels" yaml:"labels"`
	// The name of the location of the instance. This can be a region for ENTERPRISE tier instances.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.22.0/docs/resources/filestore_instance#location FilestoreInstance#location}
	Location *string `field:"optional" json:"location" yaml:"location"`
	// performance_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.22.0/docs/resources/filestore_instance#performance_config FilestoreInstance#performance_config}
	PerformanceConfig *FilestoreInstancePerformanceConfig `field:"optional" json:"performanceConfig" yaml:"performanceConfig"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.22.0/docs/resources/filestore_instance#project FilestoreInstance#project}.
	Project *string `field:"optional" json:"project" yaml:"project"`
	// Either NFSv3, for using NFS version 3 as file sharing protocol, or NFSv4.1, for using NFS version 4.1 as file sharing protocol. NFSv4.1 can be used with HIGH_SCALE_SSD, ZONAL, REGIONAL and ENTERPRISE. The default is NFSv3. Default value: "NFS_V3" Possible values: ["NFS_V3", "NFS_V4_1"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.22.0/docs/resources/filestore_instance#protocol FilestoreInstance#protocol}
	Protocol *string `field:"optional" json:"protocol" yaml:"protocol"`
	// A map of resource manager tags.
	//
	// Resource manager tag keys
	// and values have the same definition as resource manager
	// tags. Keys must be in the format tagKeys/{tag_key_id},
	// and values are in the format tagValues/456. The field is
	// ignored when empty. The field is immutable and causes
	// resource replacement when mutated. This field is only set
	// at create time and modifying this field after creation
	// will trigger recreation. To apply tags to an existing
	// resource, see the 'google_tags_tag_value' resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.22.0/docs/resources/filestore_instance#tags FilestoreInstance#tags}
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.22.0/docs/resources/filestore_instance#timeouts FilestoreInstance#timeouts}
	Timeouts *FilestoreInstanceTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
	// The name of the Filestore zone of the instance.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.22.0/docs/resources/filestore_instance#zone FilestoreInstance#zone}
	Zone *string `field:"optional" json:"zone" yaml:"zone"`
}

