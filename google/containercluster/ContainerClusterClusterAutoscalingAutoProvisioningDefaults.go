package containercluster


type ContainerClusterClusterAutoscalingAutoProvisioningDefaults struct {
	// The Customer Managed Encryption Key used to encrypt the boot disk attached to each node in the node pool.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/container_cluster#boot_disk_kms_key ContainerCluster#boot_disk_kms_key}
	BootDiskKmsKey *string `field:"optional" json:"bootDiskKmsKey" yaml:"bootDiskKmsKey"`
	// Size of the disk attached to each node, specified in GB. The smallest allowed disk size is 10GB.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/container_cluster#disk_size ContainerCluster#disk_size}
	DiskSize *float64 `field:"optional" json:"diskSize" yaml:"diskSize"`
	// Type of the disk attached to each node.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/container_cluster#disk_type ContainerCluster#disk_type}
	DiskType *string `field:"optional" json:"diskType" yaml:"diskType"`
	// The default image type used by NAP once a new node pool is being created.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/container_cluster#image_type ContainerCluster#image_type}
	ImageType *string `field:"optional" json:"imageType" yaml:"imageType"`
	// management block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/container_cluster#management ContainerCluster#management}
	Management *ContainerClusterClusterAutoscalingAutoProvisioningDefaultsManagement `field:"optional" json:"management" yaml:"management"`
	// Minimum CPU platform to be used by this instance.
	//
	// The instance may be scheduled on the specified or newer CPU platform. Applicable values are the friendly names of CPU platforms, such as Intel Haswell.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/container_cluster#min_cpu_platform ContainerCluster#min_cpu_platform}
	MinCpuPlatform *string `field:"optional" json:"minCpuPlatform" yaml:"minCpuPlatform"`
	// Scopes that are used by NAP when creating node pools.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/container_cluster#oauth_scopes ContainerCluster#oauth_scopes}
	OauthScopes *[]*string `field:"optional" json:"oauthScopes" yaml:"oauthScopes"`
	// The Google Cloud Platform Service Account to be used by the node VMs.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/container_cluster#service_account ContainerCluster#service_account}
	ServiceAccount *string `field:"optional" json:"serviceAccount" yaml:"serviceAccount"`
	// shielded_instance_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/container_cluster#shielded_instance_config ContainerCluster#shielded_instance_config}
	ShieldedInstanceConfig *ContainerClusterClusterAutoscalingAutoProvisioningDefaultsShieldedInstanceConfig `field:"optional" json:"shieldedInstanceConfig" yaml:"shieldedInstanceConfig"`
	// upgrade_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/container_cluster#upgrade_settings ContainerCluster#upgrade_settings}
	UpgradeSettings *ContainerClusterClusterAutoscalingAutoProvisioningDefaultsUpgradeSettings `field:"optional" json:"upgradeSettings" yaml:"upgradeSettings"`
}

