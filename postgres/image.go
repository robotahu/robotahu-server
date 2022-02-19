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
	defer query.Close()

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

func (db *Db) CreateImage(url, caption string) (*Image, error) {
	query, err := db.Prepare("INSERT INTO images(url, caption) VALUES(?, ?)")
	if err != nil {
		fmt.Println("CreateImage Preparation Err:", err)
		return nil, err
	}
	defer query.Close()

	row := query.QueryRow(url, caption)
	if err != nil {
		fmt.Println("CreateImage Query Err:", err)
		return nil, err
	}

	var image Image
	err = row.Scan(
		&image.ID,
		&image.Url,
		&image.Caption,
	)

	if err != nil {
		fmt.Println("Error scannig row:", err)
		return nil, err
	}

	return &image, nil
}
