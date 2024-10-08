/*
FastReport Cloud

Testing GroupUsersAPIService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package gofrcloud

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	openapiclient "github.com/fastreports/gofrcloud"
)

func Test_gofrcloud_GroupUsersAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test GroupUsersAPIService GroupUsersAddUserToGroup", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id string
		var userId string

		httpRes, err := apiClient.GroupUsersAPI.GroupUsersAddUserToGroup(context.Background(), id, userId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test GroupUsersAPIService GroupUsersGetUsersInGroup", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id string

		resp, httpRes, err := apiClient.GroupUsersAPI.GroupUsersGetUsersInGroup(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test GroupUsersAPIService GroupUsersLeaveFromGroup", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id string

		httpRes, err := apiClient.GroupUsersAPI.GroupUsersLeaveFromGroup(context.Background(), id).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test GroupUsersAPIService GroupUsersRemoveFromGroup", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id string
		var userId string

		httpRes, err := apiClient.GroupUsersAPI.GroupUsersRemoveFromGroup(context.Background(), id, userId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
