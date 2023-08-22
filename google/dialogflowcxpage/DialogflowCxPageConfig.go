package dialogflowcxpage

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DialogflowCxPageConfig struct {
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
	// The human-readable name of the page, unique within the agent.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/dialogflow_cx_page#display_name DialogflowCxPage#display_name}
	DisplayName *string `field:"required" json:"displayName" yaml:"displayName"`
	// entry_fulfillment block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/dialogflow_cx_page#entry_fulfillment DialogflowCxPage#entry_fulfillment}
	EntryFulfillment *DialogflowCxPageEntryFulfillment `field:"optional" json:"entryFulfillment" yaml:"entryFulfillment"`
	// event_handlers block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/dialogflow_cx_page#event_handlers DialogflowCxPage#event_handlers}
	EventHandlers interface{} `field:"optional" json:"eventHandlers" yaml:"eventHandlers"`
	// form block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/dialogflow_cx_page#form DialogflowCxPage#form}
	Form *DialogflowCxPageForm `field:"optional" json:"form" yaml:"form"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/dialogflow_cx_page#id DialogflowCxPage#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// The language of the following fields in page:.
	//
	// Page.entry_fulfillment.messages
	// Page.entry_fulfillment.conditional_cases
	// Page.event_handlers.trigger_fulfillment.messages
	// Page.event_handlers.trigger_fulfillment.conditional_cases
	// Page.form.parameters.fill_behavior.initial_prompt_fulfillment.messages
	// Page.form.parameters.fill_behavior.initial_prompt_fulfillment.conditional_cases
	// Page.form.parameters.fill_behavior.reprompt_event_handlers.messages
	// Page.form.parameters.fill_behavior.reprompt_event_handlers.conditional_cases
	// Page.transition_routes.trigger_fulfillment.messages
	// Page.transition_routes.trigger_fulfillment.conditional_cases
	// If not specified, the agent's default language is used. Many languages are supported. Note: languages must be enabled in the agent before they can be used.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/dialogflow_cx_page#language_code DialogflowCxPage#language_code}
	LanguageCode *string `field:"optional" json:"languageCode" yaml:"languageCode"`
	// The flow to create a page for. Format: projects/<Project ID>/locations/<Location ID>/agents/<Agent ID>/flows/<Flow ID>.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/dialogflow_cx_page#parent DialogflowCxPage#parent}
	Parent *string `field:"optional" json:"parent" yaml:"parent"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/dialogflow_cx_page#timeouts DialogflowCxPage#timeouts}
	Timeouts *DialogflowCxPageTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
	// Ordered list of TransitionRouteGroups associated with the page.
	//
	// Transition route groups must be unique within a page.
	// If multiple transition routes within a page scope refer to the same intent, then the precedence order is: page's transition route -> page's transition route group -> flow's transition routes.
	// If multiple transition route groups within a page contain the same intent, then the first group in the ordered list takes precedence.
	// Format:projects/<Project ID>/locations/<Location ID>/agents/<Agent ID>/flows/<Flow ID>/transitionRouteGroups/<TransitionRouteGroup ID>.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/dialogflow_cx_page#transition_route_groups DialogflowCxPage#transition_route_groups}
	TransitionRouteGroups *[]*string `field:"optional" json:"transitionRouteGroups" yaml:"transitionRouteGroups"`
	// transition_routes block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/dialogflow_cx_page#transition_routes DialogflowCxPage#transition_routes}
	TransitionRoutes interface{} `field:"optional" json:"transitionRoutes" yaml:"transitionRoutes"`
}

