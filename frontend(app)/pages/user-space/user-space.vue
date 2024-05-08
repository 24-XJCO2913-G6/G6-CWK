<template>
	<view>
		<!-- #ifdef MP -->
		<uni-nav-bar :shadow="false" :fixed="true" :border="false" right-text="Menu"
			@click-right="clickNavigationButton"></uni-nav-bar>
		<!-- #endif -->
		<!-- 头部 -->
		<view class="flex align-center p-3 border-bottom border-light-secondary">
			<image :src="userinfo.userpic" style="width: 180rpx;height: 180rpx;" class="rounded-circle"></image>
			<view class="pl-3 flex flex-column flex-1">
				<view class="flex align-center">
					<view class="flex-1 flex flex-column align-center justify-center" v-for="(item,index) in counts"
						:key="index">
						<text class="font-lg font-weight-bold">{{item.num|formatNum}}</text>
						<text class="font text-muted">{{item.name}}</text>
					</view>
				</view>
				<view class="flex justify-center align-center">

					<button v-if="user_id == user.id" type="default" size="mini" style="width: 400rpx;"
						@click="openUserInfo">
						Edit Profile
					</button>

					<button v-else type="default" size="mini"
						:class="userinfo.isFollow ? 'bg-light text-dark' : 'bg-main'" style="width: 400rpx;"
						@click="doFollow">
						{{userinfo.isFollow ? '' : 'Follow'}}
					</button>

				</view>
			</view>
		</view>

		<!-- tab -->
		<view class="flex align-center" style="height: 100rpx;">
			<view class="flex-1 flex align-center justify-center" v-for="(item,index) in tabBars" :key="index"
				:class="index === tabIndex ? 'font-lg font-weight-bold text-main':'font-md'" @click="changeTab(index)">
				{{item.name}}
			</view>
		</view>

		<template v-if="tabIndex === 0">
			<view class="animated fast fadeIn">
				<view class="p-3 border-bottom">
					<view class="font-md">Information</view>
					<view class="font">Age：{{userinfo.regtime}}</view>
					<view class="font">id：{{user_id}}</view>
				</view>
				<view class="p-3 border-bottom">
					<view class="font-md">Information</view>
					<view class="font">Sex: Female</view>
					<view class="font">Birthday：{{userinfo.birthday}}</view>
					<view class="font">Job：{{userinfo.job}}</view>
					<view class="font">Hometown：{{userinfo.path}}</view>
					<view class="font">Relationship Status：{{userinfo.qg}}</view>
				</view>
			</view>
		</template>
		<template v-else>
			<view class="animated fast fadeIn">
				<!-- <common-list :item="myitem" v-for="myitem in posts" :key="myitem.post_id"></common-list> -->
				<common-list :item="post" v-for="post in posts" :key="post.post_id"></common-list>
				<divider></divider>
				<!-- <load-more :loadmore="loadmore"></load-more> -->
			</view>
		</template>



		<!-- 弹出层 -->
		<uni-popup ref="popup" type="top">
			<view class="flex align-center justify-center font-md border-bottom py-2" hover-class="bg-light"
				@click="doBlack">
				<text class="iconfont icon-sousuo mr-2"></text>
				{{userinfo.isblack ? 'de-blacklist' : 'blacklist'}}
			</view>
			<view v-if="!userinfo.isblack" class="flex align-center justify-center font-md py-2" hover-class="bg-light"
				@click="openChat">
				<text class="iconfont icon-shanchu mr-2"></text> Talk
			</view>
		</uni-popup>


	</view>
</template>

<script>
	const emotionArray = ['secrecy', 'unmarried', 'married']
	import commonList from '@/components/common/common-list.vue';
	import loadMore from '@/components/common/load-more.vue';
	import uniPopup from '@/components/uni-ui/uni-popup/uni-popup.vue';
	import $T from '@/common/time.js';
	import {
		mapState
	} from 'vuex'
	import uniNavBar from '@/components/uni-ui/uni-nav-bar/uni-nav-bar.vue';
	export default {

		components: {
			commonList,
			loadMore,
			uniPopup,
			uniNavBar
		},
		data() {
			return {
				posts: [{
					post_id: 1,
					name: 'Yodo',
					visibility: 1,
					time: '2024-05-01',
					like: 19,
					userpic: '../../static/user_pic/1.jpg',
					title: 'Nice day to have a walk',
					collect: 13,
					content: '44444444444444444',
					comments_num: 0,
					center: [
						103.981051,
						30.765612
					],
					path: [
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
						],

					],
					comments: [

						{
							userpic: '../../static/user_pic/2.jpg',
							username: 'Vero',
							data: 'Good!',
							time: '2024-4-7'
						},

					],
				}],
				user_id: 9,
				userinfo: {
					userpic: "/static/user_pic/1.jpg",
					username: "",
					sex: 0,
					age: 20,
					isFollow: true,
					regtime: "20",
					birthday: "2003-08-17",
					job: "Software Engineer",
					path: "Chengdu",
					qg: "None"
				},
				counts: [{
					name: "Post",
					num: 1
				}, {
					name: "Follow",
					num: 0
				}, {
					name: "Fan",
					num: 1
				}],
				tabIndex: 0,
				tabBars: [{
					name: "Info",
				}, {
					name: "Post",
					list: [

					],
					// 1.上拉加载更多  2.加载中... 3.没有更多了
					loadmore: "Load more",
					page: 1
				}],
			}
		},
		onNavigationBarButtonTap() {
			this.clickNavigationButton()
		},
		computed: {
			...mapState({
				user: state => state.user
			}),
			list() {
				return this.tabBars[this.tabIndex].list
			},
			loadmore() {
				return this.tabBars[this.tabIndex].loadmore
			}
		},
		filters: {
			formatNum(value) {
				return value > 99 ? '99+' : value;
			}
		},
		onLoad(e) {
		
			this.user_id = e.user_id
			// 加载用户个人信息
			this.getUserInfo()
			// 获取用户相关数据
			this.getCounts()
			// 监听关注和顶踩操作
			uni.$on('updateFollowOrSupport', (e) => {
				switch (e.type) {
					case 'follow': // 关注
						this.follow(e.data.user_id)
						break;
					case 'support': // 顶踩
						this.doSupport(e.data)
						break;
				}
			})
			// 监听评论数变化
			uni.$on('updateCommentsCount', ({
				id,
				count
			}) => {
				this.tabBars.forEach(tab => {
					if (tab.list) {
						tab.list.forEach((item) => {
							if (item.id === id) {
								item.comment_count = count
							}
						})
					}
				})
			})
		},
		onUnload() {
			uni.$off('updateFollowOrSupport', (e) => {})
			uni.$off('updateCommentsCount', (e) => {})
		},
		methods: {
			clickNavigationButton() {
				if (this.user_id == this.user.id) {
					return uni.navigateTo({
						url: '../user-set/user-set',
					});
				}
				this.$refs.popup.open()
			},
	
			changeTab(index) {
				this.tabIndex = index
				this.getList()
			},

			// 加入/移出黑名单
			doBlack() {
				this.checkAuth(() => {
					let url = this.userinfo.isblack ? '/removeblack' : '/addblack'
					let msg = this.userinfo.isblack ? 'de-blacklist' : 'blacklist'
					uni.showModal({
						content: '是否要' + msg,
						success: (res) => {
							if (res.confirm) {
								this.$H.post(url, {
									id: this.user_id
								}, {
									token: true
								}).then(res => {
									this.userinfo.isblack = !this.userinfo.isblack
									uni.showToast({
										title: msg + 'Success',
										icon: 'none'
									});
								})
							}
						}
					});

				})
			},
			openChat() {
				let user = {
					user_id: this.user_id,
					username: this.userinfo.username,
					userpic: this.userinfo.userpic
				}
				this.navigateTo({
					url: "/pages/user-chat/user-chat?user=" + JSON.stringify(user)
				})
			}
		}
	}
</script>

<style>

</style>