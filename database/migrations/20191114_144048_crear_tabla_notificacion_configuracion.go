package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type CrearTablaNotificacionConfiguracion_20191114_144048 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &CrearTablaNotificacionConfiguracion_20191114_144048{}
	m.Created = "20191114_144048"

	migration.Register("CrearTablaNotificacionConfiguracion_20191114_144048", m)
}

// Run the migrations
func (m *CrearTablaNotificacionConfiguracion_20191114_144048) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	m.SQL("CREATE TABLE IF NOT EXISTS configuracion.notificacion_configuracion ( id serial NOT NULL, end_point character varying NOT NULL, metodo_http integer NOT NULL, tipo integer NOT NULL, cuerpo_notificacion json NOT NULL, aplicacion integer NOT NULL, CONSTRAINT pk_notificacion_configuracion PRIMARY KEY (id), CONSTRAINT fk_configuracion_notificacion_metodo_http FOREIGN KEY (metodo_http) REFERENCES configuracion.metodo_http (id), CONSTRAINT fk_notificacion_configuracion_aplicacion FOREIGN KEY (aplicacion) REFERENCES configuracion.aplicacion (id), CONSTRAINT fk_notificacion_configuracion_notificacion_tipo FOREIGN KEY (tipo) REFERENCES configuracion.notificacion_tipo (id) );")
	m.SQL("ALTER TABLE configuracion.notificacion_configuracion OWNER TO desarrollooas;")
	m.SQL("COMMENT ON TABLE configuracion.notificacion_configuracion IS 'Configuracion de las notificaciones que seran emitidas';")
	m.SQL("COMMENT ON COLUMN configuracion.notificacion_configuracion.id IS 'Identificador único de la configuracion generada';")
	m.SQL("COMMENT ON COLUMN configuracion.notificacion_configuracion.end_point IS 'end_point del api que genera la notificacion';")
	m.SQL("COMMENT ON COLUMN configuracion.notificacion_configuracion.metodo_http IS 'Metodo http de la peticion que genera la notificación';")
	m.SQL("COMMENT ON COLUMN configuracion.notificacion_configuracion.tipo IS 'tipo de notificación';")
	m.SQL("COMMENT ON COLUMN configuracion.notificacion_configuracion.cuerpo_notificacion IS 'cuerpo de la notificacion (configuracion para mostrar en cliente y plantilla de mensaje)';")
	m.SQL("COMMENT ON COLUMN configuracion.notificacion_configuracion.aplicacion IS 'Aplicacion que genera la notificacion';")	
}

// Reverse the migrations
func (m *CrearTablaNotificacionConfiguracion_20191114_144048) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
	m.SQL("DROP TABLE IF EXISTS configuracion.notificacion_configuracion")
}
