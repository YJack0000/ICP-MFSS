# ICP-MFS

2023 計算機概論與程式設計課期末用的檔案上傳服務，主要是希望有個我們自己的簡易的檔案上傳服務，並且可以對於不同題目過濾檔案名稱以及檔案大小。
此小專案使用 gin 作為後端框架，minio 作為資料儲存。前端使用簡易的 htmx + bootstrap 撰寫，整個應用包括跳出的通知都是 Server 端 Render 完成的。

## 怎麼跑起來

你可以看一下下面的介紹，然後稍微改一些地方之後就可以用 docker build 起來
```
$ docker build .
$ docker run -p <HostPort>:8080 \
-e MINIO_ENDPOINT=<MinioEndpoint> \
-e MINIO_ACCESS_KEY=<MinioAccessKey> \
-e MINIO_SECRET_KEY=<MinioSecretKey> \
<YourDockerImageName>
```

## 專案架構

簡單講一下，畢竟你看到這個專案你應該基本的都會了，主要有幾個部分：

### `handler/`

這個資料夾主要就是 API Endpoint 的邏輯，有兩個分別為 `/upload` 和 `/verify`，另外還有一個 `/write`。

* `/upload` 顧名思義就是讓同學上傳檔案的 API，裡面會去檢查一下檔案大小還有學號等等的，最後將檔案儲存在 minio。
* `/verify` 讓同學對於每一題可以去查詢看看目前有還有哪幾個檔案沒有上傳完成，哪幾個檔案上傳完成。
* `/write` 是讓我們自己用的，就是把這段時間在 middleware 層攔截到的 `IP` 跟 `學號` 記錄起來到 `static/ip_to_student_id.txt` 檔案，以方便我們檢查有沒有人在皮。

### `pkg/`

這裡面其實是方一些會跟其他東西互動的套件，但是這個專案只有 minio 所以裡面就只放 `minio_client`。那 `client.go` 裡面有使用 `sync.Once` 來確保他只會做一次連線，不會每次 request 都建立新的連線。
但這個地方讀 config 的方法很懶散，有興趣可以來改改。

### `static/`

這裡面就是放前端的網頁需要用的東西，比較需要注意的是 `components/` ，這是之後在 `utils/` 裡面會用到的。其他就是 htmx + bootstrap，把網頁用得票票釀釀並且可以跟後端互動。

### `utils/`

* `submission_utils.go` 比較需要注意一下，裡面的 `GetVaildSubmissionFileNames` 是用來設定說每個不同的上傳區裡面應該要有什麼檔案。沒錯就是那麼土炮，認識我的可能覺得啊幹你的 code 怎麼就不 OCP 了，那我也只能說這樣挺方便快速的，不服來發 PR。
* 另外 `*_response.go` 系列就是拿 `static/components` 裡面的模板來渲染出前端要顯示的內容，你去看一下就會懂了。
