// Prebuilt google Provider for Terraform CDK (cdktf)
package google


type ContainerNodePoolManagement struct {
	// Whether the nodes will be automatically repaired.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/container_node_pool#auto_repair ContainerNodePool#auto_repair}
	AutoRepair interface{} `field:"optional" json:"autoRepair" yaml:"autoRepair"`
	// Whether the nodes will be automatically upgraded.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/container_node_pool#auto_upgrade ContainerNodePool#auto_upgrade}
	AutoUpgrade interface{} `field:"optional" json:"autoUpgrade" yaml:"autoUpgrade"`
}

