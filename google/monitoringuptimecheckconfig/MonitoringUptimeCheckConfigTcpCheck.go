// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package monitoringuptimecheckconfig


type MonitoringUptimeCheckConfigTcpCheck struct {
	// The port to the page to run the check against.
	//
	// Will be combined with host (specified within the MonitoredResource) to construct the full URL.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.83.0/docs/resources/monitoring_uptime_check_config#port MonitoringUptimeCheckConfig#port}
	Port *float64 `field:"required" json:"port" yaml:"port"`
}

