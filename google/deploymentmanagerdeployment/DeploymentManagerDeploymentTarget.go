package deploymentmanagerdeployment


type DeploymentManagerDeploymentTarget struct {
	// config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/deployment_manager_deployment#config DeploymentManagerDeployment#config}
	Config *DeploymentManagerDeploymentTargetConfig `field:"required" json:"config" yaml:"config"`
	// imports block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/deployment_manager_deployment#imports DeploymentManagerDeployment#imports}
	Imports interface{} `field:"optional" json:"imports" yaml:"imports"`
}

