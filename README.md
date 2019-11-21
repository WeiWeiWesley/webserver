Webserver
===

## 簡介
這是一個基於[ Gin Web Framework ](https://github.com/gin-gonic/gin)建立的 web server架構，提供簡易的 http 與 websocket 路由。可透過 docker 容器啟動，若有動態擴展與自動部署需求可參照[ cloudbuild.yaml ](https://github.com/WeiWeiWesley/webserver/blob/master/cloudbuild.yaml)進行部署。

預留資料庫、RPC連線、快取、設定檔介面供伺服器回應使用，使用方法可參照 [ example branch ](https://github.com/WeiWeiWesley/webserver/tree/example)

## Related Projects
[ Gin Web Framework ](https://github.com/gin-gonic/gin) <br>
[ Gorilla WebSocket ](https://github.com/gorilla/websocket) <br>
[ Logrus ](https://github.com/sirupsen/logrus) <br>
[ Simple Redis ](https://github.com/WeiWeiWesley/simple_redis)<br>
[ Redigo ](https://github.com/gomodule/redigo) <br>
[ Simple ORM ](https://github.com/WeiWeiWesley/simple_orm)<br>
[ GORM ](https://github.com/jinzhu/gorm) <br>
[ BurntSushi/toml ](https://github.com/BurntSushi/toml) <br>
