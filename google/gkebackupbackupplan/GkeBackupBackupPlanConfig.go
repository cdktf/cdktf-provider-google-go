package gkebackupbackupplan

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type GkeBackupBackupPlanConfig struct {
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
	// The source cluster from which Backups will be created via this BackupPlan.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/gke_backup_backup_plan#cluster GkeBackupBackupPlan#cluster}
	Cluster *string `field:"required" json:"cluster" yaml:"cluster"`
	// The region of the Backup Plan.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/gke_backup_backup_plan#location GkeBackupBackupPlan#location}
	Location *string `field:"required" json:"location" yaml:"location"`
	// The full name of the BackupPlan Resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/gke_backup_backup_plan#name GkeBackupBackupPlan#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// backup_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/gke_backup_backup_plan#backup_config GkeBackupBackupPlan#backup_config}
	BackupConfig *GkeBackupBackupPlanBackupConfig `field:"optional" json:"backupConfig" yaml:"backupConfig"`
	// backup_schedule block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/gke_backup_backup_plan#backup_schedule GkeBackupBackupPlan#backup_schedule}
	BackupSchedule *GkeBackupBackupPlanBackupSchedule `field:"optional" json:"backupSchedule" yaml:"backupSchedule"`
	// This flag indicates whether this BackupPlan has been deactivated.
	//
	// Setting this field to True locks the BackupPlan such that no further updates will be allowed
	// (except deletes), including the deactivated field itself. It also prevents any new Backups
	// from being created via this BackupPlan (including scheduled Backups).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/gke_backup_backup_plan#deactivated GkeBackupBackupPlan#deactivated}
	Deactivated interface{} `field:"optional" json:"deactivated" yaml:"deactivated"`
	// User specified descriptive string for this BackupPlan.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/gke_backup_backup_plan#description GkeBackupBackupPlan#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/gke_backup_backup_plan#id GkeBackupBackupPlan#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Description: A set of custom labels supplied by the user.
	//
	// A list of key->value pairs.
	// Example: { "name": "wrench", "mass": "1.3kg", "count": "3" }.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/gke_backup_backup_plan#labels GkeBackupBackupPlan#labels}
	Labels *map[string]*string `field:"optional" json:"labels" yaml:"labels"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/gke_backup_backup_plan#project GkeBackupBackupPlan#project}.
	Project *string `field:"optional" json:"project" yaml:"project"`
	// retention_policy block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/gke_backup_backup_plan#retention_policy GkeBackupBackupPlan#retention_policy}
	RetentionPolicy *GkeBackupBackupPlanRetentionPolicy `field:"optional" json:"retentionPolicy" yaml:"retentionPolicy"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/gke_backup_backup_plan#timeouts GkeBackupBackupPlan#timeouts}
	Timeouts *GkeBackupBackupPlanTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

