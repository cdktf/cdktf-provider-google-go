// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package bigquerytable

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type BigqueryTableConfig struct {
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
	// The dataset ID to create the table in. Changing this forces a new resource to be created.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/bigquery_table#dataset_id BigqueryTable#dataset_id}
	DatasetId *string `field:"required" json:"datasetId" yaml:"datasetId"`
	// A unique ID for the resource. Changing this forces a new resource to be created.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/bigquery_table#table_id BigqueryTable#table_id}
	TableId *string `field:"required" json:"tableId" yaml:"tableId"`
	// biglake_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/bigquery_table#biglake_configuration BigqueryTable#biglake_configuration}
	BiglakeConfiguration *BigqueryTableBiglakeConfiguration `field:"optional" json:"biglakeConfiguration" yaml:"biglakeConfiguration"`
	// Specifies column names to use for data clustering.
	//
	// Up to four top-level columns are allowed, and should be specified in descending priority order.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/bigquery_table#clustering BigqueryTable#clustering}
	Clustering *[]*string `field:"optional" json:"clustering" yaml:"clustering"`
	// Whether Terraform will be prevented from destroying the instance.
	//
	// When the field is set to true or unset in Terraform state, a terraform apply or terraform destroy that would delete the table will fail. When the field is set to false, deleting the table is allowed.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/bigquery_table#deletion_protection BigqueryTable#deletion_protection}
	DeletionProtection interface{} `field:"optional" json:"deletionProtection" yaml:"deletionProtection"`
	// The field description.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/bigquery_table#description BigqueryTable#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// encryption_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/bigquery_table#encryption_configuration BigqueryTable#encryption_configuration}
	EncryptionConfiguration *BigqueryTableEncryptionConfiguration `field:"optional" json:"encryptionConfiguration" yaml:"encryptionConfiguration"`
	// The time when this table expires, in milliseconds since the epoch.
	//
	// If not present, the table will persist indefinitely. Expired tables will be deleted and their storage reclaimed.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/bigquery_table#expiration_time BigqueryTable#expiration_time}
	ExpirationTime *float64 `field:"optional" json:"expirationTime" yaml:"expirationTime"`
	// external_catalog_table_options block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/bigquery_table#external_catalog_table_options BigqueryTable#external_catalog_table_options}
	ExternalCatalogTableOptions *BigqueryTableExternalCatalogTableOptions `field:"optional" json:"externalCatalogTableOptions" yaml:"externalCatalogTableOptions"`
	// external_data_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/bigquery_table#external_data_configuration BigqueryTable#external_data_configuration}
	ExternalDataConfiguration *BigqueryTableExternalDataConfiguration `field:"optional" json:"externalDataConfiguration" yaml:"externalDataConfiguration"`
	// A descriptive name for the table.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/bigquery_table#friendly_name BigqueryTable#friendly_name}
	FriendlyName *string `field:"optional" json:"friendlyName" yaml:"friendlyName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/bigquery_table#id BigqueryTable#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Whether Terraform will prevent implicitly added columns in schema from showing diff.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/bigquery_table#ignore_auto_generated_schema BigqueryTable#ignore_auto_generated_schema}
	IgnoreAutoGeneratedSchema interface{} `field:"optional" json:"ignoreAutoGeneratedSchema" yaml:"ignoreAutoGeneratedSchema"`
	// Mention which fields in schema are to be ignored.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/bigquery_table#ignore_schema_changes BigqueryTable#ignore_schema_changes}
	IgnoreSchemaChanges *[]*string `field:"optional" json:"ignoreSchemaChanges" yaml:"ignoreSchemaChanges"`
	// A mapping of labels to assign to the resource.
	//
	// **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
	// 				Please refer to the field 'effective_labels' for all of the labels present on the resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/bigquery_table#labels BigqueryTable#labels}
	Labels *map[string]*string `field:"optional" json:"labels" yaml:"labels"`
	// materialized_view block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/bigquery_table#materialized_view BigqueryTable#materialized_view}
	MaterializedView *BigqueryTableMaterializedView `field:"optional" json:"materializedView" yaml:"materializedView"`
	// The maximum staleness of data that could be returned when the table (or stale MV) is queried.
	//
	// Staleness encoded as a string encoding of [SQL IntervalValue type](https://cloud.google.com/bigquery/docs/reference/standard-sql/data-types#interval_type).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/bigquery_table#max_staleness BigqueryTable#max_staleness}
	MaxStaleness *string `field:"optional" json:"maxStaleness" yaml:"maxStaleness"`
	// The ID of the project in which the resource belongs.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/bigquery_table#project BigqueryTable#project}
	Project *string `field:"optional" json:"project" yaml:"project"`
	// range_partitioning block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/bigquery_table#range_partitioning BigqueryTable#range_partitioning}
	RangePartitioning *BigqueryTableRangePartitioning `field:"optional" json:"rangePartitioning" yaml:"rangePartitioning"`
	// If set to true, queries over this table require a partition filter that can be used for partition elimination to be specified.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/bigquery_table#require_partition_filter BigqueryTable#require_partition_filter}
	RequirePartitionFilter interface{} `field:"optional" json:"requirePartitionFilter" yaml:"requirePartitionFilter"`
	// The tags attached to this table.
	//
	// Tag keys are globally unique. Tag key is expected to be in the namespaced format, for example "123456789012/environment" where 123456789012 is the ID of the parent organization or project resource for this tag key. Tag value is expected to be the short name, for example "Production".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/bigquery_table#resource_tags BigqueryTable#resource_tags}
	ResourceTags *map[string]*string `field:"optional" json:"resourceTags" yaml:"resourceTags"`
	// A JSON schema for the table.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/bigquery_table#schema BigqueryTable#schema}
	Schema *string `field:"optional" json:"schema" yaml:"schema"`
	// schema_foreign_type_info block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/bigquery_table#schema_foreign_type_info BigqueryTable#schema_foreign_type_info}
	SchemaForeignTypeInfo *BigqueryTableSchemaForeignTypeInfo `field:"optional" json:"schemaForeignTypeInfo" yaml:"schemaForeignTypeInfo"`
	// table_constraints block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/bigquery_table#table_constraints BigqueryTable#table_constraints}
	TableConstraints *BigqueryTableTableConstraints `field:"optional" json:"tableConstraints" yaml:"tableConstraints"`
	// View sets the optional parameter "view": Specifies the view that determines which table information is returned.
	//
	// By default, basic table information and storage statistics (STORAGE_STATS) are returned. Possible values: TABLE_METADATA_VIEW_UNSPECIFIED, BASIC, STORAGE_STATS, FULL
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/bigquery_table#table_metadata_view BigqueryTable#table_metadata_view}
	TableMetadataView *string `field:"optional" json:"tableMetadataView" yaml:"tableMetadataView"`
	// table_replication_info block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/bigquery_table#table_replication_info BigqueryTable#table_replication_info}
	TableReplicationInfo *BigqueryTableTableReplicationInfo `field:"optional" json:"tableReplicationInfo" yaml:"tableReplicationInfo"`
	// time_partitioning block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/bigquery_table#time_partitioning BigqueryTable#time_partitioning}
	TimePartitioning *BigqueryTableTimePartitioning `field:"optional" json:"timePartitioning" yaml:"timePartitioning"`
	// view block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/bigquery_table#view BigqueryTable#view}
	View *BigqueryTableView `field:"optional" json:"view" yaml:"view"`
}

