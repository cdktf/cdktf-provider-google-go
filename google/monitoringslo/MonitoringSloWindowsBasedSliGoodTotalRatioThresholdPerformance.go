package monitoringslo


type MonitoringSloWindowsBasedSliGoodTotalRatioThresholdPerformance struct {
	// distribution_cut block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/monitoring_slo#distribution_cut MonitoringSlo#distribution_cut}
	DistributionCut *MonitoringSloWindowsBasedSliGoodTotalRatioThresholdPerformanceDistributionCut `field:"optional" json:"distributionCut" yaml:"distributionCut"`
	// good_total_ratio block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/monitoring_slo#good_total_ratio MonitoringSlo#good_total_ratio}
	GoodTotalRatio *MonitoringSloWindowsBasedSliGoodTotalRatioThresholdPerformanceGoodTotalRatio `field:"optional" json:"goodTotalRatio" yaml:"goodTotalRatio"`
}

