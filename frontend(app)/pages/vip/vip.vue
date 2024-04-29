<template>
  <view class="container " style='margin-top:30px'>
    <view class="member-row">
      <view class="member-item" @click="handleClick('Month')" :class="{active: activeOption === 'Month'}">
        <view style="font-weight: bold; font-size: 25px;">Month</view>
        <view style=" font-size: 20px;"> $29.9 </view>
      </view>
      <view class="member-item" @click="handleClick('Quarter')" :class="{active: activeOption === 'Quarter'}">
		  <view style="font-weight: bold; font-size: 25px;">Quarter</view>
		  <view  style=" font-size: 20px;"> $69.9 </view>
      </view>
    </view>
    <view class="member-row">
      <view class="member-item" @click="handleClick('Year')" :class="{active: activeOption === 'Year'}">
		  <view style="font-weight: bold; font-size: 25px;">Year</view>
		  <view  style=" font-size: 20px;"> $199.9</view>
      </view>
      <view class="member-item" @click="handleClick('Permanent')" :class="{active: activeOption === 'Permanent'}">
		  <view style="font-weight: bold; font-size: 25px;"> Permanent</view>
		  <view  style=" font-size: 20px;"> $999</view>
      
      </view>
    </view>
	<view class="payment-block" style='margin-bottom: 40px;margin-top: 30px;'>
	      <view style="font-weight: bold; margin-bottom: 20px; font-size:20px">Payment</view>
	      <view class="payment-option" @click="selectPayment('Wechat')" style='padding-left:10px ;margin-bottom:15px'>
	         <img src="../../static/wechat.png"  alt="Wechat Icon" style="width: 24px; height: 24px; margin-right: 10px;"> Wechat <span v-show="selectedPayment === 'Wechat'"  style="font-size: 35px;float: right; color: blue; font-weight:bold; margin-right: 20px;">&#10003;</span>
	      </view>
	      <view class="payment-option" @click="selectPayment('Alipay')"  style='padding-left:10px'>
	         <img src="../../static/ali.png" alt="Wechat Icon" style="width: 24px; height: 24px; margin-right: 10px;"> Alipay <span v-show="selectedPayment === 'Alipay'" style="font-size: 35px;float: right; color:blue; font-weight:bold; margin-right: 20px;">&#10003;</span>
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
  height:100px;
  padding: 10px;
  text-align: center;
  background-color: #fff;
  color: #000;
  border: 1px solid #ccc;
  border-radius:10px;
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
  border:2px solid #FFAA33 ;
  background-color: #ffffe0;
  color: #FFAA33 ;
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
.payment-option{
	height:50px;
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
    };
  },
  methods: {
    handleClick(option) {
      this.activeOption = option;
    },
	selectPayment(option) {
	  this.selectedPayment = option;
	},
	openPayment(){
		uni.navigateTo({
			url: '../payment/payment',
		});
	}
  }
};
</script>
