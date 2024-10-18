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

// checks if the CreateSubaccountRequestPayload type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateSubaccountRequestPayload{}

// CreateSubaccountRequestPayload Details of the new subaccount.
type CreateSubaccountRequestPayload struct {
	// Enables the subaccount to use beta services and applications. Not to be used in a production environment. Cannot be reverted once set. Any use of beta functionality is at the customer's own risk, and SAP shall not be liable for errors or damages caused by the use of beta features.
	BetaEnabled *bool `json:"betaEnabled,omitempty"`
	// (Deprecated) User-defined labels to assign, update, and remove as key-value pairs from the subaccount.  NOTE: This field is deprecated. \"customProperties\" supports only single values per key and is now replaced by the string array \"labels\", which supports multiple values per key. The \"customProperties\" field overwrites any labels with the same key. We recommend transitioning to the \"labels\" field. Do not include \"customProperties\" and \"labels\" together in the same request (customProperties will be ignored). 
	// Deprecated
	CustomProperties []AddPropertyRequestPayload `json:"customProperties,omitempty"`
	// A description of the subaccount for customer-facing UIs.
	Description *string `json:"description,omitempty"`
	// The display name of the subaccount for customer-facing UIs.
	DisplayName string `json:"displayName" validate:"regexp=^((?![\\/]).)*$"`
	// JSON array of up to 10 user-defined labels to assign as key-value pairs to the subaccount. Each label has a name (key) that you specify, and to which you can assign up to 10 corresponding values or leave empty.  Keys and values are each limited to 63 characters.  The keys and values of labels are case-sensitive. Try to avoid creating duplicate variants of the same keys or values with a different casing (example: \"myValue\" and \"MyValue\").  Example: {   \"Cost Center\": [\"19700626\"],   \"Department\": [\"Sales\"],   \"Contacts\": [\"name1@example.com\",\"name2@example.com\"],   \"EMEA\":[] }
	Labels *map[string][]string `json:"labels,omitempty"`
	// The origin of the subaccount creation. * <b>API:</b> Subaccount is created/updated by an admin using the REST APIs of the Accounts service.
	Origin *string `json:"origin,omitempty"`
	// The unique ID subaccount’s parent entity.
	ParentGUID *string `json:"parentGUID,omitempty"`
	// The region in which the subaccount was created.
	Region string `json:"region"`
	// List of additional admins of the subaccount. You can add users only from the same user base as you (example: sap.default or custom identity provider). Do not add yourself as you are assigned by default. Enter inline as a valid JSON array. Specify the admins by user IDs for Neo subaccounts and e-mails for multi-environment subaccounts.
	SubaccountAdmins []string `json:"subaccountAdmins,omitempty"`
	// The subdomain that becomes part of the path used to access the authorization tenant of the subaccount. Must be unique within the defined region. Use only letters (a-z), digits (0-9), and hyphens (not at start or end). Maximum length is 63 characters. Cannot be changed after the subaccount has been created. Does not apply to Neo subaccounts.
	Subdomain *string `json:"subdomain,omitempty"`
	// Whether the subaccount is used for production purposes. This flag can help your cloud operator to take appropriate action when handling incidents that are related to mission-critical accounts in production systems. Do not apply for subaccounts that are used for non-production purposes, such as development, testing, and demos. Applying this setting this does not modify the subaccount. * <b>NOT_USED_FOR_PRODUCTION:</b> Subaccount is not used for production purposes. * <b>USED_FOR_PRODUCTION:</b> Subaccount is used for production purposes.  
	UsedForProduction *string `json:"usedForProduction,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _CreateSubaccountRequestPayload CreateSubaccountRequestPayload

// NewCreateSubaccountRequestPayload instantiates a new CreateSubaccountRequestPayload object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateSubaccountRequestPayload(displayName string, region string) *CreateSubaccountRequestPayload {
	this := CreateSubaccountRequestPayload{}
	this.DisplayName = displayName
	this.Region = region
	return &this
}

// NewCreateSubaccountRequestPayloadWithDefaults instantiates a new CreateSubaccountRequestPayload object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateSubaccountRequestPayloadWithDefaults() *CreateSubaccountRequestPayload {
	this := CreateSubaccountRequestPayload{}
	return &this
}

// GetBetaEnabled returns the BetaEnabled field value if set, zero value otherwise.
func (o *CreateSubaccountRequestPayload) GetBetaEnabled() bool {
	if o == nil || IsNil(o.BetaEnabled) {
		var ret bool
		return ret
	}
	return *o.BetaEnabled
}

// GetBetaEnabledOk returns a tuple with the BetaEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateSubaccountRequestPayload) GetBetaEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.BetaEnabled) {
		return nil, false
	}
	return o.BetaEnabled, true
}

// HasBetaEnabled returns a boolean if a field has been set.
func (o *CreateSubaccountRequestPayload) HasBetaEnabled() bool {
	if o != nil && !IsNil(o.BetaEnabled) {
		return true
	}

	return false
}

// SetBetaEnabled gets a reference to the given bool and assigns it to the BetaEnabled field.
func (o *CreateSubaccountRequestPayload) SetBetaEnabled(v bool) {
	o.BetaEnabled = &v
}

// GetCustomProperties returns the CustomProperties field value if set, zero value otherwise.
// Deprecated
func (o *CreateSubaccountRequestPayload) GetCustomProperties() []AddPropertyRequestPayload {
	if o == nil || IsNil(o.CustomProperties) {
		var ret []AddPropertyRequestPayload
		return ret
	}
	return o.CustomProperties
}

// GetCustomPropertiesOk returns a tuple with the CustomProperties field value if set, nil otherwise
// and a boolean to check if the value has been set.
// Deprecated
func (o *CreateSubaccountRequestPayload) GetCustomPropertiesOk() ([]AddPropertyRequestPayload, bool) {
	if o == nil || IsNil(o.CustomProperties) {
		return nil, false
	}
	return o.CustomProperties, true
}

// HasCustomProperties returns a boolean if a field has been set.
func (o *CreateSubaccountRequestPayload) HasCustomProperties() bool {
	if o != nil && !IsNil(o.CustomProperties) {
		return true
	}

	return false
}

// SetCustomProperties gets a reference to the given []AddPropertyRequestPayload and assigns it to the CustomProperties field.
// Deprecated
func (o *CreateSubaccountRequestPayload) SetCustomProperties(v []AddPropertyRequestPayload) {
	o.CustomProperties = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *CreateSubaccountRequestPayload) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateSubaccountRequestPayload) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *CreateSubaccountRequestPayload) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *CreateSubaccountRequestPayload) SetDescription(v string) {
	o.Description = &v
}

// GetDisplayName returns the DisplayName field value
func (o *CreateSubaccountRequestPayload) GetDisplayName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value
// and a boolean to check if the value has been set.
func (o *CreateSubaccountRequestPayload) GetDisplayNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DisplayName, true
}

// SetDisplayName sets field value
func (o *CreateSubaccountRequestPayload) SetDisplayName(v string) {
	o.DisplayName = v
}

// GetLabels returns the Labels field value if set, zero value otherwise.
func (o *CreateSubaccountRequestPayload) GetLabels() map[string][]string {
	if o == nil || IsNil(o.Labels) {
		var ret map[string][]string
		return ret
	}
	return *o.Labels
}

// GetLabelsOk returns a tuple with the Labels field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateSubaccountRequestPayload) GetLabelsOk() (*map[string][]string, bool) {
	if o == nil || IsNil(o.Labels) {
		return nil, false
	}
	return o.Labels, true
}

// HasLabels returns a boolean if a field has been set.
func (o *CreateSubaccountRequestPayload) HasLabels() bool {
	if o != nil && !IsNil(o.Labels) {
		return true
	}

	return false
}

// SetLabels gets a reference to the given map[string][]string and assigns it to the Labels field.
func (o *CreateSubaccountRequestPayload) SetLabels(v map[string][]string) {
	o.Labels = &v
}

// GetOrigin returns the Origin field value if set, zero value otherwise.
func (o *CreateSubaccountRequestPayload) GetOrigin() string {
	if o == nil || IsNil(o.Origin) {
		var ret string
		return ret
	}
	return *o.Origin
}

// GetOriginOk returns a tuple with the Origin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateSubaccountRequestPayload) GetOriginOk() (*string, bool) {
	if o == nil || IsNil(o.Origin) {
		return nil, false
	}
	return o.Origin, true
}

// HasOrigin returns a boolean if a field has been set.
func (o *CreateSubaccountRequestPayload) HasOrigin() bool {
	if o != nil && !IsNil(o.Origin) {
		return true
	}

	return false
}

// SetOrigin gets a reference to the given string and assigns it to the Origin field.
func (o *CreateSubaccountRequestPayload) SetOrigin(v string) {
	o.Origin = &v
}

// GetParentGUID returns the ParentGUID field value if set, zero value otherwise.
func (o *CreateSubaccountRequestPayload) GetParentGUID() string {
	if o == nil || IsNil(o.ParentGUID) {
		var ret string
		return ret
	}
	return *o.ParentGUID
}

// GetParentGUIDOk returns a tuple with the ParentGUID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateSubaccountRequestPayload) GetParentGUIDOk() (*string, bool) {
	if o == nil || IsNil(o.ParentGUID) {
		return nil, false
	}
	return o.ParentGUID, true
}

// HasParentGUID returns a boolean if a field has been set.
func (o *CreateSubaccountRequestPayload) HasParentGUID() bool {
	if o != nil && !IsNil(o.ParentGUID) {
		return true
	}

	return false
}

// SetParentGUID gets a reference to the given string and assigns it to the ParentGUID field.
func (o *CreateSubaccountRequestPayload) SetParentGUID(v string) {
	o.ParentGUID = &v
}

// GetRegion returns the Region field value
func (o *CreateSubaccountRequestPayload) GetRegion() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Region
}

// GetRegionOk returns a tuple with the Region field value
// and a boolean to check if the value has been set.
func (o *CreateSubaccountRequestPayload) GetRegionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Region, true
}

// SetRegion sets field value
func (o *CreateSubaccountRequestPayload) SetRegion(v string) {
	o.Region = v
}

// GetSubaccountAdmins returns the SubaccountAdmins field value if set, zero value otherwise.
func (o *CreateSubaccountRequestPayload) GetSubaccountAdmins() []string {
	if o == nil || IsNil(o.SubaccountAdmins) {
		var ret []string
		return ret
	}
	return o.SubaccountAdmins
}

// GetSubaccountAdminsOk returns a tuple with the SubaccountAdmins field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateSubaccountRequestPayload) GetSubaccountAdminsOk() ([]string, bool) {
	if o == nil || IsNil(o.SubaccountAdmins) {
		return nil, false
	}
	return o.SubaccountAdmins, true
}

// HasSubaccountAdmins returns a boolean if a field has been set.
func (o *CreateSubaccountRequestPayload) HasSubaccountAdmins() bool {
	if o != nil && !IsNil(o.SubaccountAdmins) {
		return true
	}

	return false
}

// SetSubaccountAdmins gets a reference to the given []string and assigns it to the SubaccountAdmins field.
func (o *CreateSubaccountRequestPayload) SetSubaccountAdmins(v []string) {
	o.SubaccountAdmins = v
}

// GetSubdomain returns the Subdomain field value if set, zero value otherwise.
func (o *CreateSubaccountRequestPayload) GetSubdomain() string {
	if o == nil || IsNil(o.Subdomain) {
		var ret string
		return ret
	}
	return *o.Subdomain
}

// GetSubdomainOk returns a tuple with the Subdomain field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateSubaccountRequestPayload) GetSubdomainOk() (*string, bool) {
	if o == nil || IsNil(o.Subdomain) {
		return nil, false
	}
	return o.Subdomain, true
}

// HasSubdomain returns a boolean if a field has been set.
func (o *CreateSubaccountRequestPayload) HasSubdomain() bool {
	if o != nil && !IsNil(o.Subdomain) {
		return true
	}

	return false
}

// SetSubdomain gets a reference to the given string and assigns it to the Subdomain field.
func (o *CreateSubaccountRequestPayload) SetSubdomain(v string) {
	o.Subdomain = &v
}

// GetUsedForProduction returns the UsedForProduction field value if set, zero value otherwise.
func (o *CreateSubaccountRequestPayload) GetUsedForProduction() string {
	if o == nil || IsNil(o.UsedForProduction) {
		var ret string
		return ret
	}
	return *o.UsedForProduction
}

// GetUsedForProductionOk returns a tuple with the UsedForProduction field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateSubaccountRequestPayload) GetUsedForProductionOk() (*string, bool) {
	if o == nil || IsNil(o.UsedForProduction) {
		return nil, false
	}
	return o.UsedForProduction, true
}

// HasUsedForProduction returns a boolean if a field has been set.
func (o *CreateSubaccountRequestPayload) HasUsedForProduction() bool {
	if o != nil && !IsNil(o.UsedForProduction) {
		return true
	}

	return false
}

// SetUsedForProduction gets a reference to the given string and assigns it to the UsedForProduction field.
func (o *CreateSubaccountRequestPayload) SetUsedForProduction(v string) {
	o.UsedForProduction = &v
}

func (o CreateSubaccountRequestPayload) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateSubaccountRequestPayload) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.BetaEnabled) {
		toSerialize["betaEnabled"] = o.BetaEnabled
	}
	if !IsNil(o.CustomProperties) {
		toSerialize["customProperties"] = o.CustomProperties
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	toSerialize["displayName"] = o.DisplayName
	if !IsNil(o.Labels) {
		toSerialize["labels"] = o.Labels
	}
	if !IsNil(o.Origin) {
		toSerialize["origin"] = o.Origin
	}
	if !IsNil(o.ParentGUID) {
		toSerialize["parentGUID"] = o.ParentGUID
	}
	toSerialize["region"] = o.Region
	if !IsNil(o.SubaccountAdmins) {
		toSerialize["subaccountAdmins"] = o.SubaccountAdmins
	}
	if !IsNil(o.Subdomain) {
		toSerialize["subdomain"] = o.Subdomain
	}
	if !IsNil(o.UsedForProduction) {
		toSerialize["usedForProduction"] = o.UsedForProduction
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *CreateSubaccountRequestPayload) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"displayName",
		"region",
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

	varCreateSubaccountRequestPayload := _CreateSubaccountRequestPayload{}

	err = json.Unmarshal(data, &varCreateSubaccountRequestPayload)

	if err != nil {
		return err
	}

	*o = CreateSubaccountRequestPayload(varCreateSubaccountRequestPayload)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "betaEnabled")
		delete(additionalProperties, "customProperties")
		delete(additionalProperties, "description")
		delete(additionalProperties, "displayName")
		delete(additionalProperties, "labels")
		delete(additionalProperties, "origin")
		delete(additionalProperties, "parentGUID")
		delete(additionalProperties, "region")
		delete(additionalProperties, "subaccountAdmins")
		delete(additionalProperties, "subdomain")
		delete(additionalProperties, "usedForProduction")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCreateSubaccountRequestPayload struct {
	value *CreateSubaccountRequestPayload
	isSet bool
}

func (v NullableCreateSubaccountRequestPayload) Get() *CreateSubaccountRequestPayload {
	return v.value
}

func (v *NullableCreateSubaccountRequestPayload) Set(val *CreateSubaccountRequestPayload) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateSubaccountRequestPayload) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateSubaccountRequestPayload) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateSubaccountRequestPayload(val *CreateSubaccountRequestPayload) *NullableCreateSubaccountRequestPayload {
	return &NullableCreateSubaccountRequestPayload{value: val, isSet: true}
}

func (v NullableCreateSubaccountRequestPayload) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateSubaccountRequestPayload) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


