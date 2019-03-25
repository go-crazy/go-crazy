#rm -rf vendor/github.com/go-crazy/*
#ln -s $PWD/../elastic vendor/github.com/go-crazy/elastic
#ln -s $PWD/../cache vendor/github.com/go-crazy/cache
# ln -s $PWD/../authentication vendor/github.com/go-crazy/authentication
go build server.go  &&  ./server