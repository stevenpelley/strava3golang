/*
Strava API v3

Testing ClubsAPIService

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

func Test_strava3golang_ClubsAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test ClubsAPIService GetClubActivitiesById", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id int64

		resp, httpRes, err := apiClient.ClubsAPI.GetClubActivitiesById(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ClubsAPIService GetClubAdminsById", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id int64

		resp, httpRes, err := apiClient.ClubsAPI.GetClubAdminsById(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ClubsAPIService GetClubById", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id int64

		resp, httpRes, err := apiClient.ClubsAPI.GetClubById(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ClubsAPIService GetClubMembersById", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id int64

		resp, httpRes, err := apiClient.ClubsAPI.GetClubMembersById(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ClubsAPIService GetLoggedInAthleteClubs", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.ClubsAPI.GetLoggedInAthleteClubs(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
