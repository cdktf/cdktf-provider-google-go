package dialogflowcxagent

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DialogflowCxAgentConfig struct {
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
	// The default language of the agent as a language tag.
	//
	// [See Language Support](https://cloud.google.com/dialogflow/cx/docs/reference/language)
	// for a list of the currently supported language codes. This field cannot be updated after creation.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/dialogflow_cx_agent#default_language_code DialogflowCxAgent#default_language_code}
	DefaultLanguageCode *string `field:"required" json:"defaultLanguageCode" yaml:"defaultLanguageCode"`
	// The human-readable name of the agent, unique within the location.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/dialogflow_cx_agent#display_name DialogflowCxAgent#display_name}
	DisplayName *string `field:"required" json:"displayName" yaml:"displayName"`
	// The name of the location this agent is located in.
	//
	// ~> **Note:** The first time you are deploying an Agent in your project you must configure location settings.
	// This is a one time step but at the moment you can only [configure location settings](https://cloud.google.com/dialogflow/cx/docs/concept/region#location-settings) via the Dialogflow CX console.
	// Another options is to use global location so you don't need to manually configure location settings.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/dialogflow_cx_agent#location DialogflowCxAgent#location}
	Location *string `field:"required" json:"location" yaml:"location"`
	// The time zone of this agent from the [time zone database](https://www.iana.org/time-zones), e.g., America/New_York, Europe/Paris.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/dialogflow_cx_agent#time_zone DialogflowCxAgent#time_zone}
	TimeZone *string `field:"required" json:"timeZone" yaml:"timeZone"`
	// The URI of the agent's avatar.
	//
	// Avatars are used throughout the Dialogflow console and in the self-hosted Web Demo integration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/dialogflow_cx_agent#avatar_uri DialogflowCxAgent#avatar_uri}
	AvatarUri *string `field:"optional" json:"avatarUri" yaml:"avatarUri"`
	// The description of this agent. The maximum length is 500 characters. If exceeded, the request is rejected.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/dialogflow_cx_agent#description DialogflowCxAgent#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Indicates if automatic spell correction is enabled in detect intent requests.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/dialogflow_cx_agent#enable_spell_correction DialogflowCxAgent#enable_spell_correction}
	EnableSpellCorrection interface{} `field:"optional" json:"enableSpellCorrection" yaml:"enableSpellCorrection"`
	// Determines whether this agent should log conversation queries.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/dialogflow_cx_agent#enable_stackdriver_logging DialogflowCxAgent#enable_stackdriver_logging}
	EnableStackdriverLogging interface{} `field:"optional" json:"enableStackdriverLogging" yaml:"enableStackdriverLogging"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/dialogflow_cx_agent#id DialogflowCxAgent#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/dialogflow_cx_agent#project DialogflowCxAgent#project}.
	Project *string `field:"optional" json:"project" yaml:"project"`
	// Name of the SecuritySettings reference for the agent. Format: projects/<Project ID>/locations/<Location ID>/securitySettings/<Security Settings ID>.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/dialogflow_cx_agent#security_settings DialogflowCxAgent#security_settings}
	SecuritySettings *string `field:"optional" json:"securitySettings" yaml:"securitySettings"`
	// speech_to_text_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/dialogflow_cx_agent#speech_to_text_settings DialogflowCxAgent#speech_to_text_settings}
	SpeechToTextSettings *DialogflowCxAgentSpeechToTextSettings `field:"optional" json:"speechToTextSettings" yaml:"speechToTextSettings"`
	// The list of all languages supported by this agent (except for the default_language_code).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/dialogflow_cx_agent#supported_language_codes DialogflowCxAgent#supported_language_codes}
	SupportedLanguageCodes *[]*string `field:"optional" json:"supportedLanguageCodes" yaml:"supportedLanguageCodes"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/dialogflow_cx_agent#timeouts DialogflowCxAgent#timeouts}
	Timeouts *DialogflowCxAgentTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

