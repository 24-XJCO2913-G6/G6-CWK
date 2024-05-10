<template>
	<view class="rank-page">
	
		<map style="width: 100%; height: 300px;" scale='15' :latitude="center[1]" :longitude="center[0]"  :polyline="routes" :markers="markers" >
		 </map>
		<view class="rank-title">Mileage Ranking</view>
		<view class="rank-list">
			<view v-for="(item, index) in rankList" :key="index" class="rank-item">
				<view class="rank-info">
					<view class="rank-number" :style="{'background-color':color_list[index%color_list.length]}">{{ index + 1 }}</view>
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
import { mapState } from 'vuex'
	export default {
		components: {
			mymap2,
		},
		data() {
			return {
				color_list:['green','blue','red','orange','purple'],
				rankList: [{
						name: 'Veroooo',
						img: 'http://120.46.81.37:8080/app/static/user_pic/2.jpg',
						distance: 4.1
					},
					{
						name: 'Yodo',
						img: 'http://120.46.81.37:8080/app/static/user_pic/1.jpg',
						distance: 3.9
					},
					{
						name: 'Pease of Summer',
						img: 'http://120.46.81.37:8080/app/static/user_pic/3.jpg',
						distance: 3.8
					},
				
				],
				markers:[
					{"longitude": 103.986166, "latitude": 30.757523, iconPath:'../../static/images/marker.png', width: 10,
					height: 10}  ,{"longitude": 103.982365, "latitude": 30.764812, iconPath:'../../static/images/marker.png', width: 10,
          height: 10}  , {"longitude": 103.981815, "latitude": 30.768302, iconPath:'../../static/images/marker.png', width: 10,
          height: 10}  
				],
				routes: [{
						points: [
  {"longitude": 103.986193, "latitude": 30.757726},
  {"longitude": 103.986179, "latitude": 30.757617},
  {"longitude": 103.986179, "latitude": 30.757559},
  {"longitude": 103.986206, "latitude": 30.757415},
  {"longitude": 103.98619, "latitude": 30.757272},
  {"longitude": 103.986195, "latitude": 30.757178},
  {"longitude": 103.986202, "latitude": 30.757023},
  {"longitude": 103.986202, "latitude": 30.75693},
  {"longitude": 103.986211, "latitude": 30.756844},
  {"longitude": 103.986247, "latitude": 30.756751},
  {"longitude": 103.986292, "latitude": 30.756662},
  {"longitude": 103.986337, "latitude": 30.756627},
  {"longitude": 103.986408, "latitude": 30.756593},
  {"longitude": 103.98644, "latitude": 30.756565},
  {"longitude": 103.986505, "latitude": 30.756551},
  {"longitude": 103.98661, "latitude": 30.756544},
  {"longitude": 103.986666, "latitude": 30.756551},
  {"longitude": 103.986812, "latitude": 30.756579},
  {"longitude": 103.986812, "latitude": 30.756607},
  {"longitude": 103.98686, "latitude": 30.756641},
  {"longitude": 103.986892, "latitude": 30.756662},
  {"longitude": 103.986933, "latitude": 30.756711},
  {"longitude": 103.986933, "latitude": 30.756746},
  {"longitude": 103.986957, "latitude": 30.756794},
  {"longitude": 103.986965, "latitude": 30.756822},
  {"longitude": 103.986973, "latitude": 30.75687},
  {"longitude": 103.986973, "latitude": 30.756926},
  {"longitude": 103.986965, "latitude": 30.757023},
  {"longitude": 103.986981, "latitude": 30.757134},
  {"longitude": 103.986989, "latitude": 30.757204},
  {"longitude": 103.986981, "latitude": 30.757273},
  {"longitude": 103.986981, "latitude": 30.757349},
  {"longitude": 103.986981, "latitude": 30.757488},
  {"longitude": 103.986981, "latitude": 30.757564},
  {"longitude": 103.986941, "latitude": 30.757655},
  {"longitude": 103.986949, "latitude": 30.757703},
  {"longitude": 103.986941, "latitude": 30.757745},
  {"longitude": 103.986892, "latitude": 30.7578},
  {"longitude": 103.98682, "latitude": 30.757869},
  {"longitude": 103.986771, "latitude": 30.757918},
  {"longitude": 103.98669, "latitude": 30.757918},
  {"longitude": 103.986602, "latitude": 30.757945},
  {"longitude": 103.986529, "latitude": 30.757931},
  {"longitude": 103.986481, "latitude": 30.757931},
  {"longitude": 103.9864, "latitude": 30.757918},
  {"longitude": 103.986351, "latitude": 30.757883},
  {"longitude": 103.986303, "latitude": 30.757841},
  {"longitude": 103.986254, "latitude": 30.7578},
  {"longitude": 103.986206, "latitude": 30.757723},
  {"longitude": 103.986166, "latitude": 30.757523}
],color: "#00aa00",
					width: 10,
					dottedLine: true,
					arrowLine: true,
					colorList: true,
					},
					{
						points:[
  {"longitude": 103.977527, "latitude": 30.767054},
  {"longitude": 103.977643, "latitude": 30.766825},
  {"longitude": 103.977727, "latitude": 30.766768},
  {"longitude": 103.977826, "latitude": 30.766625},
  {"longitude": 103.977926, "latitude": 30.766525},
  {"longitude": 103.978076, "latitude": 30.766325},
  {"longitude": 103.978375, "latitude": 30.765926},
  {"longitude": 103.978558, "latitude": 30.76574},
  {"longitude": 103.978741, "latitude": 30.765526},
  {"longitude": 103.978874, "latitude": 30.765297},
  {"longitude": 103.978973, "latitude": 30.765269},
  {"longitude": 103.97904, "latitude": 30.765255},
  {"longitude": 103.979106, "latitude": 30.765312},
  {"longitude": 103.979273, "latitude": 30.765355},
  {"longitude": 103.979406, "latitude": 30.765483},
  {"longitude": 103.979672, "latitude": 30.765583},
  {"longitude": 103.979788, "latitude": 30.765669},
  {"longitude": 103.979938, "latitude": 30.765769},
  {"longitude": 103.980154, "latitude": 30.765869},
  {"longitude": 103.980337, "latitude": 30.766012},
  {"longitude": 103.980486, "latitude": 30.766097},
  {"longitude": 103.980619, "latitude": 30.766169},
  {"longitude": 103.980669, "latitude": 30.766169},
  {"longitude": 103.980736, "latitude": 30.766026},
  {"longitude": 103.980785, "latitude": 30.76584},
  {"longitude": 103.980835, "latitude": 30.765697},
  {"longitude": 103.981051, "latitude": 30.765612},
  {"longitude": 103.981168, "latitude": 30.765569},
  {"longitude": 103.981317, "latitude": 30.765555},
  {"longitude": 103.981517, "latitude": 30.765555},
  {"longitude": 103.98175, "latitude": 30.765497},
  {"longitude": 103.981999, "latitude": 30.76534},
  {"longitude": 103.982165, "latitude": 30.765155},
  {"longitude": 103.982282, "latitude": 30.764969},
  {"longitude": 103.982365, "latitude": 30.764812}
],color: "#0000ff",
					width: 10,
					dottedLine: true,
					arrowLine: true,
					colorList: true,
			
					},
					{
						points:[
  {"longitude": 103.984276, "latitude": 30.768216},
  {"longitude": 103.984082, "latitude": 30.768406},
  {"longitude": 103.983949, "latitude": 30.76852},
  {"longitude": 103.983872, "latitude": 30.768615},
  {"longitude": 103.983684, "latitude": 30.768767},
  {"longitude": 103.983529, "latitude": 30.768891},
  {"longitude": 103.983462, "latitude": 30.768995},
  {"longitude": 103.983308, "latitude": 30.769176},
  {"longitude": 103.983263, "latitude": 30.769185},
  {"longitude": 103.98312, "latitude": 30.7691},
  {"longitude": 103.982998, "latitude": 30.768967},
  {"longitude": 103.982832, "latitude": 30.768853},
  {"longitude": 103.982511, "latitude": 30.768625},
  {"longitude": 103.982368, "latitude": 30.768511},
  {"longitude": 103.982312, "latitude": 30.768492},
  {"longitude": 103.982279, "latitude": 30.768492},
  {"longitude": 103.982235, "latitude": 30.768596},
  {"longitude": 103.982191, "latitude": 30.768663},
  {"longitude": 103.982158, "latitude": 30.768653},
  {"longitude": 103.981947, "latitude": 30.768587},
  {"longitude": 103.981925, "latitude": 30.76853},
  {"longitude": 103.981815, "latitude": 30.768463},
  {"longitude": 103.981793, "latitude": 30.768397},
  {"longitude": 103.981815, "latitude": 30.768302}
],color: "#ff0000",
					width: 10,
					dottedLine: true,
					arrowLine: true,
					colorList: true,

					}
					]
				}
			},
			computed:{
				
				center(){
					let longitude=0;
					let latitude=0;
					let sum_lon=0;
					let sum_lat=0;
					let points_num=0;
					this.routes.forEach(t=>{
						sum_lon+=t.points.reduce((sum,x)=>sum+x.longitude,0)
						sum_lat+=t.points.reduce((sum,x)=>sum+x.latitude,0)
						points_num+=t.points.length;
					})
					
					
					longitude=sum_lon/points_num;
					latitude=sum_lat/points_num;
					return [longitude,latitude]
					
				}
			},
			methods:{
				getColor(index) {
				      // 确保 index 在颜色数组的范围内
				      index = index % this.colorList.length;
				      return this.colorList[index];
				    }
			},
			  onLoad(){
				  console.log('haha')
				
				  uni.request({
				              url: 'http://120.46.81.37:8080/app/rank',
				              method: 'GET',
				              success: (res) => {
				                  this.rankList=res.rank_list
								  this.routes=res.routes

				              },
				              fail: (err) => {
				                  console.error('Error fetching person info:', err);

				              }
			  })
			}  ,
		
		}
</script>

<style scoped>
	.rank-page {
		padding: 20rpx;
		background-color: #f0f3f5;
		/* 淡灰色背景 */
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
		border-top: 1rpx solid #d9d9d9;
		/* 浅灰色分割线 */
		border-bottom: 1rpx solid #d9d9d9;
		/* 浅灰色分割线 */
	}

	.rank-number {
		width: 70rpx;
		height: 70rpx;
		line-height: 70rpx;
		/* 垂直居中 */
		text-align: center;
		font-size: 40rpx;
		font-weight: bold;
		color: #fff;
		background-color: skyblue;
		/* 蓝色背景 */
		border-radius: 50%;
		/* 圆形 */
		margin-right: 10rpx;
	}

	.rank-info {
		display: flex;
		align-items: center;
		flex: 1;
		/* 占用剩余空间 */
	}

	.rank-name {
		font-size: 40rpx;
		font-weight: bold;
		color: #333;
		/* 深灰色名字 */
	}

	.rank-steps {
		font-size: 35rpx;
		color: black;
		/* 蓝色步数 */
		text-align: right;
	}


	.avatar-image {
		width: 80rpx;
		/* 设置一个固定的宽度，你可以根据需要调整 */
		height: 80rpx;
		/* 如果你想保持正方形头像，也设置一个高度 */
		overflow: hidden;
		/* 确保图片不会超出容器 */
		border-radius: 50%;
		/* 保持图片为圆形 */
		margin-right: 20rpx;
		/* 右边距，用于与排名信息分隔 */

	}
</style>