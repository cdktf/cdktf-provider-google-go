package deploymentmanagerdeployment


type DeploymentManagerDeploymentTargetConfig struct {
	// The full YAML contents of your configuration file.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/deployment_manager_deployment#content DeploymentManagerDeployment#content}
	Content *string `field:"required" json:"content" yaml:"content"`
}

