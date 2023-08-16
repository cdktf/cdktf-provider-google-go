package documentaiwarehousedocumentschema

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DocumentAiWarehouseDocumentSchemaConfig struct {
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
	// Name of the schema given by the user.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/document_ai_warehouse_document_schema#display_name DocumentAiWarehouseDocumentSchema#display_name}
	DisplayName *string `field:"required" json:"displayName" yaml:"displayName"`
	// The location of the resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/document_ai_warehouse_document_schema#location DocumentAiWarehouseDocumentSchema#location}
	Location *string `field:"required" json:"location" yaml:"location"`
	// The unique identifier of the project.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/document_ai_warehouse_document_schema#project_number DocumentAiWarehouseDocumentSchema#project_number}
	ProjectNumber *string `field:"required" json:"projectNumber" yaml:"projectNumber"`
	// property_definitions block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/document_ai_warehouse_document_schema#property_definitions DocumentAiWarehouseDocumentSchema#property_definitions}
	PropertyDefinitions interface{} `field:"required" json:"propertyDefinitions" yaml:"propertyDefinitions"`
	// Tells whether the document is a folder or a typical document.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/document_ai_warehouse_document_schema#document_is_folder DocumentAiWarehouseDocumentSchema#document_is_folder}
	DocumentIsFolder interface{} `field:"optional" json:"documentIsFolder" yaml:"documentIsFolder"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/document_ai_warehouse_document_schema#id DocumentAiWarehouseDocumentSchema#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/document_ai_warehouse_document_schema#timeouts DocumentAiWarehouseDocumentSchema#timeouts}
	Timeouts *DocumentAiWarehouseDocumentSchemaTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

