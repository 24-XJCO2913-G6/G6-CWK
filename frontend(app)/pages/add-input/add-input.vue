<template>
	<view>
		<!-- 自定义导航 -->
		<!-- #ifdef MP -->

		<!-- #endif -->
		<!-- #ifndef MP -->
		<uni-nav-bar left-icon="back" statusBar :border="false" @click-left="goBack">

		</uni-nav-bar>
		<!-- #endif -->
		<!-- 文本域 -->
		<mymap  :path="path" :center='center' :zoom='16' :mapheight='80'></mymap>
		<br><br>


		<!-- Fixed框 -->
		<view style="position: fixed; bottom: 0; width: 100%; height: 100px; background-color: #eaeff5; padding: 12px;">
			<view v-if="!isRecording" class="flex justify-between align-center">
				<!-- Options for cycling, running, and driving -->
				<view class="option" @click="selectMode('cycling')">Cycling</view>
				<view class="option" @click="selectMode('running')">Running</view>
				<view class="option" @click="selectMode('driving')">Driving</view>

				<!-- 开始行程按钮 -->
				<button @click="startRecording" style="background-color: #4ba3c7; color: #fff;">Start</button>
			</view>
			<view v-else class="flex justify-between align-center">
				<view style="flex: 1;">
					<text style="color: #555;">里程: {{distance}} km</text>
				</view>
				<view style="flex: 1;">
					<text style="color: #555;">时间: {{formattedTime}} </text>
				</view>
				<!-- <view class="info">
					<text class="info-label" style="display: block;">Distance:</text>
					<br> 
					<text class="info-value" style="display: block;">{{ distance }} km</text>
				</view>
				<view class="info">
					<text class="info-label" style="display: block;">Time:</text>
					<br> 
					<text class="info-value" style="display: block;">{{ formattedTime }} </text>
				</view> -->
				<view style="flex: 1;">
					<!-- <button @click="toggleRecording"
						:style="{ backgroundColor: isPaused ? '#33cc33' : '#4ba3c7', color: '#fff' }">{{ isPaused ? 'Resume' : 'Pause' }}</button> -->
					<image @tap="toggleRecording"
						:src="isPaused ?  '/static/images/resume.png':'/static/images/pause.png'" mode="aspectFit"
						style="width: 50px; height: 50px;background-color: #eaeff5; " >
					</image>
				</view>
				<view style="flex: 1;">
					<button @click="stopRecording" style="background-color: #4ba3c7; color: #fff;">End</button>
				</view>
				<!-- <view class="button-group">
					<button @click="toggleRecording"
						:style="{ backgroundColor: isPaused ? '#33cc33' : '#4ba3c7', color: '#fff' }">{{ isPaused ? 'Resume' : 'Pause' }}</button>
					<button @click="stopRecording" style="background-color: #4ba3c7; color: #fff;">End Journey</button>
				</view> -->
			</view>
		</view>


		<view>
			<!-- Countdown display -->
			<div class='mybox'>
				<div v-if="showflag" class="countdown">{{currentCountdown==0?"Go":currentCountdown}}</div>
			</div>
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
				mydisabled:true,
				path: [
					[103.985895, 30.763873], // 起点坐标
					[103.986895, 30.764873], // 中间点坐标
					[103.987895, 30.765873] // 终点坐标
				],
				center: [103.985895, 30.763873],


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
				return this.isOpenArray[this.isopen]
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

			startTimer() {
				this.showflag = true;
				let t = setInterval(() => {
					this.currentCountdown--;
					if (this.currentCountdown <= -1) {
						this.showflag = false
						this.currentCountdown = 3
						clearInterval(t)
						this.mydisabled=false;
					}

				}, 1000)
			},

			startRecording() {
			
				this.startTimer()
				this.isRecording = true;
				console.log('你好')
				setTimeout(() => {
					this.timer = setInterval(() => {
						this.time++;
						console.log(this.time)
					}, 1000); // Update the current time every minute (adjust as needed)

				}, 4000)

				// You can add logic to start recording distance and time here
			},
			// Method to pause the recording
			pauseRecording() {
				// Clear the existing timer to pause the recording
				clearInterval(this.timer);
				// Update the state to indicate that recording is paused
				// You might want to store this state to resume recording later
				this.isRecording = false;
			},
			// Method to toggle recording (pause/resume)
			toggleRecording() {
				if(!this.mydisabled){
				if (this.isPaused) {
					// If paused, resume the timer
					this.timer = setInterval(() => {
						this.time++;
					}, 1000);
				} else {
					// If not paused, pause the timer
					clearInterval(this.timer);
				}
				// Toggle the paused state
				this.isPaused = !this.isPaused;
			}},
			stopRecording() {
			
				this.isRecording = false;
				this.time = 0;
				clearInterval(this.timer);
				// You can add logic to stop recording distance and time here
				// Once stopped, you can update the distance and time values accordingly
			},
			// 发布
			submit() {
				// if(this.post_class_id == 0){
				// 	return uni.showToast({
				// 		title: '请选择分类',
				// 		icon: 'none'
				// 	});
				// }
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