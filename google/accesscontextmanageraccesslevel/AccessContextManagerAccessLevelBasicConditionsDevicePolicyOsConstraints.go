package accesscontextmanageraccesslevel


type AccessContextManagerAccessLevelBasicConditionsDevicePolicyOsConstraints struct {
	// The operating system type of the device. Possible values: ["OS_UNSPECIFIED", "DESKTOP_MAC", "DESKTOP_WINDOWS", "DESKTOP_LINUX", "DESKTOP_CHROME_OS", "ANDROID", "IOS"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/access_context_manager_access_level#os_type AccessContextManagerAccessLevel#os_type}
	OsType *string `field:"required" json:"osType" yaml:"osType"`
	// The minimum allowed OS version.
	//
	// If not set, any version
	// of this OS satisfies the constraint.
	// Format: "major.minor.patch" such as "10.5.301", "9.2.1".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/access_context_manager_access_level#minimum_version AccessContextManagerAccessLevel#minimum_version}
	MinimumVersion *string `field:"optional" json:"minimumVersion" yaml:"minimumVersion"`
	// If you specify DESKTOP_CHROME_OS for osType, you can optionally include requireVerifiedChromeOs to require Chrome Verified Access.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/access_context_manager_access_level#require_verified_chrome_os AccessContextManagerAccessLevel#require_verified_chrome_os}
	RequireVerifiedChromeOs interface{} `field:"optional" json:"requireVerifiedChromeOs" yaml:"requireVerifiedChromeOs"`
}

