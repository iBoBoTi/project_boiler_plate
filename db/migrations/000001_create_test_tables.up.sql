CREATE TABLE roles (
  id VARCHAR(45) PRIMARY KEY,
  title VARCHAR(255) NOT NULL,
  description TEXT NOT NULL,
  created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
  updated_at TIMESTAMPTZ NOT NULL DEFAULT now()
);

CREATE TABLE permissions  (
  id VARCHAR(45) PRIMARY KEY,
  title VARCHAR(255) NOT NULL,
  created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
  updated_at TIMESTAMPTZ NOT NULL DEFAULT now()
);

CREATE TABLE role_permissions (
  role_id VARCHAR(45) NOT NULL REFERENCES roles(id) ON UPDATE CASCADE ON DELETE CASCADE,
  permission_id VARCHAR(45) NOT NULL REFERENCES permissions(id) ON UPDATE CASCADE ON DELETE CASCADE,
  PRIMARY KEY (role_id, permission_id)
);