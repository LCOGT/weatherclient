Las Cumbres Observatory Weather Web Interface
=============================================

This is the source code for the Las Cumbres Observatory Weather Web Interface,
which is deployed at: <https://weather.lco.global/>.

Getting Started
---------------

These instructions will get you a copy of the project up and running on your
local machine for development and testing purposes.

### Weather Server

This will build a Docker image named `weatherserver`, containing the server code:

    $ docker build --pull -t weatherserver server

This Docker image can be run by executing this command:

    $ docker run -d weatherserver

### Weather Client

This will build a Docker image named `weatherclient`, containing the client code:

    $ docker build --pull -t weatherclient client

This Docker image can be run by executing this command:

    $ docker run -d -p 80:80 weatherclient

License
-------

This project is licensed under the GNU GPL v3 License - see the
[LICENSE](LICENSE) file for details.
