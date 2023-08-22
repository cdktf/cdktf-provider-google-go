package gkebackupbackupplan


type GkeBackupBackupPlanBackupConfigSelectedApplicationsNamespacedNames struct {
	// The name of a Kubernetes Resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/gke_backup_backup_plan#name GkeBackupBackupPlan#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The namespace of a Kubernetes Resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/gke_backup_backup_plan#namespace GkeBackupBackupPlan#namespace}
	Namespace *string `field:"required" json:"namespace" yaml:"namespace"`
}

