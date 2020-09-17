package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type CrearTablaNotificacion_20191114_144307 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &CrearTablaNotificacion_20191114_144307{}
	m.Created = "20191114_144307"

	migration.Register("CrearTablaNotificacion_20191114_144307", m)
}

// Run the migrations
func (m *CrearTablaNotificacion_20191114_144307) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	m.SQL("CREATE TABLE IF NOT EXISTS configuracion.notificacion ( id serial NOT NULL, fecha_creacion timestamp with time zone NOT NULL DEFAULT now(), cuerpo_notificacion json, notificacion_configuracion integer NOT NULL, CONSTRAINT PK_notificacion PRIMARY KEY (id), CONSTRAINT fk_notificacion_notificacion_configuracion FOREIGN KEY (notificacion_configuracion) REFERENCES configuracion.notificacion_configuracion (id) );")
	m.SQL("ALTER TABLE configuracion.notificacion OWNER TO desarrollooas;")
	m.SQL("COMMENT ON TABLE configuracion.notificacion IS 'Notificaciones de sistema a usuario centralizadas para todos los sistemas';")
	m.SQL("COMMENT ON COLUMN configuracion.notificacion.id IS 'Identificador único de la notificación generada';")
	m.SQL("COMMENT ON COLUMN configuracion.notificacion.fecha_creacion IS 'Fecha y hora en la que se creó la notificación';")
	m.SQL("COMMENT ON COLUMN configuracion.notificacion.cuerpo_notificacion IS 'Cuerpo de la notificación, será un objeto JSON en el que se almacenará el mensaje, título y otra información relevante';")
	m.SQL("COMMENT ON COLUMN configuracion.notificacion.notificacion_configuracion IS 'Referencia a la configuracion de la notificacion generada';")	
}

// Reverse the migrations
func (m *CrearTablaNotificacion_20191114_144307) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
	m.SQL("DROP TABLE IF EXISTS configuracion.notificacion")
}
