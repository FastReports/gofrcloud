/*
FastReport Cloud

Testing ConfigurationApiService

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

func Test_gofrcloud_ConfigurationApiService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test ConfigurationApiService ConfigurationGet", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.ConfigurationApi.ConfigurationGet(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
