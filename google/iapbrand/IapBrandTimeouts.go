package iapbrand


type IapBrandTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/iap_brand#create IapBrand#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/iap_brand#delete IapBrand#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
}

