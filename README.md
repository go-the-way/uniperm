# uniperm

`uniperm` is a Go library designed for managing users, roles, and permissions in applications. It provides a structured approach to handle authentication and authorization using a hierarchical permission system. The library integrates with GORM for database operations and includes a CLI tool called `permctl` for administrative tasks like generating DDL scripts, example JSON, and inserting permission data from JSON files into the database.

## Features

- **User Management**: Create, update, delete users; assign roles; enable/disable users; handle login/logout.
- **Role Management**: Create, update, delete roles; assign permissions; enable/disable roles.
- **Permission Management**: Hierarchical permissions (tree structure) with support for buttons and routes; add, update, delete permissions.
- **Database Integration**: Uses GORM for MySQL; supports auto-migration, pagination, and custom validation checks.
- **CLI Tool (`permctl`)**:
    - Print DDL for permission tables.
    - Print example JSON for permissions.
    - Generate and insert permission data from JSON into the database, with options for truncation, creation, and skipping empty names.
- **Security**: Passwords are stored as MD5 hashes; super admin protection.
- **Extensibility**: Customizable pagination and database setup.

## Installation

To install the library, use `go get`:

```bash
go get github.com/go-the-way/uniperm
```

Note: The CLI tool `permctl` is built separately. To install it, navigate to the project root and run:

```bash
go install
```

This will make `permctl` available in your `$GOPATH/bin`.

## Usage

### Setting Up the Database

Uniperm uses GORM with MySQL. First, set up your database connection and configure the library:

```go
import (
    "github.com/go-the-way/uniperm"
    "gorm.io/driver/mysql"
    "gorm.io/gorm"
)

func main() {
    dsn := "user:pass@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local"
    db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
    if err != nil {
        panic(err)
    }
    uniperm.SetDB(db)
    uniperm.SetPagination(customPaginationFunc) // Optional: Provide your own pagination logic
    if err := uniperm.AutoMigrate(); err != nil {
        panic(err)
    }
    // Now you can use the services
}
```

The `AutoMigrate` function creates tables for `User`, `Role`, `RolePermission`, and `Permission` models.

#### Custom Pagination

You can provide a custom pagination function:

```go
func customPagination(db *gorm.DB, page, limit int, count *int64, list any) error {
    // Implement your pagination logic here
    return nil
}
```

### Models

Uniperm defines the following models:

- **User**: Represents application users with fields like username, password (MD5), role ID, state, business IDs, remarks, login details.
- **Role**: Represents roles with name, description, type, state.
- **Permission**: Hierarchical permissions with name, route, parent ID, and button flag.
- **RolePermission**: Maps roles to permissions.

Constants are provided for states (e.g., `UserStateEnable`, `RoleStateEnable`) and button flags.

### Services

Uniperm exposes services for users, roles, and permissions.

#### User Service

```go
// Get a paginated list of users
resp, err := uniperm.UserGetPage(uniperm.UserGetPageReq{Page: 1, Limit: 10})

// Get a single user by ID
resp, err := uniperm.UserGet(uniperm.UserGetReq{Id: 1})

// Get user permissions (routes)
resp, err := uniperm.UserGetPerm(uniperm.UserGetPermReq{Id: 1})

// Get user button permissions
resp, err := uniperm.UserGetPermButton(uniperm.UserGetPermButtonReq{Id: 1})

// Add a new user
err := uniperm.UserAdd(uniperm.UserAddReq{Username: "newuser", Password: "password", RoleId: 1})

// Update user details
err := uniperm.UserUpdate(uniperm.UserUpdateReq{Id: 1, Username: "updateduser"})

// Update user password
err := uniperm.UserUpdatePassword(uniperm.UserUpdatePasswordReq{Id: 1, Password: "newpass"})

// Update user role
err := uniperm.UserUpdateRole(uniperm.UserUpdateRoleReq{Id: 1, RoleId: 2})

// Delete a user
err := uniperm.UserDelete(uniperm.UserDeleteReq{Id: 1})

// Enable a user
err := uniperm.UserEnable(uniperm.UserEnableReq{Id: 1})

// Disable a user
err := uniperm.UserDisable(uniperm.UserDisableReq{Id: 1})

// Login (custom callback for token generation)
err := uniperm.UserLogin(uniperm.UserLoginReq{Username: "user", Password: "pass"}, func(u *uniperm.User) {
    // Generate token or session
})

// Logout
err := uniperm.UserLogout(uniperm.UserLogoutReq{})
```

#### Role Service

```go
// Get a paginated list of roles
resp, err := uniperm.RoleGetPage(uniperm.RoleGetPageReq{Page: 1, Limit: 10})

// Get a single role by ID
resp, err := uniperm.RoleGet(uniperm.RoleGetReq{Id: 1})

// Get role permissions
resp, err := uniperm.RoleGetPerm(uniperm.RoleGetPermReq{Id: 1})

// Update role permissions
err := uniperm.RoleUpdatePerm(uniperm.RoleUpdatePermReq{Id: 1, PermissionId: []uint{1, 2}})

// Add a new role
err := uniperm.RoleAdd(uniperm.RoleAddReq{Name: "newrole"})

// Update role details
err := uniperm.RoleUpdate(uniperm.RoleUpdateReq{Id: 1, Name: "updatedrole"})

// Delete a role
err := uniperm.RoleDelete(uniperm.RoleDeleteReq{Id: 1})

// Enable a role
err := uniperm.RoleEnable(uniperm.RoleEnableReq{Id: 1})

// Disable a role
err := uniperm.RoleDisable(uniperm.RoleDisableReq{Id: 1})
```

#### Permission Service

```go
// Get permission tree
resp, err := uniperm.PermissionTree(uniperm.PermissionTreeReq{PermissionId: []uint{1, 2}})

// Get a single permission by ID
resp, err := uniperm.PermissionGet(uniperm.PermissionGetReq{Id: 1})

// Add a new permission
err := uniperm.PermissionAdd(uniperm.PermissionAddReq{Name: "newperm", Route: "/new", ParentId: 0, IsButton: false})

// Update permission details
err := uniperm.PermissionUpdate(uniperm.PermissionUpdateReq{Id: 1, Name: "updatedperm"})

// Delete a permission
err := uniperm.PermissionDelete(uniperm.PermissionDeleteReq{Id: 1})
```

### CLI Tool: permctl

The `permctl` tool is used for database setup and permission data management.

#### Commands

- **ddl**: Print the DDL for the permission table.
  ```bash
  permctl ddl
  ```

- **example**: Print an example JSON for permissions.
  ```bash
  permctl example
  ```

- **generate**: Generate and insert permission data from a JSON file into the database.
  ```bash
  permctl generate --jsonFile perm.json --dbHost 127.0.0.1 --dbPort 3306 --dbUser root --dbPasswd passwd --dbName uniperm --create --truncate
  ```
  Flags:
    - `--table`: Table name (default: "uniperm_permissions")
    - `--dbHost`, `--dbPort`, `--dbUser`, `--dbPasswd`, `--dbName`: Database connection details.
    - `--dbDSNFromEnv`: Read DSN from environment variable.
    - `--dbDSNEnvName`: Environment variable name for DSN (default: "PERMCTL_DSN").
    - `--jsonFile`: Path to JSON file (default: "perm.json").
    - `--create`: Create the table if it doesn't exist.
    - `--truncate`: Truncate the table before inserting.
    - `--skipEmptyName`: Skip permissions with empty names.
    - `--verbose`: Show verbose output.
    - `--printInsertSqlTpl`: Print the insert SQL template and exit.

#### JSON Format for Permissions

The JSON file should follow this structure (recursive for hierarchies):

```json
[
  {
    "path": "/dashboard",
    "name": "Dashboard",
    "is_button": false,
    "routes": [
      {
        "path": "/dashboard/view",
        "name": "View Dashboard",
        "is_button": true
      }
    ]
  }
]
```

## Validation and Checks

The library includes built-in checks to prevent invalid operations, such as:
- Ensuring usernames are unique.
- Preventing deletion of super admin (ID 1).
- Checking for references before deleting roles or permissions.
- Validating permission types (e.g., not buttons for certain operations).

## Contributing

Contributions are welcome! Please submit issues or pull requests on the GitHub repository.

## License

Apache License. See LICENSE file for details.