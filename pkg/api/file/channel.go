package file

import (
	"github.com/dezhishen/onebot-sdk/pkg/channel"
	"github.com/dezhishen/onebot-sdk/pkg/model"
)

type ChannelApiFileClient struct {
	channel.ApiChannel
}

func NewChannelApiFileClient(channel channel.ApiChannel) (OnebotApiFileClient, error) {
	return &ChannelApiFileClient{
		channel,
	}, nil
}

// 上传群文件
// upload_group_file
// groupId 群号
// file 文件路径
// name 文件名
// folder 群文件目录ID
func (cli *ChannelApiFileClient) UploadGroupFile(groupId int64, file string, name string, folder string) error {
	return cli.PostByRequest(
		API_UPLOAD_GROUP_FILE,
		map[string]interface{}{
			"group_id": groupId,
			"file":     file,
			"name":     name,
			"folder":   folder,
		},
	)
}

// 删除群文件
// delete_group_file
// groupId 群号
// file_id 文件ID 参考 File 对象
// busid 文件类型 参考 File 对象
func (cli *ChannelApiFileClient) DeleteGroupFile(groupId int64, fileId string, busid int32) error {
	return cli.PostByRequest(
		API_DELETE_GROUP_FILE,
		map[string]interface{}{
			"group_id": groupId,
			"file_id":  fileId,
			"busid":    busid,
		},
	)
}

// 创建群文件目录
// create_group_file_folder
// groupId 群号
// name 目录名
// parentId 父目录ID
func (cli *ChannelApiFileClient) CreateGroupFileFolder(groupId int64, name string, parentId string) error {
	return cli.PostByRequest(
		API_CREATE_GROUP_FILE_FOLDER,
		map[string]interface{}{
			"group_id": groupId,
			"name":     name,
			"parent":   parentId,
		},
	)
}

// 删除群文件目录
// delete_group_folder
// groupId 群号
// folder_id 文件夹ID 参考 Folder 对象
func (cli *ChannelApiFileClient) DeleteGroupFolder(groupId int64, folderId string) error {
	return cli.PostByRequest(
		API_DELETE_GROUP_FOLDER,
		map[string]interface{}{
			"group_id":  groupId,
			"folder_id": folderId,
		},
	)
}

// 获取群文件系统信息
// get_group_file_system_info
// groupId 群号
func (cli *ChannelApiFileClient) GetGroupFileSystemInfo(groupId int64) (*model.GroupFileSystemInfoResult, error) {
	var result model.GroupFileSystemInfoResult
	if err := cli.PostByRequestForResult(
		API_GET_GROUP_FILE_SYSTEM_INFO,
		map[string]interface{}{
			"group_id": groupId,
		},
		&result,
	); err != nil {
		return nil, err
	}
	return &result, nil
}

// 获取群根目录文件列表
// get_group_root_files
// groupId 群号
func (cli *ChannelApiFileClient) GetGroupRootFiles(groupId int64) (*model.GroupFilesResult, error) {
	var result model.GroupFilesResult
	if err := cli.PostByRequestForResult(
		API_GET_GROUP_ROOT_FILES,
		map[string]interface{}{
			"group_id": groupId,
		},
		&result,
	); err != nil {
		return nil, err
	}
	return &result, nil
}

// get_group_files_by_folder
// 获取群文件列表
// groupId 群号
// folder_id 文件夹ID
func (cli *ChannelApiFileClient) GetGroupFilesByFolder(groupId int64, folderId string) (*model.GroupFilesResult, error) {
	var result model.GroupFilesResult
	if err := cli.PostByRequestForResult(
		API_GET_GROUP_FILES_BY_FOLDER,
		map[string]interface{}{
			"group_id":  groupId,
			"folder_id": folderId,
		},
		&result,
	); err != nil {
		return nil, err
	}
	return &result, nil
}

// 获取群文件资源链接
// get_group_file_url
// groupId 群号
// file_id 文件ID 参考 File 对象
// busid 文件类型 参考 File 对象
func (cli *ChannelApiFileClient) GetGroupFileUrl(groupId int64, fileId string, busid int32) (*model.FileUrlResult, error) {
	var result model.FileUrlResult
	if err := cli.PostByRequestForResult(
		API_GET_GROUP_FILE_URL,
		map[string]interface{}{
			"group_id": groupId,
			"file_id":  fileId,
			"busid":    busid,
		},
		&result,
	); err != nil {
		return nil, err
	}
	return &result, nil
}

// 上传好友文件
// upload_private_file
// userId 用户ID
// file 文件路径
// name 文件名
func (cli *ChannelApiFileClient) UploadPrivateFile(userId int64, file string, name string) error {
	return cli.PostByRequest(
		API_UPLOAD_PRIVATE_FILE,
		map[string]interface{}{
			"user_id": userId,
			"file":    file,
			"name":    name,
		},
	)
}
