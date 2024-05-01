// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datafusioninstance


type DataFusionInstanceEventPublishConfig struct {
	// Option to enable Event Publishing.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.27.0/docs/resources/data_fusion_instance#enabled DataFusionInstance#enabled}
	Enabled interface{} `field:"required" json:"enabled" yaml:"enabled"`
	// The resource name of the Pub/Sub topic. Format: projects/{projectId}/topics/{topic_id}.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.27.0/docs/resources/data_fusion_instance#topic DataFusionInstance#topic}
	Topic *string `field:"required" json:"topic" yaml:"topic"`
}

