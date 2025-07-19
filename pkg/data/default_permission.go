package data

import "github.com/Sanoy24/lyrics-rest-api/internal/models"

var DefaultPermissions = []models.Permission{
	// User permissions
	{Name: "user_create", Resource: "user", Action: "create"},
	{Name: "user_read", Resource: "user", Action: "read"},
	{Name: "user_update", Resource: "user", Action: "update"},
	{Name: "user_delete", Resource: "user", Action: "delete"},

	// Song permissions
	{Name: "song_create", Resource: "song", Action: "create"},
	{Name: "song_read", Resource: "song", Action: "read"},
	{Name: "song_update", Resource: "song", Action: "update"},
	{Name: "song_delete", Resource: "song", Action: "delete"},
	{Name: "song_verify", Resource: "song", Action: "verify"},

	// Artist permissions
	{Name: "artist_create", Resource: "artist", Action: "create"},
	{Name: "artist_read", Resource: "artist", Action: "read"},
	{Name: "artist_update", Resource: "artist", Action: "update"},
	{Name: "artist_delete", Resource: "artist", Action: "delete"},
	{Name: "artist_verify", Resource: "artist", Action: "verify"},

	// Annotation permissions
	{Name: "annotation_create", Resource: "annotation", Action: "create"},
	{Name: "annotation_read", Resource: "annotation", Action: "read"},
	{Name: "annotation_update", Resource: "annotation", Action: "update"},
	{Name: "annotation_delete", Resource: "annotation", Action: "delete"},
	{Name: "annotation_verify", Resource: "annotation", Action: "verify"},

	// Comment permissions
	{Name: "comment_create", Resource: "comment", Action: "create"},
	{Name: "comment_read", Resource: "comment", Action: "read"},
	{Name: "comment_update", Resource: "comment", Action: "update"},
	{Name: "comment_delete", Resource: "comment", Action: "delete"},

	// Vote permissions
	{Name: "vote_create", Resource: "vote", Action: "create"},
	{Name: "vote_read", Resource: "vote", Action: "read"},
	{Name: "vote_update", Resource: "vote", Action: "update"},
	{Name: "vote_delete", Resource: "vote", Action: "delete"},

	// Album permissions
	{Name: "album_create", Resource: "album", Action: "create"},
	{Name: "album_read", Resource: "album", Action: "read"},
	{Name: "album_update", Resource: "album", Action: "update"},
	{Name: "album_delete", Resource: "album", Action: "delete"},

	// Moderation permissions
	{Name: "moderate_content", Resource: "content", Action: "moderate"},
	{Name: "moderate_users", Resource: "user", Action: "moderate"},
}
