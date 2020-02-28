package models

import (
	"fmt"
	"time"
)

type FileData struct {
	Id        int
	FileName  string
	Owner     string
	Path      string
	Timestamp time.Time
}

func AllFileData(owner string) ([]*FileData, error) {
	rows, err := db.Query("SELECT * FROM file_data WHERE owner=$1", owner)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	fileData := make([]*FileData, 0)
	for rows.Next() {
		fd := new(FileData)
		err := rows.Scan(&fd.Id, &fd.FileName, &fd.Owner, &fd.Path, &fd.Timestamp)
		if err != nil {
			return nil, err
		}
		fileData = append(fileData, fd)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}
	return fileData, nil
}

func GetFileData(id int, owner string) (*FileData, error) {
	row := db.QueryRow("SELECT * FROM file_data WHERE id=$1 AND owner=$2", id, owner)

	var fd FileData

	err := row.Scan(&fd.Id, &fd.FileName, &fd.Owner, &fd.Path, &fd.Timestamp)
	if err != nil {
		return nil, err
	}

	return &fd, nil
}

func (fd *FileData) Insert() {
	sqlStatement := `INSERT INTO file_data (filename, owner, path, uploaded_at)
                    VALUES ($1, $2, $3, $4)`
	_, err := db.Exec(sqlStatement, fd.FileName, fd.Owner, fd.Path, time.Now())
	if err != nil {
		fmt.Println(err)
	}
}
