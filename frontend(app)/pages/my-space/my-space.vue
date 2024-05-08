<template>
	<view>
		<!-- #ifdef MP -->
		<uni-nav-bar :shadow="false" :fixed="true" :border="false" right-text="Menu" @click-right="clickNavigationButton"></uni-nav-bar>
		<!-- #endif -->
		<!-- 头部 -->
		<view class="flex align-center p-3 border-bottom border-light-secondary">
			<image src="../../static/user_pic/4.jpg" 
			style="width: 180rpx;height: 180rpx;"
			class="rounded-circle"></image>
			<view class="pl-3 flex flex-column flex-1">
				<view class="flex align-center">
					<view class="flex-1 flex flex-column align-center justify-center" v-for="(item,index) in counts" :key="index">
						<text class="font-lg font-weight-bold">{{item.num|formatNum}}</text>
						<text class="font text-muted">{{item.name}}</text>
					</view>
				</view>
				<view class="flex justify-center align-center">
					
					<button v-if="user_id == user.id"
					 type="default" size="mini"
					style="width: 400rpx;" @click="openUserInfo">
						Edit Profile
					</button>
			

				</view>
			</view>
		</view>
		
		

		<template >
			<view class="animated fast fadeIn">
				<common-list v-for="(item,index) in posts" :key="index" :item="item" :index="index" @follow="follow" @doSupport="doSupport"></common-list>
				<divider></divider>
				<load-more :loadmore="loadmore"></load-more>
			</view>
		</template>
		
		
		
		<!-- 弹出层 -->
		<uni-popup ref="popup" type="top">
			<view class="flex align-center justify-center font-md border-bottom py-2" hover-class="bg-light" @click="doBlack">
				<text class="iconfont icon-sousuo mr-2"></text> 
				{{userinfo.isblack ? 'de-blacklist' : 'blacklist'}}
			</view>
			<view v-if="!userinfo.isblack" class="flex align-center justify-center font-md py-2" hover-class="bg-light" @click="openChat">
				<text class="iconfont icon-shanchu mr-2"></text> Talk
			</view>
		</uni-popup>
		
		
	</view>
</template>

<script>
	const emotionArray = ['secrecy', 'unmarried', 'married']
	import commonList from '@/components/common/common-list.vue';
	import loadMore from '@/components/common/load-more.vue';
	import uniPopup from '@/components/uni-ui/uni-popup/uni-popup.vue';
	import $T from '@/common/time.js';
	import { mapState } from 'vuex'
	import uniNavBar from '@/components/uni-ui/uni-nav-bar/uni-nav-bar.vue';
	export default {
		components: {
			commonList,
			loadMore,
			uniPopup,
			uniNavBar
		},
		data() {
			return {
				posts:[
					{
					post_id:1,
					name:'Super girl',
					visibility:1,
					time:'2024-5-5',
					like:0,
					userpic:'../../static/user_pic/4.jpg',
					collect:0,
					content:'Cycling by the lake',
					title:'Cool day',
					comments_num:0,
					path:[
					  [
					    103.976838,
					    30.768159
					  ],
					  [
					    103.976871,
					    30.768094
					  ],
					  [
					    103.976905,
					    30.768044
					  ],
					  [
					    103.976921,
					    30.768023
					  ],
					  [
					    103.976938,
					    30.768001
					  ],
					  [
					    103.97703,
					    30.768073
					  ],
					  [
					    103.977097,
					    30.768109
					  ],
					  [
					    103.977155,
					    30.768166
					  ],
					  [
					    103.977222,
					    30.768195
					  ],
					  [
					    103.977339,
					    30.76831
					  ],
					  [
					    103.977439,
					    30.768382
					  ],
					  [
					    103.977548,
					    30.768439
					  ],
					  [
					    103.977673,
					    30.768511
					  ],
					  [
					    103.977799,
					    30.768597
					  ],
					  [
					    103.977983,
					    30.768726
					  ],
					  [
					    103.978108,
					    30.768827
					  ],
					  [
					    103.97825,
					    30.76892
					  ],
					  [
					    103.9784,
					    30.769042
					  ],
					  [
					    103.978526,
					    30.769143
					  ],
					  [
					    103.978651,
					    30.76925
					  ],
					  [
					    103.978801,
					    30.769372
					  ],
					  [
					    103.978969,
					    30.769523
					  ],
					  [
					    103.979127,
					    30.769652
					  ],
					  [
					    103.979303,
					    30.769439
					  ],
					  [
					    103.979345,
					    30.769381
					  ],
					  [
					    103.979445,
					    30.769252
					  ],
					  [
					    103.979545,
					    30.769108
					  ],
					  [
					    103.97962,
					    30.769008
					  ],
					  [
					    103.979796,
					    30.768843
					  ],
					  [
					    103.97998,
					    30.768606
					  ],
					  [
					    103.980113,
					    30.768448
					  ],
					  [
					    103.980205,
					    30.768276
					  ],
					  [
					    103.980222,
					    30.768204
					  ],
					  [
					    103.980272,
					    30.768211
					  ],
					  [
					    103.980356,
					    30.768247
					  ],
					  [
					    103.980447,
					    30.768326
					  ],
					  [
					    103.980573,
					    30.76839
					  ],
					  [
					    103.980673,
					    30.768441
					  ],
					  [
					    103.980715,
					    30.768484
					  ],
					  [
					    103.980665,
					    30.768577
					  ],
					  [
					    103.98059,
					    30.768721
					  ],
					  [
					    103.980498,
					    30.768879
					  ],
					  [
					    103.980297,
					    30.76908
					  ],
					  [
					    103.980105,
					    30.769302
					  ],
					  [
					    103.98003,
					    30.769431
					  ],
					  [
					    103.979971,
					    30.769525
					  ],
					  [
					    103.979946,
					    30.769661
					  ],
					  [
					    103.979896,
					    30.769733
					  ],
					  [
					    103.979904,
					    30.76979
					  ],
					  [
					    103.980038,
					    30.769776
					  ],
					  [
					    103.980172,
					    30.76969
					  ],
					  [
					    103.980272,
					    30.769546
					  ],
					  [
					    103.980272,
					    30.76941
					  ],
					  [
					    103.980305,
					    30.769302
					  ],
					  [
					    103.980322,
					    30.769209
					  ],
					  [
					    103.980322,
					    30.769159
					  ],
					  [
					    103.980389,
					    30.769116
					  ],
					  [
					    103.980506,
					    30.769015
					  ],
					  [
					    103.980581,
					    30.768979
					  ],
					  [
					    103.980706,
					    30.7689
					  ],
					  [
					    103.980757,
					    30.768871
					  ],
					  [
					    103.980832,
					    30.768814
					  ],
					  [
					    103.98094,
					    30.768771
					  ],
					  [
					    103.981007,
					    30.768713
					  ],
					  [
					    103.981074,
					    30.76867
					  ],
					  [
					    103.981183,
					    30.768584
					  ],
					  [
					    103.981258,
					    30.76852
					  ],
					  [
					    103.981283,
					    30.768469
					  ],
					  [
					    103.981308,
					    30.768426
					  ]
					],
											center: [
					    103.980322,
					    30.769209
					  ]
				},
				{	
					post_id:4,
					name:'Super girl',
					visibility:1,
					time:'2024-4-13',
					like:6,
					userpic:'../../static/user_pic/4.jpg',
					collect:3,
					title:'School walk',
					content:'Nice place to go',
					comments_num:5,
					path:[
					  [
					    103.977527,
					    30.767054
					  ],
					  [
					    103.977643,
					    30.766825
					  ],
					  [
					    103.977727,
					    30.766768
					  ],
					  [
					    103.977826,
					    30.766625
					  ],
					  [
					    103.977926,
					    30.766525
					  ],
					  [
					    103.978076,
					    30.766325
					  ],
					  [
					    103.978375,
					    30.765926
					  ],
					  [
					    103.978558,
					    30.76574
					  ],
					  [
					    103.978741,
					    30.765526
					  ],
					  [
					    103.978874,
					    30.765297
					  ],
					  [
					    103.978973,
					    30.765269
					  ],
					  [
					    103.97904,
					    30.765255
					  ],
					  [
					    103.979106,
					    30.765312
					  ],
					  [
					    103.979273,
					    30.765355
					  ],
					  [
					    103.979406,
					    30.765483
					  ],
					  [
					    103.979672,
					    30.765583
					  ],
					  [
					    103.979788,
					    30.765669
					  ],
					  [
					    103.979938,
					    30.765769
					  ],
					  [
					    103.980154,
					    30.765869
					  ],
					  [
					    103.980337,
					    30.766012
					  ],
					  [
					    103.980486,
					    30.766097
					  ],
					  [
					    103.980619,
					    30.766169
					  ],
					  [
					    103.980669,
					    30.766169
					  ],
					  [
					    103.980736,
					    30.766026
					  ],
					  [
					    103.980785,
					    30.76584
					  ],
					  [
					    103.980835,
					    30.765697
					  ],
					  [
					    103.981051,
					    30.765612
					  ],
					  [
					    103.981168,
					    30.765569
					  ],
					  [
					    103.981317,
					    30.765555
					  ],
					  [
					    103.981517,
					    30.765555
					  ],
					  [
					    103.98175,
					    30.765497
					  ],
					  [
					    103.981999,
					    30.76534
					  ],
					  [
					    103.982165,
					    30.765155
					  ],
					  [
					    103.982282,
					    30.764969
					  ],
					  [
					    103.982365,
					    30.764812
					  ]
					],
											 center:   [
					    103.981317,
					    30.765555
					  ],
				
				},	],
				user_id:0,
				userinfo:{
					userpic:"/static/default.jpg",
					username:"",
					sex:0,
					age:20,
					isFollow:false,
					regtime:"",
					birthday:"",
					job:"",
					path:"",
					qg:""
				},
				counts:[{
					name:"Post",
					num:2
				},{
					name:"Follow",
					num:1
				},{
					name:"Fan",
					num:0
				}],
				tabIndex:0,
				tabBars:[{
					name:"Post",
					list:[],
					// 1.上拉加载更多  2.加载中... 3.没有更多了
					loadmore:"Load more",
					page:1
				}],
			}
		},
		onNavigationBarButtonTap() {
			this.clickNavigationButton()
		},
		computed: {
			...mapState({
				user:state=>state.user
			}),
		
		},
		filters: {
			formatNum(value) {
				return value > 99 ? '99+' : value;
			}
		},
		onLoad(e) {
			
			
		
		},
		onUnload() {
			uni.$off('updateFollowOrSupport',(e)=>{})
			uni.$off('updateCommentsCount',(e)=>{})
		},
		methods: {
			clickNavigationButton(){
				if(this.user_id == this.user.id){
					return uni.navigateTo({
						url: '../user-set/user-set',
					});
				}
				this.$refs.popup.open()
			},
			// 获取用户相关数据
			getCounts(){
				this.$H.get('/user/getcounts/'+this.user_id).then(res=>{
					this.counts[0].num = res.post_count
					this.counts[1].num = res.withfollow_count
					this.counts[2].num = res.withfen_count
				})
			},
			// 获取用户个人信息
			getUserInfo(){
				this.$H.post('/getuserinfo',{
					user_id:this.user_id
				},{
					token:true,
					notoast:true
				}).then(res=>{
					this.userinfo = {
						userpic:res.userpic,
						username:res.username,
						sex:res.userinfo.sex,
						age:res.userinfo.age,
						isFollow:res.fens.length > 0,
						isblack:res.blacklist.length > 0,
						regtime:$T.dateFormat(new Date(res.create_time * 1000), '{Y}-{MM}-{DD}'),
						birthday:$T.getHoroscope(res.userinfo.birthday),
						job:res.userinfo.job ? res.userinfo.job : 'No',
						path:res.userinfo.path ? res.userinfo.path : 'No',
						qg:emotionArray[res.userinfo.qg],
					}
					uni.setNavigationBarTitle({
						title:this.userinfo.username
					})
				})
			},
	
			// 关注/取消关注
			doFollow(){
				this.checkAuth(()=>{
					let url = this.userinfo.isFollow ? '/unfollow' : '/follow'
					let msg = this.userinfo.isFollow ? 'Unfollow' : 'Follow'
					this.$H.post(url,{
						follow_id:this.user_id
					},{
						token:true
					}).then(res=>{
						this.userinfo.isFollow = !this.userinfo.isFollow
						uni.showToast({
							title: msg+'success',
							icon: 'none'
						});
						uni.$emit('updateIndex')
						this.getList()
					})
				})
			},
			// 进入编辑资料
			openUserInfo(){
				uni.navigateTo({
					url: '../user-userinfo/user-userinfo',
				});
			},
			// 加入/移出黑名单
			doBlack(){
				this.checkAuth(()=>{
					let url = this.userinfo.isblack ? '/removeblack':'/addblack'
					let msg = this.userinfo.isblack ? 'de-blacklist' : 'blacklist'
					uni.showModal({
						content: '是否要'+msg,
						success: (res)=> {
							if (res.confirm) {
								this.$H.post(url,{
									id:this.user_id
								},{
									token:true
								}).then(res=>{
									this.userinfo.isblack = !this.userinfo.isblack
									uni.showToast({
										title: msg+'Success',
										icon: 'none'
									});
								})
							}
						}
					});
					
				})
			},
			openChat(){
				let user = {
					user_id:this.user_id,
					username:this.userinfo.username,
					userpic:this.userinfo.userpic
				}
				this.navigateTo({
					url:"/pages/user-chat/user-chat?user="+JSON.stringify(user)
				})
			}
		}
	}
</script>

<style>

</style>
