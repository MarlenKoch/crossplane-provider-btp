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

// checks if the PropertyResponseObject type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PropertyResponseObject{}

// PropertyResponseObject (Deprecated) Contains information about the labels assigned to a specified subaccount. This field supports only single values per key and is now replaced by the string array \"labels\", which supports multiple values per key. The \"customProperties\" field returns only the first value of any label key that has multiple values assigned to it.
type PropertyResponseObject struct {
	// The unique ID for the corresponding entity.
	AccountGUID string `json:"accountGUID"`
	// The name for the label.
	Key string `json:"key"`
	// The value for the corresponding label key.
	Value string `json:"value"`
	AdditionalProperties map[string]interface{}
}

type _PropertyResponseObject PropertyResponseObject

// NewPropertyResponseObject instantiates a new PropertyResponseObject object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPropertyResponseObject(accountGUID string, key string, value string) *PropertyResponseObject {
	this := PropertyResponseObject{}
	this.AccountGUID = accountGUID
	this.Key = key
	this.Value = value
	return &this
}

// NewPropertyResponseObjectWithDefaults instantiates a new PropertyResponseObject object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPropertyResponseObjectWithDefaults() *PropertyResponseObject {
	this := PropertyResponseObject{}
	return &this
}

// GetAccountGUID returns the AccountGUID field value
func (o *PropertyResponseObject) GetAccountGUID() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AccountGUID
}

// GetAccountGUIDOk returns a tuple with the AccountGUID field value
// and a boolean to check if the value has been set.
func (o *PropertyResponseObject) GetAccountGUIDOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AccountGUID, true
}

// SetAccountGUID sets field value
func (o *PropertyResponseObject) SetAccountGUID(v string) {
	o.AccountGUID = v
}

// GetKey returns the Key field value
func (o *PropertyResponseObject) GetKey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Key
}

// GetKeyOk returns a tuple with the Key field value
// and a boolean to check if the value has been set.
func (o *PropertyResponseObject) GetKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Key, true
}

// SetKey sets field value
func (o *PropertyResponseObject) SetKey(v string) {
	o.Key = v
}

// GetValue returns the Value field value
func (o *PropertyResponseObject) GetValue() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *PropertyResponseObject) GetValueOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value
func (o *PropertyResponseObject) SetValue(v string) {
	o.Value = v
}

func (o PropertyResponseObject) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PropertyResponseObject) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["accountGUID"] = o.AccountGUID
	toSerialize["key"] = o.Key
	toSerialize["value"] = o.Value

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *PropertyResponseObject) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"accountGUID",
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

	varPropertyResponseObject := _PropertyResponseObject{}

	err = json.Unmarshal(data, &varPropertyResponseObject)

	if err != nil {
		return err
	}

	*o = PropertyResponseObject(varPropertyResponseObject)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "accountGUID")
		delete(additionalProperties, "key")
		delete(additionalProperties, "value")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePropertyResponseObject struct {
	value *PropertyResponseObject
	isSet bool
}

func (v NullablePropertyResponseObject) Get() *PropertyResponseObject {
	return v.value
}

func (v *NullablePropertyResponseObject) Set(val *PropertyResponseObject) {
	v.value = val
	v.isSet = true
}

func (v NullablePropertyResponseObject) IsSet() bool {
	return v.isSet
}

func (v *NullablePropertyResponseObject) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePropertyResponseObject(val *PropertyResponseObject) *NullablePropertyResponseObject {
	return &NullablePropertyResponseObject{value: val, isSet: true}
}

func (v NullablePropertyResponseObject) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePropertyResponseObject) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


