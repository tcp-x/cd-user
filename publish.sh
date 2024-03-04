# get current repository latest version
echo "current repository latest version:\n"
git ls-remote --tags https://github.com/tcp-x/cd-user.git

# cd-cli plugin compile 
go build -buildmode=plugin -o ./plugin/User.so ./user/user.go
go build -buildmode=plugin -o ./plugin/Session.so ./session/session.go

# set latest version
Version="v0.0.7"
go mod tidy
git submodule update --remote
git add go.mod go.sum user/user.go plugin/User.so session/session.go plugin/Session.so
git add -A
git commit -a -m "set version $Version"
git tag $Version
git push origin $Version






