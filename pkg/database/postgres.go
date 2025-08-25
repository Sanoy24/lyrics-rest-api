package database

import (
	"fmt"
	"log"
	// "os"
	"time"

	"github.com/Sanoy24/lyrics-rest-api/internal/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func InitDB(dbParam *models.PostgresParam) error {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable",
		dbParam.DB_HOST, dbParam.DB_USER, dbParam.DB_PASSWORD, dbParam.DB_NAME, dbParam.DB_PORT)

	// newLogger := logger.New(log.New(os.Stdout, "\r\n", log.LstdFlags), logger.Config{
	// 	SlowThreshold:             200 * time.Millisecond,
	// 	LogLevel:                  logger.Info,
	// 	ParameterizedQueries:      true,
	// 	IgnoreRecordNotFoundError: true,
	// 	Colorful:                  true,
	// })

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})

	if err != nil {
		return fmt.Errorf("failed to connect to database: %w", err)
	}
	psqlDB, err := db.DB()

	if err != nil {
		return fmt.Errorf("failed to get database instance: %w", err)
	}

	psqlDB.SetMaxIdleConns(10)                  // Maximum number of connections in the idle connection pool.
	psqlDB.SetMaxOpenConns(100)                 // Maximum number of open connections to the database.
	psqlDB.SetConnMaxLifetime(time.Minute * 60) // Maximum amount of time a connection may be reused.
	psqlDB.SetConnMaxIdleTime(time.Minute * 1)  // Maximum amount of time an idle connection can remain in the pool

	// Test connection
	if err := psqlDB.Ping(); err != nil {
		return fmt.Errorf("failed to ping database: %w", err)
	}

	DB = db

	log.Println("Connected to database successfully")

	// Optional: AutoMigrate your models (useful for development, consider separate migrations for production)
	// Example:
	// err = db.AutoMigrate(&models.Lyric{}, &models.Artist{}) // Replace with your actual models
	// if err != nil {
	// 	return fmt.Errorf("failed to auto migrate database: %w", err)
	// }
	// log.Println("Database auto-migration completed."

	return nil
}

func GetDB() *gorm.DB {
	if DB == nil {
		panic("DB not connected")
	}
	return DB
}

func CloseDB() {
	if DB != nil {
		psqlDB, err := DB.DB()
		if err != nil {
			log.Println("Error getting database instance:", err)
			return
		}

		if err := psqlDB.Close(); err != nil {
			log.Println("Error closing database connection:", err)
		} else {
			log.Println("Database connection closed successfully")
		}
	}
}

func Migrate(db *gorm.DB) error {
	if db == nil {
		return fmt.Errorf("database connection is not initialized")
	}

	log.Println("Running database migrations...")

	err := db.AutoMigrate(
		&models.User{},
		&models.Role{},
		&models.Permission{},
		&models.Artist{},
		&models.Album{},
		&models.Song{},
		&models.Annotation{},
		&models.Comment{},
		&models.Vote{},
		&models.SongArtist{},
		&models.SongTag{},
		&models.ArtistTag{},
	)

	if err != nil {
		return fmt.Errorf("failed to auto-migrate: %w", err)
	}
	log.Println("Database migration completed successfully")
	return nil
}

func SetupTsVector(db *gorm.DB) error {
	db.Exec(`ALTER TABLE songs ADD COLUMN IF NOT EXITS search_vector tsvector;`)
	db.Exec(`
	CREATE OR REPLACE FUNCTION songs_search_vector_update() RETURNS trigger AS $$
BEGIN
    NEW.search_vector :=
        setweight(to_tsvector('english', coalesce(NEW.title,'')),'A') ||
        setweight(to_tsvector('english', coalesce(NEW.lyrics,'')),'B') ||
        setweight(to_tsvector('english', coalesce(NEW.description,'')),'C') ||
        setweight(to_tsvector('english', coalesce(
            (SELECT title FROM albums WHERE albums.id = NEW.album_id), '')), 'B') ||
        setweight(to_tsvector('english', coalesce(
            (SELECT string_agg(artists.name, ' ')
             FROM song_artists
             JOIN artists ON artists.id = song_artists.artist_id
             WHERE song_artists.song_id = NEW.id), '')
        ), 'A');
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;
	`)

	db.Exec(`
    DROP TRIGGER IF EXISTS songs_search_vector_trigger ON songs;
    CREATE TRIGGER songs_search_vector_trigger
    BEFORE INSERT OR UPDATE ON songs
    FOR EACH ROW
    EXECUTE FUNCTION songs_search_vector_update();
    `)

	db.Exec(`CREATE INDEX IF NOT EXISTS songs_search_vector_idx ON songs USING GIN (search_vector);`)

	return nil
}
