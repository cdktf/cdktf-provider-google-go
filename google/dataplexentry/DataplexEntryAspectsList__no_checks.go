// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package dataplexentry

// Building without runtime type checking enabled, so all the below just return nil

func (d *jsiiProxy_DataplexEntryAspectsList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (d *jsiiProxy_DataplexEntryAspectsList) validateGetParameters(index *float64) error {
	return nil
}

func (d *jsiiProxy_DataplexEntryAspectsList) validateResolveParameters(context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_DataplexEntryAspectsList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_DataplexEntryAspectsList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_DataplexEntryAspectsList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_DataplexEntryAspectsList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewDataplexEntryAspectsListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

