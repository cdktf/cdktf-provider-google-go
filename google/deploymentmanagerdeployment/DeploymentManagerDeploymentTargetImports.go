package deploymentmanagerdeployment


type DeploymentManagerDeploymentTargetImports struct {
	// The full contents of the template that you want to import.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/deployment_manager_deployment#content DeploymentManagerDeployment#content}
	Content *string `field:"optional" json:"content" yaml:"content"`
	// The name of the template to import, as declared in the YAML configuration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/deployment_manager_deployment#name DeploymentManagerDeployment#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
}

