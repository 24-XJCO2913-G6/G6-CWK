<template>
	<view class="flex align-center p-2 border-bottom border-light-secondary animated fast fadeIn" hover-class="bg-light" @click="open">
		<image :src="item.avatar ? item.avatar : '/static/default.jpg'"
		style="height: 80rpx;width: 80rpx;"
		class="rounded-circle mr-2 bg-secondary"></image>
		
		<view class="flex flex-column flex-1">
			<view class="flex align-center justify-between">
				<text class="font-md">{{item.username}}</text>
				<text class="font-sm text-secondary">
				{{item.update_time | formatTime}}</text>
			</view>
			<view class="flex align-center justify-between">
				<text class="text-secondary text-ellipsis" 
				style="max-width: 500rpx;">{{item.data}}</text>
				<uni-badge :text="item.noread" type="error"></uni-badge>
			</view>
		</view>
		
	</view>
</template>

<script>
	import $T from '@/common/time.js';
	import uniBadge from '@/components/uni-ui/uni-badge/uni-badge.vue'
	export default {
		components: {
			uniBadge
		},
		props:{
			item:Object,
			index:Number
		},
		// 过滤器
		filters: {
			formatTime(value) {
				return $T.gettime(value);
			}
		},
		methods: {
			// 打开聊天页
			open() {
				let user = {
					user_id:this.item.user_id,
					username:this.item.username,
					userpic:this.item.avatar
				}
				uni.navigateTo({
					url: '../../pages/user-chat/user-chat?user='+JSON.stringify(user),
					success: () => {
						// 清除未读数，更新底部导航消息数显示
						this.$store.dispatch('readChatMessage',this.item)
					}
				});
			}
		},
	}
</script>

<style>
</style>
