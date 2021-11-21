-- object: configuracion.proceso | type: TABLE --
-- DROP TABLE IF EXISTS configuracion.proceso CASCADE;
CREATE TABLE configuracion.proceso(
	id serial NOT NULL,
	aplicacion_id integer NOT NULL,
	sigla varchar(10) NOT NULL,
	nombre varchar(100) NOT NULL,
	descripcion varchar(300),
	metadatos jsonb,
	activo bool NOT NULL DEFAULT true,
	fecha_creacion timestamptz NOT NULL,
	fecha_modificacion timestamptz NOT NULL,
	CONSTRAINT uq_sigla_aplicacion_proceso UNIQUE (aplicacion_id,sigla),
	CONSTRAINT proceso_pk PRIMARY KEY (id)
);

-- object: fk_proceso_aplicacion | type: CONSTRAINT --
-- ALTER TABLE configuracion.proceso DROP CONSTRAINT IF EXISTS fk_proceso_aplicacion CASCADE;
ALTER TABLE configuracion.proceso ADD CONSTRAINT fk_proceso_aplicacion FOREIGN KEY (aplicacion_id)
REFERENCES configuracion.aplicacion (id) MATCH FULL
ON DELETE NO ACTION ON UPDATE NO ACTION;
-- ddl-end --

-- ddl-end --
COMMENT ON TABLE configuracion.proceso IS 'Procesos asociados a una aplicacion';
-- ddl-end --
COMMENT ON COLUMN configuracion.proceso.sigla IS 'Abreviacion del proceso. Una vez creada no deberia modificarse!';
-- ddl-end --
COMMENT ON COLUMN configuracion.proceso.metadatos IS 'Propiedades que tienen sentido fuera de esta API';
-- ddl-end --
COMMENT ON CONSTRAINT uq_sigla_aplicacion_proceso ON configuracion.proceso  IS 'No mas de un mismo proceso (identificado por sigla) por aplicacion';
-- ddl-end --
-- ALTER TABLE configuracion.proceso OWNER TO postgres;
-- ddl-end --
