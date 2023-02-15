package cloudfunctionsfunction


type CloudfunctionsFunctionSourceRepository struct {
	// The URL pointing to the hosted repository where the function is defined.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/cloudfunctions_function#url CloudfunctionsFunction#url}
	Url *string `field:"required" json:"url" yaml:"url"`
}

