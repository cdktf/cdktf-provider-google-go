// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package monitoringslo


type MonitoringSloWindowsBasedSliMetricMeanInRangeRange struct {
	// max value for the range (inclusive). If not given, will be set to "infinity", defining an open range ">= range.min".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.16.0/docs/resources/monitoring_slo#max MonitoringSlo#max}
	Max *float64 `field:"optional" json:"max" yaml:"max"`
	// Min value for the range (inclusive). If not given, will be set to "-infinity", defining an open range "< range.max".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.16.0/docs/resources/monitoring_slo#min MonitoringSlo#min}
	Min *float64 `field:"optional" json:"min" yaml:"min"`
}

