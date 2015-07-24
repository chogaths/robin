package dbfw

import (
	"bytes"
	"fmt"
	"log"
	"reflect"
	"strconv"
	"strings"
)

const Token_FIELD_NAME = "$FIELD_NAME$"   // 'charname'	query
const Token_FIELD_EQUN = "$FIELD_EQUN$"   // 'coin'=100   update
const Token_FIELD_VALUE = "$FIELD_VALUE$" //  100	    insert

func rawReflectCompose(buff *bytes.Buffer, s interface{}, tagMatch string, indent int, callback func(*bytes.Buffer, reflect.Value, string) bool) bool {

	typeinfo := reflect.TypeOf(s)
	valueinfo := reflect.ValueOf(s)

	if typeinfo.Kind() == reflect.Ptr {
		typeinfo = typeinfo.Elem()
		valueinfo = valueinfo.Elem()
	}

	var validCount int

	for i := 0; i < valueinfo.NumField(); i++ {

		v := valueinfo.Field(i)

		tdef := typeinfo.Field(i)

		switch tdef.Type.Kind() {
		case reflect.Ptr, reflect.Struct:
			rawReflectCompose(buff, v.Interface(), tagMatch, indent+1, callback)
		default:
			tag := tdef.Tag.Get("db")

			// 匹配tag才可合成
			if !strings.Contains(tag, tagMatch) {
				continue
			}

			fieldName := strings.ToLower(tdef.Name)

			if strings.Contains(fieldName, "unrecognized") {
				continue
			}

			if validCount > 0 || indent > 0 {
				buff.WriteString(", ")
			}

			if !callback(buff, v, fieldName) {
				return false
			}

			validCount++
		}

	}

	return true

}

// 根据规则, 反射字段名
func reflectCompose(s interface{}, cmd, token, tagMatch string, callback func(*bytes.Buffer, reflect.Value, string) bool) string {

	if strings.Index(cmd, token) == -1 || s == nil {
		return cmd
	}

	var buff bytes.Buffer

	rawReflectCompose(&buff, s, tagMatch, 0, callback)

	return strings.Replace(cmd, token, buff.String(), 1)
}

// 使用等式反射字段
func useEquationField(buff *bytes.Buffer, fieldValue reflect.Value, fieldName string) bool {
	buff.WriteString(fmt.Sprintf("`%s`=", fieldName))

	return useFieldValue(buff, fieldValue, fieldName)
}

// 使用字段名反射字段
func useFieldName(buff *bytes.Buffer, fieldValue reflect.Value, fieldName string) bool {
	buff.WriteString(fmt.Sprintf("`%s`", fieldName))
	return true
}

// 使用字段值反射字段
func useFieldValue(buff *bytes.Buffer, fieldValue reflect.Value, fieldName string) bool {
	switch fieldValue.Kind() {
	case reflect.Int32, reflect.Int64:
		buff.WriteString(strconv.FormatInt(fieldValue.Int(), 10))
		return true
	case reflect.String:
		//@TODO mysql_escape_string
		buff.WriteString(fmt.Sprintf("'%s'", fieldValue.String()))
		return true
	}

	log.Printf("dbfw.ComposeDBCommand unsupport type: %s", fieldValue.Type())
	return false
}
