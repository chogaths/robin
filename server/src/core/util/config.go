package util

import (
	"encoding/json"
	"github.com/golang/protobuf/proto"
	"io/ioutil"
)

func LoadJson(filename string, out interface{}) error {

	content, err := ioutil.ReadFile(filename)

	if err != nil {
		return err
	}

	err = json.Unmarshal(content, &out)

	if err != nil {
		return err
	}

	return nil

}

func LoadPBTFile(filename string, msg proto.Message) error {

	content, err := ioutil.ReadFile(filename)

	if err != nil {
		return err
	}

	return proto.UnmarshalText(string(content), msg)
}
