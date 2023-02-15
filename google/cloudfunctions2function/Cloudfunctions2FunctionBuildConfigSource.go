package cloudfunctions2function


type Cloudfunctions2FunctionBuildConfigSource struct {
	// repo_source block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/cloudfunctions2_function#repo_source Cloudfunctions2Function#repo_source}
	RepoSource *Cloudfunctions2FunctionBuildConfigSourceRepoSource `field:"optional" json:"repoSource" yaml:"repoSource"`
	// storage_source block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/cloudfunctions2_function#storage_source Cloudfunctions2Function#storage_source}
	StorageSource *Cloudfunctions2FunctionBuildConfigSourceStorageSource `field:"optional" json:"storageSource" yaml:"storageSource"`
}

