package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type CrearTablaNotificacionEstado_20191114_145459 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &CrearTablaNotificacionEstado_20191114_145459{}
	m.Created = "20191114_145459"

	migration.Register("CrearTablaNotificacionEstado_20191114_145459", m)
}

// Run the migrations
func (m *CrearTablaNotificacionEstado_20191114_145459) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	m.SQL("CREATE TABLE IF NOT EXIST configuracion.notificacion_estado ( id serial NOT NULL, nombre character varying(100), codigo_abreviacion character varying(20), descripcion character varying(250), activo boolean DEFAULT true, numero_orden numeric(5,2), CONSTRAINT PK_notificacion_estado PRIMARY KEY (id) );")
	m.SQL("ALTER TABLE configuracion.notificacion_estado OWNER TO desarrollooas;")
	m.SQL("COMMENT ON TABLE configuracion.notificacion_estado IS 'Centralización de los estados de las notificaciones de los sistemas';")
	m.SQL("COMMENT ON COLUMN configuracion.notificacion_estado.id IS 'Identificador único del estado de la notificación';")
	m.SQL("COMMENT ON COLUMN configuracion.notificacion_estado.nombre IS 'Nombre del estado de la notificación';")
	m.SQL("COMMENT ON COLUMN configuracion.notificacion_estado.codigo_abreviacion IS 'Código de abreviación del estado de la notificación';")
	m.SQL("COMMENT ON COLUMN configuracion.notificacion_estado.descripcion IS 'Descripción del estado de notificación';")
	m.SQL("COMMENT ON COLUMN configuracion.notificacion_estado.activo IS 'indicador de si el registro se encuentra en estado activo';")
	m.SQL("COMMENT ON COLUMN configuracion.notificacion_estado.numero_orden IS 'Orden de los estados de notificación';")
}

// Reverse the migrations
func (m *CrearTablaNotificacionEstado_20191114_145459) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
	m.SQL("DROP TABLE IF EXISTS configuracion.notificacion_estado")
}
