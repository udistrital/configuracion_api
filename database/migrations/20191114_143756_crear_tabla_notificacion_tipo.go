package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type CrearTablaNotificacionTipo_20191114_143756 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &CrearTablaNotificacionTipo_20191114_143756{}
	m.Created = "20191114_143756"

	migration.Register("CrearTablaNotificacionTipo_20191114_143756", m)
}

// Run the migrations
func (m *CrearTablaNotificacionTipo_20191114_143756) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	m.SQL("CREATE TABLE IF NOT EXISTS configuracion.notificacion_tipo ( id serial NOT NULL, nombre character varying(100), CONSTRAINT PK_notificacion_tipo PRIMARY KEY (id) );")
	m.SQL("ALTER TABLE configuracion.notificacion_tipo OWNER TO desarrollooas;")
	m.SQL("COMMENT ON TABLE configuracion.notificacion_tipo IS 'Centralización de los tipo de las notificaciones de los sistemas';")
	m.SQL("COMMENT ON COLUMN configuracion.notificacion_tipo.id IS 'Identificador único del tipo de la notificación';")
	m.SQL("COMMENT ON COLUMN configuracion.notificacion_tipo.nombre IS 'Nombre del tipo de la notificación';")
}

// Reverse the migrations
func (m *CrearTablaNotificacionTipo_20191114_143756) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
	m.SQL("DROP TABLE IF EXISTS configuracion.notificacion_tipo")
}
