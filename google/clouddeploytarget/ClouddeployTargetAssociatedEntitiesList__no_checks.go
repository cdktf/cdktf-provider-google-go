// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package clouddeploytarget

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_ClouddeployTargetAssociatedEntitiesList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (c *jsiiProxy_ClouddeployTargetAssociatedEntitiesList) validateGetParameters(index *float64) error {
	return nil
}

func (c *jsiiProxy_ClouddeployTargetAssociatedEntitiesList) validateResolveParameters(context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_ClouddeployTargetAssociatedEntitiesList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_ClouddeployTargetAssociatedEntitiesList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_ClouddeployTargetAssociatedEntitiesList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_ClouddeployTargetAssociatedEntitiesList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewClouddeployTargetAssociatedEntitiesListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

