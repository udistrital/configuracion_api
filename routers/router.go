// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"github.com/udistrital/configuracion_api/controllers"

	"github.com/astaxie/beego"
)

func init() {
	ns := beego.NewNamespace("/v1",

		beego.NSNamespace("/perfil_x_menu_opcion",
			beego.NSInclude(
				&controllers.PerfilXMenuOpcionController{},
			),
		),

		beego.NSNamespace("/menu_opcion_padre",
			beego.NSInclude(
				&controllers.MenuOpcionPadreController{},
			),
		),

		beego.NSNamespace("/menu_opcion",
			beego.NSInclude(
				&controllers.MenuOpcionController{},
			),
		),

		beego.NSNamespace("/aplicacion",
			beego.NSInclude(
				&controllers.AplicacionController{},
			),
		),

		beego.NSNamespace("/aplicacion_rol",
			beego.NSInclude(
				&controllers.AplicacionRolController{},
			),
		),

		beego.NSNamespace("/parametro",
			beego.NSInclude(
				&controllers.ParametroController{},
			),
		),

		beego.NSNamespace("/perfil",
			beego.NSInclude(
				&controllers.PerfilController{},
			),
		),
		beego.NSNamespace("/proceso",
			beego.NSInclude(
				&controllers.ProcesoController{},
			),
		),
		beego.NSNamespace("/proceso_version",
			beego.NSInclude(
				&controllers.VersionProcesoController{},
			),
		),
		beego.NSNamespace("/proceso_estado",
			beego.NSInclude(
				&controllers.EstadoProcesoController{},
			),
		),
		beego.NSNamespace("/proceso_transicion",
			beego.NSInclude(
				&controllers.TransicionProcesoController{},
			),
		),
	)
	beego.AddNamespace(ns)
}
