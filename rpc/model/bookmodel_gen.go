// Code generated by goctl. DO NOT EDIT.

package model

import (
	"context"
	"database/sql"
	"fmt"
	"strings"

	"github.com/zeromicro/go-zero/core/stores/builder"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/core/stringx"
)

var (
	bookFieldNames          = builder.RawFieldNames(&Book{})
	bookRows                = strings.Join(bookFieldNames, ",")
	bookRowsExpectAutoSet   = strings.Join(stringx.Remove(bookFieldNames, "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), ",")
	bookRowsWithPlaceHolder = strings.Join(stringx.Remove(bookFieldNames, "`book`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), "=?,") + "=?"

	cacheBookBookPrefix = "cache:book:book:"
)

type (
	bookModel interface {
		Insert(ctx context.Context, data *Book) (sql.Result, error)
		FindOne(ctx context.Context, book string) (*Book, error)
		Update(ctx context.Context, data *Book) error
		Delete(ctx context.Context, book string) error
	}

	defaultBookModel struct {
		sqlc.CachedConn
		table string
	}

	Book struct {
		Book  string `db:"book"`  // book name
		Price int64  `db:"price"` // book price
	}
)

func newBookModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) *defaultBookModel {
	return &defaultBookModel{
		CachedConn: sqlc.NewConn(conn, c, opts...),
		table:      "`book`",
	}
}

func (m *defaultBookModel) Delete(ctx context.Context, book string) error {
	bookBookKey := fmt.Sprintf("%s%v", cacheBookBookPrefix, book)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where `book` = ?", m.table)
		return conn.ExecCtx(ctx, query, book)
	}, bookBookKey)
	return err
}

func (m *defaultBookModel) FindOne(ctx context.Context, book string) (*Book, error) {
	bookBookKey := fmt.Sprintf("%s%v", cacheBookBookPrefix, book)
	var resp Book
	err := m.QueryRowCtx(ctx, &resp, bookBookKey, func(ctx context.Context, conn sqlx.SqlConn, v any) error {
		query := fmt.Sprintf("select %s from %s where `book` = ? limit 1", bookRows, m.table)
		return conn.QueryRowCtx(ctx, v, query, book)
	})
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultBookModel) Insert(ctx context.Context, data *Book) (sql.Result, error) {
	bookBookKey := fmt.Sprintf("%s%v", cacheBookBookPrefix, data.Book)
	ret, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("insert into %s (%s) values (?, ?)", m.table, bookRowsExpectAutoSet)
		return conn.ExecCtx(ctx, query, data.Book, data.Price)
	}, bookBookKey)
	return ret, err
}

func (m *defaultBookModel) Update(ctx context.Context, data *Book) error {
	bookBookKey := fmt.Sprintf("%s%v", cacheBookBookPrefix, data.Book)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `book` = ?", m.table, bookRowsWithPlaceHolder)
		return conn.ExecCtx(ctx, query, data.Price, data.Book)
	}, bookBookKey)
	return err
}

func (m *defaultBookModel) formatPrimary(primary any) string {
	return fmt.Sprintf("%s%v", cacheBookBookPrefix, primary)
}

func (m *defaultBookModel) queryPrimary(ctx context.Context, conn sqlx.SqlConn, v, primary any) error {
	query := fmt.Sprintf("select %s from %s where `book` = ? limit 1", bookRows, m.table)
	return conn.QueryRowCtx(ctx, v, query, primary)
}

func (m *defaultBookModel) tableName() string {
	return m.table
}
