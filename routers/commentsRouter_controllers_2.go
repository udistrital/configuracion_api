package routers

import (
	"github.com/astaxie/beego"
)

func init() {


	beego.GlobalControllerRouter["configuracion_V2/controllers:PerfilXMenuOpcionController"] = append(beego.GlobalControllerRouter["configuracion/controllers:PerfilXMenuOpcionController"],
		beego.ControllerComments{
			Method: "arbolMenus",
			Router: `/arbolMenus`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

}
