package main

import (
	"encoding/json"
	"fmt"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	//"fmt"
	//"io"
	"log"
)

var (

	userCli IUserRPCServiceClient
)

func main() {

	conn, err := grpc.Dial("127.0.0.1:9002", grpc.WithInsecure())

	if err != nil {
		log.Fatal("failed to connect : ", err)
	}

	userCli = NewIUserRPCServiceClient(conn)

	GetGroupList()
	AddUser( "hhgcnet2" , "987654321")
	AddUserPlus( "hhgcnet444" , "987654321","administrators" ,"5555" )

}

/**
* 获取用户组列表
 */
func GetGroupList() {

	type datalist struct {
		Data []string `json:"Data"`
	}

	rStr, _ := userCli.GetGroupList(context.Background(), &GetGroupList_Request{})
	data := []byte(rStr.GroupList)

	var tmp datalist
	err := json.Unmarshal(data, &tmp)

	if err != nil {
		log.Println(err)

	}
	fmt.Println(tmp.Data)

	for i, val := range tmp.Data {
		fmt.Println(i+1, val)
	}

}

/**
* 获取用户组中的用户列表
 */
func GetGroupUsers(groupName string) {

	type datalist struct {
		Data []string `json:"Data"`
	}

	rStr, _ := userCli.GetGroupUsers(context.Background(), &GetGroupUsers_Request{Group: groupName})
	data := []byte(rStr.UserList)

	var tmp datalist

	err := json.Unmarshal(data, &tmp)

	if err != nil {
		log.Println(err)
	}
	fmt.Println(tmp.Data)

	for i, val := range tmp.Data {
		fmt.Println(i+1, val)
	}
}

/**
* 获取用户属于哪些用户组的列表
 */
func GetUserGroups(userName string) {

	type datalist struct {
		Data []string `json:"Data"`
	}

	rStr, _ := userCli.GetUserGroups(context.Background(), &GetUserGroups_Request{User: userName})
	data := []byte(rStr.UGroupList)

	var tmp datalist
	err := json.Unmarshal(data, &tmp)

	if err != nil {
		log.Println(err)
	}
	fmt.Println(tmp.Data)

	for i, val := range tmp.Data {
		fmt.Println(i+1, val)
	}

}

/**
* 添加用户
 */
func AddUser(userName string, password string) {

	code, _ := userCli.AddUser(context.Background(), &AddUser_Request{UserName: userName, Password: password})
	log.Printf("AddUser userName:%s password:%s code: %d\n", userName, password, code.Code)
}

func AddUserPlus(userName string, password string , groupName1 string ,groupName2 string  ) {

	code, _ := userCli.AddUserPlus(context.Background(), &AddUserPlus_Request{ UserName: userName, Password: password ,GroupName1: groupName1 ,GroupName2: groupName2})
	log.Printf("AddUserPlus userName:%s password:%s code: %d\n", userName, password, code.RCode)
}


/**
* 删除用户
 */
func DelUser(username string) {

	code, _ := userCli.DelUser(context.Background(), &DelUser_Request{UserName: username})
	log.Printf("DelUser userName:%s code:%d\n", username, code.Code)

}

/**
* 添加用户组
 */
func AddGroup(groupName string) {

	code, _ := userCli.AddGroup(context.Background(), &AddGroup_Request{GroupName: groupName})
	log.Printf("AddGroup groupName:%s code:%d\n", groupName, code.Code)

}

/**
* 删除用户组
 */
func DelGroup(groupName string) {

	code, _ := userCli.DelGroup(context.Background(), &DelGroup_Request{GroupName: groupName})
	log.Printf("DelGroup groupName:%s code:%d\n", groupName, code.Code)

}

/**
* 将用户加入到用户组
 */
func AddUserToGroup(userName string, groupName string) {

	code, _ := userCli.AddUserToGroup(context.Background(), &AddUserToGroup_Request{User: userName, Group: groupName})
	log.Printf("AddUserToGroup groupName:%s, userName:%s code:%d\n", groupName, userName, code.Code)

}

/**
* 将用户从用户组中移除
 */
func DelUserToGroup(userName string, groupName string) {

	code, _ := userCli.DelUserToGroup(context.Background(), &DelUserToGroup_Request{User: userName, Group: groupName})
	log.Printf("DelUserToGroup groupName:%s userName:%s code:%d\n", groupName, userName, code.Code)

}

/**
* 禁用和启用用户
 */
func SetUserActive(userName string, bActive int) {

	code, _ := userCli.SetUserActive(context.Background(), &SetUserActive_Request{User: userName, BActive: int32(bActive)})
	log.Printf("SetUserActive userName:%s bActive:%d code:%d\n", userName, bActive, code.Code)

}

/**
* 用户密码设置：需要验证原密码
 */
func SetUserPassword(userName string, oldPassword string, newPassword string) {

	code, _ := userCli.SetUserPassword(context.Background(), &SetUserPassword_Request{User: userName, OldPassword: oldPassword, NewPassword: newPassword})
	log.Printf("SetUserPassword userName:%s, oldpassword:%s newpassword:%s code:%d\n", userName, oldPassword, newPassword, code.Code)

}

/**
* 用户密码设置：无需验证原密码，管理员强制重置
 */
func SetUserPasswordAdmin(userName string, userPassword string) {

	code, _ := userCli.SetUserPasswordAdmin(context.Background(), &SetUserPasswordAdmin_Request{UserName: userName, UserPassword: userPassword})
	log.Printf("SetUserPasswordAdmin userName:%s, password:%s code:%d\n", userName, userPassword, code.Code)
}

/**
* 根据错误返回错误信息，有bug 待修正
 */
func ReturnErrInfo(code int) {

	sStr, _ := userCli.ReturnErrInfo(context.Background(), &ReturnErrInfo_Request{Code: int32(code)})
	log.Printf("ReturnErrInfo code:%d\n", sStr.RCode)

	fmt.Printf("ReturnErrInfo:%s\n", sStr.ErrorText)
}
