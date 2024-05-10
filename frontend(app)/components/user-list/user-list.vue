<template>
	<view class="p-2 flex align-center border-bottom border-light-secondary" hover-class="bg-light" @click="open">
		<image :src="item.user_pic"
		style="width: 100rpx;height: 100rpx;"
		class="rounded-circle mr-2"></image>
		<view class="flex flex-column flex-1">
			<text class="font-md text-dark">{{item.username}}</text>
		</view> 
		<view class="uni-icon uni-icon-checkbox-filled" 
		:class="item.isFollow ? 'text-light-muted':'text-main'"></view> 
	</view>
</template>

<script>
	import uniBadge from '@/components/uni-ui/uni-badge/uni-badge.vue';
	import { mapState } from 'vuex'
	export default {
		components: {
			uniBadge
		},
		props:{
			item:Object,
			index:Number
		},
		computed:{
			...mapState({
				loginStatus:state=>state.loginStatus,
				aToken: state => state.aToken,
				rToken: state=>state.rToken
				
			}),
		},
		methods: {
			open() {
				uni.navigateTo({
					url: '/pages/user-space/user-space?user_id='+this.item.id,
				});
			}
		},
	}
</script>

<style>
</style>
