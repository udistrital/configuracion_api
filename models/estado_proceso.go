package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/astaxie/beego/orm"
)

type EstadoProceso struct {
	Id                int64           `orm:"column(id);auto"`
	VersionProcesoId  *VersionProceso `orm:"column(version_proceso_id);rel(fk)"`
	Sigla             string          `orm:"column(sigla);size(10)"`
	Nombre            string          `orm:"column(nombre);size(50)"`
	Descripcion       string          `orm:"column(descripcion);size(300);null"`
	Metadatos         string          `orm:"column(metadatos);type(jsonb);null"`
	Activo            bool            `orm:"column(activo)"`
	FechaCreacion     time.Time       `orm:"column(fecha_creacion);auto_now_add;type(datetime)"`
	FechaModificacion time.Time       `orm:"column(fecha_modificacion);auto_now;type(datetime)"`
}

func init() {
	orm.RegisterModel(new(EstadoProceso))
}

// AddEstadoProceso insert a new EstadoProceso into database and returns
// last inserted Id on success.
func AddEstadoProceso(m *EstadoProceso) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetEstadoProcesoById retrieves EstadoProceso by Id. Returns error if
// Id doesn't exist
func GetEstadoProcesoById(id int64) (v *EstadoProceso, err error) {
	o := orm.NewOrm()
	v = &EstadoProceso{Id: id}
	if err = o.QueryTable(new(EstadoProceso)).Filter("Id", id).RelatedSel().One(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllEstadoProceso retrieves all EstadoProceso matches certain condition. Returns empty list if
// no records exist
func GetAllEstadoProceso(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(EstadoProceso))
	// query k=v
	for k, v := range query {
		// rewrite dot-notation to Object__Attribute
		k = strings.Replace(k, ".", "__", -1)
		if strings.Contains(k, "isnull") {
			qs = qs.Filter(k, (v == "true" || v == "1"))
		} else if strings.Contains(k, "__in") {
			arr := strings.Split(v, "|")
			qs = qs.Filter(k, arr)
		} else if strings.Contains(k, "__not_in") {
			k = strings.Replace(k, "__not_in", "", -1)
			qs = qs.Exclude(k, v)
		} else {
			qs = qs.Filter(k, v)
		}
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

	var l []EstadoProceso
	qs = qs.OrderBy(sortFields...).RelatedSel()
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

// UpdateEstadoProceso updates EstadoProceso by Id and returns error if
// the record to be updated doesn't exist
func UpdateEstadoProcesoById(m *EstadoProceso) (err error) {
	o := orm.NewOrm()
	v := EstadoProceso{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteEstadoProceso deletes EstadoProceso by Id and returns error if
// the record to be deleted doesn't exist
func DeleteEstadoProceso(id int64) (err error) {
	o := orm.NewOrm()
	v := EstadoProceso{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&EstadoProceso{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
