package accesscontextmanageraccesslevels


type AccessContextManagerAccessLevelsAccessLevelsBasicConditionsDevicePolicy struct {
	// A list of allowed device management levels. An empty list allows all management levels. Possible values: ["MANAGEMENT_UNSPECIFIED", "NONE", "BASIC", "COMPLETE"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/resources/access_context_manager_access_levels#allowed_device_management_levels AccessContextManagerAccessLevels#allowed_device_management_levels}
	AllowedDeviceManagementLevels *[]*string `field:"optional" json:"allowedDeviceManagementLevels" yaml:"allowedDeviceManagementLevels"`
	// A list of allowed encryptions statuses. An empty list allows all statuses. Possible values: ["ENCRYPTION_UNSPECIFIED", "ENCRYPTION_UNSUPPORTED", "UNENCRYPTED", "ENCRYPTED"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/resources/access_context_manager_access_levels#allowed_encryption_statuses AccessContextManagerAccessLevels#allowed_encryption_statuses}
	AllowedEncryptionStatuses *[]*string `field:"optional" json:"allowedEncryptionStatuses" yaml:"allowedEncryptionStatuses"`
	// os_constraints block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/resources/access_context_manager_access_levels#os_constraints AccessContextManagerAccessLevels#os_constraints}
	OsConstraints interface{} `field:"optional" json:"osConstraints" yaml:"osConstraints"`
	// Whether the device needs to be approved by the customer admin.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/resources/access_context_manager_access_levels#require_admin_approval AccessContextManagerAccessLevels#require_admin_approval}
	RequireAdminApproval interface{} `field:"optional" json:"requireAdminApproval" yaml:"requireAdminApproval"`
	// Whether the device needs to be corp owned.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/resources/access_context_manager_access_levels#require_corp_owned AccessContextManagerAccessLevels#require_corp_owned}
	RequireCorpOwned interface{} `field:"optional" json:"requireCorpOwned" yaml:"requireCorpOwned"`
	// Whether or not screenlock is required for the DevicePolicy to be true. Defaults to false.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/resources/access_context_manager_access_levels#require_screen_lock AccessContextManagerAccessLevels#require_screen_lock}
	RequireScreenLock interface{} `field:"optional" json:"requireScreenLock" yaml:"requireScreenLock"`
}

