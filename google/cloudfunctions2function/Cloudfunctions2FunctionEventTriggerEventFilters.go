package cloudfunctions2function


type Cloudfunctions2FunctionEventTriggerEventFilters struct {
	// 'Required.
	//
	// The name of a CloudEvents attribute.
	// Currently, only a subset of attributes are supported for filtering. Use the 'gcloud eventarc providers describe' command to learn more about events and their attributes.
	// Do not filter for the 'type' attribute here, as this is already achieved by the resource's 'event_type' attribute.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/cloudfunctions2_function#attribute Cloudfunctions2Function#attribute}
	Attribute *string `field:"required" json:"attribute" yaml:"attribute"`
	// Required.
	//
	// The value for the attribute.
	// If the operator field is set as 'match-path-pattern', this value can be a path pattern instead of an exact value.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/cloudfunctions2_function#value Cloudfunctions2Function#value}
	Value *string `field:"required" json:"value" yaml:"value"`
	// Optional.
	//
	// The operator used for matching the events with the value of
	// the filter. If not specified, only events that have an exact key-value
	// pair specified in the filter are matched.
	// The only allowed value is 'match-path-pattern'.
	// [See documentation on path patterns here](https://cloud.google.com/eventarc/docs/path-patterns)'
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/cloudfunctions2_function#operator Cloudfunctions2Function#operator}
	Operator *string `field:"optional" json:"operator" yaml:"operator"`
}

