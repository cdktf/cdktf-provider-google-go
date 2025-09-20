// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package monitoringalertpolicy


type MonitoringAlertPolicyConditionsConditionSqlDailyExecutionTime struct {
	// Hours of a day in 24 hour format.
	//
	// Must be greater than or equal
	// to 0 and typically must be less than or equal to 23. An API may
	// choose to allow the value "24:00:00" for scenarios like business
	// closing time.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.50.0/docs/resources/monitoring_alert_policy#hours MonitoringAlertPolicy#hours}
	Hours *float64 `field:"optional" json:"hours" yaml:"hours"`
	// Minutes of an hour. Must be greater than or equal to 0 and less than or equal to 59.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.50.0/docs/resources/monitoring_alert_policy#minutes MonitoringAlertPolicy#minutes}
	Minutes *float64 `field:"optional" json:"minutes" yaml:"minutes"`
	// Fractions of seconds, in nanoseconds. Must be greater than or equal to 0 and less than or equal to 999,999,999.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.50.0/docs/resources/monitoring_alert_policy#nanos MonitoringAlertPolicy#nanos}
	Nanos *float64 `field:"optional" json:"nanos" yaml:"nanos"`
	// Seconds of a minute.
	//
	// Must be greater than or equal to 0 and
	// typically must be less than or equal to 59. An API may allow the
	// value 60 if it allows leap-seconds.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.50.0/docs/resources/monitoring_alert_policy#seconds MonitoringAlertPolicy#seconds}
	Seconds *float64 `field:"optional" json:"seconds" yaml:"seconds"`
}

