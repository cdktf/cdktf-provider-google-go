package appengineapplicationurldispatchrules


type AppEngineApplicationUrlDispatchRulesDispatchRules struct {
	// Pathname within the host.
	//
	// Must start with a "/". A single "*" can be included at the end of the path.
	// The sum of the lengths of the domain and path may not exceed 100 characters.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/app_engine_application_url_dispatch_rules#path AppEngineApplicationUrlDispatchRules#path}
	Path *string `field:"required" json:"path" yaml:"path"`
	// Pathname within the host.
	//
	// Must start with a "/". A single "*" can be included at the end of the path.
	// The sum of the lengths of the domain and path may not exceed 100 characters.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/app_engine_application_url_dispatch_rules#service AppEngineApplicationUrlDispatchRules#service}
	Service *string `field:"required" json:"service" yaml:"service"`
	// Domain name to match against.
	//
	// The wildcard "*" is supported if specified before a period: "*.".
	// Defaults to matching all domains: "*".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/app_engine_application_url_dispatch_rules#domain AppEngineApplicationUrlDispatchRules#domain}
	Domain *string `field:"optional" json:"domain" yaml:"domain"`
}

