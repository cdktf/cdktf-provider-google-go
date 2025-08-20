// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package firestoredatabase

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type FirestoreDatabaseConfig struct {
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
	// The location of the database. Available locations are listed at https://cloud.google.com/firestore/docs/locations.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.0/docs/resources/firestore_database#location_id FirestoreDatabase#location_id}
	LocationId *string `field:"required" json:"locationId" yaml:"locationId"`
	// The ID to use for the database, which will become the final component of the database's resource name.
	//
	// This value should be 4-63
	// characters. Valid characters are /[a-z][0-9]-/ with first character
	// a letter and the last a letter or a number. Must not be
	// UUID-like /[0-9a-f]{8}(-[0-9a-f]{4}){3}-[0-9a-f]{12}/.
	// "(default)" database id is also valid.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.0/docs/resources/firestore_database#name FirestoreDatabase#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The type of the database. See https://cloud.google.com/datastore/docs/firestore-or-datastore for information about how to choose. Possible values: ["FIRESTORE_NATIVE", "DATASTORE_MODE"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.0/docs/resources/firestore_database#type FirestoreDatabase#type}
	Type *string `field:"required" json:"type" yaml:"type"`
	// The App Engine integration mode to use for this database. Possible values: ["ENABLED", "DISABLED"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.0/docs/resources/firestore_database#app_engine_integration_mode FirestoreDatabase#app_engine_integration_mode}
	AppEngineIntegrationMode *string `field:"optional" json:"appEngineIntegrationMode" yaml:"appEngineIntegrationMode"`
	// cmek_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.0/docs/resources/firestore_database#cmek_config FirestoreDatabase#cmek_config}
	CmekConfig *FirestoreDatabaseCmekConfig `field:"optional" json:"cmekConfig" yaml:"cmekConfig"`
	// The concurrency control mode to use for this database. Possible values: ["OPTIMISTIC", "PESSIMISTIC", "OPTIMISTIC_WITH_ENTITY_GROUPS"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.0/docs/resources/firestore_database#concurrency_mode FirestoreDatabase#concurrency_mode}
	ConcurrencyMode *string `field:"optional" json:"concurrencyMode" yaml:"concurrencyMode"`
	// The database edition. Possible values: ["STANDARD", "ENTERPRISE"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.0/docs/resources/firestore_database#database_edition FirestoreDatabase#database_edition}
	DatabaseEdition *string `field:"optional" json:"databaseEdition" yaml:"databaseEdition"`
	// State of delete protection for the database.
	//
	// When delete protection is enabled, this database cannot be deleted.
	// The default value is 'DELETE_PROTECTION_STATE_UNSPECIFIED', which is currently equivalent to 'DELETE_PROTECTION_DISABLED'.
	// **Note:** Additionally, to delete this database using 'terraform destroy', 'deletion_policy' must be set to 'DELETE'. Possible values: ["DELETE_PROTECTION_STATE_UNSPECIFIED", "DELETE_PROTECTION_ENABLED", "DELETE_PROTECTION_DISABLED"]
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.0/docs/resources/firestore_database#delete_protection_state FirestoreDatabase#delete_protection_state}
	DeleteProtectionState *string `field:"optional" json:"deleteProtectionState" yaml:"deleteProtectionState"`
	// Deletion behavior for this database.
	//
	// If the deletion policy is 'ABANDON', the database will be removed from Terraform state but not deleted from Google Cloud upon destruction.
	// If the deletion policy is 'DELETE', the database will both be removed from Terraform state and deleted from Google Cloud upon destruction.
	// The default value is 'ABANDON'.
	// See also 'delete_protection'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.0/docs/resources/firestore_database#deletion_policy FirestoreDatabase#deletion_policy}
	DeletionPolicy *string `field:"optional" json:"deletionPolicy" yaml:"deletionPolicy"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.0/docs/resources/firestore_database#id FirestoreDatabase#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Whether to enable the PITR feature on this database.
	//
	// If 'POINT_IN_TIME_RECOVERY_ENABLED' is selected, reads are supported on selected versions of the data from within the past 7 days.
	// versionRetentionPeriod and earliestVersionTime can be used to determine the supported versions. These include reads against any timestamp within the past hour
	// and reads against 1-minute snapshots beyond 1 hour and within 7 days.
	// If 'POINT_IN_TIME_RECOVERY_DISABLED' is selected, reads are supported on any version of the data from within the past 1 hour. Default value: "POINT_IN_TIME_RECOVERY_DISABLED" Possible values: ["POINT_IN_TIME_RECOVERY_ENABLED", "POINT_IN_TIME_RECOVERY_DISABLED"]
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.0/docs/resources/firestore_database#point_in_time_recovery_enablement FirestoreDatabase#point_in_time_recovery_enablement}
	PointInTimeRecoveryEnablement *string `field:"optional" json:"pointInTimeRecoveryEnablement" yaml:"pointInTimeRecoveryEnablement"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.0/docs/resources/firestore_database#project FirestoreDatabase#project}.
	Project *string `field:"optional" json:"project" yaml:"project"`
	// Input only.
	//
	// A map of resource manager tags. Resource manager tag keys
	// and values have the same definition as resource manager tags.
	// Keys must be in the format tagKeys/{tag_key_id}, and values are in the format tagValues/456.
	// The field is ignored when empty. The field is immutable and causes
	// resource replacement when mutated. To apply tags to an existing resource, see
	// the 'google_tags_tag_value' resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.0/docs/resources/firestore_database#tags FirestoreDatabase#tags}
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.0/docs/resources/firestore_database#timeouts FirestoreDatabase#timeouts}
	Timeouts *FirestoreDatabaseTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

