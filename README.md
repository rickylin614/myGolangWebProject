

# Golang Web專案

- 採前後分離
- 對應前端Vue : <https://github.com/ddalbert66/vue_order_bento>

# 使用技術



## 依賴包說明
### 1. gin-gonic

- 一種go的http web framework
- 使用gin來快速配置路由以及中間件(達成類似JAVA AOP)
- 在benchmark中，效能高於其他frame

### 2. uber-go/zap

- uber團隊的zap包用於日誌紀錄
- 高性能的日誌紀錄工具，效能高於其餘日誌紀錄
- 具有報錯顯示完整路徑並不會強制中斷程序的error錯誤日誌類型

### 3. viper

- viper為設定檔管理工具
- 支持JSON，TOML，YAML，HCL，envfile，Properties等等檔案
- 支持從環境變數中存取
- 支持熱讀取

### 4. gorm 

- ORM工具，方便從資料庫中存寫數據。
- 可快速的使用預設CRUD，使用方法輕鬆易懂
- 可自定義Logger，且默認Logger已有足夠的資料
- 支持事務回滾

### 5. go-redis

- 緩存存儲工具
- 使用於登入緩存、在線會員緩存、資料鎖

### 6. google/uuid

- 快速產生不會重複的唯一值
- 目前作用存取用戶識別(產生session id給予用戶)

### 6. gorilla/websocket

- 待補