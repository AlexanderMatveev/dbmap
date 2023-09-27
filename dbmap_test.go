package dbmap

import (
	"reflect"
	"testing"
	"time"
)

type Entity struct {
	Id            int       `db:"id"`
	NullableField *int      `db:"nullable_field"`
	Name          string    `db:"name"`
	CreatedAt     time.Time `db:"created_at"`
	FieldNotIdDb  int
}

func TestColumns(t *testing.T) {
	type args struct {
		t reflect.Type
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "Correct",
			args: args{
				t: reflect.TypeOf(Entity{}),
			},
			want: []string{
				"id",
				"nullable_field",
				"name",
				"created_at",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Columns(tt.args.t); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Columns() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestScan(t *testing.T) {
	type args struct {
		s any
	}
	e := Entity{
		Id:            1,
		NullableField: nil,
		Name:          "Test",
		CreatedAt:     time.Now(),
	}
	tests := []struct {
		name string
		args args
		want []any
	}{
		{
			name: "Correct",
			args: args{
				s: &e,
			},
			want: []any{
				&e.Id,
				&e.NullableField,
				&e.Name,
				&e.CreatedAt,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Scan(tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Scan() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestValues(t *testing.T) {
	type args struct {
		s any
	}
	e := Entity{
		Id:        1,
		Name:      "Test",
		CreatedAt: time.Now(),
	}
	tests := []struct {
		name string
		args args
		want []any
	}{
		{
			name: "Correct",
			args: args{
				s: e,
			},
			want: []any{
				e.Id,
				e.NullableField,
				e.Name,
				e.CreatedAt,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Values(tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Values() = %v, want %v", got, tt.want)
			}
		})
	}
}
