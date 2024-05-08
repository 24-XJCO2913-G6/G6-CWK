<template>
	<view>
		<!-- 帖子详情页详细信息 -->
		<common-list :item="info"  @follow="follow" @doSupport="doSupport" @doComment="doComment"
			@doShare="doShare">
			<image v-for="item of info.post_pic" :src="item"
						style="height: 350rpx; margin-bottom: 10px;margin-top:10px" class="rounded w-100" mode="aspectFill"></image>
					{{info.content}}
		</common-list>
		
		<divider></divider>
		<!-- <view class="p-2 font-md font-weight-bold">
			Comment {{comments.length}}
		</view> -->
		<view class="px-2">

			<view class="uni-comment-list" v-for="(item,index) in info.comments" :key="index">
				<view v-if="item.fid" style="width: 60rpx;"></view>
				<view class="flex w-100" >
					<view >
						<image :src="item.userpic" class="rounded-circle mr-2" style="width: 65rpx;height: 65rpx;" ></image>
					</view>
					<view class="uni-comment-body">
						<view class="uni-comment-top">
							<text>{{item.username}}</text>
						</view>
						<view class="uni-comment-content" @click="reply(item.id)">{{item.data}}</view>
						<view class="uni-comment-date">
							<view>{{item.time}}</view>
						</view>
					</view> 
				</view>
			</view>

		</view>

		<!-- 占位 -->
		<view style="height: 100rpx;"></view>
		<bottom-input :focus="focus" @blur="blur" @submit="submit"></bottom-input>

		<more-share ref="share"></more-share>

	</view>
</template>

<script>
	import commonList from '@/components/common/common-list.vue';
	import bottomInput from '@/components/common/bottom-input.vue';
	import moreShare from '@/components/common/more-share.vue';
	import mymap from '../../components/map/mymap.vue';
	import $T from '@/common/time.js';
	export default {
		components: {
			commonList,
			bottomInput,
			moreShare,
			mymap
		},
		data() {
			return {
				// 当前帖子信息
				info: {
					id: 15,
					user_id: 0,
					username: "昵称",
					userpic: "",
					newstime: 0,
					isFollow: false,
					title: "",
					content:"wow",
					titlepic: "",
					support: {
						type: "support", // 顶
						support_count: 0,
						unsupport_count: 0
					},
					comment_count: 0,
					share_num: 0,
				},
				images: [],
			
				focus: false,
				reply_id: 0
			}
		},
		onLoad(e) {
			
			// 初始化
			if (e.detail) {
				this.__init(JSON.parse(e.detail))
			}
			// 监听关注和顶踩操作
			uni.$on('updateFollowOrSupport', (e) => {
				switch (e.type) {
					case 'follow': // 关注
						this.follow(e.data.user_id)
						break;
					case 'support': // 顶踩
						this.doSupport(e.data)
						break;
				}
			})
		},
		mounted(){
	console.log(this.info);		
		},
		onUnload() {
			uni.$off('updateFollowOrSupport', (e) => {})
		},
		computed: {
			imagesList() {
				return this.images.map(item => item.url)
			}
		},
		onNavigationBarButtonTap() {
			this.$refs.share.open({
				title: this.info.title,
				shareText: this.info.title,
				href: "http://www.dishaxy.com",
				image: this.info.titlepic
			})
		},
		onBackPress() {
			this.$refs.share.close()
		},
		// #ifndef APP-PLUS
		// 微信小程序分享
		onShareAppMessage(res) {
			return {
				title: this.info.title,
				path: '/pages/detail/detail?detail=' + JSON.stringify(this.info)
			}
		},
		// #endif
		methods: {
			__init(data) {
				// 修改标题
				uni.setNavigationBarTitle({
					title: data.title
				})
				this.info = data
				
				
			},
			// 提交评论
			submit(data) {
				this.comments++;
				this.info.comments.unshift({userpic:'../../static/user_pic/4.jpg',username:'Super girl',data:data,time:'2024-4-9'})
			},
		
			// 输入框失去焦点
			blur() {
				this.reply_id = 0
				this.focus = false
			}
		}
	}
</script>

<style>

</style>