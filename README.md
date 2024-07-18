<!--
 Copyright 2024 ariefsetyonugroho
 
 Licensed under the Apache License, Version 2.0 (the "License");
 you may not use this file except in compliance with the License.
 You may obtain a copy of the License at
 
     https://www.apache.org/licenses/LICENSE-2.0
 
 Unless required by applicable law or agreed to in writing, software
 distributed under the License is distributed on an "AS IS" BASIS,
 WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 See the License for the specific language governing permissions and
 limitations under the License.
-->
### Create & install
- Membuat project go
```
go mod init "nameproject"
```

- Installasi gin (web framework)
```
go get -u github.com/gin-gonic/gin
```

- Installasi GORM - Golang Object-Realtional Mapping (library database)
```
go get -u gorm.io/gorm gorm.io/diver/mysql
```

### Run Project
```
go run main.go
```
Open : [Localhost](http://localhost:3000)

#### Router API
- get all data (GET)
```
http://localhost:3000/api/post
```

- post data (POST)
```
http://localhost:3000/api/post

body {
    "title": "",
    "content": "" 
}
```

- get data by id (GET)
```
http://localhost:3000/api/post/'id'
```

### Referensi
[Santri Koding](https://santrikoding.com/tutorial-restful-api-golang-1-membuat-project-golang
)