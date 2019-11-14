package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type CrearTablaPerfilXMenuOpcion_20191114_144720 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &CrearTablaPerfilXMenuOpcion_20191114_144720{}
	m.Created = "20191114_144720"

	migration.Register("CrearTablaPerfilXMenuOpcion_20191114_144720", m)
}

// Run the migrations
func (m *CrearTablaPerfilXMenuOpcion_20191114_144720) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	m.SQL("CREATE TABLE IF NOT EXIST configuracion.perfil_x_menu_opcion ( id serial NOT NULL, perfil integer NOT NULL, opcion integer NOT NULL, CONSTRAINT PK_PERFIL_X_MENU_OPCION PRIMARY KEY (id), CONSTRAINT UQ_PERFIL_X_MENU UNIQUE (perfil, opcion), CONSTRAINT FK_OPCION FOREIGN KEY (opcion) REFERENCES configuracion.menu_opcion (id), CONSTRAINT FK_PERFIL FOREIGN KEY (perfil) REFERENCES configuracion.perfil (id) );")
	m.SQL("ALTER TABLE configuracion.perfil_x_menu_opcion OWNER TO desarrollooas;")
	m.SQL("COMMENT ON TABLE configuracion.perfil_x_menu_opcion IS 'Tabla que maneja los menus que le pertenecen a ciertos perfiles';")
	m.SQL("COMMENT ON COLUMN configuracion.perfil_x_menu_opcion.id IS 'Identificador de la tabla perfil_x_menu_opcion';")
	m.SQL("COMMENT ON COLUMN configuracion.perfil_x_menu_opcion.perfil IS 'Contiene la informacion del perfil';")
	m.SQL("COMMENT ON COLUMN configuracion.perfil_x_menu_opcion.opcion IS 'Contiene la informacion del menu_opcion';")
	m.SQL("COMMENT ON CONSTRAINT PK_PERFIL_X_MENU_OPCION ON configuracion.perfil_x_menu_opcion IS 'Llave primera de la tabla de rompimiento';")
	m.SQL("COMMENT ON CONSTRAINT UQ_PERFIL_X_MENU ON configuracion.perfil_x_menu_opcion IS 'Restricción que solo deja asociar un registro de menú por perfil, garantiza que no haya replica.';")
}

// Reverse the migrations
func (m *CrearTablaPerfilXMenuOpcion_20191114_144720) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
	m.SQL("DROP TABLE IF EXISTS configuracion.perfil_x_menu_opcion")
}
