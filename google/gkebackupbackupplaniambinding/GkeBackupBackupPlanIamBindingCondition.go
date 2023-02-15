package gkebackupbackupplaniambinding


type GkeBackupBackupPlanIamBindingCondition struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/gke_backup_backup_plan_iam_binding#expression GkeBackupBackupPlanIamBinding#expression}.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/gke_backup_backup_plan_iam_binding#title GkeBackupBackupPlanIamBinding#title}.
	Title *string `field:"required" json:"title" yaml:"title"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/gke_backup_backup_plan_iam_binding#description GkeBackupBackupPlanIamBinding#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

