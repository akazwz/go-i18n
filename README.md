# go i18n

### a simple go i18n

### support translate file JSON, TOML, YAML, HCL, envfile and Java properties config files

### simple use

1. add translate files in a dir,default is ./config/i18n filename should be language code like 'en', 'zh'

en.yaml

````yaml en.yaml
test:
  name: Test
````

----
zh.yaml

````yaml zh.yaml
test:
  name: 测试
````

2. just use

````go
// Test in English is Test
// Test in Chinese is 测试
func useYaml() {
    i := i18n.I18n{}
    i.SetLang("en")
    i.SetTransFilePath("./config/i18n")
    i.SetTransFileType("yaml")

    enStr := i.Trans("test.name").ToStr()

    i.SetLang("zh")
    zhStr := i.Trans("test.name").ToStr()

    fmt.Printf("Test in English is %s\n", enStr)
    fmt.Printf("Test in Chinese is %s\n", zhStr)
}
````

### setting

#### 1. default lang is en, change bu using SetLang
 
#### 2. default translate file path is "./config/i18n", change by using SetTransFilePath

#### 3. default translate file type is yaml, change by using SetTransFileType
