package appengineflexibleappversion


type AppEngineFlexibleAppVersionHandlers struct {
	// Actions to take when the user is not logged in. Possible values: ["AUTH_FAIL_ACTION_REDIRECT", "AUTH_FAIL_ACTION_UNAUTHORIZED"].
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/app_engine_flexible_app_version#auth_fail_action AppEngineFlexibleAppVersion#auth_fail_action}
	AuthFailAction *string `field:"optional" json:"authFailAction" yaml:"authFailAction"`
	// Methods to restrict access to a URL based on login status. Possible values: ["LOGIN_OPTIONAL", "LOGIN_ADMIN", "LOGIN_REQUIRED"].
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/app_engine_flexible_app_version#login AppEngineFlexibleAppVersion#login}
	Login *string `field:"optional" json:"login" yaml:"login"`
	// 30x code to use when performing redirects for the secure field. Possible values: ["REDIRECT_HTTP_RESPONSE_CODE_301", "REDIRECT_HTTP_RESPONSE_CODE_302", "REDIRECT_HTTP_RESPONSE_CODE_303", "REDIRECT_HTTP_RESPONSE_CODE_307"].
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/app_engine_flexible_app_version#redirect_http_response_code AppEngineFlexibleAppVersion#redirect_http_response_code}
	RedirectHttpResponseCode *string `field:"optional" json:"redirectHttpResponseCode" yaml:"redirectHttpResponseCode"`
	// script block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/app_engine_flexible_app_version#script AppEngineFlexibleAppVersion#script}
	Script *AppEngineFlexibleAppVersionHandlersScript `field:"optional" json:"script" yaml:"script"`
	// Security (HTTPS) enforcement for this URL. Possible values: ["SECURE_DEFAULT", "SECURE_NEVER", "SECURE_OPTIONAL", "SECURE_ALWAYS"].
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/app_engine_flexible_app_version#security_level AppEngineFlexibleAppVersion#security_level}
	SecurityLevel *string `field:"optional" json:"securityLevel" yaml:"securityLevel"`
	// static_files block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/app_engine_flexible_app_version#static_files AppEngineFlexibleAppVersion#static_files}
	StaticFiles *AppEngineFlexibleAppVersionHandlersStaticFiles `field:"optional" json:"staticFiles" yaml:"staticFiles"`
	// URL prefix.
	//
	// Uses regular expression syntax, which means regexp special characters must be escaped, but should not contain groupings.
	// All URLs that begin with this prefix are handled by this handler, using the portion of the URL after the prefix as part of the file path.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/app_engine_flexible_app_version#url_regex AppEngineFlexibleAppVersion#url_regex}
	UrlRegex *string `field:"optional" json:"urlRegex" yaml:"urlRegex"`
}

