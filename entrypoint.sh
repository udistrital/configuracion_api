#!/usr/bin/env bash

set -e
set -u
set -o pipefail

if [ -n "${PARAMETER_STORE:-}" ]; then
  export CONFIGURACION_API__PGUSER="$(aws ssm get-parameter --name /${PARAMETER_STORE}/configuracion_api/db/username --output text --query Parameter.Value)"
  export CONFIGURACION_API__PGPASS="$(aws ssm get-parameter --with-decryption --name /${PARAMETER_STORE}/configuracion_api/db/password --output text --query Parameter.Value)"
fi

exec ./main "$@"
