package models

import (
    "fmt"
    "time"
)

type FileData struct {
    Id int
    FileName string
    Owner string
    Path string
    Timestamp time.Time
}

func (fd *FileData) Insert() {
    sqlStatement :=`INSERT INTO file_data (filename, owner, path, uploaded_at)
                    VALUES ($1, $2, $3, $4)`
    _, err := db.Exec(sqlStatement, fd.FileName, fd.Owner, fd.Path, time.Now())
    if err != nil {
        fmt.Println(err)
    }
}
