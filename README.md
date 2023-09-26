# DbMap
Tiny package to auto map struct fields to database columns and vice versa without ORM's.

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

var structColumns = dbmap.Columns(reflect.TypeOf(event.Event{}))

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