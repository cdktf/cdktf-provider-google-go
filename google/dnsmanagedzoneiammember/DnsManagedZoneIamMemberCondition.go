package dnsmanagedzoneiammember


type DnsManagedZoneIamMemberCondition struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/dns_managed_zone_iam_member#expression DnsManagedZoneIamMember#expression}.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/dns_managed_zone_iam_member#title DnsManagedZoneIamMember#title}.
	Title *string `field:"required" json:"title" yaml:"title"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/dns_managed_zone_iam_member#description DnsManagedZoneIamMember#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

