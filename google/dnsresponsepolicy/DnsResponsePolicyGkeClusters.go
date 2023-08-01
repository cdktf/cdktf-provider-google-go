package dnsresponsepolicy


type DnsResponsePolicyGkeClusters struct {
	// The resource name of the cluster to bind this ManagedZone to. This should be specified in the format like 'projects/*\/locations/*\/clusters/*'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/dns_response_policy#gke_cluster_name DnsResponsePolicy#gke_cluster_name}
	GkeClusterName *string `field:"required" json:"gkeClusterName" yaml:"gkeClusterName"`
}

