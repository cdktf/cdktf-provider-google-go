package containernodepool


type ContainerNodePoolUpgradeSettings struct {
	// blue_green_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/container_node_pool#blue_green_settings ContainerNodePool#blue_green_settings}
	BlueGreenSettings *ContainerNodePoolUpgradeSettingsBlueGreenSettings `field:"optional" json:"blueGreenSettings" yaml:"blueGreenSettings"`
	// The number of additional nodes that can be added to the node pool during an upgrade.
	//
	// Increasing max_surge raises the number of nodes that can be upgraded simultaneously. Can be set to 0 or greater.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/container_node_pool#max_surge ContainerNodePool#max_surge}
	MaxSurge *float64 `field:"optional" json:"maxSurge" yaml:"maxSurge"`
	// The number of nodes that can be simultaneously unavailable during an upgrade.
	//
	// Increasing max_unavailable raises the number of nodes that can be upgraded in parallel. Can be set to 0 or greater.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/container_node_pool#max_unavailable ContainerNodePool#max_unavailable}
	MaxUnavailable *float64 `field:"optional" json:"maxUnavailable" yaml:"maxUnavailable"`
	// Update strategy for the given nodepool.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/container_node_pool#strategy ContainerNodePool#strategy}
	Strategy *string `field:"optional" json:"strategy" yaml:"strategy"`
}

