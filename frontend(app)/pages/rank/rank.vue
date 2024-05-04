<template>  
  <view class="rank-page">  
  <mymap2 :path="path" :center='center' :zoom='16' :mapheight='25' :routes='routes'></mymap2>
    <view class="rank-title">Mileage Ranking</view>  
    <view class="rank-list">  
      <view v-for="(item, index) in rankList" :key="index" class="rank-item">  
        <view class="rank-info">  
          <view class="rank-number">{{ index + 1 }}</view>  
		  <image :src="item.img" mode="aspectFill" class="avatar-image"></image>  
          <view class="rank-name">{{ item.name }}</view>  
        </view>  
        <view class="rank-steps">{{ item.distance }} km</view>  
      </view>  
    </view>  
  </view>  
</template>

<script>  
	import mymap2 from '../../components/map/mymap2.vue';
	
export default {  
	components:{
		mymap2,
	},
  data() {  
    return {  
      rankList: [  
        { name: 'Jack', img: '/static/avatar1.png', distance: 10 },  
        { name: 'Mary', img: '/static/avatar1.png', distance: 8 },  
        { name: 'Jennifer', img: '/static/avatar1.png', distance: 5 },  
        { name: 'Kardashian', img: '/static/avatar1.png', distance: 2 },  
      ]  ,
	  routes:[
		  {
			 path:[
			     [103.985895, 30.763873], // 起点坐标
			     [103.985895, 30.764873], // 中间点坐标
			     [103.987895, 30.765873]  // 终点坐标
			 ],
			 center:[103.985895, 30.763873], 
		  },{
			  path:[
			   [103.983895, 30.763873], // 起点坐标
			   [103.987895, 30.764874], // 中间点坐标
			   [103.988894, 30.763873]  // 终点坐标
			  ],
			 center:[103.987895, 30.764874], 
		  },{
		  path:[
		   [103.982895, 30.763873], // 起点坐标
		   [103.986895, 30.764574], // 中间点坐标
		   [103.987694, 30.765872]  // 终点坐标
		  ],
		   center: [103.986895, 30.764574], },
		   {
		   path:[
		    [103.982895, 30.763873], // 起点坐标
		    [103.986895, 30.765574], // 中间点坐标
		    [103.987694, 30.766872]  // 终点坐标
		   ],
		    center: [103.982895, 30.767574], }
	  ]
    }  
  }  ,
//   mounted(){
// 	  console.log('haha')
// 	  uni.request({
// 	              url: 'http://localhost:8080/app/rank',
// 	              method: 'GET',
// 	              success: (res) => {
// 	                  this.rankList=res.rank_list
	  				
// 	              },
// 	              fail: (err) => {
// 	                  console.error('Error fetching person info:', err);
	  				
// 	              }
//   })
// }  
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