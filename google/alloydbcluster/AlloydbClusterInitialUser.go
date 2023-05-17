package alloydbcluster


type AlloydbClusterInitialUser struct {
	// The initial password for the user.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/alloydb_cluster#password AlloydbCluster#password}
	Password *string `field:"required" json:"password" yaml:"password"`
	// The database username.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/alloydb_cluster#user AlloydbCluster#user}
	User *string `field:"optional" json:"user" yaml:"user"`
}

