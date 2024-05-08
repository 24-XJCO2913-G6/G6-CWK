<template>
  <view class="like-messages-page">
    <view class="message-item" v-for="(item, index) in collects" :key="index">
      <image class="avatar" :src="item.avatar"></image>
      <view class="message-content">
        <view class="user-info">
          <text class="username">{{ item.username }}</text>
          <text class="like-action">collect your post</text>
        </view>
        <text class="title">{{ item.title }}</text>
        <text class="date">{{ item.date }}</text>
      </view>
    </view>
  </view>
</template>

<script>
export default {
  data() {
    return {
     collects:[
			    { post_id:1,username: 'Verooo', title: 'School walk', date: '04-20', avatar: '../../static/user_pic/2.jpg' },
				{ post_id:12,username: 'Yodo', title: 'School walk', date: '04-13', avatar: '../../static/user_pic/1.jpg' },
			    {post_id:13, username: 'Jelly cat', title: 'School walk', date: '04-13', avatar: '../../static/user_pic/5.jpg' },
		  // ...更多点赞消息对象
		  // ...更多点赞消息对象
          // ...更多点赞消息对象
        ]
    };
  },
  onLoad(options) {
   uni.request({
			   url: 'http://120.46.81.37:8080/app/get_collects',
			   method: 'GET',
			   data: {	 
					  'aToken': aToken,
					  'rToken':rToken,
			   },
				header: {
					'Content-Type': 'application/x-www-form-urlencoded',
				},
     success: (res) => {
   	this.collects= res.data;
     },
     fail: (err) => {
   	console.error('Failed to fetch posts:', err);
     }
       });
  },

};
</script>

<style scoped>
.like-messages-page {
  padding: 10px; /* 页面内边距 */
}

.message-item {
  display: flex;
  align-items: center;
  margin-bottom: 15px;
  padding: 10px;
  border-bottom: 1px solid #eee; /* 分割线 */
}

.avatar {
  width: 40px;
  height: 40px;
  border-radius: 50%;
  margin-right: 15px;
}

.message-content {
  flex: 1; /* 填充剩余空间 */
  display: flex;
  flex-direction: column;
  justify-content: center;
}

.user-info {
  display: flex;
  align-items: center;
  margin-bottom: 5px;
}

.username {
  font-weight: bold; /* 加粗用户名 */
  margin-right: 5px;
}

.like-action {
  color: #888; /* 点赞动作颜色 */
}

.comment {
  margin-bottom: 5px;
}

.date {
  color: #aaa; /* 日期颜色 */
  font-size: 12px;
}
</style>