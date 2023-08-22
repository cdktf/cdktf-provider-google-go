package gkebackupbackupplaniambinding


type GkeBackupBackupPlanIamBindingCondition struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/gke_backup_backup_plan_iam_binding#expression GkeBackupBackupPlanIamBinding#expression}.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/gke_backup_backup_plan_iam_binding#title GkeBackupBackupPlanIamBinding#title}.
	Title *string `field:"required" json:"title" yaml:"title"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/gke_backup_backup_plan_iam_binding#description GkeBackupBackupPlanIamBinding#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

