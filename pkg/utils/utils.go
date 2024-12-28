package utils

import (
	"derma/detect/config"
	"derma/detect/pkg/constants"
	"derma/detect/pkg/errno"
	"fmt"
	"mime/multipart"
	"regexp"
	"strconv"
	"strings"
	"time"
)

func GetMysqlDSN() string {
	host := config.MYSQL.Host
	port := config.MYSQL.Port
	dbName := config.MYSQL.DbName
	username := config.MYSQL.Username
	password := config.MYSQL.Password
	charset := config.MYSQL.Charset

	dsn := strings.Join([]string{username, ":", password, "@tcp(", host, ":", port, ")/", dbName, "?charset=" + charset + "&parseTime=true"}, "")

	return dsn
}

func CheckPhoneNum(phone int64) error {
	s := strconv.FormatInt(phone, 10)
	re, err := regexp.Compile(`^(13[0-9]|14[5|7]|15[0|1|2|3|5|6|7|8|9]|18[0|1|2|3|5|6|7|8|9])\d{8}$`)
	if err != nil {
		return err
	}

	matched := re.MatchString(s)
	if !matched {
		return errno.PhoneFormatError
	}

	return nil
}

func CheckEmail(email string) error {
	re, err := regexp.Compile(`^\w+([-+.]\w+)*@\w+([-.]\w+)*\.\w+([-.]\w+)*$`)
	if err != nil {
		return err
	}

	matched := re.MatchString(email)
	if !matched {
		return errno.EamilFormatError
	}

	return nil
}

func GenerateAvatarName(userID int64) string {
	currentTime := time.Now()
	// 获取年月日和小时分钟
	year, month, day := currentTime.Date()
	hour, minute := currentTime.Hour(), currentTime.Minute()
	return fmt.Sprintf("%v_%d%02d%02d_%02d%02d_avatar.jpg", userID, year, month, day, hour, minute)
}

func GeneratePictureName(userID int64) string {
	currentTime := time.Now()
	// 获取年月日和小时分钟
	year, month, day := currentTime.Date()
	hour, minute := currentTime.Hour(), currentTime.Minute()
	return fmt.Sprintf("%v_%d%02d%02d_%02d%02d_picture.jpg", userID, year, month, day, hour, minute)
}

func GenerateCoverName(userID int64) string {
	currentTime := time.Now()
	// 获取年月日和小时分钟
	year, month, day := currentTime.Date()
	hour, minute := currentTime.Hour(), currentTime.Minute()
	return fmt.Sprintf("%v_%d%02d%02d_%02d%02d_cover.jpg", userID, year, month, day, hour, minute)
}

func GetAvatarURL(objectKey string) string {
	return fmt.Sprintf("https://%s.%s/%s", config.OSS.BucketName, config.OSS.Endpoint, objectKey)
}

func GetPictureURL(objectKey string) string {
	return fmt.Sprintf("https://%s.%s/%s", config.OSS.BucketName, config.OSS.Endpoint, objectKey)
}

func GetDetectionURL(objectKey string) string {
	return fmt.Sprintf("https://%s.%s/%s", config.OSS.BucketName, config.OSS.Endpoint, objectKey)
}

func GetCoverURL(objectKey string) string {
	return fmt.Sprintf("https://%s.%s/%s", config.OSS.BucketName, config.OSS.Endpoint, objectKey)
}

func IsPictureFile(header *multipart.FileHeader) bool {
	contentType := header.Header.Get("Content-Type")
	if strings.HasPrefix(contentType, "image/") {
		return true
	}

	filename := header.Filename
	extensions := []string{".jpg", ".jpeg", ".png"} // Add more picture extensions if needed
	for _, ext := range extensions {
		if strings.HasSuffix(strings.ToLower(filename), ext) {
			return true
		}
	}

	return false
}

func CheckFileSize(header *multipart.FileHeader) bool {
	return header.Size <= constants.MaxFileSize
}
