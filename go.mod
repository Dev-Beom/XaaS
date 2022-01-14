module github.com/dev-beom/xaas

go 1.17

replace github.com/dev-beom/xaas/apiserver => ./apiserver

replace github.com/dev-beom/xaas/controlmanager => ./controlmanager

require (
	github.com/dev-beom/xaas/apiserver v0.0.0-00010101000000-000000000000
	github.com/dev-beom/xaas/controlmanager v0.0.0-00010101000000-000000000000
)

require (
	github.com/Microsoft/go-winio v0.4.16 // indirect
	github.com/dgrijalva/jwt-go v3.2.0+incompatible // indirect
	github.com/go-playground/locales v0.14.0 // indirect
	github.com/go-playground/universal-translator v0.18.0 // indirect
	github.com/go-playground/validator v9.31.0+incompatible // indirect
	github.com/james-barrow/golang-ipc v0.0.0-20210227130457-95e7cc81f5e2 // indirect
	github.com/joho/godotenv v1.4.0 // indirect
	github.com/labstack/echo v3.3.10+incompatible // indirect
	github.com/labstack/gommon v0.3.1 // indirect
	github.com/leodido/go-urn v1.2.1 // indirect
	github.com/mattn/go-colorable v0.1.12 // indirect
	github.com/mattn/go-isatty v0.0.14 // indirect
	github.com/valyala/bytebufferpool v1.0.0 // indirect
	github.com/valyala/fasttemplate v1.2.1 // indirect
	golang.org/x/crypto v0.0.0-20211215153901-e495a2d5b3d3 // indirect
	golang.org/x/net v0.0.0-20211216030914-fe4d6282115f // indirect
	golang.org/x/sys v0.0.0-20211216021012-1d35b9e2eb4e // indirect
	golang.org/x/text v0.3.7 // indirect
)
