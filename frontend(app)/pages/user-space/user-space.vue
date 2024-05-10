<template>
	<view>
		<!-- #ifdef MP -->
		<uni-nav-bar :shadow="false" :fixed="true" :border="false" right-text="Menu"
			@click-right="clickNavigationButton"></uni-nav-bar>
		<!-- #endif -->
		<!-- 头部 -->
		<view class="flex align-center p-3 border-bottom border-light-secondary">
			<image :src="userinfo.userpic" style="width: 180rpx;height: 180rpx;" class="rounded-circle"></image>
			<view class="pl-3 flex flex-column flex-1">
				<view class="flex align-center">
					<view class="flex-1 flex flex-column align-center justify-center" v-for="(item,index) in counts"
						:key="index">
						<text class="font-lg font-weight-bold">{{item.num|formatNum}}</text>
						<text class="font text-muted">{{item.name}}</text>
					</view>
				</view>	
			</view>
		</view>

		<!-- tab -->
		<view class="flex align-center" style="height: 100rpx;">
			<view class="flex-1 flex align-center justify-center" v-for="(item,index) in tabBars" :key="index"
				:class="index === tabIndex ? 'font-lg font-weight-bold text-main':'font-md'" @click="changeTab(index)">
				{{item.name}}
			</view>
		</view>

		<template v-if="tabIndex === 0">
			<view class="animated fast fadeIn">
				<view class="p-3 border-bottom">
					<view class="font-md">Information</view>
					<view class="font">Sex: {{userinfo.sex}}</view>
					<view class="font">Birthday：{{userinfo.birthday}}</view>
					<view class="font">Job：{{userinfo.job}}</view>
					<view class="font">Hometown：{{userinfo.hometown}}</view>
					<view class="font">Relationship Status：{{userinfo.relation}}</view>
				</view>
			</view>
		</template>
		<template v-else>
			<view class="animated fast fadeIn">
				<!-- <common-list :item="myitem" v-for="myitem in posts" :key="myitem.post_id"></common-list> -->
				<common-list :item="post" v-for="post in posts" :key="post.post_id"></common-list>
				<divider></divider>
				<!-- <load-more :loadmore="loadmore"></load-more> -->
			</view>
		</template>



		<!-- 弹出层 -->
		<uni-popup ref="popup" type="top">
			<view class="flex align-center justify-center font-md border-bottom py-2" hover-class="bg-light"
				@click="doBlack">
				<text class="iconfont icon-sousuo mr-2"></text>
				{{userinfo.isblack ? 'de-blacklist' : 'blacklist'}}
			</view>
			<view v-if="!userinfo.isblack" class="flex align-center justify-center font-md py-2" hover-class="bg-light"
				@click="openChat">
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
	import {
		mapState
	} from 'vuex'
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
				posts: [{
					post_id: 1,
					name: 'Yodo',
					visibility: 1,
					time: '2024-05-01',
					like: 19,
					islike:1,
					iscollect:0,
					userpic: 'http://120.46.81.37:8080/app/static/user_pic/1.jpg',
					title: 'Nice day to have a walk',
					collect: 13,
					content: '44444444444444444',
					comments_num: 0,
					center: [
						103.981051,
						30.765612
					],
					path:[
					{points:[
					  {"longitude": 103.977527, "latitude": 30.767054},
					  {"longitude": 103.977643, "latitude": 30.766825},
					  {"longitude": 103.977727, "latitude": 30.766768},
					  {"longitude": 103.977826, "latitude": 30.766625},
					  {"longitude": 103.977926, "latitude": 30.766525},
					  {"longitude": 103.978076, "latitude": 30.766325},
					  {"longitude": 103.978375, "latitude": 30.765926},
					  {"longitude": 103.978558, "latitude": 30.76574},
					  {"longitude": 103.978741, "latitude": 30.765526},
					  {"longitude": 103.978874, "latitude": 30.765297},
					  {"longitude": 103.978973, "latitude": 30.765269},
					  {"longitude": 103.97904, "latitude": 30.765255},
					  {"longitude": 103.979106, "latitude": 30.765312},
					  {"longitude": 103.979273, "latitude": 30.765355},
					  {"longitude": 103.979406, "latitude": 30.765483},
					  {"longitude": 103.979672, "latitude": 30.765583},
					  {"longitude": 103.979788, "latitude": 30.765669},
					  {"longitude": 103.979938, "latitude": 30.765769},
					  {"longitude": 103.980154, "latitude": 30.765869},
					  {"longitude": 103.980337, "latitude": 30.766012},
					  {"longitude": 103.980486, "latitude": 30.766097},
					  {"longitude": 103.980619, "latitude": 30.766169},
					  {"longitude": 103.980669, "latitude": 30.766169},
					  {"longitude": 103.980736, "latitude": 30.766026},
					  {"longitude": 103.980785, "latitude": 30.76584},
					  {"longitude": 103.980835, "latitude": 30.765697},
					  {"longitude": 103.981051, "latitude": 30.765612},
					  {"longitude": 103.981168, "latitude": 30.765569},
					  {"longitude": 103.981317, "latitude": 30.765555},
					  {"longitude": 103.981517, "latitude": 30.765555},
					  {"longitude": 103.98175, "latitude": 30.765497},
					  {"longitude": 103.981999, "latitude": 30.76534},
					  {"longitude": 103.982165, "latitude": 30.765155},
					  {"longitude": 103.982282, "latitude": 30.764969},
					  {"longitude": 103.982365, "latitude": 30.764812}
					],color: "#2e7efd",
										width: 10,
										dottedLine: true,
										arrowLine: true,
										colorList: true,}
					 
					],markers:[
											    {"longitude": 103.982365, "latitude": 30.764812, iconPath:'../../static/images/marker.png', width: 10,
					          height: 10} ] ,
					comments: [
						{
							userpic: 'http://120.46.81.37:8080/app/static/user_pic/2.jpg',
							username: 'Vero',
							data: 'Good!',
							time: '2024-4-7'
						},

					],
				}],
				userinfo: {
					userpic: "http://120.46.81.37:8080/app/static/user_pic/1.jpg",
					username: "Abby",
					sex: 'female',
					birthday: "2003-08-17",
					job: "Software Engineer",
					hometown: "Chengdu",
					relation: "married"
				},
				counts: [{
					name: "Post",
					num: 1
				}, {
					name: "Follow",
					num: 0
				}, {
					name: "Fan",
					num: 1
				}],
				tabIndex: 0,
				tabBars: [{
					name: "Info",
				}, {
					name: "Post",
					list: [

					],
					// 1.上拉加载更多  2.加载中... 3.没有更多了
					loadmore: "Load more",
					page: 1
				}],
			}
		},
		onNavigationBarButtonTap() {
			this.clickNavigationButton()
		},
		computed: {
			...mapState({
				user: state => state.user
			}),
			list() {
				return this.tabBars[this.tabIndex].list
			},
			loadmore() {
				return this.tabBars[this.tabIndex].loadmore
			}
		},
		filters: {
			formatNum(value) {
				return value > 99 ? '99+' : value;
			}
		},
		onLoad(e) {
		uni.request({
			  url: 'http://120.46.81.37:8080/app/profile_detail',
			  method: 'POST',
header: {
				    'Content-Type': 'application/x-www-form-urlencoded',
					'aToken': this.aToken,
					'rToken':this.rToken,
				},
			  success: (res) => {
					this.user_info.user_pic=res.user_pic
					this.user_info.username = res.username
					this.user_info.sex  =  res.sex
					this.user_info.job  = res.job
					this.user_info.birthday  =res.birthday
					this.user_info.hometown =res.hometown
					this.user_info.relation =res.relation
			  },
			  fail: (err) => {
				  console.error('Error fetching person info:', err);
				
			  }
		});
			uni.request({
					  //获取某个用户发过的所有帖子
						url: 'http://120.46.81.37:8080/app/get_posts',
			           method: 'GET',
header: {
				    'Content-Type': 'application/x-www-form-urlencoded',
					'aToken': this.aToken,
					'rToken':this.rToken,
				},
			           success: (res) => {
						   this.counts[0].num=res.data.length;
						   this.posts=res.data;
			               console.log('Post data uploaded successfully:', res);
			           },
			           fail: (err) => {
			               console.error('Error uploading post data:', err);
			           }
			       });
			uni.request({
						//获取某个用户所有关注的人
						url: 'http://120.46.81.37:8080/app/get_follow',
			           method: 'GET',
header: {
				    'Content-Type': 'application/x-www-form-urlencoded',
					'aToken': this.aToken,
					'rToken':this.rToken,
				},
			           success: (res) => {
						   this.counts[1].num=res.data.length;
			               console.log('Post data uploaded successfully:', res);
			           },
			           fail: (err) => {
			               console.error('Error uploading post data:', err);
			           }
			       });
			uni.request({
						//获取某个用户所有的粉丝
						url: 'http://120.46.81.37:8080/app/get_fan',
			           method: 'GET',
header: {
				    'Content-Type': 'application/x-www-form-urlencoded',
					'aToken': this.aToken,
					'rToken':this.rToken,
				},
			           success: (res) => {
						   this.counts[2].num=res.data.length;
			               console.log('Post data uploaded successfully:', res);
			           },
			           fail: (err) => {
			               console.error('Error uploading post data:', err);
			           }
			       });

		},
		onUnload() {
			
		},
		methods: {
			dofollow(){
				uni.request({
				    url: 'http://120.46.81.37:8080/app/follow',
				    method: 'POST',
				  data: {
				      'user_id': user_id,
				  },
header: {
				    'Content-Type': 'application/x-www-form-urlencoded',
					'aToken': this.aToken,
					'rToken':this.rToken,
				},
				    success: function (res) {
				        console.log('Post request successful:', res.data);
				        // 可以在这里处理返回的数据
				    },
				    fail: function (err) {
				        console.error('Post request failed:', err);
				    }
				});
			},
			clickNavigationButton() {
				if (this.user_id == this.user.id) {
					return uni.navigateTo({
						url: '../user-set/user-set',
					});
				}
				
			},
	
			changeTab(index) {
				this.tabIndex = index
				this.getList()
			},


		}
	}
</script>

<style>

</style>
