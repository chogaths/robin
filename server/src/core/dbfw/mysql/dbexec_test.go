package dbfw

import (
	"log"
	"testing"
)

type testData struct {
	UserName string
	Password string
	ID       int64
}

// 测试表结构
/*

SET FOREIGN_KEY_CHECKS=0;

-- ----------------------------
-- Table structure for `tb_test`
-- ----------------------------
DROP TABLE IF EXISTS `tb_test`;
CREATE TABLE `tb_test` (
  `username` varchar(64) COLLATE utf8_bin NOT NULL,
  `password` varchar(64) COLLATE utf8_bin NOT NULL,
  `id` bigint(11) NOT NULL AUTO_INCREMENT,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8 COLLATE=utf8_bin;

-- ----------------------------
-- Records of tb_test
-- ----------------------------
INSERT INTO `tb_test` VALUES ('a1', 'p1', '1');
INSERT INTO `tb_test` VALUES ('a2', 'p2', '2');


*/

// 测试数据
/*

a1	p1	1
a2	p2	2


*/

func TestDBExecutor(t *testing.T) {

	dbexec := NewDBExecutor()

	dbexec.Open("root:123456@tcp(localhost:3306)/coretest")

	var id int64 = 1
	result, err := dbexec.Query(&testData{}, "select $FIELD$ from tb_test where id = ?", id)

	if err != nil {
		t.Error(err)
	}

	if len(result) == 1 {
		dataptr := result[0].(testData)
		log.Printf("%s %s", dataptr.UserName, dataptr.Password)
	}

}
