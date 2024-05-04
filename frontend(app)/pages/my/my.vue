<template>
	<view>
		<!-- 未登录 -->
		<template v-if="!loginStatus">
			<view class="flex align-center justify-center py-2 font">
				登录社区，体验更多功能
			</view>
			<other-login></other-login>
			<view class="flex align-center justify-center py-2 font text-secondary" @click="openLogin">
				账号/邮箱/手机登录 <text class="ml-1 iconfont icon-jinru"></text>
			</view>
		</template>

		<!-- 已登录 -->
		<view v-else class="flex align-center p-2" hover-class="bg-light" @click='openSpace'>
			<image :src="user.userpic"
			style="width: 100rpx;height: 100rpx;"
			class="rounded-circle"></image>
			<view class="flex flex-column flex-1 px-2">
				<text class="font-lg font-weight-bold text-dark">{{user.username}}</text>
			</view>
			<text class="iconfont icon-jinru" style="font-weight: bold;"></text>
		</view>
		
		
		
		<view class="px-3 py-2">
			<image src="/static/demo/banner1.jpg" mode="aspectFill"
			style="height: 170rpx;width: 100%;" class="rounded"></image>
		</view>
		
		<uni-list-item title="View Routes" showExtraIcon @click="openRoutes">
			<text slot="icon" class="iconfont icon-liulan"></text>
		</uni-list-item>
		<uni-list-item title="VIP" showExtraIcon @click="openVIP">
			<text slot="icon" class="iconfont icon-huiyuanvip"></text>
		</uni-list-item>
		<uni-list-item title="Rank" showExtraIcon @click="openRank">
			<text slot="icon" class="iconfont icon-user_icon"></text>
		</uni-list-item>
		<uni-list-item title="Badge" showExtraIcon @click="openBadge">
		  <text slot="icon" class="iconfont icon-xuanze"></text>
		</uni-list-item>

		<!-- #ifdef MP -->
		<navigator url="../user-set/user-set" hover-class="none">
		<uni-list-item title="我的设置" showExtraIcon>
			<text slot="icon" class="iconfont icon-shezhi"></text>
		</uni-list-item>
		</navigator>
		<!-- #endif -->
		
		
	</view>
</template>

<script>
	import uniListItem from '@/components/uni-ui/uni-list-item/uni-list-item.vue';
	import otherLogin from '@/components/common/other-login.vue';
	import { mapState } from 'vuex'
	export default {
		components: {
			uniListItem,
			otherLogin
		},
		data() {
			return {
				myuser:{
					userpic:'',
					username:''
				},
				post_num:0,
				comments_num:0,
				fan_num:0
			}
		},
		onNavigationBarButtonTap() {
			uni.navigateTo({
				url: '../user-set/user-set'
			});
		},
		computed: {
			...mapState({
				loginStatus:state=>state.loginStatus,
				user:state=>state.user
			}),
			// 用户头像
			avatar(){
				return this.user.userpic ? this.user.userpic : '/static/default.jpg'
			}
		},
		//页面一展示，这个函数就被触发。
		onShow() {
			console.log('haha')
			this.post_num=1
		
			//必须是登录的状态，才会调用里面getCounts方法
			if(this.loginStatus){
				this.getCounts()
			}
		},
		watch: {
			loginStatus(newValue, oldValue) {
				if(newValue){
					this.getCounts()
				} else {
					this.myData.forEach(item=>{
						item.num = 0
					})
				}
			}
		},
		methods: {
			// 获取用户相关数据
			openSpace(){
				uni.navigateTo({
					url: '../my-space/my-space',
				});
			},
			getCounts(){
				console.log('调用了')
			    uni.request({
			                url: 'http://localhost:8080/app/profile',
			                method: 'GET',
			                success: (res) => {
			         
								this.myuser.username=res.user.username;
								this.myuser.userpic=res.user.userpic;
								
			                },
			                fail: (err) => {
								
			                    console.error('Error fetching person info:', err);
								
			                }
			            });
			},
			// 打开登录页
			openLogin(){
				uni.navigateTo({
					url: '../login/login',
				});
			},
			openVIP(){
				uni.navigateTo({
					url: '../vip/vip',
				});
			},
			openBadge(){
				uni.navigateTo({
					url: '../badge/badge'
				});
			},
			openRank(){
				uni.navigateTo({
					url: '../rank/rank'
				});
			},
			openRoutes(){
				uni.navigateTo({
					url: '../routes/routes'
				});
			},
		}
	}
</script>

<style>

</style>
