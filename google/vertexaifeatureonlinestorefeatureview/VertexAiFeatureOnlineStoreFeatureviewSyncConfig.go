// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package vertexaifeatureonlinestorefeatureview


type VertexAiFeatureOnlineStoreFeatureviewSyncConfig struct {
	// Cron schedule (https://en.wikipedia.org/wiki/Cron) to launch scheduled runs. To explicitly set a timezone to the cron tab, apply a prefix in the cron tab: "CRON_TZ=${IANA_TIME_ZONE}" or "TZ=${IANA_TIME_ZONE}".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.22.0/docs/resources/vertex_ai_feature_online_store_featureview#cron VertexAiFeatureOnlineStoreFeatureview#cron}
	Cron *string `field:"optional" json:"cron" yaml:"cron"`
}

