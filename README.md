webdavfs
========

[![Go Reference](https://pkg.go.dev/badge/github.com/shurcooL/webdavfs.svg)](https://pkg.go.dev/github.com/shurcooL/webdavfs)

Collection of Go packages for working with the [`webdav.FileSystem`](https://godoc.org/golang.org/x/net/webdav#FileSystem) interface.

Installation
------------

```sh
go get github.com/shurcooL/webdavfs
```

Directories
-----------

| Path                                                                 | Synopsis                                                                     |
|----------------------------------------------------------------------|------------------------------------------------------------------------------|
| [vfsutil](https://pkg.go.dev/github.com/shurcooL/webdavfs/vfsutil)   | Package vfsutil implements some I/O utility functions for webdav.FileSystem. |
| [webdavfs](https://pkg.go.dev/github.com/shurcooL/webdavfs/webdavfs) | Package webdavfs implements webdav.FileSystem using an http.FileSystem.      |

License
-------

-	[MIT License](LICENSE)
