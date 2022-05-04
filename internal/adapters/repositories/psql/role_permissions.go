package psql

import (
	"context"
	"errors"
	"github.com/iBoBoTi/project_boiler_plate/internal/core/domain"
	"github.com/iBoBoTi/project_boiler_plate/internal/core/ports"
	"github.com/jackc/pgx/v4/pgxpool"
)

type rolePermissionsRepository struct {
	db  *pgxpool.Pool
	log ports.Logger
}

func NewRolePermissionsRepository(db *pgxpool.Pool, log ports.Logger) ports.RolePermissionsRepository {
	return &rolePermissionsRepository{
		db:  db,
		log: log,
	}
}

func (r *rolePermissionsRepository) AddPermissionsToRole(rolePermissions *domain.RolePermissions) error {
	queryString := `INSERT INTO role_permissions (role_id, permission_id) VALUES ($1, $2)`
	for _, permID := range rolePermissions.PermIDs {
		_, err := r.db.Exec(context.Background(), queryString, rolePermissions.RoleID, permID)
		if err != nil {
			return err
		}
	}
	return nil
}

func (r *rolePermissionsRepository) RemovePermissionFromRole(roleId string, permissionID string) error {
	queryString := `DELETE FROM role_permissions WHERE role_id = $1 AND permission_id = $2`
	cmdTag, err := r.db.Exec(context.Background(), queryString, roleId, permissionID)
	if err != nil {
		return err
	}
	if cmdTag.RowsAffected() != 1 {
		return errors.New("permission not removed")
	}
	return nil
}

func (r *rolePermissionsRepository) GetAllPermissionsForRole(roleId string) ([]domain.Permission, error) {
	permissions := make([]domain.Permission, 0)
	queryString := `SELECT permissions.id,permissions.title,permissions.created_at, permissions.updated_at FROM role_permissions FULL OUTER JOIN permissions ON role_permissions.permission_id=permissions.id WHERE role_permissions.role_id = $1`

	rows, err := r.db.Query(context.Background(), queryString, roleId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		permission := domain.Permission{}
		err := rows.Scan(&permission.ID, &permission.Title, &permission.CreatedAt, &permission.UpdatedAt)
		if err != nil {
			return nil, err
		}
		permissions = append(permissions, permission)
	}
	return permissions, nil
}
