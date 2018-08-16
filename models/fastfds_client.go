package models

import (
	"github.com/weilaihui/fdfs_client"
	"github.com/astaxie/beego"
)

func UploadByFilename(filename string)(GroupName,RemoteFileId,err string){
	// 通过配置文件 创建
	fdfsCLient,thiserr := fdfs_client.NewFdfsClient("")
	if thiserr !=nil {
		beego.Info("fdfs_client.NewFdfsClient thiserr",thiserr)
		GroupName =""
		RemoteFileId=""
	}

	uploadRespense,err := fdfsCLient.UploadByFilename(filename,)
	if err !=nil {
		beego.Info("ffdfsCLient.UploadByFilename(filename) err", err)
		GroupName =""
		RemoteFileId=""
		return
	}
	beego.Info(uploadRespense.GroupName)
	beego.Info(uploadRespense.RemoteFileId)

	return uploadRespense.GroupName,uploadRespense.RemoteFileId,nil
}
// 功能函数 二进制fdfs 上传
func uploadbybyffer(){

	fdfsCLient,thiserr := fdfs_client.NewFdfsClient("")
	if thiserr !=nil {
		beego.Info("fdfs_client.NewFdfsClient thiserr",thiserr)
		GroupName =""
		RemoteFileId=""
	}

	uploadRespense,err := fdfsCLient.UploadByBuffer("","")
	if err !=nil {
		beego.Info("ffdfsCLient.UploadByFilename(filename) err", err)
		GroupName =""
		RemoteFileId=""
		return
	}
	beego.Info(uploadRespense.GroupName)
	beego.Info(uploadRespense.RemoteFileId)

	return uploadRespense.GroupName,uploadRespense.RemoteFileId

}
