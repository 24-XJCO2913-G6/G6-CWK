<template>
	<view>
		<uni-status-bar></uni-status-bar>
		<view>
			<view class="iconfont icon-guanbi flex align-center justify-center font-lg"
			style="width: 100rpx;height: 100rpx;" hover-class="bg-light" 
			@click="back"></view>
		</view>
		
		<view class="text-center" style="padding-top: 130rpx;padding-bottom: 70rpx;font-size: 55rpx;">
			{{ forget ? 'Forgot Password' : (status ? 'Sign up' : 'Sign in') }}
        </view>

		<view class="px-2">
			<view v-if="!forget && !status" class="mb-2">
				<input type="text" v-model="username" placeholder="Enter Email" class="border-bottom p-2"/>
			</view>
			<view v-if="!forget && !status" class="mb-2 flex align-stretch">
				<input type="text" v-model="password" placeholder="Password" class="border-bottom p-2 flex-1"/>
			</view>

			<view v-if="!forget && status" class="mb-2 flex align-stretch">
				<input type="text" v-model="username" placeholder="Enter Email" class="border-bottom p-2 flex-1"/>
			</view>
			<view v-if="!forget && status" class="mb-2 flex align-stretch">
				<input type="text" v-model="password" placeholder="Password" class="border-bottom p-2 flex-1"/>
			</view>

			<view v-if="forget" class="mb-2 flex align-stretch">
				<input type="text" v-model="email" placeholder="Email" class="border-bottom p-2 flex-1"/>
			</view>
			<view v-if="forget" class="mb-2 flex align-stretch">
				<input type="text" v-model="password" placeholder="New Password" class="border-bottom p-2 flex-1"/>
			</view>
			<view v-if="forget" class="mb-2 flex align-stretch">
			    <input type="text" v-model="code" placeholder="Verification Code" class="border-bottom p-2 flex-1"/>
				<view v-if="forget" class="border-bottom flex align-center justify-center font-sm text-white" :class="codeTime > 0 ? 'bg-main-disabled' : 'bg-main'" style="width: 180rpx;" @click="getCode">{{codeTime > 0 ? codeTime + ' s' : 'Get Code'}}</view>
			</view>
		</view>
		
		<view class="py-2 px-3">
			<button class="text-white" style="border-radius: 50rpx;border: 0;" type="primary" :disabled="disabled" :class="disabled ? 'bg-main-disabled':'bg-main'" @click="submit" :loading="loading">{{ loading ? 'Processing...' : (forget ? 'Reset Password' : (status ? 'Sign up' : 'Sign in')) }}</button>
		</view>
		
		<view v-if="!forget" class="flex align-center justify-center pt-2 pb-4">
			<view class="text-primary font-sm" @click="changeStatus">
				{{ status ? 'Already have an account? Sign in' : 'Have no account? Sign up' }}
			</view>
		</view>
		<view v-if="!forget" class="flex align-center justify-center pt-2 pb-4">
			<view class="text-primary font-sm" @click="forgotPassword">
				Forgot password?
			</view>
		</view>
		<view v-if="forget" class="flex align-center justify-center pt-2 pb-4">
			<view class="text-primary font-sm" @click="backToLogin">
				Return to Login
			</view>
		</view>
		
		<view class="flex align-center justify-center">
			<view style="height: 1rpx;background-color: #dddddd;width: 100rpx;"></view>
			<view class="mx-2 text-muted">Been</view>
			<view style="height: 1rpx;background-color: #dddddd;width: 100rpx;"></view>
		</view>
		
		<other-login back></other-login>
	</view>
</template>

<script>
	import uniStatusBar from '@/components/uni-ui/uni-status-bar/uni-status-bar.vue';
	import otherLogin from '@/components/common/other-login.vue';
	import UniPopup from '@/components/uni-ui/uni-popup/uni-popup.vue';
	export default {
		components: {
			uniStatusBar,
			otherLogin,
			UniPopup
		},
		data() {
			return {
				status: 0,
				forget: false,
				username: "",
				password: "",
				phone: "",
				code: "",
				codeTime: 0,
				loading: false,
				showModal: false
			}
		},
		onLoad() {
			
		},
		computed: {
			disabled() {
				if (!this.forget) {
					return this.username === '' || this.password === '';
				} else {
					return this.phone === '' || this.code === '';
				}
			}
		},
		methods: {
			toggleModal: function() {
				this.showModal = !this.showModal;
			},
			closeme: function() {
				this.showModal = !this.showModal;
			},
			back() {
				uni.navigateBack({
					delta: 1
				});
			},
			changeStatus() {
				// 初始化表单
				this.initForm()
				this.status = !this.status
			},
			forgotPassword() {
				this.forget = true;
			},
			backToLogin() {
				this.forget = false;
			},
			// 初始化表单
			initForm() {
				this.username = ''
				this.password = ''
				this.phone = ''
				this.code = ''
			},
			// 提交
			submit() {
				if (this.forget) {
					// 处理忘记密码逻辑
				} else {
					// 处理登录/注册逻辑
				}
			}
		}
	}
</script>

<style>
/* 样式 */
</style>
