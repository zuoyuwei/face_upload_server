package face

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"

	"github.com/flipped-aurora/gin-vue-admin/server/model"
	//"gin-vue-admin/model/response"
	//"github.com/flipped-aurora/gin-vue-admin/server/service"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"github.com/gin-gonic/gin"
	"io"
	"io/ioutil"
	"log"
	"mime/multipart"
	"net/http"
	"os"
	"strconv"
	"time"
)

type PersonalUploadFileApi struct{}

func getUserID(c *gin.Context) int {
	var uploadFile model.FtpFile
	if err := c.ShouldBindQuery(&uploadFile); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return -1
	}
	return uploadFile.Shangchuanrenid
}

// @Tags PersonalUploadFile
// @Summary 上传文件
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.FtpFile true "上传文件地址, 文件名"
// @Success 200 {object} response.Response{msg=string} "上传文件"
// @Router /fileUpload/fileUpload [post]
func (e *PersonalUploadFileApi) UploadGo(c *gin.Context) {
	fmt.Println("请求时间：", time.Now().Format("2006-01-02 15:04:05"))
	//UP_CONFIG := global.GVA_CONFIG.Upload.Url
	UP_CONFIG := "58.59.8.83"

	var uploadFile model.FtpFile
	//删除之前的临时文件
	err := os.RemoveAll("./resource/upload/temporary/")
	if err != nil {
		// 删除失败
		fmt.Println("删除失败:", err)
	} else {
		// 删除成功
		fmt.Println("删除成功")
	}

	//获取文件
	file, err := c.FormFile("file")
	if err != nil {
		response.FailWithMessage("文件获取错误", c)
		return
	}
	if file == nil {
		response.FailWithMessage("文件获取错误，文件为空", c)
		return
	}
	//获取文件名
	filename := file.Filename
	fmt.Println("filename:", filename)
	uploadFile.Wenjianming = filename
	//获取今天日期
	timeObj := time.Now()
	var riqi = timeObj.Format("2006-01-02")
	//文件路径
	lujing := "./resource/upload/temporary/" + riqi + "/"
	//判断文件夹是否存在
	var bool bool
	bool, err = utils.PathExists(lujing)
	//文件夹不存在则创建
	if !bool {
		err = os.MkdirAll(lujing, os.ModePerm)
	}

	lujing += filename
	//fmt.Println(lujing)
	uploadFile.Shangchuanrenid = int(getUserID(c))
	//result := "/resource/upload/temporary/" + riqi + "/" + filename
	//uploadFile.Lujing = result

	if err = c.SaveUploadedFile(file, lujing); err != nil {
		response.FailWithMessage("保存文件失败", c)
		return
	} else {
		fmt.Println("转发时间：", time.Now().Format("2006-01-02 15:04:05"))
		url := "http://" + UP_CONFIG + ":9999/up/"

		// key:file 里面放一个文件
		// multipart/form-data 传一个文件
		client := http.Client{}
		bodyBuf := &bytes.Buffer{}
		bodyWrite := multipart.NewWriter(bodyBuf)
		file, err := os.Open(lujing)
		defer file.Close()
		if err != nil {
			log.Println("err")
		}
		// file 为key
		fileWrite, err := bodyWrite.CreateFormFile("file", filename)
		_, err = io.Copy(fileWrite, file)
		if err != nil {
			log.Println("err")
		}

		bodyWrite.Close()
		// 创建请求
		contentType := bodyWrite.FormDataContentType()
		req, err := http.NewRequest("POST", url, bodyBuf)
		if err != nil {
			fmt.Println("req:", err)
		}
		// 设置头
		req.Header.Set("Content-Type", contentType)
		req.Header.Set("Shangchuanrenid", strconv.Itoa(uploadFile.Shangchuanrenid))
		req.Header.Set("pingtai", "ERP")
		resp, err := client.Do(req)

		if err != nil {
			fmt.Println("resp:", err)
		}
		defer resp.Body.Close()

		//把body转换成map
		var result map[string]interface{}
		b, err := ioutil.ReadAll(resp.Body)
		if err == nil {
			json.Unmarshal([]byte(string(b)), &result)
		}

		fmt.Println(result["folderName"])
		//存文件id
		uploadFile.Wenjianid = int(result["id"].(float64))
		folderName := result["folderName"]
		uploadFile.Lujing = folderName.(string)
		fmt.Println("folderName:", folderName.(string))

		if err = global.GVA_DB.Create(&uploadFile).Error; err != nil {
			response.FailWithMessage("保存文件失败", c)
			return
		} else {
			response.OkWithDetailed(response.UploadFileResponse{
				Status:     "上传文件成功",
				FolderName: folderName.(string),
			}, "文件保存成功", c)
			fmt.Println("完成时间：", time.Now().Format("2006-01-02 15:04:05"))
		}
	}
}

// @Tags PersonalUploadFile
// @Summary 下载文件
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.FtpFile true "下载文件地址, 文件名"
// @Success 200 {object} response.Response{msg=string} "下载文件"
// @Router /fileDown/fileDown [post]
func (e *PersonalUploadFileApi) DownGo(c *gin.Context) {

	fmt.Println("请求uploadDown时间：", time.Now().Format("2006-01-02 15:04:05"))
	//UP_CONFIG := global.GVA_CONFIG.Upload.Url
	UP_CONFIG := "58.59.8.83"

	path := c.Query("url")
	requrl := "http://" + UP_CONFIG + ":9999" + path

	// 创建db
	//db := global.GVA_DB.Model(&model.UploadFile{})
	db := global.GVA_DB.Model(&model.FtpFile{})

	var wenjianming string
	err := db.Debug().Raw("select wenjianming from personal_ftp_file where 1=1 and deleted_at is null and lujing=?", path).First(&wenjianming).Error
	if err != nil {
		fmt.Println("请求数据库错误：", err)
	}

	fmt.Println("wenjianming:", wenjianming)

	responseFile, err := http.Get(requrl)
	fmt.Println("请求requrl时间：", time.Now().Format("2006-01-02 15:04:05"))
	if err != nil || responseFile.StatusCode != http.StatusOK {
		fmt.Println("请求文件服务器错误：", err)
		response.FailWithMessage("请求文件服务器失败", c)
		//c.Status(http.StatusServiceUnavailable)
		return
	}

	reader := responseFile.Body
	contentLength := responseFile.ContentLength
	contentType := responseFile.Header.Get("Content-Type")

	extraHeaders := map[string]string{
		"Content-Disposition": fmt.Sprintf(`inline; filename="%s"`, wenjianming),
		//"Content-Type":contentType,
		//"Content-Disposition": "attachment; filename=\""+wenjianming+"\"",
	}

	c.DataFromReader(http.StatusOK, contentLength, contentType, reader, extraHeaders)
	defer responseFile.Body.Close()
	fmt.Println("请求完成时间：", time.Now().Format("2006-01-02 15:04:05"))

}
