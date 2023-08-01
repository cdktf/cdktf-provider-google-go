package dialogflowcxentitytype

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DialogflowCxEntityTypeConfig struct {
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
	// The human-readable name of the entity type, unique within the agent.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/dialogflow_cx_entity_type#display_name DialogflowCxEntityType#display_name}
	DisplayName *string `field:"required" json:"displayName" yaml:"displayName"`
	// entities block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/dialogflow_cx_entity_type#entities DialogflowCxEntityType#entities}
	Entities interface{} `field:"required" json:"entities" yaml:"entities"`
	// Indicates whether the entity type can be automatically expanded.
	//
	// KIND_MAP: Map entity types allow mapping of a group of synonyms to a canonical value.
	// KIND_LIST: List entity types contain a set of entries that do not map to canonical values. However, list entity types can contain references to other entity types (with or without aliases).
	// KIND_REGEXP: Regexp entity types allow to specify regular expressions in entries values. Possible values: ["KIND_MAP", "KIND_LIST", "KIND_REGEXP"]
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/dialogflow_cx_entity_type#kind DialogflowCxEntityType#kind}
	Kind *string `field:"required" json:"kind" yaml:"kind"`
	// Represents kinds of entities.
	//
	// AUTO_EXPANSION_MODE_UNSPECIFIED: Auto expansion disabled for the entity.
	// AUTO_EXPANSION_MODE_DEFAULT: Allows an agent to recognize values that have not been explicitly listed in the entity. Possible values: ["AUTO_EXPANSION_MODE_DEFAULT", "AUTO_EXPANSION_MODE_UNSPECIFIED"]
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/dialogflow_cx_entity_type#auto_expansion_mode DialogflowCxEntityType#auto_expansion_mode}
	AutoExpansionMode *string `field:"optional" json:"autoExpansionMode" yaml:"autoExpansionMode"`
	// Enables fuzzy entity extraction during classification.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/dialogflow_cx_entity_type#enable_fuzzy_extraction DialogflowCxEntityType#enable_fuzzy_extraction}
	EnableFuzzyExtraction interface{} `field:"optional" json:"enableFuzzyExtraction" yaml:"enableFuzzyExtraction"`
	// excluded_phrases block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/dialogflow_cx_entity_type#excluded_phrases DialogflowCxEntityType#excluded_phrases}
	ExcludedPhrases interface{} `field:"optional" json:"excludedPhrases" yaml:"excludedPhrases"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/dialogflow_cx_entity_type#id DialogflowCxEntityType#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// The language of the following fields in entityType: EntityType.entities.value EntityType.entities.synonyms EntityType.excluded_phrases.value If not specified, the agent's default language is used. Many languages are supported. Note: languages must be enabled in the agent before they can be used.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/dialogflow_cx_entity_type#language_code DialogflowCxEntityType#language_code}
	LanguageCode *string `field:"optional" json:"languageCode" yaml:"languageCode"`
	// The agent to create a entity type for. Format: projects/<Project ID>/locations/<Location ID>/agents/<Agent ID>.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/dialogflow_cx_entity_type#parent DialogflowCxEntityType#parent}
	Parent *string `field:"optional" json:"parent" yaml:"parent"`
	// Indicates whether parameters of the entity type should be redacted in log.
	//
	// If redaction is enabled, page parameters and intent parameters referring to the entity type will be replaced by parameter name when logging.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/dialogflow_cx_entity_type#redact DialogflowCxEntityType#redact}
	Redact interface{} `field:"optional" json:"redact" yaml:"redact"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/dialogflow_cx_entity_type#timeouts DialogflowCxEntityType#timeouts}
	Timeouts *DialogflowCxEntityTypeTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

