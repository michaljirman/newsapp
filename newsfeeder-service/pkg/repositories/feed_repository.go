package repositories

import (
	"context"
	"database/sql"

	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"

	"github.com/michaljirman/newsapp/newsfeeder-service/pkg/configs"
	"github.com/michaljirman/newsapp/newsfeeder-service/pkg/models"
)

// FeedRepository describes the persistence methods on the Feed model
type FeedRepository interface {
	CreateFeed(ctx context.Context, feed models.Feed) (uint64, error)
	GetFeeds(ctx context.Context, category, provider string) ([]models.Feed, error)
	GetFeedByID(ctx context.Context, feedID uint64) (models.Feed, error)
}

// feedRepository implements FeedRepository interface to work on the Feed model
type feedRepository struct {
	db     *sqlx.DB
	Config *configs.DBConfig
	logger *logrus.Logger
}

// NewFeedRepository creates a new instance of FeedRepository
func NewFeedRepository(cfg *configs.DBConfig, dbHandle *sql.DB, logger *logrus.Logger) (FeedRepository, error) {
	db := sqlx.NewDb(dbHandle, "postgres")
	repo := &feedRepository{
		db:     db,
		Config: cfg,
		logger: logger,
	}
	return repo, nil
}

// CreateFeed inserts a new feed to the database
func (repo *feedRepository) CreateFeed(ctx context.Context, feed models.Feed) (uint64, error) {
	tx := repo.db.MustBegin()
	var feedID uint64
	tx.QueryRowx("INSERT INTO feed (category, provider, url) VALUES ($1, $2, $3) RETURNING feed_id", feed.Category, feed.Provider, feed.URL).Scan(&feedID)
	err := tx.Commit()
	if err != nil {
		return 0, errors.Wrap(err, "failed to instert a new feed to the db")
	}
	return feedID, nil
}

// GetFeeds retrieves all feeds from the database
// TODO use https://github.com/masterminds/squirrel for more complex filtering
func (repo *feedRepository) GetFeeds(ctx context.Context, category, provider string) ([]models.Feed, error) {
	feeds := []models.Feed{}
	query := "SELECT * FROM feed"
	args := make([]interface{}, 0)
	if category != "" {
		query += " WHERE category like $1"
		args = append(args, category)
	}
	if provider != "" {
		if category != "" {
			query += " AND provider like $2"
		} else {
			query += " WHERE provider like $1"
		}
		args = append(args, provider)
	}
	err := repo.db.Select(&feeds, repo.db.Rebind(query), args...)
	if err != nil {
		return nil, errors.Wrap(err, "failed to query db for existing feeds")
	}
	return feeds, nil
}

// GetFeedByID retrieves a single feed by
func (repo *feedRepository) GetFeedByID(ctx context.Context, feedID uint64) (models.Feed, error) {
	feed := models.Feed{}
	err := repo.db.Get(&feed, "SELECT * FROM feed where feed_id = $1", feedID)
	if err != nil {
		return models.Feed{}, errors.Wrap(err, "failed to get feed by ID")
	}
	return feed, nil
}
