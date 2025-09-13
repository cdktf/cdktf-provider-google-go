// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dialogflowconversationprofile

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DialogflowConversationProfileConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktf.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// Required. Human readable name for this profile. Max length 1024 bytes.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.3/docs/resources/dialogflow_conversation_profile#display_name DialogflowConversationProfile#display_name}
	DisplayName *string `field:"required" json:"displayName" yaml:"displayName"`
	// desc.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.3/docs/resources/dialogflow_conversation_profile#location DialogflowConversationProfile#location}
	Location *string `field:"required" json:"location" yaml:"location"`
	// automated_agent_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.3/docs/resources/dialogflow_conversation_profile#automated_agent_config DialogflowConversationProfile#automated_agent_config}
	AutomatedAgentConfig *DialogflowConversationProfileAutomatedAgentConfig `field:"optional" json:"automatedAgentConfig" yaml:"automatedAgentConfig"`
	// human_agent_assistant_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.3/docs/resources/dialogflow_conversation_profile#human_agent_assistant_config DialogflowConversationProfile#human_agent_assistant_config}
	HumanAgentAssistantConfig *DialogflowConversationProfileHumanAgentAssistantConfig `field:"optional" json:"humanAgentAssistantConfig" yaml:"humanAgentAssistantConfig"`
	// human_agent_handoff_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.3/docs/resources/dialogflow_conversation_profile#human_agent_handoff_config DialogflowConversationProfile#human_agent_handoff_config}
	HumanAgentHandoffConfig *DialogflowConversationProfileHumanAgentHandoffConfig `field:"optional" json:"humanAgentHandoffConfig" yaml:"humanAgentHandoffConfig"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.3/docs/resources/dialogflow_conversation_profile#id DialogflowConversationProfile#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Language code for the conversation profile. This should be a BCP-47 language tag.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.3/docs/resources/dialogflow_conversation_profile#language_code DialogflowConversationProfile#language_code}
	LanguageCode *string `field:"optional" json:"languageCode" yaml:"languageCode"`
	// logging_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.3/docs/resources/dialogflow_conversation_profile#logging_config DialogflowConversationProfile#logging_config}
	LoggingConfig *DialogflowConversationProfileLoggingConfig `field:"optional" json:"loggingConfig" yaml:"loggingConfig"`
	// new_message_event_notification_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.3/docs/resources/dialogflow_conversation_profile#new_message_event_notification_config DialogflowConversationProfile#new_message_event_notification_config}
	NewMessageEventNotificationConfig *DialogflowConversationProfileNewMessageEventNotificationConfig `field:"optional" json:"newMessageEventNotificationConfig" yaml:"newMessageEventNotificationConfig"`
	// notification_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.3/docs/resources/dialogflow_conversation_profile#notification_config DialogflowConversationProfile#notification_config}
	NotificationConfig *DialogflowConversationProfileNotificationConfig `field:"optional" json:"notificationConfig" yaml:"notificationConfig"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.3/docs/resources/dialogflow_conversation_profile#project DialogflowConversationProfile#project}.
	Project *string `field:"optional" json:"project" yaml:"project"`
	// Name of the CX SecuritySettings reference for the agent.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.3/docs/resources/dialogflow_conversation_profile#security_settings DialogflowConversationProfile#security_settings}
	SecuritySettings *string `field:"optional" json:"securitySettings" yaml:"securitySettings"`
	// stt_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.3/docs/resources/dialogflow_conversation_profile#stt_config DialogflowConversationProfile#stt_config}
	SttConfig *DialogflowConversationProfileSttConfig `field:"optional" json:"sttConfig" yaml:"sttConfig"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.3/docs/resources/dialogflow_conversation_profile#timeouts DialogflowConversationProfile#timeouts}
	Timeouts *DialogflowConversationProfileTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
	// The time zone of this conversational profile.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.3/docs/resources/dialogflow_conversation_profile#time_zone DialogflowConversationProfile#time_zone}
	TimeZone *string `field:"optional" json:"timeZone" yaml:"timeZone"`
	// tts_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.3/docs/resources/dialogflow_conversation_profile#tts_config DialogflowConversationProfile#tts_config}
	TtsConfig *DialogflowConversationProfileTtsConfig `field:"optional" json:"ttsConfig" yaml:"ttsConfig"`
}

