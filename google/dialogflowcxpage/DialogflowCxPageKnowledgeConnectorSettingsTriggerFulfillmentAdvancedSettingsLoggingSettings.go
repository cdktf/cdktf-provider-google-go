// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dialogflowcxpage


type DialogflowCxPageKnowledgeConnectorSettingsTriggerFulfillmentAdvancedSettingsLoggingSettings struct {
	// Enables consent-based end-user input redaction, if true, a pre-defined session parameter **$session.params.conversation-redaction** will be used to determine if the utterance should be redacted.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.50.0/docs/resources/dialogflow_cx_page#enable_consent_based_redaction DialogflowCxPage#enable_consent_based_redaction}
	EnableConsentBasedRedaction interface{} `field:"optional" json:"enableConsentBasedRedaction" yaml:"enableConsentBasedRedaction"`
	// Enables DF Interaction logging.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.50.0/docs/resources/dialogflow_cx_page#enable_interaction_logging DialogflowCxPage#enable_interaction_logging}
	EnableInteractionLogging interface{} `field:"optional" json:"enableInteractionLogging" yaml:"enableInteractionLogging"`
	// Enables Google Cloud Logging.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.50.0/docs/resources/dialogflow_cx_page#enable_stackdriver_logging DialogflowCxPage#enable_stackdriver_logging}
	EnableStackdriverLogging interface{} `field:"optional" json:"enableStackdriverLogging" yaml:"enableStackdriverLogging"`
}

