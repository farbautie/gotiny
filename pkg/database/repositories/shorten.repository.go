package repositories

import "github.com/farbautie/gotiny/pkg/database/entities"

func (r *Repositories) Save(url string, shortUrl string) (*entities.Link, error) {
	return nil, nil
}

func (r *Repositories) GetByShortUrl(shortUrl string) (*entities.Link, error) {
	return nil, nil
}

func (r *Repositories) Update(id string, link *entities.Link) error {
	return nil
}

func (r *Repositories) Delete(id string) error {
	return nil
}
