package i18n

import (
	"encoding/json"
	"github.com/spf13/viper"
	"strconv"
)

type I18n struct {
	TransFilePath string
	Lang          string
	TransFileType string
	Value         interface{}
}

// SetTransFilePath set translate file path,default is "./config/i18n"
func (i *I18n) SetTransFilePath(transFilePath string) *I18n {
	i.TransFilePath = transFilePath
	return i
}

// SetLang set language, default is "en"
func (i *I18n) SetLang(lang string) *I18n {
	i.Lang = lang
	return i
}

// SetTransFileType set translate file type, default is "yaml"
func (i *I18n) SetTransFileType(transFileType string) *I18n {
	i.TransFileType = transFileType
	return i
}

// Trans translate word
func (i *I18n) Trans(word string) *I18n {
	config := viper.New()
	transFilePath := "./config/i18n"
	lang := "en"
	transFileType := "yaml"
	if i.TransFilePath != "" {
		transFilePath = i.TransFilePath
	}
	if i.Lang != "" {
		lang = i.Lang
	}
	if i.TransFileType != "" {
		transFileType = i.TransFileType
	}
	config.SetConfigName(lang)
	config.AddConfigPath(transFilePath)
	config.SetConfigType(transFileType)
	if err := config.ReadInConfig(); err != nil {
		panic(err)
	}
	value := config.Get(word)
	i.Value = value
	return i
}

// ToStr after Trans() if you need string type, use this
func (i *I18n) ToStr() string {
	value := i.Value
	var key string
	if value == nil {
		return key
	}
	switch value.(type) {
	case float64:
		ft := value.(float64)
		key = strconv.FormatFloat(ft, 'f', -1, 64)
	case float32:
		ft := value.(float32)
		key = strconv.FormatFloat(float64(ft), 'f', -1, 64)
	case int:
		it := value.(int)
		key = strconv.Itoa(it)
	case uint:
		it := value.(uint)
		key = strconv.Itoa(int(it))
	case int8:
		it := value.(int8)
		key = strconv.Itoa(int(it))
	case uint8:
		it := value.(uint8)
		key = strconv.Itoa(int(it))
	case int16:
		it := value.(int16)
		key = strconv.Itoa(int(it))
	case uint16:
		it := value.(uint16)
		key = strconv.Itoa(int(it))
	case int32:
		it := value.(int32)
		key = strconv.Itoa(int(it))
	case uint32:
		it := value.(uint32)
		key = strconv.Itoa(int(it))
	case int64:
		it := value.(int64)
		key = strconv.FormatInt(it, 10)
	case uint64:
		it := value.(uint64)
		key = strconv.FormatUint(it, 10)
	case string:
		key = value.(string)
	case []byte:
		key = string(value.([]byte))
	default:
		newValue, _ := json.Marshal(value)
		key = string(newValue)
	}
	return key
}
