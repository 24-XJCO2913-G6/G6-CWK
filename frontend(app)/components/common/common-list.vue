<template>

	<view class="p-2 animated fast fadeIn">


		<view class="flex align-center justify-between">
			<view class="flex align-center">
		
				<image class="rounded-circle mr-2" :src="item.Photo" style="width: 65rpx;height: 65rpx;" lazy-load
					@click="navigateToPage"></image>
				<view>
					<view class="font" style="line-height: 1.5;">
						{{item.Author}}
					</view>
					<text class="font-sm text-light-muted" style="line-height: 1.5;">
						{{item.Pub_time}}
					</text>
				</view>
				<view v-if="item.visibility== 1">
					<image src="http://120.46.81.37:8080/static/images/VIP.png"
						style="width: 25px; height: 25px; background-color: white; margin-bottom: 2px;">
					</image>
				</view>

			</view>

			<!-- 按钮 -->
			<view v-if='!ifme||item.isFollow' class="flex align-center justify-center rounded bg-main text-white animated faster"
				:style="[defaultcss]" hover-class="jello" @click='myfollow(item.user_id)' id='myfollow' >
				Follow
			</view>
		</view>
		<view class="font-md my-1" @click='openDetail'>{{item.title}}</view>
		<map style='height:200px;width:100%;' scale='16' :latitude="item.center[1]" :longitude="item.center[0]"  :polyline="item.path" :markers="item.markers" >

		 </map>
		                      

		<slot></slot>
		<!-- 标题 -->

		<!-- 帖子详情 -->

		<!-- 图标按钮 -->
		<view class="flex align-center">

			<!-- 顶 -->
			<view class="flex align-center justify-center flex-1 animated faster" hover-class="jello text-main">
				<text class="iconfont icon-dianzan2 mr-2" @click='mylike(item.post_id)' id='theLike'></text>
				<text>{{item.like}}</text>
			</view>
			<!-- 踩 -->
			<view class="flex align-center justify-center flex-1 animated faster" hover-class="jello text-main">
				<text class="iconfont icon-xihuan1 mr-2" @click='mycollect(item.post_id)' id='theCollect'></text> <!-- TODO:收藏图标 -->
				<text>{{item.collect}}</text>
			</view>
			<view class="flex align-center justify-center flex-1 animated faster" hover-class="jello text-main">
				<text class="iconfont icon-pinglun2 mr-2"></text>

				<text>{{item.comments_num}}</text>
			</view>

		</view>
	</view>
</template>

<script>
	import $T from '@/common/time.js';
	import {
		mapState
	} from 'vuex';
	import mymap from '../../components/map/mymap.vue';

	export default {
		components: {
			mymap,"map": "@uni/map"
		},
		data(){
			return {defaultcss:'width: 90rpx;height: 50rpx;',
					latitude: 39.909,
						longitude: 116.39742,
						covers: [{
							latitude: 39.909,
							longitude: 116.39742,
						
						}, {
							latitude: 39.90,
							longitude: 116.39,
				
						}]}
		},
		props: {
			item: Object,
			index: {
				type: Number,
				default: -1
			},
			isdetail: {
				type: Boolean,
				default: false
			},
			ifme:Boolean
		},
		filters: {
			formatTime(value) {
				return $T.gettime(value);
			}
		},
		computed: {
			...mapState({
				user: state => state.user
			})
		},
	
		methods: {
			navigateToPage() {
				// 使用 uni.navigateTo 方法进行页面跳转
				uni.navigateTo({
					url: '../../pages/user-space/user-space', // 替换为你想要跳转到的页面路径
				});
			},
			myfollow(user_id){
				console.log('关注了',user_id)
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
			mylike(post_id) {
				let login_id=1;
				console.log(login_id,'点赞了帖子',post_id)
				uni.request({
				    url: 'http://120.46.81.37:8080/app/like',
				    method: 'POST',
				    data: {
				        'post_id': post_id,
				    },
					header: {
					    'Content-Type': 'application/x-www-form-urlencoded',
						'aToken': this.aToken,
						'rToken':this.rToken,
					},
				    success: function (res) {
				        console.log('Post request successful:', res.data);
						
						if(res.data==1){
							//点赞
							this.item.like++;
							document.getElementById('theLike').classList.remove('icon-dianzan2');
							document.getElementById('theLike').classList.add('icon-dianzan1');
						}else if(res.data==0){
							//取消点赞
						this.item.like--;
							document.getElementById('theLike').classList.remove('icon-dianzan1');
							document.getElementById('theLike').classList.add('icon-dianzan2');
						}
				        // 可以在这里处理返回的数据
				    },
				    fail: function (err) {
				        console.error('Post request failed:', err);
				    }
			})
			},
			mycollect(post_id) {
				let login_id=1;
				console.log(login_id,'收藏了帖子',post_id)
				uni.request({
				    url: 'http://120.46.81.37:8080/app/collect',
				    method: 'POST',
				  data: {
				      'post_id': post_id,
				  },
			header: {
				    'Content-Type': 'application/x-www-form-urlencoded',
					'aToken': this.aToken,
					'rToken':this.rToken,
				},
				    success: function (res) {
				        console.log('Post request successful:', res.data);
						if(res.data==1){
							//收藏
							this.item.collect++;
							document.getElementById('theCollect').classList.remove('icon-xihuan1');
							document.getElementById('theCollect').classList.add('icon-xihuan');
						}else if(res.data==0){
							//取消收藏
							this.item.collect--;
							document.getElementById('theCollect').classList.remove('icon-xihuan');
							document.getElementById('theCollect').classList.add('icon-xihuan1');
						}
				    },
				    fail: function (err) {
				        console.error('Post request failed:', err);
				    }
				});
			
			},
			// 打开个人空间
			openSpace(user_id) {
				uni.navigateTo({
					url: '/pages/user-space/user-space?user_id=' + user_id,
				});
			},
			// 进入详情页
			openDetail() {
				console.log('哈')
				// 处于详情中
				if (this.isdetail) return;
				// uni.navigateTo({
				// 	url: '../../pages/detail/detail?detail='+JSON.stringify(this.item),
				// });
				uni.navigateTo({
					url: `../../pages/detail/detail?detail=${encodeURIComponent(JSON.stringify(this.item))}`,
				});
			},
			// 进入详情页
			//   async openDetail() {
			//     console.log('进入详情页');
			//     // 检查是否已经处于详情中
			//     if (this.isdetail) return;

			//     // 准备发送到后端的用户ID
			//     const userId = this.item.user_id; // 确保 this.item 有 user_id 属性

			//     try {
			//       // 发送请求到后端判断用户是否VIP
			//       const response = await uni.request({
			//         url: `/vip/${userId}`, // 后端提供的相对路径
			//         method: 'GET',
			//         header: {
			//           'content-type': 'application/json', // 根据服务器要求设置请求头
			//         },
			//       });

			//       // 检查后端返回的状态码
			//       if (response.statusCode === 200) {
			//         // 解析返回的JSON数据
			//         const result = response.data;

			//         // 判断用户是否是VIP
			//         if (result && result.isVip) {
			//           // 如果是VIP，则跳转到详情页
			//           uni.navigateTo({
			//             url: `../../pages/detail/detail?detail=${encodeURIComponent(JSON.stringify(this.item))}`,
			//           });
			//         } else {
			//           // 如果不是VIP，可以在这里处理非VIP用户的逻辑，比如提示用户
			//           uni.showToast({
			//             title: '您没有权限查看详情',
			//             icon: 'none',
			//           });
			//         }
			//       } else {
			//         // 服务器返回非200状态码的处理逻辑
			//         uni.showToast({
			//           title: '服务器异常，请稍后再试',
			//           icon: 'none',
			//         });
			//       }
			//     } catch (error) {
			//       // 网络请求出错的处理逻辑
			//       console.error('请求失败:', error);
			//       uni.showToast({
			//         title: '网络请求失败，请检查网络',
			//         icon: 'none',
			//       });
			//     }
			//   },
			// 顶踩操作
			doSupport(type) {
				this.checkAuth(() => {
					this.$H.post('/support', {
						post_id: this.item.id,
						type: type === 'support' ? 0 : 1
					}, {
						token: true,
						native: true
					}).then(res => {
						if (res.data.errorCode) {
							return uni.showToast({
								title: res.data.msg,
								icon: 'none'
							});
						}
						console.log('通知父组件');
						// 通知父组件
						uni.$emit('updateFollowOrSupport', {
							type: "support",
							data: {
								type: type,
								id: this.item.id
							}
						})
					})
				})
			},
			// 评论
			doComment() {
				this.checkAuth(() => {
					if (!this.isdetail) {
						return this.openDetail()
					}
					this.$emit('doComment')
				})
			},
			// 分享
			doShare() {
				if (!this.isdetail) {
					return this.openDetail()
				}
				this.$emit('doShare')
			}
		},
	}
</script>

<style>
	.support-active {
		color: #FF4A6A;
	}
</style>