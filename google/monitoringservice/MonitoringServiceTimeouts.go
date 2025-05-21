// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package monitoringservice


type MonitoringServiceTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.36.0/docs/resources/monitoring_service#create MonitoringService#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.36.0/docs/resources/monitoring_service#delete MonitoringService#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.36.0/docs/resources/monitoring_service#update MonitoringService#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

