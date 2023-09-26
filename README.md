# DbMap
Tiny package to auto map struct fields to database columns and vice versa without ORM's.

[![Workflow](https://github.com/AlexanderMatveev/dbmap/actions/workflows/go.yml/badge.svg)](https://github.com/AlexanderMatveev/dbmap/actions)
[![Go Report Card](https://goreportcard.com/badge/github.com/AlexanderMatveev/dbmap)](https://goreportcard.com/report/github.com/AlexanderMatveev/dbmap)

## Usage

```go
import (
    "github.com/AlexanderMatveev/dbmap"
    "github.com/Masterminds/squirrel"
)

type Entity struct {
    Id        int       `db:"id"`
    Name      string    `db:"name"`
    CreatedAt time.Time `db:"created_at"`
}

var structColumns = dbmap.Columns(reflect.TypeOf(Entity{}))

func main(){
    // ...
    squirrel.Select(structColumns...)
    // ...
    rows.Scan(dbmap.Scan(&s)...)
    // ...
    squirrel.Insert("table").Columns(structColumns...)
    // ...
}

```