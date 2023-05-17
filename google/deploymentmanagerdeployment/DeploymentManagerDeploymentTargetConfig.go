package deploymentmanagerdeployment


type DeploymentManagerDeploymentTargetConfig struct {
	// The full YAML contents of your configuration file.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/deployment_manager_deployment#content DeploymentManagerDeployment#content}
	Content *string `field:"required" json:"content" yaml:"content"`
}

