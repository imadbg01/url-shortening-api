package model

func GetAllShorteners() ([]Shorten, error) {
	var shortener []Shorten
	tx := db.Find(&shortener)
	if tx.Error != nil {
		return []Shorten{}, tx.Error
	}
	return shortener, nil
}

func GetShorten(id uint64) (Shorten, error) {
	var shorten Shorten
	tx := db.Where("id = ?", id).First(&shorten)
	if tx.Error != nil {
		return Shorten{}, tx.Error
	}
	return shorten, nil
}

func CreateShorten(shorten Shorten) error {
	tx := db.Create(&shorten)
	return tx.Error
}

func UpdateShorten(shorten Shorten) error {
	tx := db.Save(&shorten)
	return tx.Error
}

func DeleteShorten(id uint64) error {
	tx := db.Unscoped().Delete(&Shorten{}, id)
	return tx.Error
}

func FindShortenUrl(url string) (Shorten, error) {
	var shorten Shorten

	tx := db.Where("id = ?", url).First(&shorten)

	return  shorten ,tx.Error
}

