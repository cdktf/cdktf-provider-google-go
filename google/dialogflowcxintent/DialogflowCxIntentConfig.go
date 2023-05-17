package dialogflowcxintent

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DialogflowCxIntentConfig struct {
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
	// The human-readable name of the intent, unique within the agent.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/dialogflow_cx_intent#display_name DialogflowCxIntent#display_name}
	DisplayName *string `field:"required" json:"displayName" yaml:"displayName"`
	// Human readable description for better understanding an intent like its scope, content, result etc. Maximum character limit: 140 characters.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/dialogflow_cx_intent#description DialogflowCxIntent#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/dialogflow_cx_intent#id DialogflowCxIntent#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Indicates whether this is a fallback intent.
	//
	// Currently only default fallback intent is allowed in the agent, which is added upon agent creation.
	// Adding training phrases to fallback intent is useful in the case of requests that are mistakenly matched, since training phrases assigned to fallback intents act as negative examples that triggers no-match event.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/dialogflow_cx_intent#is_fallback DialogflowCxIntent#is_fallback}
	IsFallback interface{} `field:"optional" json:"isFallback" yaml:"isFallback"`
	// The key/value metadata to label an intent.
	//
	// Labels can contain lowercase letters, digits and the symbols '-' and '_'. International characters are allowed, including letters from unicase alphabets. Keys must start with a letter. Keys and values can be no longer than 63 characters and no more than 128 bytes.
	// Prefix "sys-" is reserved for Dialogflow defined labels. Currently allowed Dialogflow defined labels include: * sys-head * sys-contextual The above labels do not require value. "sys-head" means the intent is a head intent. "sys.contextual" means the intent is a contextual intent.
	// An object containing a list of "key": value pairs. Example: { "name": "wrench", "mass": "1.3kg", "count": "3" }.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/dialogflow_cx_intent#labels DialogflowCxIntent#labels}
	Labels *map[string]*string `field:"optional" json:"labels" yaml:"labels"`
	// The language of the following fields in intent: Intent.training_phrases.parts.text If not specified, the agent's default language is used. Many languages are supported. Note: languages must be enabled in the agent before they can be used.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/dialogflow_cx_intent#language_code DialogflowCxIntent#language_code}
	LanguageCode *string `field:"optional" json:"languageCode" yaml:"languageCode"`
	// parameters block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/dialogflow_cx_intent#parameters DialogflowCxIntent#parameters}
	Parameters interface{} `field:"optional" json:"parameters" yaml:"parameters"`
	// The agent to create an intent for. Format: projects/<Project ID>/locations/<Location ID>/agents/<Agent ID>.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/dialogflow_cx_intent#parent DialogflowCxIntent#parent}
	Parent *string `field:"optional" json:"parent" yaml:"parent"`
	// The priority of this intent.
	//
	// Higher numbers represent higher priorities.
	// If the supplied value is unspecified or 0, the service translates the value to 500,000, which corresponds to the Normal priority in the console.
	// If the supplied value is negative, the intent is ignored in runtime detect intent requests.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/dialogflow_cx_intent#priority DialogflowCxIntent#priority}
	Priority *float64 `field:"optional" json:"priority" yaml:"priority"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/dialogflow_cx_intent#timeouts DialogflowCxIntent#timeouts}
	Timeouts *DialogflowCxIntentTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
	// training_phrases block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/dialogflow_cx_intent#training_phrases DialogflowCxIntent#training_phrases}
	TrainingPhrases interface{} `field:"optional" json:"trainingPhrases" yaml:"trainingPhrases"`
}

