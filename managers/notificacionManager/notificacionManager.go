package rubromanager

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

func getOldNotification(profile) {
	o := orm.NewOrm()
		`SELECT A.*
		FROM (select N.*
			from 	`+ beego.AppConfig.String("PGschemas") + `.notificacion as N,
					`+ beego.AppConfig.String("PGschemas") + `.notificacion_configuracion as NC,
					`+ beego.AppConfig.String("PGschemas") + `.notificacion_configuracion_perfil as NCP,
					`+ beego.AppConfig.String("PGschemas") + `.perfil as p
			where
			NC.id = N.notificacion_`+ beego.AppConfig.String("PGschemas") + ` and
			NCP.notificacion_`+ beego.AppConfig.String("PGschemas") + ` = NC.id and
			NCP.perfil = p.id and
			p.nombre IN (` + profile +`)
			) as A
		left join `+ beego.AppConfig.String("PGschemas") + `.notificacion_estado_usuario as B on 
		A.id = B.notificacion where B.id is null`
}