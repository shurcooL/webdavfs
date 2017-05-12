webdavfs
========

[![Build Status](https://travis-ci.org/shurcooL/webdavfs.svg?branch=master)](https://travis-ci.org/shurcooL/webdavfs) [![GoDoc](https://godoc.org/github.com/shurcooL/webdavfs?status.svg)](https://godoc.org/github.com/shurcooL/webdavfs)

Collection of Go packages for working with the [`webdav.FileSystem`](https://godoc.org/golang.org/x/net/webdav#FileSystem) interface.

Installation
------------

```bash
go get -u github.com/shurcooL/webdavfs/...
```

Directories
-----------

| Path                                                                | Synopsis                                                                     |
|---------------------------------------------------------------------|------------------------------------------------------------------------------|
| [vfsutil](https://godoc.org/github.com/shurcooL/webdavfs/vfsutil)   | Package vfsutil implements some I/O utility functions for webdav.FileSystem. |
| [webdavfs](https://godoc.org/github.com/shurcooL/webdavfs/webdavfs) | Package webdavfs implements webdav.FileSystem using an http.FileSystem.      |

License
-------

-	[MIT License](https://opensource.org/licenses/mit-license.php)
