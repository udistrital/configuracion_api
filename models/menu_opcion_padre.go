package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"strconv"
	"github.com/astaxie/beego/orm"
)

type MenuOpcionPadre struct {
	Id    int         `orm:"column(id);pk;auto"`
	Padre *MenuOpcion `orm:"column(padre);rel(fk)"`
	Hijo  *MenuOpcion `orm:"column(hijo);rel(fk)"`
}

//Estructura para construir los menus
type Menu struct {
	Id int
    Nombre string
    Url string
    Opciones *[]Menu
}

func (t *MenuOpcionPadre) TableName() string {
	return "menu_opcion_padre"
}

func init() {
	orm.RegisterModel(new(MenuOpcionPadre))
}

// AddMenuOpcionPadre insert a new MenuOpcionPadre into database and returns
// last inserted Id on success.
func AddMenuOpcionPadre(m *MenuOpcionPadre) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetMenuOpcionPadreById retrieves MenuOpcionPadre by Id. Returns error if
// Id doesn't exist
func GetMenuOpcionPadreById(id int) (v *MenuOpcionPadre, err error) {
	o := orm.NewOrm()
	v = &MenuOpcionPadre{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllMenuOpcionPadre retrieves all MenuOpcionPadre matches certain condition. Returns empty list if
// no records exist
func GetAllMenuOpcionPadre(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(MenuOpcionPadre)).RelatedSel()
	// query k=v
	for k, v := range query {
		// rewrite dot-notation to Object__Attribute
		k = strings.Replace(k, ".", "__", -1)
		qs = qs.Filter(k, v)
	}
	// order by:
	var sortFields []string
	if len(sortby) != 0 {
		if len(sortby) == len(order) {
			// 1) for each sort field, there is an associated order
			for i, v := range sortby {
				orderby := ""
				if order[i] == "desc" {
					orderby = "-" + v
				} else if order[i] == "asc" {
					orderby = v
				} else {
					return nil, errors.New("Error: Invalid order. Must be either [asc|desc]")
				}
				sortFields = append(sortFields, orderby)
			}
			qs = qs.OrderBy(sortFields...)
		} else if len(sortby) != len(order) && len(order) == 1 {
			// 2) there is exactly one order, all the sorted fields will be sorted by this order
			for _, v := range sortby {
				orderby := ""
				if order[0] == "desc" {
					orderby = "-" + v
				} else if order[0] == "asc" {
					orderby = v
				} else {
					return nil, errors.New("Error: Invalid order. Must be either [asc|desc]")
				}
				sortFields = append(sortFields, orderby)
			}
		} else if len(sortby) != len(order) && len(order) != 1 {
			return nil, errors.New("Error: 'sortby', 'order' sizes mismatch or 'order' size is not 1")
		}
	} else {
		if len(order) != 0 {
			return nil, errors.New("Error: unused 'order' fields")
		}
	}

	var l []MenuOpcionPadre
	qs = qs.OrderBy(sortFields...)
	if _, err = qs.Limit(limit, offset).All(&l, fields...); err == nil {
		if len(fields) == 0 {
			for _, v := range l {
				ml = append(ml, v)
			}
		} else {
			// trim unused fields
			for _, v := range l {
				m := make(map[string]interface{})
				val := reflect.ValueOf(v)
				for _, fname := range fields {
					m[fname] = val.FieldByName(fname).Interface()
				}
				ml = append(ml, m)
			}
		}
		return ml, nil
	}
	return nil, err
}

// UpdateMenuOpcionPadre updates MenuOpcionPadre by Id and returns error if
// the record to be updated doesn't exist
func UpdateMenuOpcionPadreById(m *MenuOpcionPadre) (err error) {
	o := orm.NewOrm()
	v := MenuOpcionPadre{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteMenuOpcionPadre deletes MenuOpcionPadre by Id and returns error if
// the record to be deleted doesn't exist
func DeleteMenuOpcionPadre(id int) (err error) {
	o := orm.NewOrm()
	v := MenuOpcionPadre{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&MenuOpcionPadre{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}

//Función que construye los menús 
func ConstruirMenuPerfil(perfiles string) (menus []Menu){
	o := orm.NewOrm()		
	//Arreglo 
	var menusByPerfil []Menu		
	num, err := o.Raw(`Select mo.id AS id, mo.nombre AS nombre, mo.url AS URL, mop.padre AS padre
						FROM configuracion.perfil pe, configuracion.perfil_x_menu_opcion pmo,
						configuracion.aplicacion app, configuracion.menu_opcion 
						AS mo left join configuracion.menu_opcion_padre 
						AS mop ON mo.id = mop.hijo 
						WHERE app.id = mo.aplicacion
						AND pe.id = pmo.perfil AND mo.id = pmo.opcion
						AND pe.nombre IN ('`+ perfiles + `') AND padre IS NULL ORDER BY mo.id`).QueryRows(&menusByPerfil)


		//num, err := o.Raw("Select Distinct mo. id, mo.nombre, mo.url FROM configuracion.menu_opcion AS mo, configuracion.menu_opcion_padre AS mop, configuracion.perfil_x_menu_opcion AS pmo, configuracion.perfil pe, configuracion.aplicacion app where pe.nombre = '" + perfiles + "'AND mo.id = mop.padre AND pmo.opcion = mo.id AND pmo.perfil = pe.id AND app.id = mo.aplicacion AND mop.padre NOT IN (Select mop2.hijo FROM configuracion.menu_opcion_padre mop2) ORDER BY mo.id").QueryRows(&menusByPerfil)


		if err == nil {
	    fmt.Println("Menus padre encontrados: ", num)
		//For para que recorra los Ids en busca de hijos
			for i := 0; i < len(menusByPerfil); i++ {
				//Me verifica que los Id tengan hijos
				ConstruirSubMenusPerfil(&menusByPerfil[i], perfiles)
			}
		
		}
		return menusByPerfil
}

//Función que obtiene los menús padre de acuerdo al Id de la aplicación 
func MenusByAplicacion(app int) (menus []Menu){
	o := orm.NewOrm()

	//Conversión de entero a string
	aplicacion := strconv.Itoa(app)
	//Arreglo 
	var menusByApp []Menu		
		num, err := o.Raw(`Select Distinct mo.id AS id, mo.nombre AS nombre, mo.url AS URL, mop.padre AS padre
							FROM configuracion.aplicacion app, configuracion.menu_opcion 
							AS mo left join configuracion.menu_opcion_padre 
							AS mop ON mo.id = mop.hijo 
							WHERE `+ aplicacion + ` = mo.aplicacion
							AND padre IS NULL ORDER BY mo.id`).QueryRows(&menusByApp)

		if err == nil {
	    fmt.Println("Menus padre encontrados: ", num)
		//For para que recorra los Ids en busca de hijos
			for i := 0; i < len(menusByApp); i++ {
				//Me verifica que los Id tengan hijos
				ConstruirSubMenusPerfilApp(&menusByApp[i])
			}
		
		}
		return menusByApp
}


//Función que construye los Submenús 
func ConstruirSubMenusPerfil(Padre *Menu, perfiles string) (menus []Menu){
	o := orm.NewOrm()
	//Conversión de entero a string
	padre := strconv.Itoa(Padre.Id)
	
	//Arreglo 
	var subMenusByPerfil []Menu

		/*num, err := o.Raw(`SELECT mo.id, mo.nombre, mo.URL, mop.padre, mop.hijo 
			FROM configuracion.menu_opcion AS mo left join configuracion.menu_opcion_padre AS mop ON mo.id = mop.hijo 
			where mop.padre = "`+ padre + `" ORDER BY mo.id`).QueryRows(&subMenusByPerfil)*/

		num, err := o.Raw(`SELECT mo.id, mo.nombre, mo.URL, mop.padre, mop.hijo
			FROM configuracion.perfil_x_menu_opcion AS pmo, configuracion.perfil pe, configuracion.menu_opcion AS mo 
			left join configuracion.menu_opcion_padre AS mop ON mo.id = mop.hijo 
			where mop.padre = '`+ padre + `' AND pe.id = pmo.perfil AND mo.id = pmo.opcion AND pe.nombre IN ('`+ perfiles + `')
			ORDER BY mo.id`).QueryRows(&subMenusByPerfil)


		//Condicional si el error es nulo
		if err == nil {
			fmt.Println("Menus Hijos encontrados: ", num)

			//Llena el elemento Opciones en la estructura del menú padre
			Padre.Opciones = &subMenusByPerfil

			//For que recorre el subMenu en busca de hijos (Recursiva)
			for i := 0; i < len(subMenusByPerfil); i++ {

					//Me verifica que los Id tengan hijos
					ConstruirSubMenusPerfil(&subMenusByPerfil[i], perfiles)	
			}  
		}
		return subMenusByPerfil
}

//Función que construye los Submenús 
func ConstruirSubMenusPerfilApp(Padre *Menu) (menus []Menu){
	o := orm.NewOrm()
	//Conversión de entero a string
	padre := strconv.Itoa(Padre.Id)
	
	//Arreglo 
	var subMenusByPerfil []Menu

		num, err := o.Raw("SELECT mo.id, mo.nombre, mo.URL, mop.padre, mop.hijo FROM configuracion.menu_opcion AS mo left join configuracion.menu_opcion_padre AS mop ON mo.id = mop.hijo where mop.padre = '"+ padre + "' ORDER BY mo.id").QueryRows(&subMenusByPerfil)
		//Condicional si el error es nulo
		if err == nil {
			fmt.Println("Menus Hijos encontrados: ", num)

			//Llena el elemento Opciones en la estructura del menú padre
			Padre.Opciones = &subMenusByPerfil

			//For que recorre el subMenu en busca de hijos (Recursiva)
			for i := 0; i < len(subMenusByPerfil); i++ {

					//Me verifica que los Id tengan hijos
					ConstruirSubMenusPerfilApp(&subMenusByPerfil[i])	
			}  
		}
		return subMenusByPerfil
}
