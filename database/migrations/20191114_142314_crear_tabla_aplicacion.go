package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type CrearTablaAplicacion_20191114_142314 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &CrearTablaAplicacion_20191114_142314{}
	m.Created = "20191114_142314"

	migration.Register("CrearTablaAplicacion_20191114_142314", m)
}

// Run the migrations
func (m *CrearTablaAplicacion_20191114_142314) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	m.SQL("CREATE TABLE IF NOT EXISTS configuracion.aplicacion ( id serial NOT NULL, nombre character varying(100) NOT NULL, descripcion text NOT NULL, dominio character varying(200) NOT NULL, estado boolean DEFAULT true, alias character varying(100), estilo_icono character varying(100), CONSTRAINT PK_APLICACION PRIMARY KEY (id), CONSTRAINT uq_aplicacion_dominio UNIQUE (dominio), CONSTRAINT uq_aplicacion_nombre UNIQUE (nombre) );")
	m.SQL("ALTER TABLE configuracion.aplicacion OWNER TO desarrollooas;")
	m.SQL("COMMENT ON TABLE configuracion.aplicacion IS 'Tabla que se encarga de almacenar la informacion de cada aplicacion contenida en la base de datos';")
	m.SQL("COMMENT ON COLUMN configuracion.aplicacion.id IS 'Identificador unico de la tabla aplicacion';")
	m.SQL("COMMENT ON COLUMN configuracion.aplicacion.nombre IS 'Nombre que identifica al aplicativo';")
	m.SQL("COMMENT ON COLUMN configuracion.aplicacion.descripcion IS 'Descripcion detallada de las funcionalidades del aplicativo';")
	m.SQL("COMMENT ON COLUMN configuracion.aplicacion.dominio IS 'Dominio o url de acceso publico del aplicativo';")
	m.SQL("COMMENT ON COLUMN configuracion.aplicacion.alias IS 'Dato que se presentará a cambio del nombre de la aplicación';")
	m.SQL("COMMENT ON COLUMN configuracion.aplicacion.estilo_icono IS 'Valor de la clase css que referenciará el ícono';")
	m.SQL("COMMENT ON CONSTRAINT PK_APLICACION ON configuracion.aplicacion IS 'Llave primaria de la tabla aplicacion';")
}

// Reverse the migrations
func (m *CrearTablaAplicacion_20191114_142314) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
	m.SQL("DROP TABLE IF EXISTS configuracion.aplicacion")
}
