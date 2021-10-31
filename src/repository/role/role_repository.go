package role

import (
	"errors"
	"fmt"

	"github.com/blog-service/src/datasources/mysql"
	"github.com/blog-service/src/domain/role"
	dateutils "github.com/blog-service/src/utils/date"
	errorutils "github.com/blog-service/src/utils/errors"
	fileutils "github.com/blog-service/src/utils/file"
	"github.com/blog-service/src/utils/logger"
)

const (
	ROLE_TABLE_NAME       = "tbl_role"
	ROLE_TBL_COLUMNS      = "id,role_name,display_name,active,deleted,created_at,updated_at,deleted_at"
	INSERT_ROLE_LOC       = "./resources/domain/sql/common/Insert.sql"
	UPDATE_ROLE_LOC       = "./resources/domain/sql/common/Update.sql"
	SELECT_ROLE_BY_ID_LOC = "./resources/domain/sql/common/SelectById.sql"
	SELECT_ALL_ROLE_LOC   = "./resources/domain/sql/common/SelectAllRows.sql"
)

var (
	ErrRoleNotFound          = errors.New("role not found")
	ErrRoleNameAlreadyExists = errors.New("role name already exists")
)

type IRoleRepository interface {
	Save(*role.Role) (*role.Role, error)
	Update(*role.Role) (*role.Role, error)
	FindById(string) (*role.Role, error)
	Find(role.RoleFilter) (*role.Role, error)
	FindAll() []role.Role
	Delete(*role.Role) error
}

type roleRepository struct{}

func NewRoleRepository() IRoleRepository {
	return &roleRepository{}
}

func (repo *roleRepository) Save(role *role.Role) (*role.Role, error) {
	str, err := fileutils.LoadResourceAsString(INSERT_ROLE_LOC)
	if err != nil {
		logger.Error(fmt.Sprintf("cannot open file to this path: %s", INSERT_ROLE_LOC), err)
		return nil, err
	}

	values := "?,?,?,?,?,?,?,?"
	query := fmt.Sprintf(str, ROLE_TABLE_NAME, ROLE_TBL_COLUMNS, values)
	logger.Info(fmt.Sprintf("Query: %s", query))

	stmt, err := mysql.Client.Prepare(query)
	if err != nil {
		logger.Error(fmt.Sprintf("cannot prepare query: %s", query), err)
		return nil, errorutils.ErrInvalidQuery
	}
	defer stmt.Close()

	_, err = stmt.Exec(role.Id,
		role.RoleName,
		role.DisplayName,
		role.IsActive,
		role.IsDeleted,
		role.CreatedAt,
		role.UpdatedAt,
		role.DeletedAt)
	if err != nil {
		logger.Error("row insert failed", err)
		return nil, errorutils.ErrRowInsertFailed
	}

	return role, nil
}

func (repo *roleRepository) Update(role *role.Role) (*role.Role, error) {
	str, err := fileutils.LoadResourceAsString(UPDATE_ROLE_LOC)
	if err != nil {
		logger.Error(fmt.Sprintf("cannot open file to this path: %s", UPDATE_ROLE_LOC), err)
		return nil, err
	}

	args := make([]interface{}, 0)
	var values string
	if role.RoleName != "" && len(role.RoleName) > 0 {
		values += "role_name=?,"
		args = append(args, role.RoleName)
	}

	if role.DisplayName != "" && len(role.DisplayName) > 0 {
		values += "display_name=?,"
		args = append(args, role.DisplayName)
	}

	if role.UpdatedAt != "" && len(role.UpdatedAt) > 0 {
		values += "updated_at=?,"
		args = append(args, role.UpdatedAt)
	}

	if role.DeletedAt != "" && len(role.DeletedAt) > 0 {
		values += "deleted_at=?,"
		args = append(args, role.DeletedAt)
	}

	values += "active=?, deleted=?"
	args = append(args, role.IsActive, role.IsDeleted, role.Id)

	query := fmt.Sprintf(str, ROLE_TABLE_NAME, values)
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

	return role, nil
}

func (repo *roleRepository) FindById(id string) (*role.Role, error) {
	str, err := fileutils.LoadResourceAsString(SELECT_ROLE_BY_ID_LOC)
	if err != nil {
		logger.Error(fmt.Sprintf("cannot open file to this path: %s", SELECT_ROLE_BY_ID_LOC), err)
		return nil, err
	}

	query := fmt.Sprintf(str, ROLE_TBL_COLUMNS, ROLE_TABLE_NAME)
	logger.Info(fmt.Sprintf("Query: %s", query))

	stmt, err := mysql.Client.Prepare(query)
	if err != nil {
		return nil, errorutils.ErrInvalidQuery
	}
	defer stmt.Close()

	role := role.Role{}
	result := stmt.QueryRow(id)
	if err := result.Scan(&role.Id, &role.RoleName, &role.DisplayName, &role.IsActive, &role.IsDeleted, &role.CreatedAt, &role.UpdatedAt, &role.DeletedAt); err != nil {
		logger.Error("role not found", err)
		return nil, ErrRoleNotFound
	}

	return &role, nil
}

func (repo *roleRepository) Find(filter role.RoleFilter) (*role.Role, error) {
	str, err := fileutils.LoadResourceAsString(SELECT_ALL_ROLE_LOC)
	if err != nil {
		logger.Error(fmt.Sprintf("cannot open file to this path: %s", SELECT_ALL_ROLE_LOC), err)
		return nil, err
	}

	args := make([]interface{}, 0)
	var condition string = "1 = 1 "
	if filter.Id != "" && len(filter.Id) > 0 {
		condition += "AND id = ? "
		args = append(args, filter.Id)
	}
	if filter.RoleName != "" && len(filter.RoleName) > 0 {
		condition += "AND role_name = ? "
		args = append(args, filter.RoleName)
	}
	if filter.Active != nil {
		condition += "AND active = ? "
		args = append(args, filter.Active)
	}
	if filter.Deleted != nil {
		condition += "AND deleted = ? "
		args = append(args, filter.Deleted)
	}

	query := fmt.Sprintf(str, ROLE_TBL_COLUMNS, ROLE_TABLE_NAME, condition)
	logger.Info(query)

	stmt, err := mysql.Client.Prepare(query)
	if err != nil {
		return nil, errorutils.ErrInvalidQuery
	}
	defer stmt.Close()

	role := role.Role{}
	result := stmt.QueryRow(args...)
	if err := result.Scan(&role.Id, &role.RoleName, &role.DisplayName, &role.IsActive, &role.IsDeleted, &role.CreatedAt, &role.UpdatedAt, &role.DeletedAt); err != nil {
		logger.Error("role not found", err)
		return nil, ErrRoleNotFound
	}

	return &role, nil
}

func (repo *roleRepository) FindAll() []role.Role {
	roles := make([]role.Role, 0)
	str, err := fileutils.LoadResourceAsString(SELECT_ALL_ROLE_LOC)
	if err != nil {
		logger.Error(fmt.Sprintf("cannot open file to this path: %s", SELECT_ALL_ROLE_LOC), err)
		return roles
	}
	query := fmt.Sprintf(str, ROLE_TBL_COLUMNS, ROLE_TABLE_NAME, "1=1 and active=true and deleted=false order by updated_at desc")
	stmt, err := mysql.Client.Prepare(query)
	if err != nil {
		logger.Error("invalid query", err)
		return roles
	}
	defer stmt.Close()

	rows, err := stmt.Query()
	if err != nil {
		logger.Debug("No data present")
		return roles
	}
	defer rows.Close()

	for rows.Next() {
		role := role.Role{}
		rows.Scan(&role.Id, &role.RoleName, &role.DisplayName, &role.IsActive, &role.IsDeleted, &role.CreatedAt, &role.UpdatedAt, &role.DeletedAt)
		roles = append(roles, role)
	}

	return roles
}

func (repo *roleRepository) Delete(request *role.Role) error {
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
