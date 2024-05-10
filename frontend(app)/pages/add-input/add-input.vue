<template>
	<view>
		<!-- #ifndef MP -->
		<uni-nav-bar left-icon="back" statusBar :border="false" @click-left="goBack">

		</uni-nav-bar>
		<!-- #endif -->
        <view class="page-body">
			<view class="page-section page-section-gap">
			    <map style="width: 100%; height: 675px;" :path="path" :latitude="latitude" :longitude="longitude" :markers="covers" :polyline="polyline">
			    </map>
			</view> 
		</view>
		

		<!-- Fixed框 -->
		<view style="position: fixed; bottom: -50px; width: 100%; height: 100px; background-color: #eaeff5; padding: 12px;">
			<view v-if="!isRecording" class="flex justify-between align-center">
				<!-- Options for cycling, running, and driving -->
				<view class="option" @click="selectMode('cycling')">Cycling</view>
				<view class="option" @click="selectMode('running')">Running</view>
				<view class="option" @click="selectMode('driving')">Driving</view>
				<!-- 开始行程按钮 -->
				
				<button @click="startRecording" style="background-color: #4ba3c7; color: #fff;">Start</button>

				<button @click="resetMap" style="background-color: #4ba3c7; color: #fff;">Reset</button>

			</view>
			<view v-else class="flex justify-between align-center">
				<view style="flex: 1;">
					<text style="color: #555;">里程: {{distance}} km</text>
				</view>
				<view style="flex: 1;">
					<text style="color: #555;">时间: {{formattedTime}} </text>
				</view>
				

				<view style="flex: 1;">
					<button @click="stopRecording" style="background-color: #4ba3c7; color: #fff;">End</button>
				</view>

			</view>
			<view>
				<!-- Countdown display -->
				<div class='mybox'>
					<div v-if="showflag" class="countdown">{{currentCountdown==0?"Go":currentCountdown}}</div>
				</div>
			</view>
		</view>
		
	</view>
</template>


<script>
	const isOpenArray = ['Public', 'Only me', "Only friends"];
	import uniNavBar from '@/components/uni-ui/uni-nav-bar/uni-nav-bar.vue';
	import uploadImage from '@/components/common/upload-image.vue';
	import mymap from '../../components/map/mymap3.vue';

	export default {
		components: {
			uniNavBar,
			uploadImage,
			mymap
		},

		data() {
			return {
				isOpenArray :['Public', 'Only VIP'],
				showflag: false,
				currentCountdown: 3,
				isRecording: false,
				distance: 0,
				time: 0,
				timer: null,
				isPaused: false, // Track whether the timer is paused
				content: "",
				imageList: [],
				// 是否已经弹出提示框
				showBack: false,
				isopen: 0,
				topic: {
					id: 0,
					title: ""
				},
				title: '',
				post_class_list: [],
				post_class_index: -1,
				startDate: '',
				startTime: '',
				endTime: '',
				endDate: '',
				id: 0, // 使用 marker 点击事件需要填写 id
				title: 'map',
				latitude: 39.909,
				longitude: 116.39742,
				covers: [
				      {
				        latitude: 39.909, // 当前位置的纬度
				        longitude: 116.39742, // 当前位置的经度
				        iconPath: '../../../static/location.png' // 当前位置的图标路径
				      },
				      {
				        latitude: 39.90, // 起点的纬度
				        longitude: 116.39, // 起点的经度
				        iconPath: '../../../static/start.png' // 起点的图标路径
				      }
				    ],
				path: [],
				polyline: [{ //初始值
						arrowLine: true,
						color: '#1E90FF', 
						width: 20,
						points: [
							
						],
					}]
			}
		},

		computed: {
			formattedTime() {
				const hours = Math.floor(this.time / 3600);
				const minutes = Math.floor((this.time % 3600) / 60);
				const seconds = this.time % 60;

				return `${hours}:${minutes}:${seconds}`;
			},
			show() {
				return this.imageList.length > 0
			},
			isopenText() {
				return isOpenArray[this.isopen]
			},
			// 文章分类可选项
			range() {
				return this.post_class_list.map(item => {
					return item.classname
				})
			},
			post_class_id() {
				if (this.post_class_index !== -1) {
					return this.post_class_list[this.post_class_index].id
				}
			},
			post_class_name() {
				if (this.post_class_index !== -1) {
					return this.post_class_list[this.post_class_index].classname
				}
			},
			imgListIds() {
				return this.imageList.map(item => {
					return {
						id: item.id
					}
				})
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
						} else { 
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
			uni.getLocation({
			    type: 'wgs84',
			    success: (res) => {
			        console.log('当前位置的经度：' + res.longitude);
			        console.log('当前位置的纬度：' + res.latitude);
					this.covers[1].latitude = res.latitude - 0.0022;
					this.covers[1].longitude = res.longitude + 0.0022;
					 
			    },
			    fail: (err) => {
			        console.error('获取位置信息失败：', err);
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
			}),
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
		mounted() {
		    // 初始化地图上下文对象
		    this.mapContext = uni.createMapContext('myMap');
		    // 启动路径更新定时器
		    this.updatePath();
		},
		methods: {    	
				updatePath() {
					        uni.getLocation({
					            type: 'wgs84',
					            success: (res) => {
									console.log('当前位置的经度：' + res.longitude);
									console.log('当前位置的纬度：' + res.latitude);
									// 更新经纬度信息 
									this.longitude = res.longitude + 0.001916;
									this.latitude = res.latitude - 0.00246;
					                // 更新当前位置坐标
					                this.covers[0].latitude = res.latitude - 0.00246;
					                this.covers[0].longitude = res.longitude + 0.001916;
								
					                // 添加当前位置坐标到路径
					                this.polyline[0].points.push({latitude: res.latitude - 0.00246, longitude: res.longitude + 0.001916});
								},
					            fail: (err) => {
					                console.error('获取位置信息失败：', err);
					            }
					        });
					},
				


			
			selectMode(mode) {
			    switch (mode) {
			        case 'cycling':
			            // 设置更新路径的间隔时间
			            this.intervalTime = 5000;
			            break;
			        case 'running':
			            // 设置更新路径的间隔时间
			            this.intervalTime = 8000;
			            break;
			        case 'driving':
			            // 设置更新路径的间隔时间
			            this.intervalTime = 3000;
			            break;
			        default:
			            break;
			    }
				this.selectedMode = mode; 
			},
			
			
			
			
			startRecording() {
			    if (!this.selectedMode) {
			        uni.showToast({
			            title: 'Please select a mode first',
			            icon: 'none'
			        });
			        return; // 不执行开始记录的操作
			    }
			
			    if (!this.intervalTime) {
			        // 如果未选择模式，则不执行任何操作
			        return;
			    }
			
			    this.startTimer(); // 启动倒计时
			    this.isRecording = true; // 开始记录状态
			
			    // 启动定时器，每秒递增 time 并更新路径
			    this.timer = setInterval(() => {
			        this.updatePath();
			        this.time++;
			    }, 1000);
			
			    // 根据所选模式的间隔时间调用 updatePath() 函数
			    setTimeout(() => {
			        this.updatePath();
			    }, this.intervalTime);
			},

			
			startTimer() {
				this.showflag = true;
				let t = setInterval(() => {
					this.currentCountdown--;
					if (this.currentCountdown <= -1) {
						this.showflag = false
						this.currentCountdown = 3
						clearInterval(t)
					}
			
				}, 1000)
			},

			
			stopRecording() {
				this.isRecording = false;
				this.time = 0;
				clearInterval(this.timer);
				// You can add logic to stop recording distance and time here
				// Once stopped, you can update the distance and time values accordingly
			}, 
			


			resetMap() {
			    // 清空路径
			    this.path = [];
			    // 清空折线
			    this.polyline = [{
			        arrowLine: true,
			        color: '#1E90FF',
			        width: 20,
			        points: [],
			    }];
			    // 重置距离
			    this.distance = 0;
			    // 重置时间
			    this.time = 0;
			},


			// 发布
			submit() {
				uni.showLoading({
					title: 'Submitting...',
					mask: false
				});
				this.$H.post('/post/create', {
					imglist: JSON.stringify(this.imgListIds),
					text: this.content,
					isopen: this.isopen,
					post_class_id: this.post_class_id
				}, {
					token: true
				}).then(res => {
					uni.hideLoading()
					uni.$emit('updateIndex')
					uni.showToast({
						title: 'Submit successfully!',
						icon: "none"
					});
					this.showBack = true
					uni.navigateBack({
						delta: 1
					});
				}).catch(err => {
					uni.hideLoading()
				})
			},
			// 获取所有文章分类
			getPostClass() {
				this.$H.get('/postclass').then(res => {
					this.post_class_list = res.list
				})
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
					itemList: isOpenArray,
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
	/* Custom styles for options */
	.option {
		padding: 8px 12px;
		background-color: #fff;
		border-radius: 20px;
		margin-right: 10px;
		color: #4ba3c7;
		font-size: 14px;
		cursor: pointer;
		transition: background-color 0.3s, color 0.3s;
	}

	.button-group {
		display: flex;
		align-items: center;
	}

	.button-group button {
		margin-right: 10px;
	}

	.option:hover {
		background-color: #4ba3c7;
		color: #fff;
	}

	.info {
		display: flex;
		align-items: center;
		margin-right: 20px;
	}

	.info-label {
		color: #666;
		font-size: 14px;
	}

	.info-value {
		margin-left: 5px;
		font-size: 14px;
		white-space: pre-wrap;
		/* Allow line breaks */
		color: #333;
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

	.mybox {
		display: flex;
		justify-content: center;
	}

	.countdown {
		width: 180px;
		height: 180px;
		border-radius: 100%;
		background-color: rgba(70, 130, 180, 0.8);
		color: white;
		position: fixed;
		top: 30%;
		text-align: center;
		line-height: 180px;
		font-size: 120px;
		font-weight: bold;

	}
</style>