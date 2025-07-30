// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package loggingfoldersettings

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type LoggingFolderSettingsConfig struct {
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
	// The folder for which to retrieve settings.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.46.0/docs/resources/logging_folder_settings#folder LoggingFolderSettings#folder}
	Folder *string `field:"required" json:"folder" yaml:"folder"`
	// If set to true, the _Default sink in newly created projects and folders will created in a disabled state.
	//
	// This can be used to automatically disable log storage if there is already an aggregated sink configured in the hierarchy. The _Default sink can be re-enabled manually if needed.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.46.0/docs/resources/logging_folder_settings#disable_default_sink LoggingFolderSettings#disable_default_sink}
	DisableDefaultSink interface{} `field:"optional" json:"disableDefaultSink" yaml:"disableDefaultSink"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.46.0/docs/resources/logging_folder_settings#id LoggingFolderSettings#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// The resource name for the configured Cloud KMS key.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.46.0/docs/resources/logging_folder_settings#kms_key_name LoggingFolderSettings#kms_key_name}
	KmsKeyName *string `field:"optional" json:"kmsKeyName" yaml:"kmsKeyName"`
	// The storage location that Cloud Logging will use to create new resources when a location is needed but not explicitly provided.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.46.0/docs/resources/logging_folder_settings#storage_location LoggingFolderSettings#storage_location}
	StorageLocation *string `field:"optional" json:"storageLocation" yaml:"storageLocation"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.46.0/docs/resources/logging_folder_settings#timeouts LoggingFolderSettings#timeouts}
	Timeouts *LoggingFolderSettingsTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

