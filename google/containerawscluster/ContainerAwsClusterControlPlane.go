package containerawscluster


type ContainerAwsClusterControlPlane struct {
	// aws_services_authentication block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/container_aws_cluster#aws_services_authentication ContainerAwsCluster#aws_services_authentication}
	AwsServicesAuthentication *ContainerAwsClusterControlPlaneAwsServicesAuthentication `field:"required" json:"awsServicesAuthentication" yaml:"awsServicesAuthentication"`
	// config_encryption block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/container_aws_cluster#config_encryption ContainerAwsCluster#config_encryption}
	ConfigEncryption *ContainerAwsClusterControlPlaneConfigEncryption `field:"required" json:"configEncryption" yaml:"configEncryption"`
	// database_encryption block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/container_aws_cluster#database_encryption ContainerAwsCluster#database_encryption}
	DatabaseEncryption *ContainerAwsClusterControlPlaneDatabaseEncryption `field:"required" json:"databaseEncryption" yaml:"databaseEncryption"`
	// The name of the AWS IAM instance pofile to assign to each control plane replica.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/container_aws_cluster#iam_instance_profile ContainerAwsCluster#iam_instance_profile}
	IamInstanceProfile *string `field:"required" json:"iamInstanceProfile" yaml:"iamInstanceProfile"`
	// The list of subnets where control plane replicas will run.
	//
	// A replica will be provisioned on each subnet and up to three values can be provided. Each subnet must be in a different AWS Availability Zone (AZ).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/container_aws_cluster#subnet_ids ContainerAwsCluster#subnet_ids}
	SubnetIds *[]*string `field:"required" json:"subnetIds" yaml:"subnetIds"`
	// The Kubernetes version to run on control plane replicas (e.g. `1.19.10-gke.1000`). You can list all supported versions on a given Google Cloud region by calling .
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/container_aws_cluster#version ContainerAwsCluster#version}
	Version *string `field:"required" json:"version" yaml:"version"`
	// Optional. The AWS instance type. When unspecified, it defaults to `m5.large`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/container_aws_cluster#instance_type ContainerAwsCluster#instance_type}
	InstanceType *string `field:"optional" json:"instanceType" yaml:"instanceType"`
	// main_volume block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/container_aws_cluster#main_volume ContainerAwsCluster#main_volume}
	MainVolume *ContainerAwsClusterControlPlaneMainVolume `field:"optional" json:"mainVolume" yaml:"mainVolume"`
	// proxy_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/container_aws_cluster#proxy_config ContainerAwsCluster#proxy_config}
	ProxyConfig *ContainerAwsClusterControlPlaneProxyConfig `field:"optional" json:"proxyConfig" yaml:"proxyConfig"`
	// root_volume block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/container_aws_cluster#root_volume ContainerAwsCluster#root_volume}
	RootVolume *ContainerAwsClusterControlPlaneRootVolume `field:"optional" json:"rootVolume" yaml:"rootVolume"`
	// Optional.
	//
	// The IDs of additional security groups to add to control plane replicas. The Anthos Multi-Cloud API will automatically create and manage security groups with the minimum rules needed for a functioning cluster.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/container_aws_cluster#security_group_ids ContainerAwsCluster#security_group_ids}
	SecurityGroupIds *[]*string `field:"optional" json:"securityGroupIds" yaml:"securityGroupIds"`
	// ssh_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/container_aws_cluster#ssh_config ContainerAwsCluster#ssh_config}
	SshConfig *ContainerAwsClusterControlPlaneSshConfig `field:"optional" json:"sshConfig" yaml:"sshConfig"`
	// Optional.
	//
	// A set of AWS resource tags to propagate to all underlying managed AWS resources. Specify at most 50 pairs containing alphanumerics, spaces, and symbols (.+-=_:@/). Keys can be up to 127 Unicode characters. Values can be up to 255 Unicode characters.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/container_aws_cluster#tags ContainerAwsCluster#tags}
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
}

