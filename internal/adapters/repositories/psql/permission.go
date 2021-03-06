package psql

import (
	"context"
	"errors"
	"fmt"
	"github.com/iBoBoTi/project_boiler_plate/internal/core/domain"
	"github.com/iBoBoTi/project_boiler_plate/internal/core/helpers"
	"github.com/iBoBoTi/project_boiler_plate/internal/core/ports"
	"github.com/jackc/pgx/v4/pgxpool"
	"math"
)

type permissionRepository struct {
	db *pgxpool.Pool
}

func NewPermissionRepository(db *pgxpool.Pool) ports.PermissionRepository {
	return &permissionRepository{db: db}
}

func (p *permissionRepository) CreatePermission(permission *domain.Permission) (*domain.Permission, error) {
	queryString := `INSERT INTO permissions (id, title) VALUES ($1, $2) RETURNING *`
	result := &domain.Permission{}
	row := p.db.QueryRow(context.Background(), queryString, permission.ID, permission.Title)
	err := row.Scan(&result.ID, &result.Title, &result.CreatedAt, &result.UpdatedAt)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (p *permissionRepository) GetPermissionByTitle(title string) (*domain.Permission, error) {
	permission := domain.Permission{}
	queryString := `SELECT * FROM permissions WHERE title = $1`
	err := p.db.QueryRow(context.Background(), queryString, title).Scan(&permission.ID, &permission.Title, &permission.CreatedAt, &permission.UpdatedAt)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return &permission, nil
}

func (p *permissionRepository) DeletePermission(id string) error {
	queryString := `DELETE FROM permissions WHERE id = $1`
	cmdTag, err := p.db.Exec(context.Background(), queryString, id)
	if err != nil {
		return err
	}
	if cmdTag.RowsAffected() != 1 {
		return errors.New("permission not deleted")
	}
	return nil
}

func (p *permissionRepository) GetPermissionByID(id string) (*domain.Permission, error) {
	permission := domain.Permission{}
	queryString := `SELECT * FROM permissions WHERE id = $1`
	err := p.db.QueryRow(context.Background(), queryString, id).Scan(&permission.ID, &permission.Title, &permission.CreatedAt, &permission.UpdatedAt)

	if err != nil {
		return nil, err
	}
	return &permission, nil
}

func (p *permissionRepository) GetAllPermissions(paginator *helpers.Paginate) (*helpers.Paginate, error) {
	permissions := make([]domain.Permission, 0)
	queryString := `SELECT * FROM permissions ORDER BY created_at DESC LIMIT $1 OFFSET $2`
	rows, err := p.db.Query(context.Background(), queryString, paginator.Limit, paginator.Offset)
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

	count := 0
	_ = p.db.QueryRow(context.Background(), `SELECT COUNT(*) FROM roles`).Scan(&count)
	paginator.TotalRows = count
	paginator.Rows = permissions
	paginator.TotalPages = int(math.Ceil(float64(count) / float64(paginator.Limit)))

	return paginator, nil
}
