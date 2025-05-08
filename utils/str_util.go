package utils

import (
	"crypto/md5"
	"fmt"
	"io"
	"log"
	"regexp"
	"sort"
	"strings"

	"github.com/bytedance/sonic"
	"github.com/google/uuid"
)

func GetUUid() string {
	uuidString := strings.ReplaceAll(uuid.New().String(), "-", "")
	return uuidString
}

func JSONMarshal(v interface{}) string {
	byt, err := sonic.Marshal(v)
	if err != nil {
		log.Println(err)
	}
	return string(byt)
}

func JSONUnMarshal(data []byte, v interface{}) interface{} {
	err := sonic.Unmarshal(data, v)
	if err != nil {
		println(err.Error())
		return nil
	}
	return v
}

func ContainsEnglish(text string) bool {
	re := regexp.MustCompile("[A-Za-z]")
	return re.MatchString(text)
}

func StrToMd5(str string) string {
	hasher := md5.New()
	io.WriteString(hasher, str) // 写入字符串到MD5哈希器

	// 获取哈希值
	hashBytes := hasher.Sum(nil)
	hashStr := fmt.Sprintf("%x", hashBytes) // 转换为16进制字符串
	return hashStr
}

// RemoveDuplicateStrings 函数用于移除字符串切片中的重复元素
func RemoveDuplicateStrings(strings []string) []string {
	// 创建一个map来存储不重复的字符串
	uniqueMap := make(map[string]bool)

	// 将切片转换为一个排序后的切片
	sort.Strings(strings)
	// 遍历排序后的字符串，只保留第一次出现的字符串
	var uniqueStrings []string
	for _, str := range strings {
		if _, exists := uniqueMap[str]; !exists {
			uniqueMap[str] = true
			uniqueStrings = append(uniqueStrings, str)
		}
	}
	return uniqueStrings
}

// 字符串长度， len指字节长度
func CharCount(s string) int {
	runes := []rune(s)
	return len(runes)
}

func StrCompile(text string) string {
	re := regexp.MustCompile(`【\d+】`)
	text = re.ReplaceAllString(text, "")
	// 将所有的数字+点+空格，前面加上零宽符号
	//re = regexp.MustCompile(`(\d+\.\s+)`)
	//injectChar := "\u200c"
	//text = re.ReplaceAllString(text, injectChar+"$1")

	//尾部定位标识转换MD格式
	//text = StrReplaceNumber(text)
	// 去除尾部的所有空格、\u200c、换行
	re2 := regexp.MustCompile(`([\x{200c}\s\n\t])+$`)
	text = re2.ReplaceAllString(text, "")

	return text
}

func StrReplaceNumber(text string) string {
	// 使用正则表达式匹配目标模式【数字】
	pattern := regexp.MustCompile(`【(\d+)】`) // \【 和 】 需要被转义
	// 定义替换逻辑的函数
	replacer := func(match []byte) []byte {
		//匹配括号内的内容，即数字部分
		re := regexp.MustCompile(`\d+`)                    // 匹配一个或多个数字字符
		num := re.FindString(string(match))                // 转换为字符串
		return []byte(fmt.Sprintf("[[%s]](%s)", num, num)) // 格式化替换字符串
	}
	// 使用ReplaceAllFunc进行替换
	newStr := pattern.ReplaceAllFunc([]byte(text), replacer)
	return string(newStr)
}

func StrReplace(compile, text string) string {
	re := regexp.MustCompile(compile)
	text = re.ReplaceAllString(text, ",")
	return text
}
