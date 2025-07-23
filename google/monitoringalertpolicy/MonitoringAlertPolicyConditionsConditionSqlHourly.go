// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package monitoringalertpolicy


type MonitoringAlertPolicyConditionsConditionSqlHourly struct {
	// Number of hours between runs.
	//
	// The interval must be greater than or
	// equal to 1 hour and less than or equal to 48 hours.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/monitoring_alert_policy#periodicity MonitoringAlertPolicy#periodicity}
	Periodicity *float64 `field:"required" json:"periodicity" yaml:"periodicity"`
	// The number of minutes after the hour (in UTC) to run the query.
	//
	// Must be greater than or equal to 0 minutes and less than or equal to
	// 59 minutes.  If left unspecified, then an arbitrary offset is used.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/monitoring_alert_policy#minute_offset MonitoringAlertPolicy#minute_offset}
	MinuteOffset *float64 `field:"optional" json:"minuteOffset" yaml:"minuteOffset"`
}

