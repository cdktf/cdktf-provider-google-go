package appengineservicesplittraffic


type AppEngineServiceSplitTrafficTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/app_engine_service_split_traffic#create AppEngineServiceSplitTraffic#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/app_engine_service_split_traffic#delete AppEngineServiceSplitTraffic#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/app_engine_service_split_traffic#update AppEngineServiceSplitTraffic#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

