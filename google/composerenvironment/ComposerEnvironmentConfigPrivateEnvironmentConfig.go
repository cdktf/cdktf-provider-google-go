package composerenvironment


type ComposerEnvironmentConfigPrivateEnvironmentConfig struct {
	// When specified, the environment will use Private Service Connect instead of VPC peerings to connect to Cloud SQL in the Tenant Project, and the PSC endpoint in the Customer Project will use an IP address from this subnetwork.
	//
	// This field is supported for Cloud Composer environments in versions composer-2.*.*-airflow-*.*.* and newer.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/composer_environment#cloud_composer_connection_subnetwork ComposerEnvironment#cloud_composer_connection_subnetwork}
	CloudComposerConnectionSubnetwork *string `field:"optional" json:"cloudComposerConnectionSubnetwork" yaml:"cloudComposerConnectionSubnetwork"`
	// The CIDR block from which IP range for Cloud Composer Network in tenant project will be reserved.
	//
	// Needs to be disjoint from private_cluster_config.master_ipv4_cidr_block and cloud_sql_ipv4_cidr_block. This field is supported for Cloud Composer environments in versions composer-2.*.*-airflow-*.*.* and newer.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/composer_environment#cloud_composer_network_ipv4_cidr_block ComposerEnvironment#cloud_composer_network_ipv4_cidr_block}
	CloudComposerNetworkIpv4CidrBlock *string `field:"optional" json:"cloudComposerNetworkIpv4CidrBlock" yaml:"cloudComposerNetworkIpv4CidrBlock"`
	// The CIDR block from which IP range in tenant project will be reserved for Cloud SQL.
	//
	// Needs to be disjoint from web_server_ipv4_cidr_block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/composer_environment#cloud_sql_ipv4_cidr_block ComposerEnvironment#cloud_sql_ipv4_cidr_block}
	CloudSqlIpv4CidrBlock *string `field:"optional" json:"cloudSqlIpv4CidrBlock" yaml:"cloudSqlIpv4CidrBlock"`
	// If true, access to the public endpoint of the GKE cluster is denied.
	//
	// If this field is set to true, ip_allocation_policy.use_ip_aliases must be set to true for Cloud Composer environments in versions composer-1.*.*-airflow-*.*.*.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/composer_environment#enable_private_endpoint ComposerEnvironment#enable_private_endpoint}
	EnablePrivateEndpoint interface{} `field:"optional" json:"enablePrivateEndpoint" yaml:"enablePrivateEndpoint"`
	// When enabled, IPs from public (non-RFC1918) ranges can be used for ip_allocation_policy.cluster_ipv4_cidr_block and ip_allocation_policy.service_ipv4_cidr_block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/composer_environment#enable_privately_used_public_ips ComposerEnvironment#enable_privately_used_public_ips}
	EnablePrivatelyUsedPublicIps interface{} `field:"optional" json:"enablePrivatelyUsedPublicIps" yaml:"enablePrivatelyUsedPublicIps"`
	// The IP range in CIDR notation to use for the hosted master network.
	//
	// This range is used for assigning internal IP addresses to the cluster master or set of masters and to the internal load balancer virtual IP. This range must not overlap with any other ranges in use within the cluster's network. If left blank, the default value of '172.16.0.0/28' is used.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/composer_environment#master_ipv4_cidr_block ComposerEnvironment#master_ipv4_cidr_block}
	MasterIpv4CidrBlock *string `field:"optional" json:"masterIpv4CidrBlock" yaml:"masterIpv4CidrBlock"`
	// The CIDR block from which IP range for web server will be reserved.
	//
	// Needs to be disjoint from master_ipv4_cidr_block and cloud_sql_ipv4_cidr_block. This field is supported for Cloud Composer environments in versions composer-1.*.*-airflow-*.*.*.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/composer_environment#web_server_ipv4_cidr_block ComposerEnvironment#web_server_ipv4_cidr_block}
	WebServerIpv4CidrBlock *string `field:"optional" json:"webServerIpv4CidrBlock" yaml:"webServerIpv4CidrBlock"`
}

