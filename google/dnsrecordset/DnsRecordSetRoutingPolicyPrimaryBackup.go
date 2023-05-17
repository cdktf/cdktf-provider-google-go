package dnsrecordset


type DnsRecordSetRoutingPolicyPrimaryBackup struct {
	// backup_geo block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/dns_record_set#backup_geo DnsRecordSet#backup_geo}
	BackupGeo interface{} `field:"required" json:"backupGeo" yaml:"backupGeo"`
	// primary block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/dns_record_set#primary DnsRecordSet#primary}
	Primary *DnsRecordSetRoutingPolicyPrimaryBackupPrimary `field:"required" json:"primary" yaml:"primary"`
	// Specifies whether to enable fencing for backup geo queries.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/dns_record_set#enable_geo_fencing_for_backups DnsRecordSet#enable_geo_fencing_for_backups}
	EnableGeoFencingForBackups interface{} `field:"optional" json:"enableGeoFencingForBackups" yaml:"enableGeoFencingForBackups"`
	// Specifies the percentage of traffic to send to the backup targets even when the primary targets are healthy.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/dns_record_set#trickle_ratio DnsRecordSet#trickle_ratio}
	TrickleRatio *float64 `field:"optional" json:"trickleRatio" yaml:"trickleRatio"`
}

