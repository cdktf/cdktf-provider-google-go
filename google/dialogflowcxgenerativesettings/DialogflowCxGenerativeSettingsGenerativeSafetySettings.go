// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dialogflowcxgenerativesettings


type DialogflowCxGenerativeSettingsGenerativeSafetySettings struct {
	// banned_phrases block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.50.0/docs/resources/dialogflow_cx_generative_settings#banned_phrases DialogflowCxGenerativeSettings#banned_phrases}
	BannedPhrases interface{} `field:"optional" json:"bannedPhrases" yaml:"bannedPhrases"`
	// Optional. Default phrase match strategy for banned phrases. See [PhraseMatchStrategy](https://cloud.google.com/dialogflow/cx/docs/reference/rest/v3/GenerativeSettings#phrasematchstrategy) for valid values.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.50.0/docs/resources/dialogflow_cx_generative_settings#default_banned_phrase_match_strategy DialogflowCxGenerativeSettings#default_banned_phrase_match_strategy}
	DefaultBannedPhraseMatchStrategy *string `field:"optional" json:"defaultBannedPhraseMatchStrategy" yaml:"defaultBannedPhraseMatchStrategy"`
}

