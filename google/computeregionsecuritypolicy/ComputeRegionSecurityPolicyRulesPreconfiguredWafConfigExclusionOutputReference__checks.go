// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build !no_runtime_type_checking

package computeregionsecuritypolicy

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

func (c *jsiiProxy_ComputeRegionSecurityPolicyRulesPreconfiguredWafConfigExclusionOutputReference) validateGetAnyMapAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (c *jsiiProxy_ComputeRegionSecurityPolicyRulesPreconfiguredWafConfigExclusionOutputReference) validateGetBooleanAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (c *jsiiProxy_ComputeRegionSecurityPolicyRulesPreconfiguredWafConfigExclusionOutputReference) validateGetBooleanMapAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (c *jsiiProxy_ComputeRegionSecurityPolicyRulesPreconfiguredWafConfigExclusionOutputReference) validateGetListAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (c *jsiiProxy_ComputeRegionSecurityPolicyRulesPreconfiguredWafConfigExclusionOutputReference) validateGetNumberAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (c *jsiiProxy_ComputeRegionSecurityPolicyRulesPreconfiguredWafConfigExclusionOutputReference) validateGetNumberListAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (c *jsiiProxy_ComputeRegionSecurityPolicyRulesPreconfiguredWafConfigExclusionOutputReference) validateGetNumberMapAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (c *jsiiProxy_ComputeRegionSecurityPolicyRulesPreconfiguredWafConfigExclusionOutputReference) validateGetStringAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (c *jsiiProxy_ComputeRegionSecurityPolicyRulesPreconfiguredWafConfigExclusionOutputReference) validateGetStringMapAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (c *jsiiProxy_ComputeRegionSecurityPolicyRulesPreconfiguredWafConfigExclusionOutputReference) validateInterpolationForAttributeParameters(property *string) error {
	if property == nil {
		return fmt.Errorf("parameter property is required, but nil was provided")
	}

	return nil
}

func (c *jsiiProxy_ComputeRegionSecurityPolicyRulesPreconfiguredWafConfigExclusionOutputReference) validatePutRequestCookieParameters(value interface{}) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	switch value.(type) {
	case cdktf.IResolvable:
		// ok
	case *[]*ComputeRegionSecurityPolicyRulesPreconfiguredWafConfigExclusionRequestCookie:
		value := value.(*[]*ComputeRegionSecurityPolicyRulesPreconfiguredWafConfigExclusionRequestCookie)
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	case []*ComputeRegionSecurityPolicyRulesPreconfiguredWafConfigExclusionRequestCookie:
		value_ := value.([]*ComputeRegionSecurityPolicyRulesPreconfiguredWafConfigExclusionRequestCookie)
		value := &value_
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	default:
		if !_jsii_.IsAnonymousProxy(value) {
			return fmt.Errorf("parameter value must be one of the allowed types: cdktf.IResolvable, *[]*ComputeRegionSecurityPolicyRulesPreconfiguredWafConfigExclusionRequestCookie; received %#v (a %T)", value, value)
		}
	}

	return nil
}

func (c *jsiiProxy_ComputeRegionSecurityPolicyRulesPreconfiguredWafConfigExclusionOutputReference) validatePutRequestHeaderParameters(value interface{}) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	switch value.(type) {
	case cdktf.IResolvable:
		// ok
	case *[]*ComputeRegionSecurityPolicyRulesPreconfiguredWafConfigExclusionRequestHeader:
		value := value.(*[]*ComputeRegionSecurityPolicyRulesPreconfiguredWafConfigExclusionRequestHeader)
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	case []*ComputeRegionSecurityPolicyRulesPreconfiguredWafConfigExclusionRequestHeader:
		value_ := value.([]*ComputeRegionSecurityPolicyRulesPreconfiguredWafConfigExclusionRequestHeader)
		value := &value_
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	default:
		if !_jsii_.IsAnonymousProxy(value) {
			return fmt.Errorf("parameter value must be one of the allowed types: cdktf.IResolvable, *[]*ComputeRegionSecurityPolicyRulesPreconfiguredWafConfigExclusionRequestHeader; received %#v (a %T)", value, value)
		}
	}

	return nil
}

func (c *jsiiProxy_ComputeRegionSecurityPolicyRulesPreconfiguredWafConfigExclusionOutputReference) validatePutRequestQueryParamParameters(value interface{}) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	switch value.(type) {
	case cdktf.IResolvable:
		// ok
	case *[]*ComputeRegionSecurityPolicyRulesPreconfiguredWafConfigExclusionRequestQueryParam:
		value := value.(*[]*ComputeRegionSecurityPolicyRulesPreconfiguredWafConfigExclusionRequestQueryParam)
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	case []*ComputeRegionSecurityPolicyRulesPreconfiguredWafConfigExclusionRequestQueryParam:
		value_ := value.([]*ComputeRegionSecurityPolicyRulesPreconfiguredWafConfigExclusionRequestQueryParam)
		value := &value_
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	default:
		if !_jsii_.IsAnonymousProxy(value) {
			return fmt.Errorf("parameter value must be one of the allowed types: cdktf.IResolvable, *[]*ComputeRegionSecurityPolicyRulesPreconfiguredWafConfigExclusionRequestQueryParam; received %#v (a %T)", value, value)
		}
	}

	return nil
}

func (c *jsiiProxy_ComputeRegionSecurityPolicyRulesPreconfiguredWafConfigExclusionOutputReference) validatePutRequestUriParameters(value interface{}) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	switch value.(type) {
	case cdktf.IResolvable:
		// ok
	case *[]*ComputeRegionSecurityPolicyRulesPreconfiguredWafConfigExclusionRequestUri:
		value := value.(*[]*ComputeRegionSecurityPolicyRulesPreconfiguredWafConfigExclusionRequestUri)
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	case []*ComputeRegionSecurityPolicyRulesPreconfiguredWafConfigExclusionRequestUri:
		value_ := value.([]*ComputeRegionSecurityPolicyRulesPreconfiguredWafConfigExclusionRequestUri)
		value := &value_
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	default:
		if !_jsii_.IsAnonymousProxy(value) {
			return fmt.Errorf("parameter value must be one of the allowed types: cdktf.IResolvable, *[]*ComputeRegionSecurityPolicyRulesPreconfiguredWafConfigExclusionRequestUri; received %#v (a %T)", value, value)
		}
	}

	return nil
}

func (c *jsiiProxy_ComputeRegionSecurityPolicyRulesPreconfiguredWafConfigExclusionOutputReference) validateResolveParameters(_context cdktf.IResolveContext) error {
	if _context == nil {
		return fmt.Errorf("parameter _context is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_ComputeRegionSecurityPolicyRulesPreconfiguredWafConfigExclusionOutputReference) validateSetComplexObjectIndexParameters(val interface{}) error {
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

func (j *jsiiProxy_ComputeRegionSecurityPolicyRulesPreconfiguredWafConfigExclusionOutputReference) validateSetComplexObjectIsFromSetParameters(val *bool) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_ComputeRegionSecurityPolicyRulesPreconfiguredWafConfigExclusionOutputReference) validateSetInternalValueParameters(val interface{}) error {
	switch val.(type) {
	case cdktf.IResolvable:
		// ok
	case *ComputeRegionSecurityPolicyRulesPreconfiguredWafConfigExclusion:
		val := val.(*ComputeRegionSecurityPolicyRulesPreconfiguredWafConfigExclusion)
		if err := _jsii_.ValidateStruct(val, func() string { return "parameter val" }); err != nil {
			return err
		}
	case ComputeRegionSecurityPolicyRulesPreconfiguredWafConfigExclusion:
		val_ := val.(ComputeRegionSecurityPolicyRulesPreconfiguredWafConfigExclusion)
		val := &val_
		if err := _jsii_.ValidateStruct(val, func() string { return "parameter val" }); err != nil {
			return err
		}
	default:
		if !_jsii_.IsAnonymousProxy(val) {
			return fmt.Errorf("parameter val must be one of the allowed types: cdktf.IResolvable, *ComputeRegionSecurityPolicyRulesPreconfiguredWafConfigExclusion; received %#v (a %T)", val, val)
		}
	}

	return nil
}

func (j *jsiiProxy_ComputeRegionSecurityPolicyRulesPreconfiguredWafConfigExclusionOutputReference) validateSetTargetRuleIdsParameters(val *[]*string) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_ComputeRegionSecurityPolicyRulesPreconfiguredWafConfigExclusionOutputReference) validateSetTargetRuleSetParameters(val *string) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_ComputeRegionSecurityPolicyRulesPreconfiguredWafConfigExclusionOutputReference) validateSetTerraformAttributeParameters(val *string) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_ComputeRegionSecurityPolicyRulesPreconfiguredWafConfigExclusionOutputReference) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func validateNewComputeRegionSecurityPolicyRulesPreconfiguredWafConfigExclusionOutputReferenceParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) error {
	if terraformResource == nil {
		return fmt.Errorf("parameter terraformResource is required, but nil was provided")
	}

	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	if complexObjectIndex == nil {
		return fmt.Errorf("parameter complexObjectIndex is required, but nil was provided")
	}

	if complexObjectIsFromSet == nil {
		return fmt.Errorf("parameter complexObjectIsFromSet is required, but nil was provided")
	}

	return nil
}

