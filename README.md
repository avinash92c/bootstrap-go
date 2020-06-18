# bootstrap-go
go application bootstrap library to quickstart your application development while the library sets up everything your application needs to be 12factor compliant
This is a wip library with more features coming.

Current features in v1
- argon2 password encryption
- jwt support with cookie tokens
- MySQL and postgres connection pooling
- file source based config store with change watchers
- remote config store support with change watchers
   - supported stores etcd,consulkv,firebase
- Logrus logging support
- configurable route builder for rest endpoints.
   - currently only secure and insecure route configuration available

Pending proposed features
- opentracing support
- nosql support
- http based config store with change watcher
   - pluggable with spring cloud config, custom http endpoint  for starters
   - pluggable custom change watchers
- additional SQL databases connection pooling support
- and more
