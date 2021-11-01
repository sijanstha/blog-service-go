package user

import (
	"errors"
	"fmt"

	"github.com/blog-service/src/datasources/mysql"
	"github.com/blog-service/src/domain/user"
	dateutils "github.com/blog-service/src/utils/date"
	errorutils "github.com/blog-service/src/utils/errors"
	fileutils "github.com/blog-service/src/utils/file"
	"github.com/blog-service/src/utils/logger"
	sqlutils "github.com/blog-service/src/utils/sql"
)

const (
	USER_TABLE_DEFINITION      = "./resources/domain/table_definition/user/UserTableDefinition.json"
	INSERT_USER_LOC            = "./resources/domain/sql/common/Insert.sql"
	UPDATE_USER_LOC            = "./resources/domain/sql/common/Update.sql"
	SELECT_USER_BY_ID_LOC      = "./resources/domain/sql/common/SelectById.sql"
	SELECT_ALL_USER_LOC        = "./resources/domain/sql/common/SelectAllRows.sql"
	SELECT_USER_WITH_LIMIT_LOC = "./resources/domain/sql/common/SelectForPagination.sql"
	COUNT_USER_ROWS_LOC        = "./resources/domain/sql/common/CountRows.sql"
)

var (
	ErrUserNotFound = errors.New("user not found")
)

type IUserRepository interface {
	Save(*user.User) (*user.UserDomain, error)
	Update(*user.User) (*user.UserDomain, error)
	FindById(string) (*user.UserDomain, error)
	Find(user.UserFilter) (*user.UserDomain, error)
	FindAll() user.UserDomainList
	FindAllWithPagination(user.UserListFilter) *user.UserPaginationDetails
	Delete(*user.User) error
}

type userRepository struct{}

func NewUserRepository() IUserRepository {
	return &userRepository{}
}

func (repo *userRepository) Save(user *user.User) (*user.UserDomain, error) {
	str, err := fileutils.LoadResourceAsString(INSERT_USER_LOC)
	if err != nil {
		logger.Error(fmt.Sprintf("cannot open file to this path: %s", INSERT_USER_LOC), err)
		return nil, err
	}

	tableMetadata, err := fileutils.GetTableInformation(USER_TABLE_DEFINITION)
	if err != nil {
		logger.Error(fmt.Sprintf("cannot open file to this path: %s", USER_TABLE_DEFINITION), err)
		return nil, err
	}

	tblName, columns, values := sqlutils.GetTableMetadata(tableMetadata)
	query := fmt.Sprintf(str, tblName, columns, values)
	logger.Info(fmt.Sprintf("Query: %s", query))

	stmt, err := mysql.Client.Prepare(query)
	if err != nil {
		logger.Error(fmt.Sprintf("cannot prepare query: %s", query), err)
		return nil, errorutils.ErrInvalidQuery
	}
	defer stmt.Close()

	_, err = stmt.Exec(user.Id,
		user.FirstName,
		user.LastName,
		user.Email,
		user.PasswordHash,
		user.ProfilePictureUrl,
		user.RoleId,
		user.IsActive,
		user.IsDeleted,
		user.CreatedAt,
		user.UpdatedAt,
		user.DeletedAt)
	if err != nil {
		logger.Error("row insert failed", err)
		return nil, errorutils.ErrRowInsertFailed
	}

	return user.ToUserDomain(), nil
}

func (repo *userRepository) Update(user *user.User) (*user.UserDomain, error) {
	str, err := fileutils.LoadResourceAsString(UPDATE_USER_LOC)
	if err != nil {
		logger.Error(fmt.Sprintf("cannot open file to this path: %s", UPDATE_USER_LOC), err)
		return nil, err
	}

	tableMetadata, err := fileutils.GetTableInformation(USER_TABLE_DEFINITION)
	if err != nil {
		logger.Error(fmt.Sprintf("cannot open file to this path: %s", USER_TABLE_DEFINITION), err)
		return nil, err
	}

	args := make([]interface{}, 0)
	var values string
	if user.FirstName != "" && len(user.FirstName) > 0 {
		values += "first_name=?,"
		args = append(args, user.FirstName)
	}

	if user.LastName != "" && len(user.LastName) > 0 {
		values += "last_name=?,"
		args = append(args, user.LastName)
	}

	if user.PasswordHash != "" && len(user.PasswordHash) > 0 {
		values += "password_hash=?,"
		args = append(args, user.PasswordHash)
	}

	if user.RoleId != "" && len(user.RoleId) > 0 {
		values += "fk_role_id=?,"
		args = append(args, user.RoleId)
	}

	if user.UpdatedAt != "" && len(user.UpdatedAt) > 0 {
		values += "updated_at=?,"
		args = append(args, user.UpdatedAt)
	}

	if user.DeletedAt != "" && len(user.DeletedAt) > 0 {
		values += "deleted_at=?,"
		args = append(args, user.DeletedAt)
	}

	values += "active=?, deleted=?"
	args = append(args, user.IsActive, user.IsDeleted, user.Id)

	query := fmt.Sprintf(str, sqlutils.GetTableName(tableMetadata), values)
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

	return user.ToUserDomain(), nil
}

func (repo *userRepository) FindById(id string) (*user.UserDomain, error) {
	str, err := fileutils.LoadResourceAsString(SELECT_USER_BY_ID_LOC)
	if err != nil {
		logger.Error(fmt.Sprintf("cannot open file to this path: %s", SELECT_USER_BY_ID_LOC), err)
		return nil, err
	}

	tableMetadata, err := fileutils.GetTableInformation(USER_TABLE_DEFINITION)
	if err != nil {
		logger.Error(fmt.Sprintf("cannot open file to this path: %s", USER_TABLE_DEFINITION), err)
		return nil, err
	}

	tblName, columns, _ := sqlutils.GetTableMetadata(tableMetadata)

	query := fmt.Sprintf(str, columns, tblName)
	logger.Info(fmt.Sprintf("Query: %s", query))

	stmt, err := mysql.Client.Prepare(query)
	if err != nil {
		return nil, errorutils.ErrInvalidQuery
	}
	defer stmt.Close()

	user := user.User{}
	result := stmt.QueryRow(id)
	if err := result.Scan(&user.Id,
		&user.FirstName,
		&user.LastName,
		&user.Email,
		&user.PasswordHash,
		&user.ProfilePictureUrl,
		&user.RoleId,
		&user.IsActive,
		&user.IsDeleted,
		&user.CreatedAt,
		&user.UpdatedAt,
		&user.DeletedAt); err != nil {
		logger.Error("user not found", err)
		return nil, ErrUserNotFound
	}

	return user.ToUserDomain(), nil
}

func (repo *userRepository) Find(filter user.UserFilter) (*user.UserDomain, error) {
	str, err := fileutils.LoadResourceAsString(SELECT_ALL_USER_LOC)
	if err != nil {
		logger.Error(fmt.Sprintf("cannot open file to this path: %s", SELECT_ALL_USER_LOC), err)
		return nil, err
	}

	tableMetadata, err := fileutils.GetTableInformation(USER_TABLE_DEFINITION)
	if err != nil {
		logger.Error(fmt.Sprintf("cannot open file to this path: %s", USER_TABLE_DEFINITION), err)
		return nil, err
	}

	args := make([]interface{}, 0)
	var condition string = "1 = 1 "
	if filter.Id != "" && len(filter.Id) > 0 {
		condition += "AND id = ? "
		args = append(args, filter.Id)
	}
	if filter.Email != "" && len(filter.Email) > 0 {
		condition += "AND email = ? "
		args = append(args, filter.Email)
	}
	if filter.Password != "" && len(filter.Password) > 0 {
		condition += "AND password_hash = ? "
		args = append(args, filter.Password)
	}
	if filter.Active != nil {
		condition += "AND active = ? "
		args = append(args, filter.Active)
	}
	if filter.Deleted != nil {
		condition += "AND deleted = ? "
		args = append(args, filter.Deleted)
	}

	tblName, columns, _ := sqlutils.GetTableMetadata(tableMetadata)
	query := fmt.Sprintf(str, columns, tblName, condition)

	stmt, err := mysql.Client.Prepare(query)
	if err != nil {
		return nil, errorutils.ErrInvalidQuery
	}
	defer stmt.Close()

	user := user.User{}
	result := stmt.QueryRow(args...)
	if err := result.Scan(&user.Id,
		&user.FirstName,
		&user.LastName,
		&user.Email,
		&user.PasswordHash,
		&user.ProfilePictureUrl,
		&user.RoleId,
		&user.IsActive,
		&user.IsDeleted,
		&user.CreatedAt,
		&user.UpdatedAt,
		&user.DeletedAt); err != nil {
		logger.Error("user not found", err)
		return nil, ErrUserNotFound
	}

	return user.ToUserDomain(), nil
}

func (repo *userRepository) FindAll() user.UserDomainList {
	users := make(user.UserList, 0)
	str, err := fileutils.LoadResourceAsString(SELECT_ALL_USER_LOC)
	if err != nil {
		logger.Error(fmt.Sprintf("cannot open file to this path: %s", SELECT_ALL_USER_LOC), err)
		return user.UserDomainList{}
	}

	tableMetadata, err := fileutils.GetTableInformation(USER_TABLE_DEFINITION)
	if err != nil {
		logger.Error(fmt.Sprintf("cannot open file to this path: %s", USER_TABLE_DEFINITION), err)
		return user.UserDomainList{}
	}

	tblName, columns, _ := sqlutils.GetTableMetadata(tableMetadata)
	query := fmt.Sprintf(str, columns, tblName, "1=1")
	stmt, err := mysql.Client.Prepare(query)
	if err != nil {
		logger.Error("invalid query", err)
		return user.UserDomainList{}
	}
	defer stmt.Close()

	rows, err := stmt.Query()
	if err != nil {
		logger.Debug("No data present")
		return user.UserDomainList{}
	}
	defer rows.Close()

	for rows.Next() {
		user := user.User{}
		rows.Scan(&user.Id,
			&user.FirstName,
			&user.LastName,
			&user.Email,
			&user.PasswordHash,
			&user.ProfilePictureUrl,
			&user.RoleId,
			&user.IsActive,
			&user.IsDeleted,
			&user.CreatedAt,
			&user.UpdatedAt,
			&user.DeletedAt)
		users = append(users, user)
	}

	return users.ToUserDomainList()
}

func (repo *userRepository) FindAllWithPagination(filter user.UserListFilter) *user.UserPaginationDetails {
	result := &user.UserPaginationDetails{}
	var users user.UserList
	selectAllQueryStr, err := fileutils.LoadResourceAsString(SELECT_USER_WITH_LIMIT_LOC)
	if err != nil {
		logger.Error(fmt.Sprintf("cannot open file to this path: %s", SELECT_USER_WITH_LIMIT_LOC), err)
		return result
	}

	countRowsQueryStr, err := fileutils.LoadResourceAsString(COUNT_USER_ROWS_LOC)
	if err != nil {
		logger.Error(fmt.Sprintf("cannot open file to this path: %s", COUNT_USER_ROWS_LOC), err)
		return result
	}

	tableMetadata, err := fileutils.GetTableInformation(USER_TABLE_DEFINITION)
	if err != nil {
		logger.Error(fmt.Sprintf("cannot open file to this path: %s", USER_TABLE_DEFINITION), err)
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
	tblName, columns, _ := sqlutils.GetTableMetadata(tableMetadata)
	selectAllQuery := fmt.Sprintf(selectAllQueryStr, columns, tblName, condition, filter.SortBy, filter.Sort, filter.Limit, offset)
	countRowsQuery := fmt.Sprintf(countRowsQueryStr, "id", tblName, condition)
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
		user := user.User{}
		rows.Scan(
			&user.Id,
			&user.FirstName,
			&user.LastName,
			&user.Email,
			&user.PasswordHash,
			&user.ProfilePictureUrl,
			&user.RoleId,
			&user.IsActive,
			&user.IsDeleted,
			&user.CreatedAt,
			&user.UpdatedAt,
			&user.DeletedAt,
		)
		users = append(users, user)
	}
	result.Size = len(result.Data)
	result.Data = users.ToUserDomainList()
	return result
}

func (repo *userRepository) Delete(request *user.User) error {
	request.IsActive = false
	request.IsDeleted = true
	request.DeletedAt = dateutils.GetNow().String()
	request.UpdatedAt = request.DeletedAt

	_, err := repo.Update(request)
	if err != nil {
		return err
	}
	return nil
}
