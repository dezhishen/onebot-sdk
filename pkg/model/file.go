package model

type GroupFileSystemInfo struct {
	// file_count	int32	文件总数
	FileCount int32 `json:"file_count"`
	// limit_count	int32	文件上限
	LimitCount int32 `json:"limit_count"`
	// used_space	int64	已使用空间
	UsedSpace int64 `json:"used_space"`
	// total_space	int64	空间上限
	TotalSpace int64 `json:"total_space"`
}

func (g *GroupFileSystemInfo) ToGRPC() *GroupFileSystemInfoGRPC {
	return &GroupFileSystemInfoGRPC{
		FileCount:  g.FileCount,
		LimitCount: g.LimitCount,
		UsedSpace:  g.UsedSpace,
		TotalSpace: g.TotalSpace,
	}
}

// ToStruct
func (g *GroupFileSystemInfoGRPC) ToStruct() *GroupFileSystemInfo {
	return &GroupFileSystemInfo{
		FileCount:  g.FileCount,
		LimitCount: g.LimitCount,
		UsedSpace:  g.UsedSpace,
		TotalSpace: g.TotalSpace,
	}
}

type GroupFileSystemInfoResult struct {
	Data    *GroupFileSystemInfo `json:"data"`
	Retcode int64                `json:"retcode"`
	Status  string               `json:"status"`
	Msg     string               `json:"msg"`
	Wording string               `json:"wording"`
}

func (g *GroupFileSystemInfoResult) ToGRPC() *GroupFileSystemInfoResultGRPC {
	result := &GroupFileSystemInfoResultGRPC{
		// Data:    g.Data.ToGRPC(),
		Retcode: g.Retcode,
		Status:  g.Status,
		Msg:     g.Msg,
		Wording: g.Wording,
	}
	if g.Data != nil {
		result.Data = g.Data.ToGRPC()
	}
	return result
}

// ToStruct

func (g *GroupFileSystemInfoResultGRPC) ToStruct() *GroupFileSystemInfoResult {
	result := &GroupFileSystemInfoResult{
		Retcode: g.Retcode,
		Status:  g.Status,
		Msg:     g.Msg,
		Wording: g.Wording,
	}
	if g.Data != nil {
		result.Data = g.Data.ToStruct()
	}
	return result
}

type File struct {
	// group_id	int32	群号
	GroupId int32 `json:"group_id"`
	// file_id	string	文件ID
	FileId string `json:"file_id"`
	// file_name	string	文件名
	FileName string `json:"file_name"`
	// busid	int32	文件类型
	Busid int32 `json:"busid"`
	// file_size	int64	文件大小
	FileSize int64 `json:"file_size"`
	// upload_time	int64	上传时间
	UploadTime int64 `json:"upload_time"`
	// dead_time	int64	过期时间,永久文件恒为0
	DeadTime int64 `json:"dead_time"`
	// modify_time	int64	最后修改时间
	ModifyTime int64 `json:"modify_time"`
	// download_times	int32	下载次数
	DownloadTimes int32 `json:"download_times"`
	// uploader	int64	上传者ID
	Uploader int64 `json:"uploader"`
	// uploader_name	string	上传者名字
	UploaderName string `json:"uploader_name"`
}

func (f *File) ToGRPC() *FileGRPC {
	return &FileGRPC{
		GroupId:       f.GroupId,
		FileId:        f.FileId,
		FileName:      f.FileName,
		Busid:         f.Busid,
		FileSize:      f.FileSize,
		UploadTime:    f.UploadTime,
		DeadTime:      f.DeadTime,
		ModifyTime:    f.ModifyTime,
		DownloadTimes: f.DownloadTimes,
		Uploader:      f.Uploader,
		UploaderName:  f.UploaderName,
	}
}

// ToStruct
func (f *FileGRPC) ToStruct() *File {
	return &File{
		GroupId:       f.GroupId,
		FileId:        f.FileId,
		FileName:      f.FileName,
		Busid:         f.Busid,
		FileSize:      f.FileSize,
		UploadTime:    f.UploadTime,
		DeadTime:      f.DeadTime,
		ModifyTime:    f.ModifyTime,
		DownloadTimes: f.DownloadTimes,
		Uploader:      f.Uploader,
		UploaderName:  f.UploaderName,
	}
}

type Folder struct {
	// group_id	int32	群号
	GroupId int32 `json:"group_id"`
	// folder_id	string	文件夹ID
	FolderId string `json:"folder_id"`
	// folder_name	string	文件名
	FolderName string `json:"folder_name"`
	// create_time	int64	创建时间
	CreateTime int64 `json:"create_time"`
	// creator	int64	创建者
	Creator int64 `json:"creator"`
	// creator_name	string	创建者名字
	CreatorName string `json:"creator_name"`
	// total_file_count	int32	子文件数量
	TotalFileCount int32 `json:"total_file_count"`
}

func (f *Folder) ToGRPC() *FolderGRPC {
	return &FolderGRPC{
		GroupId:        f.GroupId,
		FolderId:       f.FolderId,
		FolderName:     f.FolderName,
		CreateTime:     f.CreateTime,
		Creator:        f.Creator,
		CreatorName:    f.CreatorName,
		TotalFileCount: f.TotalFileCount,
	}
}

// ToStruct
func (f *FolderGRPC) ToStruct() *Folder {
	return &Folder{
		GroupId:        f.GroupId,
		FolderId:       f.FolderId,
		FolderName:     f.FolderName,
		CreateTime:     f.CreateTime,
		Creator:        f.Creator,
		CreatorName:    f.CreatorName,
		TotalFileCount: f.TotalFileCount,
	}
}

type GroupFiles struct {
	Files   []*File   `json:"files"`
	Folders []*Folder `json:"folders"`
}

func (g *GroupFiles) ToGRPC() *GroupFilesGRPC {
	result := &GroupFilesGRPC{}
	for _, v := range g.Files {
		result.Files = append(result.Files, v.ToGRPC())
	}
	for _, v := range g.Folders {
		result.Folders = append(result.Folders, v.ToGRPC())
	}
	return result
}

// ToStruct
func (g *GroupFilesGRPC) ToStruct() *GroupFiles {
	result := &GroupFiles{}
	for _, v := range g.Files {
		result.Files = append(result.Files, v.ToStruct())
	}
	for _, v := range g.Folders {
		result.Folders = append(result.Folders, v.ToStruct())
	}
	return result
}

type GroupFilesResult struct {
	Data *GroupFiles `json:"data"`
	// retcode	int64	返回码
	Retcode int64 `json:"retcode"`
	// status	string	状态
	Status string `json:"status"`
	// msg	string	错误信息
	Msg string `json:"msg"`
	// wording	string	错误信息
	Wording string `json:"wording"`
}

func (g *GroupFilesResult) ToGRPC() *GroupFilesResultGRPC {
	result := &GroupFilesResultGRPC{
		Retcode: g.Retcode,
		Status:  g.Status,
		Msg:     g.Msg,
		Wording: g.Wording,
	}
	if g.Data != nil {
		result.Data = g.Data.ToGRPC()
	}
	return result
}

func (g *GroupFilesResultGRPC) ToStruct() *GroupFilesResult {
	result := &GroupFilesResult{
		Retcode: g.Retcode,
		Status:  g.Status,
		Msg:     g.Msg,
		Wording: g.Wording,
	}
	if g.Data != nil {
		result.Data = g.Data.ToStruct()
	}
	return result
}

type FileUrl struct {
	// url	string	文件下载链接
	Url string `json:"url"`
}

func (f *FileUrl) ToGRPC() *FileUrlGRPC {
	return &FileUrlGRPC{
		Url: f.Url,
	}
}

func (f *FileUrlGRPC) ToStruct() *FileUrl {
	return &FileUrl{
		Url: f.Url,
	}
}

type FileUrlResult struct {
	Data *FileUrl `json:"data"`
	// retcode	int64	返回码
	Retcode int64 `json:"retcode"`
	// status	string	状态
	Status string `json:"status"`
	// msg	string	错误信息
	Msg string `json:"msg"`
	// wording	string	错误信息
	Wording string `json:"wording"`
}

func (f *FileUrlResult) ToGRPC() *FileUrlResultGRPC {
	result := &FileUrlResultGRPC{
		Retcode: f.Retcode,
		Status:  f.Status,
		Msg:     f.Msg,
		Wording: f.Wording,
	}
	if f.Data != nil {
		result.Data = f.Data.ToGRPC()
	}
	return result
}

func (f *FileUrlResultGRPC) ToStruct() *FileUrlResult {
	result := &FileUrlResult{
		Retcode: f.Retcode,
		Status:  f.Status,
		Msg:     f.Msg,
		Wording: f.Wording,
	}
	if f.Data != nil {
		result.Data = f.Data.ToStruct()
	}
	return result
}
