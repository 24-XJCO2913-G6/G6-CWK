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
		<!-- 文本域 -->
		<mymap3 :path="path" :center='center' :zoom='16' :mapheight='30' ></mymap3>
		<view class="font-md px-2 py-1" style="display: flex; flex-wrap: wrap;">
			<view style="margin-bottom: 10px; flex-basis: 100%;">
				<h4>Route Selection:</h4>

				<!-- <picker class='custom-picker' style="font-size: 14px; background-color: white; border:none;"
				          :value='selectedRouteIndex' :range="routeOptions" @change="onRouteChange">
				    {{ routeOptions[selectedRouteIndex] }}
				  </picker> -->
				  <picker class='custom-picker' style="font-size: 14px; background-color: white; border:none;"
				            :value='selectedRouteIndex' :range="routeOptions" @change="onRouteChange" >
				      <!-- 用户选择的路线名称将通过 selectedRouteName 计算属性显示 -->
					  {{ routeOptions[selectedRouteIndex] }}

				    </picker>

	
			</view>
		</view>
		<view class="font-md px-2 py-1" style="display: flex; flex-wrap: wrap;">
					<view style="margin-bottom: 10px; flex-basis: 100%;">
						<h4>Mode Selection:</h4>
							<div class="buttons">
							<button @click="setActive('running')" :class="{ active: activeMode === 'running' }">Running</button>
							<button @click="setActive('cycling')" :class="{ active: activeMode === 'cycling' }">Cycling</button>
							<button @click="setActive('driving')" :class="{ active: activeMode === 'driving' }">Driving</button>
							</div>
				
			
					</view>
				</view>
		<h4 style='margin-left:10px;margin-top: 20px;'>Title:</h4>
		<input v-model="title" placeholder="Title" class="uni-textarea px-2"
			style='margin:5px; width:90%; height:40px;' />
		<h4 style='margin-left:10px'>Desciption:</h4>
		<textarea v-model="content" placeholder="Say something" class="uni-textarea px-2"
			style='margin:10px;width:90%' />

		<!-- 选中的话题 -->
		<!-- <view class="font-md px-2 py-1 flex">
			<view class="border px-3 py main-color main-border-color flex align-center justify-center" style="border-radius: 50rpx;">
				<text class="iconfont icon-huati mr-1"></text>
				{{topic.title ? "所属话题："+topic.title : "Choose topic"}}
			</view>
		</view> -->

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
						name: 'School walk',
						time: '15:00',
						center:[
    103.990729,
    30.762043
  ],
						path:[
  [
    103.986915,
    30.759078
  ],
  [
    103.986957,
    30.759221
  ],
  [
    103.98704,
    30.759346
  ],
  [
    103.986936,
    30.759578
  ],
  [
    103.986873,
    30.759792
  ],
  [
    103.986811,
    30.760024
  ],
  [
    103.986915,
    30.760167
  ],
  [
    103.987206,
    30.760453
  ],
  [
    103.987372,
    30.760667
  ],
  [
    103.987518,
    30.760775
  ],
  [
    103.987663,
    30.760846
  ],
  [
    103.987871,
    30.760882
  ],
  [
    103.988412,
    30.761007
  ],
  [
    103.98889,
    30.761025
  ],
  [
    103.989243,
    30.761096
  ],
  [
    103.989472,
    30.761096
  ],
  [
    103.9897,
    30.761275
  ],
  [
    103.989971,
    30.761382
  ],
  [
    103.99022,
    30.761489
  ],
  [
    103.990397,
    30.761596
  ],
  [
    103.990563,
    30.761811
  ],
  [
    103.990729,
    30.761882
  ],
  [
    103.990729,
    30.762043
  ],
  [
    103.990688,
    30.762168
  ],
  [
    103.990459,
    30.762471
  ],
  [
    103.990355,
    30.762632
  ],
  [
    103.990168,
    30.762739
  ],
  [
    103.989981,
    30.762918
  ],
  [
    103.989919,
    30.763079
  ],
  [
    103.989815,
    30.763204
  ],
  [
    103.989732,
    30.76324
  ],
  [
    103.989482,
    30.763311
  ],
  [
    103.989254,
    30.763311
  ],
  [
    103.989211,
    30.763394
  ],
  [
    103.98892,
    30.76334
  ],
  [
    103.988754,
    30.76334
  ],
  [
    103.988567,
    30.763287
  ],
  [
    103.988338,
    30.763198
  ],
  [
    103.988109,
    30.76309
  ],
  [
    103.987881,
    30.763001
  ],
  [
    103.987735,
    30.762823
  ]
]
						
					},
					{	
						route_id:2,
						name: 'Lake',
						time: '10:39',
						path:[
  [
    103.976838,
    30.768159
  ],
  [
    103.976871,
    30.768094
  ],
  [
    103.976905,
    30.768044
  ],
  [
    103.976921,
    30.768023
  ],
  [
    103.976938,
    30.768001
  ],
  [
    103.97703,
    30.768073
  ],
  [
    103.977097,
    30.768109
  ],
  [
    103.977155,
    30.768166
  ],
  [
    103.977222,
    30.768195
  ],
  [
    103.977339,
    30.76831
  ],
  [
    103.977439,
    30.768382
  ],
  [
    103.977548,
    30.768439
  ],
  [
    103.977673,
    30.768511
  ],
  [
    103.977799,
    30.768597
  ],
  [
    103.977983,
    30.768726
  ],
  [
    103.978108,
    30.768827
  ],
  [
    103.97825,
    30.76892
  ],
  [
    103.9784,
    30.769042
  ],
  [
    103.978526,
    30.769143
  ],
  [
    103.978651,
    30.76925
  ],
  [
    103.978801,
    30.769372
  ],
  [
    103.978969,
    30.769523
  ],
  [
    103.979127,
    30.769652
  ],
  [
    103.979303,
    30.769439
  ],
  [
    103.979345,
    30.769381
  ],
  [
    103.979445,
    30.769252
  ],
  [
    103.979545,
    30.769108
  ],
  [
    103.97962,
    30.769008
  ],
  [
    103.979796,
    30.768843
  ],
  [
    103.97998,
    30.768606
  ],
  [
    103.980113,
    30.768448
  ],
  [
    103.980205,
    30.768276
  ],
  [
    103.980222,
    30.768204
  ],
  [
    103.980272,
    30.768211
  ],
  [
    103.980356,
    30.768247
  ],
  [
    103.980447,
    30.768326
  ],
  [
    103.980573,
    30.76839
  ],
  [
    103.980673,
    30.768441
  ],
  [
    103.980715,
    30.768484
  ],
  [
    103.980665,
    30.768577
  ],
  [
    103.98059,
    30.768721
  ],
  [
    103.980498,
    30.768879
  ],
  [
    103.980297,
    30.76908
  ],
  [
    103.980105,
    30.769302
  ],
  [
    103.98003,
    30.769431
  ],
  [
    103.979971,
    30.769525
  ],
  [
    103.979946,
    30.769661
  ],
  [
    103.979896,
    30.769733
  ],
  [
    103.979904,
    30.76979
  ],
  [
    103.980038,
    30.769776
  ],
  [
    103.980172,
    30.76969
  ],
  [
    103.980272,
    30.769546
  ],
  [
    103.980272,
    30.76941
  ],
  [
    103.980305,
    30.769302
  ],
  [
    103.980322,
    30.769209
  ],
  [
    103.980322,
    30.769159
  ],
  [
    103.980389,
    30.769116
  ],
  [
    103.980506,
    30.769015
  ],
  [
    103.980581,
    30.768979
  ],
  [
    103.980706,
    30.7689
  ],
  [
    103.980757,
    30.768871
  ],
  [
    103.980832,
    30.768814
  ],
  [
    103.98094,
    30.768771
  ],
  [
    103.981007,
    30.768713
  ],
  [
    103.981074,
    30.76867
  ],
  [
    103.981183,
    30.768584
  ],
  [
    103.981258,
    30.76852
  ],
  [
    103.981283,
    30.768469
  ],
  [
    103.981308,
    30.768426
  ]
],
						center: [
    103.980322,
    30.769209
  ],
					},
					{	
						route_id:3,
						name: 'Go to the playground',
						time: '11:00 ',
						path:[
  [
    103.977527,
    30.767054
  ],
  [
    103.977643,
    30.766825
  ],
  [
    103.977727,
    30.766768
  ],
  [
    103.977826,
    30.766625
  ],
  [
    103.977926,
    30.766525
  ],
  [
    103.978076,
    30.766325
  ],
  [
    103.978375,
    30.765926
  ],
  [
    103.978558,
    30.76574
  ],
  [
    103.978741,
    30.765526
  ],
  [
    103.978874,
    30.765297
  ],
  [
    103.978973,
    30.765269
  ],
  [
    103.97904,
    30.765255
  ],
  [
    103.979106,
    30.765312
  ],
  [
    103.979273,
    30.765355
  ],
  [
    103.979406,
    30.765483
  ],
  [
    103.979672,
    30.765583
  ],
  [
    103.979788,
    30.765669
  ],
  [
    103.979938,
    30.765769
  ],
  [
    103.980154,
    30.765869
  ],
  [
    103.980337,
    30.766012
  ],
  [
    103.980486,
    30.766097
  ],
  [
    103.980619,
    30.766169
  ],
  [
    103.980669,
    30.766169
  ],
  [
    103.980736,
    30.766026
  ],
  [
    103.980785,
    30.76584
  ],
  [
    103.980835,
    30.765697
  ],
  [
    103.981051,
    30.765612
  ],
  [
    103.981168,
    30.765569
  ],
  [
    103.981317,
    30.765555
  ],
  [
    103.981517,
    30.765555
  ],
  [
    103.98175,
    30.765497
  ],
  [
    103.981999,
    30.76534
  ],
  [
    103.982165,
    30.765155
  ],
  [
    103.982282,
    30.764969
  ],
  [
    103.982365,
    30.764812
  ]
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
				return this.routes.map(route => route.name+'     '+route.time);
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
		
			uni.request({
			    url: 'http://120.46.81.37:8080/app/get_routes',
				header: {
					'aToken': aToken,
					'rToken':rToken,
				},
			  success: (res) => {
				this.routes = res.data;
			  },
			  fail: (err) => {
				console.error('Failed to fetch posts:', err);
			  }
			    });
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
					post_mode:this.activeMode,
					imgs:this.imageList,
				}
				console.log('提交的字典',dict)
				console.log('提交了')
				 uni.request({
							url: 'http://120.46.81.37:8080/app/upload_post',
				            method: 'POST',
				            data: {
								post:dict,
								'aToken': aToken,
								'rToken':rToken,
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