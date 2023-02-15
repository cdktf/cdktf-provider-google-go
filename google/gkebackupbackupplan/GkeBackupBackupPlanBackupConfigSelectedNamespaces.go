package gkebackupbackupplan


type GkeBackupBackupPlanBackupConfigSelectedNamespaces struct {
	// A list of Kubernetes Namespaces.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/gke_backup_backup_plan#namespaces GkeBackupBackupPlan#namespaces}
	Namespaces *[]*string `field:"required" json:"namespaces" yaml:"namespaces"`
}

