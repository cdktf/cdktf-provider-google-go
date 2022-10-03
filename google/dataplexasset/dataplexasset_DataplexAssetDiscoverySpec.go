package dataplexasset


type DataplexAssetDiscoverySpec struct {
	// Required. Whether discovery is enabled.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/dataplex_asset#enabled DataplexAsset#enabled}
	Enabled interface{} `field:"required" json:"enabled" yaml:"enabled"`
	// csv_options block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/dataplex_asset#csv_options DataplexAsset#csv_options}
	CsvOptions *DataplexAssetDiscoverySpecCsvOptions `field:"optional" json:"csvOptions" yaml:"csvOptions"`
	// Optional.
	//
	// The list of patterns to apply for selecting data to exclude during discovery. For Cloud Storage bucket assets, these are interpreted as glob patterns used to match object names. For BigQuery dataset assets, these are interpreted as patterns to match table names.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/dataplex_asset#exclude_patterns DataplexAsset#exclude_patterns}
	ExcludePatterns *[]*string `field:"optional" json:"excludePatterns" yaml:"excludePatterns"`
	// Optional.
	//
	// The list of patterns to apply for selecting data to include during discovery if only a subset of the data should considered. For Cloud Storage bucket assets, these are interpreted as glob patterns used to match object names. For BigQuery dataset assets, these are interpreted as patterns to match table names.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/dataplex_asset#include_patterns DataplexAsset#include_patterns}
	IncludePatterns *[]*string `field:"optional" json:"includePatterns" yaml:"includePatterns"`
	// json_options block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/dataplex_asset#json_options DataplexAsset#json_options}
	JsonOptions *DataplexAssetDiscoverySpecJsonOptions `field:"optional" json:"jsonOptions" yaml:"jsonOptions"`
	// Optional.
	//
	// Cron schedule (https://en.wikipedia.org/wiki/Cron) for running discovery periodically. Successive discovery runs must be scheduled at least 60 minutes apart. The default value is to run discovery every 60 minutes. To explicitly set a timezone to the cron tab, apply a prefix in the cron tab: "CRON_TZ=${IANA_TIME_ZONE}" or TZ=${IANA_TIME_ZONE}". The ${IANA_TIME_ZONE} may only be a valid string from IANA time zone database. For example, "CRON_TZ=America/New_York 1 * * * *", or "TZ=America/New_York 1 * * * *".
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/dataplex_asset#schedule DataplexAsset#schedule}
	Schedule *string `field:"optional" json:"schedule" yaml:"schedule"`
}

