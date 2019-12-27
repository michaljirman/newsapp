package repositories

import (
	"context"
	"reflect"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"

	"github.com/michaljirman/newsapp/newsfeeder-service/pkg/configs"
	"github.com/michaljirman/newsapp/newsfeeder-service/pkg/models"
)

func setup(t *testing.T) (*configs.Config, *logrus.Logger, context.Context) {
	godotenv.Load("../.env")

	//todo setup method
	cfg, err := configs.Get()
	if err != nil {
		t.Fatalf("failed to load config: %+v", err)
	}

	// setup global logger settings
	ll, err := logrus.ParseLevel("debug")
	if err != nil {
		t.Fatalf("invalid value for loglevel (could not be parsed): %+v", err)
	}
	logger := logrus.New()
	logger.SetReportCaller(true)
	logger.SetLevel(ll)
	return cfg, logger, context.Background()
}

func prepareTestFeeds() ([]models.Feed, *sqlmock.Rows) {
	timeNow := time.Now()
	expectedRows := sqlmock.NewRows([]string{"feed_id", "created_at", "updated_at", "category", "provider", "url"}).
		AddRow(1, timeNow, timeNow, "UK", "BBC", "http://feeds.bbci.co.uk/news/uk/rss.xml").
		AddRow(2, timeNow, timeNow, "Technology", "BBC", "http://feeds.bbci.co.uk/news/technology/rss.xml").
		AddRow(3, timeNow, timeNow, "UK", "Reuters", "http://feeds.reuters.com/reuters/UKdomesticNews?format=xml").
		AddRow(4, timeNow, timeNow, "Technology", "Reuters", "http://feeds.reuters.com/reuters/UKdomesticNews?format=xml")

	// Expected results
	var expectedFeeds []models.Feed
	expectedFeed1 := models.Feed{
		ID:       1,
		Created:  timeNow,
		Updated:  timeNow,
		Category: "UK",
		Provider: "BBC",
		URL:      "http://feeds.bbci.co.uk/news/uk/rss.xml",
	}

	expectedFeed2 := models.Feed{
		ID:       2,
		Created:  timeNow,
		Updated:  timeNow,
		Category: "Technology",
		Provider: "BBC",
		URL:      "http://feeds.bbci.co.uk/news/technology/rss.xml",
	}

	expectedFeed3 := models.Feed{
		ID:       3,
		Created:  timeNow,
		Updated:  timeNow,
		Category: "UK",
		Provider: "Reuters",
		URL:      "http://feeds.reuters.com/reuters/UKdomesticNews?format=xml",
	}

	expectedFeed4 := models.Feed{
		ID:       4,
		Created:  timeNow,
		Updated:  timeNow,
		Category: "Technology",
		Provider: "Reuters",
		URL:      "http://feeds.reuters.com/reuters/UKdomesticNews?format=xml",
	}
	expectedFeeds = append(expectedFeeds, expectedFeed1)
	expectedFeeds = append(expectedFeeds, expectedFeed2)
	expectedFeeds = append(expectedFeeds, expectedFeed3)
	expectedFeeds = append(expectedFeeds, expectedFeed4)
	return expectedFeeds, expectedRows
}

func TestGetFeeds(t *testing.T) {
	cfg, logger, ctx := setup(t)
	mockDB, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer mockDB.Close()

	repo, err := NewFeedRepository(&cfg.Db, mockDB, logger)
	if err != nil {
		t.Fatal("failed to create a new repository")
	}

	expectedFeeds, expectedRows := prepareTestFeeds()
	mock.ExpectQuery("^SELECT (.+) FROM feed").WillReturnRows(expectedRows)

	feeds, err := repo.GetFeeds(ctx, "", "")
	if err != nil {
		t.Fatalf("failed to get feeds from db mock %+v", err)
	}

	if !reflect.DeepEqual(feeds, expectedFeeds) {
		t.Errorf("got %v want %v", feeds, expectedFeeds)
	}
}

func TestGetFeedByCategory(t *testing.T) {
	cfg, logger, ctx := setup(t)
	mockDB, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer mockDB.Close()

	repo, err := NewFeedRepository(&cfg.Db, mockDB, logger)
	if err != nil {
		t.Fatal("failed to create a new repository")
	}

	categoryArg := "UK"
	expectedFeeds, expectedRows := prepareTestFeeds()

	// mock.ExpectPrepare("^SELECT (.+) FROM feed WHERE").ExpectQuery().WithArgs(categoryArg).WillReturnRows(expectedRows)
	mock.ExpectQuery("^SELECT (.+) FROM feed WHERE").WithArgs(categoryArg).WillReturnRows(expectedRows)

	feeds, err := repo.GetFeeds(ctx, categoryArg, "")
	if err != nil {
		t.Fatalf("failed to get feeds from db mock %+v", err)
	}

	if !reflect.DeepEqual(feeds, expectedFeeds) {
		t.Errorf("got %v want %v", feeds, expectedFeeds)
	}
}

func TestGetFeedByCategoryAndProvider(t *testing.T) {
	cfg, logger, ctx := setup(t)
	mockDB, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer mockDB.Close()

	repo, err := NewFeedRepository(&cfg.Db, mockDB, logger)
	if err != nil {
		t.Fatal("failed to create a new repository")
	}

	categoryArg := "UK"
	providerArg := "BBC"
	expectedFeeds, expectedRows := prepareTestFeeds()

	// mock.ExpectPrepare("^SELECT (.+) FROM feed WHERE").ExpectQuery().WithArgs(categoryArg).WillReturnRows(expectedRows)
	mock.ExpectQuery("^SELECT (.+) FROM feed WHERE").WithArgs(categoryArg, providerArg).WillReturnRows(expectedRows)

	feeds, err := repo.GetFeeds(ctx, categoryArg, providerArg)
	if err != nil {
		t.Fatalf("failed to get feeds from db mock %+v", err)
	}

	if !reflect.DeepEqual(feeds, expectedFeeds) {
		t.Errorf("got %v want %v", feeds, expectedFeeds)
	}
}
