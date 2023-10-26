package controllers

import (
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"

	"github.com/udistrital/configuracion_api/models"
)

// MenuOpcionPadreController oprations for MenuOpcionPadre
type MenuOpcionPadreController struct {
	beego.Controller
}

// URLMapping ...
func (c *MenuOpcionPadreController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("GetOne", c.GetOne)
	c.Mapping("GetAll", c.GetAll)
	c.Mapping("Put", c.Put)
	c.Mapping("Delete", c.Delete)
	c.Mapping("ArbolMenus", c.ArbolMenus)
}

// Post ...
// @Title Post
// @Description create MenuOpcionPadre
// @Param	body		body 	models.MenuOpcionPadre	true		"body for MenuOpcionPadre content"
// @Success 201 {object} models.MenuOpcionPadre
// @Failure 400 the request contains incorrect syntax
// @router / [post]
func (c *MenuOpcionPadreController) Post() {
	var v models.MenuOpcionPadre
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {
		if _, err := models.AddMenuOpcionPadre(&v); err == nil {
			c.Ctx.Output.SetStatus(201)
			c.Data["json"] = v
		} else {
			logs.Error(err)
			//c.Data["development"] = map[string]interface{}{"Code": "000", "Body": err.Error(), "Type": "error"}
			c.Data["system"] = err
			c.Abort("400")
			fmt.Println(err)
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
// @Description get MenuOpcionPadre by id
// @Param	id		path 	int	true		"The key for staticblock"
// @Success 200 {object} models.MenuOpcionPadre
// @Failure 404 not found resource
// @router /:id [get]
func (c *MenuOpcionPadreController) GetOne() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	v, err := models.GetMenuOpcionPadreById(id)
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
// @Description get MenuOpcionPadre
// @Param	query	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// @Param	fields	query	string	false	"Fields returned. e.g. col1,col2 ..."
// @Param	sortby	query	string	false	"Sorted-by fields. e.g. col1,col2 ..."
// @Param	order	query	string	false	"Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ..."
// @Param	limit	query	string	false	"Limit the size of result set. Must be an integer"
// @Param	offset	query	string	false	"Start position of result set. Must be an integer"
// @Success 200 {object} []models.MenuOpcionPadre
// @Failure 404 not found resource
// @router / [get]
func (c *MenuOpcionPadreController) GetAll() {
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

	l, err := models.GetAllMenuOpcionPadre(query, fields, sortby, order, offset, limit)
	if err != nil {
		logs.Error(err)
		//c.Data["development"] = map[string]interface{}{"Code": "000", "Body": err.Error(), "Type": "error"}
		c.Data["system"] = err
		c.Abort("404")
	} else {
		if l == nil {
			l = []interface{}{}
		}
		c.Data["json"] = l
	}
	c.ServeJSON()
}

// Put ...
// @Title Put
// @Description update the MenuOpcionPadre
// @Param	id		path 	int	true		"The id you want to update"
// @Param	body		body 	models.MenuOpcionPadre	true		"body for MenuOpcionPadre content"
// @Success 200 {object} models.MenuOpcionPadre
// @Failure 400 the request contains incorrect syntax
// @router /:id [put]
func (c *MenuOpcionPadreController) Put() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	v := models.MenuOpcionPadre{Id: id}
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {
		if err := models.UpdateMenuOpcionPadreById(&v); err == nil {
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
// @Description delete the MenuOpcionPadre
// @Param	id		path 	int	true		"The id you want to delete"
// @Success 200 {string} delete success!
// @Failure 404 not found resource
// @router /:id [delete]
func (c *MenuOpcionPadreController) Delete() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	if err := models.DeleteMenuOpcionPadre(id); err == nil {
		c.Data["json"] = map[string]interface{}{"Id": id}
	} else {
		logs.Error(err)
		//c.Data["development"] = map[string]interface{}{"Code": "000", "Body": err.Error(), "Type": "error"}
		c.Data["system"] = err
		c.Abort("404")
	}
	c.ServeJSON()
}

// ArbolMenus ...
// @Title ArbolMenus
// @Description Genera árbol de menus de acuerdo a roles y app
// @Param	roles		path 	string	true		"Los roles a los que se necesita que se desplieguen las opciones"
// @Param	app		path 	string	true		"Aplicación a la cual pertenecen las opciones"
// @Success 200 {object} []models.Menu
// @Failure 403 :roles is empty
// @Failure 403 :app is empty
// @router /ArbolMenus/:roles/:app [get]
func (c *MenuOpcionPadreController) ArbolMenus() {
	roles := c.Ctx.Input.Param(":roles")
	app := c.Ctx.Input.Param(":app")
	perfilesR := strings.NewReplacer(",", "','")

	//Construcción Json menus
	l := models.ConstruirMenuPerfil(perfilesR.Replace(roles), app)
	fmt.Println("Este es el resultado de la consulta")
	fmt.Println(l)

	if l == nil {
		c.Data["json"] = []interface{}{}
	} else {
		c.Data["json"] = l
	}
	//Generera el Json con los datos obtenidos
	c.ServeJSON()
}

// ArbolMenuRoles
// @Title ArbolMenuRoles
// @Description Genera árbol de permisos concedidos a una lista de roles en una apliación determinada. Los parámetros se pasan como parámetros de tipo query y no path
// @Param	app		query	string	true	"Aplicación de la que se consulta el menú"
// @Param	roles	query	string	true	"Lista de roles a los que se desea consultar el menú"
// @Success 200 {object} []models.Menu
// @Failure 400 roles is empty
// @Failure 400 app is empty
// @router /menu_roles [get]
func (c *MenuOpcionPadreController) ArbolMenuRoles() {
	roles := c.GetString("roles")
	app := c.GetString("app")

	if app == "" || roles == "" {
		err := "Debe especificar una aplicación y por lo menos un rol"
		logs.Error(err)
		c.Data["system"] = err
		c.Abort("401")
	}

	perfilesR := strings.NewReplacer(",", "','")
	l := models.ConstruirMenuPerfil(perfilesR.Replace(roles), app)
	if l == nil {
		c.Data["json"] = []interface{}{}
	} else {
		c.Data["json"] = l
	}

	c.ServeJSON()
}
