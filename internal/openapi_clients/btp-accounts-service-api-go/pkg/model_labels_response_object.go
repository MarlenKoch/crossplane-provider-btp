/*
Accounts Service

The Accounts service provides REST APIs that are responsible for the management of global accounts, and the creation and management of directories, subaccounts, and their custom properties/tags.  Global accounts represent a business entity and contain contract information, including customer details and purchased entitlements. The global account is the context for billing each customer.  Use the subaccount APIs to structure your global account according to your organization's and project's requirements regarding members, authorizations, and quotas. This service also provides you with APIs for creating and managing directories. While the use of directories is optional, they allow you to further organize and manage your subaccounts according to your specific technical and business needs. The service also lets you manage the custom properties/tags that you associate with your directories and subaccounts. NOTE: These APIs are relevant only for cloud management tools feature set B. For details and information about whether this applies to your global account, see [Cloud Management Tools - Feature Set Overview](https://help.sap.com/viewer/65de2977205c403bbc107264b8eccf4b/Cloud/en-US/caf4e4e23aef4666ad8f125af393dfb2.html).  See also: * [Authorization](https://help.sap.com/viewer/65de2977205c403bbc107264b8eccf4b/latest/en-US/3670474a58c24ac2b082e76cbbd9dc19.html) * [Rate Limiting](https://help.sap.com/viewer/65de2977205c403bbc107264b8eccf4b/latest/en-US/77b217b3f57a45b987eb7fbc3305ce1e.html) * [Error Response Format](https://help.sap.com/viewer/65de2977205c403bbc107264b8eccf4b/latest/en-US/77fef2fb104b4b1795e2e6cee790e8b8.html) * [Asynchronous Jobs](https://help.sap.com/viewer/65de2977205c403bbc107264b8eccf4b/latest/en-US/0a0a6ab0ad114d72a6611c1c6b21683e.html)

API version: 1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the LabelsResponseObject type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &LabelsResponseObject{}

// LabelsResponseObject Labels assigned as key-value pairs to the entity.
type LabelsResponseObject struct {
	// User-defined labels that are assigned as key-value pairs in a JSON array to the entity.  Example: {    \"Cost Center\": [\"19700626\"],    \"Department\": [\"Sales\"],    \"Contacts\": [\"name1@example.com\",\"name2@example.com\"],    \"EMEA\":[] }
	Labels *map[string][]string `json:"labels,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _LabelsResponseObject LabelsResponseObject

// NewLabelsResponseObject instantiates a new LabelsResponseObject object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLabelsResponseObject() *LabelsResponseObject {
	this := LabelsResponseObject{}
	return &this
}

// NewLabelsResponseObjectWithDefaults instantiates a new LabelsResponseObject object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLabelsResponseObjectWithDefaults() *LabelsResponseObject {
	this := LabelsResponseObject{}
	return &this
}

// GetLabels returns the Labels field value if set, zero value otherwise.
func (o *LabelsResponseObject) GetLabels() map[string][]string {
	if o == nil || IsNil(o.Labels) {
		var ret map[string][]string
		return ret
	}
	return *o.Labels
}

// GetLabelsOk returns a tuple with the Labels field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LabelsResponseObject) GetLabelsOk() (*map[string][]string, bool) {
	if o == nil || IsNil(o.Labels) {
		return nil, false
	}
	return o.Labels, true
}

// HasLabels returns a boolean if a field has been set.
func (o *LabelsResponseObject) HasLabels() bool {
	if o != nil && !IsNil(o.Labels) {
		return true
	}

	return false
}

// SetLabels gets a reference to the given map[string][]string and assigns it to the Labels field.
func (o *LabelsResponseObject) SetLabels(v map[string][]string) {
	o.Labels = &v
}

func (o LabelsResponseObject) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o LabelsResponseObject) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Labels) {
		toSerialize["labels"] = o.Labels
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *LabelsResponseObject) UnmarshalJSON(data []byte) (err error) {
	varLabelsResponseObject := _LabelsResponseObject{}

	err = json.Unmarshal(data, &varLabelsResponseObject)

	if err != nil {
		return err
	}

	*o = LabelsResponseObject(varLabelsResponseObject)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "labels")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableLabelsResponseObject struct {
	value *LabelsResponseObject
	isSet bool
}

func (v NullableLabelsResponseObject) Get() *LabelsResponseObject {
	return v.value
}

func (v *NullableLabelsResponseObject) Set(val *LabelsResponseObject) {
	v.value = val
	v.isSet = true
}

func (v NullableLabelsResponseObject) IsSet() bool {
	return v.isSet
}

func (v *NullableLabelsResponseObject) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLabelsResponseObject(val *LabelsResponseObject) *NullableLabelsResponseObject {
	return &NullableLabelsResponseObject{value: val, isSet: true}
}

func (v NullableLabelsResponseObject) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLabelsResponseObject) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


