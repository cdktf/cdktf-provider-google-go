// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package monitoringalertpolicy


type MonitoringAlertPolicyAlertStrategyNotificationRateLimit struct {
	// Not more than one notification per period.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.15.0/docs/resources/monitoring_alert_policy#period MonitoringAlertPolicy#period}
	Period *string `field:"optional" json:"period" yaml:"period"`
}

