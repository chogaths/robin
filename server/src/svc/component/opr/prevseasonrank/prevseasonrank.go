package prevseasonrank

import (
	"core/martinihttp"
	"core/netfw"
	"github.com/martini-contrib/binding"
	"github.com/martini-contrib/render"
	"github.com/martini-contrib/sessionauth"
	"svc/component/opr/dbopr"
)

type dbLastSeason struct {
	ID        int32  `db:"q"`
	T_sid     int32  `db:"q"`
	T_imei    string `db:"q"`
	Area_id   int32  `db:"q"`
	Last_rank int32  `db:"q"`
	Status    int32  `db:"q"`
}

type zoneInfo struct {
	Name     string
	ShowName string
}

type queryZone struct {
	Zone string `form:"zone"`
}

type oprPrevSeasonRankCom struct {
}

func (self *oprPrevSeasonRankCom) Start(peer netfw.IPeer) {

	m := peer.(martinihttp.IMartiniAcceptor).GetInterface()

	m.Get("/seasonrank", sessionauth.LoginRequired, binding.Bind(queryZone{}), func(r render.Render, msg queryZone) {

		var zones []zoneInfo

		dbopr.IterateShowDB(func(name, showname string) {
			zones = append(zones, zoneInfo{
				Name:     name,
				ShowName: showname,
			})
		})

		r.HTML(200, "index", nil)
		r.HTML(200, "seasonrank", zones)

	})

	m.Post("/rankdata", sessionauth.LoginRequired, binding.Bind(queryZone{}), func(r render.Render, msg queryZone) {

		res, err := dbopr.ExecuteSQL(msg.Zone, &dbLastSeason{}, "select $FIELD_NAME$ from tb_last_season")
		if err != nil {
			r.Data(200, []byte(err.Error()))
			return
		}

		r.JSON(200, res)

	})

}

func init() {

	netfw.RegisterComponent("seasonrank", func(p netfw.IPeer) netfw.IComponent {
		return &oprPrevSeasonRankCom{}
	})

}
