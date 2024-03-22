# This script is necessitated by vscode inability to upgrade modules due to cache issues
go clean -cache
go clean -modcache
# rm -rf ~/.config/Code/Cache/*
# rm -rf ~/.config/Code/CachedData/*

# clean mod files: remove get github.com/tcp-x modules
bash -c 'sh clean-mod-files.sh "/media/emp-06/disk-02/projects/test-projects/cd-user/server/go.sum" "github.com/tcp-x/"'
bash -c 'sh clean-mod-files.sh "/media/emp-06/disk-02/projects/test-projects/cd-user/server/go.mod" "github.com/tcp-x/"'
bash -c 'sh clean-mod-files.sh "/media/emp-06/disk-02/projects/test-projects/cd-user/client/go.sum" "github.com/tcp-x/"'
bash -c 'sh clean-mod-files.sh "/media/emp-06/disk-02/projects/test-projects/cd-user/client/go.mod" "github.com/tcp-x/"'

go mod tidy

# get the latests
cd /media/emp-06/disk-02/projects/test-projects/cd-user/server
go get github.com/tcp-x/cd-core/sys/base
go get github.com/tcp-x/cd-core/sys/user
go get github.com/tcp-x/cd-rpc/service

cd /media/emp-06/disk-02/projects/test-projects/cd-user/client
go get github.com/tcp-x/cd-core/sys/base
go get github.com/tcp-x/cd-core/sys/user
go get github.com/tcp-x/cd-rpc/service