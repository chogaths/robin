package stringfilter

import (
	"bufio"
	"core/util"
	"github.com/golang/protobuf/proto"
	"io"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"protos/gamedef"
	"strings"
	"testing"
)

func replaceWordTest(fileName string, stringFilter *StringFilter) int32 {
	f, err := os.Open(fileName)
	var lineCount int32
	defer f.Close()

	if err == nil {
		buff := bufio.NewReader(f)

		// 按行读取文件
		for {
			l, err := buff.ReadString('\n')

			// 去掉回车
			line := strings.TrimSpace(l)

			if err != nil || io.EOF == err {
				break
			}
			lineCount++

			// 调用回调处理每行逻辑
			if stringFilter != nil {
				log.Println(stringFilter.Replace(line))
			}
		}
	} else {
		log.Println("open file error")
	}

	return lineCount
}

func TestListFileAndFilter(b *testing.T) {
	log.Println("TestListFileAndFilter")
	t := NewStringFilter()
	t.ListFile.Load("word.txt", t.Insert)
	s := t.Replace("xxxmb")
	log.Println(s)
	log.Println(t.Exist("mb"))
	log.Println(t.Exist("m"))
}

func TestStringFilter(b *testing.T) {
	log.Println("TestListFileAndFilter")
	t := NewStringFilter()
	t.LoadBlogFile("../cfg/chatfilter.txt")
	replaceWordTest("word.txt", t)
}

func loadConfig(name string, m proto.Message) error {

	dataPath := getDataPath()

	var ext = ".pbt"

	final := filepath.Join(dataPath, name) + ext

	if err := util.LoadPBTFile(final, m); err != nil {

		log.Printf("%s %s", final, err)
		return err
	}

	log.Printf("%s loaded", final)

	return nil
}

const (
	ConfigDev = "../../../../../table/obj"
)

func getDataPath() string {
	content, err := ioutil.ReadFile("../cfg/datapath.cfg")

	if err == nil {
		return strings.Trim(string(content), "\r\n ")
	}

	return ConfigDev
}

func TestNameString(b *testing.T) {
	log.Println("TestNameTable string")

	// 加载Name表
	var nameFile gamedef.NameFile
	if loadConfig("Name", &nameFile) != nil {
		return
	}

	t := NewStringFilter()
	t.LoadBlogFile("../../cfg/charnamefilter.txt")
	for _, v := range nameFile.Name {
		//log.Println(v.GetName())
		if t.Exist(v.GetName()) {
			log.Println(v.GetName())
			log.Println(t.Replace(v.GetName()))
		}
	}
}
