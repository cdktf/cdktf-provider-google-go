// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package chroniclerule

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_ChronicleRuleSeverityList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (c *jsiiProxy_ChronicleRuleSeverityList) validateGetParameters(index *float64) error {
	return nil
}

func (c *jsiiProxy_ChronicleRuleSeverityList) validateResolveParameters(context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_ChronicleRuleSeverityList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_ChronicleRuleSeverityList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_ChronicleRuleSeverityList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewChronicleRuleSeverityListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

