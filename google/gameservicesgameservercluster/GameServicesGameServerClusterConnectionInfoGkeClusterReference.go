package gameservicesgameservercluster


type GameServicesGameServerClusterConnectionInfoGkeClusterReference struct {
	// The full or partial name of a GKE cluster, using one of the following forms:.
	//
	// 'projects/{project_id}/locations/{location}/clusters/{cluster_id}'
	// 'locations/{location}/clusters/{cluster_id}'
	// '{cluster_id}'
	//
	// If project and location are not specified, the project and location of the
	// GameServerCluster resource are used to generate the full name of the
	// GKE cluster.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/game_services_game_server_cluster#cluster GameServicesGameServerCluster#cluster}
	Cluster *string `field:"required" json:"cluster" yaml:"cluster"`
}

