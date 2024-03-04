# get current repository latest version
echo "current repository latest version:\n"
git ls-remote --tags https://github.com/tcp-x/cd-user.git
# cd-cli plugin compile 
go build -buildmode=plugin -o User.so
# set latest version
Version="v0.0.4"
go mod tidy
git submodule update --remote
git add go.mod go.sum user.go User.so
git add -A
git commit -a -m "set version $Version"
git tag $Version
git push origin $Version



