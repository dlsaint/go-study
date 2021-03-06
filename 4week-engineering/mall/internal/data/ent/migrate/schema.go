// Code generated by entc, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// CommentsColumns holds the columns for the "comments" table.
	CommentsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt64, Increment: true},
		{Name: "name", Type: field.TypeString},
		{Name: "content", Type: field.TypeString},
		{Name: "created_at", Type: field.TypeTime, SchemaType: map[string]string{"mysql": "datetime"}},
		{Name: "updated_at", Type: field.TypeTime, SchemaType: map[string]string{"mysql": "datetime"}},
		{Name: "product_comments", Type: field.TypeInt64, Nullable: true},
	}
	// CommentsTable holds the schema information for the "comments" table.
	CommentsTable = &schema.Table{
		Name:       "comments",
		Columns:    CommentsColumns,
		PrimaryKey: []*schema.Column{CommentsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "comments_products_comments",
				Columns:    []*schema.Column{CommentsColumns[5]},
				RefColumns: []*schema.Column{ProductsColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// ProductsColumns holds the columns for the "products" table.
	ProductsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt64, Increment: true},
		{Name: "name", Type: field.TypeString},
		{Name: "price", Type: field.TypeInt64},
		{Name: "created_at", Type: field.TypeTime, SchemaType: map[string]string{"mysql": "datetime"}},
		{Name: "updated_at", Type: field.TypeTime, SchemaType: map[string]string{"mysql": "datetime"}},
	}
	// ProductsTable holds the schema information for the "products" table.
	ProductsTable = &schema.Table{
		Name:       "products",
		Columns:    ProductsColumns,
		PrimaryKey: []*schema.Column{ProductsColumns[0]},
	}
	// TagsColumns holds the columns for the "tags" table.
	TagsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt64, Increment: true},
		{Name: "slug", Type: field.TypeString},
		{Name: "name", Type: field.TypeString},
		{Name: "created_at", Type: field.TypeTime, SchemaType: map[string]string{"mysql": "datetime"}},
		{Name: "updated_at", Type: field.TypeTime, SchemaType: map[string]string{"mysql": "datetime"}},
	}
	// TagsTable holds the schema information for the "tags" table.
	TagsTable = &schema.Table{
		Name:       "tags",
		Columns:    TagsColumns,
		PrimaryKey: []*schema.Column{TagsColumns[0]},
	}
	// TagPostsColumns holds the columns for the "tag_posts" table.
	TagPostsColumns = []*schema.Column{
		{Name: "tag_id", Type: field.TypeInt64},
		{Name: "product_id", Type: field.TypeInt64},
	}
	// TagPostsTable holds the schema information for the "tag_posts" table.
	TagPostsTable = &schema.Table{
		Name:       "tag_posts",
		Columns:    TagPostsColumns,
		PrimaryKey: []*schema.Column{TagPostsColumns[0], TagPostsColumns[1]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "tag_posts_tag_id",
				Columns:    []*schema.Column{TagPostsColumns[0]},
				RefColumns: []*schema.Column{TagsColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "tag_posts_product_id",
				Columns:    []*schema.Column{TagPostsColumns[1]},
				RefColumns: []*schema.Column{ProductsColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		CommentsTable,
		ProductsTable,
		TagsTable,
		TagPostsTable,
	}
)

func init() {
	CommentsTable.ForeignKeys[0].RefTable = ProductsTable
	TagPostsTable.ForeignKeys[0].RefTable = TagsTable
	TagPostsTable.ForeignKeys[1].RefTable = ProductsTable
}
