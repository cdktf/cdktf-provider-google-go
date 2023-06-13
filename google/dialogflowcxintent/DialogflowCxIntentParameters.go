package dialogflowcxintent


type DialogflowCxIntentParameters struct {
	// The entity type of the parameter.
	//
	// Format: projects/-/locations/-/agents/-/entityTypes/<System Entity Type ID> for system entity types (for example, projects/-/locations/-/agents/-/entityTypes/sys.date), or projects/<Project ID>/locations/<Location ID>/agents/<Agent ID>/entityTypes/<Entity Type ID> for developer entity types.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/resources/dialogflow_cx_intent#entity_type DialogflowCxIntent#entity_type}
	EntityType *string `field:"required" json:"entityType" yaml:"entityType"`
	// The unique identifier of the parameter. This field is used by training phrases to annotate their parts.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/resources/dialogflow_cx_intent#id DialogflowCxIntent#id}
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"required" json:"id" yaml:"id"`
	// Indicates whether the parameter represents a list of values.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/resources/dialogflow_cx_intent#is_list DialogflowCxIntent#is_list}
	IsList interface{} `field:"optional" json:"isList" yaml:"isList"`
	// Indicates whether the parameter content should be redacted in log.
	//
	// If redaction is enabled, the parameter content will be replaced by parameter name during logging.
	// Note: the parameter content is subject to redaction if either parameter level redaction or entity type level redaction is enabled.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/resources/dialogflow_cx_intent#redact DialogflowCxIntent#redact}
	Redact interface{} `field:"optional" json:"redact" yaml:"redact"`
}

