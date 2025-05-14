// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package monitoringmetricdescriptor


type MonitoringMetricDescriptorTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.35.0/docs/resources/monitoring_metric_descriptor#create MonitoringMetricDescriptor#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.35.0/docs/resources/monitoring_metric_descriptor#delete MonitoringMetricDescriptor#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.35.0/docs/resources/monitoring_metric_descriptor#update MonitoringMetricDescriptor#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

