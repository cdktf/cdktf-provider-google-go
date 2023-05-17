package composerenvironment


type ComposerEnvironmentConfigNodeConfig struct {
	// The disk size in GB used for node VMs.
	//
	// Minimum size is 20GB. If unspecified, defaults to 100GB. Cannot be updated. This field is supported for Cloud Composer environments in versions composer-1.*.*-airflow-*.*.*.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/composer_environment#disk_size_gb ComposerEnvironment#disk_size_gb}
	DiskSizeGb *float64 `field:"optional" json:"diskSizeGb" yaml:"diskSizeGb"`
	// Deploys 'ip-masq-agent' daemon set in the GKE cluster and defines nonMasqueradeCIDRs equals to pod IP range so IP masquerading is used for all destination addresses, except between pods traffic.
	//
	// See: https://cloud.google.com/kubernetes-engine/docs/how-to/ip-masquerade-agent
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/composer_environment#enable_ip_masq_agent ComposerEnvironment#enable_ip_masq_agent}
	EnableIpMasqAgent interface{} `field:"optional" json:"enableIpMasqAgent" yaml:"enableIpMasqAgent"`
	// Configuration for controlling how IPs are allocated in the GKE cluster. Cannot be updated.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/composer_environment#ip_allocation_policy ComposerEnvironment#ip_allocation_policy}
	IpAllocationPolicy interface{} `field:"optional" json:"ipAllocationPolicy" yaml:"ipAllocationPolicy"`
	// The Compute Engine machine type used for cluster instances, specified as a name or relative resource name.
	//
	// For example: "projects/{project}/zones/{zone}/machineTypes/{machineType}". Must belong to the enclosing environment's project and region/zone. This field is supported for Cloud Composer environments in versions composer-1.*.*-airflow-*.*.*.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/composer_environment#machine_type ComposerEnvironment#machine_type}
	MachineType *string `field:"optional" json:"machineType" yaml:"machineType"`
	// The Compute Engine machine type used for cluster instances, specified as a name or relative resource name.
	//
	// For example: "projects/{project}/zones/{zone}/machineTypes/{machineType}". Must belong to the enclosing environment's project and region/zone. The network must belong to the environment's project. If unspecified, the "default" network ID in the environment's project is used. If a Custom Subnet Network is provided, subnetwork must also be provided.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/composer_environment#network ComposerEnvironment#network}
	Network *string `field:"optional" json:"network" yaml:"network"`
	// The set of Google API scopes to be made available on all node VMs.
	//
	// Cannot be updated. If empty, defaults to ["https://www.googleapis.com/auth/cloud-platform"]. This field is supported for Cloud Composer environments in versions composer-1.*.*-airflow-*.*.*.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/composer_environment#oauth_scopes ComposerEnvironment#oauth_scopes}
	OauthScopes *[]*string `field:"optional" json:"oauthScopes" yaml:"oauthScopes"`
	// The Google Cloud Platform Service Account to be used by the node VMs.
	//
	// If a service account is not specified, the "default" Compute Engine service account is used. Cannot be updated. If given, note that the service account must have roles/composer.worker for any GCP resources created under the Cloud Composer Environment.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/composer_environment#service_account ComposerEnvironment#service_account}
	ServiceAccount *string `field:"optional" json:"serviceAccount" yaml:"serviceAccount"`
	// The Compute Engine subnetwork to be used for machine communications, , specified as a self-link, relative resource name (e.g. "projects/{project}/regions/{region}/subnetworks/{subnetwork}"), or by name. If subnetwork is provided, network must also be provided and the subnetwork must belong to the enclosing environment's project and region.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/composer_environment#subnetwork ComposerEnvironment#subnetwork}
	Subnetwork *string `field:"optional" json:"subnetwork" yaml:"subnetwork"`
	// The list of instance tags applied to all node VMs.
	//
	// Tags are used to identify valid sources or targets for network firewalls. Each tag within the list must comply with RFC1035. Cannot be updated.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/composer_environment#tags ComposerEnvironment#tags}
	Tags *[]*string `field:"optional" json:"tags" yaml:"tags"`
	// The Compute Engine zone in which to deploy the VMs running the Apache Airflow software, specified as the zone name or relative resource name (e.g. "projects/{project}/zones/{zone}"). Must belong to the enclosing environment's project and region. This field is supported for Cloud Composer environments in versions composer-1.*.*-airflow-*.*.*.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/composer_environment#zone ComposerEnvironment#zone}
	Zone *string `field:"optional" json:"zone" yaml:"zone"`
}

