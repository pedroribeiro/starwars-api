# starwars-api

## About the project

starwars-api is an api to manage Star wars Planets.

## API docs

GET ROUTES

- /plantes (list all planets)
- /plantes?name=_planet name_ (get one planets by name)
- /plantes?id=_planet id_ (get one planets by name)

POST ROUTES

- /plantes (add planet)

```
body schema exemple:
{
	"name": "Earth",
	"climate": "Mixed",
	"terrain": ""
}
```

DELETE ROUTES

- /plantes?id=_planet id_ (delete planet by ID)

## Install

To build startwars-api You need:

- Go [version 1.16 or greater](https://golang.org/doc/install).
- MongoDB [version 5.0.6 or greater](https://docs.mongodb.com/manual/installation/)

Clone this repository and build using `make build`:

    $ mkdir -p $GOPATH/src/github.com/pedroribeiro
    $ cd $GOPATH/src/github.com/pedroribeiro
    $ git clone https://github.com/pedroribeiro/starwars-api.git
    $ cd starwars-api
    $ create a .env file on project root folder and add your DB connection string and port:

	exemple:
	     DB_URL=mongodb://127.0.0.1:27017
	     PORT=:3030

    $ make build
    $ ./starwars-api

## Contributing

* Pedro Ribeiro
	* https://github.com/pedroribeiro
	* https://www.linkedin.com/in/ribeiro-m-pedro/


## License

Apache License 2.0, see [LICENSE](https://github.com/prometheus/prometheus/blob/main/LICENSE).
