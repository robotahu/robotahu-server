package postgres

import "fmt"

type Image struct {
	ID      int
	Url     string
	Caption string
}

func (db *Db) GetImages() ([]Image, error) {
	query, err := db.Prepare("SELECT * FROM images")
	if err != nil {
		fmt.Println("GetImages Preparation Err:", err)
		return nil, err
	}

	rows, err := query.Query()
	if err != nil {
		fmt.Println("GetImages Query Err:", err)
		return nil, err
	}

	var image Image
	images := []Image{}

	for rows.Next() {
		err = rows.Scan(
			&image.ID,
			&image.Url,
			&image.Caption,
		)

		if err != nil {
			fmt.Println("Error scannig row:", err)
			return nil, err
		}

		images = append(images, image)
	}

	return images, nil
}
