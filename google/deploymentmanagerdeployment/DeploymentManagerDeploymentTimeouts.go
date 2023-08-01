package deploymentmanagerdeployment


type DeploymentManagerDeploymentTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/deployment_manager_deployment#create DeploymentManagerDeployment#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/deployment_manager_deployment#delete DeploymentManagerDeployment#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/deployment_manager_deployment#update DeploymentManagerDeployment#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

