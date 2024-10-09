package postgres

import (
    "github.com/uozi-tech/cosy/settings"
    "gorm.io/driver/postgres"
    "testing"
)

func TestOpen(t *testing.T) {
    dbs := &settings.DataBase{
        User:     "cosy",
        Password: "cosy",
        Host:     "127.0.0.1",
        Port:     5432,
        Name:     "cosy",
    }

    dialector := Open(dbs)

    d, ok := dialector.(*postgres.Dialector)
    if !ok {
        t.Fatal("dialector is not *Dialector")
    }
    if d.DSN == "" {
        t.Error("dialector.DSN is empty")
    }
}
