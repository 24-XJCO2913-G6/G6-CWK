<template>
	<view>
		  <view class="modal" v-if="showModal" style='padding:10px; '>
		     <view class="modal-content">
		       <text style='font-size: 30px;font-weight: bold;text-shadow: 2px 2px black;'>Welcome to Been</text>
			   <p style='padding:10px;font-size:20px; background-color: rgba(255,255,255,0.5);color:black'>We are thrilled to have you join our vibrant community of travelers and sports enthusiasts. Explore, connect, and share your adventures with like-minded individuals around the world. Let's embark on unforgettable journeys together. Enjoy your stay! 🌍🎉</p>
		       <button @click="showModal=false">Close</button>
		     </view>
		   </view>
		<!-- #ifdef MP -->
		<uni-nav-bar :shadow="false" :border="false" 
		@click-left="clickLeft" @click-right="clickRight">
			<!-- 左边图标 -->
			<block slot="left">
				<view class="iconfont icon-qiandao ml-2 mr-2" style="font-size: 22px;color: #FF9619;"></view>
			</block>
			<!-- 中间搜索框 -->
			<view class="flex justify-center align-center rounded text-muted bg-light flex-1 mt-1" style="margin-left: -46upx;height: 60upx;" @tap="openSearch">
				<view class="iconfont icon-sousuo mr-1"></view>搜索帖子
			</view>
			<!-- 右边图标 -->
			<block slot="right">
				<view class="icon iconfont icon-bianji1 text-dark" style="font-size: 22px;"></view>
			</block>
		</uni-nav-bar>
		<!-- #endif -->
		<!-- 顶部选项卡 -->
	<!-- 	<scroll-view scroll-x :scroll-into-view="scrollInto" scroll-with-animation
		class="scroll-row border-bottom border-light-secondary" 
		style="height: 100rpx;">
			<view v-for="(item,index) in tabBars" :key="index" 
			class="scroll-row-item px-3 py-2 font-md" :id="'tab'+index"
			:class="tabIndex === index?'text-main font-lg font-weight-bold':''"
			@click="changeTab(index)">{{item.classname}}</view>
		</scroll-view> -->
		<view>
			
	<common-list :item="myitem" v-for="myitem in posts" :key="myitem.post_id"></common-list>
		111

		</view>
		 
		
	</view>
</template>

<script>
	import commonList from '@/components/common/common-list.vue';
	import loadMore from '@/components/common/load-more.vue';
	import uniNavBar from '@/components/uni-ui/uni-nav-bar/uni-nav-bar.vue';
	import uniPopup  from '@/components/uni-ui/uni-popup/uni-popup.vue';
	export default {
		components: {
			commonList,
			loadMore,
			uniNavBar,
			uniPopup,
		},
		data() {
			return {
				showModal:true,
				posts:[
					{
					post_id:1,
					name:'Alice',
					visibility:1,
					time:'2022-01-12',
					like:19,
					userpic:'',
					title: 'WOw',
					collect:13,
					content:'44444444444444444',
					comments:5,
					center:[103.985895, 30.765873]	,
					path:[
									    [103.985895, 30.763873], // 起点坐标
									    [103.986895, 30.764873], // 中间很多点坐标
									    [103.987895, 30.766573]],
				},
				{	
					post_id:2,
					name:'Jack',
					visibility:0,
					time:'2022-01-12',
					like:20,
					title: 'Nice place to go',
					userpic:'',
					collect:3,
					content:'33333333333333333',
					comments:5,
					path:[
					  [103.983895, 30.763873], // 起点坐标
					  [103.987895, 30.764874], // 中间点坐标
					  [103.988894, 30.763873]  // 终点坐标
					 ],
					center:[103.987895, 30.764874], 
				
										
					
				
				},	{
					post_id:3,
					name:'Mike',
					visibility:1,
					time:'2022-01-12',
					like:5,
					title: 'I like it',
					userpic:'',
					collect:8,
					content:'121212121212121',
					comments:5,
					path:[
					 [103.982895, 30.763873], // 起点坐标
					 [103.986895, 30.765574], // 中间点坐标
					 [103.987694, 30.766872]  // 终点坐标
					],
					 center: [103.983895, 30.765574], 
										
					
				
				}],
				// 列表高度
				scrollH:600,
				// 顶部选项卡
				scrollInto:"",
				tabIndex:0,
				tabBars: [],
				newsList:[]
			}
		},
		// 监听点击导航栏搜索框
		onNavigationBarSearchInputClicked() {
			uni.navigateTo({
				url: '../search/search?type=post',
			})
		},
		// 监听导航按钮点击事件
		onNavigationBarButtonTap() {
			this.navigateTo({
				url: '../add-input/add-input',
			})
		},
		
		onLoad() {
			uni.getSystemInfo({
				success:res=>{
					this.scrollH = res.windowHeight - uni.upx2px(101)
					// #ifdef MP
					this.scrollH -= 44
					// #endif
				}
			})
			// 监听刷新首页
			uni.$on('updateIndex',()=>{
				this.getData()
			})
			// 根据选项生成列表
			this.getData()
			// 监听关注和顶踩操作
			uni.$on('updateFollowOrSupport',(e)=>{
				// console.log('接收到了');
				switch (e.type){
					case 'follow': // 关注
					this.follow(e.data.user_id)
						break;
					case 'support': // 顶踩
					this.doSupport(e.data)
						break;
				}
			})
			// 监听评论数变化
			uni.$on('updateCommentsCount',({id,count})=>{
				this.newsList.forEach(tab=>{
					tab.list.forEach((item)=>{
						if(item.id === id){
							item.comment_count = count
						}
					})
				})
			})
		},
		onUnload() {
			uni.$off('updateFollowOrSupport',(e)=>{})
			uni.$off('updateIndex',(e)=>{})
			uni.$off('updateCommentsCount',(e)=>{})
		},
		methods: {
			 closeModal() {
			      this.showModal = false;
			    },
			// #ifndef APP-PLUS
			clickLeft(){
				// console.log('左边事件')
			},
			clickRight(){
				// 打开发布页面
				this.navigateTo({
					url: '../add-input/add-input',
				})
			},
			openSearch(){
				uni.navigateTo({
					url: '../search/search',
				});
			},
			// #endif
			// 获取数据
			getData(){
				// 获取分类
				this.$H.get('/postclass').then(res=>{
					this.tabBars = res.list
					// 根据分类生成列表
					var arr = []
					for (let i = 0; i < this.tabBars.length; i++) {
						// 生成列表模板
						arr.push({
							// 1.上拉加载更多  2.加载中... 3.没有更多了
							loadmore:"上拉加载更多",
							list:[],
							page:1,
							firstLoad:false
						})
					}
					this.newsList = arr
					console.log(arr[0])
					console.log(arr[0].list)
					// 获取第一个分类的数据
					if (this.tabBars.length) {
						this.getList()
					}
				})

			},
			// 获取指定分类下的列表
			getList(){
				let index = this.tabIndex
				let id = this.tabBars[index].id
				let page = this.newsList[index].page
				let isrefresh = page === 1
				this.$H.get('/postclass/'+id+'/post/'+page,{},{
					token:true,
					noCheck:true
				})
				.then(res=>{
					let list = res.list.map(v=>{
						return this.$U.formatCommonList(v)
					})

					this.newsList[index].list = isrefresh ? list : [...this.newsList[index].list,...list];
					
					this.newsList[index].loadmore  = list.length < 10 ? '没有更多了' : '上拉加载更多';
					
					if (isrefresh) {
						this.newsList[index].firstLoad = true
					}
				}).catch(err=>{
					if(!isrefresh){
						this.newsList[index].page--;
					}
				})
			},
			// 监听滑动
			onChangeTab(e){
				this.changeTab(e.detail.current)
			},
			// 切换选项
			changeTab(index){
				if (this.tabIndex === index) {
					return;
				}
				this.tabIndex = index
				// 滚动到指定元素
				this.scrollInto = 'tab'+index
				// 获取当前分类下的列表数据
				if (!this.newsList[this.tabIndex].firstLoad) {
					this.getList()
				}
			},
			// 关注
			follow(user_id){
				// 找到当前作者的所有列表
				this.newsList.forEach(tab=>{
					tab.list.forEach((item)=>{
						if(item.user_id === user_id){
							item.isFollow = true
						}
					})
				})
				uni.showToast({ title: '关注成功' })
			},
			// 顶踩操作
			doSupport(e){
				// 拿到当前的选项卡对应的list
				this.newsList[this.tabIndex].list.forEach(item=>{
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
			// 上拉加载更多
			loadmore(index){
				// 拿到当前列表
				let item = this.newsList[index]
				// 判断是否处于可加载状态（上拉加载更多）
				if (item.loadmore !== '上拉加载更多') return;
				// 修改当前列表加载状态
				item.loadmore = '加载中...'
				// 请求数据
				item.page++;
				this.getList()
			}
		}
	}
</script>

<style>

	.modal {
	    z-index: 100;
	    position: fixed;
	    top: 50%;
	    left: 50%;
	    transform: translate(-50%, -50%);
	    width: 80%;
	    height: 50%;
	    background-color: rgba(255, 255, 255, 0.9);
	    border: blue solid 1px;
	    border-radius: 10px;
	    display: flex;
	    justify-content: center;
	    color: white;
	    align-items: center;
	   
	}
	
	.modal::before {
	    content: "";
	    position: absolute;
	    top: 0;
	    left: 0;
	    width: 100%;
	    height: 100%;
	    background-image: url("../../static/images/landing.jpeg");
	    background-size: 100% 100%;
	    filter: blur(1px); /* 调整模糊程度 */
	    z-index: -1; /* 将虚化的背景置于底层 */
	}
	
	
	.modal-content {
	  text-align: center;
	}
	
	.text-welcome {
	  font-size: 24px;
	}
	
	button {
	  width: 100px;
	  padding: 5px 10px;
	  margin-top: 10px;
	}
</style>
