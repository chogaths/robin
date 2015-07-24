package userdata

import (
	"core/martinihttp"
	"core/netfw"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"github.com/martini-contrib/binding"
	"github.com/martini-contrib/render"
	"log"
	"svc/component/opr/dbopr"
	//"github.com/martini-contrib/sessionauth"
)

type dbUserData struct {
	T_sid        int32  `db:"q" form:"t_sid"`
	T_imei       string `db:"q"`
	T_nickname   string `db:"q"`
	T_victory    int32  `db:"q|u" form:"t_victory"`
	T_yuxi_value string `db:"q|u" form:"t_yuxi_value"`
	T_password   string `form:"t_password"`
	Zone         string `form:"zone"`
	Service      int32  `form:"service"`
}

type dataInfo struct {
	Zone    string
	Service int32
	Page    int32
	MaxPage int32
	Data    []dbUserData
}

type queryByZone struct {
	Zone    string `form:"zone"`
	Service int32  `form:"service"`
	Page    int32  `form:"page"`
	ID      int32  `form:"id"`
	IMEI    string `form:"imei"`
}

type dbZone struct {
	Area_id int32 `db:"q"`
}

type zoneInfo struct {
	Name     string
	ShowName string
	AreaID   []int32
}

type oprUserDataCom struct {
}

func (self *oprUserDataCom) Start(peer netfw.IPeer) {

	m := peer.(martinihttp.IMartiniAcceptor).GetInterface()

	m.Get("/userdata", func(r render.Render) {

		var zonesInfo []zoneInfo

		dbopr.IterateShowDB(func(name, showName string) {

			zone := zoneInfo{
				Name:     name,
				ShowName: showName,
			}

			res, err := dbopr.ExecuteSQL(name, &dbZone{}, "select $FIELD_NAME$ from tb_area")
			if err != nil {
				log.Println(err)
			}

			for _, b := range res {
				area := b.(dbZone)
				zone.AreaID = append(zone.AreaID, area.Area_id)
			}

			if len(zone.AreaID) > 0 {
				zonesInfo = append(zonesInfo, zone)
			}

		})

		r.HTML(200, "index", nil)
		r.HTML(200, "userdata", zonesInfo)
	})

	m.Post("/queryuserdata", binding.Bind(queryByZone{}), func(msg queryByZone, r render.Render) {

		table := fmt.Sprintf("tb_rank_%d", msg.Service)

		if !dbopr.IsTableExist(msg.Zone, table) {
			r.Data(200, []byte(fmt.Sprintf("zone(%s) is not exist!", table)))
			return
		}

		maxpage := dbopr.GetRecordPageCount(msg.Zone, table)

		var err error
		var res []interface{}

		if msg.Zone == "db_td2ios" {
			if msg.ID != 0 {
				res, err = dbopr.ExecuteSQL(msg.Zone, &dbUserData{}, "select %s.t_sid, %s.t_imei, %s.t_nickname, %s.t_victory, %s.t_yuxi_value, tb_user.t_password from %s, tb_user where %s.t_sid=tb_user.t_sid and %s.t_sid=%d", table, table, table, table, table, table, table, table, msg.ID)
			} else if msg.IMEI != "" {
				res, err = dbopr.ExecuteSQL(msg.Zone, &dbUserData{}, "select %s.t_sid, %s.t_imei, %s.t_nickname, %s.t_victory, %s.t_yuxi_value, tb_user.t_password from %s, tb_user where %s.t_sid=tb_user.t_sid and %s.t_imei='%s'", table, table, table, table, table, table, table, table, msg.IMEI)
			} else {
				if msg.Page <= 0 {
					msg.Page = 1
				} else if msg.Page > maxpage {
					msg.Page = maxpage
				}
				res, err = dbopr.ExecuteSQL(msg.Zone, &dbUserData{}, "select %s.t_sid, %s.t_imei, %s.t_nickname, %s.t_victory, %s.t_yuxi_value, tb_user.t_password from %s, tb_user where %s.t_sid=tb_user.t_sid limit %d,%d", table, table, table, table, table, table, table, (msg.Page-1)*100, 100)
			}
		} else {
			if msg.ID != 0 {
				res, err = dbopr.ExecuteSQL(msg.Zone, &dbUserData{}, "select $FIELD_NAME$ from %s where t_sid=%d", table, msg.ID)
			} else if msg.IMEI != "" {
				res, err = dbopr.ExecuteSQL(msg.Zone, &dbUserData{}, "select $FIELD_NAME$ from where t_imei='%s'", table, msg.IMEI)
			} else {
				if msg.Page <= 0 {
					msg.Page = 1
				} else if msg.Page > maxpage {
					msg.Page = maxpage
				}
				res, err = dbopr.ExecuteSQL(msg.Zone, &dbUserData{}, "select $FIELD_NAME$ from %s limit %d,%d", table, (msg.Page-1)*100, 100)
			}
		}

		if err != nil {
			r.Data(200, []byte(err.Error()))
			return
		}

		info := dataInfo{
			Zone:    msg.Zone,
			Service: msg.Service,
			Page:    msg.Page,
			MaxPage: maxpage,
		}

		for _, v := range res {
			user := v.(dbUserData)
			info.Data = append(info.Data, user)
		}

		r.JSON(200, info)

	})

	m.Post("/userdata", binding.Bind(dbUserData{}), func(msg dbUserData, r render.Render) {

		table := fmt.Sprintf("tb_rank_%d", msg.Service)

		if !dbopr.IsTableExist(msg.Zone, table) {
			r.Data(200, []byte(fmt.Sprintf("zone(%s) is not exist!", table)))
			return
		}

		_, err := dbopr.ExecuteSQL(msg.Zone, &msg, "update %s set $FIELD_EQUN$ where t_sid=%d", table, msg.T_sid)
		if err != nil {
			log.Println(err)
			r.Data(200, []byte(err.Error()))
			return
		}

		if msg.T_password != "" {

			hash := md5.New()
			_, err = hash.Write([]byte(msg.T_password))
			if err != nil {
				r.JSON(200, []byte(err.Error()))
				return
			}

			msg.T_password = hex.EncodeToString(hash.Sum(nil))

			log.Println("update user password:", msg.T_password)
			_, err = dbopr.ExecuteSQL(msg.Zone, nil, "update tb_user set t_password='%s' where t_sid=%d", msg.T_password, msg.T_sid)
			if err != nil {
				log.Println(err)
				r.JSON(200, []byte(err.Error()))
				return
			} else {
				r.JSON(200, msg)
				return
			}

		}

		r.JSON(200, "execute success")

	})

	m.Post("/deleteuser", binding.Bind(dbUserData{}), func(msg dbUserData, r render.Render) {

		table := fmt.Sprintf("tb_rank_%d", msg.Service)

		if !dbopr.IsTableExist(msg.Zone, table) {
			r.Data(200, []byte(fmt.Sprintf("zone(%s) is not exist!", table)))
			return
		}

		_, err := dbopr.ExecuteSQL(msg.Zone, nil, "delete from %s where t_sid=%d", table, msg.T_sid)
		if err != nil {
			log.Println(err)
			r.Data(200, []byte(err.Error()))
			return
		}

		_, err = dbopr.ExecuteSQL(msg.Zone, nil, "delete from tb_user where t_sid=%d", msg.T_sid)
		if err != nil {
			log.Println(err)
			r.Data(200, []byte(err.Error()))
			return
		}

		r.JSON(200, "execute success")

	})

}

func init() {

	netfw.RegisterComponent("userdata", func(p netfw.IPeer) netfw.IComponent {
		return &oprUserDataCom{}
	})

}
