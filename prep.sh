git init .
git add -A .
git commit -m HelloWorld

go get github.com/tools/godep
godep save
git add -A .
git commit -m Deps

godep go build
./go-on-heroku
