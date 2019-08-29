package models

import (
	"fmt"

	"github.com/astaxie/beego/orm"
)

// GetAplicacionById retrieves Aplicacion by Id. Returns error if
// Id doesn't exist
func GetAplicacionByRol(roles []Perfil) (v []Aplicacion, err error) {

	var query string = ""

	// Get a QueryBuilder object. Takes DB driver name as parameter
	// Second return value is error, ignored here
	qb, _ := orm.NewQueryBuilder("mysql")

	for i, _ := range roles {

		if i < len(roles)-1 {
			query += "p.nombre='" + roles[i].Nombre + "' OR "
		} else {
			query += "p.nombre='" + roles[i].Nombre + "'"
		}

	}

	qb.Select("distinct app.*").
		From("configuracion.aplicacion app").
		InnerJoin("configuracion.perfil p").
		On("app.id = p.aplicacion").
		Where(query + " AND UPPER(app.dominio) NOT LIKE '%API%' ")

	sql := qb.String()

	fmt.Println(query)

	o := orm.NewOrm()
	var a []Aplicacion

	if _, err := o.Raw(sql).QueryRows(&a); err == nil {
		return a, err
	}

	return nil, err
}
