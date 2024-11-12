// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package computenodetemplate

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_ComputeNodeTemplateDisksList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (c *jsiiProxy_ComputeNodeTemplateDisksList) validateGetParameters(index *float64) error {
	return nil
}

func (c *jsiiProxy_ComputeNodeTemplateDisksList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_ComputeNodeTemplateDisksList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_ComputeNodeTemplateDisksList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_ComputeNodeTemplateDisksList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_ComputeNodeTemplateDisksList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewComputeNodeTemplateDisksListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

