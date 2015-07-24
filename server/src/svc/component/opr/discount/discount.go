package icloud

import (
	"core/martinihttp"
	"core/netfw"
	"github.com/martini-contrib/binding"
	"github.com/martini-contrib/render"
	"github.com/martini-contrib/sessionauth"
	"log"
	"svc/component/opr/dbopr"
)

type dbDiscountInfo struct {
	Id        int32  `db:"q" form:"id"`
	Pstyle    int32  `db:"q" form:"pstyle"`
	Pvalue    int32  `db:"q" form:"pvalue"`
	Pcount    int32  `db:"q" form:"pcount"`
	Prlevel   int32  `db:"q" form:"prlevel"`
	Pjilv     int32  `db:"q" form:"pjilv"`
	Pcurrency int32  `db:"q" form:"pcurrency"`
	Pprice    int32  `db:"q|u" form:"pprice"`
	Popen     int32  `db:"q" form:"popen"`
	Ptitle    string `db:"q" form:"ptitle"`
}

type allDiscount struct {
	Discount float32 `form:"discount"`
}

const dbName = "db_discount"

type oprDiscountCom struct {
}

func (self *oprDiscountCom) Start(peer netfw.IPeer) {

	m := peer.(martinihttp.IMartiniAcceptor).GetInterface()

	m.Get("/discount", sessionauth.LoginRequired, func(r render.Render) {

		r.HTML(200, "index", nil)
		r.HTML(200, "discount", nil)

	})

	m.Post("/getdiscount", sessionauth.LoginRequired, func(r render.Render) {

		res, err := dbopr.ExecuteSQL(dbName, &dbDiscountInfo{}, "select $FIELD_NAME$ from tb_onlinestore")
		if err != nil {
			log.Println(err)
			r.Data(200, []byte(err.Error()))
			return
		}

		r.JSON(200, res)

	})

	m.Post("/discount", binding.Bind(dbDiscountInfo{}), sessionauth.LoginRequired, func(r render.Render, msg dbDiscountInfo) {

		_, err := dbopr.ExecuteSQL(dbName, &msg, "update tb_onlinestore set $FIELD_EQUN$ where id=%d", msg.Id)
		if err != nil {
			log.Println(err)
			r.Data(200, []byte(err.Error()))
			return
		}

		r.JSON(200, "execute success")

	})

	m.Post("/alldiscount", binding.Bind(allDiscount{}), sessionauth.LoginRequired, func(r render.Render, msg allDiscount) {

		_, err := dbopr.ExecuteSQL(dbName, nil, "update tb_onlinestore set pprice=pprice*%f", msg.Discount)
		if err != nil {
			log.Println(err)
			r.Data(200, []byte(err.Error()))
			return
		}

		r.JSON(200, "execute success")

	})

}

func init() {

	netfw.RegisterComponent("discount", func(p netfw.IPeer) netfw.IComponent {
		return &oprDiscountCom{}
	})

}
