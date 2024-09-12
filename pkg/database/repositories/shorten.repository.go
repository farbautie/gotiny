package repositories

import (
	"github.com/farbautie/gotiny/pkg/database/entities"
)

func (r *Repositories) Save(url string, shortUrl string) (*entities.Link, error) {
	query := `INSERT INTO links (url, short_url) VALUES ($1, $2) RETURNING id, url, short_url, created_at, updated_at;`
	var link entities.Link
	result, err := r.pool.Query(query, url, shortUrl)
	if err != nil {
		return nil, err
	}
	defer result.Close()

	if !result.Next() {
		return nil, nil
	}

	result.Scan(&link.ID, &link.Url, &link.ShortUrl, &link.CreatedAt, &link.UpdatedAt)

	return &link, nil
}

func (r *Repositories) GetByShortUrl(shortUrl string) (*entities.Link, error) {
	query := `SELECT id, url, short_url, created_at, updated_at FROM links WHERE short_url = $1 AND deleted_at IS NULL;`
	var link entities.Link
	result, err := r.pool.Query(query, shortUrl)
	if err != nil {
		return nil, err
	}
	defer result.Close()
	if !result.Next() {
		return nil, nil
	}

	result.Scan(&link.ID, &link.Url, &link.ShortUrl, &link.CreatedAt, &link.UpdatedAt)

	return &link, nil
}

func (r *Repositories) Update(id string, link *entities.Link) error {
	query := `UPDATE links SET url = $1, short_url = $2, updated_at = NOW() WHERE id = $3;`
	result, err := r.pool.Query(query, link.Url, link.ShortUrl, id)
	if err != nil {
		return err
	}
	defer result.Close()

	if !result.Next() {
		return nil
	}

	return nil
}

func (r *Repositories) UpdateStats(id string) (*entities.Link, error) {
	var link entities.Link
	query := `UPDATE links SET hits = hits + 1 WHERE id = $1 AND deleted_at IS NULL RETURNING id, url, short_url, created_at, updated_at;`
	result, err := r.pool.Query(query, id)
	if err != nil {
		return nil, err
	}
	defer result.Close()

	if !result.Next() {
		return nil, nil
	}

	err = result.Scan(&link.ID, &link.Url, &link.ShortUrl, &link.CreatedAt, &link.UpdatedAt)
	if err != nil {
		return nil, err
	}

	return &link, nil
}

func (r *Repositories) Delete(id string) error {
	query := `UPDATE links SET deleted_at = NOW() WHERE id = $1;`
	result, err := r.pool.Query(query, id)
	if err != nil {
		return err
	}
	defer result.Close()

	if !result.Next() {
		return nil
	}

	return nil
}

func (r *Repositories) GetStats(id string) (*entities.Link, error) {
	var link entities.Link
	query := `SELECT id, url, short_url, hits, created_at, updated_at FROM links WHERE id = $1 AND deleted_at IS NULL;`
	result, err := r.pool.Query(query, id)
	if err != nil {
		return nil, err
	}
	defer result.Close()
	if !result.Next() {
		return nil, nil
	}

	err = result.Scan(&link.ID, &link.Url, &link.ShortUrl, &link.Hits, &link.CreatedAt, &link.UpdatedAt)
	if err != nil {
		return nil, err
	}

	return &link, nil
}
