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

// checks if the CreateDirectoryRequestPayload type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateDirectoryRequestPayload{}

// CreateDirectoryRequestPayload Details of the directory to create.
type CreateDirectoryRequestPayload struct {
	// (Deprecated) User-defined labels to assign, update, and remove as key-value pairs from the directory.  NOTE: This field is deprecated. \"customProperties\" supports only single values per key and is now replaced by the string array \"labels\", which supports multiple values per key. The \"customProperties\" field overwrites any labels with the same key. We recommend transitioning to the \"labels\" field. Do not include \"customProperties\" and \"labels\" together in the same request (customProperties will be ignored). 
	// Deprecated
	CustomProperties []AddPropertyRequestPayload `json:"customProperties,omitempty"`
	// A description of the directory.
	Description *string `json:"description,omitempty"`
	// Additional admins of the directory. Applies only to directories that have the user authorization management feature enabled. Do not add yourself as you are assigned as a directory admin by default. Example: [\"admin1@example.com\", \"admin2@example.com\"]
	DirectoryAdmins []string `json:"directoryAdmins,omitempty"`
	// <b>The features to be enabled in the directory. The available features are:</b> - <b>DEFAULT</b>: (Mandatory) All directories provide the following basic features: (1) Group and filter subaccounts for reports and filters, (2) monitor usage and costs on a directory level (costs only available for contracts that use the consumption-based commercial model), and (3) set custom properties and tags to the directory for identification and reporting purposes. - <b>ENTITLEMENTS</b>: (Optional) Enables the assignment of a quota for services and applications to the directory from the global account quota for distribution to the subaccounts under this directory.  - <b>AUTHORIZATIONS</b>: (Optional) Allows you to assign users as administrators or viewers of this directory. You must apply this feature in combination with the ENTITLEMENTS feature.   IMPORTANT: Your multi-level account hierarchy can have more than one directory enabled with user authorization and/or entitlement management; however, only one directory in any directory path can have these features enabled. In other words, other directories above or below this directory in the same path can only have the default features specified. If you are not sure which features to enable, we recommend that you set only the default features, and then add features later on as they are needed.  <br/><b>Valid values:</b>  [DEFAULT] [DEFAULT,ENTITLEMENTS] [DEFAULT,ENTITLEMENTS,AUTHORIZATIONS]<br/>
	DirectoryFeatures []string `json:"directoryFeatures,omitempty"`
	// The display name of the directory.
	DisplayName string `json:"displayName" validate:"regexp=^((?![\\/]).)*$"`
	// JSON array of up to 10 user-defined labels to assign as key-value pairs to the directory. Each label has a name (key) that you specify, and to which you can assign up to 10 corresponding values or leave empty.  Keys and values are each limited to 63 characters.  Label keys and values are case-sensitive. Try to avoid creating duplicate variants of the same keys or values with a different casing (example: \"myValue\" and \"MyValue\").  Example:  {   \"Cost Center\": [\"19700626\"],   \"Department\": [\"Sales\"],   \"Contacts\": [\"name1@example.com\",\"name2@example.com\"],   \"EMEA\":[] } 
	Labels *map[string][]string `json:"labels,omitempty"`
	// Applies only to directories that have the user authorization management feature enabled. The subdomain becomes part of the path used to access the authorization tenant of the directory. Must be unique in the defined region. Use only letters (a-z), digits (0-9), and hyphens (not at start or end). Maximum length is 63 characters. Cannot be changed after the directory has been created.
	Subdomain *string `json:"subdomain,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _CreateDirectoryRequestPayload CreateDirectoryRequestPayload

// NewCreateDirectoryRequestPayload instantiates a new CreateDirectoryRequestPayload object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateDirectoryRequestPayload(displayName string) *CreateDirectoryRequestPayload {
	this := CreateDirectoryRequestPayload{}
	this.DisplayName = displayName
	return &this
}

// NewCreateDirectoryRequestPayloadWithDefaults instantiates a new CreateDirectoryRequestPayload object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateDirectoryRequestPayloadWithDefaults() *CreateDirectoryRequestPayload {
	this := CreateDirectoryRequestPayload{}
	return &this
}

// GetCustomProperties returns the CustomProperties field value if set, zero value otherwise.
// Deprecated
func (o *CreateDirectoryRequestPayload) GetCustomProperties() []AddPropertyRequestPayload {
	if o == nil || IsNil(o.CustomProperties) {
		var ret []AddPropertyRequestPayload
		return ret
	}
	return o.CustomProperties
}

// GetCustomPropertiesOk returns a tuple with the CustomProperties field value if set, nil otherwise
// and a boolean to check if the value has been set.
// Deprecated
func (o *CreateDirectoryRequestPayload) GetCustomPropertiesOk() ([]AddPropertyRequestPayload, bool) {
	if o == nil || IsNil(o.CustomProperties) {
		return nil, false
	}
	return o.CustomProperties, true
}

// HasCustomProperties returns a boolean if a field has been set.
func (o *CreateDirectoryRequestPayload) HasCustomProperties() bool {
	if o != nil && !IsNil(o.CustomProperties) {
		return true
	}

	return false
}

// SetCustomProperties gets a reference to the given []AddPropertyRequestPayload and assigns it to the CustomProperties field.
// Deprecated
func (o *CreateDirectoryRequestPayload) SetCustomProperties(v []AddPropertyRequestPayload) {
	o.CustomProperties = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *CreateDirectoryRequestPayload) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateDirectoryRequestPayload) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *CreateDirectoryRequestPayload) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *CreateDirectoryRequestPayload) SetDescription(v string) {
	o.Description = &v
}

// GetDirectoryAdmins returns the DirectoryAdmins field value if set, zero value otherwise.
func (o *CreateDirectoryRequestPayload) GetDirectoryAdmins() []string {
	if o == nil || IsNil(o.DirectoryAdmins) {
		var ret []string
		return ret
	}
	return o.DirectoryAdmins
}

// GetDirectoryAdminsOk returns a tuple with the DirectoryAdmins field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateDirectoryRequestPayload) GetDirectoryAdminsOk() ([]string, bool) {
	if o == nil || IsNil(o.DirectoryAdmins) {
		return nil, false
	}
	return o.DirectoryAdmins, true
}

// HasDirectoryAdmins returns a boolean if a field has been set.
func (o *CreateDirectoryRequestPayload) HasDirectoryAdmins() bool {
	if o != nil && !IsNil(o.DirectoryAdmins) {
		return true
	}

	return false
}

// SetDirectoryAdmins gets a reference to the given []string and assigns it to the DirectoryAdmins field.
func (o *CreateDirectoryRequestPayload) SetDirectoryAdmins(v []string) {
	o.DirectoryAdmins = v
}

// GetDirectoryFeatures returns the DirectoryFeatures field value if set, zero value otherwise.
func (o *CreateDirectoryRequestPayload) GetDirectoryFeatures() []string {
	if o == nil || IsNil(o.DirectoryFeatures) {
		var ret []string
		return ret
	}
	return o.DirectoryFeatures
}

// GetDirectoryFeaturesOk returns a tuple with the DirectoryFeatures field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateDirectoryRequestPayload) GetDirectoryFeaturesOk() ([]string, bool) {
	if o == nil || IsNil(o.DirectoryFeatures) {
		return nil, false
	}
	return o.DirectoryFeatures, true
}

// HasDirectoryFeatures returns a boolean if a field has been set.
func (o *CreateDirectoryRequestPayload) HasDirectoryFeatures() bool {
	if o != nil && !IsNil(o.DirectoryFeatures) {
		return true
	}

	return false
}

// SetDirectoryFeatures gets a reference to the given []string and assigns it to the DirectoryFeatures field.
func (o *CreateDirectoryRequestPayload) SetDirectoryFeatures(v []string) {
	o.DirectoryFeatures = v
}

// GetDisplayName returns the DisplayName field value
func (o *CreateDirectoryRequestPayload) GetDisplayName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value
// and a boolean to check if the value has been set.
func (o *CreateDirectoryRequestPayload) GetDisplayNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DisplayName, true
}

// SetDisplayName sets field value
func (o *CreateDirectoryRequestPayload) SetDisplayName(v string) {
	o.DisplayName = v
}

// GetLabels returns the Labels field value if set, zero value otherwise.
func (o *CreateDirectoryRequestPayload) GetLabels() map[string][]string {
	if o == nil || IsNil(o.Labels) {
		var ret map[string][]string
		return ret
	}
	return *o.Labels
}

// GetLabelsOk returns a tuple with the Labels field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateDirectoryRequestPayload) GetLabelsOk() (*map[string][]string, bool) {
	if o == nil || IsNil(o.Labels) {
		return nil, false
	}
	return o.Labels, true
}

// HasLabels returns a boolean if a field has been set.
func (o *CreateDirectoryRequestPayload) HasLabels() bool {
	if o != nil && !IsNil(o.Labels) {
		return true
	}

	return false
}

// SetLabels gets a reference to the given map[string][]string and assigns it to the Labels field.
func (o *CreateDirectoryRequestPayload) SetLabels(v map[string][]string) {
	o.Labels = &v
}

// GetSubdomain returns the Subdomain field value if set, zero value otherwise.
func (o *CreateDirectoryRequestPayload) GetSubdomain() string {
	if o == nil || IsNil(o.Subdomain) {
		var ret string
		return ret
	}
	return *o.Subdomain
}

// GetSubdomainOk returns a tuple with the Subdomain field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateDirectoryRequestPayload) GetSubdomainOk() (*string, bool) {
	if o == nil || IsNil(o.Subdomain) {
		return nil, false
	}
	return o.Subdomain, true
}

// HasSubdomain returns a boolean if a field has been set.
func (o *CreateDirectoryRequestPayload) HasSubdomain() bool {
	if o != nil && !IsNil(o.Subdomain) {
		return true
	}

	return false
}

// SetSubdomain gets a reference to the given string and assigns it to the Subdomain field.
func (o *CreateDirectoryRequestPayload) SetSubdomain(v string) {
	o.Subdomain = &v
}

func (o CreateDirectoryRequestPayload) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateDirectoryRequestPayload) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CustomProperties) {
		toSerialize["customProperties"] = o.CustomProperties
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.DirectoryAdmins) {
		toSerialize["directoryAdmins"] = o.DirectoryAdmins
	}
	if !IsNil(o.DirectoryFeatures) {
		toSerialize["directoryFeatures"] = o.DirectoryFeatures
	}
	toSerialize["displayName"] = o.DisplayName
	if !IsNil(o.Labels) {
		toSerialize["labels"] = o.Labels
	}
	if !IsNil(o.Subdomain) {
		toSerialize["subdomain"] = o.Subdomain
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *CreateDirectoryRequestPayload) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"displayName",
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

	varCreateDirectoryRequestPayload := _CreateDirectoryRequestPayload{}

	err = json.Unmarshal(data, &varCreateDirectoryRequestPayload)

	if err != nil {
		return err
	}

	*o = CreateDirectoryRequestPayload(varCreateDirectoryRequestPayload)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "customProperties")
		delete(additionalProperties, "description")
		delete(additionalProperties, "directoryAdmins")
		delete(additionalProperties, "directoryFeatures")
		delete(additionalProperties, "displayName")
		delete(additionalProperties, "labels")
		delete(additionalProperties, "subdomain")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCreateDirectoryRequestPayload struct {
	value *CreateDirectoryRequestPayload
	isSet bool
}

func (v NullableCreateDirectoryRequestPayload) Get() *CreateDirectoryRequestPayload {
	return v.value
}

func (v *NullableCreateDirectoryRequestPayload) Set(val *CreateDirectoryRequestPayload) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateDirectoryRequestPayload) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateDirectoryRequestPayload) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateDirectoryRequestPayload(val *CreateDirectoryRequestPayload) *NullableCreateDirectoryRequestPayload {
	return &NullableCreateDirectoryRequestPayload{value: val, isSet: true}
}

func (v NullableCreateDirectoryRequestPayload) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateDirectoryRequestPayload) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


