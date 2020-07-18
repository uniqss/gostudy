module servers

go 1.14

require (
	github.com/cihub/seelog v0.0.0-20170130134532-f561c5e57575
	github.com/sirupsen/logrus v1.6.0 // indirect
	library_wrappers v0.0.0-00010101000000-000000000000
)

replace library_wrappers => ../library_wrappers
