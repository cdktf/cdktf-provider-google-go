// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dialogflowcxsecuritysettings

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DialogflowCxSecuritySettingsConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktf.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// The human-readable name of the security settings, unique within the location.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.42.0/docs/resources/dialogflow_cx_security_settings#display_name DialogflowCxSecuritySettings#display_name}
	DisplayName *string `field:"required" json:"displayName" yaml:"displayName"`
	// The location these settings are located in.
	//
	// Settings can only be applied to an agent in the same location.
	// See [Available Regions](https://cloud.google.com/dialogflow/cx/docs/concept/region#avail) for a list of supported locations.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.42.0/docs/resources/dialogflow_cx_security_settings#location DialogflowCxSecuritySettings#location}
	Location *string `field:"required" json:"location" yaml:"location"`
	// audio_export_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.42.0/docs/resources/dialogflow_cx_security_settings#audio_export_settings DialogflowCxSecuritySettings#audio_export_settings}
	AudioExportSettings *DialogflowCxSecuritySettingsAudioExportSettings `field:"optional" json:"audioExportSettings" yaml:"audioExportSettings"`
	// [DLP](https://cloud.google.com/dlp/docs) deidentify template name. Use this template to define de-identification configuration for the content. If empty, Dialogflow replaces sensitive info with [redacted] text. Note: deidentifyTemplate must be located in the same region as the SecuritySettings. Format: projects/<Project ID>/locations/<Location ID>/deidentifyTemplates/<Template ID> OR organizations/<Organization ID>/locations/<Location ID>/deidentifyTemplates/<Template ID>.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.42.0/docs/resources/dialogflow_cx_security_settings#deidentify_template DialogflowCxSecuritySettings#deidentify_template}
	DeidentifyTemplate *string `field:"optional" json:"deidentifyTemplate" yaml:"deidentifyTemplate"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.42.0/docs/resources/dialogflow_cx_security_settings#id DialogflowCxSecuritySettings#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// insights_export_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.42.0/docs/resources/dialogflow_cx_security_settings#insights_export_settings DialogflowCxSecuritySettings#insights_export_settings}
	InsightsExportSettings *DialogflowCxSecuritySettingsInsightsExportSettings `field:"optional" json:"insightsExportSettings" yaml:"insightsExportSettings"`
	// [DLP](https://cloud.google.com/dlp/docs) inspect template name. Use this template to define inspect base settings. If empty, we use the default DLP inspect config. Note: inspectTemplate must be located in the same region as the SecuritySettings. Format: projects/<Project ID>/locations/<Location ID>/inspectTemplates/<Template ID> OR organizations/<Organization ID>/locations/<Location ID>/inspectTemplates/<Template ID>.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.42.0/docs/resources/dialogflow_cx_security_settings#inspect_template DialogflowCxSecuritySettings#inspect_template}
	InspectTemplate *string `field:"optional" json:"inspectTemplate" yaml:"inspectTemplate"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.42.0/docs/resources/dialogflow_cx_security_settings#project DialogflowCxSecuritySettings#project}.
	Project *string `field:"optional" json:"project" yaml:"project"`
	// List of types of data to remove when retention settings triggers purge. Possible values: ["DIALOGFLOW_HISTORY"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.42.0/docs/resources/dialogflow_cx_security_settings#purge_data_types DialogflowCxSecuritySettings#purge_data_types}
	PurgeDataTypes *[]*string `field:"optional" json:"purgeDataTypes" yaml:"purgeDataTypes"`
	// Defines what types of data to redact.
	//
	// If not set, defaults to not redacting any kind of data.
	// * REDACT_DISK_STORAGE: On data to be written to disk or similar devices that are capable of holding data even if power is disconnected. This includes data that are temporarily saved on disk. Possible values: ["REDACT_DISK_STORAGE"]
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.42.0/docs/resources/dialogflow_cx_security_settings#redaction_scope DialogflowCxSecuritySettings#redaction_scope}
	RedactionScope *string `field:"optional" json:"redactionScope" yaml:"redactionScope"`
	// Defines how we redact data.
	//
	// If not set, defaults to not redacting.
	// * REDACT_WITH_SERVICE: Call redaction service to clean up the data to be persisted. Possible values: ["REDACT_WITH_SERVICE"]
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.42.0/docs/resources/dialogflow_cx_security_settings#redaction_strategy DialogflowCxSecuritySettings#redaction_strategy}
	RedactionStrategy *string `field:"optional" json:"redactionStrategy" yaml:"redactionStrategy"`
	// Defines how long we retain persisted data that contains sensitive info.
	//
	// Only one of 'retention_window_days' and 'retention_strategy' may be set.
	// * REMOVE_AFTER_CONVERSATION: Removes data when the conversation ends. If there is no conversation explicitly established, a default conversation ends when the corresponding Dialogflow session ends. Possible values: ["REMOVE_AFTER_CONVERSATION"]
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.42.0/docs/resources/dialogflow_cx_security_settings#retention_strategy DialogflowCxSecuritySettings#retention_strategy}
	RetentionStrategy *string `field:"optional" json:"retentionStrategy" yaml:"retentionStrategy"`
	// Retains the data for the specified number of days.
	//
	// User must set a value lower than Dialogflow's default 365d TTL (30 days for Agent Assist traffic), higher value will be ignored and use default. Setting a value higher than that has no effect. A missing value or setting to 0 also means we use default TTL.
	// Only one of 'retention_window_days' and 'retention_strategy' may be set.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.42.0/docs/resources/dialogflow_cx_security_settings#retention_window_days DialogflowCxSecuritySettings#retention_window_days}
	RetentionWindowDays *float64 `field:"optional" json:"retentionWindowDays" yaml:"retentionWindowDays"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.42.0/docs/resources/dialogflow_cx_security_settings#timeouts DialogflowCxSecuritySettings#timeouts}
	Timeouts *DialogflowCxSecuritySettingsTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

