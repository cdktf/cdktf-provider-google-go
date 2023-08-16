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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/alloydb_cluster#cluster_id AlloydbCluster#cluster_id}
	ClusterId *string `field:"required" json:"clusterId" yaml:"clusterId"`
	// The location where the alloydb cluster should reside.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/alloydb_cluster#location AlloydbCluster#location}
	Location *string `field:"required" json:"location" yaml:"location"`
	// The relative resource name of the VPC network on which the instance can be accessed.
	//
	// It is specified in the following form:
	//
	// "projects/{projectNumber}/global/networks/{network_id}".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/alloydb_cluster#network AlloydbCluster#network}
	Network *string `field:"required" json:"network" yaml:"network"`
	// automated_backup_policy block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/alloydb_cluster#automated_backup_policy AlloydbCluster#automated_backup_policy}
	AutomatedBackupPolicy *AlloydbClusterAutomatedBackupPolicy `field:"optional" json:"automatedBackupPolicy" yaml:"automatedBackupPolicy"`
	// continuous_backup_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/alloydb_cluster#continuous_backup_config AlloydbCluster#continuous_backup_config}
	ContinuousBackupConfig *AlloydbClusterContinuousBackupConfig `field:"optional" json:"continuousBackupConfig" yaml:"continuousBackupConfig"`
	// User-settable and human-readable display name for the Cluster.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/alloydb_cluster#display_name AlloydbCluster#display_name}
	DisplayName *string `field:"optional" json:"displayName" yaml:"displayName"`
	// encryption_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/alloydb_cluster#encryption_config AlloydbCluster#encryption_config}
	EncryptionConfig *AlloydbClusterEncryptionConfig `field:"optional" json:"encryptionConfig" yaml:"encryptionConfig"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/alloydb_cluster#id AlloydbCluster#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// initial_user block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/alloydb_cluster#initial_user AlloydbCluster#initial_user}
	InitialUser *AlloydbClusterInitialUser `field:"optional" json:"initialUser" yaml:"initialUser"`
	// User-defined labels for the alloydb cluster.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/alloydb_cluster#labels AlloydbCluster#labels}
	Labels *map[string]*string `field:"optional" json:"labels" yaml:"labels"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/alloydb_cluster#project AlloydbCluster#project}.
	Project *string `field:"optional" json:"project" yaml:"project"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/alloydb_cluster#timeouts AlloydbCluster#timeouts}
	Timeouts *AlloydbClusterTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

