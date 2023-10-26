package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

    beego.GlobalControllerRouter["github.com/udistrital/configuracion_api/controllers:AplicacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/configuracion_api/controllers:AplicacionController"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/configuracion_api/controllers:AplicacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/configuracion_api/controllers:AplicacionController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/configuracion_api/controllers:AplicacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/configuracion_api/controllers:AplicacionController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: "/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/configuracion_api/controllers:AplicacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/configuracion_api/controllers:AplicacionController"],
        beego.ControllerComments{
            Method: "Put",
            Router: "/:id",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/configuracion_api/controllers:AplicacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/configuracion_api/controllers:AplicacionController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/:id",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/configuracion_api/controllers:AplicacionRolController"] = append(beego.GlobalControllerRouter["github.com/udistrital/configuracion_api/controllers:AplicacionRolController"],
        beego.ControllerComments{
            Method: "GetAplicacionByRol",
            Router: "/aplicacion_rol",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/configuracion_api/controllers:EstadoProcesoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/configuracion_api/controllers:EstadoProcesoController"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/configuracion_api/controllers:EstadoProcesoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/configuracion_api/controllers:EstadoProcesoController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/configuracion_api/controllers:EstadoProcesoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/configuracion_api/controllers:EstadoProcesoController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: "/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/configuracion_api/controllers:EstadoProcesoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/configuracion_api/controllers:EstadoProcesoController"],
        beego.ControllerComments{
            Method: "Put",
            Router: "/:id",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/configuracion_api/controllers:EstadoProcesoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/configuracion_api/controllers:EstadoProcesoController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/:id",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/configuracion_api/controllers:MenuOpcionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/configuracion_api/controllers:MenuOpcionController"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/configuracion_api/controllers:MenuOpcionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/configuracion_api/controllers:MenuOpcionController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/configuracion_api/controllers:MenuOpcionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/configuracion_api/controllers:MenuOpcionController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: "/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/configuracion_api/controllers:MenuOpcionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/configuracion_api/controllers:MenuOpcionController"],
        beego.ControllerComments{
            Method: "Put",
            Router: "/:id",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/configuracion_api/controllers:MenuOpcionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/configuracion_api/controllers:MenuOpcionController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/:id",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/configuracion_api/controllers:MenuOpcionPadreController"] = append(beego.GlobalControllerRouter["github.com/udistrital/configuracion_api/controllers:MenuOpcionPadreController"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/configuracion_api/controllers:MenuOpcionPadreController"] = append(beego.GlobalControllerRouter["github.com/udistrital/configuracion_api/controllers:MenuOpcionPadreController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/configuracion_api/controllers:MenuOpcionPadreController"] = append(beego.GlobalControllerRouter["github.com/udistrital/configuracion_api/controllers:MenuOpcionPadreController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: "/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/configuracion_api/controllers:MenuOpcionPadreController"] = append(beego.GlobalControllerRouter["github.com/udistrital/configuracion_api/controllers:MenuOpcionPadreController"],
        beego.ControllerComments{
            Method: "Put",
            Router: "/:id",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/configuracion_api/controllers:MenuOpcionPadreController"] = append(beego.GlobalControllerRouter["github.com/udistrital/configuracion_api/controllers:MenuOpcionPadreController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/:id",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/configuracion_api/controllers:MenuOpcionPadreController"] = append(beego.GlobalControllerRouter["github.com/udistrital/configuracion_api/controllers:MenuOpcionPadreController"],
        beego.ControllerComments{
            Method: "ArbolMenus",
            Router: "/ArbolMenus/:roles/:app",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/configuracion_api/controllers:MenuOpcionPadreController"] = append(beego.GlobalControllerRouter["github.com/udistrital/configuracion_api/controllers:MenuOpcionPadreController"],
        beego.ControllerComments{
            Method: "ArbolMenuRoles",
            Router: "/menu_roles",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/configuracion_api/controllers:ParametroController"] = append(beego.GlobalControllerRouter["github.com/udistrital/configuracion_api/controllers:ParametroController"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/configuracion_api/controllers:ParametroController"] = append(beego.GlobalControllerRouter["github.com/udistrital/configuracion_api/controllers:ParametroController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/configuracion_api/controllers:ParametroController"] = append(beego.GlobalControllerRouter["github.com/udistrital/configuracion_api/controllers:ParametroController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: "/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/configuracion_api/controllers:ParametroController"] = append(beego.GlobalControllerRouter["github.com/udistrital/configuracion_api/controllers:ParametroController"],
        beego.ControllerComments{
            Method: "Put",
            Router: "/:id",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/configuracion_api/controllers:ParametroController"] = append(beego.GlobalControllerRouter["github.com/udistrital/configuracion_api/controllers:ParametroController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/:id",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/configuracion_api/controllers:PerfilController"] = append(beego.GlobalControllerRouter["github.com/udistrital/configuracion_api/controllers:PerfilController"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/configuracion_api/controllers:PerfilController"] = append(beego.GlobalControllerRouter["github.com/udistrital/configuracion_api/controllers:PerfilController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/configuracion_api/controllers:PerfilController"] = append(beego.GlobalControllerRouter["github.com/udistrital/configuracion_api/controllers:PerfilController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: "/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/configuracion_api/controllers:PerfilController"] = append(beego.GlobalControllerRouter["github.com/udistrital/configuracion_api/controllers:PerfilController"],
        beego.ControllerComments{
            Method: "Put",
            Router: "/:id",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/configuracion_api/controllers:PerfilController"] = append(beego.GlobalControllerRouter["github.com/udistrital/configuracion_api/controllers:PerfilController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/:id",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/configuracion_api/controllers:PerfilXMenuOpcionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/configuracion_api/controllers:PerfilXMenuOpcionController"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/configuracion_api/controllers:PerfilXMenuOpcionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/configuracion_api/controllers:PerfilXMenuOpcionController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/configuracion_api/controllers:PerfilXMenuOpcionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/configuracion_api/controllers:PerfilXMenuOpcionController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: "/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/configuracion_api/controllers:PerfilXMenuOpcionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/configuracion_api/controllers:PerfilXMenuOpcionController"],
        beego.ControllerComments{
            Method: "Put",
            Router: "/:id",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/configuracion_api/controllers:PerfilXMenuOpcionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/configuracion_api/controllers:PerfilXMenuOpcionController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/:id",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/configuracion_api/controllers:PerfilXMenuOpcionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/configuracion_api/controllers:PerfilXMenuOpcionController"],
        beego.ControllerComments{
            Method: "MenusPorAplicacion",
            Router: "/MenusPorAplicacion/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/configuracion_api/controllers:ProcesoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/configuracion_api/controllers:ProcesoController"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/configuracion_api/controllers:ProcesoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/configuracion_api/controllers:ProcesoController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/configuracion_api/controllers:ProcesoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/configuracion_api/controllers:ProcesoController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: "/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/configuracion_api/controllers:ProcesoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/configuracion_api/controllers:ProcesoController"],
        beego.ControllerComments{
            Method: "Put",
            Router: "/:id",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/configuracion_api/controllers:ProcesoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/configuracion_api/controllers:ProcesoController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/:id",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/configuracion_api/controllers:TransicionProcesoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/configuracion_api/controllers:TransicionProcesoController"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/configuracion_api/controllers:TransicionProcesoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/configuracion_api/controllers:TransicionProcesoController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/configuracion_api/controllers:TransicionProcesoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/configuracion_api/controllers:TransicionProcesoController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: "/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/configuracion_api/controllers:TransicionProcesoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/configuracion_api/controllers:TransicionProcesoController"],
        beego.ControllerComments{
            Method: "Put",
            Router: "/:id",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/configuracion_api/controllers:TransicionProcesoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/configuracion_api/controllers:TransicionProcesoController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/:id",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/configuracion_api/controllers:VersionProcesoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/configuracion_api/controllers:VersionProcesoController"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/configuracion_api/controllers:VersionProcesoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/configuracion_api/controllers:VersionProcesoController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/configuracion_api/controllers:VersionProcesoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/configuracion_api/controllers:VersionProcesoController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: "/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/configuracion_api/controllers:VersionProcesoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/configuracion_api/controllers:VersionProcesoController"],
        beego.ControllerComments{
            Method: "Put",
            Router: "/:id",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/configuracion_api/controllers:VersionProcesoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/configuracion_api/controllers:VersionProcesoController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/:id",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
