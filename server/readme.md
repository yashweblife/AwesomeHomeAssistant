# Information

There is a `main.go` file which is the top level of the server, it should only contain api server level code

Files that follow `*Manager.go` format are used as controllers between the database and api server

`dbms.go` is where we do exclusive interaction with the db, dont call the db from outside this. The functions in here are used by `*Manager.go` files

`types.go` contains all the `structs` and `interfaces` used in the project, its a good habit to do this



#### Api

/auth/login - POST
/auth/register - POST
/auth/users - GET
/auth/user - GET
/auth/user - DELETE

/device/ - GET
/device/value - GET
/device/register - POST
/device/ - DELETE
/device/devices - GET