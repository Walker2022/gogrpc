package main

import (

	"fmt"
	"log"
	"strings"

	"golang.org/x/net/context"
)

func NewUserRPCService() *UserRPCService {
	return &UserRPCService{}
}


type UserRPCService struct {
}

/**
* 获取用户组列表
*/
func (s *UserRPCService)GetGroupList( ctx context.Context, req *GetGroupList_Request) (*GetGroupList_Response,error ) {

	var buf []byte = make([]byte, 65536)
	var bufsize int = 0
	code, _, _ := fnGetGroupList.Call(newBuffer(buf), newBufferSize(&bufsize))

	log.Printf("GetGroupList bufsize:%d code:%d\n", bufsize, code)

	rStr := string( buf[:bufsize] )

	return &GetGroupList_Response{ GroupList:rStr },nil
}

/**
* 获取用户组中的用户列表
 */
func (s *UserRPCService)GetGroupUsers( ctx context.Context, req *GetGroupUsers_Request) (*GetGroupUsers_Response,error ){

	var buf []byte = make([]byte, 65536)
	var bufsize int = 0
	code, _, _ := fnGetGroupUsers.Call(AStrPtr(req.Group), newBuffer(buf), newBufferSize(&bufsize))

	log.Printf("GetGroupUsers bufsize:%d code:%d\n", bufsize, code)

	rStr := string( buf[:bufsize] )

	return &GetGroupUsers_Response{ UserList:rStr },nil
}

/**
* 获取用户属于哪些用户组的列表
 */
func (s *UserRPCService)GetUserGroups( ctx context.Context, req *GetUserGroups_Request) (*GetUserGroups_Response,error ) {

	var buf []byte = make([]byte, 65536)
	var bufsize int = 0
	code, _, _ := fnGetUserGroups.Call(AStrPtr(req.User), newBuffer(buf), newBufferSize(&bufsize))

	log.Printf("GetUserGroups bufsize:%d code:%d\n", bufsize, code)

	rStr := string( buf[:bufsize] )

	return &GetUserGroups_Response{ UGroupList:rStr },nil
}

/**
* 添加用户
 */
func (s *UserRPCService)AddUser( ctx context.Context, req *AddUser_Request) (*AddUser_Response,error ){

	code, _, _ := fnAddUser.Call(AStrPtr(req.UserName), AStrPtr(req.Password))
	log.Printf("AddUser userName:%s password:%s code: %d\n", req.UserName, "********", code)

	return &AddUser_Response{ Code: int32(code)},nil
}


func (s *UserRPCService)AddUserPlus( ctx context.Context, req *AddUserPlus_Request) (*AddUserPlus_Response,error ){

	code, _, _ := fnAddUserPlus.Call(AStrPtr(req.UserName), AStrPtr(req.Password) , AStrPtr(req.GroupName1) ,AStrPtr(req.GroupName2) )
	log.Printf("AddUserPlus userName:%s password:%s code: %d\n", req.UserName, "********", code)

	return &AddUserPlus_Response{ RCode: int32(code)},nil
}


/**
* 删除用户
 */
func (s *UserRPCService)DelUser( ctx context.Context, req *DelUser_Request) (*DelUser_Response,error ) {

	code, _, _ := fnDelUser.Call(AStrPtr(req.UserName))
	log.Printf("DelUser userName:%s code:%d\n", req.UserName, code)

	return &DelUser_Response{Code: int32(code)},nil
}

/**
* 添加用户组
 */
func (s *UserRPCService)AddGroup( ctx context.Context, req *AddGroup_Request) (*AddGroup_Response,error ) {

	code, _, _ := fnAddGroup.Call(AStrPtr(req.GroupName))
	log.Printf("AddGroup groupName:%s code:%d\n", req.GroupName, code)

	return &AddGroup_Response{ Code: int32(code)},nil
}

/**
* 删除用户组
 */
func (s *UserRPCService)DelGroup( ctx context.Context, req *DelGroup_Request) (*DelGroup_Response,error ){

	code, _, _ := fnDelGroup.Call(AStrPtr(req.GroupName))
	log.Printf("DelGroup groupName:%s code:%d\n", req.GroupName, code)

	return &DelGroup_Response{ Code: int32(code)},nil
}

/**
* 将用户加入到用户组
 */
func (s *UserRPCService)AddUserToGroup( ctx context.Context, req *AddUserToGroup_Request) (*AddUserToGroup_Response,error ){

	code, _, _ := fnAddUserToGroup.Call(AStrPtr(req.Group), AStrPtr(req.User))
	log.Printf("AddUserToGroup groupName:%s, userName:%s code:%d\n", req.Group, req.User, code)

	return &AddUserToGroup_Response{ Code: int32(code)},nil
}

/**
* 将用户从用户组中移除
 */
func (s *UserRPCService)DelUserToGroup( ctx context.Context, req *DelUserToGroup_Request) (*DelUserToGroup_Response,error ){

	code, _, _ := fnDelUserToGroup.Call(AStrPtr(req.Group), AStrPtr(req.User))
	log.Printf("DelUserToGroup groupName:%s userName:%s code:%d\n", req.Group, req.User, code)

	return &DelUserToGroup_Response{ Code: int32(code)},nil
}

/**
* 禁用和启用用户
 */
func (s *UserRPCService)SetUserActive( ctx context.Context, req *SetUserActive_Request) (*SetUserActive_Response,error ) {

	code, _, _ := fnSetUserActive.Call(AStrPtr(req.User), IntPtr( int(req.BActive)))
	log.Printf("SetUserActive userName:%s bActive:%d code:%d\n", req.User, req.BActive, code)

	return &SetUserActive_Response{ Code: int32(code)},nil
}

/**
* 用户密码设置：需要验证原密码
 */
func (s *UserRPCService)SetUserPassword( ctx context.Context, req *SetUserPassword_Request) (*SetUserPassword_Response,error ){

	code, _, _ := fnSetUserPassword.Call(AStrPtr(req.User), AStrPtr(req.OldPassword), AStrPtr(req.NewPassword))
	log.Printf("SetUserPassword userName:%s, oldpassword:%s newpassword:%s code:%d\n", req.User, req.OldPassword, req.NewPassword, code)

	return &SetUserPassword_Response{ Code: int32(code)},nil
}

/**
* 用户密码设置：无需验证原密码，管理员强制重置
 */
func (s *UserRPCService)SetUserPasswordAdmin( ctx context.Context, req *SetUserPasswordAdmin_Request) (*SetUserPasswordAdmin_Response,error ){

	code, _, _ := fnSetUserPasswordAdmin.Call(AStrPtr(req.UserName), AStrPtr(req.UserPassword))
	log.Printf("SetUserPasswordAdmin userName:%s, password:%s code:%d\n", req.UserName, req.UserPassword, code)

	return &SetUserPasswordAdmin_Response{ Code: int32(code)},nil
}

/**
* 根据错误返回错误信息，有bug 待修正
 */
func (s *UserRPCService)ReturnErrInfo( ctx context.Context, req *ReturnErrInfo_Request) (*ReturnErrInfo_Response,error ){

	var buf []byte = make([]byte, 65536)
	var bufsize int = 0

	scode, _, _ := fnReturnErrInfo.Call(IntPtr( int( req.Code) ), newBuffer(buf), newBufferSize(&bufsize))
	log.Printf("SetUserPasswordAdmin lcode:%d ReturnErrInfo-bufsize:%d\n", scode, bufsize)

	outText := strings.Trim( string(buf[:bufsize]) ,"\r\n")
	fmt.Printf("ReturnErrInfo:%s\n", outText )

	return &ReturnErrInfo_Response{ RCode: int32(scode), ErrorText: outText},nil
}