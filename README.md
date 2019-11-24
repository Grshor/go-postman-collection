# go-postman-collection

[![Test](https://github.com/rbretecher/go-postman-collection/workflows/test/badge.svg)](https://github.com/rbretecher/go-postman-collection/actions?query=workflow=test)
[![Go Report Card](https://goreportcard.com/badge/github.com/rbretecher/go-postman-collection)](https://goreportcard.com/report/github.com/rbretecher/go-postman-collection)

Go module to work with Postman Collections.

Development is under progress and for now it only supports partially Postman Collection format v2.1.0

### Current support

#### Postman Collection Schema

| Schema Version | Supported |
| -------------- | --------- |
| 1.0            | No        |
| 2.0 < 3.0      | No        |
| 2.1 >= 3.0     | Partial   |

#### Postman Objects

|  Object            | Supported |
| ------------------ | --------- |
| Collection         | Yes       |
| ItemGroup (Folder) | Partial   |
| Item               | Partial   |
| Request            | Partial   |
| Response           | No        |
| Event              | No        |
| Variable           | No        |
| Auth               | Yes       |

### Useful resources

https://schema.getpostman.com/json/collection/latest/docs/index.html
