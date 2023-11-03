package domain

type Gallery struct {
	Id          string
	Name        string
	Description string
	CoverUrl    string
}

type Image struct {
	Id          string
	Name        string
	Description string
	ImageUrl    string
	GalleryId   string
}
