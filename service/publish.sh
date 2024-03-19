# get current repository latest version
echo "current repository latest version:\n"
                     
git ls-remote --tags https://github.com/tcp-x/cd-rpc.git
# set latest version
Version="v0.0.4"

# cd $projDir
go mod tidy
git add go.mod go.sum multiplication_service.go user_service.go
git commit -am "set version $Version"
git tag $Version
git push origin $Version

