package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type CrearTablaPerfil_20191114_144511 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &CrearTablaPerfil_20191114_144511{}
	m.Created = "20191114_144511"

	migration.Register("CrearTablaPerfil_20191114_144511", m)
}

// Run the migrations
func (m *CrearTablaPerfil_20191114_144511) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	m.SQL("CREATE TABLE IF NOT EXIST configuracion.perfil ( id serial NOT NULL, nombre character varying(50) NOT NULL, aplicacion integer NOT NULL, CONSTRAINT "PK_PERFIL" PRIMARY KEY (id), CONSTRAINT "FK_PERFIL_APP" FOREIGN KEY (aplicacion) REFERENCES configuracion.aplicacion (id) );")
	m.SQL("ALTER TABLE configuracion.perfil OWNER TO desarrollooas;")
	m.SQL("COMMENT ON TABLE configuracion.perfil IS 'Contiene la informacion de un perfil';")
	m.SQL("COMMENT ON COLUMN configuracion.perfil.id IS 'Identificador de un perfil';")
	m.SQL("COMMENT ON COLUMN configuracion.perfil.nombre IS 'Nombre del perfil';")
	m.SQL("COMMENT ON COLUMN configuracion.perfil.aplicacion IS 'Contiene el id de la aplicacion asociada al perfil';")
	m.SQL("COMMENT ON CONSTRAINT PK_PERFIL ON configuracion.perfil IS 'llave primaria de la tabla perfil';")
	
}

// Reverse the migrations
func (m *CrearTablaPerfil_20191114_144511) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
	m.SQL("DROP TABLE IF EXISTS configuracion.perfil")
}
