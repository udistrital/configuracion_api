package routers

import (
	"github.com/astaxie/beego"
)

func init() {

	beego.GlobalControllerRouter["configuracion_V2/controllers:AplicacionController"] = append(beego.GlobalControllerRouter["configuracion_V2/controllers:AplicacionController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["configuracion_V2/controllers:AplicacionController"] = append(beego.GlobalControllerRouter["configuracion_V2/controllers:AplicacionController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["configuracion_V2/controllers:AplicacionController"] = append(beego.GlobalControllerRouter["configuracion_V2/controllers:AplicacionController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["configuracion_V2/controllers:AplicacionController"] = append(beego.GlobalControllerRouter["configuracion_V2/controllers:AplicacionController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["configuracion_V2/controllers:AplicacionController"] = append(beego.GlobalControllerRouter["configuracion_V2/controllers:AplicacionController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["configuracion_V2/controllers:MenuOpcionController"] = append(beego.GlobalControllerRouter["configuracion_V2/controllers:MenuOpcionController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["configuracion_V2/controllers:MenuOpcionController"] = append(beego.GlobalControllerRouter["configuracion_V2/controllers:MenuOpcionController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["configuracion_V2/controllers:MenuOpcionController"] = append(beego.GlobalControllerRouter["configuracion_V2/controllers:MenuOpcionController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["configuracion_V2/controllers:MenuOpcionController"] = append(beego.GlobalControllerRouter["configuracion_V2/controllers:MenuOpcionController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["configuracion_V2/controllers:MenuOpcionController"] = append(beego.GlobalControllerRouter["configuracion_V2/controllers:MenuOpcionController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["configuracion_V2/controllers:MenuOpcionPadreController"] = append(beego.GlobalControllerRouter["configuracion_V2/controllers:MenuOpcionPadreController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["configuracion_V2/controllers:MenuOpcionPadreController"] = append(beego.GlobalControllerRouter["configuracion_V2/controllers:MenuOpcionPadreController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["configuracion_V2/controllers:MenuOpcionPadreController"] = append(beego.GlobalControllerRouter["configuracion_V2/controllers:MenuOpcionPadreController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["configuracion_V2/controllers:MenuOpcionPadreController"] = append(beego.GlobalControllerRouter["configuracion_V2/controllers:MenuOpcionPadreController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["configuracion_V2/controllers:MenuOpcionPadreController"] = append(beego.GlobalControllerRouter["configuracion_V2/controllers:MenuOpcionPadreController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["configuracion_V2/controllers:ParametroController"] = append(beego.GlobalControllerRouter["configuracion_V2/controllers:ParametroController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["configuracion_V2/controllers:ParametroController"] = append(beego.GlobalControllerRouter["configuracion_V2/controllers:ParametroController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["configuracion_V2/controllers:ParametroController"] = append(beego.GlobalControllerRouter["configuracion_V2/controllers:ParametroController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["configuracion_V2/controllers:ParametroController"] = append(beego.GlobalControllerRouter["configuracion_V2/controllers:ParametroController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["configuracion_V2/controllers:ParametroController"] = append(beego.GlobalControllerRouter["configuracion_V2/controllers:ParametroController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["configuracion_V2/controllers:PerfilController"] = append(beego.GlobalControllerRouter["configuracion_V2/controllers:PerfilController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["configuracion_V2/controllers:PerfilController"] = append(beego.GlobalControllerRouter["configuracion_V2/controllers:PerfilController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["configuracion_V2/controllers:PerfilController"] = append(beego.GlobalControllerRouter["configuracion_V2/controllers:PerfilController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["configuracion_V2/controllers:PerfilController"] = append(beego.GlobalControllerRouter["configuracion_V2/controllers:PerfilController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["configuracion_V2/controllers:PerfilController"] = append(beego.GlobalControllerRouter["configuracion_V2/controllers:PerfilController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["configuracion_V2/controllers:PerfilXMenuOpcionController"] = append(beego.GlobalControllerRouter["configuracion_V2/controllers:PerfilXMenuOpcionController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["configuracion_V2/controllers:PerfilXMenuOpcionController"] = append(beego.GlobalControllerRouter["configuracion_V2/controllers:PerfilXMenuOpcionController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["configuracion_V2/controllers:PerfilXMenuOpcionController"] = append(beego.GlobalControllerRouter["configuracion_V2/controllers:PerfilXMenuOpcionController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["configuracion_V2/controllers:PerfilXMenuOpcionController"] = append(beego.GlobalControllerRouter["configuracion_V2/controllers:PerfilXMenuOpcionController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["configuracion_V2/controllers:PerfilXMenuOpcionController"] = append(beego.GlobalControllerRouter["configuracion_V2/controllers:PerfilXMenuOpcionController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

}
