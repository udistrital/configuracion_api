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
// @Failure 410 Controlador archivado
// @Param	body		body 	models.NotificacionEstadoUsuario	true		"body for NotificacionEstadoUsuario content"
// @Success 201 {int} models.NotificacionEstadoUsuario
// @Failure 400 the request contains incorrect syntax
// @router / [post]
func (c *NotificacionEstadoUsuarioController) Post() {
	c.Ctx.Output.SetStatus(http.StatusGone)
	c.ServeJSON()
	return

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
// @Failure 410 Controlador archivado
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.NotificacionEstadoUsuario
// @Failure 404 not found resource
// @router /:id [get]
func (c *NotificacionEstadoUsuarioController) GetOne() {
	c.Ctx.Output.SetStatus(http.StatusGone)
	c.ServeJSON()
	return

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
// @Failure 410 Controlador archivado
// @router / [get]
func (c *NotificacionEstadoUsuarioController) GetAll() {
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
// @Failure 410 Controlador archivado
// @router /:id [put]
func (c *NotificacionEstadoUsuarioController) Put() {
	c.Ctx.Output.SetStatus(http.StatusGone)
	c.ServeJSON()
	return

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
// @Failure 410 Controlador archivado
// @router /:id [delete]
func (c *NotificacionEstadoUsuarioController) Delete() {
	c.Ctx.Output.SetStatus(http.StatusGone)
	c.ServeJSON()
	return

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
// @Failure 410 Controlador archivado
// @router /getOldNotification/:roles/:usuario [get]
func (c *NotificacionEstadoUsuarioController) GetOldNotification() {
	c.Ctx.Output.SetStatus(http.StatusGone)
	c.ServeJSON()
	return

	rolesStr := c.Ctx.Input.Param(":roles")
	userStr := c.Ctx.Input.Param(":usuario")
	c.Data["json"] = notimanager.GetOldNotification(rolesStr, userStr)
}

// pushNotificationUser ...
// @Title pushNotificationUser
// @Description create pushNotificationUser
// @Failure 410 Controlador archivado
// @router /pushNotificationUser [post]
func (c *NotificacionEstadoUsuarioController) PushNotificationUser() {
	c.Ctx.Output.SetStatus(http.StatusGone)
	c.ServeJSON()
	return

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
// @Failure 410 Controlador archivado
// @router /changeStateNoView/:usuario [post]
func (c *NotificacionEstadoUsuarioController) ChangeStateNoView() {
	c.Ctx.Output.SetStatus(http.StatusGone)
	c.ServeJSON()
	return

	userStr := c.Ctx.Input.Param(":usuario")
	c.Data["json"] = notimanager.ChangeStateNoView(userStr)
}

// ChangeStateToView ...
// @Title changeStateNoView
// @Description create NotificacionEstadoUsuario
// @Failure 410 Controlador archivado
// @router /changeStateToView/:id [get]
func (c *NotificacionEstadoUsuarioController) ChangeStateToView() {
	c.Ctx.Output.SetStatus(http.StatusGone)
	c.ServeJSON()
	return

	idStr := c.Ctx.Input.Param(":id")
	c.Data["json"] = notimanager.ChangeStateToView(idStr)
}
