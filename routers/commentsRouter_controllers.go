package routers

import (
	"github.com/astaxie/beego"
)

func init() {

	beego.GlobalControllerRouter["github.com/jsreyes/configuracion_api/controllers:AplicacionController"] = append(beego.GlobalControllerRouter["github.com/jsreyes/configuracion_api/controllers:AplicacionController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/jsreyes/configuracion_api/controllers:AplicacionController"] = append(beego.GlobalControllerRouter["github.com/jsreyes/configuracion_api/controllers:AplicacionController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/jsreyes/configuracion_api/controllers:AplicacionController"] = append(beego.GlobalControllerRouter["github.com/jsreyes/configuracion_api/controllers:AplicacionController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/jsreyes/configuracion_api/controllers:AplicacionController"] = append(beego.GlobalControllerRouter["github.com/jsreyes/configuracion_api/controllers:AplicacionController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/jsreyes/configuracion_api/controllers:AplicacionController"] = append(beego.GlobalControllerRouter["github.com/jsreyes/configuracion_api/controllers:AplicacionController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/jsreyes/configuracion_api/controllers:MenuOpcionController"] = append(beego.GlobalControllerRouter["github.com/jsreyes/configuracion_api/controllers:MenuOpcionController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/jsreyes/configuracion_api/controllers:MenuOpcionController"] = append(beego.GlobalControllerRouter["github.com/jsreyes/configuracion_api/controllers:MenuOpcionController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/jsreyes/configuracion_api/controllers:MenuOpcionController"] = append(beego.GlobalControllerRouter["github.com/jsreyes/configuracion_api/controllers:MenuOpcionController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/jsreyes/configuracion_api/controllers:MenuOpcionController"] = append(beego.GlobalControllerRouter["github.com/jsreyes/configuracion_api/controllers:MenuOpcionController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/jsreyes/configuracion_api/controllers:MenuOpcionController"] = append(beego.GlobalControllerRouter["github.com/jsreyes/configuracion_api/controllers:MenuOpcionController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/jsreyes/configuracion_api/controllers:MenuOpcionPadreController"] = append(beego.GlobalControllerRouter["github.com/jsreyes/configuracion_api/controllers:MenuOpcionPadreController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/jsreyes/configuracion_api/controllers:MenuOpcionPadreController"] = append(beego.GlobalControllerRouter["github.com/jsreyes/configuracion_api/controllers:MenuOpcionPadreController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/jsreyes/configuracion_api/controllers:MenuOpcionPadreController"] = append(beego.GlobalControllerRouter["github.com/jsreyes/configuracion_api/controllers:MenuOpcionPadreController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/jsreyes/configuracion_api/controllers:MenuOpcionPadreController"] = append(beego.GlobalControllerRouter["github.com/jsreyes/configuracion_api/controllers:MenuOpcionPadreController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/jsreyes/configuracion_api/controllers:MenuOpcionPadreController"] = append(beego.GlobalControllerRouter["github.com/jsreyes/configuracion_api/controllers:MenuOpcionPadreController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/jsreyes/configuracion_api/controllers:ParametroController"] = append(beego.GlobalControllerRouter["github.com/jsreyes/configuracion_api/controllers:ParametroController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/jsreyes/configuracion_api/controllers:ParametroController"] = append(beego.GlobalControllerRouter["github.com/jsreyes/configuracion_api/controllers:ParametroController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/jsreyes/configuracion_api/controllers:ParametroController"] = append(beego.GlobalControllerRouter["github.com/jsreyes/configuracion_api/controllers:ParametroController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/jsreyes/configuracion_api/controllers:ParametroController"] = append(beego.GlobalControllerRouter["github.com/jsreyes/configuracion_api/controllers:ParametroController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/jsreyes/configuracion_api/controllers:ParametroController"] = append(beego.GlobalControllerRouter["github.com/jsreyes/configuracion_api/controllers:ParametroController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/jsreyes/configuracion_api/controllers:PerfilController"] = append(beego.GlobalControllerRouter["github.com/jsreyes/configuracion_api/controllers:PerfilController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/jsreyes/configuracion_api/controllers:PerfilController"] = append(beego.GlobalControllerRouter["github.com/jsreyes/configuracion_api/controllers:PerfilController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/jsreyes/configuracion_api/controllers:PerfilController"] = append(beego.GlobalControllerRouter["github.com/jsreyes/configuracion_api/controllers:PerfilController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/jsreyes/configuracion_api/controllers:PerfilController"] = append(beego.GlobalControllerRouter["github.com/jsreyes/configuracion_api/controllers:PerfilController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/jsreyes/configuracion_api/controllers:PerfilController"] = append(beego.GlobalControllerRouter["github.com/jsreyes/configuracion_api/controllers:PerfilController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/jsreyes/configuracion_api/controllers:PerfilXMenuOpcionController"] = append(beego.GlobalControllerRouter["github.com/jsreyes/configuracion_api/controllers:PerfilXMenuOpcionController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/jsreyes/configuracion_api/controllers:PerfilXMenuOpcionController"] = append(beego.GlobalControllerRouter["github.com/jsreyes/configuracion_api/controllers:PerfilXMenuOpcionController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/jsreyes/configuracion_api/controllers:PerfilXMenuOpcionController"] = append(beego.GlobalControllerRouter["github.com/jsreyes/configuracion_api/controllers:PerfilXMenuOpcionController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/jsreyes/configuracion_api/controllers:PerfilXMenuOpcionController"] = append(beego.GlobalControllerRouter["github.com/jsreyes/configuracion_api/controllers:PerfilXMenuOpcionController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/jsreyes/configuracion_api/controllers:PerfilXMenuOpcionController"] = append(beego.GlobalControllerRouter["github.com/jsreyes/configuracion_api/controllers:PerfilXMenuOpcionController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

}
