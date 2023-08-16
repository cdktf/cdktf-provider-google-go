package appengineflexibleappversion


type AppEngineFlexibleAppVersionHandlersStaticFiles struct {
	// Whether files should also be uploaded as code data.
	//
	// By default, files declared in static file handlers are
	// uploaded as static data and are only served to end users; they cannot be read by the application. If enabled,
	// uploads are charged against both your code and static data storage resource quotas.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/app_engine_flexible_app_version#application_readable AppEngineFlexibleAppVersion#application_readable}
	ApplicationReadable interface{} `field:"optional" json:"applicationReadable" yaml:"applicationReadable"`
	// Time a static file served by this handler should be cached by web proxies and browsers.
	//
	// A duration in seconds with up to nine fractional digits, terminated by 's'. Example "3.5s".
	// Default is '0s'
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/app_engine_flexible_app_version#expiration AppEngineFlexibleAppVersion#expiration}
	Expiration *string `field:"optional" json:"expiration" yaml:"expiration"`
	// HTTP headers to use for all responses from these URLs. An object containing a list of "key:value" value pairs.".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/app_engine_flexible_app_version#http_headers AppEngineFlexibleAppVersion#http_headers}
	HttpHeaders *map[string]*string `field:"optional" json:"httpHeaders" yaml:"httpHeaders"`
	// MIME type used to serve all files served by this handler.
	//
	// Defaults to file-specific MIME types, which are derived from each file's filename extension.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/app_engine_flexible_app_version#mime_type AppEngineFlexibleAppVersion#mime_type}
	MimeType *string `field:"optional" json:"mimeType" yaml:"mimeType"`
	// Path to the static files matched by the URL pattern, from the application root directory.
	//
	// The path can refer to text matched in groupings in the URL pattern.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/app_engine_flexible_app_version#path AppEngineFlexibleAppVersion#path}
	Path *string `field:"optional" json:"path" yaml:"path"`
	// Whether this handler should match the request if the file referenced by the handler does not exist.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/app_engine_flexible_app_version#require_matching_file AppEngineFlexibleAppVersion#require_matching_file}
	RequireMatchingFile interface{} `field:"optional" json:"requireMatchingFile" yaml:"requireMatchingFile"`
	// Regular expression that matches the file paths for all files that should be referenced by this handler.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/app_engine_flexible_app_version#upload_path_regex AppEngineFlexibleAppVersion#upload_path_regex}
	UploadPathRegex *string `field:"optional" json:"uploadPathRegex" yaml:"uploadPathRegex"`
}

