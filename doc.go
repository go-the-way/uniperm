// Copyright 2025 uniperm Author. All Rights Reserved.
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//      http://www.apache.org/licenses/LICENSE-2.0
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Package uniperm provides a comprehensive permission management system for Go applications.
// It includes models for users, roles, and permissions, along with services for managing these entities.
// The package supports database operations using GORM and includes a CLI tool (permctl) for generating and managing permission data.
//
// Key Features:
//   - User management: Create, update, delete users; manage roles and states.
//   - Role management: Create, update, delete roles; assign permissions.
//   - Permission management: Hierarchical permissions with tree structure support.
//   - Database integration: Auto-migration, pagination, and custom checks.
//   - CLI tool: Generate DDL, examples, and insert permission data from JSON.
//
// For detailed usage, see the README.md file.
package uniperm
