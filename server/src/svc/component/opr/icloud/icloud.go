package icloud

import (
	"core/martinihttp"
	"core/netfw"
	"crypto/md5"
	"encoding/hex"
	"github.com/martini-contrib/binding"
	"github.com/martini-contrib/render"
	"github.com/martini-contrib/sessionauth"
	"svc/component/opr/dbopr"
)

const dbName = "db_icloud"

type DBIcloudData struct {
	Id                      int32  `db:"q" form:"id"`
	Imei                    string `db:"q"`
	Email                   string `db:"q"`
	Password                string `db:"q|u" form:"password"`
	Remainder_restore_times int32  `db:"q|u" form:"restore"`
	Enable                  int32  `db:"q"`
}

type queryByZone struct {
	Zone  int32  `form:"zone"`
	Page  int32  `form:"page"`
	Id    int32  `form:"id"`
	Imei  string `form:"imei"`
	Email string `form:"email"`
}

type oprIcloudCom struct {
}

func (self *oprIcloudCom) Start(peer netfw.IPeer) {

	m := peer.(martinihttp.IMartiniAcceptor).GetInterface()

	m.Get("/icloud", sessionauth.LoginRequired, func(r render.Render) {

		page := dbopr.GetRecordPageCount(dbName, "tb_td2_backup")

		r.HTML(200, "index", nil)
		r.HTML(200, "icloud", page)

	})

	m.Post("/queryicloud", binding.Bind(queryByZone{}), sessionauth.LoginRequired, func(msg queryByZone, r render.Render) {

		var res []interface{}
		var err error

		if msg.Id != 0 {
			res, err = dbopr.ExecuteSQL(dbName, &DBIcloudData{}, "select $FIELD_NAME$ from tb_td2_backup where id=%d", msg.Id)
		} else if msg.Imei != "" {
			res, err = dbopr.ExecuteSQL(dbName, &DBIcloudData{}, "select $FIELD_NAME$ from tb_td2_backup where imei='%s'", msg.Imei)
		} else if msg.Email != "" {
			res, err = dbopr.ExecuteSQL(dbName, &DBIcloudData{}, "select $FIELD_NAME$ from tb_td2_backup where email='%s'", msg.Email)
		} else {
			res, err = dbopr.ExecuteSQL(dbName, &DBIcloudData{}, "select $FIELD_NAME$ from tb_td2_backup limit %d,%d", msg.Page*dbopr.PageRecordCount, dbopr.PageRecordCount)
		}

		if err != nil {
			r.Data(200, []byte(err.Error()))
			return
		}

		r.JSON(200, res)

	})

	m.Post("/icloud", binding.Bind(DBIcloudData{}), func(msg DBIcloudData, r render.Render) {

		_, err := dbopr.ExecuteSQL(dbName, nil, "delete from tb_td2_backup where id=%d", msg.Id)
		if err != nil {
			r.Data(200, []byte(err.Error()))
			return
		}

		r.JSON(200, "execute success")

	})

	m.Post("/updateicloud", binding.Bind(DBIcloudData{}), func(msg DBIcloudData, r render.Render) {

		var err error

		if msg.Password != "" {

			hash := md5.New()
			_, err = hash.Write([]byte(msg.Password))
			if err != nil {
				r.JSON(200, map[string]string{"Result": err.Error()})
				return
			}

			msg.Password = hex.EncodeToString(hash.Sum(nil))

			_, err = dbopr.ExecuteSQL(dbName, &msg, "update tb_td2_backup set $FIELD_EQUN$ where id=%d", msg.Id)

		} else {

			_, err = dbopr.ExecuteSQL(dbName, nil, "update tb_td2_backup set remainder_restore_times=%d where id=%d", msg.Remainder_restore_times, msg.Id)

		}

		if err != nil {
			r.JSON(200, map[string]string{"Result": err.Error()})
			return
		}

		r.JSON(200, map[string]string{
			"Result":   "execute success",
			"Password": msg.Password,
		})
	})

}

func init() {

	netfw.RegisterComponent("icloud", func(p netfw.IPeer) netfw.IComponent {
		return &oprIcloudCom{}
	})

}
