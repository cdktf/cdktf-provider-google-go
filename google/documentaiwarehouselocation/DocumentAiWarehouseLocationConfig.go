package documentaiwarehouselocation

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DocumentAiWarehouseLocationConfig struct {
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
	// The access control mode for accessing the customer data. Possible values: ["ACL_MODE_DOCUMENT_LEVEL_ACCESS_CONTROL_GCI", "ACL_MODE_DOCUMENT_LEVEL_ACCESS_CONTROL_BYOID", "ACL_MODE_UNIVERSAL_ACCESS"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/document_ai_warehouse_location#access_control_mode DocumentAiWarehouseLocation#access_control_mode}
	AccessControlMode *string `field:"required" json:"accessControlMode" yaml:"accessControlMode"`
	// The type of database used to store customer data. Possible values: ["DB_INFRA_SPANNER", "DB_CLOUD_SQL_POSTGRES"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/document_ai_warehouse_location#database_type DocumentAiWarehouseLocation#database_type}
	DatabaseType *string `field:"required" json:"databaseType" yaml:"databaseType"`
	// The location in which the instance is to be provisioned. It takes the form projects/{projectNumber}/locations/{location}.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/document_ai_warehouse_location#location DocumentAiWarehouseLocation#location}
	Location *string `field:"required" json:"location" yaml:"location"`
	// The unique identifier of the project.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/document_ai_warehouse_location#project_number DocumentAiWarehouseLocation#project_number}
	ProjectNumber *string `field:"required" json:"projectNumber" yaml:"projectNumber"`
	// The default role for the person who create a document. Possible values: ["DOCUMENT_ADMIN", "DOCUMENT_EDITOR", "DOCUMENT_VIEWER"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/document_ai_warehouse_location#document_creator_default_role DocumentAiWarehouseLocation#document_creator_default_role}
	DocumentCreatorDefaultRole *string `field:"optional" json:"documentCreatorDefaultRole" yaml:"documentCreatorDefaultRole"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/document_ai_warehouse_location#id DocumentAiWarehouseLocation#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// The KMS key used for CMEK encryption.
	//
	// It is required that
	// the kms key is in the same region as the endpoint. The
	// same key will be used for all provisioned resources, if
	// encryption is available. If the kmsKey is left empty, no
	// encryption will be enforced.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/document_ai_warehouse_location#kms_key DocumentAiWarehouseLocation#kms_key}
	KmsKey *string `field:"optional" json:"kmsKey" yaml:"kmsKey"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/document_ai_warehouse_location#timeouts DocumentAiWarehouseLocation#timeouts}
	Timeouts *DocumentAiWarehouseLocationTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

