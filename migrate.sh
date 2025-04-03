#!/usr/bin/env bash

# = = = = = = = =
#   Utilities
# = = = = = = = =

form_postgres_url() {
  # [capturable]

  source .env
  [ $? -ne 0 ] && return 1

  printf "postgresql://%s:%s@%s:%s/%s?sslmode=%s" \
         "$DB_USER" "$DB_PASSWORD" "$DB_HOST" "$DB_PORT" "$DB_NAME" "$DB_SSL_MODE"
}

get_migrations_path() {
  # [capturable]

  printf "./platform/migrations"
}


# = = = = = = = = = = = =
#    Command Handlers
# = = = = = = = = = = = =

up_command() {
  # [control]

  local path
  path="$(get_migrations_path)"
  [ $? -ne 0 ] && return 1

  local url
  url="$(form_postgres_url)"
  [ $? -ne 0 ] && return 1

  migrate -path "$path" -database "$url" -verbose up
}

down_command() {
  # [control]

  local path
  path="$(get_migrations_path)"
  [ $? -ne 0 ] && return 1

  local url
  url="$(form_postgres_url)"
  [ $? -ne 0 ] && return 1

  migrate -path "$path" -database "$url" -verbose down
}

create_command() {
  # [control]

  if [[ $# -eq 0 ]]; then
    printf "Usage: migrate.sh create <migration_name>\n" >&2
    return 1
  fi

  local name="$1"
  shift

  local path
  path="$(get_migrations_path)"
  [ $? -ne 0 ] && return 1

  migrate create -ext sql -dir "$path" -seq "$name"
}

url_command() {
  # [control]

  local result
  result="$(form_postgres_url)"
  [ $? -ne 0 ] && return 1

  printf %s "$result"
}


# = = = = = = = = = = =
#    Routing Section
# = = = = = = = = = = =

print_help_message() {
  local -a message=(
    'Usage: migrate.sh <command>'
    ''
    'A tool for automating the golang-migrate command.'
    ''
    'Commands:'
    '    up       Apply all up migrations.'
    ''
    '    down     Apply all down migrations.'
    ''
    '    url      Do not apply migrations just return'
    '             the database connection url.'
    ''
    '    create   Create a new sequential migration.'
    '             Requires a single argument that represents'
    '             the name of new migration files.'
    ''
    '    help     Print this message.'
  )

  for line in "${message[@]}"; do
    echo "$line"
  done
}

if [[ $# -eq 0 ]]; then
  print_help_message >&2
  exit 1
fi


declare selected_command
selected_command="$1"
shift

case "$selected_command" in
  'up') up_command;;
  'down') down_command;;
  'create') create_command "$@";;
  'url') url_command;;
  'help') print_help_message; exit 0;;
  *)
    printf "Command '%s' not recognized. See 'migrate.sh help' for more info.\n" "${selected_command}" >&2
    exit 1
    ;;
esac