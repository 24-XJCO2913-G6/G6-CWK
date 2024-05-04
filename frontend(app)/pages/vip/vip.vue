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
				<img src="../../static/wechat.png" alt="Wechat Icon"
					style="width: 24px; height: 24px; margin-right: 10px;"> Wechat <span
					v-show="selectedPayment === 'Wechat'"
					style="font-size: 35px;float: right; color: blue; font-weight:bold; margin-right: 20px;">&#10003;</span>
			</view>
			<view class="payment-option" @click="selectPayment('Alipay')" style='padding-left:10px'>
				<img src="../../static/ali.png" alt="Wechat Icon"
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
				isVip: true, // 初始时假设用户不是会员
			};
		},
		methods: {
			handleClick(option) {
				this.activeOption = option;
			},
			selectPayment(option) {
				this.selectedPayment = option;
			},
			checkVipStatus() {
				// 这里使用uni.request来发送请求，您需要根据实际情况替换URL和参数
				uni.request({
					url: '你的后端服务地址/vip/用户的唯一标识符', // 替换为实际的URL和用户标识符
					method: 'GET',
					success: (res) => {
						if (res.statusCode === 200) {
							const data = res.data;
							if (data && data.isVip) {
								this.isVip = true; // 如果返回isVip为true，则将isVip设置为true
							} else {
								this.isVip = false; // 否则，设置为false
							}
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
			cancelVip() {
			      // 向用户确认是否取消会员
			      const confirm = uni.getSystemInfoSync().platform === 'android' ? 'confirm' : 'alert';
			      uni[confirm]({
			        title: '取消会员',
			        message: '您确定要取消会员吗？',
			        success: (res) => {
			          if (res.confirm) {
			            // 发送取消会员请求
			            uni.request({
			              url: '你的后端服务地址/vip/取消', // 替换为实际的取消会员的后端API地址
			              method: 'POST',
			              // 如果需要认证信息，可以在header中添加
			              // header: {
			              //   'Authorization': 'Bearer 用户的令牌'
			              // },
			              success: (res) => {
			                if (res.statusCode === 200) {
			                  uni.showToast({
			                    title: '取消会员成功',
			                  });
			                  // 更新会员状态
			                  this.isVip = false;
			                  // 清空会员到期时间
			                  this.expireDate = '';
			                } else {
			                  uni.showToast({
			                    title: '取消会员失败',
			                    icon: 'none',
			                  });
			                }
			              },
			              fail: (error) => {
			                console.error('取消会员失败', error);
			                uni.showToast({
			                  title: '取消会员失败',
			                  icon: 'none',
			                });
			              }
			            });
			          }
			        },
			      });
			    },
			// 假设这是获取会员到期时间的方法
			getExpireDate() {
				// 这里应该是一个API请求，获取用户的会员到期时间
				// 请替换为实际的后端API地址
				uni.request({
					url: '你的后端服务地址/api/getVipExpireDate',
					method: 'GET',
					// 其他配置...
					success: (res) => {
						if (res.statusCode === 200) {
							
							// this.expireDate = res.data.expireDate; // 假设返回的到期时间在res.data.expireDate
						} else {
							console.error('获取会员到期时间失败，状态码：', res.statusCode);
						}
					},
					fail: (error) => {
						console.error('获取会员到期时间失败', error);
					}
				});
				// 模拟返回一个到期时间字符串
				this.expireDate = '2024-06-30';
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
	  background-color:  #007bff;
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