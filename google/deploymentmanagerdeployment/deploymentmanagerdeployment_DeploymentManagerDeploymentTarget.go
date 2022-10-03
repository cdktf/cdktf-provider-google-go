package deploymentmanagerdeployment


type DeploymentManagerDeploymentTarget struct {
	// config block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/deployment_manager_deployment#config DeploymentManagerDeployment#config}
	Config *DeploymentManagerDeploymentTargetConfig `field:"required" json:"config" yaml:"config"`
	// imports block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/deployment_manager_deployment#imports DeploymentManagerDeployment#imports}
	Imports interface{} `field:"optional" json:"imports" yaml:"imports"`
}

