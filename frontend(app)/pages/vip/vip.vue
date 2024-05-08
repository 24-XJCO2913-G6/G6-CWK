<template>

	<view class="container " style='margin-top:30px'>
		<!-- 新增显示会员到期时间的组件 -->
		<view v-if='isVip'>
			<view class="expire-info" v-if="expireDate">
				<text class="expire-label">Membership valid until：</text>
				<text class="expire-value" style="display: inline-block; text-align: center;">{{ expireDate }}</text>
			</view>
		</view>

		<view v-if="isVip" class="cancel-vip">
			<button class="cancel-button" @click="cancelVip">Cancel Membership</button>
		</view>

		<view class="member-row">
			<view class="member-item" @click="handleClick('Month')" :class="{active: activeOption === 'Month'}">
				<view style="font-weight: bold; font-size: 25px;">Month</view>
				<view style=" font-size: 20px;"> $29.9 </view>
			</view>
			<view class="member-item" @click="handleClick('Quarter')" :class="{active: activeOption === 'Quarter'}">
				<view style="font-weight: bold; font-size: 25px;">Quarter</view>
				<view style=" font-size: 20px;"> $69.9 </view>
			</view>
		</view>
		<view class="member-row">
			<view class="member-item" @click="handleClick('Year')" :class="{active: activeOption === 'Year'}">
				<view style="font-weight: bold; font-size: 25px;">Year</view>
				<view style=" font-size: 20px;"> $199.9</view>
			</view>
			<view class="member-item" @click="handleClick('Permanent')" :class="{active: activeOption === 'Permanent'}">
				<view style="font-weight: bold; font-size: 25px;"> Permanent</view>
				<view style=" font-size: 20px;"> $999</view>

			</view>
		</view>
		<view class="payment-block" style='margin-bottom: 40px;margin-top: 30px;'>
			<view style="font-weight: bold; margin-bottom: 20px; font-size:20px">Payment</view>
			<view class="payment-option" @click="selectPayment('Wechat')" style='padding-left:10px ;margin-bottom:15px'>
				<img src="http://120.46.81.37:8080/app/static/common/wechat.png" alt="Wechat Icon"
					style="width: 24px; height: 24px; margin-right: 10px;"> Wechat <span
					v-show="selectedPayment === 'Wechat'"
					style="font-size: 35px;float: right; color: blue; font-weight:bold; margin-right: 20px;">&#10003;</span>
			</view>
			<view class="payment-option" @click="selectPayment('Alipay')" style='padding-left:10px'>
				<img src="http://120.46.81.37:8080/app/static/common/ali.png" alt="Wechat Icon"
					style="width: 24px; height: 24px; margin-right: 10px;"> Alipay <span
					v-show="selectedPayment === 'Alipay'"
					style="font-size: 35px;float: right; color:blue; font-weight:bold; margin-right: 20px;">&#10003;</span>
			</view>
		</view>

		<button class="confirm-button" @click="openPayment">Confirm</button>
		<view style='margin-top: 30px; color: grey; font-size: 25px; text-align: center;'>
			<view style='font-weight:bold; font-size:30px'>What you can get?</view>
			<view>View Vip Posts</view>
			<view>Promote Your Posts for Free</view>
			<view>Personalized setting</view>
		</view>

	</view>
</template>

<style>
	.container {
		display: flex;
		flex-wrap: wrap;
	}

	.member-item {
		width: 100%;
		height: 100px;
		padding: 10px;
		text-align: center;
		background-color: #fff;
		color: #000;
		border: 1px solid #ccc;
		border-radius: 10px;
		margin-bottom: 10px;
	}

	.member-item.active {
		border: 1px solid #ffd700;
		background-color: #ffffe0;
		color: #ffd700;
	}

	.payment-option.active {
		border: 1px solid #ffd700;
		background-color: #ffffe0;
		color: #ffd700;
	}

	.container {
		display: flex;
		flex-direction: column;
	}

	.member-row {
		display: flex;
	}

	.member-item {
		flex: 1;
		padding: 10px;
		text-align: center;
		background-color: #fff;
		color: #000;
		border: 1px solid #ccc;
	}

	.member-item.active {
		border: 2px solid #FFAA33;
		background-color: #ffffe0;
		color: #FFAA33;
	}

	.confirm-button {
		width: 100%;
		padding: 10px;
		text-align: center;
		background-color: #007bff;
		color: #fff;
		border: none;
		margin-top: 10px;
	}

	.payment-option {
		height: 50px;
		border: 1px grey solid;
		border-radius: 10px;
		line-height: 50px;
	}
</style>

<script>
	export default {
		data() {
			return {
				activeOption: null,
				selectedPayment: '',
				// 新增会员到期时间属性
				expireDate: '', // 假设这是从后端获取的会员到期时间
				isVip: false, // 初始时假设用户不是会员
			};
		},
		onLoad() {
			// 在组件加载时检查会员状态
			this.checkVipStatus();
		},
		methods: {
			handleClick(option) {
				this.activeOption = option;
			},
			selectPayment(option) {
				this.selectedPayment = option;
			},
			// checkVipStatus() {
			// 	uni.request({
			// 		url: 'http://120.46.81.37:8080/app/vip', // 替换为实际的URL和用户标识符
			// 		method: 'GET',
			// 		success: (res) => {
			// 			if (res.statusCode === 200) {
			// 				const data = res.data;
			// 				if (data && data.isVip) {
			// 					this.isVip = true; // 如果返回isVip为true，则将isVip设置为true
			// 				} else {
			// 					this.isVip = false; // 否则，设置为false
			// 				}
			// 			}
			// 		},
			// 		fail: (error) => {
			// 			console.error('检查会员状态失败', error);
			// 		}
			// 	});
			// },
			checkVipStatus() {
				uni.request({
					url: 'http://120.46.81.37:8080/app/isvip',
					method: 'GET',
					data: {
						'aToken': this.aToken, // 替换为实际的aToken
						'rToken': this.rToken, // 替换为实际的rToken
					},
					header: {
						'Content-Type': 'application/x-www-form-urlencoded', // 根据后端要求设置请求头
					},
					success: (res) => {
						if (res.statusCode === 200) {
							const data = res.data;
							if (data && data.isVip) {
								this.isVip = true; // 如果返回isVip为true，则将isVip设置为true
							} else {
								this.isVip = false; // 否则，设置为false
							}
						} else {
							console.error('检查会员状态失败，状态码：', res.statusCode);
						}
					},
					fail: (error) => {
						console.error('检查会员状态失败', error);
					}
				});
			},
			openPayment() {
				uni.navigateTo({
					url: '../payment/payment',
				});
			},
			// 纯前端的取消会员
			// cancelVip() {
			//   uni.showModal({
			//     title: 'Cancel Membership',
			//     message: 'Are you sure you want to cancel your membership?',
			//     success: (res) => {
			//       if (res.confirm) {
			//         // 模拟显示加载提示
			//         uni.showLoading({
			//           title: 'Processing',
			//         });
			//         // 模拟取消会员成功的场景
			//         setTimeout(() => {
			//           // 隐藏加载提示
			//           uni.hideLoading();
			//           // 更新UI状态，模拟取消会员成功
			//           this.isVip = false;
			//           this.expireDate = '';
			//           // 显示成功提示
			//           uni.showToast({
			//             title: 'Membership cancelled successfully',
			//           });
			//           // 如果需要，也可以模拟显示退款成功的提示
			//           uni.showToast({
			//             title: 'Refund processed successfully',
			//             icon: 'success',
			//           });
			//         }, 1500); // 延迟1.5秒模拟网络请求
			//       }
			//     },
			//     fail: (error) => {
			//       console.error('Modal Error', error);
			//     },
			//   });
			// },
			// 添加后端请求，取消vip会员
			cancelVip() {
				uni.showModal({
					title: 'Cancel Membership',
					message: 'Are you sure you want to cancel your membership?',
					success: (res) => {
						if (res.confirm) {
							uni.showLoading({
								title: 'Processing',
							});
							uni.request({
								url: 'http://120.46.81.37:8080/cancelVip', // API URL
								method: 'POST',
								data: {
									'aToken': aToken,
									'rToken': rToken,
								},
								header: {
									'Content-Type': 'application/x-www-form-urlencoded',
								},
								success: (res) => {
									uni.hideLoading();
									if (res.statusCode === 200) {
										uni.showToast({
											title: 'Membership cancelled successfully',
										});
										// 更新UI状态
										this.isVip = false;
										this.expireDate = '';
									} else {
										uni.showToast({
											title: 'Failed to cancel membership',
											icon: 'none',
										});
									}
								},
								fail: (error) => {
									console.error('Failed to cancel membership', error);
									uni.hideLoading();
									uni.showToast({
										title: 'Failed to cancel membership',
										icon: 'none',
									});
								},
							});
						}
					},
					fail: (error) => {
						console.error('Modal Error', error);
					},
				});
			},
			getExpireDate() {
				uni.request({
					url: 'http://120.46.81.37:8080/app/getVipExpireDate',
					method: 'GET',
					data: {
						'aToken': aToken,
						'rToken': rToken,
					},
					header: {
						'Content-Type': 'application/x-www-form-urlencoded',
					},
					success: (res) => {
						if (res.statusCode === 200) {
							// 假设返回的到期时间在res.data.expireDate
							this.expireDate = res.data.expireDate;
						} else {
							console.error('获取会员到期时间失败，状态码：', res.statusCode);
						}
					},
					fail: (error) => {
						console.error('获取会员到期时间失败', error);
					}
				});
			}
		},
		// 在组件创建时获取会员到期时间
		created() {
			this.getExpireDate();
		}
	};
</script>
<style>
	.expire-info {
		display: flex;
		align-items: center;
		/* 垂直居中 */
		justify-content: flex-start;
		/* 水平开始 */
		margin-bottom: 20px;
	}

	.expire-label {
		font-weight: bold;
		font-size: 20px;
		margin-right: 10px;
		/* 标签和时间之间的间距 */
	}

	.cancel-vip {
		margin-top: 5px;
		margin-bottom: 45px;
		text-align: center;
	}

	.cancel-button {
		/* 您的样式 */
		background-color: #007bff;
		color: #fff;
		/* ... */
	}

	.expire-value {
		flex-grow: 1;
		/* 让时间部分占据剩余空间 */
		font-size: 20px;
		text-align: center;
		/* 时间居中 */
		border: 1px solid #007bff;
		/* 边框样式示例 */
		padding: 5px 10px;
		/* 内边距 */
		border-radius: 5px;
		/* 圆角 */
	}
</style>