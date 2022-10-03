package monitoringslo


type MonitoringSloRequestBasedSli struct {
	// distribution_cut block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/monitoring_slo#distribution_cut MonitoringSlo#distribution_cut}
	DistributionCut *MonitoringSloRequestBasedSliDistributionCut `field:"optional" json:"distributionCut" yaml:"distributionCut"`
	// good_total_ratio block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/monitoring_slo#good_total_ratio MonitoringSlo#good_total_ratio}
	GoodTotalRatio *MonitoringSloRequestBasedSliGoodTotalRatio `field:"optional" json:"goodTotalRatio" yaml:"goodTotalRatio"`
}

