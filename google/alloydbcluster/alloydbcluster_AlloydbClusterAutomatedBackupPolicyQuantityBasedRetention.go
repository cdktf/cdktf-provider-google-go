package alloydbcluster


type AlloydbClusterAutomatedBackupPolicyQuantityBasedRetention struct {
	// The number of backups to retain.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/alloydb_cluster#count AlloydbCluster#count}
	Count *float64 `field:"optional" json:"count" yaml:"count"`
}

