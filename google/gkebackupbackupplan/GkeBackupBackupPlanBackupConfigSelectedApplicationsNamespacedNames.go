package gkebackupbackupplan


type GkeBackupBackupPlanBackupConfigSelectedApplicationsNamespacedNames struct {
	// The name of a Kubernetes Resource.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/gke_backup_backup_plan#name GkeBackupBackupPlan#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The namespace of a Kubernetes Resource.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/gke_backup_backup_plan#namespace GkeBackupBackupPlan#namespace}
	Namespace *string `field:"required" json:"namespace" yaml:"namespace"`
}

