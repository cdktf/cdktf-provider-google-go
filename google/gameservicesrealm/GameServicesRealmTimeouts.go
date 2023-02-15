package gameservicesrealm


type GameServicesRealmTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/game_services_realm#create GameServicesRealm#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/game_services_realm#delete GameServicesRealm#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/game_services_realm#update GameServicesRealm#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

