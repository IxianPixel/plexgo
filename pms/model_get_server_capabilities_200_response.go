/*
Plex-API

An Open API Spec for interacting with Plex.tv and Plex Servers

API version: 0.0.3
Contact: Lukeslakemail@gmail.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package pms

import (
	"encoding/json"
)

// checks if the GetServerCapabilities200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetServerCapabilities200Response{}

// GetServerCapabilities200Response struct for GetServerCapabilities200Response
type GetServerCapabilities200Response struct {
	MediaContainer *GetServerCapabilities200ResponseMediaContainer `json:"MediaContainer,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _GetServerCapabilities200Response GetServerCapabilities200Response

// NewGetServerCapabilities200Response instantiates a new GetServerCapabilities200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetServerCapabilities200Response() *GetServerCapabilities200Response {
	this := GetServerCapabilities200Response{}
	return &this
}

// NewGetServerCapabilities200ResponseWithDefaults instantiates a new GetServerCapabilities200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetServerCapabilities200ResponseWithDefaults() *GetServerCapabilities200Response {
	this := GetServerCapabilities200Response{}
	return &this
}

// GetMediaContainer returns the MediaContainer field value if set, zero value otherwise.
func (o *GetServerCapabilities200Response) GetMediaContainer() GetServerCapabilities200ResponseMediaContainer {
	if o == nil || isNil(o.MediaContainer) {
		var ret GetServerCapabilities200ResponseMediaContainer
		return ret
	}
	return *o.MediaContainer
}

// GetMediaContainerOk returns a tuple with the MediaContainer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetServerCapabilities200Response) GetMediaContainerOk() (*GetServerCapabilities200ResponseMediaContainer, bool) {
	if o == nil || isNil(o.MediaContainer) {
		return nil, false
	}
	return o.MediaContainer, true
}

// HasMediaContainer returns a boolean if a field has been set.
func (o *GetServerCapabilities200Response) HasMediaContainer() bool {
	if o != nil && !isNil(o.MediaContainer) {
		return true
	}

	return false
}

// SetMediaContainer gets a reference to the given GetServerCapabilities200ResponseMediaContainer and assigns it to the MediaContainer field.
func (o *GetServerCapabilities200Response) SetMediaContainer(v GetServerCapabilities200ResponseMediaContainer) {
	o.MediaContainer = &v
}

func (o GetServerCapabilities200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetServerCapabilities200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.MediaContainer) {
		toSerialize["MediaContainer"] = o.MediaContainer
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *GetServerCapabilities200Response) UnmarshalJSON(bytes []byte) (err error) {
	varGetServerCapabilities200Response := _GetServerCapabilities200Response{}

	if err = json.Unmarshal(bytes, &varGetServerCapabilities200Response); err == nil {
		*o = GetServerCapabilities200Response(varGetServerCapabilities200Response)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "MediaContainer")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableGetServerCapabilities200Response struct {
	value *GetServerCapabilities200Response
	isSet bool
}

func (v NullableGetServerCapabilities200Response) Get() *GetServerCapabilities200Response {
	return v.value
}

func (v *NullableGetServerCapabilities200Response) Set(val *GetServerCapabilities200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGetServerCapabilities200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGetServerCapabilities200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetServerCapabilities200Response(val *GetServerCapabilities200Response) *NullableGetServerCapabilities200Response {
	return &NullableGetServerCapabilities200Response{value: val, isSet: true}
}

func (v NullableGetServerCapabilities200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetServerCapabilities200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


