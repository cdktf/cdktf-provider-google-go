package alloydbcluster


type AlloydbClusterInitialUser struct {
	// The initial password for the user.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/alloydb_cluster#password AlloydbCluster#password}
	Password *string `field:"required" json:"password" yaml:"password"`
	// The database username.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/alloydb_cluster#user AlloydbCluster#user}
	User *string `field:"optional" json:"user" yaml:"user"`
}

