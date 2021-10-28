package comment

import (
	"errors"
	"fmt"

	"github.com/blog-service/src/datasources/mysql"
	"github.com/blog-service/src/domain/comment"
	dateutils "github.com/blog-service/src/utils/date"
	errorutils "github.com/blog-service/src/utils/errors"
	fileutils "github.com/blog-service/src/utils/file"
	"github.com/blog-service/src/utils/logger"
)

const (
	COMMENT_TBL_NAME              = "tbl_comments"
	COMMENT_TBL_COLUMNS           = "id,review,fk_post_id,active,deleted,created_at,updated_at,deleted_at"
	INSERT_COMMENT_LOC            = "./resources/domain/sql/common/Insert.sql"
	UPDATE_COMMENT_LOC            = "./resources/domain/sql/common/Update.sql"
	SELECT_COMMENT_BY_ID_LOC      = "./resources/domain/sql/common/SelectById.sql"
	SELECT_ALL_COMMENT_LOC        = "./resources/domain/sql/common/SelectAllRows.sql"
	SELECT_COMMENT_WITH_LIMIT_LOC = "./resources/domain/sql/common/SelectForPagination.sql"
	COUNT_COMMENT_ROWS_LOC        = "./resources/domain/sql/common/CountRows.sql"
)

var (
	ErrCommentNotFound = errors.New("comment not found")
)

type ICommentRepository interface {
	Save(comment.Comment) (*comment.Comment, error)
	Update(comment.Comment) (*comment.Comment, error)
	FindById(string) (*comment.Comment, error)
	Find(comment.CommentFilter) (*comment.Comment, error)
	FindAll(comment.CommentFilter) []comment.Comment
	FindAllWithPagination(comment.CommentListFilter) *comment.CommentPaginationDetails
	Delete(*comment.Comment) error
}

type commentRepository struct{}

func NewCommentRepository() ICommentRepository {
	return &commentRepository{}
}

func (c *commentRepository) Save(comment comment.Comment) (*comment.Comment, error) {
	str, err := fileutils.LoadResourceAsString(INSERT_COMMENT_LOC)
	if err != nil {
		logger.Error(fmt.Sprintf("cannot open file to this path: %s", INSERT_COMMENT_LOC), err)
		return nil, err
	}

	values := "?,?,?,?,?,?,?,?"
	query := fmt.Sprintf(str, COMMENT_TBL_NAME, COMMENT_TBL_COLUMNS, values)
	logger.Info(fmt.Sprintf("Query: %s", query))

	stmt, err := mysql.Client.Prepare(query)
	if err != nil {
		logger.Error(fmt.Sprintf("cannot prepare query: %s", query), err)
		return nil, errorutils.ErrInvalidQuery
	}
	defer stmt.Close()

	_, err = stmt.Exec(comment.Id,
		comment.Review,
		comment.PostId,
		comment.IsActive,
		comment.IsDeleted,
		comment.CreatedAt,
		comment.UpdatedAt,
		comment.DeletedAt)
	if err != nil {
		logger.Error("row insert failed", err)
		return nil, errorutils.ErrRowInsertFailed
	}

	return &comment, nil
}

func (c *commentRepository) Update(comment comment.Comment) (*comment.Comment, error) {
	str, err := fileutils.LoadResourceAsString(UPDATE_COMMENT_LOC)
	if err != nil {
		logger.Error(fmt.Sprintf("cannot open file to this path: %s", UPDATE_COMMENT_LOC), err)
		return nil, err
	}

	args := make([]interface{}, 0)
	var values string
	if comment.Review != "" && len(comment.Review) > 0 {
		values += "review=?,"
		args = append(args, comment.Review)
	}

	if comment.UpdatedAt != "" && len(comment.UpdatedAt) > 0 {
		values += "updated_at=?,"
		args = append(args, comment.UpdatedAt)
	}

	if comment.DeletedAt != "" && len(comment.DeletedAt) > 0 {
		values += "deleted_at=?,"
		args = append(args, comment.DeletedAt)
	}

	values += "active=?, deleted=?"
	args = append(args, comment.IsActive, comment.IsDeleted, comment.Id)

	query := fmt.Sprintf(str, COMMENT_TBL_NAME, values)
	logger.Info(fmt.Sprintf("Query: %s", query))

	stmt, err := mysql.Client.Prepare(query)
	if err != nil {
		return nil, errorutils.ErrInvalidQuery
	}
	defer stmt.Close()

	_, err = stmt.Exec(args...)
	if err != nil {
		logger.Error("row insert failed", err)
		return nil, errorutils.ErrRowInsertFailed
	}

	return &comment, nil
}

func (c *commentRepository) FindById(id string) (*comment.Comment, error) {
	str, err := fileutils.LoadResourceAsString(SELECT_COMMENT_BY_ID_LOC)
	if err != nil {
		logger.Error(fmt.Sprintf("cannot open file to this path: %s", SELECT_COMMENT_BY_ID_LOC), err)
		return nil, err
	}

	query := fmt.Sprintf(str, COMMENT_TBL_COLUMNS, COMMENT_TBL_NAME)
	logger.Info(fmt.Sprintf("Query: %s", query))

	stmt, err := mysql.Client.Prepare(query)
	if err != nil {
		return nil, errorutils.ErrInvalidQuery
	}
	defer stmt.Close()

	comment := comment.Comment{}
	result := stmt.QueryRow(id)
	if err := result.Scan(&comment.Id, &comment.Review, &comment.PostId, &comment.IsActive, &comment.IsDeleted, &comment.CreatedAt, &comment.UpdatedAt, &comment.DeletedAt); err != nil {
		logger.Error("Comment not found", err)
		return nil, ErrCommentNotFound
	}

	return &comment, nil
}

func (c *commentRepository) Find(filter comment.CommentFilter) (*comment.Comment, error) {
	str, err := fileutils.LoadResourceAsString(SELECT_ALL_COMMENT_LOC)
	if err != nil {
		logger.Error(fmt.Sprintf("cannot open file to this path: %s", SELECT_ALL_COMMENT_LOC), err)
		return nil, err
	}

	args := make([]interface{}, 0)
	var condition string = "1 = 1 "
	if filter.Id != "" && len(filter.Id) > 0 {
		condition += "AND id = ? "
		args = append(args, filter.Id)
	}
	if filter.PostId != "" && len(filter.PostId) > 0 {
		condition += "AND fk_post_id = ? "
		args = append(args, filter.PostId)
	}
	if filter.Active != nil {
		condition += "AND active = ? "
		args = append(args, filter.Active)
	}
	if filter.Deleted != nil {
		condition += "AND deleted = ? "
		args = append(args, filter.Deleted)
	}

	query := fmt.Sprintf(str, COMMENT_TBL_COLUMNS, COMMENT_TBL_NAME, condition)

	stmt, err := mysql.Client.Prepare(query)
	if err != nil {
		return nil, errorutils.ErrInvalidQuery
	}
	defer stmt.Close()

	comment := comment.Comment{}
	result := stmt.QueryRow(args...)
	if err := result.Scan(&comment.Id, &comment.Review, &comment.PostId, &comment.IsActive, &comment.IsDeleted, &comment.CreatedAt, &comment.UpdatedAt, &comment.DeletedAt); err != nil {
		logger.Error("Comment not found", err)
		return nil, ErrCommentNotFound
	}

	return &comment, nil
}

func (c *commentRepository) FindAll(filter comment.CommentFilter) []comment.Comment {
	comments := make([]comment.Comment, 0)
	str, err := fileutils.LoadResourceAsString(SELECT_ALL_COMMENT_LOC)
	if err != nil {
		logger.Error(fmt.Sprintf("cannot open file to this path: %s", SELECT_ALL_COMMENT_LOC), err)
		return comments
	}

	args := make([]interface{}, 0)
	var condition string = "1 = 1 "
	if filter.Id != "" && len(filter.Id) > 0 {
		condition += "AND id = ? "
		args = append(args, filter.Id)
	}
	if filter.PostId != "" && len(filter.PostId) > 0 {
		condition += "AND fk_post_id = ? "
		args = append(args, filter.PostId)
	}
	if filter.Active != nil {
		condition += "AND active = ? "
		args = append(args, filter.Active)
	}
	if filter.Deleted != nil {
		condition += "AND deleted = ? "
		args = append(args, filter.Deleted)
	}

	query := fmt.Sprintf(str, COMMENT_TBL_COLUMNS, COMMENT_TBL_NAME, condition)
	stmt, err := mysql.Client.Prepare(query)
	if err != nil {
		logger.Error("invalid query", err)
		return comments
	}
	defer stmt.Close()

	rows, err := stmt.Query(args...)
	if err != nil {
		logger.Debug("No data present")
		return comments
	}
	defer rows.Close()

	for rows.Next() {
		comment := comment.Comment{}
		rows.Scan(&comment.Id, &comment.Review, &comment.PostId, &comment.IsActive, &comment.IsDeleted, &comment.CreatedAt, &comment.UpdatedAt, &comment.DeletedAt)
		comments = append(comments, comment)
	}

	return comments
}

func (c *commentRepository) FindAllWithPagination(filter comment.CommentListFilter) *comment.CommentPaginationDetails {
	result := &comment.CommentPaginationDetails{
		Data: []comment.Comment{},
	}
	selectAllQueryStr, err := fileutils.LoadResourceAsString(SELECT_COMMENT_WITH_LIMIT_LOC)
	if err != nil {
		logger.Error(fmt.Sprintf("cannot open file to this path: %s", SELECT_COMMENT_WITH_LIMIT_LOC), err)
		return result
	}

	countRowsQueryStr, err := fileutils.LoadResourceAsString(COUNT_COMMENT_ROWS_LOC)
	if err != nil {
		logger.Error(fmt.Sprintf("cannot open file to this path: %s", COUNT_COMMENT_ROWS_LOC), err)
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
	if filter.Filter.PostId != "" && len(filter.Filter.PostId) > 0 {
		condition += "AND fk_post_id = ? "
		args = append(args, filter.Filter.PostId)
	}

	offset := (filter.Page - 1) * filter.Limit
	selectAllQuery := fmt.Sprintf(selectAllQueryStr, COMMENT_TBL_COLUMNS, COMMENT_TBL_NAME, condition, filter.SortBy, filter.Sort, filter.Limit, offset)
	countRowsQuery := fmt.Sprintf(countRowsQueryStr, "id", COMMENT_TBL_NAME, condition)
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
		logger.Error("No any comment record in database", err)
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
		comment := comment.Comment{}
		rows.Scan(&comment.Id, &comment.Review, &comment.PostId, &comment.IsActive, &comment.IsDeleted, &comment.CreatedAt, &comment.UpdatedAt, &comment.DeletedAt)
		result.Data = append(result.Data, comment)
	}
	result.Size = len(result.Data)

	return result
}

func (c *commentRepository) Delete(request *comment.Comment) error {
	request.IsActive = false
	request.IsDeleted = true
	request.DeletedAt = dateutils.GetNow().String()
	request.UpdatedAt = request.DeletedAt

	_, err := c.Update(*request)
	if err != nil {
		return err
	}
	return nil
}
