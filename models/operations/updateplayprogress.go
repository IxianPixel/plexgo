// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
)

type UpdatePlayProgressRequest struct {
	// the media key
	Key string `queryParam:"style=form,explode=true,name=key"`
	// The time, in milliseconds, used to set the media playback progress.
	Time float64 `queryParam:"style=form,explode=true,name=time"`
	// The playback state of the media item.
	State string `queryParam:"style=form,explode=true,name=state"`
}

func (o *UpdatePlayProgressRequest) GetKey() string {
	if o == nil {
		return ""
	}
	return o.Key
}

func (o *UpdatePlayProgressRequest) GetTime() float64 {
	if o == nil {
		return 0.0
	}
	return o.Time
}

func (o *UpdatePlayProgressRequest) GetState() string {
	if o == nil {
		return ""
	}
	return o.State
}

type UpdatePlayProgressResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *UpdatePlayProgressResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *UpdatePlayProgressResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *UpdatePlayProgressResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}