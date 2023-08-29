/*
Strava API v3

Testing SegmentsAPIService

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

func Test_strava3golang_SegmentsAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test SegmentsAPIService ExploreSegments", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.SegmentsAPI.ExploreSegments(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SegmentsAPIService GetLoggedInAthleteStarredSegments", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.SegmentsAPI.GetLoggedInAthleteStarredSegments(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SegmentsAPIService GetSegmentById", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id int64

		resp, httpRes, err := apiClient.SegmentsAPI.GetSegmentById(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SegmentsAPIService StarSegment", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id int64

		resp, httpRes, err := apiClient.SegmentsAPI.StarSegment(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
