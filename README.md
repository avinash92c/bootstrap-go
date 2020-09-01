# bootstrap-go

[![made-with-Go](https://img.shields.io/badge/Made%20with-Go-1f425f.svg)](http://golang.org)
[![GoDoc](https://godoc.org/github.com/sirupsen/logrus?status.svg)](https://pkg.go.dev/github.com/avinash92c/bootstrap-go/?tab=doc)
[![License](https://img.shields.io/badge/License-Apache%202.0-blue.svg)](https://github.com/avinash92c/bootstrap-go/blob/master/LICENSE)
[![Documentation](https://readthedocs.org/projects/ansicolortags/badge/?version=latest)](https://github.com/avinash92c/bootstrap-go/blob/master/docs/readme.md)

- [USAGE-DOCS](https://github.com/avinash92c/bootstrap-go/blob/master/docs/readme.md)

go application bootstrap library to quickstart your application development while the library sets up everything your application needs to be 12factor compliant
This is a wip library with more features coming.

Current features in v1

- argon2 password encryption
- jwt support with cookie tokens
- MySQL and postgres connection pooling
- file source based config store with change watchers
- remote config store support with change watchers
  - supported stores etcd,consulkv,firebase
- Logrus logging support with following hooks
  - File
  - Graylog
  - Logstash
- configurable route builder for rest endpoints.
  - currently only secure and insecure route configuration available
- Profiling With PProf Http Endpoints. Refer Docs

Pending proposed features

- opentracing support
- caching local + remote
- opentelemetry support
- nosql support
- http based config store with change watcher
  - pluggable with spring cloud config, custom http endpoint for starters
  - pluggable custom change watchers
- additional SQL databases connection pooling support
- and more

Known Issues

- Rolling File Logging Hook Not Writing Data to File
