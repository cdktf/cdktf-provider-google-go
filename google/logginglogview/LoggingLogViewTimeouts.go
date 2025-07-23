// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package logginglogview


type LoggingLogViewTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/logging_log_view#create LoggingLogView#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/logging_log_view#delete LoggingLogView#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/logging_log_view#update LoggingLogView#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

