package file

import (
	"github.com/dezhishen/onebot-sdk/pkg/model"
)

type OnebotApiFileClient interface {

	// 上传群文件
	// upload_group_file
	// groupId 群号
	// file 文件路径
	// name 文件名
	// folder 群文件目录ID
	UploadGroupFile(groupId int64, file string, name string, folder string) error

	// 删除群文件
	// delete_group_file
	// groupId 群号
	// file_id 文件ID 参考 File 对象
	// busid 文件类型 参考 File 对象
	DeleteGroupFile(groupId int64, fileId string, busid int32) error

	// 创建群文件目录
	// create_group_file_folder
	// groupId 群号
	// name 目录名
	// parentId 父目录ID
	CreateGroupFileFolder(groupId int64, name string, parentId string) error

	// 删除群文件目录
	// delete_group_folder
	// groupId 群号
	// folder_id 文件夹ID 参考 Folder 对象
	DeleteGroupFolder(groupId int64, folderId string) error

	// 获取群文件系统信息
	// get_group_file_system_info
	// groupId 群号
	GetGroupFileSystemInfo(groupId int64) (*model.GroupFileSystemInfoResult, error)

	// 获取群根目录文件列表
	// get_group_root_files
	// groupId 群号
	GetGroupRootFiles(groupId int64) (*model.GroupFilesResult, error)

	// get_group_files_by_folder
	// 获取群文件列表
	// groupId 群号
	// folder_id 文件夹ID
	GetGroupFilesByFolder(groupId int64, folderId string) (*model.GroupFilesResult, error)

	// 获取群文件资源链接
	// get_group_file_url
	// groupId 群号
	// file_id 文件ID 参考 File 对象
	// busid 文件类型 参考 File 对象
	GetGroupFileUrl(groupId int64, fileId string, busid int32) (*model.FileUrlResult, error)

	// 上传好友文件
	// upload_private_file
	// userId 用户ID
	// file 文件路径
	// name 文件名
	UploadPrivateFile(userId int64, file string, name string) error
}
