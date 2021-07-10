# Gobeer

**Gobeer** is an API developed for studies purposes. If you want to create your own API, follow along with [this tutorial series](https://youtu.be/MNE_grboFPM) that explains everything about this API.

## Setting up your database

This API uses a `sqlite3` database because of its simplicity. Please, follow the steps below to create your databases.

Firstly, create the main `sqlite3` database:
```
$ sqlite3 data/beer.db
sqlite> CREATE TABLE beer(id INTEGER PRIMARY KEY AUTOINCREMENT, name text NOT NULL, type integer NOT NULL, style integer not null);
sqlite> .quit
```

Then, create the database that will be used on tests:
```
$ sqlite3 data/beer_test.db
sqlite> CREATE TABLE beer(id INTEGER PRIMARY KEY AUTOINCREMENT, name text NOT NULL, type integer NOT NULL, style integer not null);
sqlite> .quit
```

## Make commands

This repository offers a set of `make` commands to help the developing process. 

1. `make test` : runs all unit tests developed in the repository. It will show the tests coverage in the terminal.
2. `make coverage` : Same as `make test`, but a pop-up will appear where you can check which files and lines the test cases are covering.
3. `make run` : runs the application.
4. `make create` : makes a `POST` request to the API to create one Beer. It expects those mandatory parameters: `name`, `type` and `style`.
> Example:
> ```
> make create name=Heineken type=2 style=6
> ```

5. `make list` : makes a `GET` request to the API to list all Beers.
6. `make get` : makes a `GET` request to the API to get a specific Beer. It expects the `id` parameter.
> Example:
> ```
> make get id=1
> ```

7. `make update` : makes a `PUT` request to the API to update a Beer. It expects those mandatory parameters: `id`, `name`, `type` and `style`.
> Example:
> ```
> make update id=1 name=Beck's type=2 style=12
> ```

8. `make delete` : makes a `DELETE` request to the API to delete a Beer. It expects the `id` parameter.
> Example:
> ```
> make delete id=1
> ```

9. `make container` : builds and runs the docker container.
