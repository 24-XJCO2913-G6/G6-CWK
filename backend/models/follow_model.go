package models

import "fmt"

type Follow struct {
	Fid        int64 `xorm:"pk autoincr"` // 轨迹的唯一标识符，被标记为主键
	FollowId   int64
	FollowById int64
	Time       string
}

type Friend struct {
	Uid  int64  // 用户的唯一标识符
	Name string // 用户名
	Pic  string // 头像
}

func BuildModelFollow() {
	err := Db.Sync2(new(Follow))
	if err != nil {
		panic(err)
	}
}

func AddFollow(FollowId int64, FollowById int64, Time string) int64 {
	follow := &Follow{
		FollowId:   FollowId,
		FollowById: FollowById,
		Time:       Time,
	}

	Fid, err := Db.Insert(follow)
	if err != nil {
		return -1
	}
	return Fid
}

// GetFans 返回指定用户关注的用户
func GetFans(uid int64) ([]Friend, error) {
	var fans []Friend

	// 查询该用户关注的用户
	var fanUIDs []int64
	err := Db.Table("follow").Cols("follow_by_id").Where("follow_id = ?", uid).Find(&fanUIDs)
	if err != nil {
		return nil, err
	}

	// 查询关注的用户的信息
	if len(fanUIDs) != 0 {
		for _, followingUID := range fanUIDs {
			var tmpUser User
			has, err := Db.Table("user").Where("uid = ?", followingUID).Get(&tmpUser)
			fmt.Println(tmpUser)
			friend := Friend{Uid: tmpUser.Uid, Name: tmpUser.Name, Pic: tmpUser.ProfilePhot}
			if err != nil {
				return nil, err
			}
			if has {
				fans = append(fans, friend)
			}
		}
		return fans, nil
	}

	return nil, nil
}

// GetFans 返回指定用户的粉丝
func GetFollowings(uid int64) ([]Friend, error) {
	var followings []Friend

	// 查询关注该用户的用户，即该用户的粉丝
	var followingUIDs []int64
	err := Db.Table("follow").Cols("follow_id").Where("follow_by_id = ?", uid).Find(&followingUIDs)
	if err != nil {
		return nil, err
	}
	fmt.Println(followingUIDs)
	// 查询粉丝的信息
	if len(followingUIDs) != 0 {
		for _, followerUID := range followingUIDs {
			var tmpUser User
			has, err := Db.Table("user").Where("uid = ?", followerUID).Get(&tmpUser)
			fmt.Println(tmpUser)
			friend := Friend{Uid: tmpUser.Uid, Name: tmpUser.Name, Pic: tmpUser.ProfilePhot}
			if err != nil {
				return nil, err
			}
			if has {
				followings = append(followings, friend)
			}
		}
		return followings, nil
	}
	return nil, nil
}

func GetFriends(Uid int64) ([]Friend, error) {
	var friends []Friend

	// 查询所有关注该用户的关注者
	followings, _ := GetFollowings(Uid)
	fans, _ := GetFans(Uid)

	//fmt.Println("followings:")
	//fmt.Println(followings)
	//fmt.Println("fans:")
	//fmt.Println(fans)

	if min(len(fans), len(followings)) == 0 {
		return nil, nil
	}
	for _, fan := range fans {
		for _, following := range followings {
			if fan.Uid == following.Uid {
				friends = append(friends, fan)
			}
		}
	}
	//fmt.Println(friends)
	return friends, nil
}
