//author brzhu
package config

import (
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"log"
)

const (
	Yaml = "yaml"
	Json = "json"
)

var RegisterFile = make(map[string]string)

// 用户的配置文件变更需要回调的方法，需要实现这个接口
type ICallBack interface {
	InitConfig(filePath string)
}

//加载配置文件得到viper对象，通过viper.get各种方法获取需要的参数,支持动态加载
//filePath 文件路径
//fileType 文件类型，比如yaml,json等viper支持的文件类型
//callBack 回掉函数，如果文件发生变动，会回调这个函数，使用方加载自己的逻辑代码
//obj 需要反序列化成的对象，注意一定要传指针，否则无法改变原对象的值
func LoadUmshal(filePath string, fileType string, callBack ICallBack, obj interface{}) error {
	vip, err := Load(filePath, fileType, callBack)
	if err != nil {
		return err
	}
	//fmt.Println(vip.GetString("Engines"))
	if err := vip.Unmarshal(obj); err != nil {
		return err
	}
	return nil
}

var i = 1

// 加载配置文件得到viper对象，通过viper.get各种方法获取需要的参数，支持动态加载
//filePath 文件路径
//fileType 文件类型，比如yaml,json等viper支持的文件类型
//callBack 回掉函数，如果文件发生变动，会回调这个函数，使用方加载自己的逻辑代码
func Load(filePath string, fileType string, callBack ICallBack) (*viper.Viper, error) {
	v := viper.New()
	v.SetConfigFile(filePath)
	v.SetConfigType(fileType)
	// viper解析配置文件
	if err := v.ReadInConfig(); err != nil {
		return nil, err
	}
	if callBack != nil {
		// 控制每个文件只用注册一次，多次会回调多次
		if _, ok := RegisterFile[filePath]; ok {
			log.Printf("[RegisterFile中已注册已监听，本次不监听文件变动]%s", filePath)
		} else {
			log.Printf("[RegisterFile中未注册监听，开始注册并监听文件变动]%s", filePath)
			// 如果不在已记录的列表中，才注册监听，否则已注册过
			v.WatchConfig()
			v.OnConfigChange(func(e fsnotify.Event) {
				log.Printf("第%d次，Config file changed: %s\n", i, filePath)
				callBack.InitConfig(filePath)
				i++
			})
			RegisterFile[filePath] = ""
		}
	}
	return v, nil
}
