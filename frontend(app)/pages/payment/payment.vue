<template>  
    <view class="password-set-page">  
        <cu-custom is-back></cu-custom>  
        <view class="text-xxl text-center text-bold margin-top-xl"></view>  
        <view class="text-df text-center text-gray margin-tb">Enter your payment password</view>  
        <view class="code flex align-center justify-center" style="margin-top: 10vh;">  
            <view class="flex align-center justify-center">  
                <view v-for="(item, index) in 6" :key="index">{{ password[index] && '●' || '' }}</view>  
            </view>  
        </view>  
        <view class="keyboard flex flex-wrap">  
            <button v-for="(item, index) in 9" :key="index" @click="key(index + 1)">  
                <text>{{ index + 1 }}</text>  
            </button>  
            <button class="hide"></button>  
            <button @click="key(0)">  
                <text>0</text>  
            </button>  
            <button @click="del()">  
                ×  
            </button>  
        </view>  
    </view>  
</template>  
  
<script>  
    export default {  
        data() {  
            return {  
                password: ''  
            }  
        },  
        methods: {  
            showToast2() {  
                uni.showToast({  
                    title: 'Successfully Pay',  
                    icon: 'success', // 将值设置为 success 或者 ''  
                    duration: 2000 // 持续时间为 2秒  
                })  
            },  

            key(key) {  
                    if (this.password.length < 6) {  
                        this.password += key;  
                    } else if (this.password.length == 6) {  
                        // 密码验证操作（这里假设密码验证总是成功，实际应用中需要添加验证逻辑）  
                        this.showToast2(); // 当密码长度为6时，显示支付成功的弹窗  
              
                        // 清空密码字段，防止重复触发  
                        this.password = ''; // 清空密码字段  
                        // 调用页面跳转函数  
						setTimeout(()=>{
							this.navigateToMyPage();  
						},1000)
                        
                    }  
                },  
			navigateToMyPage() {  
			        uni.navigateTo({  
			            url: '../vip/vip' // 替换为你想要跳转的页面的路径  
			        }).catch(error => {  
			            console.error('页面跳转失败:', error);  
			            // 这里可以添加错误处理逻辑，比如显示一个错误提示等  
			        });  
			    },  	
            del() {  
                if (this.password.length > 0) {  
                    this.password = this.password.substring(0, this.password.length - 1);  
                }  
            }  
        }  
    }  
</script>




<style lang="scss">
    page {
        background: #FFFFFF;
    }
    .password-set-page {
        .code {
            >view {
                border: 1px solid #DDDDDD;
                border-radius: 8rpx;
                overflow: hidden;
                view {
                    display: flex;
                    align-items: center;
                    justify-content: center;
                    width: 100rpx;
                    height: 100rpx;
                    background: #FAFAFA;
                    font-size: 60rpx;
                    &:not(:last-child) {
                        border-right: 1px solid #DDDDDD;
                    }
                }
            }
        }
        .keyboard {
            position: fixed;
            bottom: 0;
            width: 100%;
            background: #EEEEEE;
            button {
                display: flex;
                align-items: center;
                justify-content: center;
                width: calc(100vw / 3 - 1px);
                background: #FFFFFF;
                border-radius: 0;
                margin-top: 1px;
                font-size: 50rpx;
                height: 120rpx;
                &.button-hover:not(.hide) {
                    background: #EEEEEE;
                }
                image {
                    width: 52rpx;
                    height: 38rpx;
                }
            }
        }
    }
	.text-df {
	  font-weight: bold;
	  font-size: 1.2em; // Adjust the size as needed
	  &.text-center {
	    text-align: center;
	  }
	  &.text-gray {
	    color: gray;
	  }
	  &.margin-tb {
	    margin-top: 10px; // Adjust top margin as needed
	    margin-bottom: 10px; // Adjust bottom margin as needed
	  }
	}

</style>