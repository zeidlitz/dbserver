package trashdatabase_test

import (
	"github.com/zeidlitz/dbserver/internal/trashdatabase"
	"testing"
)

func TestTrashDBConnection(t *testing.T) {
	var db trashdatabase.TrashDB
	err := db.Connect("trashconnection")
	if err != nil {
		t.Fatal("Failed to connect to TrashDB")
	}

}
