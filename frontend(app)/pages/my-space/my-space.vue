<template>
	<view>
		<!-- #ifdef MP -->
		<uni-nav-bar :shadow="false" :fixed="true" :border="false" right-text="Menu" @click-right="clickNavigationButton"></uni-nav-bar>
		<!-- #endif -->
		<!-- 头部 -->
		<view class="flex align-center p-3 border-bottom border-light-secondary">
			<image :src="userinfo.userpic" 
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
					name:'Alice',
					visibility:1,
					time:'2022-01-12',
					like:19,
					userpic:'',
					collect:13,
					content:'WOw',
					comments:5,
					center:[103.985895, 30.763873]	,
					path:[
									    [103.985895, 30.763873], // 起点坐标
									    [103.986895, 30.764873], // 中间很多点坐标
									    [103.987895, 30.765873]],
				},
				{	
					post_id:2,
					name:'Jack',
					visibility:0,
					time:'2022-01-12',
					like:20,
					userpic:'',
					collect:3,
					content:'Nice place to go',
					comments:5,
					center:[103.985805, 30.763883]	,
					path:[
									    [103.985895, 31.763873], // 起点坐标
									    [103.986895, 31.764873], // 中间很多点坐标
									    [103.987895, 30.765873]],
				
										
					
				
				},	{
					post_id:3,
					name:'Mike',
					visibility:1,
					time:'2022-01-12',
					like:5,
					userpic:'',
					collect:8,
					content:'I like it',
					comments:5,
					center:[103.985895, 30.763873]	,
					path:[
									    [103.985895, 30.763873], // 起点坐标
									    [103.986895, 30.764873], // 中间很多点坐标
									    [103.987895, 30.765873]],
				
										
					
				
				}],
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
					num:0
				},{
					name:"Follow",
					num:0
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
			if(!e.user_id){
				return uni.showToast({
					title: 'Invalid parameter',
					icon: 'none'
				});
			}
			this.user_id = e.user_id
			// 加载用户个人信息
			this.getUserInfo()
			// 获取用户相关数据
			this.getCounts()
			// 监听关注和顶踩操作
			uni.$on('updateFollowOrSupport',(e)=>{
				switch (e.type){
					case 'follow': // 关注
					this.follow(e.data.user_id)
						break;
					case 'support': // 顶踩
					this.doSupport(e.data)
						break;
				}
			})
		
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
