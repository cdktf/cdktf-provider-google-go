// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package sccv2organizationnotificationconfig


type SccV2OrganizationNotificationConfigTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.18.0/docs/resources/scc_v2_organization_notification_config#create SccV2OrganizationNotificationConfig#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.18.0/docs/resources/scc_v2_organization_notification_config#delete SccV2OrganizationNotificationConfig#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.18.0/docs/resources/scc_v2_organization_notification_config#update SccV2OrganizationNotificationConfig#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

