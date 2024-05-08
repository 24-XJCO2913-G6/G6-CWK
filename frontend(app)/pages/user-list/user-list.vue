<template>
	<view>
		<!-- #ifdef MP -->
		<uni-nav-bar :shadow="false" :border="false" @click-right="goBack">
			<!-- 中间搜索框 -->
			<view class="flex justify-center align-center rounded text-muted bg-light flex-1 mt-1" style="margin-left: -46upx;height: 60upx;" @tap="openSearch">
				<view class="iconfont icon-sousuo mr-1"></view>搜索帖子
			</view>
			<!-- 右边图标 -->
			<block slot="right">
				<view class="text-dark font">取消</view>
			</block>
		</uni-nav-bar>
		<!-- #endif -->
		<!-- tab -->
		<view class="flex align-center" style="height: 100rpx;">
			<view class="flex-1 flex align-center justify-center"
			v-for="(item,index) in tabBars" :key="index"
			:class="index === tabIndex ? 'font-lg font-weight-bold text-main':'font-md'"
			@click="changeTab(index)">
			{{item.name}} <text v-if="item.num > 0" class="ml-2">{{item.num|formatNum}}</text></view>
		</view>
		
		<!-- 列表 -->
		<swiper :duration="150" :current="tabIndex" @change="onChangeTab"
		:style="'height:'+scrollH+'px;'">
			<swiper-item v-for="(item,index) in newsList" :key="index">
				<scroll-view scroll-y="true" :style="'height:'+scrollH+'px;'"
				@scrolltolower="loadmore(index)">
					<template v-if="item.list.length > 0">
						<!-- 列表 -->
						<block v-for="(item2,index2) in item.list" :key="index2">
							
							<!-- 列表样式 -->
							<user-list :item="item2" :index="index"></user-list>
							
							
						</block>
						<!-- 上拉加载 -->
						<load-more v-if="item.list.length > 10" :loadmore="item.loadmore"></load-more>
					</template>
					<!-- 无数据 -->
					<template v-else>
						<no-thing></no-thing>
					</template>
				</scroll-view>
			</swiper-item>
		</swiper>
		
		
	</view>
</template>

<script>
	
	import loadMore from '@/components/common/load-more.vue';
	import noThing from '@/components/common/no-thing.vue';
	import userList from '@/components/user-list/user-list.vue';
	import uniNavBar from '@/components/uni-ui/uni-nav-bar/uni-nav-bar.vue';
	export default {
		components: {
			loadMore,
			noThing,
			userList,
			uniNavBar
		},
		data() {
			return {
				// 列表高度
				scrollH:600,
				tabIndex:0,
				tabBars:[{
					name:"Friends",
					num:0,
					key:"friends"
				},{
					name:"Follows",
					num:0,
					key:"follows"
				},{
					name:"Fans",
					num:0,
					key:"fens"
				}],
				
				newsList:[]
			}
		},
		// 监听点击输入框事件
		onNavigationBarSearchInputClicked() {
			this.openSearch()
		},
		onNavigationBarButtonTap() {
			uni.navigateBack({
				delta: 1
			});
		},
		onLoad() {
			uni.getSystemInfo({
				success:res=>{
					this.scrollH = res.windowHeight - uni.upx2px(100)
					// #ifdef MP
					this.scrollH -= 44
					// #endif
				}
			})
			// 根据选项生成列表
			this.getData()
		},
		filters: {
			formatNum(value) {
				return value > 99 ? '99+' : value;
			}
		},
		methods: {
			openSearch(){
				uni.navigateTo({
					url: '../search/search?type=user',
				});
			},
			// 获取用户相关数据
			getCounts(){
				this.$H.get('/user/getcounts/'+this.user.id,{},{
					token:true,
					notoast:true
				}).then(res=>{
					this.tabBars[0].num = res.friend_count
					this.tabBars[1].num = res.withfollow_count
					this.tabBars[2].num = res.withfen_count
				})
			},
			// 获取数据
			getData(){
				var arr = []
				for (let i = 0; i < this.tabBars.length; i++) {
					// 生成列表模板
					let obj = {
						// 1.上拉加载更多  2.加载中... 3.没有更多了
						loadmore:"上拉加载更多",
						list:[],
						page:1,
						firstLoad:false
					}
					arr.push(obj)
				}
				this.newsList = arr
				this.getList()
			},
			getList(){
				let index = this.tabIndex
				let id = this.tabBars[index].id
				let page = this.newsList[index].page
				let isrefresh = page === 1
				this.$H.get('/'+this.tabBars[index].key+'/'+page,{},{
					token:true,
					noCheck:true
				})
				.then(res=>{
					console.log(res);
					let list = res.list.map(v=>{
						return {
							id:v.id,
							avatar:v.userpic,
							username:v.username,
							sex:v.userinfo.sex,
							age:v.userinfo.age,
							isFollow:index !== 2
						}
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
			// tab切换
			changeTab(index){
				this.tabIndex = index
			},
			// 监听滑动
			onChangeTab(e){
				this.changeTab(e.detail.current)
				if(!this.newsList[e.detail.current].firstLoad){
					this.getList()
				}
			},
			// 上拉加载更多
			loadmore(index){
				// 拿到当前列表
				let item = this.newsList[index]
				// 判断是否处于可加载状态（上拉加载更多）
				if (item.loadmore !== '上拉加载更多') return;
				// 修改当前列表加载状态
				item.loadmore = '加载中...'
				// 数据请求
				item.page++
				this.getList()
			},
			goBack(){
				uni.navigateBack({
					delta: 1
				});
			}
		}
	}
</script>

<style>

</style>
