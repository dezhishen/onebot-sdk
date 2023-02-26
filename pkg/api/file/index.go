package file

import (
	"github.com/dezhishen/onebot-sdk/pkg/model"
)

type OnebotApiFileClient interface {

	// upload_group_file
	// 上传群文件
	// groupId 群号
	// file 文件路径
	// name 文件名
	// folder 群文件目录ID
	UploadGroupFile(groupId int64, file string, name string, folder string) error

	// delete_group_file
	// 删除群文件
	// groupId 群号
	// file_id 文件ID 参考 File 对象
	// busid 文件类型 参考 File 对象
	DeleteGroupFile(groupId int64, fileId string, busid int32) error

	// create_group_file_folder
	// 创建群文件目录
	// groupId 群号
	// name 目录名
	// parentId 父目录ID
	CreateGroupFileFolder(groupId int64, name string, parentId string) error

	// delete_group_folder
	// 删除群文件目录
	// groupId 群号
	// folder_id 文件夹ID 参考 Folder 对象
	DeleteGroupFolder(groupId int64, folderId string) error

	// get_group_file_system_info
	// 获取群文件系统信息
	// groupId 群号
	GetGroupFileSystemInfo(groupId int64) (*model.GroupFileSystemInfoResult, error)

	// get_group_root_files
	// 获取群根目录文件列表
	// groupId 群号
	GetGroupRootFiles(groupId int64) (*model.GroupFilesResult, error)

	// get_group_files_by_folder
	// 获取群文件列表
	// groupId 群号
	// folder_id 文件夹ID
	GetGroupFilesByFolder(groupId int64, folderId string) (*model.GroupFilesResult, error)
	// get_group_file_url
	// 获取群文件资源链接
	// groupId 群号
	// file_id 文件ID 参考 File 对象
	// busid 文件类型 参考 File 对象
	GetGroupFileUrl(groupId int64, fileId string, busid int32) (*model.FileUrlResult, error)

	// upload_private_file
	// 上传好友文件
	// userId 用户ID
	// file 文件路径
	// name 文件名
	UploadPrivateFile(userId int64, file string, name string) error
}
