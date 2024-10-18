/*
Accounts Service

The Accounts service provides REST APIs that are responsible for the management of global accounts, and the creation and management of directories, subaccounts, and their custom properties/tags.  Global accounts represent a business entity and contain contract information, including customer details and purchased entitlements. The global account is the context for billing each customer.  Use the subaccount APIs to structure your global account according to your organization's and project's requirements regarding members, authorizations, and quotas. This service also provides you with APIs for creating and managing directories. While the use of directories is optional, they allow you to further organize and manage your subaccounts according to your specific technical and business needs. The service also lets you manage the custom properties/tags that you associate with your directories and subaccounts. NOTE: These APIs are relevant only for cloud management tools feature set B. For details and information about whether this applies to your global account, see [Cloud Management Tools - Feature Set Overview](https://help.sap.com/viewer/65de2977205c403bbc107264b8eccf4b/Cloud/en-US/caf4e4e23aef4666ad8f125af393dfb2.html).  See also: * [Authorization](https://help.sap.com/viewer/65de2977205c403bbc107264b8eccf4b/latest/en-US/3670474a58c24ac2b082e76cbbd9dc19.html) * [Rate Limiting](https://help.sap.com/viewer/65de2977205c403bbc107264b8eccf4b/latest/en-US/77b217b3f57a45b987eb7fbc3305ce1e.html) * [Error Response Format](https://help.sap.com/viewer/65de2977205c403bbc107264b8eccf4b/latest/en-US/77fef2fb104b4b1795e2e6cee790e8b8.html) * [Asynchronous Jobs](https://help.sap.com/viewer/65de2977205c403bbc107264b8eccf4b/latest/en-US/0a0a6ab0ad114d72a6611c1c6b21683e.html)

API version: 1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"fmt"
)

// checks if the UpdateEntitySettingsRequestPayload type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateEntitySettingsRequestPayload{}

// UpdateEntitySettingsRequestPayload Settings as key-value pairs to assign, update, and remove from the entity.
type UpdateEntitySettingsRequestPayload struct {
	// A key for the property. Limited to 200 characters.
	Key string `json:"key"`
	// A value for the corresponding key. Limited to 2000 characters.
	Value map[string]interface{} `json:"value"`
	AdditionalProperties map[string]interface{}
}

type _UpdateEntitySettingsRequestPayload UpdateEntitySettingsRequestPayload

// NewUpdateEntitySettingsRequestPayload instantiates a new UpdateEntitySettingsRequestPayload object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateEntitySettingsRequestPayload(key string, value map[string]interface{}) *UpdateEntitySettingsRequestPayload {
	this := UpdateEntitySettingsRequestPayload{}
	this.Key = key
	this.Value = value
	return &this
}

// NewUpdateEntitySettingsRequestPayloadWithDefaults instantiates a new UpdateEntitySettingsRequestPayload object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateEntitySettingsRequestPayloadWithDefaults() *UpdateEntitySettingsRequestPayload {
	this := UpdateEntitySettingsRequestPayload{}
	return &this
}

// GetKey returns the Key field value
func (o *UpdateEntitySettingsRequestPayload) GetKey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Key
}

// GetKeyOk returns a tuple with the Key field value
// and a boolean to check if the value has been set.
func (o *UpdateEntitySettingsRequestPayload) GetKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Key, true
}

// SetKey sets field value
func (o *UpdateEntitySettingsRequestPayload) SetKey(v string) {
	o.Key = v
}

// GetValue returns the Value field value
func (o *UpdateEntitySettingsRequestPayload) GetValue() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}

	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *UpdateEntitySettingsRequestPayload) GetValueOk() (map[string]interface{}, bool) {
	if o == nil {
		return map[string]interface{}{}, false
	}
	return o.Value, true
}

// SetValue sets field value
func (o *UpdateEntitySettingsRequestPayload) SetValue(v map[string]interface{}) {
	o.Value = v
}

func (o UpdateEntitySettingsRequestPayload) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateEntitySettingsRequestPayload) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["key"] = o.Key
	toSerialize["value"] = o.Value

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *UpdateEntitySettingsRequestPayload) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"key",
		"value",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varUpdateEntitySettingsRequestPayload := _UpdateEntitySettingsRequestPayload{}

	err = json.Unmarshal(data, &varUpdateEntitySettingsRequestPayload)

	if err != nil {
		return err
	}

	*o = UpdateEntitySettingsRequestPayload(varUpdateEntitySettingsRequestPayload)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "key")
		delete(additionalProperties, "value")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableUpdateEntitySettingsRequestPayload struct {
	value *UpdateEntitySettingsRequestPayload
	isSet bool
}

func (v NullableUpdateEntitySettingsRequestPayload) Get() *UpdateEntitySettingsRequestPayload {
	return v.value
}

func (v *NullableUpdateEntitySettingsRequestPayload) Set(val *UpdateEntitySettingsRequestPayload) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateEntitySettingsRequestPayload) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateEntitySettingsRequestPayload) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateEntitySettingsRequestPayload(val *UpdateEntitySettingsRequestPayload) *NullableUpdateEntitySettingsRequestPayload {
	return &NullableUpdateEntitySettingsRequestPayload{value: val, isSet: true}
}

func (v NullableUpdateEntitySettingsRequestPayload) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateEntitySettingsRequestPayload) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


