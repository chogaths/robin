package stringfilter

import (
	"bufio"
	"io"
	"log"
	"os"
	"strings"
)

type ListFile struct{}

func (self *ListFile) Load(fileName string, doLine func(strline string)) (int32, error) {
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
			if doLine != nil {
				doLine(line)
			}
		}
	} else {
		log.Println("open file error %s", fileName)
		return 0, err
	}

	return lineCount, nil
}
