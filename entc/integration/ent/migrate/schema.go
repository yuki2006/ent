// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// CardsColumns holds the columns for the "cards" table.
	CardsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "create_time", Type: field.TypeTime},
		{Name: "update_time", Type: field.TypeTime},
		{Name: "balance", Type: field.TypeFloat64, Default: 0},
		{Name: "number", Type: field.TypeString},
		{Name: "name", Type: field.TypeString, Nullable: true},
		{Name: "user_card", Type: field.TypeInt, Unique: true, Nullable: true},
	}
	// CardsTable holds the schema information for the "cards" table.
	CardsTable = &schema.Table{
		Name:       "cards",
		Columns:    CardsColumns,
		PrimaryKey: []*schema.Column{CardsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "cards_users_card",
				Columns:    []*schema.Column{CardsColumns[6]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
		Indexes: []*schema.Index{
			{
				Name:    "card_id",
				Unique:  false,
				Columns: []*schema.Column{CardsColumns[0]},
			},
			{
				Name:    "card_number",
				Unique:  true,
				Columns: []*schema.Column{CardsColumns[4]},
			},
			{
				Name:    "card_id_name_number",
				Unique:  false,
				Columns: []*schema.Column{CardsColumns[0], CardsColumns[5], CardsColumns[4]},
			},
		},
	}
	// CommentsColumns holds the columns for the "comments" table.
	CommentsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "unique_int", Type: field.TypeInt, Unique: true},
		{Name: "unique_float", Type: field.TypeFloat64, Unique: true},
		{Name: "nillable_int", Type: field.TypeInt, Nullable: true},
		{Name: "table", Type: field.TypeString, Nullable: true},
		{Name: "dir", Type: field.TypeJSON, Nullable: true},
	}
	// CommentsTable holds the schema information for the "comments" table.
	CommentsTable = &schema.Table{
		Name:       "comments",
		Columns:    CommentsColumns,
		PrimaryKey: []*schema.Column{CommentsColumns[0]},
	}
	// FieldTypesColumns holds the columns for the "field_types" table.
	FieldTypesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "int", Type: field.TypeInt},
		{Name: "int8", Type: field.TypeInt8},
		{Name: "int16", Type: field.TypeInt16},
		{Name: "int32", Type: field.TypeInt32},
		{Name: "int64", Type: field.TypeInt64},
		{Name: "optional_int", Type: field.TypeInt, Nullable: true},
		{Name: "optional_int8", Type: field.TypeInt8, Nullable: true},
		{Name: "optional_int16", Type: field.TypeInt16, Nullable: true},
		{Name: "optional_int32", Type: field.TypeInt32, Nullable: true},
		{Name: "optional_int64", Type: field.TypeInt64, Nullable: true},
		{Name: "nillable_int", Type: field.TypeInt, Nullable: true},
		{Name: "nillable_int8", Type: field.TypeInt8, Nullable: true},
		{Name: "nillable_int16", Type: field.TypeInt16, Nullable: true},
		{Name: "nillable_int32", Type: field.TypeInt32, Nullable: true},
		{Name: "nillable_int64", Type: field.TypeInt64, Nullable: true},
		{Name: "validate_optional_int32", Type: field.TypeInt32, Nullable: true},
		{Name: "optional_uint", Type: field.TypeUint, Nullable: true},
		{Name: "optional_uint8", Type: field.TypeUint8, Nullable: true},
		{Name: "optional_uint16", Type: field.TypeUint16, Nullable: true},
		{Name: "optional_uint32", Type: field.TypeUint32, Nullable: true},
		{Name: "optional_uint64", Type: field.TypeUint64, Nullable: true},
		{Name: "state", Type: field.TypeEnum, Nullable: true, Enums: []string{"on", "off"}},
		{Name: "optional_float", Type: field.TypeFloat64, Nullable: true},
		{Name: "optional_float32", Type: field.TypeFloat32, Nullable: true},
		{Name: "text", Type: field.TypeString, Nullable: true, Size: 2147483647, SchemaType: map[string]string{"mysql": "mediumtext"}},
		{Name: "datetime", Type: field.TypeTime, Nullable: true, SchemaType: map[string]string{"mysql": "datetime", "postgres": "date"}},
		{Name: "decimal", Type: field.TypeFloat64, Nullable: true, SchemaType: map[string]string{"mysql": "decimal(6,2)", "postgres": "numeric"}},
		{Name: "link_other", Type: field.TypeOther, Nullable: true, SchemaType: map[string]string{"mysql": "varchar(255)", "postgres": "varchar", "sqlite3": "varchar(255)"}},
		{Name: "link_other_func", Type: field.TypeOther, Nullable: true, SchemaType: map[string]string{"mysql": "varchar(255)", "postgres": "varchar", "sqlite3": "varchar(255)"}},
		{Name: "mac", Type: field.TypeString, Nullable: true, SchemaType: map[string]string{"postgres": "macaddr"}},
		{Name: "string_array", Type: field.TypeOther, Nullable: true, SchemaType: map[string]string{"mysql": "blob", "postgres": "text[]", "sqlite3": "json"}},
		{Name: "password", Type: field.TypeString, Nullable: true, SchemaType: map[string]string{"mysql": "char(32)"}},
		{Name: "string_scanner", Type: field.TypeString, Nullable: true},
		{Name: "duration", Type: field.TypeInt64, Nullable: true},
		{Name: "dir", Type: field.TypeString},
		{Name: "ndir", Type: field.TypeString, Nullable: true},
		{Name: "str", Type: field.TypeString, Nullable: true},
		{Name: "null_str", Type: field.TypeString, Nullable: true},
		{Name: "link", Type: field.TypeString, Nullable: true},
		{Name: "null_link", Type: field.TypeString, Nullable: true},
		{Name: "active", Type: field.TypeBool, Nullable: true},
		{Name: "null_active", Type: field.TypeBool, Nullable: true},
		{Name: "deleted", Type: field.TypeBool, Nullable: true},
		{Name: "deleted_at", Type: field.TypeTime, Nullable: true},
		{Name: "raw_data", Type: field.TypeBytes, Nullable: true, Size: 20},
		{Name: "sensitive", Type: field.TypeBytes, Nullable: true},
		{Name: "ip", Type: field.TypeBytes, Nullable: true},
		{Name: "null_int64", Type: field.TypeInt, Nullable: true},
		{Name: "schema_int", Type: field.TypeInt, Nullable: true},
		{Name: "schema_int8", Type: field.TypeInt8, Nullable: true},
		{Name: "schema_int64", Type: field.TypeInt64, Nullable: true},
		{Name: "schema_float", Type: field.TypeFloat64, Nullable: true},
		{Name: "schema_float32", Type: field.TypeFloat32, Nullable: true},
		{Name: "null_float", Type: field.TypeFloat64, Nullable: true},
		{Name: "role", Type: field.TypeEnum, Enums: []string{"ADMIN", "OWNER", "USER", "READ", "WRITE"}, Default: "READ"},
		{Name: "priority", Type: field.TypeEnum, Nullable: true, Enums: []string{"UNKNOWN", "LOW", "HIGH"}},
		{Name: "optional_uuid", Type: field.TypeUUID, Nullable: true},
		{Name: "nillable_uuid", Type: field.TypeUUID, Nullable: true},
		{Name: "strings", Type: field.TypeJSON, Nullable: true},
		{Name: "pair", Type: field.TypeBytes},
		{Name: "nil_pair", Type: field.TypeBytes, Nullable: true},
		{Name: "vstring", Type: field.TypeString},
		{Name: "triple", Type: field.TypeString},
		{Name: "big_int", Type: field.TypeInt, Nullable: true},
		{Name: "password_other", Type: field.TypeOther, Nullable: true, SchemaType: map[string]string{"mysql": "char(32)", "postgres": "varchar", "sqlite3": "char(32)"}},
		{Name: "file_field", Type: field.TypeInt, Nullable: true},
	}
	// FieldTypesTable holds the schema information for the "field_types" table.
	FieldTypesTable = &schema.Table{
		Name:       "field_types",
		Columns:    FieldTypesColumns,
		PrimaryKey: []*schema.Column{FieldTypesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "field_types_files_field",
				Columns:    []*schema.Column{FieldTypesColumns[66]},
				RefColumns: []*schema.Column{FilesColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// FilesColumns holds the columns for the "files" table.
	FilesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "fsize", Type: field.TypeInt, Default: 2147483647},
		{Name: "name", Type: field.TypeString},
		{Name: "user", Type: field.TypeString, Nullable: true},
		{Name: "group", Type: field.TypeString, Nullable: true},
		{Name: "op", Type: field.TypeBool, Nullable: true},
		{Name: "file_type_files", Type: field.TypeInt, Nullable: true},
		{Name: "group_files", Type: field.TypeInt, Nullable: true},
		{Name: "user_files", Type: field.TypeInt, Nullable: true},
	}
	// FilesTable holds the schema information for the "files" table.
	FilesTable = &schema.Table{
		Name:       "files",
		Columns:    FilesColumns,
		PrimaryKey: []*schema.Column{FilesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "files_file_types_files",
				Columns:    []*schema.Column{FilesColumns[6]},
				RefColumns: []*schema.Column{FileTypesColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:     "files_groups_files",
				Columns:    []*schema.Column{FilesColumns[7]},
				RefColumns: []*schema.Column{GroupsColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:     "files_users_files",
				Columns:    []*schema.Column{FilesColumns[8]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
		Indexes: []*schema.Index{
			{
				Name:    "file_name_size",
				Unique:  false,
				Columns: []*schema.Column{FilesColumns[2], FilesColumns[1]},
			},
			{
				Name:    "file_name_user",
				Unique:  true,
				Columns: []*schema.Column{FilesColumns[2], FilesColumns[3]},
			},
			{
				Name:    "file_user_files_file_type_files",
				Unique:  false,
				Columns: []*schema.Column{FilesColumns[8], FilesColumns[6]},
			},
			{
				Name:    "file_name_user_files_file_type_files",
				Unique:  true,
				Columns: []*schema.Column{FilesColumns[2], FilesColumns[8], FilesColumns[6]},
			},
			{
				Name:    "file_name_user_files",
				Unique:  false,
				Columns: []*schema.Column{FilesColumns[2], FilesColumns[8]},
			},
		},
	}
	// FileTypesColumns holds the columns for the "file_types" table.
	FileTypesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "name", Type: field.TypeString, Unique: true},
		{Name: "type", Type: field.TypeEnum, Enums: []string{"png", "svg", "jpg"}, Default: "png"},
		{Name: "state", Type: field.TypeEnum, Enums: []string{"ON", "OFF"}, Default: "ON"},
	}
	// FileTypesTable holds the schema information for the "file_types" table.
	FileTypesTable = &schema.Table{
		Name:       "file_types",
		Columns:    FileTypesColumns,
		PrimaryKey: []*schema.Column{FileTypesColumns[0]},
	}
	// GoodsColumns holds the columns for the "goods" table.
	GoodsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
	}
	// GoodsTable holds the schema information for the "goods" table.
	GoodsTable = &schema.Table{
		Name:       "goods",
		Columns:    GoodsColumns,
		PrimaryKey: []*schema.Column{GoodsColumns[0]},
	}
	// GroupsColumns holds the columns for the "groups" table.
	GroupsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "active", Type: field.TypeBool, Default: true},
		{Name: "expire", Type: field.TypeTime},
		{Name: "type", Type: field.TypeString, Nullable: true, Size: 255},
		{Name: "max_users", Type: field.TypeInt, Nullable: true, Default: 10},
		{Name: "name", Type: field.TypeString},
		{Name: "group_info", Type: field.TypeInt},
	}
	// GroupsTable holds the schema information for the "groups" table.
	GroupsTable = &schema.Table{
		Name:       "groups",
		Columns:    GroupsColumns,
		PrimaryKey: []*schema.Column{GroupsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "groups_group_infos_info",
				Columns:    []*schema.Column{GroupsColumns[6]},
				RefColumns: []*schema.Column{GroupInfosColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
	}
	// GroupInfosColumns holds the columns for the "group_infos" table.
	GroupInfosColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "desc", Type: field.TypeString},
		{Name: "max_users", Type: field.TypeInt, Default: 10000},
	}
	// GroupInfosTable holds the schema information for the "group_infos" table.
	GroupInfosTable = &schema.Table{
		Name:       "group_infos",
		Columns:    GroupInfosColumns,
		PrimaryKey: []*schema.Column{GroupInfosColumns[0]},
	}
	// ItemsColumns holds the columns for the "items" table.
	ItemsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeString, Size: 64},
		{Name: "text", Type: field.TypeString, Unique: true, Nullable: true, Size: 128},
	}
	// ItemsTable holds the schema information for the "items" table.
	ItemsTable = &schema.Table{
		Name:       "items",
		Columns:    ItemsColumns,
		PrimaryKey: []*schema.Column{ItemsColumns[0]},
	}
	// LicensesColumns holds the columns for the "licenses" table.
	LicensesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true, SchemaType: map[string]string{"postgres": "bigserial"}},
	}
	// LicensesTable holds the schema information for the "licenses" table.
	LicensesTable = &schema.Table{
		Name:       "licenses",
		Columns:    LicensesColumns,
		PrimaryKey: []*schema.Column{LicensesColumns[0]},
	}
	// NodesColumns holds the columns for the "nodes" table.
	NodesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "value", Type: field.TypeInt, Nullable: true},
		{Name: "node_next", Type: field.TypeInt, Unique: true, Nullable: true},
	}
	// NodesTable holds the schema information for the "nodes" table.
	NodesTable = &schema.Table{
		Name:       "nodes",
		Columns:    NodesColumns,
		PrimaryKey: []*schema.Column{NodesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "nodes_nodes_next",
				Columns:    []*schema.Column{NodesColumns[2]},
				RefColumns: []*schema.Column{NodesColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// PetColumns holds the columns for the "pet" table.
	PetColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "age", Type: field.TypeFloat64, Default: 0},
		{Name: "name", Type: field.TypeString},
		{Name: "uuid", Type: field.TypeUUID, Nullable: true},
		{Name: "nickname", Type: field.TypeString, Nullable: true},
		{Name: "user_pets", Type: field.TypeInt, Nullable: true},
		{Name: "user_team", Type: field.TypeInt, Unique: true, Nullable: true},
	}
	// PetTable holds the schema information for the "pet" table.
	PetTable = &schema.Table{
		Name:       "pet",
		Columns:    PetColumns,
		PrimaryKey: []*schema.Column{PetColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "pet_users_pets",
				Columns:    []*schema.Column{PetColumns[5]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:     "pet_users_team",
				Columns:    []*schema.Column{PetColumns[6]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
		Indexes: []*schema.Index{
			{
				Name:    "pet_name_user_pets",
				Unique:  false,
				Columns: []*schema.Column{PetColumns[2], PetColumns[5]},
			},
			{
				Name:    "pet_nickname",
				Unique:  true,
				Columns: []*schema.Column{PetColumns[4]},
			},
		},
	}
	// SpecsColumns holds the columns for the "specs" table.
	SpecsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
	}
	// SpecsTable holds the schema information for the "specs" table.
	SpecsTable = &schema.Table{
		Name:       "specs",
		Columns:    SpecsColumns,
		PrimaryKey: []*schema.Column{SpecsColumns[0]},
	}
	// TasksColumns holds the columns for the "tasks" table.
	TasksColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "priority", Type: field.TypeInt, Default: 1},
	}
	// TasksTable holds the schema information for the "tasks" table.
	TasksTable = &schema.Table{
		Name:       "tasks",
		Columns:    TasksColumns,
		PrimaryKey: []*schema.Column{TasksColumns[0]},
	}
	// UsersColumns holds the columns for the "users" table.
	UsersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "optional_int", Type: field.TypeInt, Nullable: true},
		{Name: "age", Type: field.TypeInt},
		{Name: "name", Type: field.TypeString},
		{Name: "last", Type: field.TypeString, Default: "unknown"},
		{Name: "nickname", Type: field.TypeString, Unique: true, Nullable: true},
		{Name: "address", Type: field.TypeString, Nullable: true},
		{Name: "phone", Type: field.TypeString, Unique: true, Nullable: true},
		{Name: "password", Type: field.TypeString, Nullable: true},
		{Name: "role", Type: field.TypeEnum, Enums: []string{"user", "admin", "free-user", "test user"}, Default: "user"},
		{Name: "employment", Type: field.TypeEnum, Enums: []string{"Full-Time", "Part-Time", "Contract"}, Default: "Full-Time"},
		{Name: "sso_cert", Type: field.TypeString, Nullable: true},
		{Name: "group_blocked", Type: field.TypeInt, Nullable: true},
		{Name: "user_spouse", Type: field.TypeInt, Unique: true, Nullable: true},
		{Name: "user_parent", Type: field.TypeInt, Nullable: true},
	}
	// UsersTable holds the schema information for the "users" table.
	UsersTable = &schema.Table{
		Name:       "users",
		Columns:    UsersColumns,
		PrimaryKey: []*schema.Column{UsersColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "users_groups_blocked",
				Columns:    []*schema.Column{UsersColumns[12]},
				RefColumns: []*schema.Column{GroupsColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:     "users_users_spouse",
				Columns:    []*schema.Column{UsersColumns[13]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:     "users_users_parent",
				Columns:    []*schema.Column{UsersColumns[14]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// SpecCardColumns holds the columns for the "spec_card" table.
	SpecCardColumns = []*schema.Column{
		{Name: "spec_id", Type: field.TypeInt},
		{Name: "card_id", Type: field.TypeInt},
	}
	// SpecCardTable holds the schema information for the "spec_card" table.
	SpecCardTable = &schema.Table{
		Name:       "spec_card",
		Columns:    SpecCardColumns,
		PrimaryKey: []*schema.Column{SpecCardColumns[0], SpecCardColumns[1]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "spec_card_spec_id",
				Columns:    []*schema.Column{SpecCardColumns[0]},
				RefColumns: []*schema.Column{SpecsColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "spec_card_card_id",
				Columns:    []*schema.Column{SpecCardColumns[1]},
				RefColumns: []*schema.Column{CardsColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// UserGroupsColumns holds the columns for the "user_groups" table.
	UserGroupsColumns = []*schema.Column{
		{Name: "user_id", Type: field.TypeInt},
		{Name: "group_id", Type: field.TypeInt},
	}
	// UserGroupsTable holds the schema information for the "user_groups" table.
	UserGroupsTable = &schema.Table{
		Name:       "user_groups",
		Columns:    UserGroupsColumns,
		PrimaryKey: []*schema.Column{UserGroupsColumns[0], UserGroupsColumns[1]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "user_groups_user_id",
				Columns:    []*schema.Column{UserGroupsColumns[0]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "user_groups_group_id",
				Columns:    []*schema.Column{UserGroupsColumns[1]},
				RefColumns: []*schema.Column{GroupsColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// UserFriendsColumns holds the columns for the "user_friends" table.
	UserFriendsColumns = []*schema.Column{
		{Name: "user_id", Type: field.TypeInt},
		{Name: "friend_id", Type: field.TypeInt},
	}
	// UserFriendsTable holds the schema information for the "user_friends" table.
	UserFriendsTable = &schema.Table{
		Name:       "user_friends",
		Columns:    UserFriendsColumns,
		PrimaryKey: []*schema.Column{UserFriendsColumns[0], UserFriendsColumns[1]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "user_friends_user_id",
				Columns:    []*schema.Column{UserFriendsColumns[0]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "user_friends_friend_id",
				Columns:    []*schema.Column{UserFriendsColumns[1]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// UserFollowingColumns holds the columns for the "user_following" table.
	UserFollowingColumns = []*schema.Column{
		{Name: "user_id", Type: field.TypeInt},
		{Name: "follower_id", Type: field.TypeInt},
	}
	// UserFollowingTable holds the schema information for the "user_following" table.
	UserFollowingTable = &schema.Table{
		Name:       "user_following",
		Columns:    UserFollowingColumns,
		PrimaryKey: []*schema.Column{UserFollowingColumns[0], UserFollowingColumns[1]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "user_following_user_id",
				Columns:    []*schema.Column{UserFollowingColumns[0]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "user_following_follower_id",
				Columns:    []*schema.Column{UserFollowingColumns[1]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		CardsTable,
		CommentsTable,
		FieldTypesTable,
		FilesTable,
		FileTypesTable,
		GoodsTable,
		GroupsTable,
		GroupInfosTable,
		ItemsTable,
		LicensesTable,
		NodesTable,
		PetTable,
		SpecsTable,
		TasksTable,
		UsersTable,
		SpecCardTable,
		UserGroupsTable,
		UserFriendsTable,
		UserFollowingTable,
	}
)

func init() {
	CardsTable.ForeignKeys[0].RefTable = UsersTable
	FieldTypesTable.ForeignKeys[0].RefTable = FilesTable
	FilesTable.ForeignKeys[0].RefTable = FileTypesTable
	FilesTable.ForeignKeys[1].RefTable = GroupsTable
	FilesTable.ForeignKeys[2].RefTable = UsersTable
	GroupsTable.ForeignKeys[0].RefTable = GroupInfosTable
	NodesTable.ForeignKeys[0].RefTable = NodesTable
	PetTable.ForeignKeys[0].RefTable = UsersTable
	PetTable.ForeignKeys[1].RefTable = UsersTable
	PetTable.Annotation = &entsql.Annotation{
		Table: "pet",
	}
	UsersTable.ForeignKeys[0].RefTable = GroupsTable
	UsersTable.ForeignKeys[1].RefTable = UsersTable
	UsersTable.ForeignKeys[2].RefTable = UsersTable
	SpecCardTable.ForeignKeys[0].RefTable = SpecsTable
	SpecCardTable.ForeignKeys[1].RefTable = CardsTable
	UserGroupsTable.ForeignKeys[0].RefTable = UsersTable
	UserGroupsTable.ForeignKeys[1].RefTable = GroupsTable
	UserFriendsTable.ForeignKeys[0].RefTable = UsersTable
	UserFriendsTable.ForeignKeys[1].RefTable = UsersTable
	UserFollowingTable.ForeignKeys[0].RefTable = UsersTable
	UserFollowingTable.ForeignKeys[1].RefTable = UsersTable
}
