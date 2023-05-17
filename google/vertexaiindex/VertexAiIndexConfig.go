package vertexaiindex

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type VertexAiIndexConfig struct {
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
	// The display name of the Index.
	//
	// The name can be up to 128 characters long and can consist of any UTF-8 characters.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/vertex_ai_index#display_name VertexAiIndex#display_name}
	DisplayName *string `field:"required" json:"displayName" yaml:"displayName"`
	// The description of the Index.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/vertex_ai_index#description VertexAiIndex#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/vertex_ai_index#id VertexAiIndex#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// The update method to use with this Index.
	//
	// The value must be the followings. If not set, BATCH_UPDATE will be used by default.
	// BATCH_UPDATE: user can call indexes.patch with files on Cloud Storage of datapoints to update.
	// STREAM_UPDATE: user can call indexes.upsertDatapoints/DeleteDatapoints to update the Index and the updates will be applied in corresponding DeployedIndexes in nearly real-time.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/vertex_ai_index#index_update_method VertexAiIndex#index_update_method}
	IndexUpdateMethod *string `field:"optional" json:"indexUpdateMethod" yaml:"indexUpdateMethod"`
	// The labels with user-defined metadata to organize your Indexes.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/vertex_ai_index#labels VertexAiIndex#labels}
	Labels *map[string]*string `field:"optional" json:"labels" yaml:"labels"`
	// metadata block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/vertex_ai_index#metadata VertexAiIndex#metadata}
	Metadata *VertexAiIndexMetadata `field:"optional" json:"metadata" yaml:"metadata"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/vertex_ai_index#project VertexAiIndex#project}.
	Project *string `field:"optional" json:"project" yaml:"project"`
	// The region of the index. eg us-central1.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/vertex_ai_index#region VertexAiIndex#region}
	Region *string `field:"optional" json:"region" yaml:"region"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/vertex_ai_index#timeouts VertexAiIndex#timeouts}
	Timeouts *VertexAiIndexTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

