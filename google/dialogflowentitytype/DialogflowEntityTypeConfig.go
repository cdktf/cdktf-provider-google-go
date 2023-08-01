package dialogflowentitytype

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DialogflowEntityTypeConfig struct {
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
	// The name of this entity type to be displayed on the console.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/dialogflow_entity_type#display_name DialogflowEntityType#display_name}
	DisplayName *string `field:"required" json:"displayName" yaml:"displayName"`
	// Indicates the kind of entity type.
	//
	// KIND_MAP: Map entity types allow mapping of a group of synonyms to a reference value.
	// KIND_LIST: List entity types contain a set of entries that do not map to reference values. However, list entity
	// types can contain references to other entity types (with or without aliases).
	// KIND_REGEXP: Regexp entity types allow to specify regular expressions in entries values. Possible values: ["KIND_MAP", "KIND_LIST", "KIND_REGEXP"]
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/dialogflow_entity_type#kind DialogflowEntityType#kind}
	Kind *string `field:"required" json:"kind" yaml:"kind"`
	// Enables fuzzy entity extraction during classification.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/dialogflow_entity_type#enable_fuzzy_extraction DialogflowEntityType#enable_fuzzy_extraction}
	EnableFuzzyExtraction interface{} `field:"optional" json:"enableFuzzyExtraction" yaml:"enableFuzzyExtraction"`
	// entities block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/dialogflow_entity_type#entities DialogflowEntityType#entities}
	Entities interface{} `field:"optional" json:"entities" yaml:"entities"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/dialogflow_entity_type#id DialogflowEntityType#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/dialogflow_entity_type#project DialogflowEntityType#project}.
	Project *string `field:"optional" json:"project" yaml:"project"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/dialogflow_entity_type#timeouts DialogflowEntityType#timeouts}
	Timeouts *DialogflowEntityTypeTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

