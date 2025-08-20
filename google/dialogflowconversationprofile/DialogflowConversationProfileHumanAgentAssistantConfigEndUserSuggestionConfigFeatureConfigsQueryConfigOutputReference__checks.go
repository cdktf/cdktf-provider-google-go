// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build !no_runtime_type_checking

package dialogflowconversationprofile

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

func (d *jsiiProxy_DialogflowConversationProfileHumanAgentAssistantConfigEndUserSuggestionConfigFeatureConfigsQueryConfigOutputReference) validateGetAnyMapAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (d *jsiiProxy_DialogflowConversationProfileHumanAgentAssistantConfigEndUserSuggestionConfigFeatureConfigsQueryConfigOutputReference) validateGetBooleanAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (d *jsiiProxy_DialogflowConversationProfileHumanAgentAssistantConfigEndUserSuggestionConfigFeatureConfigsQueryConfigOutputReference) validateGetBooleanMapAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (d *jsiiProxy_DialogflowConversationProfileHumanAgentAssistantConfigEndUserSuggestionConfigFeatureConfigsQueryConfigOutputReference) validateGetListAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (d *jsiiProxy_DialogflowConversationProfileHumanAgentAssistantConfigEndUserSuggestionConfigFeatureConfigsQueryConfigOutputReference) validateGetNumberAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (d *jsiiProxy_DialogflowConversationProfileHumanAgentAssistantConfigEndUserSuggestionConfigFeatureConfigsQueryConfigOutputReference) validateGetNumberListAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (d *jsiiProxy_DialogflowConversationProfileHumanAgentAssistantConfigEndUserSuggestionConfigFeatureConfigsQueryConfigOutputReference) validateGetNumberMapAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (d *jsiiProxy_DialogflowConversationProfileHumanAgentAssistantConfigEndUserSuggestionConfigFeatureConfigsQueryConfigOutputReference) validateGetStringAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (d *jsiiProxy_DialogflowConversationProfileHumanAgentAssistantConfigEndUserSuggestionConfigFeatureConfigsQueryConfigOutputReference) validateGetStringMapAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (d *jsiiProxy_DialogflowConversationProfileHumanAgentAssistantConfigEndUserSuggestionConfigFeatureConfigsQueryConfigOutputReference) validateInterpolationForAttributeParameters(property *string) error {
	if property == nil {
		return fmt.Errorf("parameter property is required, but nil was provided")
	}

	return nil
}

func (d *jsiiProxy_DialogflowConversationProfileHumanAgentAssistantConfigEndUserSuggestionConfigFeatureConfigsQueryConfigOutputReference) validatePutContextFilterSettingsParameters(value *DialogflowConversationProfileHumanAgentAssistantConfigEndUserSuggestionConfigFeatureConfigsQueryConfigContextFilterSettings) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(value, func() string { return "parameter value" }); err != nil {
		return err
	}

	return nil
}

func (d *jsiiProxy_DialogflowConversationProfileHumanAgentAssistantConfigEndUserSuggestionConfigFeatureConfigsQueryConfigOutputReference) validatePutDialogflowQuerySourceParameters(value *DialogflowConversationProfileHumanAgentAssistantConfigEndUserSuggestionConfigFeatureConfigsQueryConfigDialogflowQuerySource) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(value, func() string { return "parameter value" }); err != nil {
		return err
	}

	return nil
}

func (d *jsiiProxy_DialogflowConversationProfileHumanAgentAssistantConfigEndUserSuggestionConfigFeatureConfigsQueryConfigOutputReference) validatePutDocumentQuerySourceParameters(value *DialogflowConversationProfileHumanAgentAssistantConfigEndUserSuggestionConfigFeatureConfigsQueryConfigDocumentQuerySource) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(value, func() string { return "parameter value" }); err != nil {
		return err
	}

	return nil
}

func (d *jsiiProxy_DialogflowConversationProfileHumanAgentAssistantConfigEndUserSuggestionConfigFeatureConfigsQueryConfigOutputReference) validatePutKnowledgeBaseQuerySourceParameters(value *DialogflowConversationProfileHumanAgentAssistantConfigEndUserSuggestionConfigFeatureConfigsQueryConfigKnowledgeBaseQuerySource) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(value, func() string { return "parameter value" }); err != nil {
		return err
	}

	return nil
}

func (d *jsiiProxy_DialogflowConversationProfileHumanAgentAssistantConfigEndUserSuggestionConfigFeatureConfigsQueryConfigOutputReference) validatePutSectionsParameters(value *DialogflowConversationProfileHumanAgentAssistantConfigEndUserSuggestionConfigFeatureConfigsQueryConfigSections) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(value, func() string { return "parameter value" }); err != nil {
		return err
	}

	return nil
}

func (d *jsiiProxy_DialogflowConversationProfileHumanAgentAssistantConfigEndUserSuggestionConfigFeatureConfigsQueryConfigOutputReference) validateResolveParameters(_context cdktf.IResolveContext) error {
	if _context == nil {
		return fmt.Errorf("parameter _context is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_DialogflowConversationProfileHumanAgentAssistantConfigEndUserSuggestionConfigFeatureConfigsQueryConfigOutputReference) validateSetComplexObjectIndexParameters(val interface{}) error {
	switch val.(type) {
	case *string:
		// ok
	case string:
		// ok
	case *float64:
		// ok
	case float64:
		// ok
	case *int:
		// ok
	case int:
		// ok
	case *uint:
		// ok
	case uint:
		// ok
	case *int8:
		// ok
	case int8:
		// ok
	case *int16:
		// ok
	case int16:
		// ok
	case *int32:
		// ok
	case int32:
		// ok
	case *int64:
		// ok
	case int64:
		// ok
	case *uint8:
		// ok
	case uint8:
		// ok
	case *uint16:
		// ok
	case uint16:
		// ok
	case *uint32:
		// ok
	case uint32:
		// ok
	case *uint64:
		// ok
	case uint64:
		// ok
	default:
		return fmt.Errorf("parameter val must be one of the allowed types: *string, *float64; received %#v (a %T)", val, val)
	}

	return nil
}

func (j *jsiiProxy_DialogflowConversationProfileHumanAgentAssistantConfigEndUserSuggestionConfigFeatureConfigsQueryConfigOutputReference) validateSetComplexObjectIsFromSetParameters(val *bool) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_DialogflowConversationProfileHumanAgentAssistantConfigEndUserSuggestionConfigFeatureConfigsQueryConfigOutputReference) validateSetConfidenceThresholdParameters(val *float64) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_DialogflowConversationProfileHumanAgentAssistantConfigEndUserSuggestionConfigFeatureConfigsQueryConfigOutputReference) validateSetInternalValueParameters(val *DialogflowConversationProfileHumanAgentAssistantConfigEndUserSuggestionConfigFeatureConfigsQueryConfig) error {
	if err := _jsii_.ValidateStruct(val, func() string { return "parameter val" }); err != nil {
		return err
	}

	return nil
}

func (j *jsiiProxy_DialogflowConversationProfileHumanAgentAssistantConfigEndUserSuggestionConfigFeatureConfigsQueryConfigOutputReference) validateSetMaxResultsParameters(val *float64) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_DialogflowConversationProfileHumanAgentAssistantConfigEndUserSuggestionConfigFeatureConfigsQueryConfigOutputReference) validateSetTerraformAttributeParameters(val *string) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_DialogflowConversationProfileHumanAgentAssistantConfigEndUserSuggestionConfigFeatureConfigsQueryConfigOutputReference) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func validateNewDialogflowConversationProfileHumanAgentAssistantConfigEndUserSuggestionConfigFeatureConfigsQueryConfigOutputReferenceParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) error {
	if terraformResource == nil {
		return fmt.Errorf("parameter terraformResource is required, but nil was provided")
	}

	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

