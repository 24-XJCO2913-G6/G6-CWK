<template>
	<view>
		<uni-list-item title="Avatar" @click="changeUserpic">
			<view class="flex align-center" slot="right">
				<image :src="user.userpic ? user.userpic : 'http://120.46.81.37:8080/app/static/default.jpg'"
				style="width: 100rpx;height: 100rpx;"
				class="rounded-circle"></image>
				<text class="iconfont icon-bianji1 ml-2"></text>
			</view>
		</uni-list-item>
		<uni-list-item title="Nickname">
			<view class="flex align-center" slot="right">
				<input class="uni-input text-right" v-model="username" />
				<text class="iconfont icon-bianji1 ml-2"></text>
			</view>
		</uni-list-item>
		<uni-list-item title="Sex" >
			<view class="flex align-center" slot="right">
				<input class="uni-input text-right" v-model="sex" />
				<text class="iconfont icon-bianji1 ml-2"></text>
			</view>
		</uni-list-item>
		<picker mode="date" :value="birthday" @change="onDateChange">
			<uni-list-item title="Birthday">
				<view class="flex align-center" slot="right">
					<text>{{birthday}}</text>
					<text class="iconfont icon-bianji1 ml-2"></text>
				</view>
			</uni-list-item>
		</picker>
		<uni-list-item title="Relationship Status" >
			<view class="flex align-center" slot="right">
				<input class="uni-input text-right" v-model="relationship" />
				<text class="iconfont icon-bianji1 ml-2"></text>
			</view>
		</uni-list-item>
		<uni-list-item title="Job" >
			<view class="flex align-center" slot="right">
				<input class="uni-input text-right" v-model="job" />
				<text class="iconfont icon-bianji1 ml-2"></text>
			</view>
		</uni-list-item>
		<uni-list-item title="Hometown" >
			<view class="flex align-center" slot="right">
				<input class="uni-input text-right" v-model="hometown" />
				<text class="iconfont icon-bianji1 ml-2"></text>
			</view>
		</uni-list-item>
		
		<view class="py-2 px-3">
			<button class="bg-main text-white" style="border-radius: 50rpx;border: 0;" type="primary" @click="submit">Complete</button>
		</view>
		
		
		<mpvue-city-picker :themeColor="themeColor" ref="mpvueCityPicker" :pickerValueDefault="cityPickerValueDefault" @onConfirm="onConfirm"></mpvue-city-picker>
		
	</view>
</template>

<script>
	import uniListItem from '@/components/uni-ui/uni-list-item/uni-list-item.vue';
	
	import mpvueCityPicker from '@/components/uni-ui/mpvue-citypicker/mpvueCityPicker.vue';
	
	import { mapState } from 'vuex'
	
	export default {
		components: {
			uniListItem,
			mpvueCityPicker
		},
		data() {
			return {
				themeColor: '#007AFF',
				cityPickerValueDefault: [0, 0, 1],
				avatar:'',
				username:"昵称",
				sex:'',
				job:"",
				birthday:"2019-03-18",
				hometown:'',
				relationship:'',
				emotion:''
			
				
			}
		},
		computed: {
			...mapState({
				aToken:state=>state.aToken,
				rToken:state=>state.rToken,
				
			}),
		
		},
		// 监听返回
		onBackPress() {
		  if (this.$refs.mpvueCityPicker.showPicker) {
		  	this.$refs.mpvueCityPicker.pickerCancel();
		    return true;
		  }
		},
		// 监听页面卸载
		onUnload() {
			if (this.$refs.mpvueCityPicker.showPicker) {
				this.$refs.mpvueCityPicker.pickerCancel()
			}
		},
		onLoad() {
			console.log('aa');
			console.log('token', this.aToken, this.rToken)
			uni.request({
				  url: 'http://localhost:8080/app/profile/149',
				  method: 'GET',
				  data:{
					  'aToken': this.aToken,
					  'rToken':this.rToken,
				  },
				  header: {
				      'Content-Type': 'application/x-www-form-urlencoded',
				  },
				  success: (res) => {
					  console.log('获取的个人信息',res)
						this.user_pic=res.user_pic
						this.username = res.username
						this.sex =  res.sex
						this.job  = res.job
						this.birthday  =res.birthday
						this.hometown=res.hometown
						this.relation=res.relation
				  },
				  fail: (err) => {
					  console.error('Error fetching person info:', err);
					
				  }
			})
		},

		methods: {
			showCityPicker(){
				this.$refs.mpvueCityPicker.show()
			},
			// 三级联动提交事件
			onConfirm(e) {
				this.pickerText = e.label
			},
			// 修改生日
			onDateChange(e){
				this.birthday = e.detail.value
			},
			// 修改头像
			changeUserpic(){
				uni.chooseImage({
					count:1,
					sizeType:["compressed"],
					sourceType:["album","camera"],
					success: (res) => {
						this.$H.upload('/edituserpic',{
							filePath: res.tempFilePaths[0],
							name: 'userpic',
							token:true
						}).then(result=>{
							console.log(result);
							this.$store.commit('editUserInfo',{
								key:"userpic",
								value:result.data
							})
							uni.showToast({
								title: '修改头像成功',
								icon: 'none'
							});
						}).catch(err=>{
							console.log(err);
						})
					}
				})
			},
			// 修改性别
			changeSex(){
				uni.showActionSheet({
				    itemList: sexArray,
				    success:(res)=>{
				        this.sex = res.tapIndex
				    }
				});
			},
			// 修改情感
			changeEmotion(){
				uni.showActionSheet({
				    itemList: emotionArray,
				    success:(res)=>{
				        this.emotion = res.tapIndex
				    }
				});
			},
			// 提交
			submit(){
				let obj = {
					user_pic:this.user_pic,
					username:this.username,
					sex:this.sex,
					relationship:this.relationship,
					job:this.job,
					birthday:this.birthday,
					hometown:this.hometown
				}
				console.log(obj);
				uni.request({
				url: 'http://120.46.81.37:8080/app/edit_profile',
				method: 'POST',
				data: obj,
				header: {
				    'Content-Type': 'application/x-www-form-urlencoded',
					'aToken': this.aToken,
					'rToken':this.rToken,
				},
				      success: (res) => {
					   this.user_pic=res.user_pic
				       this.username = res.username
				       this.sex =  res.sex
				       this.job  = res.job
				       this.birthday  =res.birthday
				       this.hometown=res.hometown
				       this.relationship=res.relationship
				      },
				      fail: (err) => {
				        console.error('Failed to fetch posts:', err);
				      }
				    });
			
				
				
			}
		}
	}
</script>

<style>

</style>
