// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package eventarcpipeline


type EventarcPipelineRetryPolicy struct {
	// The maximum number of delivery attempts for any message.
	//
	// The value must
	// be between 1 and 100.
	// The default value for this field is 5.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.47.0/docs/resources/eventarc_pipeline#max_attempts EventarcPipeline#max_attempts}
	MaxAttempts *float64 `field:"optional" json:"maxAttempts" yaml:"maxAttempts"`
	// The maximum amount of seconds to wait between retry attempts.
	//
	// The value
	// must be between 1 and 600.
	// The default value for this field is 60.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.47.0/docs/resources/eventarc_pipeline#max_retry_delay EventarcPipeline#max_retry_delay}
	MaxRetryDelay *string `field:"optional" json:"maxRetryDelay" yaml:"maxRetryDelay"`
	// The minimum amount of seconds to wait between retry attempts.
	//
	// The value
	// must be between 1 and 600.
	// The default value for this field is 5.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.47.0/docs/resources/eventarc_pipeline#min_retry_delay EventarcPipeline#min_retry_delay}
	MinRetryDelay *string `field:"optional" json:"minRetryDelay" yaml:"minRetryDelay"`
}

