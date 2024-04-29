<template>  
  <view class="rank-page">  
    <view class="rank-title">Mileage Ranking</view>  
    <view class="rank-list">  
      <view v-for="(item, index) in rankList" :key="index" class="rank-item">  
        <view class="rank-info">  
          <view class="rank-number">{{ index + 1 }}</view>  
		  <image :src="item.avatar" mode="aspectFill" class="avatar-image"></image>  
          <view class="rank-name">{{ item.name }}</view>  
        </view>  
        <view class="rank-steps">{{ item.steps }} km</view>  
      </view>  
    </view>  
  </view>  
</template>

<script>  
export default {  
  data() {  
    return {  
      rankList: [  
        { name: 'Jack', avatar: '/static/avatar1.png', steps: 10 },  
        { name: 'Mary', avatar: '/static/avatar1.png', steps: 8 },  
        { name: 'Jennifer', avatar: '/static/avatar1.png', steps: 5 },  
        { name: 'Kardashian', avatar: '/static/avatar1.png', steps: 2 },  
      ]  
    }  
  }  
}  
</script>  
  
<style scoped>  
.rank-page {  
  padding: 20rpx;  
  background-color: #f0f3f5; /* 淡灰色背景 */  
}  
  
.rank-title {  
  text-align: center;  
  font-size: 36rpx;  
  font-weight: bold;  
  color: black;
  margin-bottom: 30rpx;  
}  
  
.rank-list {  
  display: flex;  
  flex-direction: column;  
}  
  
.rank-item {  
  display: flex;  
  align-items: center;  
  justify-content: space-between;  
  padding: 20rpx 0;  
  border-top: 1rpx solid #d9d9d9; /* 浅灰色分割线 */  
  border-bottom: 1rpx solid #d9d9d9; /* 浅灰色分割线 */  
}  
  
.rank-number {  
  width: 70rpx;  
  height: 70rpx;  
  line-height: 70rpx; /* 垂直居中 */  
  text-align: center;  
  font-size: 40rpx;  
  font-weight: bold;  
  color: #fff;  
  background-color: skyblue;/* 蓝色背景 */  
  border-radius: 50%; /* 圆形 */  
  margin-right: 10rpx;
}  
  
.rank-info {  
  display: flex;  
  align-items: center;  
  flex: 1; /* 占用剩余空间 */  
}  
  
.rank-name {  
  font-size: 40rpx;  
  font-weight: bold;  
  color: #333; /* 深灰色名字 */  
}  
  
.rank-steps {  
  font-size: 35rpx;  
  color: black; /* 蓝色步数 */  
  text-align: right;  
}  


.avatar-image {  
  width: 80rpx; /* 设置一个固定的宽度，你可以根据需要调整 */  
  height: 80rpx; /* 如果你想保持正方形头像，也设置一个高度 */  
  overflow: hidden; /* 确保图片不会超出容器 */  
  border-radius: 50%; /* 保持图片为圆形 */  
  margin-right: 20rpx; /* 右边距，用于与排名信息分隔 */  

}  

</style>