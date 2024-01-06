// Package httpctype list some common http content-type
package httpctype

const Key = "Content-Type"

// HTTP Content-Type
const (
	CSS      = "text/css; charset=utf-8"
	HTML     = "text/html; charset=utf-8"
	Text     = "text/plain; charset=utf-8"
	Plain    = Text
	XML2     = "text/xml; charset=utf-8"
	XML      = "application/xml; charset=utf-8"
	YML      = "application/x-yaml; charset=utf-8"
	YAML     = "application/x-yaml; charset=utf-8"
	JSON     = "application/json; charset=utf-8"
	JSONP    = "application/javascript; charset=utf-8"
	JS       = JSONP
	JS2      = "text/javascript; charset=utf-8"
	MSGPACK  = "application/x-msgpack; charset=utf-8"
	MSGPACK2 = "application/msgpack; charset=utf-8"
	PROTOBUF = "application/x-protobuf"
	Form     = "application/x-www-form-urlencoded"
	FormData = "multipart/form-data"
	Binary   = "application/octet-stream"
)

// Content-Type MIME
const (
	MimeHTML          = "text/html"
	MimeText          = "text/plain"
	MimePlain         = MimeText
	MimeJSON          = "application/json"
	MimeXML           = "application/xml"
	MimeXML2          = "text/xml"
	MimeYAML          = "application/x-yaml"
	MimePOSTForm      = "application/x-www-form-urlencoded"
	MimeMultiDataForm = "multipart/form-data"
	MimePROTOBUF      = "application/x-protobuf"
	MimeMSGPACK       = "application/x-msgpack"
	MimeMSGPACK2      = "application/msgpack"
)
