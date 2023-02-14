package alloydbcluster


type AlloydbClusterAutomatedBackupPolicyTimeBasedRetention struct {
	// The retention period. A duration in seconds with up to nine fractional digits, terminated by 's'. Example: "3.5s".
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/alloydb_cluster#retention_period AlloydbCluster#retention_period}
	RetentionPeriod *string `field:"optional" json:"retentionPeriod" yaml:"retentionPeriod"`
}

