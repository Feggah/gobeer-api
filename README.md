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
