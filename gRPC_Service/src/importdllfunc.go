package main

import "syscall"

/**
UserManager.dll 函数声明：
EXPORT ULONG AddUser(const char* userName, const char* password);
EXPORT ULONG DelUser(const char* userName);
EXPORT ULONG AddGroup(const char* groupName);
EXPORT ULONG DelGroup(const char* groupName);
EXPORT ULONG AddUserToGroup(const char* group, const char* user);
EXPORT ULONG DelUserToGroup(const char* group, const char* user);
EXPORT ULONG GetGroupList(char* GroupList,int& size);
EXPORT ULONG GetGroupUsers(const char* group, char* UserList, int& size);
EXPORT ULONG GetUserGroups(const char* user, char* uGroupList, int& size);
EXPORT ULONG SetUserActive(const char* user, bool bActive);
EXPORT ULONG SetUserPassword(const char* user, const char* oldpassword, const char* newpassword);
EXPORT ULONG SetUserPasswordAdmin(const char* UserName, const char* UserPassword);
EXPORT unsigned ReturnErrInfo(ULONG info, char* buf, int& len);
**/

var (
	//os.PathSeparator
	dllPath = "E:/UAM2020.NEW/WINDLL/UserManager.dll"

	handle = syscall.NewLazyDLL(dllPath)

	fnAddUser              = handle.NewProc("AddUser")              //添加用户
	fnDelUser              = handle.NewProc("DelUser")              //删除用户
	fnAddGroup             = handle.NewProc("AddGroup")             //创建用户组
	fnDelGroup             = handle.NewProc("DelGroup")             //删除用户组
	fnAddUserToGroup       = handle.NewProc("AddUserToGroup")       //将用户加入组
	fnDelUserToGroup       = handle.NewProc("DelUserToGroup")       //将用户从用户组中移除
	fnGetGroupList         = handle.NewProc("GetGroupList")         //获取用户组列表
	fnGetGroupUsers        = handle.NewProc("GetGroupUsers")        //获取指定用户组中的所有用户列表
	fnGetUserGroups        = handle.NewProc("GetUserGroups")        //获取用户所属的用户组
	fnSetUserActive        = handle.NewProc("SetUserActive")        //禁用或启用用户
	fnSetUserPassword      = handle.NewProc("SetUserPassword")      //用户设置密码：需原密码
	fnSetUserPasswordAdmin = handle.NewProc("SetUserPasswordAdmin") //重置用户密码：无需原密码

	fnAddUserPlus = handle.NewProc("AddUserPlus")
	//EXPORT ULONG  AddUserPlus(const char* userName, const char* password,const char* groupName1,const char* groupName2);

	fnReturnErrInfo = handle.NewProc("ReturnErrInfo") //获取错误信息

)
