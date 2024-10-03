### Installation

#### Install Dependencies
```bash
go get github.com/gorilla/mux
go get github.com/go-sql-driver/mysql@latest
```

#### Create Database
```sql
CREATE DATABASE movies_db;
USE movies_db;
CREATE TABLE movies(id int NOT NULL AUTO_INCREMENT, title VARCHAR(255) NOT NULL,genre VARCHAR(255), rating DECIMAL(3, 1) CHECK (rating >= 1.0 AND rating <= 10.0), PRIMARY KEY (id));
INSERT INTO movies values(1, "Taare Zameen Par", "Drama", 9.4);
INSERT INTO movies values(2, "3 Idiots", "Comedy", 9.8);
SELECT * FROM movies;
```

#### DEFINE DB Details in const.go file
```go
package main

const DBUser = "[DATABASE USER]"
const DBName = "[DATABASE NAME]"
const DBPassword = "[DATABASE PASSWORD]"
const DBHost = "[DATABASE HOST]"
```

#### Running application
```bash
go build
./movies-app-crud
```
