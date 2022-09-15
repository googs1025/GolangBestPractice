package DoSomethingWithFile

import (
	"log"
	"os"
)

// 主要分为：
// 创建文件、打开文件、关闭文件、改变文件权限

func DoSomethingForFile(filename string) {

	// 创建文件
	file, err := os.Create(filename)
	if err != nil {
		log.Fatalf("创建文件失败 err=%s\n", err)
	}

	// 取得文件信息
	fileInfo, err := file.Stat()
	if err != nil {
		log.Fatalf("获取文件信息失败 err=%s\n", err)
	}

	log.Printf("文件名 %s\n", fileInfo.Name())
	log.Printf("文件权限等级 %s\n", fileInfo.Mode())
	log.Printf("文件修改时间 %s\n", fileInfo.ModTime())

	// 改变文件权限
	err = file.Chmod(0777)
	if err != nil {
		log.Fatalf("修改文件权限失败！ err=%s\n", err)
	}

	// 改变文件拥有者
	err = file.Chown(os.Getuid(), os.Getgid())
	if err != nil{
		log.Fatalf("改变拥有者失败！ err=%s\n", err)
	}

	//fileInfo, err = file.Stat()
	//if err != nil{
	//	log.Fatalf("get file info second failed err=%s\n", err)
	//}
	//log.Printf("File change Permissions is %s\n", fileInfo.Mode())

	// 关闭文件
	err = file.Close()
	if err != nil{
		log.Fatalf("关闭文件失败 err=%s\n", err)
	}

	// 删除文件
	err = os.Remove(filename)
	if err != nil{
		log.Fatalf("删除文件失败 err=%s\n", err)
	}

}
