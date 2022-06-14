-- Diff code generated with pgModeler (PostgreSQL Database Modeler)
-- pgModeler version: 0.9.1-beta
-- Diff date: 2022-02-18 20:06:34
-- Source model: test
-- Database: teest
-- PostgreSQL version: 10.0

-- [ Diff summary ]
-- Dropped objects: 21
-- Created objects: 0
-- Changed objects: 4
-- Truncated tables: 0

SET search_path=public,configuracion;
-- ddl-end --


-- [ Dropped objects ] --
ALTER TABLE IF EXISTS configuracion.notificacion DROP CONSTRAINT IF EXISTS fk_notificacion_notificacion_configuracion CASCADE;
-- ddl-end --
ALTER TABLE IF EXISTS configuracion.notificacion_estado_usuario DROP CONSTRAINT IF EXISTS fk_notificacion_estado_usuario_notificacion_estado CASCADE;
-- ddl-end --
ALTER TABLE IF EXISTS configuracion.notificacion_estado_usuario DROP CONSTRAINT IF EXISTS fk_notificacion_estado_usuario_notificacion CASCADE;
-- ddl-end --
ALTER TABLE IF EXISTS configuracion.notificacion_configuracion_perfil DROP CONSTRAINT IF EXISTS fk_notificacion_configuracion_perfil_perfil CASCADE;
-- ddl-end --
ALTER TABLE IF EXISTS configuracion.notificacion_configuracion_perfil DROP CONSTRAINT IF EXISTS fk_notificacion_configuracion_perfil_notificacion_configuracion CASCADE;
-- ddl-end --
ALTER TABLE IF EXISTS configuracion.notificacion_configuracion DROP CONSTRAINT IF EXISTS fk_notificacion_configuracion_notificacion_tipo CASCADE;
-- ddl-end --
ALTER TABLE IF EXISTS configuracion.notificacion_configuracion DROP CONSTRAINT IF EXISTS fk_notificacion_configuracion_aplicacion CASCADE;
-- ddl-end --
ALTER TABLE IF EXISTS configuracion.notificacion_configuracion DROP CONSTRAINT IF EXISTS fk_configuracion_notificacion_metodo_http CASCADE;
-- ddl-end --
DROP TABLE IF EXISTS configuracion.notificacion_tipo CASCADE;
-- ddl-end --
DROP SEQUENCE IF EXISTS configuracion.notificacion_tipo_id_seq CASCADE;
-- ddl-end --
DROP TABLE IF EXISTS configuracion.notificacion CASCADE;
-- ddl-end --
DROP TABLE IF EXISTS configuracion.notificacion_estado_usuario CASCADE;
-- ddl-end --
DROP SEQUENCE IF EXISTS configuracion.notificacion_estado_usuario_id_seq CASCADE;
-- ddl-end --
DROP TABLE IF EXISTS configuracion.notificacion_estado CASCADE;
-- ddl-end --
DROP SEQUENCE IF EXISTS configuracion.notificacion_estado_id_seq CASCADE;
-- ddl-end --
DROP TABLE IF EXISTS configuracion.notificacion_configuracion_perfil CASCADE;
-- ddl-end --
DROP SEQUENCE IF EXISTS configuracion.notificacion_configuracion_perfil_id_seq CASCADE;
-- ddl-end --
DROP TABLE IF EXISTS configuracion.notificacion_configuracion CASCADE;
-- ddl-end --
DROP SEQUENCE IF EXISTS configuracion.notificacion_configuracion_id_seq CASCADE;
-- ddl-end --
DROP SEQUENCE IF EXISTS configuracion.notificacion_id_seq CASCADE;
-- ddl-end --
DROP TABLE IF EXISTS configuracion.metodo_http CASCADE;
-- ddl-end --
