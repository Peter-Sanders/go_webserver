#!/bin/bash 

direction=$1
# query=$(cat "sql/migrate-${direction}.sql")
fname="sql/migrate-${direction}.sql"
db="db/$(awk -F "=" '/dbname/ {print $2}' env.ini)"

sqlite3 $db ".read ${fname}"


if [ $direction == "up" ]; then
  adminfname=".param set :fname $(awk -F "=" '/adminfname/ {print $2}' env.ini)"
  adminlname=".param set :lname $(awk -F "=" '/adminlname/ {print $2}' env.ini)"
  adminusername=".param set :username $(awk -F "=" '/adminusername/ {print $2}' env.ini)"
  adminphone=".param set :phone $(awk -F "=" '/adminphone/ {print $2}' env.ini)"
  adminemail=".param set :email $(awk -F "=" '/adminemail/ {print $2}' env.ini)"
  adminpassword=".param set :password $(echo -n $(awk -F "=" '/adminpassword/ {print $2}' env.ini) | sha256sum)"

  sqlite3 $db  ".param init" "${adminfname}" "${adminlname}" "${adminusername}" "${adminphone}" "${adminemail}" "${adminpassword::-3}" ".read sql/admin/insert_admin_user.sql"

fi

exit 0


