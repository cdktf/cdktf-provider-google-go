package alloydbbackup

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type AlloydbBackupConfig struct {
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
	// The ID of the alloydb backup.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/alloydb_backup#backup_id AlloydbBackup#backup_id}
	BackupId *string `field:"required" json:"backupId" yaml:"backupId"`
	// The full resource name of the backup source cluster (e.g., projects/{project}/locations/{location}/clusters/{clusterId}).
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/alloydb_backup#cluster_name AlloydbBackup#cluster_name}
	ClusterName *string `field:"required" json:"clusterName" yaml:"clusterName"`
	// User-provided description of the backup.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/alloydb_backup#description AlloydbBackup#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/alloydb_backup#id AlloydbBackup#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// User-defined labels for the alloydb backup.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/alloydb_backup#labels AlloydbBackup#labels}
	Labels *map[string]*string `field:"optional" json:"labels" yaml:"labels"`
	// The location where the alloydb backup should reside.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/alloydb_backup#location AlloydbBackup#location}
	Location *string `field:"optional" json:"location" yaml:"location"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/alloydb_backup#project AlloydbBackup#project}.
	Project *string `field:"optional" json:"project" yaml:"project"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/alloydb_backup#timeouts AlloydbBackup#timeouts}
	Timeouts *AlloydbBackupTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}
