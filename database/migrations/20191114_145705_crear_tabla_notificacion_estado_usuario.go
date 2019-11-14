package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type CrearTablaNotificacionEstadoUsuario_20191114_145705 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &CrearTablaNotificacionEstadoUsuario_20191114_145705{}
	m.Created = "20191114_145705"

	migration.Register("CrearTablaNotificacionEstadoUsuario_20191114_145705", m)
}

// Run the migrations
func (m *CrearTablaNotificacionEstadoUsuario_20191114_145705) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	m.SQL("CREATE TABLE IF NOT EXIST configuracion.notificacion_estado_usuario ( id serial NOT NULL, notificacion integer NOT NULL, notificacion_estado integer NOT NULL, fecha timestamp without time zone NOT NULL, usuario character varying(100), activo boolean DEFAULT true, CONSTRAINT pk_notificacion_estado_usuario PRIMARY KEY (id), CONSTRAINT fk_notificacion_estado_usuario_notificacion FOREIGN KEY (notificacion) REFERENCES configuracion.notificacion (id), CONSTRAINT fk_notificacion_estado_usuario_notificacion_estado FOREIGN KEY (notificacion_estado) REFERENCES configuracion.notificacion_estado (id) );")
	m.SQL("ALTER TABLE configuracion.notificacion_estado_usuario OWNER TO desarrollooas;")
	m.SQL("COMMENT ON TABLE configuracion.notificacion_estado_usuario IS 'Tabla que relaciona una notificación, un estado y un usuario. Permite así conservar el histórico de la notificación desde que se crea hasta que es leída';")
	m.SQL("COMMENT ON COLUMN configuracion.notificacion_estado_usuario.id IS 'Identificador único de notificacion_estado_usuario';")
	m.SQL("COMMENT ON COLUMN configuracion.notificacion_estado_usuario.notificacion IS 'Id de la notificación registrada';")
	m.SQL("COMMENT ON COLUMN configuracion.notificacion_estado_usuario.notificacion_estado IS 'Estado de la notificación registrada';")
	m.SQL("COMMENT ON COLUMN configuracion.notificacion_estado_usuario.fecha IS 'Fecha de la notificación ';")
	m.SQL("COMMENT ON COLUMN configuracion.notificacion_estado_usuario.usuario IS 'Usuario de WSO2 que interactúa con la notificación registrada';")
	m.SQL("COMMENT ON COLUMN configuracion.notificacion_estado_usuario.activo IS 'indicador de si el registro se encuentra en estado activo';")
}

// Reverse the migrations
func (m *CrearTablaNotificacionEstadoUsuario_20191114_145705) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
	m.SQL("DROP TABLE IF EXISTS configuracion.notificacion_estado_usuario")
}
