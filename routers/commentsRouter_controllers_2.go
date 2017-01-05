package routers

import (
	"github.com/astaxie/beego"
)

func init() {


	beego.GlobalControllerRouter["github.com/jsreyes/configuracion_api/controllers:PerfilXMenuOpcionController"] = append(beego.GlobalControllerRouter["configuracion/controllers:PerfilXMenuOpcionController"],
		beego.ControllerComments{
			Method: "arbolMenus",
			Router: `/arbolMenus`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

}
