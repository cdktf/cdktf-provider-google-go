package computeglobalforwardingrule


type ComputeGlobalForwardingRuleMetadataFilters struct {
	// filter_labels block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/compute_global_forwarding_rule#filter_labels ComputeGlobalForwardingRule#filter_labels}
	FilterLabels interface{} `field:"required" json:"filterLabels" yaml:"filterLabels"`
	// Specifies how individual `filterLabel` matches within the list of `filterLabel`s contribute towards the overall `metadataFilter` match.
	//
	// Supported values are:
	//
	// MATCH_ANY: At least one of the `filterLabels` must have a matching label in the provided metadata.
	// MATCH_ALL: All `filterLabels` must have matching labels in the provided metadata. Possible values: NOT_SET, MATCH_ALL, MATCH_ANY
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/compute_global_forwarding_rule#filter_match_criteria ComputeGlobalForwardingRule#filter_match_criteria}
	FilterMatchCriteria *string `field:"required" json:"filterMatchCriteria" yaml:"filterMatchCriteria"`
}

