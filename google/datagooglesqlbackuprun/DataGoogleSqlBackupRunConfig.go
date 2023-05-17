package datagooglesqlbackuprun

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataGoogleSqlBackupRunConfig struct {
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
	// Name of the database instance.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/data-sources/sql_backup_run#instance DataGoogleSqlBackupRun#instance}
	Instance *string `field:"required" json:"instance" yaml:"instance"`
	// The identifier for this backup run.
	//
	// Unique only for a specific Cloud SQL instance. If left empty and multiple backups exist for the instance, most_recent must be set to true.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/data-sources/sql_backup_run#backup_id DataGoogleSqlBackupRun#backup_id}
	BackupId *float64 `field:"optional" json:"backupId" yaml:"backupId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/data-sources/sql_backup_run#id DataGoogleSqlBackupRun#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Toggles use of the most recent backup run if multiple backups exist for a Cloud SQL instance.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/data-sources/sql_backup_run#most_recent DataGoogleSqlBackupRun#most_recent}
	MostRecent interface{} `field:"optional" json:"mostRecent" yaml:"mostRecent"`
	// Project ID of the project that contains the instance.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/data-sources/sql_backup_run#project DataGoogleSqlBackupRun#project}
	Project *string `field:"optional" json:"project" yaml:"project"`
}

