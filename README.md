# uniperm
The unified permission package

# Install permctl
```
go install github.com/go-the-way/uniperm/cmd/permctl@latest
```

# Services

## User 用户模块
- user.GetPage 分页查询
- user.Get 查询详情
- user.GetPerm 查询用户权限
- user.GetPermButton 查询用户按钮权限
- user.Add 新增
- user.Update 修改
- user.UpdatePassword 修改密码
- user.UpdateRole 修改用户角色
- user.Delete 删除
- user.Enable 启用
- user.Disable 禁用
- user.Login 用户登录
- user.Logout 用户注销

## Role 角色模块
- role.GetPage 分页查询
- role.Get 查询详情
- role.GetPerm 查询角色权限
- role.UpdatePerm 修改角色权限
- role.Add 新增
- role.Update 修改
- role.Delete 删除
- role.Enable 启用
- role.Disable 禁用

## Permission 权限模块
- permission.Tree 查询权限树
- permission.Get 查询
- permission.Add 新增
- permission.Update 编辑
- permission.Delete 删除
 
# Models
- models.User/models.UnipermUser 用户
- models.Role/models.UnipermRole 角色
- models.Permission/models.UnipermPermission 权限
- models.RolePermission/models.UnipermRolePermission 角色权限