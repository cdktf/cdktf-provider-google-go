package dialogflowagent

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DialogflowAgentConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count *float64 `field:"optional" json:"count" yaml:"count"`
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
	// [See Language Support](https://cloud.google.com/dialogflow/docs/reference/language)
	// for a list of the currently supported language codes. This field cannot be updated after creation.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/dialogflow_agent#default_language_code DialogflowAgent#default_language_code}
	DefaultLanguageCode *string `field:"required" json:"defaultLanguageCode" yaml:"defaultLanguageCode"`
	// The name of this agent.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/dialogflow_agent#display_name DialogflowAgent#display_name}
	DisplayName *string `field:"required" json:"displayName" yaml:"displayName"`
	// The time zone of this agent from the [time zone database](https://www.iana.org/time-zones), e.g., America/New_York, Europe/Paris.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/dialogflow_agent#time_zone DialogflowAgent#time_zone}
	TimeZone *string `field:"required" json:"timeZone" yaml:"timeZone"`
	// API version displayed in Dialogflow console.
	//
	// If not specified, V2 API is assumed. Clients are free to query
	// different service endpoints for different API versions. However, bots connectors and webhook calls will follow
	// the specified API version.
	// API_VERSION_V1: Legacy V1 API.
	// API_VERSION_V2: V2 API.
	// API_VERSION_V2_BETA_1: V2beta1 API. Possible values: ["API_VERSION_V1", "API_VERSION_V2", "API_VERSION_V2_BETA_1"]
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/dialogflow_agent#api_version DialogflowAgent#api_version}
	ApiVersion *string `field:"optional" json:"apiVersion" yaml:"apiVersion"`
	// The URI of the agent's avatar, which are used throughout the Dialogflow console.
	//
	// When an image URL is entered
	// into this field, the Dialogflow will save the image in the backend. The address of the backend image returned
	// from the API will be shown in the [avatarUriBackend] field.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/dialogflow_agent#avatar_uri DialogflowAgent#avatar_uri}
	AvatarUri *string `field:"optional" json:"avatarUri" yaml:"avatarUri"`
	// To filter out false positive results and still get variety in matched natural language inputs for your agent, you can tune the machine learning classification threshold.
	//
	// If the returned score value is less than the threshold
	// value, then a fallback intent will be triggered or, if there are no fallback intents defined, no intent will be
	// triggered. The score values range from 0.0 (completely uncertain) to 1.0 (completely certain). If set to 0.0, the
	// default of 0.3 is used.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/dialogflow_agent#classification_threshold DialogflowAgent#classification_threshold}
	ClassificationThreshold *float64 `field:"optional" json:"classificationThreshold" yaml:"classificationThreshold"`
	// The description of this agent. The maximum length is 500 characters. If exceeded, the request is rejected.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/dialogflow_agent#description DialogflowAgent#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Determines whether this agent should log conversation queries.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/dialogflow_agent#enable_logging DialogflowAgent#enable_logging}
	EnableLogging interface{} `field:"optional" json:"enableLogging" yaml:"enableLogging"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/dialogflow_agent#id DialogflowAgent#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Determines how intents are detected from user queries.
	//
	// MATCH_MODE_HYBRID: Best for agents with a small number of examples in intents and/or wide use of templates
	// syntax and composite entities.
	// MATCH_MODE_ML_ONLY: Can be used for agents with a large number of examples in intents, especially the ones
	// using @sys.any or very large developer entities. Possible values: ["MATCH_MODE_HYBRID", "MATCH_MODE_ML_ONLY"]
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/dialogflow_agent#match_mode DialogflowAgent#match_mode}
	MatchMode *string `field:"optional" json:"matchMode" yaml:"matchMode"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/dialogflow_agent#project DialogflowAgent#project}.
	Project *string `field:"optional" json:"project" yaml:"project"`
	// The list of all languages supported by this agent (except for the defaultLanguageCode).
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/dialogflow_agent#supported_language_codes DialogflowAgent#supported_language_codes}
	SupportedLanguageCodes *[]*string `field:"optional" json:"supportedLanguageCodes" yaml:"supportedLanguageCodes"`
	// The agent tier.
	//
	// If not specified, TIER_STANDARD is assumed.
	// TIER_STANDARD: Standard tier.
	// TIER_ENTERPRISE: Enterprise tier (Essentials).
	// TIER_ENTERPRISE_PLUS: Enterprise tier (Plus).
	// NOTE: Due to consistency issues, the provider will not read this field from the API. Drift is possible between
	// the Terraform state and Dialogflow if the agent tier is changed outside of Terraform. Possible values: ["TIER_STANDARD", "TIER_ENTERPRISE", "TIER_ENTERPRISE_PLUS"]
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/dialogflow_agent#tier DialogflowAgent#tier}
	Tier *string `field:"optional" json:"tier" yaml:"tier"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/dialogflow_agent#timeouts DialogflowAgent#timeouts}
	Timeouts *DialogflowAgentTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

