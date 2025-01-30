// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package sccv2projectnotificationconfig


type SccV2ProjectNotificationConfigTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.18.1/docs/resources/scc_v2_project_notification_config#create SccV2ProjectNotificationConfig#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.18.1/docs/resources/scc_v2_project_notification_config#delete SccV2ProjectNotificationConfig#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.18.1/docs/resources/scc_v2_project_notification_config#update SccV2ProjectNotificationConfig#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

