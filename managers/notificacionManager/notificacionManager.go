package notimanager

import (
	"fmt"
	"time"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/udistrital/configuracion_api/models"
)

func GetConfiguracion(endpoint, metodohttp,tipo,aplicacion,appname string) []models.NotificacionConfiguracion{
	o := orm.NewOrm()
	o.Begin()
	var configuracion []models.NotificacionConfiguracion
	beego.Info(endpoint)
	_, err := o.Raw(
		`SELECT T0.*, T0.*, T1.*, T2.*, T3.* FROM 
		configuracion.notificacion_configuracion T0 INNER JOIN 
		configuracion.metodo_http T1 ON T1.id = T0.metodo_http INNER JOIN 
		configuracion.notificacion_tipo T2 ON T2.id = T0.tipo INNER JOIN
		configuracion.aplicacion T3 ON T3.id = T0.aplicacion 
		WHERE T0.end_point LIKE '%notificacion_tipo%' AND T1.nombre = 'POST' AND T2.nombre = 'success' AND T3.nombre = 'configuracion_api' LIMIT 10;`).QueryRows(&configuracion)
	beego.Info(configuracion)
	if err != nil {
		fmt.Println("Error al consultar la configuracion")
	}

	if len(configuracion) == 0 || configuracion == nil{
		fmt.Println("No existen las configuraciones")

	} else {
		
	}
	return configuracion
}

func GetOldNotification(profile string, user string) []models.Notificacion {
	o := orm.NewOrm()
	o.Begin()
	var notificaciones []models.Notificacion

	_, err := o.Raw(
		`select N.* from ` + beego.AppConfig.String("PGschemas") + `.notificacion as N, (select N.id
			from 	` + beego.AppConfig.String("PGschemas") + `.notificacion as N,
					` + beego.AppConfig.String("PGschemas") + `.notificacion_configuracion as NC,
					` + beego.AppConfig.String("PGschemas") + `.notificacion_configuracion_perfil as NCP,
					` + beego.AppConfig.String("PGschemas") + `.perfil as p
			where
			NC.id = N.notificacion_configuracion and
			NCP.notificacion_configuracion = NC.id and
			NCP.perfil = p.id and
			p.nombre IN ('` + profile + `')
		except 
		select N.id from 
		` + beego.AppConfig.String("PGschemas") + `.notificacion as N, 
		` + beego.AppConfig.String("PGschemas") + `.notificacion_estado_usuario as NEU 
		where N.id = NEU.notificacion and NEU.usuario = '` + user + `') as IDS
		where IDS.id = N.id`).QueryRows(&notificaciones)
	beego.Info(notificaciones)
	if err != nil {
		fmt.Println("Error al consultar las notificaciones antiguas")
	}

	if len(notificaciones) == 0 || notificaciones == nil {
		fmt.Println("No existen notificaciones creadas")

	} else {
		for _, idx := range notificaciones {
			beego.Info("Entro por old notifications")
			relation := models.NotificacionEstadoUsuario{}
			relation.Activo = true
			relation.Fecha = time.Now().Local()
			relation.Notificacion = &models.Notificacion{Id: int(idx.Id)}
			relation.NotificacionEstado = &models.NotificacionEstado{Id: 1}
			relation.Usuario = user
			_, err = o.Insert(&relation)
			if err != nil {
				o.Rollback()
				panic("Error al insertar las notificaciones 2")
			}
		}
		if err != nil {
			o.Rollback()
			fmt.Println("Error al insertar las notificaciones antiguas")
		} else {
			o.Commit()
		}
	}
	return notificaciones
}

func PushNotificationUser(N *models.NotificacionUsuarioMasiva) error {
	o := orm.NewOrm()
	o.Begin()
	N.Notificacion.FechaCreacion = time.Now().Local()
	idN, err := o.Insert(N.Notificacion)
	if err != nil {
		o.Rollback()
		panic(err)
	}
	for _, idx := range N.Usuarios {
		relation := models.NotificacionEstadoUsuario{}
		relation.Activo = true
		relation.Fecha = time.Now().Local()
		relation.Notificacion = &models.Notificacion{Id: int(idN)}
		relation.NotificacionEstado = &models.NotificacionEstado{Id: 1}
		relation.Usuario = idx
		_, err = o.Insert(&relation)
		if err != nil {
			o.Rollback()
			panic("Error al insertar las notificaciones 2")
		}
	}

	if err != nil {
		o.Rollback()
		panic("Error al insertar las notificaciones 3")
	} else {
		o.Commit()
	}
	return err
}

func ChangeStateNoView(user string) []models.NotificacionEstadoUsuario {
	o := orm.NewOrm()
	o.Begin()
	var notificaciones []models.NotificacionEstadoUsuario
	_, err := o.Raw(
		`select NEU.* from 
		` + beego.AppConfig.String("PGschemas") + `.notificacion_estado_usuario as NEU 
		where NEU.usuario = '` + user + `' and
		activo = true and
		notificacion_estado = 1`).QueryRows(&notificaciones)
	beego.Info(notificaciones)
	if err != nil {
		fmt.Println("Error al actualizar las notificaciones antiguas")
	}

	if len(notificaciones) == 0 || notificaciones == nil {
		fmt.Println("No existen notificaciones_estado_usuario creadas")

	} else {
		for _, n := range notificaciones {
			notification := n
			if err = o.Read(&notification); err == nil {
				notification.Activo = false
				var num int64
				if num, err = o.Update(&notification); err == nil {
					fmt.Println("Number of records updated in database:", num)
				}
			}

			new_notification := models.NotificacionEstadoUsuario{}
			new_notification.Activo = true
			new_notification.Fecha = time.Now().Local()
			new_notification.Notificacion = n.Notificacion
			new_notification.NotificacionEstado = &models.NotificacionEstado{Id: 2}
			new_notification.Usuario = n.Usuario
			id, err := o.Insert(&new_notification)
			if err == nil {
				fmt.Println(id)
			}
		}
	}

	if err != nil {
		o.Rollback()
		panic("Error al actualizar las notificaciones ")
	} else {
		o.Commit()
	}

	return notificaciones
}
