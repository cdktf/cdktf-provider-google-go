package dataproccluster


type DataprocClusterClusterConfigGceClusterConfig struct {
	// By default, clusters are not restricted to internal IP addresses, and will have ephemeral external IP addresses assigned to each instance.
	//
	// If set to true, all instances in the cluster will only have internal IP addresses. Note: Private Google Access (also known as privateIpGoogleAccess) must be enabled on the subnetwork that the cluster will be launched in.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/dataproc_cluster#internal_ip_only DataprocCluster#internal_ip_only}
	InternalIpOnly interface{} `field:"optional" json:"internalIpOnly" yaml:"internalIpOnly"`
	// A map of the Compute Engine metadata entries to add to all instances.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/dataproc_cluster#metadata DataprocCluster#metadata}
	Metadata *map[string]*string `field:"optional" json:"metadata" yaml:"metadata"`
	// The name or self_link of the Google Compute Engine network to the cluster will be part of.
	//
	// Conflicts with subnetwork. If neither is specified, this defaults to the "default" network.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/dataproc_cluster#network DataprocCluster#network}
	Network *string `field:"optional" json:"network" yaml:"network"`
	// node_group_affinity block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/dataproc_cluster#node_group_affinity DataprocCluster#node_group_affinity}
	NodeGroupAffinity *DataprocClusterClusterConfigGceClusterConfigNodeGroupAffinity `field:"optional" json:"nodeGroupAffinity" yaml:"nodeGroupAffinity"`
	// reservation_affinity block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/dataproc_cluster#reservation_affinity DataprocCluster#reservation_affinity}
	ReservationAffinity *DataprocClusterClusterConfigGceClusterConfigReservationAffinity `field:"optional" json:"reservationAffinity" yaml:"reservationAffinity"`
	// The service account to be used by the Node VMs. If not specified, the "default" service account is used.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/dataproc_cluster#service_account DataprocCluster#service_account}
	ServiceAccount *string `field:"optional" json:"serviceAccount" yaml:"serviceAccount"`
	// The set of Google API scopes to be made available on all of the node VMs under the service_account specified.
	//
	// These can be either FQDNs, or scope aliases.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/dataproc_cluster#service_account_scopes DataprocCluster#service_account_scopes}
	ServiceAccountScopes *[]*string `field:"optional" json:"serviceAccountScopes" yaml:"serviceAccountScopes"`
	// shielded_instance_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/dataproc_cluster#shielded_instance_config DataprocCluster#shielded_instance_config}
	ShieldedInstanceConfig *DataprocClusterClusterConfigGceClusterConfigShieldedInstanceConfig `field:"optional" json:"shieldedInstanceConfig" yaml:"shieldedInstanceConfig"`
	// The name or self_link of the Google Compute Engine subnetwork the cluster will be part of. Conflicts with network.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/dataproc_cluster#subnetwork DataprocCluster#subnetwork}
	Subnetwork *string `field:"optional" json:"subnetwork" yaml:"subnetwork"`
	// The list of instance tags applied to instances in the cluster.
	//
	// Tags are used to identify valid sources or targets for network firewalls.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/dataproc_cluster#tags DataprocCluster#tags}
	Tags *[]*string `field:"optional" json:"tags" yaml:"tags"`
	// The GCP zone where your data is stored and used (i.e. where the master and the worker nodes will be created in). If region is set to 'global' (default) then zone is mandatory, otherwise GCP is able to make use of Auto Zone Placement to determine this automatically for you. Note: This setting additionally determines and restricts which computing resources are available for use with other configs such as cluster_config.master_config.machine_type and cluster_config.worker_config.machine_type.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/dataproc_cluster#zone DataprocCluster#zone}
	Zone *string `field:"optional" json:"zone" yaml:"zone"`
}

