// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package monitoringuptimecheckconfig


type MonitoringUptimeCheckConfigSyntheticMonitor struct {
	// cloud_function_v2 block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.50.0/docs/resources/monitoring_uptime_check_config#cloud_function_v2 MonitoringUptimeCheckConfig#cloud_function_v2}
	CloudFunctionV2 *MonitoringUptimeCheckConfigSyntheticMonitorCloudFunctionV2 `field:"required" json:"cloudFunctionV2" yaml:"cloudFunctionV2"`
}

