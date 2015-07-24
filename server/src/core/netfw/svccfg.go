package netfw

import (
	"core/util"
	"encoding/json"
	"github.com/robertkrimen/otto"
	"log"
	"path/filepath"
	"reflect"
	"sync"
)

var (
	requireMap map[string]bool = make(map[string]bool)
	jsVM                       = otto.New()
	jsVMGuard  sync.Mutex
)

func buildin_require(filename string) {

	// 已经包含过了
	if _, ok := requireMap[filename]; ok {
		return
	}

	final := filepath.Join(DefaultConfigPath, filename+".js")

	log.Printf("%s parsing...", final)
	if err := execFile(final); err != nil {
		log.Println(final, err)
		return
	}

	requireMap[filename] = true
}

func buildin_panic(desc string) {
	panic(desc)
}

func execFile(filename string) error {
	script, err := jsVM.Compile(filename, nil)

	if err != nil {
		return err
	}

	if _, err := jsVM.Run(script); err != nil {
		return err
	}

	return nil
}

func execScript(script string) error {

	_, err := jsVM.Run(script)

	return err
}

func initConfigEnv(svcname string, param svcParam) {
	jsVM.Set("require", buildin_require)

	jsVM.Set("PortOffset", util.PortOffset)
	jsVM.Set("ReplaceIP", util.ReplaceIP)
	jsVM.Set("SvcName", SvcName)
	jsVM.Set("SvcIndex", SvcIndex)
	jsVM.Set("panic", buildin_panic)

	buildin_require(svcname)

	if param.configFile != "" {
		log.Printf("exec file: %s", param.configFile)
		execFile(param.configFile)
	}

	if param.configStr != "" {
		log.Printf("exec script: %s", param.configStr)
		execScript(param.configStr)
	}

	log.Println("start FetchConfig")
	fetchConfig("FetchConfig", "", &SvcConfig)
}

func fetchConfig(funcname string, name string, config interface{}) bool {

	var ret otto.Value
	var err error

	if name == "" {
		ret, err = jsVM.Call(funcname, nil)
	} else {
		ret, err = jsVM.Call(funcname, nil, name)
	}

	if err != nil {
		log.Printf("call failed, in '%s', %v", funcname, err)
		return false
	}

	jsonConfig, err := ret.ToString()

	if err != nil {
		log.Println(err)
		return false
	}

	if err := json.Unmarshal([]byte(jsonConfig), config); err != nil {
		log.Printf("json.Unmarshal failed, in '%s', %v", funcname, err)
		return false
	}

	return true
}

// 根据configName获取js全局空间里的configObject, 并序列化到pb结构体
func GetConfig(configName string, msg interface{}) bool {

	jsVMGuard.Lock()

	defer jsVMGuard.Unlock()

	return fetchConfig("GetConfig", configName, msg)
}

func fillStructValue(fieldValue reflect.Value, sourceValue reflect.Value, fieldName string) {
	switch fieldValue.Kind() {
	case reflect.Invalid:
		log.Printf("field not match '%s'", fieldName)
		return
	case reflect.Ptr:

		newPtrValue := reflect.New(fieldValue.Type().Elem())

		if newPtrValue.Elem().Kind() != reflect.Struct {
			if newPtrValue.Type().Elem() != sourceValue.Type() {
				sourceValue = sourceValue.Convert(newPtrValue.Type().Elem())
			}

			// 设置new出来的值
			newPtrValue.Elem().Set(sourceValue)

		} else {

			obj := sourceValue.Interface().(map[string]interface{})

			fillStruct(obj, newPtrValue)
		}

		// 设置new出来的ptr给field
		fieldValue.Set(newPtrValue)

	case reflect.Slice:

		if sourceValue.Kind() != reflect.Slice {
			log.Printf("should add [] to repeated field '%s'", fieldName)
			return
		}

		sliceValue := reflect.MakeSlice(fieldValue.Type(), sourceValue.Len(), sourceValue.Len())

		for i := 0; i < sourceValue.Len(); i++ {

			arrValue := sourceValue.Index(i)

			arrField := sliceValue.Index(i)

			if arrValue.Kind() == reflect.Interface {
				arrValue = reflect.ValueOf(arrValue.Interface())
			}

			arrField.Set(arrValue)
		}

		fieldValue.Set(sliceValue)
	}
}

func fillStruct(obj map[string]interface{}, structValue reflect.Value) {

	if structValue.IsNil() {
		return
	}

	if structValue.Kind() == reflect.Ptr {
		structValue = structValue.Elem()
	}

	for k, v := range obj {

		fieldValue := structValue.FieldByName(k)

		sourceValue := reflect.ValueOf(v)

		log.Printf("set %s=%v  %v|%v", k, v, fieldValue.Kind().String(), sourceValue.Kind().String())

		fillStructValue(fieldValue, sourceValue, k)
	}

}
