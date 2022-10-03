package cloudbuildtrigger


type CloudbuildTriggerBuildSource struct {
	// repo_source block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/cloudbuild_trigger#repo_source CloudbuildTrigger#repo_source}
	RepoSource *CloudbuildTriggerBuildSourceRepoSource `field:"optional" json:"repoSource" yaml:"repoSource"`
	// storage_source block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/cloudbuild_trigger#storage_source CloudbuildTrigger#storage_source}
	StorageSource *CloudbuildTriggerBuildSourceStorageSource `field:"optional" json:"storageSource" yaml:"storageSource"`
}

