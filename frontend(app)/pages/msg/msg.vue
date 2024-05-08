<template>
	<view>
		<!-- 导航 -->


		<view class="icon-box">
			<view class="icon-text-group" @click="handleClick('like')">
				<image class="icon" src="../../static/images/like.png"></image>
				<text class="icon-text">Like</text>
			</view>
			<view class="icon-text-group" @click="handleClick('collect')">
				<image class="icon" src="../../static/images/collect.png"></image>
				<text class="icon-text">Collect</text>
			</view>
			<view class="icon-text-group" @click="handleClick('comment')">
				<image class="icon" src="../../static/images/comment.png"></image>
				<text class="icon-text">Comment</text>
			</view>
		</view>

		<template v-if="list.length > 0">
			<!-- 消息列表 -->
			<block v-for="(item,index) in list" :key="index">
				<msg-list :item="item" :index="index"></msg-list>
			</block>
		</template>
		<template v-else>
			<no-thing></no-thing>
		</template>

		<!-- 弹出层 -->
		<uni-popup ref="popup" type="top">
			<view class="flex align-center justify-center font-md border-bottom py-2" hover-class="bg-light"
				@click="popupEvent('friend')">
				<text class="iconfont icon-sousuo mr-2"></text> Add Friends
			</view>
			<view class="flex align-center justify-center font-md py-2" hover-class="bg-light"
				@click="popupEvent('clear')">
				<text class="iconfont icon-shanchu mr-2"></text> Clear List
			</view>
		</uni-popup>

	</view>
</template>

<script>
	import msgList from '@/components/msg/msg-list.vue';
	import noThing from '@/components/common/no-thing.vue';
	import uniPopup from '@/components/uni-ui/uni-popup/uni-popup.vue';
	import uniNavBar from '@/components/uni-ui/uni-nav-bar/uni-nav-bar.vue';
	import {
		mapState
	} from 'vuex'
	export default {
		components: {
			msgList,
			noThing,
			uniPopup,
			uniNavBar
		},
		data() {
			return {

			}
		},
		// 页面加载
		onLoad() {

		},
		computed: {
			...mapState({
				list: state => state.chatList
			})
		},
		// 监听下拉刷新
		onPullDownRefresh() {
			this.refresh()
		},
		// 监听原生导航栏按钮事件
		onNavigationBarButtonTap(e) {
			switch (e.index) {
				case 0: // 左边
					uni.navigateTo({
						url: '../user-list/user-list',
					});
					// 关闭弹出层
					this.$refs.popup.close()
					break;
				case 1: // 右边
					this.$refs.popup.open()
					break;
			}
		},
		methods: {
			// #ifndef APP-PLUS
			clickLeft() {
				this.navigateTo({
					url: '../user-list/user-list',
				})
				this.$refs.popup.close()
			},
			clickRight() {
				this.$refs.popup.open()
			},
			// #endif
			// 下拉刷新
			handleClick(type) {
				switch (type) {
					case 'like':
						uni.navigateTo({
							url: `/pages/like/like` // 替换为点赞消息页面的实际路径
						});
						break;
					case 'collect':
						uni.navigateTo({
							url: `/pages/collect/collect` // 替换为点赞消息页面的实际路径
						});
						break;
					case 'comment':
						uni.navigateTo({
							url: `/pages/comment/comment` // 替换为点赞消息页面的实际路径
						});
						break;
						// 其他case...
				}
			},
			refresh() {
				setTimeout(() => {
					this.list = demo
					// 停止下拉刷新
					uni.stopPullDownRefresh()
				}, 2000)
			},
			// 弹出层选项点击事件
			popupEvent(e) {
				switch (e) {
					case 'friend':
						uni.navigateTo({
							url: '../search/search?type=user',
						});
						break;
					case 'clear':
						this.$store.commit('clearChatList')
						uni.showToast({
							title: 'Clear Successfully',
							icon: 'none'
						});
						console.log('Clear List');
						break;
				}
				// 关闭弹出层
				this.$refs.popup.close()
			}
		}
	}
</script>

<style>
	.icon-box {
		display: flex;
		justify-content: space-around;
		padding: 10px;
	}

	.icon-text-group {
		display: flex;
		flex-direction: column;
		align-items: center;
		/* 垂直居中对齐 */
		cursor: pointer;
		/* 鼠标悬停效果 */
	}

	.icon {
		width: 60px;
		height: 60px;
		background-color: white;
		margin-bottom: 5px;
		/* 图标和文字之间的间距 */
	}

	.icon-text {
		font-size: 14px;
		text-align: center;
	}
</style>