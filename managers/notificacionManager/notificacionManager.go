package notimanager

import (
	"fmt"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/udistrital/configuracion_api/models"
)

func GetOldNotification(profile string) []models.Notificacion {
	o := orm.NewOrm()
	o.Begin()
	var notificaciones []models.Notificacion

	_, err := o.Raw(`SELECT A.*
		FROM (select N.*
			from 	` + beego.AppConfig.String("PGschemas") + `.notificacion as N,
					` + beego.AppConfig.String("PGschemas") + `.notificacion_configuracion as NC,
					` + beego.AppConfig.String("PGschemas") + `.notificacion_configuracion_perfil as NCP,
					` + beego.AppConfig.String("PGschemas") + `.perfil as p
			where
			NC.id = N.notificacion_configuracion and
			NCP.notificacion_configuracion = NC.id and
			NCP.perfil = p.id and
			p.nombre IN ('` + profile + `')
			) as A
		left join ` + beego.AppConfig.String("PGschemas") + `.notificacion_estado_usuario as B on 
		A.id = B.notificacion where B.id is null`).QueryRows(&notificaciones)
	beego.Info(notificaciones)
	if err != nil {
		fmt.Println("Error al consultar las notificaciones antiguas")
	}

	if len(notificaciones) == 0 || notificaciones == nil {
		fmt.Println("No existen notificaciones creadas")

	}
	return notificaciones
}
