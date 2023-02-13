/*
EJBCA REST Interface

Testing V1ConfigdumpApiService

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

func Test_ejbca_V1ConfigdumpApiService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test V1ConfigdumpApiService GetJsonConfigdump", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.V1ConfigdumpApi.GetJsonConfigdump(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test V1ConfigdumpApiService GetJsonConfigdumpForType", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var type_ string

		resp, httpRes, err := apiClient.V1ConfigdumpApi.GetJsonConfigdumpForType(context.Background(), type_).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test V1ConfigdumpApiService GetJsonConfigdumpForTypeAndSetting", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var type_ string
		var setting string

		resp, httpRes, err := apiClient.V1ConfigdumpApi.GetJsonConfigdumpForTypeAndSetting(context.Background(), type_, setting).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test V1ConfigdumpApiService GetZipExport", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.V1ConfigdumpApi.GetZipExport(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test V1ConfigdumpApiService PostJsonImport", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.V1ConfigdumpApi.PostJsonImport(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test V1ConfigdumpApiService PostZipImport", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.V1ConfigdumpApi.PostZipImport(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test V1ConfigdumpApiService Status4", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.V1ConfigdumpApi.Status4(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
