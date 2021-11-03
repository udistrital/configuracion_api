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

		beego.NSNamespace("/notificacion_estado_usuario",
			beego.NSInclude(
				&controllers.NotificacionEstadoUsuarioController{},
			),
		),

		beego.NSNamespace("/notificacion_configuracion",
			beego.NSInclude(
				&controllers.NotificacionConfiguracionController{},
			),
		),

		beego.NSNamespace("/notificacion_configuracion_perfil",
			beego.NSInclude(
				&controllers.NotificacionConfiguracionPerfilController{},
			),
		),

		beego.NSNamespace("/metodo_http",
			beego.NSInclude(
				&controllers.MetodoHttpController{},
			),
		),
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
		beego.NSNamespace("/notificacion_tipo",

			beego.NSInclude(
				&controllers.NotificacionTipoController{},
			),
		),

		beego.NSNamespace("/notificacion",
			beego.NSInclude(
				&controllers.NotificacionController{},
			),
		),

		beego.NSNamespace("/notificacion_estado",
			beego.NSInclude(
				&controllers.NotificacionEstadoController{},
			),
		),
		beego.NSNamespace("/perfil",
			beego.NSInclude(
				&controllers.PerfilController{},
			),
		),
	)
	beego.AddNamespace(ns)
}
