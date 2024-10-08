/*
FastReport Cloud

Testing TemplatesAPIService

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

func Test_gofrcloud_TemplatesAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test TemplatesAPIService TemplateFolderAndFileClearRecycleBin", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var subscriptionId string

		httpRes, err := apiClient.TemplatesAPI.TemplateFolderAndFileClearRecycleBin(context.Background(), subscriptionId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test TemplatesAPIService TemplateFolderAndFileCopyFiles", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var subscriptionId string

		httpRes, err := apiClient.TemplatesAPI.TemplateFolderAndFileCopyFiles(context.Background(), subscriptionId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test TemplatesAPIService TemplateFolderAndFileCountRecycleBinFoldersAndFiles", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var subscriptionId string

		resp, httpRes, err := apiClient.TemplatesAPI.TemplateFolderAndFileCountRecycleBinFoldersAndFiles(context.Background(), subscriptionId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test TemplatesAPIService TemplateFolderAndFileDeleteFiles", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var subscriptionId string

		httpRes, err := apiClient.TemplatesAPI.TemplateFolderAndFileDeleteFiles(context.Background(), subscriptionId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test TemplatesAPIService TemplateFolderAndFileGetCount", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id string

		resp, httpRes, err := apiClient.TemplatesAPI.TemplateFolderAndFileGetCount(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test TemplatesAPIService TemplateFolderAndFileGetFoldersAndFiles", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id string

		resp, httpRes, err := apiClient.TemplatesAPI.TemplateFolderAndFileGetFoldersAndFiles(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test TemplatesAPIService TemplateFolderAndFileGetRecycleBinFoldersAndFiles", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var subscriptionId string

		resp, httpRes, err := apiClient.TemplatesAPI.TemplateFolderAndFileGetRecycleBinFoldersAndFiles(context.Background(), subscriptionId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test TemplatesAPIService TemplateFolderAndFileMoveFiles", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var subscriptionId string

		httpRes, err := apiClient.TemplatesAPI.TemplateFolderAndFileMoveFiles(context.Background(), subscriptionId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test TemplatesAPIService TemplateFolderAndFileMoveFilesToBin", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var subscriptionId string

		httpRes, err := apiClient.TemplatesAPI.TemplateFolderAndFileMoveFilesToBin(context.Background(), subscriptionId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test TemplatesAPIService TemplateFolderAndFileRecoverAllFromRecycleBin", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var subscriptionId string

		httpRes, err := apiClient.TemplatesAPI.TemplateFolderAndFileRecoverAllFromRecycleBin(context.Background(), subscriptionId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test TemplatesAPIService TemplateFolderAndFileRecoverFiles", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var subscriptionId string

		httpRes, err := apiClient.TemplatesAPI.TemplateFolderAndFileRecoverFiles(context.Background(), subscriptionId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test TemplatesAPIService TemplateFoldersCalculateFolderSize", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id string

		resp, httpRes, err := apiClient.TemplatesAPI.TemplateFoldersCalculateFolderSize(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test TemplatesAPIService TemplateFoldersCopyFolder", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id string
		var folderId string

		resp, httpRes, err := apiClient.TemplatesAPI.TemplateFoldersCopyFolder(context.Background(), id, folderId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test TemplatesAPIService TemplateFoldersDeleteFolder", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id string

		httpRes, err := apiClient.TemplatesAPI.TemplateFoldersDeleteFolder(context.Background(), id).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test TemplatesAPIService TemplateFoldersExport", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id string

		resp, httpRes, err := apiClient.TemplatesAPI.TemplateFoldersExport(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test TemplatesAPIService TemplateFoldersGetBreadcrumbs", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id string

		resp, httpRes, err := apiClient.TemplatesAPI.TemplateFoldersGetBreadcrumbs(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test TemplatesAPIService TemplateFoldersGetFolder", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id string

		resp, httpRes, err := apiClient.TemplatesAPI.TemplateFoldersGetFolder(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test TemplatesAPIService TemplateFoldersGetFolders", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id string

		resp, httpRes, err := apiClient.TemplatesAPI.TemplateFoldersGetFolders(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test TemplatesAPIService TemplateFoldersGetFoldersCount", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id string

		resp, httpRes, err := apiClient.TemplatesAPI.TemplateFoldersGetFoldersCount(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test TemplatesAPIService TemplateFoldersGetOrCreate", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.TemplatesAPI.TemplateFoldersGetOrCreate(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test TemplatesAPIService TemplateFoldersGetPermissions", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id string

		resp, httpRes, err := apiClient.TemplatesAPI.TemplateFoldersGetPermissions(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test TemplatesAPIService TemplateFoldersGetRootFolder", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.TemplatesAPI.TemplateFoldersGetRootFolder(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test TemplatesAPIService TemplateFoldersMoveFolder", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id string
		var folderId string

		resp, httpRes, err := apiClient.TemplatesAPI.TemplateFoldersMoveFolder(context.Background(), id, folderId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test TemplatesAPIService TemplateFoldersMoveFolderToBin", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id string

		httpRes, err := apiClient.TemplatesAPI.TemplateFoldersMoveFolderToBin(context.Background(), id).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test TemplatesAPIService TemplateFoldersPostFolder", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id string

		resp, httpRes, err := apiClient.TemplatesAPI.TemplateFoldersPostFolder(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test TemplatesAPIService TemplateFoldersPrepare", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id string

		resp, httpRes, err := apiClient.TemplatesAPI.TemplateFoldersPrepare(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test TemplatesAPIService TemplateFoldersRecoverFolder", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id string

		httpRes, err := apiClient.TemplatesAPI.TemplateFoldersRecoverFolder(context.Background(), id).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test TemplatesAPIService TemplateFoldersRenameFolder", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id string

		resp, httpRes, err := apiClient.TemplatesAPI.TemplateFoldersRenameFolder(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test TemplatesAPIService TemplateFoldersUpdateIcon", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id string

		resp, httpRes, err := apiClient.TemplatesAPI.TemplateFoldersUpdateIcon(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test TemplatesAPIService TemplateFoldersUpdatePermissions", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id string

		httpRes, err := apiClient.TemplatesAPI.TemplateFoldersUpdatePermissions(context.Background(), id).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test TemplatesAPIService TemplateFoldersUpdateTags", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id string

		resp, httpRes, err := apiClient.TemplatesAPI.TemplateFoldersUpdateTags(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test TemplatesAPIService TemplatesCopyFile", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id string
		var folderId string

		resp, httpRes, err := apiClient.TemplatesAPI.TemplatesCopyFile(context.Background(), id, folderId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test TemplatesAPIService TemplatesCreateSharingKey", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id string

		resp, httpRes, err := apiClient.TemplatesAPI.TemplatesCreateSharingKey(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test TemplatesAPIService TemplatesDeleteFile", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id string

		httpRes, err := apiClient.TemplatesAPI.TemplatesDeleteFile(context.Background(), id).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test TemplatesAPIService TemplatesDeleteSharingKey", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id string
		var key string

		httpRes, err := apiClient.TemplatesAPI.TemplatesDeleteSharingKey(context.Background(), id, key).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test TemplatesAPIService TemplatesExport", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id string

		resp, httpRes, err := apiClient.TemplatesAPI.TemplatesExport(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test TemplatesAPIService TemplatesGetFile", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id string

		resp, httpRes, err := apiClient.TemplatesAPI.TemplatesGetFile(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test TemplatesAPIService TemplatesGetFileHistory", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id string

		resp, httpRes, err := apiClient.TemplatesAPI.TemplatesGetFileHistory(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test TemplatesAPIService TemplatesGetFilesCount", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id string

		resp, httpRes, err := apiClient.TemplatesAPI.TemplatesGetFilesCount(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test TemplatesAPIService TemplatesGetFilesList", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id string

		resp, httpRes, err := apiClient.TemplatesAPI.TemplatesGetFilesList(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test TemplatesAPIService TemplatesGetPermissions", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id string

		resp, httpRes, err := apiClient.TemplatesAPI.TemplatesGetPermissions(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test TemplatesAPIService TemplatesGetSharingKeys", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id string

		resp, httpRes, err := apiClient.TemplatesAPI.TemplatesGetSharingKeys(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test TemplatesAPIService TemplatesMoveFile", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id string
		var folderId string

		resp, httpRes, err := apiClient.TemplatesAPI.TemplatesMoveFile(context.Background(), id, folderId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test TemplatesAPIService TemplatesMoveFileToBin", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id string

		httpRes, err := apiClient.TemplatesAPI.TemplatesMoveFileToBin(context.Background(), id).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test TemplatesAPIService TemplatesPrepare", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id string

		resp, httpRes, err := apiClient.TemplatesAPI.TemplatesPrepare(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test TemplatesAPIService TemplatesRecoverFile", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id string

		httpRes, err := apiClient.TemplatesAPI.TemplatesRecoverFile(context.Background(), id).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test TemplatesAPIService TemplatesRenameFile", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id string

		resp, httpRes, err := apiClient.TemplatesAPI.TemplatesRenameFile(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test TemplatesAPIService TemplatesStaticPreview", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id string

		resp, httpRes, err := apiClient.TemplatesAPI.TemplatesStaticPreview(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test TemplatesAPIService TemplatesUpdateContent", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id string

		httpRes, err := apiClient.TemplatesAPI.TemplatesUpdateContent(context.Background(), id).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test TemplatesAPIService TemplatesUpdateContentV2", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id string

		httpRes, err := apiClient.TemplatesAPI.TemplatesUpdateContentV2(context.Background(), id).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test TemplatesAPIService TemplatesUpdateIcon", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id string

		resp, httpRes, err := apiClient.TemplatesAPI.TemplatesUpdateIcon(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test TemplatesAPIService TemplatesUpdatePermissions", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id string

		httpRes, err := apiClient.TemplatesAPI.TemplatesUpdatePermissions(context.Background(), id).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test TemplatesAPIService TemplatesUpdateTags", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id string

		resp, httpRes, err := apiClient.TemplatesAPI.TemplatesUpdateTags(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test TemplatesAPIService TemplatesUploadFile", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id string

		resp, httpRes, err := apiClient.TemplatesAPI.TemplatesUploadFile(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test TemplatesAPIService TemplatesUploadFileV2", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id string

		resp, httpRes, err := apiClient.TemplatesAPI.TemplatesUploadFileV2(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
