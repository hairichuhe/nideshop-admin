package servers

import (
	"fmt"
	"io"
	"mime/multipart"
	"nideshop-admin/config"
	"nideshop-admin/tools"
	"os"
	"time"
)

var domain string = config.GinVueAdminconfig.UpLoad.Domain // 访问的域名设置，为了在数据库中储存使用
var filepath string = config.GinVueAdminconfig.UpLoad.Path // 你在七牛云的secretKey  这里是我个人测试号的key 仅供测试使用 恳请大家不要乱传东西

// 接收两个参数 一个文件流 一个 bucket 你的七牛云标准空间的名字
func Upload(file *multipart.FileHeader, bucket string) (err error, path string, key string) {
	f, e := file.Open()
	if e != nil {
		fmt.Println(e)
		return e, "", ""
	}

	suffix := tools.GetSuffix(file.Filename)
	mscond := int(time.Now().UnixNano()/1000000) - 1581341955458
	fileKey := fmt.Sprintf("%d%s", mscond, suffix) // 文件名格式 自己可以改 建议保证唯一性
	dst := filepath + bucket + "/" + fileKey       //本地服务器文件保存位置
	defer f.Close()
	//创建 dst 文件
	err = os.MkdirAll(filepath+bucket, 0755)
	if err != nil {
		return err, "", ""
	}
	out, err := os.Create(dst)
	if err != nil {
		return err, "", ""
	}
	defer out.Close()
	// 拷贝文件
	_, err = io.Copy(out, f)
	if err != nil {
		return err, "", ""
	}
	return err, domain + bucket + "/" + fileKey, fileKey
}

func DeleteFile(bucket string, key string) error {

	src := filepath + bucket + "/" + key
	err := os.Remove(src)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}
