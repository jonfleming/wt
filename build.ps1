# Web-to-text wt.exe for Windows
# jonfleming@hotmail.com
$version = git describe --tags --always
go build -ldflags "-X main.version=$version" -o wt.exe