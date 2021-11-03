package controllers

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"
	"strings"

	"github.com/udistrital/configuracion_api/models"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	notimanager "github.com/udistrital/configuracion_api/managers/notificacionManager"
)

// NotificacionConfiguracionController operations for NotificacionConfiguracion
type NotificacionConfiguracionController struct {
	beego.Controller
}

// URLMapping ...
func (c *NotificacionConfiguracionController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("GetOne", c.GetOne)
	c.Mapping("GetAll", c.GetAll)
	c.Mapping("Put", c.Put)
	c.Mapping("Delete", c.Delete)
}

// Post ...
// @Title Post
// @Description create NotificacionConfiguracion
// @Failure 410 Controlador archivado
// @router / [post]
func (c *NotificacionConfiguracionController) Post() {
	c.Ctx.Output.SetStatus(http.StatusGone)
	c.ServeJSON()
	return

	var v models.NotificacionConfiguracion
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {
		if _, err := models.AddNotificacionConfiguracion(&v); err == nil {
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

// GetOne ...
// @Title Get One
// @Description get NotificacionConfiguracion by id
// @Failure 410 Controlador archivado
// @router /:id [get]
func (c *NotificacionConfiguracionController) GetOne() {
	c.Ctx.Output.SetStatus(http.StatusGone)
	c.ServeJSON()
	return

	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	v, err := models.GetNotificacionConfiguracionById(id)
	if err != nil {
		logs.Error(err)
		//c.Data["development"] = map[string]interface{}{"Code": "000", "Body": err.Error(), "Type": "error"}
		c.Data["system"] = err
		c.Abort("404")
	} else {
		c.Data["json"] = v
	}
	c.ServeJSON()
}

// GetConfiguracion ...
// @Title getConfiguracion
// @Description get a configuration
// @Failure 410 Controlador archivado
// @router /getConfiguracion/ [post]
func (c *NotificacionConfiguracionController) GetConfiguracion() {
	c.Ctx.Output.SetStatus(http.StatusGone)
	c.ServeJSON()
	return

	var v map[string]interface{}
	// fields: col1,col2,entity.col3
	beego.Info(c.Ctx.Input.RequestBody)
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {

		c.Data["json"] = notimanager.GetConfiguracion(v["EndPoint"].(string), v["MetodoHttp"].(string), v["Tipo"].(string), v["Aplicacion"].(string))
		beego.Info(c.Data["json"])
	} else {
		beego.Error(err)
		logs.Error(err)
		//c.Data["development"] = map[string]interface{}{"Code": "000", "Body": err.Error(), "Type": "error"}
		c.Data["system"] = err
		c.Abort("400")
	}
	c.ServeJSON()
}

// GetAll ...
// @Title Get All
// @Description get NotificacionConfiguracion
// @Failure 410 Controlador archivado
// @router / [get]
func (c *NotificacionConfiguracionController) GetAll() {
	c.Ctx.Output.SetStatus(http.StatusGone)
	c.ServeJSON()
	return

	var fields []string
	var sortby []string
	var order []string
	var query = make(map[string]string)
	var limit int64 = 10
	var offset int64

	// fields: col1,col2,entity.col3
	if v := c.GetString("fields"); v != "" {
		fields = strings.Split(v, ",")
	}
	// limit: 10 (default is 10)
	if v, err := c.GetInt64("limit"); err == nil {
		limit = v
	}
	// offset: 0 (default is 0)
	if v, err := c.GetInt64("offset"); err == nil {
		offset = v
	}
	// sortby: col1,col2
	if v := c.GetString("sortby"); v != "" {
		sortby = strings.Split(v, ",")
	}
	// order: desc,asc
	if v := c.GetString("order"); v != "" {
		order = strings.Split(v, ",")
	}
	// query: k:v,k:v
	if v := c.GetString("query"); v != "" {
		for _, cond := range strings.Split(v, ",") {
			kv := strings.SplitN(cond, ":", 2)
			if len(kv) != 2 {
				c.Data["json"] = errors.New("Error: invalid query key/value pair")
				c.ServeJSON()
				return
			}
			k, v := kv[0], kv[1]
			query[k] = v
		}
	}

	l, err := models.GetAllNotificacionConfiguracion(query, fields, sortby, order, offset, limit)
	if err != nil {
		logs.Error(err)
		//c.Data["development"] = map[string]interface{}{"Code": "000", "Body": err.Error(), "Type": "error"}
		c.Data["system"] = err
		c.Abort("404")
	} else {
		if l == nil {
			l = append(l, map[string]interface{}{})
		}
		c.Data["json"] = l
	}
	c.ServeJSON()
}

// Put ...
// @Title Put
// @Description update the NotificacionConfiguracion
// @Failure 410 Controlador archivado
// @router /:id [put]
func (c *NotificacionConfiguracionController) Put() {
	c.Ctx.Output.SetStatus(http.StatusGone)
	c.ServeJSON()
	return

	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	v := models.NotificacionConfiguracion{Id: id}
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {
		if err := models.UpdateNotificacionConfiguracionById(&v); err == nil {
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

// Delete ...
// @Title Delete
// @Description delete the NotificacionConfiguracion
// @Failure 410 Controlador archivado
// @router /:id [delete]
func (c *NotificacionConfiguracionController) Delete() {
	c.Ctx.Output.SetStatus(http.StatusGone)
	c.ServeJSON()
	return

	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	if err := models.DeleteNotificacionConfiguracion(id); err == nil {
		c.Data["json"] = map[string]interface{}{"Id": id}
	} else {
		logs.Error(err)
		//c.Data["development"] = map[string]interface{}{"Code": "000", "Body": err.Error(), "Type": "error"}
		c.Data["system"] = err
		c.Abort("404")
	}
	c.ServeJSON()
}
