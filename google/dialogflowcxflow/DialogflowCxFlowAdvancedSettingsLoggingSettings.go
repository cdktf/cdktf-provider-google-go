// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dialogflowcxflow


type DialogflowCxFlowAdvancedSettingsLoggingSettings struct {
	// Enables consent-based end-user input redaction, if true, a pre-defined session parameter **$session.params.conversation-redaction** will be used to determine if the utterance should be redacted.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.46.0/docs/resources/dialogflow_cx_flow#enable_consent_based_redaction DialogflowCxFlow#enable_consent_based_redaction}
	EnableConsentBasedRedaction interface{} `field:"optional" json:"enableConsentBasedRedaction" yaml:"enableConsentBasedRedaction"`
	// Enables DF Interaction logging.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.46.0/docs/resources/dialogflow_cx_flow#enable_interaction_logging DialogflowCxFlow#enable_interaction_logging}
	EnableInteractionLogging interface{} `field:"optional" json:"enableInteractionLogging" yaml:"enableInteractionLogging"`
	// Enables Google Cloud Logging.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.46.0/docs/resources/dialogflow_cx_flow#enable_stackdriver_logging DialogflowCxFlow#enable_stackdriver_logging}
	EnableStackdriverLogging interface{} `field:"optional" json:"enableStackdriverLogging" yaml:"enableStackdriverLogging"`
}

