# go-crazy 
A Golang Framework For Web Artisans

#install
1.Create your project folder and cd inside

    $ mkdir -p $GOPATH/src/github.com/myusername/project && cd "$_"

2.Download src:

    git clone https://github.com/QylinFly/go-crazy.git ./

3.Vendor init your project

    govendor init

4.Run your project

    $ go run main.go


### vendor 

##### config file
     https://github.com/jinzhu/configor

##### ORM:gorm 
    https://github.com/jinzhu/gorm

##### log
    https://github.com/uber-go/zap


### manage
#### graceful-shutdown
    http://localhost:8080/down

### elastic 6.1.1
#### github.com/olivere/elastic
    https://github.com/olivere/elastic


### shell
    govendor fetch github.com/gorilla/websocket@v1.2.0
    go build index.go  &&  ./index

    go get -u golang.org/x/net