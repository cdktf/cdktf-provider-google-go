// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package clouddeployautomation

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_ClouddeployAutomationRulesList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (c *jsiiProxy_ClouddeployAutomationRulesList) validateGetParameters(index *float64) error {
	return nil
}

func (c *jsiiProxy_ClouddeployAutomationRulesList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_ClouddeployAutomationRulesList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_ClouddeployAutomationRulesList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_ClouddeployAutomationRulesList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_ClouddeployAutomationRulesList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewClouddeployAutomationRulesListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

