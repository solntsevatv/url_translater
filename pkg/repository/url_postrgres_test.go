package repository

import (
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jmoiron/sqlx"
	"github.com/solntsevatv/url_translater/internal/url_translater"
	"github.com/stretchr/testify/assert"
)

func TestUrlPostgres_CreateShortURL(t *testing.T) {
	mockDB, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer mockDB.Close()
	db := sqlx.NewDb(mockDB, "sqlmock")

	r := NewUrlPostgres(db)

	tests := []struct {
		name    string
		mock    func()
		input   url_translater.URL
		want    string
		wantErr bool
	}{
		{
			name: "Ok",
			mock: func() {
				rows := sqlmock.NewRows([]string{"id"}).AddRow(1)
				mock.ExpectQuery("INSERT INTO links").
					WithArgs("some_long_url", "A").WillReturnRows(rows)
			},
			input: url_translater.URL{
				Id:       1,
				LongUrl:  "some_long_url",
				ShortURL: "A",
			},
			want: "A",
		},
		{
			name: "Empty Fields",
			mock: func() {
				rows := sqlmock.NewRows([]string{"id"})
				mock.ExpectQuery("INSERT INTO links").
					WithArgs("some_long_url", "").WillReturnRows(rows)
			},
			input: url_translater.URL{
				Id:       1,
				LongUrl:  "some_long_url",
				ShortURL: "",
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mock()

			got, err := r.CreateShortURL(tt.input)
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.want, got)
			}
			assert.NoError(t, mock.ExpectationsWereMet())
		})
	}
}

func TestAuthPostgres_GetLongURL(t *testing.T) {
	mockDB, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer mockDB.Close()
	db := sqlx.NewDb(mockDB, "sqlmock")

	r := NewUrlPostgres(db)

	type args struct {
		short_url string
	}

	tests := []struct {
		name    string
		mock    func()
		input   url_translater.ShortURL
		want    string
		wantErr bool
	}{
		{
			name: "Ok",
			mock: func() {
				rows := sqlmock.NewRows([]string{"id", "long_url"}).AddRow(64, "AB")
				mock.ExpectQuery("SELECT (.+) FROM links").WithArgs("AB").WillReturnRows(rows)
			},
			input: url_translater.ShortURL{
				Id:      64,
				LinkUrl: "AB",
			},
			want:    "AB",
			wantErr: false,
		},
		{
			name: "Not Found",
			mock: func() {
				rows := sqlmock.NewRows([]string{"id", "long_url"})
				mock.ExpectQuery("SELECT (.+) FROM links").
					WithArgs("not found").WillReturnRows(rows)
			},

			input: url_translater.ShortURL{
				Id:      1,
				LinkUrl: "not found",
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mock()

			got, err := r.GetLongURL(tt.input)
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.want, got)
			}
			assert.NoError(t, mock.ExpectationsWereMet())
		})
	}
}
