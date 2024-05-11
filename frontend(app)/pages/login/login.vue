<template>
	<view>
		<uni-status-bar></uni-status-bar>
		<view>
			<view class="iconfont icon-guanbi flex align-center justify-center font-lg"
				style="width: 100rpx;height: 100rpx;" hover-class="bg-light" @click="back"></view>
		</view>

		<view class="text-center" style="padding-top: 130rpx;padding-bottom: 70rpx;font-size: 55rpx;">
			{{ forget ? 'Forgot Password' : (status ? 'Sign up' : 'Sign in') }}
		</view>

		<view class="px-2">
			<view v-if="!forget && !status" class="mb-2">
				<input type="text" v-model="email" placeholder="Enter Email" class="border-bottom p-2" />
			</view>
			<view v-if="!forget && !status" class="mb-2 flex align-stretch">
				<input type="password" v-model="password" placeholder="Password" class="border-bottom p-2 flex-1" />
			</view>

			<view v-if="!forget && status" class="mb-2 flex align-stretch">
				<input type="text" v-model="email" placeholder="Enter Email" class="border-bottom p-2 flex-1" />
			</view>
			<view v-if="!forget && status" class="mb-2 flex align-stretch">
				<input type="password" v-model="password" placeholder="Password" class="border-bottom p-2 flex-1" />
			</view>
			<view v-if="!forget && status" class="mb-2 flex align-stretch">
				<input type="text" v-model="code" placeholder="Verification Code" class="border-bottom p-2 flex-1" />
				<view v-if="!forget && status" class="border-bottom flex align-center justify-center font-sm text-white"
					:class="codeTime > 0 ? 'bg-main-disabled' : 'bg-main'" style="width: 180rpx;" @click="getCode">
					{{codeTime > 0 ? codeTime + ' s' : 'Get Code'}}
				</view>
			</view>

			<view v-if="forget" class="mb-2 flex align-stretch">
				<input type="text" v-model="email" placeholder="Email" class="border-bottom p-2 flex-1" />
			</view>
			<view v-if="forget" class="mb-2 flex align-stretch">
				<input type="password" v-model="password" placeholder="New Password" class="border-bottom p-2 flex-1" />
			</view>
			<view v-if="forget" class="mb-2 flex align-stretch">
				<input type="text" v-model="code" placeholder="Verification Code" class="border-bottom p-2 flex-1" />
				<view v-if="forget" class="border-bottom flex align-center justify-center font-sm text-white"
					:class="codeTime > 0 ? 'bg-main-disabled' : 'bg-main'" style="width: 180rpx;" @click="getCode">
					{{codeTime > 0 ? codeTime + ' s' : 'Get Code'}}
				</view>
			</view>
		</view>

		<view class="py-2 px-3">
			<button class="text-white" style="border-radius: 50rpx;border: 0;" type="primary" @click="submit"
				:loading="loading">{{ loading ? 'Processing...' : (forget ? 'Reset Password' : (status ? 'Sign up' : 'Sign in')) }}</button>
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
				email: "",
				password: "",
				phone: "",
				code: "",
				codeTime: 0,
				loading: false,
				showModal: false
			}
		},
		onLoad() {

			var regExp = /^\w{5,10}@\w+\.(com|net|org)$/;
			var flag = regExp.test(email);

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
			getCode() {
				if (this.codeTime > 0) return;

				uni.request({
					url: 'http://120.46.81.37:8080/requestCode',
					method: 'POST',
					data: {
						email: this.email,
					},
			header: {
				    'Content-Type': 'application/x-www-form-urlencoded',
					'aToken': this.aToken,
					'rToken':this.rToken,
				},
					success: (res) => {
						// 倒计时60秒
						this.codeTime = 60;
						let interval = setInterval(() => {
							if (this.codeTime <= 0) {
								clearInterval(interval);
								this.codeTime = 0;
							} else {
								this.codeTime--;
							}
						}, 1000);
					},
					fail: (err) => {
						// console.log(this.email)
						uni.showToast({
							title: 'The verification code request failed. ',
							icon: 'none',
						});
					}
				});
			},

			// 提交
			// submit() {
			// 	let regExp = /^\w{5,10}@\w+\.(com|net|org)$/;
			// 	let flag = regExp.test(this.email);
			// 	regExp.test(flag);
			// 	console.log(flag)
			// 	console.log(this.email, this.password);
			// 	uni.request({
			// 		url: 'http://120.46.81.37:8080/login',
			// 		method: 'POST',
			// 		data: {
			// 			email: this.email,
			// 			password: this.password,
			// 		},
			// 		header: {
			// 			'Content-Type': 'application/x-www-form-urlencoded',
			// 		},
			// 		success: (res) => {

			// 			let data = res.data;
			// 			let aToken = data.aToken;
			// 			let rToken = data.rToken;
			// 			console.log(aToken, rToken);


			// 		},
			// 		fail: (err) => {
			// 			console.error('Failed to fetch posts:', err);
			// 		}
			// 	});

			// }
			submit() {
				if (!this.forget) {
					// 检查邮箱格式
					let regExp = /^\w{5,10}@\w+\.(com|net|org)$/;
					let flag = regExp.test(this.email);
					if (!flag) {
						uni.showToast({
							title: 'Invalid email format.',
							icon: 'none',
						});
						return;
					}
					console.log(this.email);
					console.log(this.password);
					console.log(this.status)

					// 登录或注册
					this.loading = true;
					uni.request({
						url: this.status ? 'http://120.46.81.37:8080/register' :
							'http://120.46.81.37:8080/login',
						method: 'POST',
						data: {
							email: this.email,
							password: this.password,
						},
						header: {
							'Content-Type': 'application/x-www-form-urlencoded',
						},
						success: (res) => {
							this.loading = false;
							// 处理成功逻辑，例如保存token等
							let data = res.data;
							let aToken = data.aToken;
							let rToken = data.rToken;
							this.$store.commit('setTokens', { aToken, rToken });
							
							// 根据实际情况进行后续操作，例如跳转到另一个页面
						},
						fail: (err) => {
							this.loading = false;
							// 处理失败逻辑
							uni.showToast({
								title: 'Sign in/Sign up failed.',
								icon: 'none',
							});
						}
					});
				} else {
					this.loading = true;
					console.log(this.email);
					console.log(this.password);
					console.log(this.code)
					uni.request({
						url: 'http://120.46.81.37:8080/app/resetPassword',
						method: 'POST',
						data: {
							email: this.email,
							newPassword: this.password,
							verificationCode: this.code,
						},
						header: {
							'Content-Type': 'application/x-www-form-urlencoded',
						},
						success: (res) => {
							this.loading = false;
							// 处理重置密码成功的逻辑
							uni.showToast({
								title: 'Password reset successful.',
							});
							// 清空表单数据，返回登录界面等
							this.initForm();
							this.forget = false;
						},
						fail: (err) => {
							this.loading = false;
							// 处理重置密码失败的逻辑
							uni.showToast({
								title: 'Password reset failed.',
								icon: 'none',
							});
						}
					});

				}
			}
		}
	}
</script>

<style>
	/* 样式 */
</style>