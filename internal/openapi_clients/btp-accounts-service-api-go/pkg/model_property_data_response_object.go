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

// checks if the PropertyDataResponseObject type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PropertyDataResponseObject{}

// PropertyDataResponseObject The response object containing information about the data.
type PropertyDataResponseObject struct {
	// The sub-categories associated with the metadata definitions.
	Classification *string `json:"classification,omitempty"`
	// A default value for the corresponding entity type.
	DefaultValue *string `json:"defaultValue,omitempty"`
	// A description for the corresponding metadata definition.
	Description *string `json:"description,omitempty"`
	// Whether the dynamic property is calculated in runtime if it exists in the validation schema for the metadata.
	DynamicProperty *bool `json:"dynamicProperty,omitempty"`
	// The type of entity, namely, global account, subaccount, directory and so on.
	EntityType string `json:"entityType"`
	// The unique ID of the entity type.
	EntityTypeGuid *string `json:"entityTypeGuid,omitempty"`
	// A name for the custom property of the entity type. Limited to 255 characters.
	Key string `json:"key"`
	// The metadata for the properties of the settings object and the valid values.
	ValidationSchema map[string]map[string]interface{} `json:"validationSchema,omitempty"`
	// The user-defined value for the corresponding key. Limited to 1024 characters.
	Value map[string]interface{} `json:"value,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _PropertyDataResponseObject PropertyDataResponseObject

// NewPropertyDataResponseObject instantiates a new PropertyDataResponseObject object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPropertyDataResponseObject(entityType string, key string) *PropertyDataResponseObject {
	this := PropertyDataResponseObject{}
	this.EntityType = entityType
	this.Key = key
	return &this
}

// NewPropertyDataResponseObjectWithDefaults instantiates a new PropertyDataResponseObject object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPropertyDataResponseObjectWithDefaults() *PropertyDataResponseObject {
	this := PropertyDataResponseObject{}
	return &this
}

// GetClassification returns the Classification field value if set, zero value otherwise.
func (o *PropertyDataResponseObject) GetClassification() string {
	if o == nil || IsNil(o.Classification) {
		var ret string
		return ret
	}
	return *o.Classification
}

// GetClassificationOk returns a tuple with the Classification field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PropertyDataResponseObject) GetClassificationOk() (*string, bool) {
	if o == nil || IsNil(o.Classification) {
		return nil, false
	}
	return o.Classification, true
}

// HasClassification returns a boolean if a field has been set.
func (o *PropertyDataResponseObject) HasClassification() bool {
	if o != nil && !IsNil(o.Classification) {
		return true
	}

	return false
}

// SetClassification gets a reference to the given string and assigns it to the Classification field.
func (o *PropertyDataResponseObject) SetClassification(v string) {
	o.Classification = &v
}

// GetDefaultValue returns the DefaultValue field value if set, zero value otherwise.
func (o *PropertyDataResponseObject) GetDefaultValue() string {
	if o == nil || IsNil(o.DefaultValue) {
		var ret string
		return ret
	}
	return *o.DefaultValue
}

// GetDefaultValueOk returns a tuple with the DefaultValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PropertyDataResponseObject) GetDefaultValueOk() (*string, bool) {
	if o == nil || IsNil(o.DefaultValue) {
		return nil, false
	}
	return o.DefaultValue, true
}

// HasDefaultValue returns a boolean if a field has been set.
func (o *PropertyDataResponseObject) HasDefaultValue() bool {
	if o != nil && !IsNil(o.DefaultValue) {
		return true
	}

	return false
}

// SetDefaultValue gets a reference to the given string and assigns it to the DefaultValue field.
func (o *PropertyDataResponseObject) SetDefaultValue(v string) {
	o.DefaultValue = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *PropertyDataResponseObject) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PropertyDataResponseObject) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *PropertyDataResponseObject) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *PropertyDataResponseObject) SetDescription(v string) {
	o.Description = &v
}

// GetDynamicProperty returns the DynamicProperty field value if set, zero value otherwise.
func (o *PropertyDataResponseObject) GetDynamicProperty() bool {
	if o == nil || IsNil(o.DynamicProperty) {
		var ret bool
		return ret
	}
	return *o.DynamicProperty
}

// GetDynamicPropertyOk returns a tuple with the DynamicProperty field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PropertyDataResponseObject) GetDynamicPropertyOk() (*bool, bool) {
	if o == nil || IsNil(o.DynamicProperty) {
		return nil, false
	}
	return o.DynamicProperty, true
}

// HasDynamicProperty returns a boolean if a field has been set.
func (o *PropertyDataResponseObject) HasDynamicProperty() bool {
	if o != nil && !IsNil(o.DynamicProperty) {
		return true
	}

	return false
}

// SetDynamicProperty gets a reference to the given bool and assigns it to the DynamicProperty field.
func (o *PropertyDataResponseObject) SetDynamicProperty(v bool) {
	o.DynamicProperty = &v
}

// GetEntityType returns the EntityType field value
func (o *PropertyDataResponseObject) GetEntityType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.EntityType
}

// GetEntityTypeOk returns a tuple with the EntityType field value
// and a boolean to check if the value has been set.
func (o *PropertyDataResponseObject) GetEntityTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EntityType, true
}

// SetEntityType sets field value
func (o *PropertyDataResponseObject) SetEntityType(v string) {
	o.EntityType = v
}

// GetEntityTypeGuid returns the EntityTypeGuid field value if set, zero value otherwise.
func (o *PropertyDataResponseObject) GetEntityTypeGuid() string {
	if o == nil || IsNil(o.EntityTypeGuid) {
		var ret string
		return ret
	}
	return *o.EntityTypeGuid
}

// GetEntityTypeGuidOk returns a tuple with the EntityTypeGuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PropertyDataResponseObject) GetEntityTypeGuidOk() (*string, bool) {
	if o == nil || IsNil(o.EntityTypeGuid) {
		return nil, false
	}
	return o.EntityTypeGuid, true
}

// HasEntityTypeGuid returns a boolean if a field has been set.
func (o *PropertyDataResponseObject) HasEntityTypeGuid() bool {
	if o != nil && !IsNil(o.EntityTypeGuid) {
		return true
	}

	return false
}

// SetEntityTypeGuid gets a reference to the given string and assigns it to the EntityTypeGuid field.
func (o *PropertyDataResponseObject) SetEntityTypeGuid(v string) {
	o.EntityTypeGuid = &v
}

// GetKey returns the Key field value
func (o *PropertyDataResponseObject) GetKey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Key
}

// GetKeyOk returns a tuple with the Key field value
// and a boolean to check if the value has been set.
func (o *PropertyDataResponseObject) GetKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Key, true
}

// SetKey sets field value
func (o *PropertyDataResponseObject) SetKey(v string) {
	o.Key = v
}

// GetValidationSchema returns the ValidationSchema field value if set, zero value otherwise.
func (o *PropertyDataResponseObject) GetValidationSchema() map[string]map[string]interface{} {
	if o == nil || IsNil(o.ValidationSchema) {
		var ret map[string]map[string]interface{}
		return ret
	}
	return o.ValidationSchema
}

// GetValidationSchemaOk returns a tuple with the ValidationSchema field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PropertyDataResponseObject) GetValidationSchemaOk() (map[string]map[string]interface{}, bool) {
	if o == nil || IsNil(o.ValidationSchema) {
		return map[string]map[string]interface{}{}, false
	}
	return o.ValidationSchema, true
}

// HasValidationSchema returns a boolean if a field has been set.
func (o *PropertyDataResponseObject) HasValidationSchema() bool {
	if o != nil && !IsNil(o.ValidationSchema) {
		return true
	}

	return false
}

// SetValidationSchema gets a reference to the given map[string]map[string]interface{} and assigns it to the ValidationSchema field.
func (o *PropertyDataResponseObject) SetValidationSchema(v map[string]map[string]interface{}) {
	o.ValidationSchema = v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *PropertyDataResponseObject) GetValue() map[string]interface{} {
	if o == nil || IsNil(o.Value) {
		var ret map[string]interface{}
		return ret
	}
	return o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PropertyDataResponseObject) GetValueOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Value) {
		return map[string]interface{}{}, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *PropertyDataResponseObject) HasValue() bool {
	if o != nil && !IsNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given map[string]interface{} and assigns it to the Value field.
func (o *PropertyDataResponseObject) SetValue(v map[string]interface{}) {
	o.Value = v
}

func (o PropertyDataResponseObject) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PropertyDataResponseObject) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Classification) {
		toSerialize["classification"] = o.Classification
	}
	if !IsNil(o.DefaultValue) {
		toSerialize["defaultValue"] = o.DefaultValue
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.DynamicProperty) {
		toSerialize["dynamicProperty"] = o.DynamicProperty
	}
	toSerialize["entityType"] = o.EntityType
	if !IsNil(o.EntityTypeGuid) {
		toSerialize["entityTypeGuid"] = o.EntityTypeGuid
	}
	toSerialize["key"] = o.Key
	if !IsNil(o.ValidationSchema) {
		toSerialize["validationSchema"] = o.ValidationSchema
	}
	if !IsNil(o.Value) {
		toSerialize["value"] = o.Value
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *PropertyDataResponseObject) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"entityType",
		"key",
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

	varPropertyDataResponseObject := _PropertyDataResponseObject{}

	err = json.Unmarshal(data, &varPropertyDataResponseObject)

	if err != nil {
		return err
	}

	*o = PropertyDataResponseObject(varPropertyDataResponseObject)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "classification")
		delete(additionalProperties, "defaultValue")
		delete(additionalProperties, "description")
		delete(additionalProperties, "dynamicProperty")
		delete(additionalProperties, "entityType")
		delete(additionalProperties, "entityTypeGuid")
		delete(additionalProperties, "key")
		delete(additionalProperties, "validationSchema")
		delete(additionalProperties, "value")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePropertyDataResponseObject struct {
	value *PropertyDataResponseObject
	isSet bool
}

func (v NullablePropertyDataResponseObject) Get() *PropertyDataResponseObject {
	return v.value
}

func (v *NullablePropertyDataResponseObject) Set(val *PropertyDataResponseObject) {
	v.value = val
	v.isSet = true
}

func (v NullablePropertyDataResponseObject) IsSet() bool {
	return v.isSet
}

func (v *NullablePropertyDataResponseObject) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePropertyDataResponseObject(val *PropertyDataResponseObject) *NullablePropertyDataResponseObject {
	return &NullablePropertyDataResponseObject{value: val, isSet: true}
}

func (v NullablePropertyDataResponseObject) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePropertyDataResponseObject) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


