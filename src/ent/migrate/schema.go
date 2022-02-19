// Code generated by entc, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// CommentsColumns holds the columns for the "comments" table.
	CommentsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID},
		{Name: "created", Type: field.TypeTime},
		{Name: "name", Type: field.TypeString, Size: 256},
		{Name: "message", Type: field.TypeString, Size: 1024},
		{Name: "menu", Type: field.TypeEnum, Enums: []string{"CLASSIC", "VEGETARIAN", "VEGAN"}},
		{Name: "event_comments", Type: field.TypeUUID, Nullable: true},
	}
	// CommentsTable holds the schema information for the "comments" table.
	CommentsTable = &schema.Table{
		Name:       "comments",
		Columns:    CommentsColumns,
		PrimaryKey: []*schema.Column{CommentsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "comments_events_comments",
				Columns:    []*schema.Column{CommentsColumns[5]},
				RefColumns: []*schema.Column{EventsColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// EventsColumns holds the columns for the "events" table.
	EventsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID},
		{Name: "created", Type: field.TypeTime},
		{Name: "last_modified", Type: field.TypeTime},
		{Name: "theme", Type: field.TypeString, Size: 256},
		{Name: "title_image", Type: field.TypeString},
		{Name: "starting_time", Type: field.TypeTime},
		{Name: "closing_time", Type: field.TypeTime, Nullable: true},
		{Name: "starter", Type: field.TypeString, Size: 512},
		{Name: "main_dish", Type: field.TypeString, Size: 512},
		{Name: "dessert", Type: field.TypeString, Size: 512},
		{Name: "description", Type: field.TypeString, Size: 2048},
	}
	// EventsTable holds the schema information for the "events" table.
	EventsTable = &schema.Table{
		Name:       "events",
		Columns:    EventsColumns,
		PrimaryKey: []*schema.Column{EventsColumns[0]},
	}
	// ImagesColumns holds the columns for the "images" table.
	ImagesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID},
		{Name: "created", Type: field.TypeTime},
		{Name: "file_name", Type: field.TypeString, Size: 256},
		{Name: "event_images", Type: field.TypeUUID, Nullable: true},
	}
	// ImagesTable holds the schema information for the "images" table.
	ImagesTable = &schema.Table{
		Name:       "images",
		Columns:    ImagesColumns,
		PrimaryKey: []*schema.Column{ImagesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "images_events_images",
				Columns:    []*schema.Column{ImagesColumns[3]},
				RefColumns: []*schema.Column{EventsColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// ParticipantsColumns holds the columns for the "participants" table.
	ParticipantsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID},
		{Name: "created", Type: field.TypeTime},
		{Name: "name", Type: field.TypeString, Size: 256},
		{Name: "message", Type: field.TypeString, Size: 256},
		{Name: "classic_menu", Type: field.TypeInt},
		{Name: "vegetarian_menu", Type: field.TypeInt},
		{Name: "vegan_menu", Type: field.TypeInt},
		{Name: "event_participants", Type: field.TypeUUID, Nullable: true},
	}
	// ParticipantsTable holds the schema information for the "participants" table.
	ParticipantsTable = &schema.Table{
		Name:       "participants",
		Columns:    ParticipantsColumns,
		PrimaryKey: []*schema.Column{ParticipantsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "participants_events_participants",
				Columns:    []*schema.Column{ParticipantsColumns[7]},
				RefColumns: []*schema.Column{EventsColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// UsersColumns holds the columns for the "users" table.
	UsersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID},
		{Name: "created", Type: field.TypeTime},
		{Name: "name", Type: field.TypeString, Unique: true, Size: 256},
		{Name: "password", Type: field.TypeBytes},
	}
	// UsersTable holds the schema information for the "users" table.
	UsersTable = &schema.Table{
		Name:       "users",
		Columns:    UsersColumns,
		PrimaryKey: []*schema.Column{UsersColumns[0]},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		CommentsTable,
		EventsTable,
		ImagesTable,
		ParticipantsTable,
		UsersTable,
	}
)

func init() {
	CommentsTable.ForeignKeys[0].RefTable = EventsTable
	ImagesTable.ForeignKeys[0].RefTable = EventsTable
	ParticipantsTable.ForeignKeys[0].RefTable = EventsTable
}
