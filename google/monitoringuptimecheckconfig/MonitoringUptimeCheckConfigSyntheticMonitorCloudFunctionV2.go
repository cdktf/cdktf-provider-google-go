// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package monitoringuptimecheckconfig


type MonitoringUptimeCheckConfigSyntheticMonitorCloudFunctionV2 struct {
	// The fully qualified name of the cloud function resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.36.0/docs/resources/monitoring_uptime_check_config#name MonitoringUptimeCheckConfig#name}
	Name *string `field:"required" json:"name" yaml:"name"`
}

