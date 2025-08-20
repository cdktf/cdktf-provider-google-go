// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dialogflowconversationprofile

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"@cdktf/provider-google.dialogflowConversationProfile.DialogflowConversationProfile",
		reflect.TypeOf((*DialogflowConversationProfile)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addMoveTarget", GoMethod: "AddMoveTarget"},
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberProperty{JsiiProperty: "automatedAgentConfig", GoGetter: "AutomatedAgentConfig"},
			_jsii_.MemberProperty{JsiiProperty: "automatedAgentConfigInput", GoGetter: "AutomatedAgentConfigInput"},
			_jsii_.MemberProperty{JsiiProperty: "cdktfStack", GoGetter: "CdktfStack"},
			_jsii_.MemberProperty{JsiiProperty: "connection", GoGetter: "Connection"},
			_jsii_.MemberProperty{JsiiProperty: "constructNodeMetadata", GoGetter: "ConstructNodeMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "count", GoGetter: "Count"},
			_jsii_.MemberProperty{JsiiProperty: "dependsOn", GoGetter: "DependsOn"},
			_jsii_.MemberProperty{JsiiProperty: "displayName", GoGetter: "DisplayName"},
			_jsii_.MemberProperty{JsiiProperty: "displayNameInput", GoGetter: "DisplayNameInput"},
			_jsii_.MemberProperty{JsiiProperty: "forEach", GoGetter: "ForEach"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberProperty{JsiiProperty: "friendlyUniqueId", GoGetter: "FriendlyUniqueId"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "hasResourceMove", GoMethod: "HasResourceMove"},
			_jsii_.MemberProperty{JsiiProperty: "humanAgentAssistantConfig", GoGetter: "HumanAgentAssistantConfig"},
			_jsii_.MemberProperty{JsiiProperty: "humanAgentAssistantConfigInput", GoGetter: "HumanAgentAssistantConfigInput"},
			_jsii_.MemberProperty{JsiiProperty: "humanAgentHandoffConfig", GoGetter: "HumanAgentHandoffConfig"},
			_jsii_.MemberProperty{JsiiProperty: "humanAgentHandoffConfigInput", GoGetter: "HumanAgentHandoffConfigInput"},
			_jsii_.MemberProperty{JsiiProperty: "id", GoGetter: "Id"},
			_jsii_.MemberProperty{JsiiProperty: "idInput", GoGetter: "IdInput"},
			_jsii_.MemberMethod{JsiiMethod: "importFrom", GoMethod: "ImportFrom"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "languageCode", GoGetter: "LanguageCode"},
			_jsii_.MemberProperty{JsiiProperty: "languageCodeInput", GoGetter: "LanguageCodeInput"},
			_jsii_.MemberProperty{JsiiProperty: "lifecycle", GoGetter: "Lifecycle"},
			_jsii_.MemberProperty{JsiiProperty: "location", GoGetter: "Location"},
			_jsii_.MemberProperty{JsiiProperty: "locationInput", GoGetter: "LocationInput"},
			_jsii_.MemberProperty{JsiiProperty: "loggingConfig", GoGetter: "LoggingConfig"},
			_jsii_.MemberProperty{JsiiProperty: "loggingConfigInput", GoGetter: "LoggingConfigInput"},
			_jsii_.MemberMethod{JsiiMethod: "moveFromId", GoMethod: "MoveFromId"},
			_jsii_.MemberMethod{JsiiMethod: "moveTo", GoMethod: "MoveTo"},
			_jsii_.MemberMethod{JsiiMethod: "moveToId", GoMethod: "MoveToId"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "newMessageEventNotificationConfig", GoGetter: "NewMessageEventNotificationConfig"},
			_jsii_.MemberProperty{JsiiProperty: "newMessageEventNotificationConfigInput", GoGetter: "NewMessageEventNotificationConfigInput"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "notificationConfig", GoGetter: "NotificationConfig"},
			_jsii_.MemberProperty{JsiiProperty: "notificationConfigInput", GoGetter: "NotificationConfigInput"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "project", GoGetter: "Project"},
			_jsii_.MemberProperty{JsiiProperty: "projectInput", GoGetter: "ProjectInput"},
			_jsii_.MemberProperty{JsiiProperty: "provider", GoGetter: "Provider"},
			_jsii_.MemberProperty{JsiiProperty: "provisioners", GoGetter: "Provisioners"},
			_jsii_.MemberMethod{JsiiMethod: "putAutomatedAgentConfig", GoMethod: "PutAutomatedAgentConfig"},
			_jsii_.MemberMethod{JsiiMethod: "putHumanAgentAssistantConfig", GoMethod: "PutHumanAgentAssistantConfig"},
			_jsii_.MemberMethod{JsiiMethod: "putHumanAgentHandoffConfig", GoMethod: "PutHumanAgentHandoffConfig"},
			_jsii_.MemberMethod{JsiiMethod: "putLoggingConfig", GoMethod: "PutLoggingConfig"},
			_jsii_.MemberMethod{JsiiMethod: "putNewMessageEventNotificationConfig", GoMethod: "PutNewMessageEventNotificationConfig"},
			_jsii_.MemberMethod{JsiiMethod: "putNotificationConfig", GoMethod: "PutNotificationConfig"},
			_jsii_.MemberMethod{JsiiMethod: "putSttConfig", GoMethod: "PutSttConfig"},
			_jsii_.MemberMethod{JsiiMethod: "putTimeouts", GoMethod: "PutTimeouts"},
			_jsii_.MemberMethod{JsiiMethod: "putTtsConfig", GoMethod: "PutTtsConfig"},
			_jsii_.MemberProperty{JsiiProperty: "rawOverrides", GoGetter: "RawOverrides"},
			_jsii_.MemberMethod{JsiiMethod: "resetAutomatedAgentConfig", GoMethod: "ResetAutomatedAgentConfig"},
			_jsii_.MemberMethod{JsiiMethod: "resetHumanAgentAssistantConfig", GoMethod: "ResetHumanAgentAssistantConfig"},
			_jsii_.MemberMethod{JsiiMethod: "resetHumanAgentHandoffConfig", GoMethod: "ResetHumanAgentHandoffConfig"},
			_jsii_.MemberMethod{JsiiMethod: "resetId", GoMethod: "ResetId"},
			_jsii_.MemberMethod{JsiiMethod: "resetLanguageCode", GoMethod: "ResetLanguageCode"},
			_jsii_.MemberMethod{JsiiMethod: "resetLoggingConfig", GoMethod: "ResetLoggingConfig"},
			_jsii_.MemberMethod{JsiiMethod: "resetNewMessageEventNotificationConfig", GoMethod: "ResetNewMessageEventNotificationConfig"},
			_jsii_.MemberMethod{JsiiMethod: "resetNotificationConfig", GoMethod: "ResetNotificationConfig"},
			_jsii_.MemberMethod{JsiiMethod: "resetOverrideLogicalId", GoMethod: "ResetOverrideLogicalId"},
			_jsii_.MemberMethod{JsiiMethod: "resetProject", GoMethod: "ResetProject"},
			_jsii_.MemberMethod{JsiiMethod: "resetSecuritySettings", GoMethod: "ResetSecuritySettings"},
			_jsii_.MemberMethod{JsiiMethod: "resetSttConfig", GoMethod: "ResetSttConfig"},
			_jsii_.MemberMethod{JsiiMethod: "resetTimeouts", GoMethod: "ResetTimeouts"},
			_jsii_.MemberMethod{JsiiMethod: "resetTimeZone", GoMethod: "ResetTimeZone"},
			_jsii_.MemberMethod{JsiiMethod: "resetTtsConfig", GoMethod: "ResetTtsConfig"},
			_jsii_.MemberProperty{JsiiProperty: "securitySettings", GoGetter: "SecuritySettings"},
			_jsii_.MemberProperty{JsiiProperty: "securitySettingsInput", GoGetter: "SecuritySettingsInput"},
			_jsii_.MemberProperty{JsiiProperty: "sttConfig", GoGetter: "SttConfig"},
			_jsii_.MemberProperty{JsiiProperty: "sttConfigInput", GoGetter: "SttConfigInput"},
			_jsii_.MemberMethod{JsiiMethod: "synthesizeAttributes", GoMethod: "SynthesizeAttributes"},
			_jsii_.MemberMethod{JsiiMethod: "synthesizeHclAttributes", GoMethod: "SynthesizeHclAttributes"},
			_jsii_.MemberProperty{JsiiProperty: "terraformGeneratorMetadata", GoGetter: "TerraformGeneratorMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "terraformMetaArguments", GoGetter: "TerraformMetaArguments"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResourceType", GoGetter: "TerraformResourceType"},
			_jsii_.MemberProperty{JsiiProperty: "timeouts", GoGetter: "Timeouts"},
			_jsii_.MemberProperty{JsiiProperty: "timeoutsInput", GoGetter: "TimeoutsInput"},
			_jsii_.MemberProperty{JsiiProperty: "timeZone", GoGetter: "TimeZone"},
			_jsii_.MemberProperty{JsiiProperty: "timeZoneInput", GoGetter: "TimeZoneInput"},
			_jsii_.MemberMethod{JsiiMethod: "toHclTerraform", GoMethod: "ToHclTerraform"},
			_jsii_.MemberMethod{JsiiMethod: "toMetadata", GoMethod: "ToMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberMethod{JsiiMethod: "toTerraform", GoMethod: "ToTerraform"},
			_jsii_.MemberProperty{JsiiProperty: "ttsConfig", GoGetter: "TtsConfig"},
			_jsii_.MemberProperty{JsiiProperty: "ttsConfigInput", GoGetter: "TtsConfigInput"},
		},
		func() interface{} {
			j := jsiiProxy_DialogflowConversationProfile{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfTerraformResource)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-google.dialogflowConversationProfile.DialogflowConversationProfileAutomatedAgentConfig",
		reflect.TypeOf((*DialogflowConversationProfileAutomatedAgentConfig)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-google.dialogflowConversationProfile.DialogflowConversationProfileAutomatedAgentConfigOutputReference",
		reflect.TypeOf((*DialogflowConversationProfileAutomatedAgentConfigOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "agent", GoGetter: "Agent"},
			_jsii_.MemberProperty{JsiiProperty: "agentInput", GoGetter: "AgentInput"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "resetSessionTtl", GoMethod: "ResetSessionTtl"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "sessionTtl", GoGetter: "SessionTtl"},
			_jsii_.MemberProperty{JsiiProperty: "sessionTtlInput", GoGetter: "SessionTtlInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_DialogflowConversationProfileAutomatedAgentConfigOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-google.dialogflowConversationProfile.DialogflowConversationProfileConfig",
		reflect.TypeOf((*DialogflowConversationProfileConfig)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-google.dialogflowConversationProfile.DialogflowConversationProfileHumanAgentAssistantConfig",
		reflect.TypeOf((*DialogflowConversationProfileHumanAgentAssistantConfig)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-google.dialogflowConversationProfile.DialogflowConversationProfileHumanAgentAssistantConfigEndUserSuggestionConfig",
		reflect.TypeOf((*DialogflowConversationProfileHumanAgentAssistantConfigEndUserSuggestionConfig)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-google.dialogflowConversationProfile.DialogflowConversationProfileHumanAgentAssistantConfigEndUserSuggestionConfigFeatureConfigs",
		reflect.TypeOf((*DialogflowConversationProfileHumanAgentAssistantConfigEndUserSuggestionConfigFeatureConfigs)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-google.dialogflowConversationProfile.DialogflowConversationProfileHumanAgentAssistantConfigEndUserSuggestionConfigFeatureConfigsConversationModelConfig",
		reflect.TypeOf((*DialogflowConversationProfileHumanAgentAssistantConfigEndUserSuggestionConfigFeatureConfigsConversationModelConfig)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-google.dialogflowConversationProfile.DialogflowConversationProfileHumanAgentAssistantConfigEndUserSuggestionConfigFeatureConfigsConversationModelConfigOutputReference",
		reflect.TypeOf((*DialogflowConversationProfileHumanAgentAssistantConfigEndUserSuggestionConfigFeatureConfigsConversationModelConfigOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "baselineModelVersion", GoGetter: "BaselineModelVersion"},
			_jsii_.MemberProperty{JsiiProperty: "baselineModelVersionInput", GoGetter: "BaselineModelVersionInput"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "model", GoGetter: "Model"},
			_jsii_.MemberProperty{JsiiProperty: "modelInput", GoGetter: "ModelInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetBaselineModelVersion", GoMethod: "ResetBaselineModelVersion"},
			_jsii_.MemberMethod{JsiiMethod: "resetModel", GoMethod: "ResetModel"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_DialogflowConversationProfileHumanAgentAssistantConfigEndUserSuggestionConfigFeatureConfigsConversationModelConfigOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-google.dialogflowConversationProfile.DialogflowConversationProfileHumanAgentAssistantConfigEndUserSuggestionConfigFeatureConfigsConversationProcessConfig",
		reflect.TypeOf((*DialogflowConversationProfileHumanAgentAssistantConfigEndUserSuggestionConfigFeatureConfigsConversationProcessConfig)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-google.dialogflowConversationProfile.DialogflowConversationProfileHumanAgentAssistantConfigEndUserSuggestionConfigFeatureConfigsConversationProcessConfigOutputReference",
		reflect.TypeOf((*DialogflowConversationProfileHumanAgentAssistantConfigEndUserSuggestionConfigFeatureConfigsConversationProcessConfigOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "recentSentencesCount", GoGetter: "RecentSentencesCount"},
			_jsii_.MemberProperty{JsiiProperty: "recentSentencesCountInput", GoGetter: "RecentSentencesCountInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetRecentSentencesCount", GoMethod: "ResetRecentSentencesCount"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_DialogflowConversationProfileHumanAgentAssistantConfigEndUserSuggestionConfigFeatureConfigsConversationProcessConfigOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-google.dialogflowConversationProfile.DialogflowConversationProfileHumanAgentAssistantConfigEndUserSuggestionConfigFeatureConfigsList",
		reflect.TypeOf((*DialogflowConversationProfileHumanAgentAssistantConfigEndUserSuggestionConfigFeatureConfigsList)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "allWithMapKey", GoMethod: "AllWithMapKey"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "get", GoMethod: "Get"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "wrapsSet", GoGetter: "WrapsSet"},
		},
		func() interface{} {
			j := jsiiProxy_DialogflowConversationProfileHumanAgentAssistantConfigEndUserSuggestionConfigFeatureConfigsList{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexList)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-google.dialogflowConversationProfile.DialogflowConversationProfileHumanAgentAssistantConfigEndUserSuggestionConfigFeatureConfigsOutputReference",
		reflect.TypeOf((*DialogflowConversationProfileHumanAgentAssistantConfigEndUserSuggestionConfigFeatureConfigsOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "conversationModelConfig", GoGetter: "ConversationModelConfig"},
			_jsii_.MemberProperty{JsiiProperty: "conversationModelConfigInput", GoGetter: "ConversationModelConfigInput"},
			_jsii_.MemberProperty{JsiiProperty: "conversationProcessConfig", GoGetter: "ConversationProcessConfig"},
			_jsii_.MemberProperty{JsiiProperty: "conversationProcessConfigInput", GoGetter: "ConversationProcessConfigInput"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "disableAgentQueryLogging", GoGetter: "DisableAgentQueryLogging"},
			_jsii_.MemberProperty{JsiiProperty: "disableAgentQueryLoggingInput", GoGetter: "DisableAgentQueryLoggingInput"},
			_jsii_.MemberProperty{JsiiProperty: "enableConversationAugmentedQuery", GoGetter: "EnableConversationAugmentedQuery"},
			_jsii_.MemberProperty{JsiiProperty: "enableConversationAugmentedQueryInput", GoGetter: "EnableConversationAugmentedQueryInput"},
			_jsii_.MemberProperty{JsiiProperty: "enableEventBasedSuggestion", GoGetter: "EnableEventBasedSuggestion"},
			_jsii_.MemberProperty{JsiiProperty: "enableEventBasedSuggestionInput", GoGetter: "EnableEventBasedSuggestionInput"},
			_jsii_.MemberProperty{JsiiProperty: "enableQuerySuggestionOnly", GoGetter: "EnableQuerySuggestionOnly"},
			_jsii_.MemberProperty{JsiiProperty: "enableQuerySuggestionOnlyInput", GoGetter: "EnableQuerySuggestionOnlyInput"},
			_jsii_.MemberProperty{JsiiProperty: "enableQuerySuggestionWhenNoAnswer", GoGetter: "EnableQuerySuggestionWhenNoAnswer"},
			_jsii_.MemberProperty{JsiiProperty: "enableQuerySuggestionWhenNoAnswerInput", GoGetter: "EnableQuerySuggestionWhenNoAnswerInput"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "putConversationModelConfig", GoMethod: "PutConversationModelConfig"},
			_jsii_.MemberMethod{JsiiMethod: "putConversationProcessConfig", GoMethod: "PutConversationProcessConfig"},
			_jsii_.MemberMethod{JsiiMethod: "putQueryConfig", GoMethod: "PutQueryConfig"},
			_jsii_.MemberMethod{JsiiMethod: "putSuggestionFeature", GoMethod: "PutSuggestionFeature"},
			_jsii_.MemberMethod{JsiiMethod: "putSuggestionTriggerSettings", GoMethod: "PutSuggestionTriggerSettings"},
			_jsii_.MemberProperty{JsiiProperty: "queryConfig", GoGetter: "QueryConfig"},
			_jsii_.MemberProperty{JsiiProperty: "queryConfigInput", GoGetter: "QueryConfigInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetConversationModelConfig", GoMethod: "ResetConversationModelConfig"},
			_jsii_.MemberMethod{JsiiMethod: "resetConversationProcessConfig", GoMethod: "ResetConversationProcessConfig"},
			_jsii_.MemberMethod{JsiiMethod: "resetDisableAgentQueryLogging", GoMethod: "ResetDisableAgentQueryLogging"},
			_jsii_.MemberMethod{JsiiMethod: "resetEnableConversationAugmentedQuery", GoMethod: "ResetEnableConversationAugmentedQuery"},
			_jsii_.MemberMethod{JsiiMethod: "resetEnableEventBasedSuggestion", GoMethod: "ResetEnableEventBasedSuggestion"},
			_jsii_.MemberMethod{JsiiMethod: "resetEnableQuerySuggestionOnly", GoMethod: "ResetEnableQuerySuggestionOnly"},
			_jsii_.MemberMethod{JsiiMethod: "resetEnableQuerySuggestionWhenNoAnswer", GoMethod: "ResetEnableQuerySuggestionWhenNoAnswer"},
			_jsii_.MemberMethod{JsiiMethod: "resetQueryConfig", GoMethod: "ResetQueryConfig"},
			_jsii_.MemberMethod{JsiiMethod: "resetSuggestionFeature", GoMethod: "ResetSuggestionFeature"},
			_jsii_.MemberMethod{JsiiMethod: "resetSuggestionTriggerSettings", GoMethod: "ResetSuggestionTriggerSettings"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "suggestionFeature", GoGetter: "SuggestionFeature"},
			_jsii_.MemberProperty{JsiiProperty: "suggestionFeatureInput", GoGetter: "SuggestionFeatureInput"},
			_jsii_.MemberProperty{JsiiProperty: "suggestionTriggerSettings", GoGetter: "SuggestionTriggerSettings"},
			_jsii_.MemberProperty{JsiiProperty: "suggestionTriggerSettingsInput", GoGetter: "SuggestionTriggerSettingsInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_DialogflowConversationProfileHumanAgentAssistantConfigEndUserSuggestionConfigFeatureConfigsOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-google.dialogflowConversationProfile.DialogflowConversationProfileHumanAgentAssistantConfigEndUserSuggestionConfigFeatureConfigsQueryConfig",
		reflect.TypeOf((*DialogflowConversationProfileHumanAgentAssistantConfigEndUserSuggestionConfigFeatureConfigsQueryConfig)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-google.dialogflowConversationProfile.DialogflowConversationProfileHumanAgentAssistantConfigEndUserSuggestionConfigFeatureConfigsQueryConfigContextFilterSettings",
		reflect.TypeOf((*DialogflowConversationProfileHumanAgentAssistantConfigEndUserSuggestionConfigFeatureConfigsQueryConfigContextFilterSettings)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-google.dialogflowConversationProfile.DialogflowConversationProfileHumanAgentAssistantConfigEndUserSuggestionConfigFeatureConfigsQueryConfigContextFilterSettingsOutputReference",
		reflect.TypeOf((*DialogflowConversationProfileHumanAgentAssistantConfigEndUserSuggestionConfigFeatureConfigsQueryConfigContextFilterSettingsOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "dropHandoffMessages", GoGetter: "DropHandoffMessages"},
			_jsii_.MemberProperty{JsiiProperty: "dropHandoffMessagesInput", GoGetter: "DropHandoffMessagesInput"},
			_jsii_.MemberProperty{JsiiProperty: "dropIvrMessages", GoGetter: "DropIvrMessages"},
			_jsii_.MemberProperty{JsiiProperty: "dropIvrMessagesInput", GoGetter: "DropIvrMessagesInput"},
			_jsii_.MemberProperty{JsiiProperty: "dropVirtualAgentMessages", GoGetter: "DropVirtualAgentMessages"},
			_jsii_.MemberProperty{JsiiProperty: "dropVirtualAgentMessagesInput", GoGetter: "DropVirtualAgentMessagesInput"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "resetDropHandoffMessages", GoMethod: "ResetDropHandoffMessages"},
			_jsii_.MemberMethod{JsiiMethod: "resetDropIvrMessages", GoMethod: "ResetDropIvrMessages"},
			_jsii_.MemberMethod{JsiiMethod: "resetDropVirtualAgentMessages", GoMethod: "ResetDropVirtualAgentMessages"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_DialogflowConversationProfileHumanAgentAssistantConfigEndUserSuggestionConfigFeatureConfigsQueryConfigContextFilterSettingsOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-google.dialogflowConversationProfile.DialogflowConversationProfileHumanAgentAssistantConfigEndUserSuggestionConfigFeatureConfigsQueryConfigDialogflowQuerySource",
		reflect.TypeOf((*DialogflowConversationProfileHumanAgentAssistantConfigEndUserSuggestionConfigFeatureConfigsQueryConfigDialogflowQuerySource)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-google.dialogflowConversationProfile.DialogflowConversationProfileHumanAgentAssistantConfigEndUserSuggestionConfigFeatureConfigsQueryConfigDialogflowQuerySourceHumanAgentSideConfig",
		reflect.TypeOf((*DialogflowConversationProfileHumanAgentAssistantConfigEndUserSuggestionConfigFeatureConfigsQueryConfigDialogflowQuerySourceHumanAgentSideConfig)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-google.dialogflowConversationProfile.DialogflowConversationProfileHumanAgentAssistantConfigEndUserSuggestionConfigFeatureConfigsQueryConfigDialogflowQuerySourceHumanAgentSideConfigOutputReference",
		reflect.TypeOf((*DialogflowConversationProfileHumanAgentAssistantConfigEndUserSuggestionConfigFeatureConfigsQueryConfigDialogflowQuerySourceHumanAgentSideConfigOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "agent", GoGetter: "Agent"},
			_jsii_.MemberProperty{JsiiProperty: "agentInput", GoGetter: "AgentInput"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "resetAgent", GoMethod: "ResetAgent"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_DialogflowConversationProfileHumanAgentAssistantConfigEndUserSuggestionConfigFeatureConfigsQueryConfigDialogflowQuerySourceHumanAgentSideConfigOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-google.dialogflowConversationProfile.DialogflowConversationProfileHumanAgentAssistantConfigEndUserSuggestionConfigFeatureConfigsQueryConfigDialogflowQuerySourceOutputReference",
		reflect.TypeOf((*DialogflowConversationProfileHumanAgentAssistantConfigEndUserSuggestionConfigFeatureConfigsQueryConfigDialogflowQuerySourceOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "agent", GoGetter: "Agent"},
			_jsii_.MemberProperty{JsiiProperty: "agentInput", GoGetter: "AgentInput"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "humanAgentSideConfig", GoGetter: "HumanAgentSideConfig"},
			_jsii_.MemberProperty{JsiiProperty: "humanAgentSideConfigInput", GoGetter: "HumanAgentSideConfigInput"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "putHumanAgentSideConfig", GoMethod: "PutHumanAgentSideConfig"},
			_jsii_.MemberMethod{JsiiMethod: "resetHumanAgentSideConfig", GoMethod: "ResetHumanAgentSideConfig"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_DialogflowConversationProfileHumanAgentAssistantConfigEndUserSuggestionConfigFeatureConfigsQueryConfigDialogflowQuerySourceOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-google.dialogflowConversationProfile.DialogflowConversationProfileHumanAgentAssistantConfigEndUserSuggestionConfigFeatureConfigsQueryConfigDocumentQuerySource",
		reflect.TypeOf((*DialogflowConversationProfileHumanAgentAssistantConfigEndUserSuggestionConfigFeatureConfigsQueryConfigDocumentQuerySource)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-google.dialogflowConversationProfile.DialogflowConversationProfileHumanAgentAssistantConfigEndUserSuggestionConfigFeatureConfigsQueryConfigDocumentQuerySourceOutputReference",
		reflect.TypeOf((*DialogflowConversationProfileHumanAgentAssistantConfigEndUserSuggestionConfigFeatureConfigsQueryConfigDocumentQuerySourceOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "documents", GoGetter: "Documents"},
			_jsii_.MemberProperty{JsiiProperty: "documentsInput", GoGetter: "DocumentsInput"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_DialogflowConversationProfileHumanAgentAssistantConfigEndUserSuggestionConfigFeatureConfigsQueryConfigDocumentQuerySourceOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-google.dialogflowConversationProfile.DialogflowConversationProfileHumanAgentAssistantConfigEndUserSuggestionConfigFeatureConfigsQueryConfigKnowledgeBaseQuerySource",
		reflect.TypeOf((*DialogflowConversationProfileHumanAgentAssistantConfigEndUserSuggestionConfigFeatureConfigsQueryConfigKnowledgeBaseQuerySource)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-google.dialogflowConversationProfile.DialogflowConversationProfileHumanAgentAssistantConfigEndUserSuggestionConfigFeatureConfigsQueryConfigKnowledgeBaseQuerySourceOutputReference",
		reflect.TypeOf((*DialogflowConversationProfileHumanAgentAssistantConfigEndUserSuggestionConfigFeatureConfigsQueryConfigKnowledgeBaseQuerySourceOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "knowledgeBases", GoGetter: "KnowledgeBases"},
			_jsii_.MemberProperty{JsiiProperty: "knowledgeBasesInput", GoGetter: "KnowledgeBasesInput"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_DialogflowConversationProfileHumanAgentAssistantConfigEndUserSuggestionConfigFeatureConfigsQueryConfigKnowledgeBaseQuerySourceOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-google.dialogflowConversationProfile.DialogflowConversationProfileHumanAgentAssistantConfigEndUserSuggestionConfigFeatureConfigsQueryConfigOutputReference",
		reflect.TypeOf((*DialogflowConversationProfileHumanAgentAssistantConfigEndUserSuggestionConfigFeatureConfigsQueryConfigOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "confidenceThreshold", GoGetter: "ConfidenceThreshold"},
			_jsii_.MemberProperty{JsiiProperty: "confidenceThresholdInput", GoGetter: "ConfidenceThresholdInput"},
			_jsii_.MemberProperty{JsiiProperty: "contextFilterSettings", GoGetter: "ContextFilterSettings"},
			_jsii_.MemberProperty{JsiiProperty: "contextFilterSettingsInput", GoGetter: "ContextFilterSettingsInput"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "dialogflowQuerySource", GoGetter: "DialogflowQuerySource"},
			_jsii_.MemberProperty{JsiiProperty: "dialogflowQuerySourceInput", GoGetter: "DialogflowQuerySourceInput"},
			_jsii_.MemberProperty{JsiiProperty: "documentQuerySource", GoGetter: "DocumentQuerySource"},
			_jsii_.MemberProperty{JsiiProperty: "documentQuerySourceInput", GoGetter: "DocumentQuerySourceInput"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "knowledgeBaseQuerySource", GoGetter: "KnowledgeBaseQuerySource"},
			_jsii_.MemberProperty{JsiiProperty: "knowledgeBaseQuerySourceInput", GoGetter: "KnowledgeBaseQuerySourceInput"},
			_jsii_.MemberProperty{JsiiProperty: "maxResults", GoGetter: "MaxResults"},
			_jsii_.MemberProperty{JsiiProperty: "maxResultsInput", GoGetter: "MaxResultsInput"},
			_jsii_.MemberMethod{JsiiMethod: "putContextFilterSettings", GoMethod: "PutContextFilterSettings"},
			_jsii_.MemberMethod{JsiiMethod: "putDialogflowQuerySource", GoMethod: "PutDialogflowQuerySource"},
			_jsii_.MemberMethod{JsiiMethod: "putDocumentQuerySource", GoMethod: "PutDocumentQuerySource"},
			_jsii_.MemberMethod{JsiiMethod: "putKnowledgeBaseQuerySource", GoMethod: "PutKnowledgeBaseQuerySource"},
			_jsii_.MemberMethod{JsiiMethod: "putSections", GoMethod: "PutSections"},
			_jsii_.MemberMethod{JsiiMethod: "resetConfidenceThreshold", GoMethod: "ResetConfidenceThreshold"},
			_jsii_.MemberMethod{JsiiMethod: "resetContextFilterSettings", GoMethod: "ResetContextFilterSettings"},
			_jsii_.MemberMethod{JsiiMethod: "resetDialogflowQuerySource", GoMethod: "ResetDialogflowQuerySource"},
			_jsii_.MemberMethod{JsiiMethod: "resetDocumentQuerySource", GoMethod: "ResetDocumentQuerySource"},
			_jsii_.MemberMethod{JsiiMethod: "resetKnowledgeBaseQuerySource", GoMethod: "ResetKnowledgeBaseQuerySource"},
			_jsii_.MemberMethod{JsiiMethod: "resetMaxResults", GoMethod: "ResetMaxResults"},
			_jsii_.MemberMethod{JsiiMethod: "resetSections", GoMethod: "ResetSections"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "sections", GoGetter: "Sections"},
			_jsii_.MemberProperty{JsiiProperty: "sectionsInput", GoGetter: "SectionsInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_DialogflowConversationProfileHumanAgentAssistantConfigEndUserSuggestionConfigFeatureConfigsQueryConfigOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-google.dialogflowConversationProfile.DialogflowConversationProfileHumanAgentAssistantConfigEndUserSuggestionConfigFeatureConfigsQueryConfigSections",
		reflect.TypeOf((*DialogflowConversationProfileHumanAgentAssistantConfigEndUserSuggestionConfigFeatureConfigsQueryConfigSections)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-google.dialogflowConversationProfile.DialogflowConversationProfileHumanAgentAssistantConfigEndUserSuggestionConfigFeatureConfigsQueryConfigSectionsOutputReference",
		reflect.TypeOf((*DialogflowConversationProfileHumanAgentAssistantConfigEndUserSuggestionConfigFeatureConfigsQueryConfigSectionsOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "resetSectionTypes", GoMethod: "ResetSectionTypes"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "sectionTypes", GoGetter: "SectionTypes"},
			_jsii_.MemberProperty{JsiiProperty: "sectionTypesInput", GoGetter: "SectionTypesInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_DialogflowConversationProfileHumanAgentAssistantConfigEndUserSuggestionConfigFeatureConfigsQueryConfigSectionsOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-google.dialogflowConversationProfile.DialogflowConversationProfileHumanAgentAssistantConfigEndUserSuggestionConfigFeatureConfigsSuggestionFeature",
		reflect.TypeOf((*DialogflowConversationProfileHumanAgentAssistantConfigEndUserSuggestionConfigFeatureConfigsSuggestionFeature)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-google.dialogflowConversationProfile.DialogflowConversationProfileHumanAgentAssistantConfigEndUserSuggestionConfigFeatureConfigsSuggestionFeatureOutputReference",
		reflect.TypeOf((*DialogflowConversationProfileHumanAgentAssistantConfigEndUserSuggestionConfigFeatureConfigsSuggestionFeatureOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "resetType", GoMethod: "ResetType"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "type", GoGetter: "Type"},
			_jsii_.MemberProperty{JsiiProperty: "typeInput", GoGetter: "TypeInput"},
		},
		func() interface{} {
			j := jsiiProxy_DialogflowConversationProfileHumanAgentAssistantConfigEndUserSuggestionConfigFeatureConfigsSuggestionFeatureOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-google.dialogflowConversationProfile.DialogflowConversationProfileHumanAgentAssistantConfigEndUserSuggestionConfigFeatureConfigsSuggestionTriggerSettings",
		reflect.TypeOf((*DialogflowConversationProfileHumanAgentAssistantConfigEndUserSuggestionConfigFeatureConfigsSuggestionTriggerSettings)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-google.dialogflowConversationProfile.DialogflowConversationProfileHumanAgentAssistantConfigEndUserSuggestionConfigFeatureConfigsSuggestionTriggerSettingsOutputReference",
		reflect.TypeOf((*DialogflowConversationProfileHumanAgentAssistantConfigEndUserSuggestionConfigFeatureConfigsSuggestionTriggerSettingsOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "noSmallTalk", GoGetter: "NoSmallTalk"},
			_jsii_.MemberProperty{JsiiProperty: "noSmallTalkInput", GoGetter: "NoSmallTalkInput"},
			_jsii_.MemberProperty{JsiiProperty: "onlyEndUser", GoGetter: "OnlyEndUser"},
			_jsii_.MemberProperty{JsiiProperty: "onlyEndUserInput", GoGetter: "OnlyEndUserInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetNoSmallTalk", GoMethod: "ResetNoSmallTalk"},
			_jsii_.MemberMethod{JsiiMethod: "resetOnlyEndUser", GoMethod: "ResetOnlyEndUser"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_DialogflowConversationProfileHumanAgentAssistantConfigEndUserSuggestionConfigFeatureConfigsSuggestionTriggerSettingsOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-google.dialogflowConversationProfile.DialogflowConversationProfileHumanAgentAssistantConfigEndUserSuggestionConfigOutputReference",
		reflect.TypeOf((*DialogflowConversationProfileHumanAgentAssistantConfigEndUserSuggestionConfigOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "disableHighLatencyFeaturesSyncDelivery", GoGetter: "DisableHighLatencyFeaturesSyncDelivery"},
			_jsii_.MemberProperty{JsiiProperty: "disableHighLatencyFeaturesSyncDeliveryInput", GoGetter: "DisableHighLatencyFeaturesSyncDeliveryInput"},
			_jsii_.MemberProperty{JsiiProperty: "featureConfigs", GoGetter: "FeatureConfigs"},
			_jsii_.MemberProperty{JsiiProperty: "featureConfigsInput", GoGetter: "FeatureConfigsInput"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberProperty{JsiiProperty: "generators", GoGetter: "Generators"},
			_jsii_.MemberProperty{JsiiProperty: "generatorsInput", GoGetter: "GeneratorsInput"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "groupSuggestionResponses", GoGetter: "GroupSuggestionResponses"},
			_jsii_.MemberProperty{JsiiProperty: "groupSuggestionResponsesInput", GoGetter: "GroupSuggestionResponsesInput"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "putFeatureConfigs", GoMethod: "PutFeatureConfigs"},
			_jsii_.MemberMethod{JsiiMethod: "resetDisableHighLatencyFeaturesSyncDelivery", GoMethod: "ResetDisableHighLatencyFeaturesSyncDelivery"},
			_jsii_.MemberMethod{JsiiMethod: "resetFeatureConfigs", GoMethod: "ResetFeatureConfigs"},
			_jsii_.MemberMethod{JsiiMethod: "resetGenerators", GoMethod: "ResetGenerators"},
			_jsii_.MemberMethod{JsiiMethod: "resetGroupSuggestionResponses", GoMethod: "ResetGroupSuggestionResponses"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_DialogflowConversationProfileHumanAgentAssistantConfigEndUserSuggestionConfigOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-google.dialogflowConversationProfile.DialogflowConversationProfileHumanAgentAssistantConfigHumanAgentSuggestionConfig",
		reflect.TypeOf((*DialogflowConversationProfileHumanAgentAssistantConfigHumanAgentSuggestionConfig)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-google.dialogflowConversationProfile.DialogflowConversationProfileHumanAgentAssistantConfigHumanAgentSuggestionConfigFeatureConfigs",
		reflect.TypeOf((*DialogflowConversationProfileHumanAgentAssistantConfigHumanAgentSuggestionConfigFeatureConfigs)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-google.dialogflowConversationProfile.DialogflowConversationProfileHumanAgentAssistantConfigHumanAgentSuggestionConfigFeatureConfigsConversationModelConfig",
		reflect.TypeOf((*DialogflowConversationProfileHumanAgentAssistantConfigHumanAgentSuggestionConfigFeatureConfigsConversationModelConfig)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-google.dialogflowConversationProfile.DialogflowConversationProfileHumanAgentAssistantConfigHumanAgentSuggestionConfigFeatureConfigsConversationModelConfigOutputReference",
		reflect.TypeOf((*DialogflowConversationProfileHumanAgentAssistantConfigHumanAgentSuggestionConfigFeatureConfigsConversationModelConfigOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "baselineModelVersion", GoGetter: "BaselineModelVersion"},
			_jsii_.MemberProperty{JsiiProperty: "baselineModelVersionInput", GoGetter: "BaselineModelVersionInput"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "model", GoGetter: "Model"},
			_jsii_.MemberProperty{JsiiProperty: "modelInput", GoGetter: "ModelInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetBaselineModelVersion", GoMethod: "ResetBaselineModelVersion"},
			_jsii_.MemberMethod{JsiiMethod: "resetModel", GoMethod: "ResetModel"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_DialogflowConversationProfileHumanAgentAssistantConfigHumanAgentSuggestionConfigFeatureConfigsConversationModelConfigOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-google.dialogflowConversationProfile.DialogflowConversationProfileHumanAgentAssistantConfigHumanAgentSuggestionConfigFeatureConfigsConversationProcessConfig",
		reflect.TypeOf((*DialogflowConversationProfileHumanAgentAssistantConfigHumanAgentSuggestionConfigFeatureConfigsConversationProcessConfig)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-google.dialogflowConversationProfile.DialogflowConversationProfileHumanAgentAssistantConfigHumanAgentSuggestionConfigFeatureConfigsConversationProcessConfigOutputReference",
		reflect.TypeOf((*DialogflowConversationProfileHumanAgentAssistantConfigHumanAgentSuggestionConfigFeatureConfigsConversationProcessConfigOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "recentSentencesCount", GoGetter: "RecentSentencesCount"},
			_jsii_.MemberProperty{JsiiProperty: "recentSentencesCountInput", GoGetter: "RecentSentencesCountInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetRecentSentencesCount", GoMethod: "ResetRecentSentencesCount"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_DialogflowConversationProfileHumanAgentAssistantConfigHumanAgentSuggestionConfigFeatureConfigsConversationProcessConfigOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-google.dialogflowConversationProfile.DialogflowConversationProfileHumanAgentAssistantConfigHumanAgentSuggestionConfigFeatureConfigsList",
		reflect.TypeOf((*DialogflowConversationProfileHumanAgentAssistantConfigHumanAgentSuggestionConfigFeatureConfigsList)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "allWithMapKey", GoMethod: "AllWithMapKey"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "get", GoMethod: "Get"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "wrapsSet", GoGetter: "WrapsSet"},
		},
		func() interface{} {
			j := jsiiProxy_DialogflowConversationProfileHumanAgentAssistantConfigHumanAgentSuggestionConfigFeatureConfigsList{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexList)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-google.dialogflowConversationProfile.DialogflowConversationProfileHumanAgentAssistantConfigHumanAgentSuggestionConfigFeatureConfigsOutputReference",
		reflect.TypeOf((*DialogflowConversationProfileHumanAgentAssistantConfigHumanAgentSuggestionConfigFeatureConfigsOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "conversationModelConfig", GoGetter: "ConversationModelConfig"},
			_jsii_.MemberProperty{JsiiProperty: "conversationModelConfigInput", GoGetter: "ConversationModelConfigInput"},
			_jsii_.MemberProperty{JsiiProperty: "conversationProcessConfig", GoGetter: "ConversationProcessConfig"},
			_jsii_.MemberProperty{JsiiProperty: "conversationProcessConfigInput", GoGetter: "ConversationProcessConfigInput"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "disableAgentQueryLogging", GoGetter: "DisableAgentQueryLogging"},
			_jsii_.MemberProperty{JsiiProperty: "disableAgentQueryLoggingInput", GoGetter: "DisableAgentQueryLoggingInput"},
			_jsii_.MemberProperty{JsiiProperty: "enableConversationAugmentedQuery", GoGetter: "EnableConversationAugmentedQuery"},
			_jsii_.MemberProperty{JsiiProperty: "enableConversationAugmentedQueryInput", GoGetter: "EnableConversationAugmentedQueryInput"},
			_jsii_.MemberProperty{JsiiProperty: "enableEventBasedSuggestion", GoGetter: "EnableEventBasedSuggestion"},
			_jsii_.MemberProperty{JsiiProperty: "enableEventBasedSuggestionInput", GoGetter: "EnableEventBasedSuggestionInput"},
			_jsii_.MemberProperty{JsiiProperty: "enableQuerySuggestionOnly", GoGetter: "EnableQuerySuggestionOnly"},
			_jsii_.MemberProperty{JsiiProperty: "enableQuerySuggestionOnlyInput", GoGetter: "EnableQuerySuggestionOnlyInput"},
			_jsii_.MemberProperty{JsiiProperty: "enableQuerySuggestionWhenNoAnswer", GoGetter: "EnableQuerySuggestionWhenNoAnswer"},
			_jsii_.MemberProperty{JsiiProperty: "enableQuerySuggestionWhenNoAnswerInput", GoGetter: "EnableQuerySuggestionWhenNoAnswerInput"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "putConversationModelConfig", GoMethod: "PutConversationModelConfig"},
			_jsii_.MemberMethod{JsiiMethod: "putConversationProcessConfig", GoMethod: "PutConversationProcessConfig"},
			_jsii_.MemberMethod{JsiiMethod: "putQueryConfig", GoMethod: "PutQueryConfig"},
			_jsii_.MemberMethod{JsiiMethod: "putSuggestionFeature", GoMethod: "PutSuggestionFeature"},
			_jsii_.MemberMethod{JsiiMethod: "putSuggestionTriggerSettings", GoMethod: "PutSuggestionTriggerSettings"},
			_jsii_.MemberProperty{JsiiProperty: "queryConfig", GoGetter: "QueryConfig"},
			_jsii_.MemberProperty{JsiiProperty: "queryConfigInput", GoGetter: "QueryConfigInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetConversationModelConfig", GoMethod: "ResetConversationModelConfig"},
			_jsii_.MemberMethod{JsiiMethod: "resetConversationProcessConfig", GoMethod: "ResetConversationProcessConfig"},
			_jsii_.MemberMethod{JsiiMethod: "resetDisableAgentQueryLogging", GoMethod: "ResetDisableAgentQueryLogging"},
			_jsii_.MemberMethod{JsiiMethod: "resetEnableConversationAugmentedQuery", GoMethod: "ResetEnableConversationAugmentedQuery"},
			_jsii_.MemberMethod{JsiiMethod: "resetEnableEventBasedSuggestion", GoMethod: "ResetEnableEventBasedSuggestion"},
			_jsii_.MemberMethod{JsiiMethod: "resetEnableQuerySuggestionOnly", GoMethod: "ResetEnableQuerySuggestionOnly"},
			_jsii_.MemberMethod{JsiiMethod: "resetEnableQuerySuggestionWhenNoAnswer", GoMethod: "ResetEnableQuerySuggestionWhenNoAnswer"},
			_jsii_.MemberMethod{JsiiMethod: "resetQueryConfig", GoMethod: "ResetQueryConfig"},
			_jsii_.MemberMethod{JsiiMethod: "resetSuggestionFeature", GoMethod: "ResetSuggestionFeature"},
			_jsii_.MemberMethod{JsiiMethod: "resetSuggestionTriggerSettings", GoMethod: "ResetSuggestionTriggerSettings"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "suggestionFeature", GoGetter: "SuggestionFeature"},
			_jsii_.MemberProperty{JsiiProperty: "suggestionFeatureInput", GoGetter: "SuggestionFeatureInput"},
			_jsii_.MemberProperty{JsiiProperty: "suggestionTriggerSettings", GoGetter: "SuggestionTriggerSettings"},
			_jsii_.MemberProperty{JsiiProperty: "suggestionTriggerSettingsInput", GoGetter: "SuggestionTriggerSettingsInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_DialogflowConversationProfileHumanAgentAssistantConfigHumanAgentSuggestionConfigFeatureConfigsOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-google.dialogflowConversationProfile.DialogflowConversationProfileHumanAgentAssistantConfigHumanAgentSuggestionConfigFeatureConfigsQueryConfig",
		reflect.TypeOf((*DialogflowConversationProfileHumanAgentAssistantConfigHumanAgentSuggestionConfigFeatureConfigsQueryConfig)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-google.dialogflowConversationProfile.DialogflowConversationProfileHumanAgentAssistantConfigHumanAgentSuggestionConfigFeatureConfigsQueryConfigContextFilterSettings",
		reflect.TypeOf((*DialogflowConversationProfileHumanAgentAssistantConfigHumanAgentSuggestionConfigFeatureConfigsQueryConfigContextFilterSettings)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-google.dialogflowConversationProfile.DialogflowConversationProfileHumanAgentAssistantConfigHumanAgentSuggestionConfigFeatureConfigsQueryConfigContextFilterSettingsOutputReference",
		reflect.TypeOf((*DialogflowConversationProfileHumanAgentAssistantConfigHumanAgentSuggestionConfigFeatureConfigsQueryConfigContextFilterSettingsOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "dropHandoffMessages", GoGetter: "DropHandoffMessages"},
			_jsii_.MemberProperty{JsiiProperty: "dropHandoffMessagesInput", GoGetter: "DropHandoffMessagesInput"},
			_jsii_.MemberProperty{JsiiProperty: "dropIvrMessages", GoGetter: "DropIvrMessages"},
			_jsii_.MemberProperty{JsiiProperty: "dropIvrMessagesInput", GoGetter: "DropIvrMessagesInput"},
			_jsii_.MemberProperty{JsiiProperty: "dropVirtualAgentMessages", GoGetter: "DropVirtualAgentMessages"},
			_jsii_.MemberProperty{JsiiProperty: "dropVirtualAgentMessagesInput", GoGetter: "DropVirtualAgentMessagesInput"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "resetDropHandoffMessages", GoMethod: "ResetDropHandoffMessages"},
			_jsii_.MemberMethod{JsiiMethod: "resetDropIvrMessages", GoMethod: "ResetDropIvrMessages"},
			_jsii_.MemberMethod{JsiiMethod: "resetDropVirtualAgentMessages", GoMethod: "ResetDropVirtualAgentMessages"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_DialogflowConversationProfileHumanAgentAssistantConfigHumanAgentSuggestionConfigFeatureConfigsQueryConfigContextFilterSettingsOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-google.dialogflowConversationProfile.DialogflowConversationProfileHumanAgentAssistantConfigHumanAgentSuggestionConfigFeatureConfigsQueryConfigDialogflowQuerySource",
		reflect.TypeOf((*DialogflowConversationProfileHumanAgentAssistantConfigHumanAgentSuggestionConfigFeatureConfigsQueryConfigDialogflowQuerySource)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-google.dialogflowConversationProfile.DialogflowConversationProfileHumanAgentAssistantConfigHumanAgentSuggestionConfigFeatureConfigsQueryConfigDialogflowQuerySourceHumanAgentSideConfig",
		reflect.TypeOf((*DialogflowConversationProfileHumanAgentAssistantConfigHumanAgentSuggestionConfigFeatureConfigsQueryConfigDialogflowQuerySourceHumanAgentSideConfig)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-google.dialogflowConversationProfile.DialogflowConversationProfileHumanAgentAssistantConfigHumanAgentSuggestionConfigFeatureConfigsQueryConfigDialogflowQuerySourceHumanAgentSideConfigOutputReference",
		reflect.TypeOf((*DialogflowConversationProfileHumanAgentAssistantConfigHumanAgentSuggestionConfigFeatureConfigsQueryConfigDialogflowQuerySourceHumanAgentSideConfigOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "agent", GoGetter: "Agent"},
			_jsii_.MemberProperty{JsiiProperty: "agentInput", GoGetter: "AgentInput"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "resetAgent", GoMethod: "ResetAgent"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_DialogflowConversationProfileHumanAgentAssistantConfigHumanAgentSuggestionConfigFeatureConfigsQueryConfigDialogflowQuerySourceHumanAgentSideConfigOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-google.dialogflowConversationProfile.DialogflowConversationProfileHumanAgentAssistantConfigHumanAgentSuggestionConfigFeatureConfigsQueryConfigDialogflowQuerySourceOutputReference",
		reflect.TypeOf((*DialogflowConversationProfileHumanAgentAssistantConfigHumanAgentSuggestionConfigFeatureConfigsQueryConfigDialogflowQuerySourceOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "agent", GoGetter: "Agent"},
			_jsii_.MemberProperty{JsiiProperty: "agentInput", GoGetter: "AgentInput"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "humanAgentSideConfig", GoGetter: "HumanAgentSideConfig"},
			_jsii_.MemberProperty{JsiiProperty: "humanAgentSideConfigInput", GoGetter: "HumanAgentSideConfigInput"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "putHumanAgentSideConfig", GoMethod: "PutHumanAgentSideConfig"},
			_jsii_.MemberMethod{JsiiMethod: "resetHumanAgentSideConfig", GoMethod: "ResetHumanAgentSideConfig"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_DialogflowConversationProfileHumanAgentAssistantConfigHumanAgentSuggestionConfigFeatureConfigsQueryConfigDialogflowQuerySourceOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-google.dialogflowConversationProfile.DialogflowConversationProfileHumanAgentAssistantConfigHumanAgentSuggestionConfigFeatureConfigsQueryConfigOutputReference",
		reflect.TypeOf((*DialogflowConversationProfileHumanAgentAssistantConfigHumanAgentSuggestionConfigFeatureConfigsQueryConfigOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "confidenceThreshold", GoGetter: "ConfidenceThreshold"},
			_jsii_.MemberProperty{JsiiProperty: "confidenceThresholdInput", GoGetter: "ConfidenceThresholdInput"},
			_jsii_.MemberProperty{JsiiProperty: "contextFilterSettings", GoGetter: "ContextFilterSettings"},
			_jsii_.MemberProperty{JsiiProperty: "contextFilterSettingsInput", GoGetter: "ContextFilterSettingsInput"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "dialogflowQuerySource", GoGetter: "DialogflowQuerySource"},
			_jsii_.MemberProperty{JsiiProperty: "dialogflowQuerySourceInput", GoGetter: "DialogflowQuerySourceInput"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "maxResults", GoGetter: "MaxResults"},
			_jsii_.MemberProperty{JsiiProperty: "maxResultsInput", GoGetter: "MaxResultsInput"},
			_jsii_.MemberMethod{JsiiMethod: "putContextFilterSettings", GoMethod: "PutContextFilterSettings"},
			_jsii_.MemberMethod{JsiiMethod: "putDialogflowQuerySource", GoMethod: "PutDialogflowQuerySource"},
			_jsii_.MemberMethod{JsiiMethod: "putSections", GoMethod: "PutSections"},
			_jsii_.MemberMethod{JsiiMethod: "resetConfidenceThreshold", GoMethod: "ResetConfidenceThreshold"},
			_jsii_.MemberMethod{JsiiMethod: "resetContextFilterSettings", GoMethod: "ResetContextFilterSettings"},
			_jsii_.MemberMethod{JsiiMethod: "resetDialogflowQuerySource", GoMethod: "ResetDialogflowQuerySource"},
			_jsii_.MemberMethod{JsiiMethod: "resetMaxResults", GoMethod: "ResetMaxResults"},
			_jsii_.MemberMethod{JsiiMethod: "resetSections", GoMethod: "ResetSections"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "sections", GoGetter: "Sections"},
			_jsii_.MemberProperty{JsiiProperty: "sectionsInput", GoGetter: "SectionsInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_DialogflowConversationProfileHumanAgentAssistantConfigHumanAgentSuggestionConfigFeatureConfigsQueryConfigOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-google.dialogflowConversationProfile.DialogflowConversationProfileHumanAgentAssistantConfigHumanAgentSuggestionConfigFeatureConfigsQueryConfigSections",
		reflect.TypeOf((*DialogflowConversationProfileHumanAgentAssistantConfigHumanAgentSuggestionConfigFeatureConfigsQueryConfigSections)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-google.dialogflowConversationProfile.DialogflowConversationProfileHumanAgentAssistantConfigHumanAgentSuggestionConfigFeatureConfigsQueryConfigSectionsOutputReference",
		reflect.TypeOf((*DialogflowConversationProfileHumanAgentAssistantConfigHumanAgentSuggestionConfigFeatureConfigsQueryConfigSectionsOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "resetSectionTypes", GoMethod: "ResetSectionTypes"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "sectionTypes", GoGetter: "SectionTypes"},
			_jsii_.MemberProperty{JsiiProperty: "sectionTypesInput", GoGetter: "SectionTypesInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_DialogflowConversationProfileHumanAgentAssistantConfigHumanAgentSuggestionConfigFeatureConfigsQueryConfigSectionsOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-google.dialogflowConversationProfile.DialogflowConversationProfileHumanAgentAssistantConfigHumanAgentSuggestionConfigFeatureConfigsSuggestionFeature",
		reflect.TypeOf((*DialogflowConversationProfileHumanAgentAssistantConfigHumanAgentSuggestionConfigFeatureConfigsSuggestionFeature)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-google.dialogflowConversationProfile.DialogflowConversationProfileHumanAgentAssistantConfigHumanAgentSuggestionConfigFeatureConfigsSuggestionFeatureOutputReference",
		reflect.TypeOf((*DialogflowConversationProfileHumanAgentAssistantConfigHumanAgentSuggestionConfigFeatureConfigsSuggestionFeatureOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "resetType", GoMethod: "ResetType"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "type", GoGetter: "Type"},
			_jsii_.MemberProperty{JsiiProperty: "typeInput", GoGetter: "TypeInput"},
		},
		func() interface{} {
			j := jsiiProxy_DialogflowConversationProfileHumanAgentAssistantConfigHumanAgentSuggestionConfigFeatureConfigsSuggestionFeatureOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-google.dialogflowConversationProfile.DialogflowConversationProfileHumanAgentAssistantConfigHumanAgentSuggestionConfigFeatureConfigsSuggestionTriggerSettings",
		reflect.TypeOf((*DialogflowConversationProfileHumanAgentAssistantConfigHumanAgentSuggestionConfigFeatureConfigsSuggestionTriggerSettings)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-google.dialogflowConversationProfile.DialogflowConversationProfileHumanAgentAssistantConfigHumanAgentSuggestionConfigFeatureConfigsSuggestionTriggerSettingsOutputReference",
		reflect.TypeOf((*DialogflowConversationProfileHumanAgentAssistantConfigHumanAgentSuggestionConfigFeatureConfigsSuggestionTriggerSettingsOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "noSmallTalk", GoGetter: "NoSmallTalk"},
			_jsii_.MemberProperty{JsiiProperty: "noSmallTalkInput", GoGetter: "NoSmallTalkInput"},
			_jsii_.MemberProperty{JsiiProperty: "onlyEndUser", GoGetter: "OnlyEndUser"},
			_jsii_.MemberProperty{JsiiProperty: "onlyEndUserInput", GoGetter: "OnlyEndUserInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetNoSmallTalk", GoMethod: "ResetNoSmallTalk"},
			_jsii_.MemberMethod{JsiiMethod: "resetOnlyEndUser", GoMethod: "ResetOnlyEndUser"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_DialogflowConversationProfileHumanAgentAssistantConfigHumanAgentSuggestionConfigFeatureConfigsSuggestionTriggerSettingsOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-google.dialogflowConversationProfile.DialogflowConversationProfileHumanAgentAssistantConfigHumanAgentSuggestionConfigOutputReference",
		reflect.TypeOf((*DialogflowConversationProfileHumanAgentAssistantConfigHumanAgentSuggestionConfigOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "disableHighLatencyFeaturesSyncDelivery", GoGetter: "DisableHighLatencyFeaturesSyncDelivery"},
			_jsii_.MemberProperty{JsiiProperty: "disableHighLatencyFeaturesSyncDeliveryInput", GoGetter: "DisableHighLatencyFeaturesSyncDeliveryInput"},
			_jsii_.MemberProperty{JsiiProperty: "featureConfigs", GoGetter: "FeatureConfigs"},
			_jsii_.MemberProperty{JsiiProperty: "featureConfigsInput", GoGetter: "FeatureConfigsInput"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberProperty{JsiiProperty: "generators", GoGetter: "Generators"},
			_jsii_.MemberProperty{JsiiProperty: "generatorsInput", GoGetter: "GeneratorsInput"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "groupSuggestionResponses", GoGetter: "GroupSuggestionResponses"},
			_jsii_.MemberProperty{JsiiProperty: "groupSuggestionResponsesInput", GoGetter: "GroupSuggestionResponsesInput"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "putFeatureConfigs", GoMethod: "PutFeatureConfigs"},
			_jsii_.MemberMethod{JsiiMethod: "resetDisableHighLatencyFeaturesSyncDelivery", GoMethod: "ResetDisableHighLatencyFeaturesSyncDelivery"},
			_jsii_.MemberMethod{JsiiMethod: "resetFeatureConfigs", GoMethod: "ResetFeatureConfigs"},
			_jsii_.MemberMethod{JsiiMethod: "resetGenerators", GoMethod: "ResetGenerators"},
			_jsii_.MemberMethod{JsiiMethod: "resetGroupSuggestionResponses", GoMethod: "ResetGroupSuggestionResponses"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_DialogflowConversationProfileHumanAgentAssistantConfigHumanAgentSuggestionConfigOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-google.dialogflowConversationProfile.DialogflowConversationProfileHumanAgentAssistantConfigMessageAnalysisConfig",
		reflect.TypeOf((*DialogflowConversationProfileHumanAgentAssistantConfigMessageAnalysisConfig)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-google.dialogflowConversationProfile.DialogflowConversationProfileHumanAgentAssistantConfigMessageAnalysisConfigOutputReference",
		reflect.TypeOf((*DialogflowConversationProfileHumanAgentAssistantConfigMessageAnalysisConfigOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "enableEntityExtraction", GoGetter: "EnableEntityExtraction"},
			_jsii_.MemberProperty{JsiiProperty: "enableEntityExtractionInput", GoGetter: "EnableEntityExtractionInput"},
			_jsii_.MemberProperty{JsiiProperty: "enableSentimentAnalysis", GoGetter: "EnableSentimentAnalysis"},
			_jsii_.MemberProperty{JsiiProperty: "enableSentimentAnalysisInput", GoGetter: "EnableSentimentAnalysisInput"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "resetEnableEntityExtraction", GoMethod: "ResetEnableEntityExtraction"},
			_jsii_.MemberMethod{JsiiMethod: "resetEnableSentimentAnalysis", GoMethod: "ResetEnableSentimentAnalysis"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_DialogflowConversationProfileHumanAgentAssistantConfigMessageAnalysisConfigOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-google.dialogflowConversationProfile.DialogflowConversationProfileHumanAgentAssistantConfigNotificationConfig",
		reflect.TypeOf((*DialogflowConversationProfileHumanAgentAssistantConfigNotificationConfig)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-google.dialogflowConversationProfile.DialogflowConversationProfileHumanAgentAssistantConfigNotificationConfigOutputReference",
		reflect.TypeOf((*DialogflowConversationProfileHumanAgentAssistantConfigNotificationConfigOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "messageFormat", GoGetter: "MessageFormat"},
			_jsii_.MemberProperty{JsiiProperty: "messageFormatInput", GoGetter: "MessageFormatInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetMessageFormat", GoMethod: "ResetMessageFormat"},
			_jsii_.MemberMethod{JsiiMethod: "resetTopic", GoMethod: "ResetTopic"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberProperty{JsiiProperty: "topic", GoGetter: "Topic"},
			_jsii_.MemberProperty{JsiiProperty: "topicInput", GoGetter: "TopicInput"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_DialogflowConversationProfileHumanAgentAssistantConfigNotificationConfigOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-google.dialogflowConversationProfile.DialogflowConversationProfileHumanAgentAssistantConfigOutputReference",
		reflect.TypeOf((*DialogflowConversationProfileHumanAgentAssistantConfigOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "endUserSuggestionConfig", GoGetter: "EndUserSuggestionConfig"},
			_jsii_.MemberProperty{JsiiProperty: "endUserSuggestionConfigInput", GoGetter: "EndUserSuggestionConfigInput"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "humanAgentSuggestionConfig", GoGetter: "HumanAgentSuggestionConfig"},
			_jsii_.MemberProperty{JsiiProperty: "humanAgentSuggestionConfigInput", GoGetter: "HumanAgentSuggestionConfigInput"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "messageAnalysisConfig", GoGetter: "MessageAnalysisConfig"},
			_jsii_.MemberProperty{JsiiProperty: "messageAnalysisConfigInput", GoGetter: "MessageAnalysisConfigInput"},
			_jsii_.MemberProperty{JsiiProperty: "notificationConfig", GoGetter: "NotificationConfig"},
			_jsii_.MemberProperty{JsiiProperty: "notificationConfigInput", GoGetter: "NotificationConfigInput"},
			_jsii_.MemberMethod{JsiiMethod: "putEndUserSuggestionConfig", GoMethod: "PutEndUserSuggestionConfig"},
			_jsii_.MemberMethod{JsiiMethod: "putHumanAgentSuggestionConfig", GoMethod: "PutHumanAgentSuggestionConfig"},
			_jsii_.MemberMethod{JsiiMethod: "putMessageAnalysisConfig", GoMethod: "PutMessageAnalysisConfig"},
			_jsii_.MemberMethod{JsiiMethod: "putNotificationConfig", GoMethod: "PutNotificationConfig"},
			_jsii_.MemberMethod{JsiiMethod: "resetEndUserSuggestionConfig", GoMethod: "ResetEndUserSuggestionConfig"},
			_jsii_.MemberMethod{JsiiMethod: "resetHumanAgentSuggestionConfig", GoMethod: "ResetHumanAgentSuggestionConfig"},
			_jsii_.MemberMethod{JsiiMethod: "resetMessageAnalysisConfig", GoMethod: "ResetMessageAnalysisConfig"},
			_jsii_.MemberMethod{JsiiMethod: "resetNotificationConfig", GoMethod: "ResetNotificationConfig"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_DialogflowConversationProfileHumanAgentAssistantConfigOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-google.dialogflowConversationProfile.DialogflowConversationProfileHumanAgentHandoffConfig",
		reflect.TypeOf((*DialogflowConversationProfileHumanAgentHandoffConfig)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-google.dialogflowConversationProfile.DialogflowConversationProfileHumanAgentHandoffConfigLivePersonConfig",
		reflect.TypeOf((*DialogflowConversationProfileHumanAgentHandoffConfigLivePersonConfig)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-google.dialogflowConversationProfile.DialogflowConversationProfileHumanAgentHandoffConfigLivePersonConfigOutputReference",
		reflect.TypeOf((*DialogflowConversationProfileHumanAgentHandoffConfigLivePersonConfigOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "accountNumber", GoGetter: "AccountNumber"},
			_jsii_.MemberProperty{JsiiProperty: "accountNumberInput", GoGetter: "AccountNumberInput"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_DialogflowConversationProfileHumanAgentHandoffConfigLivePersonConfigOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-google.dialogflowConversationProfile.DialogflowConversationProfileHumanAgentHandoffConfigOutputReference",
		reflect.TypeOf((*DialogflowConversationProfileHumanAgentHandoffConfigOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "livePersonConfig", GoGetter: "LivePersonConfig"},
			_jsii_.MemberProperty{JsiiProperty: "livePersonConfigInput", GoGetter: "LivePersonConfigInput"},
			_jsii_.MemberMethod{JsiiMethod: "putLivePersonConfig", GoMethod: "PutLivePersonConfig"},
			_jsii_.MemberMethod{JsiiMethod: "resetLivePersonConfig", GoMethod: "ResetLivePersonConfig"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_DialogflowConversationProfileHumanAgentHandoffConfigOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-google.dialogflowConversationProfile.DialogflowConversationProfileLoggingConfig",
		reflect.TypeOf((*DialogflowConversationProfileLoggingConfig)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-google.dialogflowConversationProfile.DialogflowConversationProfileLoggingConfigOutputReference",
		reflect.TypeOf((*DialogflowConversationProfileLoggingConfigOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "enableStackdriverLogging", GoGetter: "EnableStackdriverLogging"},
			_jsii_.MemberProperty{JsiiProperty: "enableStackdriverLoggingInput", GoGetter: "EnableStackdriverLoggingInput"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "resetEnableStackdriverLogging", GoMethod: "ResetEnableStackdriverLogging"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_DialogflowConversationProfileLoggingConfigOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-google.dialogflowConversationProfile.DialogflowConversationProfileNewMessageEventNotificationConfig",
		reflect.TypeOf((*DialogflowConversationProfileNewMessageEventNotificationConfig)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-google.dialogflowConversationProfile.DialogflowConversationProfileNewMessageEventNotificationConfigOutputReference",
		reflect.TypeOf((*DialogflowConversationProfileNewMessageEventNotificationConfigOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "messageFormat", GoGetter: "MessageFormat"},
			_jsii_.MemberProperty{JsiiProperty: "messageFormatInput", GoGetter: "MessageFormatInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetMessageFormat", GoMethod: "ResetMessageFormat"},
			_jsii_.MemberMethod{JsiiMethod: "resetTopic", GoMethod: "ResetTopic"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberProperty{JsiiProperty: "topic", GoGetter: "Topic"},
			_jsii_.MemberProperty{JsiiProperty: "topicInput", GoGetter: "TopicInput"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_DialogflowConversationProfileNewMessageEventNotificationConfigOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-google.dialogflowConversationProfile.DialogflowConversationProfileNotificationConfig",
		reflect.TypeOf((*DialogflowConversationProfileNotificationConfig)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-google.dialogflowConversationProfile.DialogflowConversationProfileNotificationConfigOutputReference",
		reflect.TypeOf((*DialogflowConversationProfileNotificationConfigOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "messageFormat", GoGetter: "MessageFormat"},
			_jsii_.MemberProperty{JsiiProperty: "messageFormatInput", GoGetter: "MessageFormatInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetMessageFormat", GoMethod: "ResetMessageFormat"},
			_jsii_.MemberMethod{JsiiMethod: "resetTopic", GoMethod: "ResetTopic"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberProperty{JsiiProperty: "topic", GoGetter: "Topic"},
			_jsii_.MemberProperty{JsiiProperty: "topicInput", GoGetter: "TopicInput"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_DialogflowConversationProfileNotificationConfigOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-google.dialogflowConversationProfile.DialogflowConversationProfileSttConfig",
		reflect.TypeOf((*DialogflowConversationProfileSttConfig)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-google.dialogflowConversationProfile.DialogflowConversationProfileSttConfigOutputReference",
		reflect.TypeOf((*DialogflowConversationProfileSttConfigOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "audioEncoding", GoGetter: "AudioEncoding"},
			_jsii_.MemberProperty{JsiiProperty: "audioEncodingInput", GoGetter: "AudioEncodingInput"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "enableWordInfo", GoGetter: "EnableWordInfo"},
			_jsii_.MemberProperty{JsiiProperty: "enableWordInfoInput", GoGetter: "EnableWordInfoInput"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "languageCode", GoGetter: "LanguageCode"},
			_jsii_.MemberProperty{JsiiProperty: "languageCodeInput", GoGetter: "LanguageCodeInput"},
			_jsii_.MemberProperty{JsiiProperty: "model", GoGetter: "Model"},
			_jsii_.MemberProperty{JsiiProperty: "modelInput", GoGetter: "ModelInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetAudioEncoding", GoMethod: "ResetAudioEncoding"},
			_jsii_.MemberMethod{JsiiMethod: "resetEnableWordInfo", GoMethod: "ResetEnableWordInfo"},
			_jsii_.MemberMethod{JsiiMethod: "resetLanguageCode", GoMethod: "ResetLanguageCode"},
			_jsii_.MemberMethod{JsiiMethod: "resetModel", GoMethod: "ResetModel"},
			_jsii_.MemberMethod{JsiiMethod: "resetSampleRateHertz", GoMethod: "ResetSampleRateHertz"},
			_jsii_.MemberMethod{JsiiMethod: "resetSpeechModelVariant", GoMethod: "ResetSpeechModelVariant"},
			_jsii_.MemberMethod{JsiiMethod: "resetUseTimeoutBasedEndpointing", GoMethod: "ResetUseTimeoutBasedEndpointing"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "sampleRateHertz", GoGetter: "SampleRateHertz"},
			_jsii_.MemberProperty{JsiiProperty: "sampleRateHertzInput", GoGetter: "SampleRateHertzInput"},
			_jsii_.MemberProperty{JsiiProperty: "speechModelVariant", GoGetter: "SpeechModelVariant"},
			_jsii_.MemberProperty{JsiiProperty: "speechModelVariantInput", GoGetter: "SpeechModelVariantInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "useTimeoutBasedEndpointing", GoGetter: "UseTimeoutBasedEndpointing"},
			_jsii_.MemberProperty{JsiiProperty: "useTimeoutBasedEndpointingInput", GoGetter: "UseTimeoutBasedEndpointingInput"},
		},
		func() interface{} {
			j := jsiiProxy_DialogflowConversationProfileSttConfigOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-google.dialogflowConversationProfile.DialogflowConversationProfileTimeouts",
		reflect.TypeOf((*DialogflowConversationProfileTimeouts)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-google.dialogflowConversationProfile.DialogflowConversationProfileTimeoutsOutputReference",
		reflect.TypeOf((*DialogflowConversationProfileTimeoutsOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "create", GoGetter: "Create"},
			_jsii_.MemberProperty{JsiiProperty: "createInput", GoGetter: "CreateInput"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "delete", GoGetter: "Delete"},
			_jsii_.MemberProperty{JsiiProperty: "deleteInput", GoGetter: "DeleteInput"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "resetCreate", GoMethod: "ResetCreate"},
			_jsii_.MemberMethod{JsiiMethod: "resetDelete", GoMethod: "ResetDelete"},
			_jsii_.MemberMethod{JsiiMethod: "resetUpdate", GoMethod: "ResetUpdate"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "update", GoGetter: "Update"},
			_jsii_.MemberProperty{JsiiProperty: "updateInput", GoGetter: "UpdateInput"},
		},
		func() interface{} {
			j := jsiiProxy_DialogflowConversationProfileTimeoutsOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-google.dialogflowConversationProfile.DialogflowConversationProfileTtsConfig",
		reflect.TypeOf((*DialogflowConversationProfileTtsConfig)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-google.dialogflowConversationProfile.DialogflowConversationProfileTtsConfigOutputReference",
		reflect.TypeOf((*DialogflowConversationProfileTtsConfigOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "effectsProfileId", GoGetter: "EffectsProfileId"},
			_jsii_.MemberProperty{JsiiProperty: "effectsProfileIdInput", GoGetter: "EffectsProfileIdInput"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "pitch", GoGetter: "Pitch"},
			_jsii_.MemberProperty{JsiiProperty: "pitchInput", GoGetter: "PitchInput"},
			_jsii_.MemberMethod{JsiiMethod: "putVoice", GoMethod: "PutVoice"},
			_jsii_.MemberMethod{JsiiMethod: "resetEffectsProfileId", GoMethod: "ResetEffectsProfileId"},
			_jsii_.MemberMethod{JsiiMethod: "resetPitch", GoMethod: "ResetPitch"},
			_jsii_.MemberMethod{JsiiMethod: "resetSpeakingRate", GoMethod: "ResetSpeakingRate"},
			_jsii_.MemberMethod{JsiiMethod: "resetVoice", GoMethod: "ResetVoice"},
			_jsii_.MemberMethod{JsiiMethod: "resetVolumeGainDb", GoMethod: "ResetVolumeGainDb"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "speakingRate", GoGetter: "SpeakingRate"},
			_jsii_.MemberProperty{JsiiProperty: "speakingRateInput", GoGetter: "SpeakingRateInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "voice", GoGetter: "Voice"},
			_jsii_.MemberProperty{JsiiProperty: "voiceInput", GoGetter: "VoiceInput"},
			_jsii_.MemberProperty{JsiiProperty: "volumeGainDb", GoGetter: "VolumeGainDb"},
			_jsii_.MemberProperty{JsiiProperty: "volumeGainDbInput", GoGetter: "VolumeGainDbInput"},
		},
		func() interface{} {
			j := jsiiProxy_DialogflowConversationProfileTtsConfigOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-google.dialogflowConversationProfile.DialogflowConversationProfileTtsConfigVoice",
		reflect.TypeOf((*DialogflowConversationProfileTtsConfigVoice)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-google.dialogflowConversationProfile.DialogflowConversationProfileTtsConfigVoiceOutputReference",
		reflect.TypeOf((*DialogflowConversationProfileTtsConfigVoiceOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "nameInput", GoGetter: "NameInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetName", GoMethod: "ResetName"},
			_jsii_.MemberMethod{JsiiMethod: "resetSsmlGender", GoMethod: "ResetSsmlGender"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "ssmlGender", GoGetter: "SsmlGender"},
			_jsii_.MemberProperty{JsiiProperty: "ssmlGenderInput", GoGetter: "SsmlGenderInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_DialogflowConversationProfileTtsConfigVoiceOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
}
