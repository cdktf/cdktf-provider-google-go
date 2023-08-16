package vertexaifeaturestoreentitytypefeature

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type VertexAiFeaturestoreEntitytypeFeatureConfig struct {
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
	// The name of the Featurestore to use, in the format projects/{project}/locations/{location}/featurestores/{featurestore}/entityTypes/{entitytype}.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/vertex_ai_featurestore_entitytype_feature#entitytype VertexAiFeaturestoreEntitytypeFeature#entitytype}
	Entitytype *string `field:"required" json:"entitytype" yaml:"entitytype"`
	// Type of Feature value. Immutable. https://cloud.google.com/vertex-ai/docs/reference/rest/v1/projects.locations.featurestores.entityTypes.features#ValueType.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/vertex_ai_featurestore_entitytype_feature#value_type VertexAiFeaturestoreEntitytypeFeature#value_type}
	ValueType *string `field:"required" json:"valueType" yaml:"valueType"`
	// Description of the feature.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/vertex_ai_featurestore_entitytype_feature#description VertexAiFeaturestoreEntitytypeFeature#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/vertex_ai_featurestore_entitytype_feature#id VertexAiFeaturestoreEntitytypeFeature#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// A set of key/value label pairs to assign to the feature.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/vertex_ai_featurestore_entitytype_feature#labels VertexAiFeaturestoreEntitytypeFeature#labels}
	Labels *map[string]*string `field:"optional" json:"labels" yaml:"labels"`
	// The name of the feature.
	//
	// The feature can be up to 64 characters long and can consist only of ASCII Latin letters A-Z and a-z, underscore(_), and ASCII digits 0-9 starting with a letter. The value will be unique given an entity type.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/vertex_ai_featurestore_entitytype_feature#name VertexAiFeaturestoreEntitytypeFeature#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/vertex_ai_featurestore_entitytype_feature#timeouts VertexAiFeaturestoreEntitytypeFeature#timeouts}
	Timeouts *VertexAiFeaturestoreEntitytypeFeatureTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

