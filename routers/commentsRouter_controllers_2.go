package routers

import (
	"github.com/astaxie/beego"
)

func init() {


	beego.GlobalControllerRouter["github.com/jsreyes/configuracion_api/controllers:MenuOpcionPadreController"] = append(beego.GlobalControllerRouter["configuracion/controllers:MenuOpcionPadreController"],
		beego.ControllerComments{
			Method: "ArbolMenus",
			Router: `/ArbolMenus/:perfil`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/jsreyes/configuracion_api/controllers:PerfilXMenuOpcionController"] = append(beego.GlobalControllerRouter["configuracion/controllers:PerfilXMenuOpcionController"],
		beego.ControllerComments{
			Method: "MenusPorAplicacion",
			Router: `/MenusPorAplicacion/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

}


