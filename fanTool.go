package fanTool

import (
	"crypto/md5"
	"crypto/rand"
	"encoding/base32"
	"encoding/hex"
	"fanTool/data"
	"math/big"
	"regexp"
	"strconv"
	"strings"
	"time"
)

// IsValidMobileNumber 正则表达式匹配以13、14、15、16、17、18、19开头的11位数字
func IsValidMobileNumber(mobileNumber string) bool {
	pattern := `^1[3-9]\d{9}$`
	match, _ := regexp.MatchString(pattern, mobileNumber)
	return match
}

// MaskMobileNumber 保留电话号码的前三位和后四位
func MaskMobileNumber(mobileNumber string) string {
	s := "****"
	if !IsValidMobileNumber(mobileNumber) {
		return s
	}
	return mobileNumber[:3] + s + mobileNumber[len(mobileNumber)-4:]
}

// IsValidIDNumber 正则表达式匹配 18 位数字，最后一位可能是数字或字母 X
func IsValidIDNumber(IDNumber string) bool {
	IDNumber = strings.ToUpper(IDNumber)
	pattern := `^\d{17}[\d|X]$`
	match, _ := regexp.MatchString(pattern, IDNumber)
	if match == false {
		return false
	}
	mustCompile := "((.{6})(.{8})(.{3}))(.)"
	subMatch := regexp.MustCompile(mustCompile).FindStringSubmatch(IDNumber)
	_, err := time.Parse("20060102", subMatch[3])
	return err == nil
}

// MaskIDNumber 保留身份证号前六位和后四位
func MaskIDNumber(IDNumber string) string {
	s := "**********"
	if !IsValidIDNumber(IDNumber) {
		return s
	}
	return IDNumber[:6] + s + IDNumber[len(IDNumber)-4:]
}

// IDNumberInfo 解析身份证号，返回生日、性别、性别名称
func IDNumberInfo(IDNumber string) *data.IDNumberInfo {
	info := &data.IDNumberInfo{}
	if !IsValidIDNumber(IDNumber) {
		return info
	}
	mustCompile := "((.{6})(.{8})(.{3}))(.)"
	subMatch := regexp.MustCompile(mustCompile).FindStringSubmatch(IDNumber)
	loc, _ := time.LoadLocation("Asia/Shanghai")
	birthday, _ := time.ParseInLocation("20060102", subMatch[3], loc)
	info.Birthday = birthday
	info.BirthdayString = subMatch[3]
	info.Sex = 1
	info.SexName = "男"
	order, _ := strconv.Atoi(subMatch[4])
	if (order % 2) == 0 {
		info.Sex = 2
		info.SexName = "女"
	}
	return info
}

// MD5 md5加密
func MD5(s string) string {
	md5Ctx := md5.New()
	md5Ctx.Write([]byte(s))
	return hex.EncodeToString(md5Ctx.Sum(nil))
}

// RandomString 随机字符串
func RandomString(length int) string {
	str := ""
	codeAlphabet := "ABCDEFGHIJKLMNOPQRSTUVWXYZ" + "abcdefghijklmnopqrstuvwxyz" + "0123456789"
	for i := 0; i < length; i++ {
		random, _ := rand.Int(rand.Reader, big.NewInt(int64(len(codeAlphabet))))
		str += string(codeAlphabet[random.Int64()])
	}
	return str
}

// RandomEncodeString 随机字符串
func RandomEncodeString(length int) string {
	randomBytes := make([]byte, 64)
	_, _ = rand.Read(randomBytes)
	return base32.StdEncoding.EncodeToString(randomBytes)[:length]
}

// GenerateGuid 生成GUID
func GenerateGuid() string {
	s := RandomString(20) + strconv.FormatInt(time.Now().UnixNano(), 10) + RandomEncodeString(20)
	return strings.ToUpper(MD5("guid" + s))
}
