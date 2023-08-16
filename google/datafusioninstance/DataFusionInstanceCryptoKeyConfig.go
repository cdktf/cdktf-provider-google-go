package datafusioninstance


type DataFusionInstanceCryptoKeyConfig struct {
	// The name of the key which is used to encrypt/decrypt customer data.
	//
	// For key in Cloud KMS, the key should be in the format of projects/*\/locations/*\/keyRings/*\/cryptoKeys/*.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/data_fusion_instance#key_reference DataFusionInstance#key_reference}
	KeyReference *string `field:"required" json:"keyReference" yaml:"keyReference"`
}

