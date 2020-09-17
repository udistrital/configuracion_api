package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type CrearTablaNotificacionConfiguracionPerfil_20191114_145258 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &CrearTablaNotificacionConfiguracionPerfil_20191114_145258{}
	m.Created = "20191114_145258"

	migration.Register("CrearTablaNotificacionConfiguracionPerfil_20191114_145258", m)
}

// Run the migrations
func (m *CrearTablaNotificacionConfiguracionPerfil_20191114_145258) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	m.SQL("CREATE TABLE IF NOT EXISTS configuracion.notificacion_configuracion_perfil ( id serial NOT NULL, notificacion_configuracion integer NOT NULL, perfil integer NOT NULL, CONSTRAINT pk_notificacion_configuracion_perfil PRIMARY KEY (id), CONSTRAINT fk_notificacion_configuracion_perfil_notificacion_configuracion FOREIGN KEY (notificacion_configuracion) REFERENCES configuracion.notificacion_configuracion (id), CONSTRAINT fk_notificacion_configuracion_perfil_perfil FOREIGN KEY (perfil) REFERENCES configuracion.perfil (id) );")
	m.SQL("ALTER TABLE configuracion.notificacion_configuracion_perfil OWNER TO desarrollooas;")
	m.SQL("COMMENT ON TABLE configuracion.notificacion_configuracion_perfil IS 'Tipos de metodo http usados';")
	m.SQL("COMMENT ON COLUMN configuracion.notificacion_configuracion_perfil.id IS 'Identificador único de la tabla';")
	m.SQL("COMMENT ON COLUMN configuracion.notificacion_configuracion_perfil.notificacion_configuracion IS 'Referencia a la configuracion de la notificacion';")
	m.SQL("COMMENT ON COLUMN configuracion.notificacion_configuracion_perfil.perfil IS 'referencia al perfil que debe recibir la notificacion';")
}

// Reverse the migrations
func (m *CrearTablaNotificacionConfiguracionPerfil_20191114_145258) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
	m.SQL("DROP TABLE IF EXISTS configuracion.notificacion_configuracion_perfil")
}
