// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package monitoringuptimecheckconfig


type MonitoringUptimeCheckConfigHttpCheckAuthInfo struct {
	// The username to authenticate.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.39.0/docs/resources/monitoring_uptime_check_config#username MonitoringUptimeCheckConfig#username}
	Username *string `field:"required" json:"username" yaml:"username"`
	// The password to authenticate.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.39.0/docs/resources/monitoring_uptime_check_config#password MonitoringUptimeCheckConfig#password}
	Password *string `field:"optional" json:"password" yaml:"password"`
	// The password to authenticate.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.39.0/docs/resources/monitoring_uptime_check_config#password_wo MonitoringUptimeCheckConfig#password_wo}
	PasswordWo *string `field:"optional" json:"passwordWo" yaml:"passwordWo"`
	// The password write-only version.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.39.0/docs/resources/monitoring_uptime_check_config#password_wo_version MonitoringUptimeCheckConfig#password_wo_version}
	PasswordWoVersion *string `field:"optional" json:"passwordWoVersion" yaml:"passwordWoVersion"`
}

