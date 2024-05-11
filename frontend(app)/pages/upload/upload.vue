<template>
	<view class="mybody" style="height: 100vh;">
		<!-- 自定义导航 -->
			<!-- #ifdef MP -->
				<uni-nav-bar :border="false">
					<view class="flex align-center j·1·	ustify-center w-100" @click="changeIsopen">
						{{post_type}}<text class="iconfont icon-shezhi"></text>
					</view>
				</uni-nav-bar>
				<!-- #endif -->
				<!-- #ifndef MP -->
				<uni-nav-bar left-icon="back" statusBar :border="false" @click-left="goBack">
					<view class="flex align-center justify-center w-100" @click="changeIsopen">
						{{post_type}}<text class="iconfont icon-shezhi"></text>
					</view>
				</uni-nav-bar>
				<!-- #endif -->
	<map style="width: 100%; height: 200px;" scale='16' :latitude="center[1]" :longitude="center[0]"  :polyline="path"  >
		 </map>
		<view class="font-md px-2 py-1" style="display: flex; flex-wrap: wrap;">
			<view style="margin-bottom: 10px; flex-basis: 100%;">
				<h4>Route Selection:</h4>

				    <picker class='custom-picker' style="font-size: 14px; background-color: white; border:none;"
				            :value='selectedRouteIndex' :range="routeOptions" @change="onRouteChange" >
				      <!-- 用户选择的路线名称将通过 selectedRouteName 计算属性显示 -->
					  {{ routeOptions[selectedRouteIndex] }}

				    </picker>

	
			</view>
		</view>

		<h4 style='margin-left:10px;margin-top: 20px;'>Title:</h4>
		<input v-model="title" placeholder="Title" class="uni-textarea px-2"
			style='margin:5px; width:90%; height:40px;' />
		<h4 style='margin-left:10px'>Desciption:</h4>
		<textarea v-model="content" placeholder="Say something" class="uni-textarea px-2"
			style='margin:10px;width:90%' />


		<!-- 多图上传 -->
		<upload-image :show="show" ref="uploadImage" :list="imageList" @change="changeImage"></upload-image>
		<br><br>
		<!-- 底部操作条 -->
		<view class="fixed-bottom bg-white flex align-center" style="height: 85rpx;">
	
			<!-- 		
			<view class="iconfont icon-huati footer-btn animated"
			hover-class="jello" @click="chooseTopic"></view> -->
			<view class="iconfont icon-tupian footer-btn animated" hover-class="jello"
				@click="iconClickEvent('uploadImage')"></view>

			<view class="bg-main text-white ml-auto flex justify-center align-center rounded mr-2 animated"
				hover-class="jello" style="width: 140rpx;height: 60rpx;" @click="submit">Submit</view>
		</view>

	</view>
</template>

<script>
	import uniNavBar from '@/components/uni-ui/uni-nav-bar/uni-nav-bar.vue';
	import uploadImage from '@/components/common/upload-image.vue';
	import mymap3 from '../../components/map/mymap3.vue';
import { mapState } from 'vuex'

	export default {
		components: {
			uniNavBar,
			uploadImage,
			mymap3
		},
		data() {
			return {
				isOpenArray : ['Public', 'VIP only'],
				activeMode:'running',
				selectedRoute: '',
				selected_route_id:1,
				selectedRouteIndex: 0, // 用于picker的value属性，存储选中的索引
				routes: [{
						route_id:9,
						end_date: '2024-5-9',
						end_time:'11:55',
						center:[
    103.990729,
    30.762043
  ],
						path:[
  {points:[
  {"longitude": 103.986915, "latitude": 30.759078},
  {"longitude": 103.986957, "latitude": 30.759221},
  {"longitude": 103.98704, "latitude": 30.759346},
  {"longitude": 103.986936, "latitude": 30.759578},
  {"longitude": 103.986873, "latitude": 30.759792},
  {"longitude": 103.986811, "latitude": 30.760024},
  {"longitude": 103.986915, "latitude": 30.760167},
  {"longitude": 103.987206, "latitude": 30.760453},
  {"longitude": 103.987372, "latitude": 30.760667},
  {"longitude": 103.987518, "latitude": 30.760775},
  {"longitude": 103.987663, "latitude": 30.760846},
  {"longitude": 103.987871, "latitude": 30.760882},
  {"longitude": 103.988412, "latitude": 30.761007},
  {"longitude": 103.98889, "latitude": 30.761025},
  {"longitude": 103.989243, "latitude": 30.761096},
  {"longitude": 103.989472, "latitude": 30.761096},
  {"longitude": 103.9897, "latitude": 30.761275},
  {"longitude": 103.989971, "latitude": 30.761382},
  {"longitude": 103.99022, "latitude": 30.761489},
  {"longitude": 103.990397, "latitude": 30.761596},
  {"longitude": 103.990563, "latitude": 30.761811},
  {"longitude": 103.990729, "latitude": 30.761882},
  {"longitude": 103.990729, "latitude": 30.762043},
  {"longitude": 103.990688, "latitude": 30.762168},
  {"longitude": 103.990459, "latitude": 30.762471},
  {"longitude": 103.990355, "latitude": 30.762632},
  {"longitude": 103.990168, "latitude": 30.762739},
  {"longitude": 103.989981, "latitude": 30.762918},
  {"longitude": 103.989919, "latitude": 30.763079},
  {"longitude": 103.989815, "latitude": 30.763204},
  {"longitude": 103.989732, "latitude": 30.76324},
  {"longitude": 103.989482, "latitude": 30.763311},
  {"longitude": 103.989254, "latitude": 30.763311},
  {"longitude": 103.989211, "latitude": 30.763394},
  {"longitude": 103.98892, "latitude": 30.76334},
  {"longitude": 103.988754, "latitude": 30.76334},
  {"longitude": 103.988567, "latitude": 30.763287},
  {"longitude": 103.988338, "latitude": 30.763198},
  {"longitude": 103.988109, "latitude": 30.76309},
  {"longitude": 103.987881, "latitude": 30.763001},
  {"longitude": 103.987735, "latitude": 30.762823}
],color: "#2e7efd",
					width: 10,
					dottedLine: true,
					arrowLine: true,
					colorList: true,}
],markers:[
						    {"longitude": 103.987735, "latitude": 30.762823, iconPath:'../../static/images/marker.png', width: 10,
          height: 10}  
					],
						
					},
					{	
						route_id:2,
						end_date: '2024-5-9',
						end_time:'2:55',
						path:[
							{points:[
  {"longitude": 103.976838, "latitude": 30.768159},
  {"longitude": 103.976871, "latitude": 30.768094},
  {"longitude": 103.976905, "latitude": 30.768044},
  {"longitude": 103.976921, "latitude": 30.768023},
  {"longitude": 103.976938, "latitude": 30.768001},
  {"longitude": 103.97703, "latitude": 30.768073},
  {"longitude": 103.977097, "latitude": 30.768109},
  {"longitude": 103.977155, "latitude": 30.768166},
  {"longitude": 103.977222, "latitude": 30.768195},
  {"longitude": 103.977339, "latitude": 30.76831},
  {"longitude": 103.977439, "latitude": 30.768382},
  {"longitude": 103.977548, "latitude": 30.768439},
  {"longitude": 103.977673, "latitude": 30.768511},
  {"longitude": 103.977799, "latitude": 30.768597},
  {"longitude": 103.977983, "latitude": 30.768726},
  {"longitude": 103.978108, "latitude": 30.768827},
  {"longitude": 103.97825, "latitude": 30.76892},
  {"longitude": 103.9784, "latitude": 30.769042},
  {"longitude": 103.978526, "latitude": 30.769143},
  {"longitude": 103.978651, "latitude": 30.76925},
  {"longitude": 103.978801, "latitude": 30.769372},
  {"longitude": 103.978969, "latitude": 30.769523},
  {"longitude": 103.979127, "latitude": 30.769652},
  {"longitude": 103.979303, "latitude": 30.769439},
  {"longitude": 103.979345, "latitude": 30.769381},
  {"longitude": 103.979445, "latitude": 30.769252},
  {"longitude": 103.979545, "latitude": 30.769108},
  {"longitude": 103.97962, "latitude": 30.769008},
  {"longitude": 103.979796, "latitude": 30.768843},
  {"longitude": 103.97998, "latitude": 30.768606},
  {"longitude": 103.980113, "latitude": 30.768448},
  {"longitude": 103.980205, "latitude": 30.768276},
  {"longitude": 103.980222, "latitude": 30.768204},
  {"longitude": 103.980272, "latitude": 30.768211},
  {"longitude": 103.980356, "latitude": 30.768247},
  {"longitude": 103.980447, "latitude": 30.768326},
  {"longitude": 103.980573, "latitude": 30.76839},
  {"longitude": 103.980673, "latitude": 30.768441},
  {"longitude": 103.980715, "latitude": 30.768484},
  {"longitude": 103.980665, "latitude": 30.768577},
  {"longitude": 103.98059, "latitude": 30.768721},
  {"longitude": 103.980498, "latitude": 30.768879},
  {"longitude": 103.980297, "latitude": 30.76908},
  {"longitude": 103.980105, "latitude": 30.769302},
  {"longitude": 103.98003, "latitude": 30.769431},
  {"longitude": 103.979971, "latitude": 30.769525},
  {"longitude": 103.979946, "latitude": 30.769661},
  {"longitude": 103.979896, "latitude": 30.769733},
  {"longitude": 103.979904, "latitude": 30.76979},
  {"longitude": 103.980038, "latitude": 30.769776},
  {"longitude": 103.980172, "latitude": 30.76969},
  {"longitude": 103.980272, "latitude": 30.769546},
  {"longitude": 103.980272, "latitude": 30.76941},
  {"longitude": 103.980305, "latitude": 30.769302},
  {"longitude": 103.980322, "latitude": 30.769209},
  {"longitude": 103.980322, "latitude": 30.769159},
  {"longitude": 103.980389, "latitude": 30.769116},
  {"longitude": 103.980506, "latitude": 30.769015},
  {"longitude": 103.980581, "latitude": 30.768979},
  {"longitude": 103.980706, "latitude": 30.7689},
  {"longitude": 103.980757, "latitude": 30.768871},
  {"longitude": 103.980832, "latitude": 30.768814},
  {"longitude": 103.98094, "latitude": 30.768771},
  {"longitude": 103.981007, "latitude": 30.768713},
  {"longitude": 103.981074, "latitude": 30.76867},
  {"longitude": 103.981183, "latitude": 30.768584},
  {"longitude": 103.981258, "latitude": 30.76852},
  {"longitude": 103.981283, "latitude": 30.768469},
  {"longitude": 103.981308, "latitude": 30.768426}
],color: "#2e7efd",
					width: 10,
					dottedLine: true,
					arrowLine: true,
					colorList: true,

}
  
],markers:[
						    {"longitude": 103.981308, "latitude": 30.768426, iconPath:'../../static/images/marker.png', width: 10,
          height: 10}  
					],
						center: [
    103.980322,
    30.769209
  ],
					},
					{	
						route_id:3,
						end_date: '2024-5-9',
						end_time:'14:55',
						path:[{points:[
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
],color: "#2e7efd",
					width: 10,
					dottedLine: true,
					arrowLine: true,
					colorList: true,
}],
markers:[
						    {"longitude": 103.982365, "latitude": 30.764812, iconPath:'../../static/images/marker.png', width: 10,
          height: 10}  
					],

						 center:   [
    103.981317,
    30.765555
  ],
					},
					// ... 其他路线
				],

				imageList: [],
				// 是否已经弹出提示框
				showBack: false,
				isopen: 1,
				topic: {
					id: 0,
					title: ""
				},
				title: '',
				content: "",
				startDate: '',
				startTime: '',
				endTime: '',
				endDate: '',
				path: [
					
				],
				center: [],
				isopen:0
			


			}
		},
		computed: {
			post_type() {
							return this.isOpenArray[this.isopen]
						},
			show() {
				return this.imageList.length > 0
			},
			imgListIds() {
				return this.imageList.map(item => {
					return {
						id: item.id
					}
				})
			},
			selectedRouteName() {
				return this.routes[this.selectedRouteIndex] ? this.routes[this.selectedRouteIndex].name : '';
			},
			// 将路线名称作为picker的选项显示
			routeOptions() {
				return this.routes.map(route => route.end_date+'     '+route.end_time);
			}
		},
		// 监听返回
		onBackPress() {
			if ((this.content !== '' || this.imageList.length > 0) && !this.showBack) {
				uni.showModal({
					content: 'Do you want to save as a draft?',
					showCancel: true,
					cancelText: "Don't save",
					confirmText: 'Save',
					success: res => {
						// 点击确认
						if (res.confirm) {
							this.store()
						} else { // 点击取消，清除缓存
							uni.removeStorage({
								key: "add-input"
							})
						}
						// 手动执行返回
						uni.navigateBack({
							delta: 1
						});
					},
				});
				this.showBack = true
				return true
			}
		},
		// 页面加载时
		onLoad() {
			this.path=this.routes[0].path;
			this.center=this.routes[0].center;
			this.selected_route_id=this.routes[0].route_id;
		
			// uni.request({
			// url: 'http://120.46.81.37:8080/app/blogPublishTrack',
			// 	data: {
			// 		'aToken': this.aToken,
			// 		'rToken':this.rToken,
			// 	},
			//   success: (res) => {
			// 	this.routes = res.data;
			//   },
			//   fail: (err) => {
			// 	console.error('Failed to fetch posts:', err);
			//   }
			//     });
			
		
			uni.getStorage({
				key: "add-input",
				success: (res) => {
					if (res.data) {
						let result = JSON.parse(res.data)
						this.content = result.content
						this.imageList = result.imageList
					}
				}
			})
			
			
			// 监听选择话题事件
			uni.$on('chooseTopic', (e) => {
				this.topic.id = e.id
				this.topic.title = e.title
			})
			// 获取所有分类
			this.getPostClass()
		},
		beforeDestroy() {
			uni.$off('chooseTopic', (e) => {})
		},
		methods: {
			setActive(str){
					console.log('ha');
					this.activeMode=str;
			},
			// 发布
			submit() {
			
				let dict={
					post_type:this.post_type=='Public'?0:1,
					title:this.title,
					content:this.content,
					route_id:this.selected_route_id,
					// imgs:this.imageList,
				}
				console.log('提交的字典',dict)
				console.log('提交了')
				 uni.request({
							url: 'http://120.46.81.37:8080/app/upload_post',
				            method: 'POST',
				            data: {
								post:dict,
								'aToken':this.aToken,
								'rToken':this.rToken,
							},
				            header: {
				                'Content-Type': 'application/x-www-form-urlencoded',
								
				            },
				            success: (res) => {
				                console.log('Post data uploaded successfully:', res);
				            },
				            fail: (err) => {
				                console.error('Error uploading post data:', err);
				            }
				        });
			
			},
			// 获取所有文章分类
			getPostClass() {
				this.$H.get('/postclass').then(res => {
					this.post_class_list = res.list
				})
			},
			onRouteChange(event) {
				// 更新选中的路线
				this.selectedRouteIndex = event.target.value;
				console.log('index',this.selectedRouteIndex,this.routes)
				this.center=this.routes[this.selectedRouteIndex].center;
				this.path=this.routes[this.selectedRouteIndex].path;
				this.selected_route_id=this.routes[this.selectedRouteIndex].route_id;
				
				console.log('中心，path',this.center,this.path)
				// 这里可以添加更多的逻辑处理
			},
			// 选择文章分类
			choosePostClass(e) {
				this.post_class_index = e.detail.value
			},
			// 选择话题
			chooseTopic() {
				uni.navigateTo({
					url: '../topic-nav/topic-nav?choose=true',
				});
			},
			// 切换可见性
			changeIsopen() {
				uni.showActionSheet({
					itemList: this.isOpenArray,
					success: (res) => {
						this.isopen = res.tapIndex
					}
				});
			},
			
			// 底部图片点击事件
			iconClickEvent(e) {
				switch (e) {
					case 'uploadImage':
						this.$refs.uploadImage.chooseImage()
						break;
				}
			},
			// 返回上一步
			goBack() {
				// uni.navigateBack({ delta: 1 });
				uni.reLaunch({
					url: '/pages/index/index' // 首页的路径
				});
			},
			// 选中图片
			changeImage(e) {
				this.imageList = e
			},
			onStartDateChange(e) {
				this.startDate = e.detail.value
			},
			onStartTimeChange(e) {
				this.startTime = e.detail.value
			},
			onEndDateChange(e) {
				this.endDate = e.detail.value
			},
			onEndTimeChange(e) {
				this.endTime = e.detail.value
			},
			// 保存操作
			store() {
				// 保存为本地存储
				let obj = {
					content: this.content,
					imageList: this.imageList
				}
				uni.setStorage({
					key: 'add-input',
					data: JSON.stringify(obj)
				})
			}
		}
	}
</script>

<style>
	.mybody {
		background-color: #EFF4F5;
	}

	.footer-btn {
		width: 86rpx;
		height: 86rpx;
		display: flex;
		justify-content: center;
		align-content: center;
		font-size: 50rpx;
	}

	.custom-picker {
		font-size: 14px;
		border: 1px solid #333;
		/* Change border color to a darker shade */
		border-radius: 5px;
		/* Add border radius for rounded corners */
		padding: 5px;
		/* Add padding inside the picker */
		margin-right: 10px;
		/* Spacing between pickers */
		background-color: #f2f2f2;
		/* Set a light background color */
	}
	.mode-selection {
	  display: flex;
	  flex-direction: column;
	  align-items: center;
	}
	.buttons {
	  display: flex;
	}
	button {

	  margin: 10px;
	  border: none;
	  background-color: white;
	}
	button.active {
	  background-color:#4ba3c7;
	  color:white;
	}
</style>