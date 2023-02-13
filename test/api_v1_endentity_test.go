/*
EJBCA REST Interface

Testing V1EndentityApiService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package ejbca

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	openapiclient "github.com/Keyfactor/ejbca-go-client-sdk"
)

func Test_ejbca_V1EndentityApiService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test V1EndentityApiService Add", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		httpRes, err := apiClient.V1EndentityApi.Add(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test V1EndentityApiService Delete", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var endentityName string

		httpRes, err := apiClient.V1EndentityApi.Delete(context.Background(), endentityName).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test V1EndentityApiService Revoke", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var endentityName string

		httpRes, err := apiClient.V1EndentityApi.Revoke(context.Background(), endentityName).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test V1EndentityApiService Search", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.V1EndentityApi.Search(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test V1EndentityApiService Setstatus", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var endentityName string

		httpRes, err := apiClient.V1EndentityApi.Setstatus(context.Background(), endentityName).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test V1EndentityApiService Status6", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.V1EndentityApi.Status6(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
