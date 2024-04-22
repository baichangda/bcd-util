package support_sql

import (
	"database/sql"
	_ "modernc.org/sqlite"
	"testing"
)

func Test(t *testing.T) {
	open, err := sql.Open("sqlite", ":memory:")
	//open, err := sql.Open("sqlite", "test.db")
	if err != nil {
		t.Errorf("%+v", err)
		return
	}
	defer open.Close()
	_, err = open.Exec("create table test(id int primary key ,name varchar(50) not null)")
	if err != nil {
		t.Errorf("%+v", err)
		return
	}
	_, err = open.Exec("INSERT INTO test(id,name) values (1,'test1'),(2,'test2')")
	if err != nil {
		t.Errorf("%+v", err)
		return
	}
	query, err := open.Query("select * from test")
	if err != nil {
		t.Errorf("%+v", err)
		return
	}
	for query.Next() {
		var id int
		var name string
		err = query.Scan(&id, &name)
		if err != nil {
			t.Errorf("%+v", err)
			return
		}
		t.Logf("%d %s", id, name)
	}
}
