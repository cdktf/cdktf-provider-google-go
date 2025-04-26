// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package eventarcmessagebus


type EventarcMessageBusLoggingConfig struct {
	// Optional.
	//
	// The minimum severity of logs that will be sent to Stackdriver/Platform
	// Telemetry. Logs at severitiy ≥ this value will be sent, unless it is NONE. Possible values: ["NONE", "DEBUG", "INFO", "NOTICE", "WARNING", "ERROR", "CRITICAL", "ALERT", "EMERGENCY"]
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.32.0/docs/resources/eventarc_message_bus#log_severity EventarcMessageBus#log_severity}
	LogSeverity *string `field:"optional" json:"logSeverity" yaml:"logSeverity"`
}

