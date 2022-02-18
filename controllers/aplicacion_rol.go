package controllers

import (
	"encoding/json"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/udistrital/configuracion_api/models"
)

// AplicacionRolController operations for Aplicacion
type AplicacionRolController struct {
	beego.Controller
}

// URLMapping ...
func (c *AplicacionRolController) URLMapping() {
	c.Mapping("GetAplicacionByRol", c.GetAplicacionByRol)
}

// Post ...
// @Title Post
// @Description create Aplicacion
// @Param	body		body 	[]models.Perfil	true		"body for Perfil content"
// @Success 201 {object} []models.Perfil
// @Failure 400 the request contains incorrect syntax
// @router /aplicacion_rol [post]
func (c *AplicacionRolController) GetAplicacionByRol() {

	var roles []models.Perfil
	//var v models.Aplicacion

	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &roles); err == nil {

		if v, err := models.GetAplicacionByRol(roles); err == nil {

			c.Ctx.Output.SetStatus(201)
			c.Data["json"] = v
		} else {
			logs.Error(err)
			//c.Data["development"] = map[string]interface{}{"Code": "000", "Body": err.Error(), "Type": "error"}
			c.Data["system"] = err
			c.Abort("400")
		}
	} else {
		logs.Error(err)
		//c.Data["development"] = map[string]interface{}{"Code": "000", "Body": err.Error(), "Type": "error"}
		c.Data["system"] = err
		c.Abort("400")
	}
	c.ServeJSON()
}
