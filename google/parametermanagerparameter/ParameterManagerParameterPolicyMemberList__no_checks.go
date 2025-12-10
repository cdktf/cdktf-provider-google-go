// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package parametermanagerparameter

// Building without runtime type checking enabled, so all the below just return nil

func (p *jsiiProxy_ParameterManagerParameterPolicyMemberList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (p *jsiiProxy_ParameterManagerParameterPolicyMemberList) validateGetParameters(index *float64) error {
	return nil
}

func (p *jsiiProxy_ParameterManagerParameterPolicyMemberList) validateResolveParameters(context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_ParameterManagerParameterPolicyMemberList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_ParameterManagerParameterPolicyMemberList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_ParameterManagerParameterPolicyMemberList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewParameterManagerParameterPolicyMemberListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

