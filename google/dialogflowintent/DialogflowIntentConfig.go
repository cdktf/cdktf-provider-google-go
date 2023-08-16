package dialogflowintent

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DialogflowIntentConfig struct {
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
	// The name of this intent to be displayed on the console.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/dialogflow_intent#display_name DialogflowIntent#display_name}
	DisplayName *string `field:"required" json:"displayName" yaml:"displayName"`
	// The name of the action associated with the intent. Note: The action name must not contain whitespaces.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/dialogflow_intent#action DialogflowIntent#action}
	Action *string `field:"optional" json:"action" yaml:"action"`
	// The list of platforms for which the first responses will be copied from the messages in PLATFORM_UNSPECIFIED (i.e. default platform). Possible values: ["FACEBOOK", "SLACK", "TELEGRAM", "KIK", "SKYPE", "LINE", "VIBER", "ACTIONS_ON_GOOGLE", "GOOGLE_HANGOUTS"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/dialogflow_intent#default_response_platforms DialogflowIntent#default_response_platforms}
	DefaultResponsePlatforms *[]*string `field:"optional" json:"defaultResponsePlatforms" yaml:"defaultResponsePlatforms"`
	// The collection of event names that trigger the intent.
	//
	// If the collection of input contexts is not empty, all of
	// the contexts must be present in the active user session for an event to trigger this intent. See the
	// [events reference](https://cloud.google.com/dialogflow/docs/events-overview) for more details.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/dialogflow_intent#events DialogflowIntent#events}
	Events *[]*string `field:"optional" json:"events" yaml:"events"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/dialogflow_intent#id DialogflowIntent#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// The list of context names required for this intent to be triggered. Format: projects/<Project ID>/agent/sessions/-/contexts/<Context ID>.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/dialogflow_intent#input_context_names DialogflowIntent#input_context_names}
	InputContextNames *[]*string `field:"optional" json:"inputContextNames" yaml:"inputContextNames"`
	// Indicates whether this is a fallback intent.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/dialogflow_intent#is_fallback DialogflowIntent#is_fallback}
	IsFallback interface{} `field:"optional" json:"isFallback" yaml:"isFallback"`
	// Indicates whether Machine Learning is disabled for the intent.
	//
	// Note: If mlDisabled setting is set to true, then this intent is not taken into account during inference in ML
	// ONLY match mode. Also, auto-markup in the UI is turned off.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/dialogflow_intent#ml_disabled DialogflowIntent#ml_disabled}
	MlDisabled interface{} `field:"optional" json:"mlDisabled" yaml:"mlDisabled"`
	// The unique identifier of the parent intent in the chain of followup intents. Format: projects/<Project ID>/agent/intents/<Intent ID>.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/dialogflow_intent#parent_followup_intent_name DialogflowIntent#parent_followup_intent_name}
	ParentFollowupIntentName *string `field:"optional" json:"parentFollowupIntentName" yaml:"parentFollowupIntentName"`
	// The priority of this intent.
	//
	// Higher numbers represent higher priorities.
	// - If the supplied value is unspecified or 0, the service translates the value to 500,000, which corresponds
	// to the Normal priority in the console.
	// - If the supplied value is negative, the intent is ignored in runtime detect intent requests.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/dialogflow_intent#priority DialogflowIntent#priority}
	Priority *float64 `field:"optional" json:"priority" yaml:"priority"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/dialogflow_intent#project DialogflowIntent#project}.
	Project *string `field:"optional" json:"project" yaml:"project"`
	// Indicates whether to delete all contexts in the current session when this intent is matched.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/dialogflow_intent#reset_contexts DialogflowIntent#reset_contexts}
	ResetContexts interface{} `field:"optional" json:"resetContexts" yaml:"resetContexts"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/dialogflow_intent#timeouts DialogflowIntent#timeouts}
	Timeouts *DialogflowIntentTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
	// Indicates whether webhooks are enabled for the intent.
	//
	// WEBHOOK_STATE_ENABLED: Webhook is enabled in the agent and in the intent.
	// WEBHOOK_STATE_ENABLED_FOR_SLOT_FILLING: Webhook is enabled in the agent and in the intent. Also, each slot
	// filling prompt is forwarded to the webhook. Possible values: ["WEBHOOK_STATE_ENABLED", "WEBHOOK_STATE_ENABLED_FOR_SLOT_FILLING"]
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/dialogflow_intent#webhook_state DialogflowIntent#webhook_state}
	WebhookState *string `field:"optional" json:"webhookState" yaml:"webhookState"`
}

