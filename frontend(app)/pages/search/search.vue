<template>
	<view>
		<!-- 自定义导航 -->
		<!-- #ifndef APP-PLUS -->
		<view style="height: 88rpx;z-index: 9999;"
		class="flex align-center px-2 position-fixed left-0 top-0 right-0 bg-white">
			<view class="iconfont icon-sousuo position-absolute text-muted" 
			style="left: 30rpx;"></view>
			<input class="flex-1 rounded bg-light" style="padding: 5rpx 0 5rpx 50rpx;" type="text" v-model="searchText" @confirm="searchEvent"
			placeholder="search" 
			placeholder-style="color: #CCCCCC;"/>
			
			<text class="pl-2" @click="goBack">取消</text>
		</view>
		<view style="height: 88rpx;">
		</view>
		<!-- #endif -->
		<template >
			<!-- 数据列表 -->
			<block v-for="(item,index) in searchList" :key="index">		
				<template >
					<!-- 用户 -->
					<user-list :item="item" :index="index"></user-list>
				</template>
			</block>
		</template>
		
	</view>
</template>

<script>

	import commonList from '@/components/common/common-list.vue';
	import topicList from '@/components/news/topic-list.vue';
	import userList from '@/components/user-list/user-list.vue';
	import loadMore from '@/components/common/load-more.vue';
	import { mapState } from 'vuex'
	export default {
		components: {
			commonList,
			topicList,
			userList,
			loadMore
		},
		data() {
			return {
				searchText:"",
				userList:[
					{
					user_id:1,
					username:'Yodo',
					user_pic:'http://120.46.81.37:8080/app/static/user_pic/1.jpg',
					},
					{
					user_id:2,
					username:'Verooo',
					user_pic:'http://120.46.81.37:8080/app/static/user_pic/2.jpg',
					},
					{
					user_id:3,
					username:'Peace of summer',
					user_pic:'http://120.46.81.37:8080/app/static/user_pic/3.jpg',
					}
					],

				type:"post",
				loadmore:"上拉加载更多",
				page:1
			}
		},
		computed:{
			searchList(){
				if(this.searchText){
				return this.userList.filter(x=>{
					return x.username.includes(this.searchText);
				})
				}
				else{
					return []
				}
				
				
			}
			
		},
		// 监听导航输入
		onNavigationBarSearchInputChanged(e){
			this.searchText = e.text
		},
		// 监听点击导航搜索按钮
		onNavigationBarButtonTap(e) {
			if (e.index === 0) {
				this.searchEvent()
			}
		},
		onLoad() {
			uni.request({
			url: 'http://120.46.81.37:8080/app/all_users',
				method: 'GET',
				data: {	 
						  'aToken': this.aToken,
						  'rToken':this.rToken,
				},
					header: {
						'Content-Type': 'application/x-www-form-urlencoded',
					},
			  success: (res) => {
				this.userList= res.data;
			  },
			  fail: (err) => {
				console.error('Failed to fetch posts:', err);
			  }
			    })
		},
		onUnload() {
			if(this.type === 'post'){
				uni.$off('updateFollowOrSupport',(e)=>{})
			}
		},
		// 监听下拉刷新
		onPullDownRefresh() {
			if(this.searchText === ''){
				return uni.stopPullDownRefresh()
			}
			this.getData(true,()=>{
				// 关闭下拉刷新状态
				uni.stopPullDownRefresh()
			})
		},
		// 监听上拉加载
		onReachBottom() {
			if(this.loadmore !== '上拉加载更多'){
				return;
			}
			this.loadmore = "加载中..."
			this.getData(false)
		},
		methods: {
			// #ifndef APP-PLUS
			goBack(){
				uni.navigateBack({
					delta: 1
				});
			},
			// #endif
			// 关注
			follow(user_id){
				// 找到当前作者的所有列表
				this.searchList.forEach((item)=>{
					if(item.user_id === user_id){
						item.isFollow = true
					}
				})
				uni.showToast({ title: '关注成功' })
			},
			// 顶踩操作
			doSupport(e){
				// 拿到当前的选项卡对应的list
				this.searchList.forEach(item=>{
					if(item.id === e.id){
						// 之前没有操作过
						if (item.support.type === '') {
							item.support[e.type+'_count']++
						} else if (item.support.type ==='support' && e.type === 'unsupport') {
							// 顶 - 1
							item.support.support_count--;
							// 踩 + 1
							item.support.unsupport_count++;
						} else if(item.support.type ==='unsupport' && e.type === 'support'){ 					// 之前踩了
							// 顶 + 1
							item.support.support_count++;
							// 踩 - 1
							item.support.unsupport_count--;
						}
						item.support.type = e.type
					}
				})
				let msg = e.type === 'support' ? '顶' : '踩'
				uni.showToast({ title: msg + '成功' });
			},
			// 点击搜索历史
			clickSearchHistory(text){
				this.searchText = text
				this.searchEvent()
			},
			// 搜索事件
			searchEvent(){
				// 收起键盘
				uni.hideKeyboard()
				// 加入搜索历史
				let index = this.list.findIndex(v => v===this.searchText)
				if(index !== -1){
					this.$U.__toFirst(this.list,index)
				} else {
					this.list.unshift(this.searchText)
				}
				uni.setStorageSync('historySeachText',JSON.stringify(this.list))
				// 请求搜索
				this.getData()
			},
			// 获取数据
			getData(isrefresh = true,callback = false){
				// 显示loading状态
				uni.showLoading({
					title: '加载中...',
					mask: false
				})
				// 请求搜索
				this.page = isrefresh ? 1 : (this.page + 1)
				this.$H.post('/search/'+this.type,{
					keyword:this.searchText,
					page:this.page
				}).then(res=>{
					// 整理格式
					let list = []
					switch (this.type){
						case 'post':
						list = res.list.map(v=>{
							return this.$U.formatCommonList(v)
						})
							break;
						case 'topic':
						list = res.list.map(v=>{
							return {
								id:v.id,
								cover:v.titlepic,
								title:v.title,
								desc:v.desc,
								today_count:v.todaypost_count,
								news_count:v.post_count
							}
						})
							break;
						case 'user':
						list = res.list.map(v=>{
							return {
								id:v.id,
								avatar:v.userpic,
								username:v.username,
								sex:v.userinfo.sex,
								age:v.userinfo.age,
								isFollow:false
							}
						})
							break;
					}
					console.log(JSON.stringify(list));
					// 渲染页面
					this.searchList = isrefresh ? [...list] : [...this.searchList,...list]
					// 加载情况
					this.loadmore = list.length < 10 ? '没有更多了' : '上拉加载更多'
					//隐藏loading
					uni.hideLoading()
					if(typeof callback === 'function'){
						callback()
					}
				}).catch(err=>{
					this.page--
					//隐藏loading
					uni.hideLoading()
					if(typeof callback === 'function'){
						callback()
					}
				})
			}
		}
	}
</script>

<style>

</style>
