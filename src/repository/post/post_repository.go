package post

import (
	"errors"
	"fmt"

	"github.com/blog-service/src/datasources/mysql"
	"github.com/blog-service/src/domain/post"
	fileutils "github.com/blog-service/src/utils/file"
	"github.com/blog-service/src/utils/logger"
)

const (
	POST_TBL_NAME              = "tbl_posts"
	POST_TBL_COLUMNS           = "id,title,description,active,deleted,created_at,updated_at,deleted_at"
	INSERT_POST_LOC            = "./resources/domain/sql/common/Insert.sql"
	UPDATE_POST_LOC            = "./resources/domain/sql/common/Update.sql"
	SELECT_POST_BY_ID_LOC      = "./resources/domain/sql/common/SelectById.sql"
	SELECT_ALL_POST_LOC        = "./resources/domain/sql/common/SelectAllRows.sql"
	SELECT_POST_WITH_LIMIT_LOC = "./resources/domain/sql/common/SelectForPagination.sql"
	COUNT_POST_ROWS_LOC        = "./resources/domain/sql/common/CountRows.sql"
)

var (
	ErrRowInsertFailed = errors.New("cannot insert post at this moment")
	ErrInvalidQuery    = errors.New("invalid query")
	ErrPostNotFound    = errors.New("post not found")
)

type IPostRepository interface {
	Save(*post.Post) (*post.Post, error)
	Update(*post.Post) (*post.Post, error)
	FindById(string) (*post.Post, error)
	Find(post.PostFilter) (*post.Post, error)
	FindAll() []post.Post
	FindAllWithPagination(post.PostListFilter) *post.PostPaginationDetails
	Delete(string) error
}

type postRepository struct{}

func NewPostRepository() IPostRepository {
	return &postRepository{}
}

func (p *postRepository) Save(post *post.Post) (*post.Post, error) {
	str, err := fileutils.LoadResourceAsString(INSERT_POST_LOC)
	if err != nil {
		logger.Error(fmt.Sprintf("cannot open file to this path: %s", INSERT_POST_LOC), err)
		return nil, err
	}

	values := "?,?,?,?,?,?,?,?"
	query := fmt.Sprintf(str, POST_TBL_NAME, POST_TBL_COLUMNS, values)
	logger.Info(fmt.Sprintf("Query: %s", query))

	stmt, err := mysql.Client.Prepare(query)
	if err != nil {
		logger.Error(fmt.Sprintf("cannot prepare query: %s", query), err)
		return nil, ErrInvalidQuery
	}
	defer stmt.Close()

	_, err = stmt.Exec(post.Id,
		post.Title,
		post.Description,
		post.IsActive,
		post.IsDeleted,
		post.CreatedAt,
		post.UpdatedAt,
		post.DeletedAt)
	if err != nil {
		logger.Error("row insert failed", err)
		return nil, ErrRowInsertFailed
	}

	return post, nil
}

func (p *postRepository) Update(post *post.Post) (*post.Post, error) {
	str, err := fileutils.LoadResourceAsString(UPDATE_POST_LOC)
	if err != nil {
		logger.Error(fmt.Sprintf("cannot open file to this path: %s", INSERT_POST_LOC), err)
		return nil, err
	}

	args := make([]interface{}, 0)
	var values string
	if post.Title != "" && len(post.Title) > 0 {
		values += "title=?,"
		args = append(args, post.Title)
	}

	if post.Description != "" && len(post.Description) > 0 {
		values += "description=?,"
		args = append(args, post.Description)
	}

	if post.UpdatedAt != "" && len(post.UpdatedAt) > 0 {
		values += "updated_at=?,"
		args = append(args, post.UpdatedAt)
	}

	values += "active=?, deleted=?"
	args = append(args, post.IsActive, post.IsDeleted, post.Id)

	query := fmt.Sprintf(str, POST_TBL_NAME, values)
	logger.Info(fmt.Sprintf("Query: %s", query))

	stmt, err := mysql.Client.Prepare(query)
	if err != nil {
		return nil, ErrInvalidQuery
	}
	defer stmt.Close()

	_, err = stmt.Exec(args...)
	if err != nil {
		logger.Error("row insert failed", err)
		return nil, ErrRowInsertFailed
	}

	return post, nil
}

func (p *postRepository) FindById(id string) (*post.Post, error) {
	str, err := fileutils.LoadResourceAsString(SELECT_POST_BY_ID_LOC)
	if err != nil {
		logger.Error(fmt.Sprintf("cannot open file to this path: %s", SELECT_POST_BY_ID_LOC), err)
		return nil, err
	}

	query := fmt.Sprintf(str, POST_TBL_COLUMNS, POST_TBL_NAME)
	logger.Info(fmt.Sprintf("Query: %s", query))

	stmt, err := mysql.Client.Prepare(query)
	if err != nil {
		return nil, ErrInvalidQuery
	}
	defer stmt.Close()

	post := post.Post{}
	result := stmt.QueryRow(id)
	if err := result.Scan(&post.Id, &post.Title, &post.Description, &post.IsActive, &post.IsDeleted, &post.CreatedAt, &post.UpdatedAt, &post.DeletedAt); err != nil {
		logger.Error("Post not found", err)
		return nil, ErrPostNotFound
	}

	return &post, nil
}

func (p *postRepository) Find(filter post.PostFilter) (*post.Post, error) {
	str, err := fileutils.LoadResourceAsString(SELECT_ALL_POST_LOC)
	if err != nil {
		logger.Error(fmt.Sprintf("cannot open file to this path: %s", INSERT_POST_LOC), err)
		return nil, err
	}

	args := make([]interface{}, 0)
	var condition string = "1 = 1 "
	if filter.Id != "" && len(filter.Id) > 0 {
		condition += "AND id = ? "
		args = append(args, filter.Id)
	}
	if filter.Title != "" && len(filter.Title) > 0 {
		condition += "AND title = ? "
		args = append(args, filter.Title)
	}
	if filter.Active != nil {
		condition += "AND active = ? "
		args = append(args, filter.Active)
	}
	if filter.Deleted != nil {
		condition += "AND deleted = ? "
		args = append(args, filter.Deleted)
	}

	query := fmt.Sprintf(str, POST_TBL_COLUMNS, POST_TBL_NAME, condition)

	stmt, err := mysql.Client.Prepare(query)
	if err != nil {
		return nil, ErrInvalidQuery
	}
	defer stmt.Close()

	post := post.Post{}
	result := stmt.QueryRow(args...)
	if err := result.Scan(&post.Id, &post.Title, &post.Description, &post.IsActive, &post.IsDeleted, &post.CreatedAt, &post.UpdatedAt, &post.DeletedAt); err != nil {
		logger.Error("Post not found", err)
		return nil, ErrPostNotFound
	}

	return &post, nil
}

func (p *postRepository) FindAll() []post.Post {
	posts := make([]post.Post, 0)
	str, err := fileutils.LoadResourceAsString(SELECT_ALL_POST_LOC)
	if err != nil {
		logger.Error(fmt.Sprintf("cannot open file to this path: %s", INSERT_POST_LOC), err)
		return posts
	}
	query := fmt.Sprintf(str, POST_TBL_COLUMNS, POST_TBL_NAME, "1=1")
	stmt, err := mysql.Client.Prepare(query)
	if err != nil {
		logger.Error("invalid query", err)
		return posts
	}
	defer stmt.Close()

	rows, err := stmt.Query()
	if err != nil {
		logger.Debug("No data present")
		return posts
	}
	defer rows.Close()

	for rows.Next() {
		post := post.Post{}
		rows.Scan(&post.Id, &post.Title, &post.Description, &post.IsActive, &post.IsDeleted, &post.CreatedAt, &post.UpdatedAt, &post.DeletedAt)
		posts = append(posts, post)
	}

	return posts
}

func (p *postRepository) FindAllWithPagination(filter post.PostListFilter) *post.PostPaginationDetails {
	result := &post.PostPaginationDetails{}
	selectAllQueryStr, err := fileutils.LoadResourceAsString(SELECT_POST_WITH_LIMIT_LOC)
	if err != nil {
		logger.Error(fmt.Sprintf("cannot open file to this path: %s", INSERT_POST_LOC), err)
		return result
	}

	countRowsQueryStr, err := fileutils.LoadResourceAsString(COUNT_POST_ROWS_LOC)
	if err != nil {
		logger.Error(fmt.Sprintf("cannot open file to this path: %s", INSERT_POST_LOC), err)
		return result
	}

	condition := "1=1 "
	args := make([]interface{}, 0)
	if filter.Filter.Active != nil {
		condition += "AND active = ? "
		args = append(args, filter.Filter.Active)
	}
	if filter.Filter.Deleted != nil {
		condition += "AND deleted = ? "
		args = append(args, filter.Filter.Deleted)
	}
	if filter.CreatedAt != "" {
		condition += "AND created_at = ? "
		args = append(args, filter.CreatedAt)
	}

	offset := (filter.Page - 1) * filter.Limit
	selectAllQuery := fmt.Sprintf(selectAllQueryStr, POST_TBL_COLUMNS, POST_TBL_NAME, condition, filter.SortBy, filter.Sort, filter.Limit, offset)
	countRowsQuery := fmt.Sprintf(countRowsQueryStr, "id", POST_TBL_NAME, condition)
	logger.Info(fmt.Sprintf("SelectQuery: %s \n CountQuery: %s\n", selectAllQuery, countRowsQuery))

	selectStmt, err := mysql.Client.Prepare(selectAllQuery)
	if err != nil {
		logger.Error("invalid query", err)
		return result
	}
	defer selectStmt.Close()

	countRowsStmt, err := mysql.Client.Prepare(countRowsQuery)
	if err != nil {
		logger.Error("invalid query", err)
		return result
	}
	defer countRowsStmt.Close()

	var count int
	countResult := countRowsStmt.QueryRow(args...)
	if err := countResult.Scan(&count); err != nil {
		logger.Error("No any post record in database", err)
		return result
	}

	rows, err := selectStmt.Query(args...)
	if err != nil {
		logger.Debug("No data present")
		return result
	}
	defer rows.Close()

	result.Page = int(filter.Page)
	result.Total = count
	for rows.Next() {
		post := post.Post{}
		rows.Scan(&post.Id, &post.Title, &post.Description, &post.IsActive, &post.IsDeleted, &post.CreatedAt, &post.UpdatedAt, &post.DeletedAt)
		result.Data = append(result.Data, post)
	}
	result.Size = len(result.Data)

	return result
}

func (p *postRepository) Delete(id string) error {
	return nil
}
