# Storage System

Fullstack application for managing a storage/inventory system.

> Built with local development in mind. May not be suitable for online production use.

## Features

- REST API
- MongoDB

## Running MongoDB service

Before running the API, a _MongoDB_ service must be running on port _27017_ (default).

**Systemd**:

```shell
# systemctl start mongodb.service
```

**_MongoDB_ daemon (using config file)**:

```shell
# mongod --config /etc/mongodb.conf
```

**_MongoDB_ daemon (using flag)**:

```shell
# mongod --port 27017
```

## Running the API

**WARNING:** If the _MongoDB_ service is not running, the API will **NOT** work.

If you are on Unix, you can run the API with:

```shell
$ cd api/
$ make build run
```

## Running Web UI

**TODO**

## License

Currently using MIT license. [View here](LICENSE)
