// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package eventarcgoogleapisource


type EventarcGoogleApiSourceLoggingConfig struct {
	// The minimum severity of logs that will be sent to Stackdriver/Platform Telemetry.
	//
	// Logs at severitiy ≥ this value will be sent, unless it is NONE. Possible values: ["NONE", "DEBUG", "INFO", "NOTICE", "WARNING", "ERROR", "CRITICAL", "ALERT", "EMERGENCY"]
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.44.0/docs/resources/eventarc_google_api_source#log_severity EventarcGoogleApiSource#log_severity}
	LogSeverity *string `field:"optional" json:"logSeverity" yaml:"logSeverity"`
}

