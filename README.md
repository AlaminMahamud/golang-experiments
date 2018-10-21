# Golang Experiments

Codebase for my experiments with golang

## Getting Started

These instructions will get you a copy of the project up and running on your local machine for development and testing purposes. See deployment for notes on how to deploy the project on a live system.

### Prerequisites
golang
``` go
go version
```


### Install
``` shell
go get -u github.com/alaminmahamud/golang_experiments
```

## Running the tests
```bash
go test
```

## Deployment

Add additional notes about how to deploy this on a live system

## Structure

``` bash
Go Directories
===============
/cmd

# Main applications for this project.

# The directory name for each application should match the name of the executable you want to have (e.g., /cmd/myapp).

# Don't put a lot of code in the application directory. If you think the code can be imported and used in other projects, then it should live in the /pkg directory. If the code is not reusable or if you don't want others to reuse it, put that code in the /internal directory. You'll be surprised what others will do, so be explicit about your intentions!

# It's common to have a small main function that imports and invokes the code from the /internal and /pkg directories and nothing else.


/internal
/pkg
/vendor

Service Level Directories
=========================
/api

Web Application Directories
===========================
/web

Common Application Directories
==============================
/configs
/init
/scripts
/build
/deployments
/test


Other Directories
=================
/docs
/tools
/examples
/third-party
/githooks
/assets
/website

Directories you shouldn't have
==============================
/src

```

## Built With

* [unittest](https://docs.python.org/3/library/unittest.html) - builtins `unittest` framework is used.

## Contributing

Please read [CONTRIBUTING.md](CONTRIBUTING.md) for details on our code of conduct, and the process for submitting pull requests to us.

## Versioning

We use [SemVer](http://semver.org/) for versioning. For the versions available, see the [tags on this repository](https://github.com/your/project/tags). 

## Authors

* **Alamin Mahamud** - *Initial work* - [alamin.rocks](https://alamin-rocks.herokuapp.com)

See also the list of [contributors](https://github.com/your/project/contributors) who participated in this project.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details

## Acknowledgments

* Hat tip to anyone whose code was used
* Inspiration - vinta and other awesome repo guys.
