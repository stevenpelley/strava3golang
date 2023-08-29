/*
Strava API v3

Testing UploadsAPIService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package strava3golang

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	openapiclient "github.com/stevenpelley/strava3golang"
)

func Test_strava3golang_UploadsAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test UploadsAPIService CreateUpload", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.UploadsAPI.CreateUpload(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test UploadsAPIService GetUploadById", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var uploadId int64

		resp, httpRes, err := apiClient.UploadsAPI.GetUploadById(context.Background(), uploadId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}