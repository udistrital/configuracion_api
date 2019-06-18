package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

    beego.GlobalControllerRouter["github.com/udistrital/configuracion_api/controllers:AplicacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/configuracion_api/controllers:AplicacionController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/configuracion_api/controllers:AplicacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/configuracion_api/controllers:AplicacionController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/configuracion_api/controllers:AplicacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/configuracion_api/controllers:AplicacionController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/configuracion_api/controllers:AplicacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/configuracion_api/controllers:AplicacionController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/configuracion_api/controllers:AplicacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/configuracion_api/controllers:AplicacionController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/configuracion_api/controllers:MenuOpcionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/configuracion_api/controllers:MenuOpcionController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/configuracion_api/controllers:MenuOpcionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/configuracion_api/controllers:MenuOpcionController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/configuracion_api/controllers:MenuOpcionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/configuracion_api/controllers:MenuOpcionController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/configuracion_api/controllers:MenuOpcionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/configuracion_api/controllers:MenuOpcionController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/configuracion_api/controllers:MenuOpcionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/configuracion_api/controllers:MenuOpcionController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/configuracion_api/controllers:MenuOpcionPadreController"] = append(beego.GlobalControllerRouter["github.com/udistrital/configuracion_api/controllers:MenuOpcionPadreController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/configuracion_api/controllers:MenuOpcionPadreController"] = append(beego.GlobalControllerRouter["github.com/udistrital/configuracion_api/controllers:MenuOpcionPadreController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/configuracion_api/controllers:MenuOpcionPadreController"] = append(beego.GlobalControllerRouter["github.com/udistrital/configuracion_api/controllers:MenuOpcionPadreController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/configuracion_api/controllers:MenuOpcionPadreController"] = append(beego.GlobalControllerRouter["github.com/udistrital/configuracion_api/controllers:MenuOpcionPadreController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/configuracion_api/controllers:MenuOpcionPadreController"] = append(beego.GlobalControllerRouter["github.com/udistrital/configuracion_api/controllers:MenuOpcionPadreController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/configuracion_api/controllers:MenuOpcionPadreController"] = append(beego.GlobalControllerRouter["github.com/udistrital/configuracion_api/controllers:MenuOpcionPadreController"],
        beego.ControllerComments{
            Method: "ArbolMenus",
            Router: `/ArbolMenus/:roles/:app`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/configuracion_api/controllers:MetodoHttpController"] = append(beego.GlobalControllerRouter["github.com/udistrital/configuracion_api/controllers:MetodoHttpController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/configuracion_api/controllers:MetodoHttpController"] = append(beego.GlobalControllerRouter["github.com/udistrital/configuracion_api/controllers:MetodoHttpController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/configuracion_api/controllers:MetodoHttpController"] = append(beego.GlobalControllerRouter["github.com/udistrital/configuracion_api/controllers:MetodoHttpController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/configuracion_api/controllers:MetodoHttpController"] = append(beego.GlobalControllerRouter["github.com/udistrital/configuracion_api/controllers:MetodoHttpController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/configuracion_api/controllers:MetodoHttpController"] = append(beego.GlobalControllerRouter["github.com/udistrital/configuracion_api/controllers:MetodoHttpController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/configuracion_api/controllers:NotificacionConfiguracionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/configuracion_api/controllers:NotificacionConfiguracionController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/configuracion_api/controllers:NotificacionConfiguracionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/configuracion_api/controllers:NotificacionConfiguracionController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/configuracion_api/controllers:NotificacionConfiguracionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/configuracion_api/controllers:NotificacionConfiguracionController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/configuracion_api/controllers:NotificacionConfiguracionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/configuracion_api/controllers:NotificacionConfiguracionController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/configuracion_api/controllers:NotificacionConfiguracionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/configuracion_api/controllers:NotificacionConfiguracionController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/configuracion_api/controllers:NotificacionConfiguracionPerfilController"] = append(beego.GlobalControllerRouter["github.com/udistrital/configuracion_api/controllers:NotificacionConfiguracionPerfilController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/configuracion_api/controllers:NotificacionConfiguracionPerfilController"] = append(beego.GlobalControllerRouter["github.com/udistrital/configuracion_api/controllers:NotificacionConfiguracionPerfilController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/configuracion_api/controllers:NotificacionConfiguracionPerfilController"] = append(beego.GlobalControllerRouter["github.com/udistrital/configuracion_api/controllers:NotificacionConfiguracionPerfilController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/configuracion_api/controllers:NotificacionConfiguracionPerfilController"] = append(beego.GlobalControllerRouter["github.com/udistrital/configuracion_api/controllers:NotificacionConfiguracionPerfilController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/configuracion_api/controllers:NotificacionConfiguracionPerfilController"] = append(beego.GlobalControllerRouter["github.com/udistrital/configuracion_api/controllers:NotificacionConfiguracionPerfilController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/configuracion_api/controllers:NotificacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/configuracion_api/controllers:NotificacionController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/configuracion_api/controllers:NotificacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/configuracion_api/controllers:NotificacionController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/configuracion_api/controllers:NotificacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/configuracion_api/controllers:NotificacionController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/configuracion_api/controllers:NotificacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/configuracion_api/controllers:NotificacionController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/configuracion_api/controllers:NotificacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/configuracion_api/controllers:NotificacionController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/configuracion_api/controllers:NotificacionEstadoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/configuracion_api/controllers:NotificacionEstadoController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/configuracion_api/controllers:NotificacionEstadoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/configuracion_api/controllers:NotificacionEstadoController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/configuracion_api/controllers:NotificacionEstadoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/configuracion_api/controllers:NotificacionEstadoController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/configuracion_api/controllers:NotificacionEstadoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/configuracion_api/controllers:NotificacionEstadoController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/configuracion_api/controllers:NotificacionEstadoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/configuracion_api/controllers:NotificacionEstadoController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/configuracion_api/controllers:NotificacionEstadoUsuarioController"] = append(beego.GlobalControllerRouter["github.com/udistrital/configuracion_api/controllers:NotificacionEstadoUsuarioController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/configuracion_api/controllers:NotificacionEstadoUsuarioController"] = append(beego.GlobalControllerRouter["github.com/udistrital/configuracion_api/controllers:NotificacionEstadoUsuarioController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/configuracion_api/controllers:NotificacionEstadoUsuarioController"] = append(beego.GlobalControllerRouter["github.com/udistrital/configuracion_api/controllers:NotificacionEstadoUsuarioController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/configuracion_api/controllers:NotificacionEstadoUsuarioController"] = append(beego.GlobalControllerRouter["github.com/udistrital/configuracion_api/controllers:NotificacionEstadoUsuarioController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/configuracion_api/controllers:NotificacionEstadoUsuarioController"] = append(beego.GlobalControllerRouter["github.com/udistrital/configuracion_api/controllers:NotificacionEstadoUsuarioController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/configuracion_api/controllers:NotificacionEstadoUsuarioController"] = append(beego.GlobalControllerRouter["github.com/udistrital/configuracion_api/controllers:NotificacionEstadoUsuarioController"],
        beego.ControllerComments{
            Method: "GetOldNotification",
            Router: `/getOldNotification/:profile`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/configuracion_api/controllers:NotificacionTipoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/configuracion_api/controllers:NotificacionTipoController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/configuracion_api/controllers:NotificacionTipoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/configuracion_api/controllers:NotificacionTipoController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/configuracion_api/controllers:NotificacionTipoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/configuracion_api/controllers:NotificacionTipoController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/configuracion_api/controllers:NotificacionTipoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/configuracion_api/controllers:NotificacionTipoController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/configuracion_api/controllers:NotificacionTipoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/configuracion_api/controllers:NotificacionTipoController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/configuracion_api/controllers:ParametroController"] = append(beego.GlobalControllerRouter["github.com/udistrital/configuracion_api/controllers:ParametroController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/configuracion_api/controllers:ParametroController"] = append(beego.GlobalControllerRouter["github.com/udistrital/configuracion_api/controllers:ParametroController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/configuracion_api/controllers:ParametroController"] = append(beego.GlobalControllerRouter["github.com/udistrital/configuracion_api/controllers:ParametroController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/configuracion_api/controllers:ParametroController"] = append(beego.GlobalControllerRouter["github.com/udistrital/configuracion_api/controllers:ParametroController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/configuracion_api/controllers:ParametroController"] = append(beego.GlobalControllerRouter["github.com/udistrital/configuracion_api/controllers:ParametroController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/configuracion_api/controllers:PerfilController"] = append(beego.GlobalControllerRouter["github.com/udistrital/configuracion_api/controllers:PerfilController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/configuracion_api/controllers:PerfilController"] = append(beego.GlobalControllerRouter["github.com/udistrital/configuracion_api/controllers:PerfilController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/configuracion_api/controllers:PerfilController"] = append(beego.GlobalControllerRouter["github.com/udistrital/configuracion_api/controllers:PerfilController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/configuracion_api/controllers:PerfilController"] = append(beego.GlobalControllerRouter["github.com/udistrital/configuracion_api/controllers:PerfilController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/configuracion_api/controllers:PerfilController"] = append(beego.GlobalControllerRouter["github.com/udistrital/configuracion_api/controllers:PerfilController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/configuracion_api/controllers:PerfilXMenuOpcionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/configuracion_api/controllers:PerfilXMenuOpcionController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/configuracion_api/controllers:PerfilXMenuOpcionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/configuracion_api/controllers:PerfilXMenuOpcionController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/configuracion_api/controllers:PerfilXMenuOpcionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/configuracion_api/controllers:PerfilXMenuOpcionController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/configuracion_api/controllers:PerfilXMenuOpcionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/configuracion_api/controllers:PerfilXMenuOpcionController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/configuracion_api/controllers:PerfilXMenuOpcionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/configuracion_api/controllers:PerfilXMenuOpcionController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/configuracion_api/controllers:PerfilXMenuOpcionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/configuracion_api/controllers:PerfilXMenuOpcionController"],
        beego.ControllerComments{
            Method: "MenusPorAplicacion",
            Router: `/MenusPorAplicacion/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
