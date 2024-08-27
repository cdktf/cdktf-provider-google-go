// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package monitoringcustomservice


type MonitoringCustomServiceTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.43.0/docs/resources/monitoring_custom_service#create MonitoringCustomService#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.43.0/docs/resources/monitoring_custom_service#delete MonitoringCustomService#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.43.0/docs/resources/monitoring_custom_service#update MonitoringCustomService#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

