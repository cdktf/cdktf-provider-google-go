// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package workbenchinstance

// Building without runtime type checking enabled, so all the below just return nil

func (w *jsiiProxy_WorkbenchInstanceHealthInfoList) validateGetParameters(index *float64) error {
	return nil
}

func (w *jsiiProxy_WorkbenchInstanceHealthInfoList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_WorkbenchInstanceHealthInfoList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_WorkbenchInstanceHealthInfoList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_WorkbenchInstanceHealthInfoList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewWorkbenchInstanceHealthInfoListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

