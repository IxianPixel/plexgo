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

// checks if the GetServerIdentity200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetServerIdentity200Response{}

// GetServerIdentity200Response struct for GetServerIdentity200Response
type GetServerIdentity200Response struct {
	MediaContainer *GetServerIdentity200ResponseMediaContainer `json:"MediaContainer,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _GetServerIdentity200Response GetServerIdentity200Response

// NewGetServerIdentity200Response instantiates a new GetServerIdentity200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetServerIdentity200Response() *GetServerIdentity200Response {
	this := GetServerIdentity200Response{}
	return &this
}

// NewGetServerIdentity200ResponseWithDefaults instantiates a new GetServerIdentity200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetServerIdentity200ResponseWithDefaults() *GetServerIdentity200Response {
	this := GetServerIdentity200Response{}
	return &this
}

// GetMediaContainer returns the MediaContainer field value if set, zero value otherwise.
func (o *GetServerIdentity200Response) GetMediaContainer() GetServerIdentity200ResponseMediaContainer {
	if o == nil || isNil(o.MediaContainer) {
		var ret GetServerIdentity200ResponseMediaContainer
		return ret
	}
	return *o.MediaContainer
}

// GetMediaContainerOk returns a tuple with the MediaContainer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetServerIdentity200Response) GetMediaContainerOk() (*GetServerIdentity200ResponseMediaContainer, bool) {
	if o == nil || isNil(o.MediaContainer) {
		return nil, false
	}
	return o.MediaContainer, true
}

// HasMediaContainer returns a boolean if a field has been set.
func (o *GetServerIdentity200Response) HasMediaContainer() bool {
	if o != nil && !isNil(o.MediaContainer) {
		return true
	}

	return false
}

// SetMediaContainer gets a reference to the given GetServerIdentity200ResponseMediaContainer and assigns it to the MediaContainer field.
func (o *GetServerIdentity200Response) SetMediaContainer(v GetServerIdentity200ResponseMediaContainer) {
	o.MediaContainer = &v
}

func (o GetServerIdentity200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetServerIdentity200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.MediaContainer) {
		toSerialize["MediaContainer"] = o.MediaContainer
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *GetServerIdentity200Response) UnmarshalJSON(bytes []byte) (err error) {
	varGetServerIdentity200Response := _GetServerIdentity200Response{}

	if err = json.Unmarshal(bytes, &varGetServerIdentity200Response); err == nil {
		*o = GetServerIdentity200Response(varGetServerIdentity200Response)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "MediaContainer")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableGetServerIdentity200Response struct {
	value *GetServerIdentity200Response
	isSet bool
}

func (v NullableGetServerIdentity200Response) Get() *GetServerIdentity200Response {
	return v.value
}

func (v *NullableGetServerIdentity200Response) Set(val *GetServerIdentity200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGetServerIdentity200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGetServerIdentity200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetServerIdentity200Response(val *GetServerIdentity200Response) *NullableGetServerIdentity200Response {
	return &NullableGetServerIdentity200Response{value: val, isSet: true}
}

func (v NullableGetServerIdentity200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetServerIdentity200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


