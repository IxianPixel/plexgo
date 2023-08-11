/*
Plex-API

Testing LibraryApiService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package pms

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	openapiclient "github.com/lukehagar/plexgo"
)

func Test_pms_LibraryApiService(t *testing.T) {

	configuration := openapiclient.NewDefaultConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test LibraryApiService DeleteLibrary", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var sectionId interface{}

		httpRes, err := apiClient.PMS.LibraryApi.DeleteLibrary(context.Background(), sectionId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test LibraryApiService GetCommonLibraryItems", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var sectionId interface{}

		httpRes, err := apiClient.PMS.LibraryApi.GetCommonLibraryItems(context.Background(), sectionId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test LibraryApiService GetFileHash", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		httpRes, err := apiClient.PMS.LibraryApi.GetFileHash(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test LibraryApiService GetLatestLibraryItems", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var sectionId interface{}

		httpRes, err := apiClient.PMS.LibraryApi.GetLatestLibraryItems(context.Background(), sectionId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test LibraryApiService GetLibraries", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		httpRes, err := apiClient.PMS.LibraryApi.GetLibraries(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test LibraryApiService GetLibrary", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var sectionId interface{}

		httpRes, err := apiClient.PMS.LibraryApi.GetLibrary(context.Background(), sectionId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test LibraryApiService GetLibraryItems", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var sectionId interface{}

		httpRes, err := apiClient.PMS.LibraryApi.GetLibraryItems(context.Background(), sectionId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test LibraryApiService GetMetadata", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var ratingKey interface{}

		httpRes, err := apiClient.PMS.LibraryApi.GetMetadata(context.Background(), ratingKey).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test LibraryApiService GetMetadataChildren", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var ratingKey interface{}

		httpRes, err := apiClient.PMS.LibraryApi.GetMetadataChildren(context.Background(), ratingKey).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test LibraryApiService GetOnDeck", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.PMS.LibraryApi.GetOnDeck(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test LibraryApiService GetRecentlyAdded", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.PMS.LibraryApi.GetRecentlyAdded(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test LibraryApiService RefreshLibrary", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var sectionId interface{}

		httpRes, err := apiClient.PMS.LibraryApi.RefreshLibrary(context.Background(), sectionId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
