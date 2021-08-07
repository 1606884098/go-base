package models

import (
	"encoding/json"
	"fmt"
	"github.com/gomodule/redigo/redis"
	"go-base/study_project/go_web/chatroom/common/message"
)

type UserDao struct {
	pool *redis.Pool
}

var (
	MyUserDao *UserDao
)

//工厂模式
func NewUserDao(pool *redis.Pool) (userDao *UserDao) {
	userDao = &UserDao{
		pool: pool,
	}
	return
}

//根据用户id查询用户
//HSet  HGet
func (this *UserDao) getUserById(conn redis.Conn, id int) (user *message.User, err error) {
	//通过给定的id去redis数据库查询这个用户
	res, err := redis.String(conn.Do("HGet", "users", id))
	if err != nil {
		//表示在users哈希表中没有找到对应的id
		if err == redis.ErrNil {
			err = ERROR_USER_NOEXISTS
		}
		return
	}
	user = &message.User{}
	//将从数据库中得到的数据反序列化成User类型的变量
	err = json.Unmarshal([]byte(res), user)
	if err != nil {
		fmt.Println("json unmarshal err = ", err)
		return
	}
	return
}

func (this *UserDao) Register(user *message.User) (err error) {
	//从连接池中获取一个连接
	conn := this.pool.Get()
	defer conn.Close()
	_, err = this.getUserById(conn, user.UserId)
	if err == nil {
		err = ERROR_USER_EXISTS
		return
	}
	fmt.Println("redis没有该用户!")
	//对用户进行序列化
	data, err := json.Marshal(user)
	if err != nil {
		return
	}
	//在users哈希表中存入键值对
	_, err = conn.Do("HSet", "users", user.UserId, string(data))
	if err != nil {
		fmt.Println("保存注册用户错误, err  = ", err)
		return
	}
	return
}

func (this *UserDao) Login(userId int, userPwd string) (user *message.User, err error) {
	conn := this.pool.Get()
	defer conn.Close()
	user, err = this.getUserById(conn, userId)
	if err != nil {
		return
	}
	//判断密码是否正确
	if user.UserPwd != userPwd {
		err = ERROR_USER_PWD
		return
	}
	return
}
