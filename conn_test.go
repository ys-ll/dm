// Copyright © 2021 Bin Liu <bin.liu@enmotech.com>

package dm

import (
	"database/sql"
	"encoding/hex"
	"fmt"
	"testing"
)

func Test_Connect(t *testing.T) {
	dsn := "dm://sysdba:SYSDBA001@127.0.0.1:5236?appName=mtk&connectTimeout=3000&logLevel=all&clientEncoding=GB18030"
	db, err := sql.Open("dm", dsn)
	if err != nil {
		t.Error(err)
		return
	}
	sqlText := "select id,col2 from mtk.tab_gbk"
	var id, col2 string

	err = db.QueryRow(sqlText).Scan(&id, &col2)
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Println(id, hex.Dump([]byte(col2)))
}
