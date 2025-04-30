// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package sccfoldernotificationconfig

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type SccFolderNotificationConfigConfig struct {
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
	// This must be unique within the organization.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.33.0/docs/resources/scc_folder_notification_config#config_id SccFolderNotificationConfig#config_id}
	ConfigId *string `field:"required" json:"configId" yaml:"configId"`
	// Numerical ID of the parent folder.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.33.0/docs/resources/scc_folder_notification_config#folder SccFolderNotificationConfig#folder}
	Folder *string `field:"required" json:"folder" yaml:"folder"`
	// The Pub/Sub topic to send notifications to. Its format is "projects/[project_id]/topics/[topic]".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.33.0/docs/resources/scc_folder_notification_config#pubsub_topic SccFolderNotificationConfig#pubsub_topic}
	PubsubTopic *string `field:"required" json:"pubsubTopic" yaml:"pubsubTopic"`
	// streaming_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.33.0/docs/resources/scc_folder_notification_config#streaming_config SccFolderNotificationConfig#streaming_config}
	StreamingConfig *SccFolderNotificationConfigStreamingConfig `field:"required" json:"streamingConfig" yaml:"streamingConfig"`
	// The description of the notification config (max of 1024 characters).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.33.0/docs/resources/scc_folder_notification_config#description SccFolderNotificationConfig#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.33.0/docs/resources/scc_folder_notification_config#id SccFolderNotificationConfig#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.33.0/docs/resources/scc_folder_notification_config#timeouts SccFolderNotificationConfig#timeouts}
	Timeouts *SccFolderNotificationConfigTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

