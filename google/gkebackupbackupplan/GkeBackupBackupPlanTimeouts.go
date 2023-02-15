package gkebackupbackupplan


type GkeBackupBackupPlanTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/gke_backup_backup_plan#create GkeBackupBackupPlan#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/gke_backup_backup_plan#delete GkeBackupBackupPlan#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/gke_backup_backup_plan#update GkeBackupBackupPlan#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

