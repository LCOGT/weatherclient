# Las Cumbres Observatory Weather Web Interface

This is the source code for the Las Cumbres Observatory Weather Web Interface,
which is deployed at: <https://weather.lco.global/>.

## Production Deployment

This project is built automatically by the [LCO Jenkins Server](http://jenkins.lco.gtn/).
Please see the [Jenkinsfile](Jenkinsfile) for further details.

This project is deployed to the LCO Kubernetes Cluster. Please see the
[LCO Helm Charts Repository](https://github.com/LCOGT/helm-charts) for further details.

## Getting Started

These instructions will get you a copy of the project up and running on your
local machine for development and testing purposes.

### Weather API Server

Please see the [Weather API Server Github Repository](https://github.com/LCOGT/weatherapi)
for details about how to use the Weather Server.

### Weather Client

This will build a Docker image named `weatherclient`, containing the client code:

```
$ docker build --pull -t my-weather-client:latest .
```

This Docker image can be run by executing this command:

```
$ docker run --rm -it -p 8080:8080 my-weather-client:latest
```

### Local development

``` bash
# install dependencies
npm install

# serve with hot reload at localhost:8080
npm run dev

# build for production with minification
npm run build
```

For detailed explanation on how things work, consult the [docs for vue-loader](https://vue-loader-v14.vuejs.org/en/).

## License

This project is licensed under the GNU GPL v3 License - see the
[LICENSE](LICENSE) file for details.
