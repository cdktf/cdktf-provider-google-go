// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package monitoringuptimecheckconfig


type MonitoringUptimeCheckConfigHttpCheckAuthInfo struct {
	// The username to authenticate.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.0/docs/resources/monitoring_uptime_check_config#username MonitoringUptimeCheckConfig#username}
	Username *string `field:"required" json:"username" yaml:"username"`
	// The password to authenticate.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.0/docs/resources/monitoring_uptime_check_config#password MonitoringUptimeCheckConfig#password}
	Password *string `field:"optional" json:"password" yaml:"password"`
	// The password to authenticate.
	//
	// Note: This property is write-only and will not be read from the API. For more info see [updating write-only attributes](/docs/providers/google/guides/using_write_only_attributes.html#updating-write-only-attributes)
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.0/docs/resources/monitoring_uptime_check_config#password_wo MonitoringUptimeCheckConfig#password_wo}
	PasswordWo *string `field:"optional" json:"passwordWo" yaml:"passwordWo"`
	// Triggers update of password_wo write-only. For more info see [updating write-only attributes](/docs/providers/google/guides/using_write_only_attributes.html#updating-write-only-attributes).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.0/docs/resources/monitoring_uptime_check_config#password_wo_version MonitoringUptimeCheckConfig#password_wo_version}
	PasswordWoVersion *string `field:"optional" json:"passwordWoVersion" yaml:"passwordWoVersion"`
}

