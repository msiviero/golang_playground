package user

import (
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/assert"
)

func TestUserGetHandler(t *testing.T) {

	mockDb, sqlMock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer mockDb.Close()

	underTest := NewUserService(sqlx.NewDb(mockDb, "mocksql"))

	sql := regexp.QuoteMeta("SELECT * FROM users WHERE id=?")

	sqlMock.ExpectQuery(sql).WillReturnRows(
		sqlmock.NewRows([]string{"id", "age", "name", "email"}).
			AddRow(1, 40, "marco", "m.siviero83@gmail.com"),
	)

	actual, _ := underTest.GetUser(1)

	assert.Equal(t, actual.Age, int32(40))
	assert.Equal(t, actual.Name, "marco")
	assert.Equal(t, actual.Email, "m.siviero83@gmail.com")
}
