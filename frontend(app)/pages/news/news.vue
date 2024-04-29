<template>
	<view>
		<!-- 自定义导航 -->
		<!-- #ifdef MP -->
		<uni-nav-bar :border="false">
			<view class="flex align-center j·1·	ustify-center w-100" @click="changeIsopen">
				{{isopenText}}<text class="iconfont icon-shezhi"></text>
			</view>
		</uni-nav-bar>
		<!-- #endif -->
		<!-- #ifndef MP -->
		<uni-nav-bar left-icon="back" statusBar :border="false" @click-left="goBack">
			<view class="flex align-center justify-center w-100" @click="changeIsopen">
				{{isopenText}}<text class="iconfont icon-shezhi"></text>
			</view>
		</uni-nav-bar>
		<!-- #endif -->
		<!-- 文本域 -->
		<mymap :path="path" :center='center' :zoom='16' :mapheight='30'></mymap>
		<br><br>




		<input v-model="title" placeholder="Title" class="uni-textarea px-2"
			style='border:1px dotted;margin:5px; width:200px;' />
		<textarea v-model="content" placeholder="Say something" class="uni-textarea px-2"
			style='border:1px dotted;margin:5px' />
		<!-- 选中的分类 -->
		<view class="font-md px-2 py-1 flex">
			<view class="border px-3 py main-color main-border-color flex align-center justify-center"
				style="border-radius: 50rpx;">
				<text class="iconfont icon-huati mr-1"></text>
				{{post_class_name ? "所属分类："+post_class_name : "Choose classification"}}
			</view>
		</view>
		<view class="font-md px-2 py-1">
			<view style="margin-bottom: 10px;">
				<h6>Start date:</h6>
				<picker mode="date" class='custom-picker'
					style="font-size: 14px; border: 1px solid black; margin-right: 10px;" :value='startDate'
					@change="onStartDateChange">{{startDate==''?'select':startDate}}</picker>
				<h6>Start time:</h6>
				<picker mode="time" class='custom-picker' style="font-size: 14px; border: 1px solid black;"
					:value='startTime' @change="onStartTimeChange">{{startTime==''?'select':startTime}}</picker>
			</view>
			<view style="margin-bottom: 10px;">
				<h6>End date:</h6>
				<picker mode="date" class='custom-picker'
					style="font-size: 14px; border: 1px solid black; margin-right: 10px;" :value='endDate'
					@change="onEndDateChange">{{endDate==''?'select':endDate}}</picker>
				<h6>End time:</h6>
				<picker mode="time" class='custom-picker' style="font-size: 14px; border: 1px solid black;"
					:value='endTime' @change="onEndTimeChange">{{endTime==''?'select':endTime}}</picker>
			</view>
		</view>

		<!-- 选中的话题 -->
		<!-- <view class="font-md px-2 py-1 flex">
			<view class="border px-3 py main-color main-border-color flex align-center justify-center" style="border-radius: 50rpx;">
				<text class="iconfont icon-huati mr-1"></text>
				{{topic.title ? "所属话题："+topic.title : "Choose topic"}}
			</view>
		</view> -->

		<!-- 多图上传 -->
		<upload-image :show="show" ref="uploadImage" :list="imageList" @change="changeImage"></upload-image>

		<!-- 底部操作条 -->
		<view class="fixed-bottom bg-white flex align-center" style="height: 85rpx;">
			<picker mode="selector" :range="range" @change="choosePostClass">
				<view class="iconfont icon-caidan footer-btn animated" hover-class="jello"></view>
			</picker>
			<!-- 		
			<view class="iconfont icon-huati footer-btn animated"
			hover-class="jello" @click="chooseTopic"></view> -->
			<view class="iconfont icon-tupian footer-btn animated" hover-class="jello"
				@click="iconClickEvent('uploadImage')"></view>
			<view class="iconfont icon-dizhi footer-btn animated" hover-class="jello" @click=""></view>
			<view class="bg-main text-white ml-auto flex justify-center align-center rounded mr-2 animated"
				hover-class="jello" style="width: 140rpx;height: 60rpx;" @click="submit">Submit</view>
		</view>

	</view>
</template>

<script>
	const isOpenArray = ['Public', 'Only me', "Only friends"];
	import uniNavBar from '@/components/uni-ui/uni-nav-bar/uni-nav-bar.vue';
	import uploadImage from '@/components/common/upload-image.vue';
	import mymap from '../../components/map/mymap.vue';


	export default {
		components: {
			uniNavBar,
			uploadImage,
			mymap
		},
		data() {
			return {
				content: "",
				imageList: [],
				// 是否已经弹出提示框
				showBack: false,
				isopen: 1,
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
				path: [
					[103.985895, 30.763873], // 起点坐标
					[103.986895, 30.764873], // 中间点坐标
					[103.987895, 30.765873] // 终点坐标
				],
				center: [103.985895, 30.763873],


			}
		},
		computed: {
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
				// uni.showActionSheet({
				//     itemList: isOpenArray,
				//     success: (res)=> {
				//         this.isopen = res.tapIndex
				//     }
				// });
				uni.navigateTo({
					url: '../index/index',
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
</style>