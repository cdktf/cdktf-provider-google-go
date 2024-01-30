// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package alloydbcluster

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type AlloydbClusterConfig struct {
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
	// The ID of the alloydb cluster.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.14.0/docs/resources/alloydb_cluster#cluster_id AlloydbCluster#cluster_id}
	ClusterId *string `field:"required" json:"clusterId" yaml:"clusterId"`
	// The location where the alloydb cluster should reside.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.14.0/docs/resources/alloydb_cluster#location AlloydbCluster#location}
	Location *string `field:"required" json:"location" yaml:"location"`
	// Annotations to allow client tools to store small amount of arbitrary data.
	//
	// This is distinct from labels. https://google.aip.dev/128
	// An object containing a list of "key": value pairs. Example: { "name": "wrench", "mass": "1.3kg", "count": "3" }.
	//
	//
	// **Note**: This field is non-authoritative, and will only manage the annotations present in your configuration.
	// Please refer to the field 'effective_annotations' for all of the annotations present on the resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.14.0/docs/resources/alloydb_cluster#annotations AlloydbCluster#annotations}
	Annotations *map[string]*string `field:"optional" json:"annotations" yaml:"annotations"`
	// automated_backup_policy block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.14.0/docs/resources/alloydb_cluster#automated_backup_policy AlloydbCluster#automated_backup_policy}
	AutomatedBackupPolicy *AlloydbClusterAutomatedBackupPolicy `field:"optional" json:"automatedBackupPolicy" yaml:"automatedBackupPolicy"`
	// The type of cluster. If not set, defaults to PRIMARY. Default value: "PRIMARY" Possible values: ["PRIMARY", "SECONDARY"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.14.0/docs/resources/alloydb_cluster#cluster_type AlloydbCluster#cluster_type}
	ClusterType *string `field:"optional" json:"clusterType" yaml:"clusterType"`
	// continuous_backup_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.14.0/docs/resources/alloydb_cluster#continuous_backup_config AlloydbCluster#continuous_backup_config}
	ContinuousBackupConfig *AlloydbClusterContinuousBackupConfig `field:"optional" json:"continuousBackupConfig" yaml:"continuousBackupConfig"`
	// The database engine major version.
	//
	// This is an optional field and it's populated at the Cluster creation time. This field cannot be changed after cluster creation.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.14.0/docs/resources/alloydb_cluster#database_version AlloydbCluster#database_version}
	DatabaseVersion *string `field:"optional" json:"databaseVersion" yaml:"databaseVersion"`
	// Policy to determine if the cluster should be deleted forcefully.
	//
	// Deleting a cluster forcefully, deletes the cluster and all its associated instances within the cluster.
	// Deleting a Secondary cluster with a secondary instance REQUIRES setting deletion_policy = "FORCE" otherwise an error is returned. This is needed as there is no support to delete just the secondary instance, and the only way to delete secondary instance is to delete the associated secondary cluster forcefully which also deletes the secondary instance.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.14.0/docs/resources/alloydb_cluster#deletion_policy AlloydbCluster#deletion_policy}
	DeletionPolicy *string `field:"optional" json:"deletionPolicy" yaml:"deletionPolicy"`
	// User-settable and human-readable display name for the Cluster.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.14.0/docs/resources/alloydb_cluster#display_name AlloydbCluster#display_name}
	DisplayName *string `field:"optional" json:"displayName" yaml:"displayName"`
	// encryption_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.14.0/docs/resources/alloydb_cluster#encryption_config AlloydbCluster#encryption_config}
	EncryptionConfig *AlloydbClusterEncryptionConfig `field:"optional" json:"encryptionConfig" yaml:"encryptionConfig"`
	// For Resource freshness validation (https://google.aip.dev/154).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.14.0/docs/resources/alloydb_cluster#etag AlloydbCluster#etag}
	Etag *string `field:"optional" json:"etag" yaml:"etag"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.14.0/docs/resources/alloydb_cluster#id AlloydbCluster#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// initial_user block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.14.0/docs/resources/alloydb_cluster#initial_user AlloydbCluster#initial_user}
	InitialUser *AlloydbClusterInitialUser `field:"optional" json:"initialUser" yaml:"initialUser"`
	// User-defined labels for the alloydb cluster.
	//
	// **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
	// Please refer to the field 'effective_labels' for all of the labels present on the resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.14.0/docs/resources/alloydb_cluster#labels AlloydbCluster#labels}
	Labels *map[string]*string `field:"optional" json:"labels" yaml:"labels"`
	// The relative resource name of the VPC network on which the instance can be accessed.
	//
	// It is specified in the following form:
	//
	// "projects/{projectNumber}/global/networks/{network_id}".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.14.0/docs/resources/alloydb_cluster#network AlloydbCluster#network}
	Network *string `field:"optional" json:"network" yaml:"network"`
	// network_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.14.0/docs/resources/alloydb_cluster#network_config AlloydbCluster#network_config}
	NetworkConfig *AlloydbClusterNetworkConfig `field:"optional" json:"networkConfig" yaml:"networkConfig"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.14.0/docs/resources/alloydb_cluster#project AlloydbCluster#project}.
	Project *string `field:"optional" json:"project" yaml:"project"`
	// restore_backup_source block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.14.0/docs/resources/alloydb_cluster#restore_backup_source AlloydbCluster#restore_backup_source}
	RestoreBackupSource *AlloydbClusterRestoreBackupSource `field:"optional" json:"restoreBackupSource" yaml:"restoreBackupSource"`
	// restore_continuous_backup_source block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.14.0/docs/resources/alloydb_cluster#restore_continuous_backup_source AlloydbCluster#restore_continuous_backup_source}
	RestoreContinuousBackupSource *AlloydbClusterRestoreContinuousBackupSource `field:"optional" json:"restoreContinuousBackupSource" yaml:"restoreContinuousBackupSource"`
	// secondary_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.14.0/docs/resources/alloydb_cluster#secondary_config AlloydbCluster#secondary_config}
	SecondaryConfig *AlloydbClusterSecondaryConfig `field:"optional" json:"secondaryConfig" yaml:"secondaryConfig"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.14.0/docs/resources/alloydb_cluster#timeouts AlloydbCluster#timeouts}
	Timeouts *AlloydbClusterTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

