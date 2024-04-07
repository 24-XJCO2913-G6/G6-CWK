<template>
	<view>
		<!-- 自定义导航 -->
		<!-- #ifndef APP-PLUS -->
		<view style="height: 88rpx;z-index: 9999;"
		class="flex align-center px-2 position-fixed left-0 top-0 right-0 bg-white">
			<view class="iconfont icon-sousuo position-absolute text-muted" 
			style="left: 30rpx;"></view>
			<input class="flex-1 rounded bg-light" style="padding: 5rpx 0 5rpx 50rpx;" type="text" v-model="searchText" @confirm="searchEvent"
			placeholder="搜索" 
			placeholder-style="color: #CCCCCC;"/>
			<text class="pl-2" @click="goBack">取消</text>
		</view>
		<view style="height: 88rpx;"></view>
		<!-- #endif -->
		
		<template v-if="searchList.length === 0">
			<!-- 搜索历史 -->
			<view class="py-2 font-md px-2">搜索历史</view>
			<view class="flex flex-wrap">
				<view class="border rounded font mx-2 my-1 px-2" 
				v-for="(item,index) in list" :key="index"
				hover-class="bg-light"
				@click="clickSearchHistory(item)">{{item}}</view>
			</view>
		</template>
		<template v-else>
			<!-- 数据列表 -->
			<block v-for="(item,index) in searchList" :key="index">
				<template v-if="type ==='post'">
					<!-- 帖子 -->
					<common-list :item="item" :index="index"></common-list>
				</template>
				<template v-else-if="type === 'topic'">
					<!-- 话题 -->
					<topic-list :item="item" :index="index"></topic-list>
				</template>
				<template v-else>
					<!-- 用户 -->
					<user-list :item="item" :index="index"></user-list>
				</template>
			</block>
			
			<!-- 上拉加载 -->
			<load-more :loadmore="loadmore"></load-more>
		</template>
		
	</view>
</template>

<script>

	import commonList from '@/components/common/common-list.vue';
	import topicList from '@/components/news/topic-list.vue';
	import userList from '@/components/user-list/user-list.vue';
	import loadMore from '@/components/common/load-more.vue';
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
				list:[],
				// 搜索结果
				searchList:[],
				// 当前搜索类型
				type:"post",
				loadmore:"上拉加载更多",
				page:1
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
		onLoad(e) {
			if (e.type) {
				this.type = e.type
			}
			let pageTitle = '帖子'
			switch (this.type){
				case 'post':
				pageTitle = '帖子'
				// 监听关注和顶踩操作
				uni.$on('updateFollowOrSupport',(e)=>{
					switch (e.type){
						case 'follow': // 关注
						this.follow(e.data.user_id)
							break;
						default:
							break;
					}
				})
					break;
				case 'topic':
				pageTitle = '话题'
					break;
				case 'user':
				pageTitle = '用户'
					break;
			}
			// 修改搜索占位
			// #ifdef APP-PLUS
			let currentWebview = this.$mp.page.$getAppWebview();
			let tn = currentWebview.getStyle().titleNView; 
			tn.searchInput.placeholder = '搜索'+pageTitle; 
			currentWebview.setStyle({
				titleNView: tn  
			})
			// #endif
			// 取出搜索历史
			let list = uni.getStorageSync('historySeachText')
			if(list){
				this.list = JSON.parse(list)
			}
			
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
