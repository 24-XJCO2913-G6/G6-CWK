<template>
	<view style="height: 100rpx;"
	class="fixed-bottom flex align-center border-top bg-white">
	
		<input type="text" v-model="content" class="flex-1 rounded bg-light ml-2" style="padding: 5rpx;" :focus="focus" placeholder="Speak politely" @confirm="submit" @blur="$emit('blur')"/>
		
		<view class="iconfont icon-fabu flex align-center justify-center font-lg animated" hover-class="jello text-main" style="width: 100rpx;" @click="submit"></view>
		
	</view>
</template>

<script>
	import { mapState } from 'vuex'
	export default {
		props: {
			focus: {
				type: Boolean,
				default: false
			},
		},
		data() {
			return {
				content: ""
			}
		},
		computed:{
			...mapState({
				loginStatus:state=>state.loginStatus,
				aToken: state => state.aToken,
				rToken: state=>state.rToken
				
			}),
		},
		methods: {
			submit() {
				// 是否为空
				if (this.content === '') {
					return uni.showToast({
						title: '消息不能为空',
						icon: 'none'
					});
				}
				this.$emit('submit',this.content)
				// 清空输入框
				this.content = ''
			}
		},
	}
</script>

<style>
</style>
