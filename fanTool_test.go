package fanTool

import (
	"log"
	"testing"
)

func Test(t *testing.T) {
	log.Println(IsValidMobileNumber("13800138000"))
	log.Println(IsValidMobileNumber("12800138000"))

	log.Println(MaskMobileNumber("13800138000"))
	log.Println(MaskMobileNumber("12800138000"))

	log.Println(IsValidIDNumber("44258419900101481X"))
	log.Println(IsValidIDNumber("44152446885258451x"))

	log.Println(MaskIDNumber("442584199001014815"))
	log.Println(MaskIDNumber("44152446885258451X"))

	log.Println(IDNumberInfo("44258419901201481x"))

	log.Println(RandomString(6))
	log.Println(RandomEncodeString(6))
	log.Println(MD5("123"))
	log.Println(GenerateGuid())

	log.Println(ThisMorningUnixTime())
	log.Println(YesterdayMorningUnixTime())
	log.Println(StringToUnixTime("2024-01-02 15:04:05"))
	log.Println(StringToUnixTime("2024-01-02"))
	log.Println(StringToUnixTime("2024/01/02"))
	log.Println(StringToUnixTime("20240102"))
	log.Println(StringToTime("2024-01-02 15:04:05"))

}
