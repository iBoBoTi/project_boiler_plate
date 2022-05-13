package psql

import (
	"context"
	"errors"
	"github.com/iBoBoTi/project_boiler_plate/internal/core/domain"
	"github.com/iBoBoTi/project_boiler_plate/internal/core/helpers"
	"github.com/iBoBoTi/project_boiler_plate/internal/core/ports"
	"github.com/jackc/pgx/v4/pgxpool"
	"math"
)

type roleRepository struct {
	db *pgxpool.Pool
}

func NewRoleRepository(db *pgxpool.Pool) ports.RoleRepository {
	return &roleRepository{db}
}

func (r *roleRepository) GetAllRoles(paginator *helpers.Paginate) (*helpers.Paginate, error) {
	roles := make([]domain.Role, 0)
	queryString := `SELECT * FROM roles ORDER BY created_at DESC LIMIT $1 OFFSET $2`
	rows, err := r.db.Query(context.Background(), queryString, paginator.Limit, paginator.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		role := domain.Role{}
		err := rows.Scan(&role.ID, &role.Title, &role.Description, &role.CreatedAt, &role.UpdatedAt)
		if err != nil {
			return nil, err
		}
		roles = append(roles, role)
	}

	count := 0
	_ = r.db.QueryRow(context.Background(), `SELECT COUNT(*) FROM roles`).Scan(&count)
	paginator.TotalRows = count
	paginator.Rows = roles
	paginator.TotalPages = int(math.Ceil(float64(count) / float64(paginator.Limit)))

	return paginator, nil
}

func (r *roleRepository) GetRoleByID(id string) (*domain.Role, error) {
	role := domain.Role{}
	queryString := `SELECT * FROM roles WHERE id = $1`
	err := r.db.QueryRow(context.Background(), queryString, id).Scan(&role.ID, &role.Title, &role.Description, &role.CreatedAt, &role.UpdatedAt)
	if err != nil {
		return nil, err
	}
	return &role, nil
}

func (r *roleRepository) GetRoleByName(name string) (*domain.Role, error) {
	role := domain.Role{}
	queryString := `SELECT * FROM roles WHERE title = $1`
	err := r.db.QueryRow(context.Background(), queryString, name).Scan(&role.ID, &role.Title, &role.Description, &role.CreatedAt, &role.UpdatedAt)
	if err != nil {
		return nil, err
	}
	return &role, nil
}

func (r *roleRepository) CreateRole(role *domain.Role) (*domain.Role, error) {
	queryString := `INSERT INTO roles (id, title, description) VALUES ($1, $2, $3) RETURNING *`
	result := &domain.Role{}
	row := r.db.QueryRow(context.Background(), queryString, role.ID, role.Title, role.Description)
	err := row.Scan(&result.ID, &result.Title, &result.Description, &result.CreatedAt, &result.UpdatedAt)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (r *roleRepository) DeleteRole(id string) error {
	queryString := `DELETE FROM roles WHERE id = $1`
	cmdTag, err := r.db.Exec(context.Background(), queryString, id)
	if err != nil {
		return err
	}
	if cmdTag.RowsAffected() != 1 {
		return errors.New("role not deleted")
	}
	return nil
}
