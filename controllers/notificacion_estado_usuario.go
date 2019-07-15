package controllers

import (
	"encoding/json"
	"errors"
	"strconv"
	"strings"

	"github.com/udistrital/configuracion_api/models"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	notimanager "github.com/udistrital/configuracion_api/managers/notificacionManager"
)

// NotificacionEstadoUsuarioController operations for NotificacionEstadoUsuario
type NotificacionEstadoUsuarioController struct {
	beego.Controller
}

// URLMapping ...
func (c *NotificacionEstadoUsuarioController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("GetOne", c.GetOne)
	c.Mapping("GetAll", c.GetAll)
	c.Mapping("Put", c.Put)
	c.Mapping("Delete", c.Delete)
}

// Post ...
// @Title Post
// @Description create NotificacionEstadoUsuario
// @Param	body		body 	models.NotificacionEstadoUsuario	true		"body for NotificacionEstadoUsuario content"
// @Success 201 {int} models.NotificacionEstadoUsuario
// @Failure 400 the request contains incorrect syntax
// @router / [post]
func (c *NotificacionEstadoUsuarioController) Post() {
	var v models.NotificacionEstadoUsuario
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {
		if _, err := models.AddNotificacionEstadoUsuario(&v); err == nil {
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
// @Description get NotificacionEstadoUsuario by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.NotificacionEstadoUsuario
// @Failure 404 not found resource
// @router /:id [get]
func (c *NotificacionEstadoUsuarioController) GetOne() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	v, err := models.GetNotificacionEstadoUsuarioById(id)
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

// GetAll ...
// @Title Get All
// @Description get NotificacionEstadoUsuario
// @Param	query	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// @Param	fields	query	string	false	"Fields returned. e.g. col1,col2 ..."
// @Param	sortby	query	string	false	"Sorted-by fields. e.g. col1,col2 ..."
// @Param	order	query	string	false	"Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ..."
// @Param	limit	query	string	false	"Limit the size of result set. Must be an integer"
// @Param	offset	query	string	false	"Start position of result set. Must be an integer"
// @Success 200 {object} models.NotificacionEstadoUsuario
// @Failure 404 not found resource
// @router / [get]
func (c *NotificacionEstadoUsuarioController) GetAll() {
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

	l, err := models.GetAllNotificacionEstadoUsuario(query, fields, sortby, order, offset, limit)
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
// @Description update the NotificacionEstadoUsuario
// @Param	id		path 	string	true		"The id you want to update"
// @Param	body		body 	models.NotificacionEstadoUsuario	true		"body for NotificacionEstadoUsuario content"
// @Success 200 {object} models.NotificacionEstadoUsuario
// @Failure 400 the request contains incorrect syntax
// @router /:id [put]
func (c *NotificacionEstadoUsuarioController) Put() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	v := models.NotificacionEstadoUsuario{Id: id}
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {
		if err := models.UpdateNotificacionEstadoUsuarioById(&v); err == nil {
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
// @Description delete the NotificacionEstadoUsuario
// @Param	id		path 	string	true		"The id you want to delete"
// @Success 200 {string} delete success!
// @Failure 404 not found resource
// @router /:id [delete]
func (c *NotificacionEstadoUsuarioController) Delete() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	if err := models.DeleteNotificacionEstadoUsuario(id); err == nil {
		c.Data["json"] = map[string]interface{}{"Id": id}
	} else {
		logs.Error(err)
		//c.Data["development"] = map[string]interface{}{"Code": "000", "Body": err.Error(), "Type": "error"}
		c.Data["system"] = err
		c.Abort("404")
	}
	c.ServeJSON()
}

// getOldNotification ...
// @Title getOldNotification
// @Description get all old notification
// @Param	roles		path 	string	true		"The profile you want to consult"
// @Param	usuario		path 	string	true		"The user you want to consult"
// @Success 200 {string} get success!
// @Failure 403 profile is empty
// @router /getOldNotification/:roles/:usuario [get]
func (c *NotificacionEstadoUsuarioController) GetOldNotification() {
	rolesStr := c.Ctx.Input.Param(":roles")
	userStr := c.Ctx.Input.Param(":usuario")
	c.Data["json"] = notimanager.GetOldNotification(rolesStr, userStr)
}

// pushNotificationUser ...
// @Title pushNotificationUser
// @Description create pushNotificationUser
// @Param	body		body 	models.NotificacionUsarioMasiva	true		"body for NotificacionEstadoUsuario content"
// @Success 201 {int} models.NotificacionEstadoUsuario
// @Failure 403 body is empty
// @router /pushNotificationUser [post]
func (c *NotificacionEstadoUsuarioController) PushNotificationUser() {
	var v models.NotificacionUsuarioMasiva
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {
		if err := notimanager.PushNotificationUser(&v); err == nil {
			c.Ctx.Output.SetStatus(201)
			c.Data["json"] = &v
		} else {
			c.Data["json"] = err.Error()
		}
	} else {
		c.Data["json"] = err.Error()
	}
}

// changeStateNoView ...
// @Title changeStateNoView
// @Description change state of notifications
// @Param	usuario		path 	string	true		"The user you want to update"
// @Success 200 {string} get success!
// @Failure 403 body is empty
// @router /changeStateNoView/:usuario [post]
func (c *NotificacionEstadoUsuarioController) ChangeStateNoView() {
	userStr := c.Ctx.Input.Param(":usuario")
	c.Data["json"] = notimanager.ChangeStateNoView(userStr)
}
