// Prebuilt google Provider for Terraform CDK (cdktf)
package google


type GameServicesGameServerClusterConnectionInfo struct {
	// gke_cluster_reference block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/game_services_game_server_cluster#gke_cluster_reference GameServicesGameServerCluster#gke_cluster_reference}
	GkeClusterReference *GameServicesGameServerClusterConnectionInfoGkeClusterReference `field:"required" json:"gkeClusterReference" yaml:"gkeClusterReference"`
	// Namespace designated on the game server cluster where the game server instances will be created.
	//
	// The namespace existence will be validated
	// during creation.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/game_services_game_server_cluster#namespace GameServicesGameServerCluster#namespace}
	Namespace *string `field:"required" json:"namespace" yaml:"namespace"`
}

