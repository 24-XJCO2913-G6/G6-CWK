<template>
  <view class="like-messages-page">
    <view class="message-item" v-for="(item, index) in messages" :key="index">
      <image class="avatar" :src="item.avatar"></image>
      <view class="message-content">
        <text class="username">{{ item.username }}</text>
        <text class="comment">{{ item.comment }}</text>
        <text class="like-action">liked your comment</text>
        <text class="date">{{ item.date }}</text>
      </view>
    </view>
  </view>
</template>

<script>
export default {
  data() {
    return {
      messages: [] // 用于存储点赞消息的数组
    };
  },
  onLoad(options) {
    // 这里可以获取从上一个页面传递过来的参数，例如:
    const { someParameter } = options;
    // 调用获取点赞消息列表的API，并填充到messages数组中
    this.fetchLikeMessages(someParameter);
  },
  methods: {
    fetchLikeMessages(parameter) {
      // 这里是获取点赞消息列表的API调用示例
      // 请替换为实际的API请求逻辑
      const fakeApiCall = () => {
        // 模拟从API获取数据
        return [
          { username: '团子', comment: '评论内容', date: '04-20' },
          // ...更多点赞消息对象
        ];
      };
      this.messages = fakeApiCall(parameter);
    }
  }
};
</script>

<style>
.like-messages-page {
  /* 页面样式 */
}
.message-item {
  /* 消息条目样式 */
  display: flex;
  align-items: start;
  margin-bottom: 10px;
}
.avatar {
  width: 40px; /* 头像大小 */
  height: 40px;
  border-radius: 50%; /* 圆形头像 */
  margin-right: 10px;
}
.message-content {
  /* 消息内容样式 */
}
.username {
  /* 用户名样式 */
}
.comment {
  /* 评论样式 */
}
.like-action {
  /* 点赞动作样式 */
}
.date {
  /* 日期样式 */
}
</style>