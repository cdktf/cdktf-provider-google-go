package dnsmanagedzoneiambinding


type DnsManagedZoneIamBindingCondition struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/dns_managed_zone_iam_binding#expression DnsManagedZoneIamBinding#expression}.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/dns_managed_zone_iam_binding#title DnsManagedZoneIamBinding#title}.
	Title *string `field:"required" json:"title" yaml:"title"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/dns_managed_zone_iam_binding#description DnsManagedZoneIamBinding#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}
