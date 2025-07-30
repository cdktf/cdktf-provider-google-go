// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package sccv2projectnotificationconfig

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type SccV2ProjectNotificationConfigConfig struct {
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
	// This must be unique within the project.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.46.0/docs/resources/scc_v2_project_notification_config#config_id SccV2ProjectNotificationConfig#config_id}
	ConfigId *string `field:"required" json:"configId" yaml:"configId"`
	// streaming_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.46.0/docs/resources/scc_v2_project_notification_config#streaming_config SccV2ProjectNotificationConfig#streaming_config}
	StreamingConfig *SccV2ProjectNotificationConfigStreamingConfig `field:"required" json:"streamingConfig" yaml:"streamingConfig"`
	// The description of the notification config (max of 1024 characters).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.46.0/docs/resources/scc_v2_project_notification_config#description SccV2ProjectNotificationConfig#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.46.0/docs/resources/scc_v2_project_notification_config#id SccV2ProjectNotificationConfig#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Location ID of the parent organization. Only global is supported at the moment.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.46.0/docs/resources/scc_v2_project_notification_config#location SccV2ProjectNotificationConfig#location}
	Location *string `field:"optional" json:"location" yaml:"location"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.46.0/docs/resources/scc_v2_project_notification_config#project SccV2ProjectNotificationConfig#project}.
	Project *string `field:"optional" json:"project" yaml:"project"`
	// The Pub/Sub topic to send notifications to. Its format is "projects/[project_id]/topics/[topic]".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.46.0/docs/resources/scc_v2_project_notification_config#pubsub_topic SccV2ProjectNotificationConfig#pubsub_topic}
	PubsubTopic *string `field:"optional" json:"pubsubTopic" yaml:"pubsubTopic"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.46.0/docs/resources/scc_v2_project_notification_config#timeouts SccV2ProjectNotificationConfig#timeouts}
	Timeouts *SccV2ProjectNotificationConfigTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

