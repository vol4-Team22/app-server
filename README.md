# Mikke-Server
> [!CAUTION]
> 実運用する際は鍵を新たに生成して、秘密鍵は絶対に公開しないでください。
## Start the server
Create a Docker image
```zsh
$ make build-local
```
Use Docker Compose to start each service.
```zsh
$ make up
```
Perform a migration to MySQL.
```zsh
$ make migrate
```

## Endpoint

[//]: # (- `/login`)

[//]: # (<details>)

[//]: # (<summary>input</summary>)

[//]: # ()
[//]: # (```json)

[//]: # ({)

[//]: # (  "title": "ここにタイトルが入る",)

[//]: # (  "comment": "ここにコメントが入る")

[//]: # (})

[//]: # (```)

[//]: # (</details>)

### `/post`
- 投稿する
<details>
<summary>input</summary>

```json
{
  "title": "ここにタイトルが入る",
  "comment": "ここにコメントが入る"
}
```
</details>

### `/questions`
- すべての投稿を取得する
<details>
<summary>output</summary>

```json
[
  {
    "post_id": 1,
    "title": "ここにタイトルが入る",
    "created_at": "2024-01-01 17:51:04.789463"
  },
  {
    "post_id": 2,
    "title": "2つ目の投稿のタイトル",
    "created_at": "2024-01-02 12:34:56.789010"
  },
  {
    "post_id": 3,
    "title": "最後の投稿のタイトル",
    "created_at": "2024-01-03 09:21:45.123456"
  }
]
```
</details>

### `/question?postid={id}`
- クエリパラメータ(`post_id`)で詳細取得
<details>
<summary>output</summary>

```json
{
  "user_id": "E724D1CE-396C-4C67-B8E7-495F9E842AEB",
  "post_id":  1,
  "title": "ここにタイトルが入る",
  "comment": "ここにコメントが入る",
  "created_at": "2024-01-01 17:51:04.789463",
  "updated_at": "2024-01-01 17:51:04.789463"
}
```
</details>