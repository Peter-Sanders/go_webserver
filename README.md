# go_webserver
Golang + HTMX = Profit?


## init

need to create a file 'env.ini' in the base dir

it should contain this information, adding what you please

dbname=
adminfname=
adminlname=
adminusername=
adminphone=
adminemail=
adminpassword=

to start:
    make run

to build the db:
    make migrate-up

to wipe the db (THIS WILL NUKE ALL TABLES AND CONTENT SO BE CAREFUL!!!!):
    make migrate-down
