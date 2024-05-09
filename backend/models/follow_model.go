package models

type Follow struct {
	Fid        int64 `xorm:"pk"` // 轨迹的唯一标识符，被标记为主键
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

// GetFollowings 返回指定用户关注的用户
func GetFollowings(uid int64) ([]Friend, error) {
	var followings []Friend

	// 查询该用户关注的用户
	var followingUIDs []int64
	err := Db.Table("follow").Cols("follow_by_id").Where("follow_id = ?", uid).Find(&followingUIDs)
	if err != nil {
		return nil, err
	}

	// 查询关注的用户的信息
	for _, followingUID := range followingUIDs {
		var friend Friend
		has, err := Db.Table("user").Cols("uid", "name", "ProfilePho").Where("uid = ?", followingUID).Get(&friend)
		if err != nil {
			return nil, err
		}
		if has {
			followings = append(followings, friend)
		}
	}

	return followings, nil
}

// GetFans 返回指定用户的粉丝
func GetFans(uid int64) ([]Friend, error) {
	var fans []Friend

	// 查询关注该用户的用户，即该用户的粉丝
	var followerUIDs []int64
	err := Db.Table("follow").Cols("follow_id").Where("follow_by_id = ?", uid).Find(&followerUIDs)
	if err != nil {
		return nil, err
	}

	// 查询粉丝的信息
	for _, followerUID := range followerUIDs {
		var friend Friend
		has, err := Db.Table("user").Cols("uid", "name", "ProfilePho").Where("uid = ?", followerUID).Get(&friend)
		if err != nil {
			return nil, err
		}
		if has {
			fans = append(fans, friend)
		}
	}

	return fans, nil
}

func GetFriends(uid int64) ([]Friend, error) {
	var friends []Friend

	// 查询所有关注该用户的关注者
	var follows []Follow
	err := Db.Where("follow_id = ?", uid).Find(&follows)
	if err != nil {
		return nil, err
	}

	// 查询这些关注者的信息，即用户的朋友信息
	for _, follow := range follows {
		var friend Friend
		has, err := Db.Table("user").Cols("uid", "name", "ProfilePho").Where("uid = ?", follow.FollowById).Get(&friend)
		if err != nil {
			return nil, err
		}
		if has {
			friends = append(friends, friend)
		}
	}

	return friends, nil
}
