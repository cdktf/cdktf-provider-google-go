package datacatalogentry

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataCatalogEntryConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count *float64 `field:"optional" json:"count" yaml:"count"`
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
	// The name of the entry group this entry is in.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/data_catalog_entry#entry_group DataCatalogEntry#entry_group}
	EntryGroup *string `field:"required" json:"entryGroup" yaml:"entryGroup"`
	// The id of the entry to create.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/data_catalog_entry#entry_id DataCatalogEntry#entry_id}
	EntryId *string `field:"required" json:"entryId" yaml:"entryId"`
	// Entry description, which can consist of several sentences or paragraphs that describe entry contents.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/data_catalog_entry#description DataCatalogEntry#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Display information such as title and description.
	//
	// A short name to identify the entry,
	// for example, "Analytics Data - Jan 2011".
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/data_catalog_entry#display_name DataCatalogEntry#display_name}
	DisplayName *string `field:"optional" json:"displayName" yaml:"displayName"`
	// gcs_fileset_spec block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/data_catalog_entry#gcs_fileset_spec DataCatalogEntry#gcs_fileset_spec}
	GcsFilesetSpec *DataCatalogEntryGcsFilesetSpec `field:"optional" json:"gcsFilesetSpec" yaml:"gcsFilesetSpec"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/data_catalog_entry#id DataCatalogEntry#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// The resource this metadata entry refers to.
	//
	// For Google Cloud Platform resources, linkedResource is the full name of the resource.
	// For example, the linkedResource for a table resource from BigQuery is:
	// //bigquery.googleapis.com/projects/projectId/datasets/datasetId/tables/tableId
	// Output only when Entry is of type in the EntryType enum. For entries with userSpecifiedType,
	// this field is optional and defaults to an empty string.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/data_catalog_entry#linked_resource DataCatalogEntry#linked_resource}
	LinkedResource *string `field:"optional" json:"linkedResource" yaml:"linkedResource"`
	// Schema of the entry (e.g. BigQuery, GoogleSQL, Avro schema), as a json string. An entry might not have any schema attached to it. See https://cloud.google.com/data-catalog/docs/reference/rest/v1/projects.locations.entryGroups.entries#schema for what fields this schema can contain.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/data_catalog_entry#schema DataCatalogEntry#schema}
	Schema *string `field:"optional" json:"schema" yaml:"schema"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/data_catalog_entry#timeouts DataCatalogEntry#timeouts}
	Timeouts *DataCatalogEntryTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
	// The type of the entry.
	//
	// Only used for Entries with types in the EntryType enum.
	// Currently, only FILESET enum value is allowed. All other entries created through Data Catalog must use userSpecifiedType. Possible values: ["FILESET"]
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/data_catalog_entry#type DataCatalogEntry#type}
	Type *string `field:"optional" json:"type" yaml:"type"`
	// This field indicates the entry's source system that Data Catalog does not integrate with.
	//
	// userSpecifiedSystem strings must begin with a letter or underscore and can only contain letters, numbers,
	// and underscores; are case insensitive; must be at least 1 character and at most 64 characters long.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/data_catalog_entry#user_specified_system DataCatalogEntry#user_specified_system}
	UserSpecifiedSystem *string `field:"optional" json:"userSpecifiedSystem" yaml:"userSpecifiedSystem"`
	// Entry type if it does not fit any of the input-allowed values listed in EntryType enum above.
	//
	// When creating an entry, users should check the enum values first, if nothing matches the entry
	// to be created, then provide a custom value, for example "my_special_type".
	// userSpecifiedType strings must begin with a letter or underscore and can only contain letters,
	// numbers, and underscores; are case insensitive; must be at least 1 character and at most 64 characters long.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/data_catalog_entry#user_specified_type DataCatalogEntry#user_specified_type}
	UserSpecifiedType *string `field:"optional" json:"userSpecifiedType" yaml:"userSpecifiedType"`
}

