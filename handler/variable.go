package handler

const (
	// VERSION version
	VERSION                   = "0.1 alpha"

	//ADD operation in handler
	ADD                       = iota
	DELETE
	READ
	UNIQID_SIZE               = 32

	API_CONTENT_HEADER        = "application/json;charset=utf-8"
	ETCD_CONTENT_HEADER       = "application/x-www-form-urlencoded"

	ADD_METHOD                = "PUT"
	DELETE_METHOD             = "DELETE"
	CONTENT_HEADER            = "Content-Type"

	MACEDON_ADD_LOC           = "/add"
	MACEDON_DELETE_LOC        = "/delete"
	MACEDON_READ_LOC          = "/read"
	MACEDON_SCAN_LOC          = "/scan"

	UROUTER_DEFAULT_LOC       = "/urouter"
	DEFAULT_REDIS_PORT        = "6379"

	UROUTER_DEFAULT_TTL       = 60

	UROUTER_DEFAULT_TIMEOUT   = 5

	MACEDON_TOKEN             = "macedon_token"
)
