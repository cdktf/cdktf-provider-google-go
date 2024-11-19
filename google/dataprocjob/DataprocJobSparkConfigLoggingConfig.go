// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dataprocjob


type DataprocJobSparkConfigLoggingConfig struct {
	// Optional.
	//
	// The per-package log levels for the driver. This may include 'root' package name to configure rootLogger. Examples: 'com.google = FATAL', 'root = INFO', 'org.apache = DEBUG'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.12.0/docs/resources/dataproc_job#driver_log_levels DataprocJob#driver_log_levels}
	DriverLogLevels *map[string]*string `field:"required" json:"driverLogLevels" yaml:"driverLogLevels"`
}

