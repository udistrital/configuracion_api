package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type CrearTablaParametro_20191114_145932 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &CrearTablaParametro_20191114_145932{}
	m.Created = "20191114_145932"

	migration.Register("CrearTablaParametro_20191114_145932", m)
}

// Run the migrations
func (m *CrearTablaParametro_20191114_145932) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	m.SQL("CREATE TABLE IF NOT EXIST configuracion.parametro ( id serial NOT NULL, nombre character varying(200) NOT NULL, valor text NOT NULL, aplicacion integer NOT NULL, CONSTRAINT PK_PARAMETRO PRIMARY KEY (id), CONSTRAINT FK_PARAMETRO_APP FOREIGN KEY (aplicacion) REFERENCES configuracion.aplicacion (id) );")
	m.SQL("ALTER TABLE configuracion.parametro OWNER TO desarrollooas;")
	m.SQL("COMMENT ON TABLE configuracion.parametro IS 'Tabla que  describe todos los parametros asociados a una aplicacion';")
	m.SQL("COMMENT ON COLUMN configuracion.parametro.id IS 'Identificador de la tabla parametro';")
	m.SQL("COMMENT ON COLUMN configuracion.parametro.nombre IS 'Nombre o llave del parametro ';")
	m.SQL("COMMENT ON COLUMN configuracion.parametro.valor IS 'Valor del parametro';")
	m.SQL("COMMENT ON COLUMN configuracion.parametro.aplicacion IS 'Id de la aplicación';")
	m.SQL("COMMENT ON CONSTRAINT PK_PARAMETRO ON configuracion.parametro IS 'Llave primaria del parametro';")
}

// Reverse the migrations
func (m *CrearTablaParametro_20191114_145932) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
	m.SQL("DROP TABLE IF EXISTS configuracion.parametro")
}
