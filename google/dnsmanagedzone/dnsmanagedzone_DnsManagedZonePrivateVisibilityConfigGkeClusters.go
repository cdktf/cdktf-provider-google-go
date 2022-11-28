package dnsmanagedzone


type DnsManagedZonePrivateVisibilityConfigGkeClusters struct {
	// The resource name of the cluster to bind this ManagedZone to.
	//
	// This should be specified in the format like
	// 'projects/*\/locations/*\/clusters/*'
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/dns_managed_zone#gke_cluster_name DnsManagedZone#gke_cluster_name}
	GkeClusterName *string `field:"required" json:"gkeClusterName" yaml:"gkeClusterName"`
}

