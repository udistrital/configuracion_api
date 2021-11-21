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

-- object: configuracion.version_proceso | type: TABLE --
-- DROP TABLE IF EXISTS configuracion.version_proceso CASCADE;
CREATE TABLE configuracion.version_proceso(
	id serial NOT NULL,
	proceso_id integer NOT NULL,
	version smallint NOT NULL,
	metadatos jsonb,
	activo boolean NOT NULL DEFAULT true,
	fecha_creacion timestamptz NOT NULL,
	fecha_modificacion timestamptz NOT NULL,
	CONSTRAINT uq_version_proceso UNIQUE (proceso_id,version),
	CONSTRAINT pk_version_proceso PRIMARY KEY (id)
);
-- ddl-end --
COMMENT ON TABLE configuracion.version_proceso IS 'La ultima/mayor sera de trabajo, las anteriores deberan considerarse finales, de solo lectura';
-- ddl-end --
COMMENT ON COLUMN configuracion.version_proceso.metadatos IS 'Propiedades que tienen sentido fuera de esta API';
-- ddl-end --
COMMENT ON CONSTRAINT uq_version_proceso ON configuracion.version_proceso  IS 'Version unica por cada proceso';
-- ddl-end --
-- ALTER TABLE configuracion.version_proceso OWNER TO postgres;
-- ddl-end --

-- object: fk_version_proceso_proceso | type: CONSTRAINT --
-- ALTER TABLE configuracion.version_proceso DROP CONSTRAINT IF EXISTS fk_version_proceso_proceso CASCADE;
ALTER TABLE configuracion.version_proceso ADD CONSTRAINT fk_version_proceso_proceso FOREIGN KEY (proceso_id)
REFERENCES configuracion.proceso (id) MATCH FULL
ON DELETE NO ACTION ON UPDATE NO ACTION;
-- ddl-end --

-- object: configuracion.estado_proceso | type: TABLE --
-- DROP TABLE IF EXISTS configuracion.estado_proceso CASCADE;
CREATE TABLE configuracion.estado_proceso(
	id serial NOT NULL,
	version_proceso_id integer NOT NULL,
	sigla varchar(10) NOT NULL,
	nombre varchar(50) NOT NULL,
	descripcion varchar(300),
	metadatos jsonb,
	activo bool NOT NULL DEFAULT true,
	fecha_creacion timestamptz NOT NULL,
	fecha_modificacion timestamptz NOT NULL,
	CONSTRAINT uq_sigla_version_proceso UNIQUE (version_proceso_id,sigla),
	CONSTRAINT pk_estado_proceso PRIMARY KEY (id)
);
-- ddl-end --
COMMENT ON COLUMN configuracion.estado_proceso.metadatos IS 'Propiedades que tienen sentido fuera de esta API';
-- ddl-end --
COMMENT ON CONSTRAINT uq_sigla_version_proceso ON configuracion.estado_proceso  IS 'Estado (mediante su sigla) unico entre la misma version';
-- ddl-end --
-- ALTER TABLE configuracion.estado_proceso OWNER TO postgres;
-- ddl-end --

-- object: fk_estado_proceso_version_proceso | type: CONSTRAINT --
-- ALTER TABLE configuracion.estado_proceso DROP CONSTRAINT IF EXISTS fk_estado_proceso_version_proceso CASCADE;
ALTER TABLE configuracion.estado_proceso ADD CONSTRAINT fk_estado_proceso_version_proceso FOREIGN KEY (version_proceso_id)
REFERENCES configuracion.version_proceso (id) MATCH FULL
ON DELETE NO ACTION ON UPDATE NO ACTION;
-- ddl-end --

-- object: configuracion.transicion_proceso | type: TABLE --
-- DROP TABLE IF EXISTS configuracion.transicion_proceso CASCADE;
CREATE TABLE configuracion.transicion_proceso(
	id serial NOT NULL,
	estado_proceso_id_anterior integer NOT NULL,
	estado_proceso_id_siguiente integer NOT NULL,
	sigla varchar(10) NOT NULL,
	nombre varchar(100) NOT NULL,
	descripcion varchar(300),
	metadatos jsonb,
	activo bool NOT NULL DEFAULT true,
	fecha_creacion timestamptz NOT NULL,
	fecha_modificacion timestamptz NOT NULL,
	CONSTRAINT pk_transicion PRIMARY KEY (id),
	CONSTRAINT uq_sigla_estados UNIQUE (estado_proceso_id_anterior,estado_proceso_id_siguiente,sigla)
);
-- ddl-end --
COMMENT ON COLUMN configuracion.transicion_proceso.metadatos IS 'Propiedades que tienen sentido fuera de esta API';
-- ddl-end --
COMMENT ON CONSTRAINT uq_sigla_estados ON configuracion.transicion_proceso  IS 'Combinacion unica entre la sigla de la transicion y los estados involucrados';
-- ddl-end --
-- ALTER TABLE configuracion.transicion_proceso OWNER TO postgres;
-- ddl-end --

-- object: fk_transicion_proceso_estado_proceso_siguiente | type: CONSTRAINT --
-- ALTER TABLE configuracion.transicion_proceso DROP CONSTRAINT IF EXISTS fk_transicion_proceso_estado_proceso_siguiente CASCADE;
ALTER TABLE configuracion.transicion_proceso ADD CONSTRAINT fk_transicion_proceso_estado_proceso_siguiente FOREIGN KEY (estado_proceso_id_siguiente)
REFERENCES configuracion.estado_proceso (id) MATCH FULL
ON DELETE NO ACTION ON UPDATE NO ACTION;
-- ddl-end --

-- object: fk_transicion_proceso_estado_proceso_anterior | type: CONSTRAINT --
-- ALTER TABLE configuracion.transicion_proceso DROP CONSTRAINT IF EXISTS fk_transicion_proceso_estado_proceso_anterior CASCADE;
ALTER TABLE configuracion.transicion_proceso ADD CONSTRAINT fk_transicion_proceso_estado_proceso_anterior FOREIGN KEY (estado_proceso_id_anterior)
REFERENCES configuracion.estado_proceso (id) MATCH FULL
ON DELETE NO ACTION ON UPDATE NO ACTION;
-- ddl-end --
