package database

import (
	"log"

	"github.com/Sanoy24/lyrics-rest-api/internal/models"
	"github.com/Sanoy24/lyrics-rest-api/pkg/data"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type SeedData struct {
	Permissions []models.Permission
	Roles       []models.Role
}

var seedData = SeedData{
	Permissions: data.DefaultPermissions,
	Roles:       data.DefaultRoles,
}

func Seed(db *gorm.DB) error {
	log.Println("Seeding database...")

	// Seed permissions
	if err := seedPermissions(db); err != nil {
		return err
	}

	// Seed roles
	if err := seedRoles(db); err != nil {
		return err
	}

	// Assign permissions to admin role
	if err := assignPermissionsToAdmin(db); err != nil {
		return err
	}

	// Assign permissions to other roles
	if err := assignAllRolePermissions(db); err != nil {
		return err
	}

	log.Println("Database seeding completed successfully.")
	return nil
}

func seedPermissions(db *gorm.DB) error {
	for _, perm := range seedData.Permissions {
		err := db.Clauses(clause.OnConflict{
			Columns:   []clause.Column{{Name: "name"}},
			DoNothing: true,
		}).Create(&perm).Error
		if err != nil {
			return err
		}
	}
	return nil
}

func seedRoles(db *gorm.DB) error {
	for _, role := range seedData.Roles {
		err := db.Clauses(clause.OnConflict{
			Columns:   []clause.Column{{Name: "name"}},
			DoNothing: true,
		}).Create(&role).Error
		if err != nil {
			return err
		}
	}
	return nil
}

func assignPermissionsToAdmin(db *gorm.DB) error {
	var adminRole models.Role
	if err := db.Where("name=?", "admin").First(&adminRole).Error; err != nil {
		return err
	}
	var allPermissions []models.Permission
	if err := db.Find(&allPermissions).Error; err != nil {
		return err
	}
	return db.Model(&adminRole).Association("Permissions").Replace(&allPermissions)
}

func assignAllRolePermissions(db *gorm.DB) error {
	// Moderator: limited moderation
	if err := assignPermissionsToRole(db, "moderator", []string{
		"moderate_content", "moderate_users",
	}); err != nil {
		return err
	}

	// Verified User: Can contribute and verify content
	if err := assignPermissionsToRole(db, "verified_user", []string{
		"song_create", "song_update", "annotation_create", "annotation_update",
		"comment_create", "comment_update", "vote_create", "vote_update",
	}); err != nil {
		return err
	}

	// User: can read and interact
	if err := assignPermissionsToRole(db, "user", []string{
		"song_read", "artist_read", "annotation_read", "comment_read", "vote_read",
		"comment_create", "vote_create",
	}); err != nil {
		return err
	}

	// Guest: read-only
	if err := assignPermissionsToRole(db, "guest", []string{
		"song_read", "artist_read", "annotation_read", "comment_read", "vote_read",
	}); err != nil {
		return err
	}

	return nil
}

func assignPermissionsToRole(db *gorm.DB, roleName string, permissionNames []string) error {
	var role models.Role
	if err := db.Where("name=?", roleName).First(&role).Error; err != nil {
		return err
	}
	var permissions []models.Permission
	if err := db.Where("name IN ?", permissionNames).Find(&permissions).Error; err != nil {
		return err
	}
	return db.Model(&role).Association("Permissions").Replace(&permissions)
}
