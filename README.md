# DbMap
Tiny package to auto map struct fields to database columns and vice versa without ORM's.

[![Go Report Card](https://goreportcard.com/badge/github.com/AlexanderMatveev/dbmap)](https://goreportcard.com/report/github.com/AlexanderMatveev/dbmap)

## Usage

```go
import (
    "github.com/AlexanderMatveev/dbmap"
    "github.com/Masterminds/squirrel"
)

type Entity struct {
    Id            int       `db:"id"`
    Name          string    `db:"name"`
    DateTimeField time.Time `db:"created_at"`
    FieldNotIdDb  int
    NullableField *int
}

var structColumns = dbmap.Columns(reflect.TypeOf(Entity{}))

func main(){
    var e Entity
    // ...
    squirrel.Select(structColumns...)
    // ...
    rows.Scan(dbmap.Scan(&e)...)
    // ...
    squirrel.Insert("table").Columns(structColumns...).Values(dbmap.Values(e)...)
    // ...
}

```
