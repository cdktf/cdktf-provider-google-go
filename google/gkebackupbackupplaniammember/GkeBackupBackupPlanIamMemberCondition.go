package gkebackupbackupplaniammember


type GkeBackupBackupPlanIamMemberCondition struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/gke_backup_backup_plan_iam_member#expression GkeBackupBackupPlanIamMember#expression}.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/gke_backup_backup_plan_iam_member#title GkeBackupBackupPlanIamMember#title}.
	Title *string `field:"required" json:"title" yaml:"title"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/gke_backup_backup_plan_iam_member#description GkeBackupBackupPlanIamMember#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

