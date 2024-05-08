<template>
	<view>
		
		<!-- #ifdef MP -->

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
			
	<common-list :item="myitem" v-for="myitem in posts" :key="myitem.post_id">

	</common-list>
		<br><br>

		</view>
		 
		
	</view>
</template>

<script>
	import commonList from '@/components/common/common-list.vue';
	import loadMore from '@/components/common/load-more.vue';
	import uniNavBar from '@/components/uni-ui/uni-nav-bar/uni-nav-bar.vue';
	import uniPopup  from '@/components/uni-ui/uni-popup/uni-popup.vue';
	import search from '@/pages/search/search.vue'
	
	export default {
		components: {
			commonList,
			loadMore,
			uniNavBar,
			uniPopup,
			search
		},
		data() {
			return {
				search_text:'',
				showModal:true,
				posts:[
					{
					user_id:8,
					post_id:4,
					name:'Super girl',
					visibility:1,
					time:'2024-05-05',
					like:0,
					userpic:'../../static/user_pic/4.jpg',
					title: 'Cool day',
					collect:0,
					content:'Cycling by the lake',
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
					    30.769209]
					},
					{
					user_id:7,
					post_id:1,
					name:'Yodo',
					visibility:1,
					time:'2024-05-01',
					like:7,
					userpic:'../../static/user_pic/1.jpg',
					title: 'Nice day to have a walk',
					collect:7,
					content:"It's really nice",
					comments_num:0,
					center: [
    103.981051,
    30.765612
  ],
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
  ],
 
],comments: [
				
					{userpic:'../../static/user_pic/2.jpg',username:'Vero',data:'Good!',time:'2024-4-7'},
					
				],
				 post_pic:[]},
				{	user_id:6,
					post_id:2,
					name:'Veroooo',
					visibility:0,
					time:'2024-04-13',
					like:6,
					title: 'Running on the playground',
					userpic:'../../static/user_pic/2.jpg',
					collect:3,
					content:"It's so hot!",
					comments_num:1,
					path:[
  [
    103.986193,
    30.757726
  ],
  [
    103.986179,
    30.757617
  ],
  [
    103.986179,
    30.757559
  ],
  [
    103.986206,
    30.757415
  ],
  [
    103.98619,
    30.757272
  ],
  [
    103.986195,
    30.757178
  ],
  [
    103.986202,
    30.757023
  ],
  [
    103.986202,
    30.75693
  ],
  [
    103.986211,
    30.756844
  ],
  [
    103.986247,
    30.756751
  ],
  [
    103.986292,
    30.756662
  ],
  [
    103.986337,
    30.756627
  ],
  [
    103.986408,
    30.756593
  ],
  [
    103.98644,
    30.756565
  ],
  [
    103.986505,
    30.756551
  ],
  [
    103.98661,
    30.756544
  ],
  [
    103.986666,
    30.756551
  ],
  [
    103.986812,
    30.756579
  ],
  [
    103.986812,
    30.756607
  ],
  [
    103.98686,
    30.756641
  ],
  [
    103.986892,
    30.756662
  ],
  [
    103.986933,
    30.756711
  ],
  [
    103.986933,
    30.756746
  ],
  [
    103.986957,
    30.756794
  ],
  [
    103.986965,
    30.756822
  ],
  [
    103.986973,
    30.75687
  ],
  [
    103.986973,
    30.756926
  ],
  [
    103.986965,
    30.757023
  ],
  [
    103.986981,
    30.757134
  ],
  [
    103.986989,
    30.757204
  ],
  [
    103.986981,
    30.757273
  ],
  [
    103.986981,
    30.757349
  ],
  [
    103.986981,
    30.757488
  ],
  [
    103.986981,
    30.757564
  ],
  [
    103.986941,
    30.757655
  ],
  [
    103.986949,
    30.757703
  ],
  [
    103.986941,
    30.757745
  ],
  [
    103.986892,
    30.7578
  ],
  [
    103.98682,
    30.757869
  ],
  [
    103.986771,
    30.757918
  ],
  [
    103.98669,
    30.757918
  ],
  [
    103.986602,
    30.757945
  ],
  [
    103.986529,
    30.757931
  ],
  [
    103.986481,
    30.757931
  ],
  [
    103.9864,
    30.757918
  ],
  [
    103.986351,
    30.757883
  ],
  [
    103.986303,
    30.757841
  ],
  [
    103.986254,
    30.7578
  ],
  [
    103.986206,
    30.757723
  ],
  [
    103.986166,
    30.757523
  ],
  
],
					center:[
    103.9864,
    30.757418
  ],
  post_pic:['../../static/post_pic/3.jpg'],
  comments: [
  					{userpic:'../../static/user_pic/1.jpg',username:'Yodo',data:'Running is cool!',time:'2024-4-14'},
  				
  					
  				],
				
										
					
				
				},	{
					user_id:2,
					post_id:3,
					name:'Peace of summer',
					visibility:1,
					time:'2024-04-06',
					like:5,
					title: 'I like lovely swan',
					userpic:'../../static/user_pic/3.jpg',
					post_pic:['../../static/post_pic/1.jpg','../../static/post_pic/2.jpg'],
					collect:8,
					content:'Swan cygnets are absolutely adorable! Their fluffy down feathers and graceful movements make them a delightful sight to behold. Watching them glide across the water with their parents is truly a heartwarming experience',
					comments_num:2,
					path:[
  [
    103.984276,
    30.768216
  ],
  [
    103.984082,
    30.768406
  ],
  [
    103.983949,
    30.76852
  ],
  [
    103.983872,
    30.768615
  ],
  [
    103.983684,
    30.768767
  ],
  [
    103.983529,
    30.768891
  ],
  [
    103.983462,
    30.768995
  ],
  [
    103.983308,
    30.769176
  ],
  [
    103.983263,
    30.769185
  ],
  [
    103.98312,
    30.7691
  ],
  [
    103.982998,
    30.768967
  ],
  [
    103.982832,
    30.768853
  ],
  [
    103.982511,
    30.768625
  ],
  [
    103.982368,
    30.768511
  ],
  [
    103.982312,
    30.768492
  ],
  [
    103.982279,
    30.768492
  ],
  [
    103.982235,
    30.768596
  ],
  [
    103.982191,
    30.768663
  ],
  [
    103.982158,
    30.768653
  ],
  [
    103.981947,
    30.768587
  ],
  [
    103.981925,
    30.76853
  ],
  [
    103.981815,
    30.768463
  ],
  [
    103.981793,
    30.768397
  ],
  [
    103.981815,
    30.768302
  ]
],
					 center: [
    103.982191,
    30.768663
  ],comments: [
					{userpic:'../../static/user_pic/1.jpg',username:'Yodo',data:'hahaha',time:'2024-4-8'},
					{userpic:'../../static/user_pic/2.jpg',username:'Vero',data:'Beautiful swan!',time:'2024-4-7'},
					
				],
										
					
				
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
		
		onLoad() {
			let user =this.$store.state.user;
			console.log('获取user',user)
	uni.request({
		url: 'http://120.46.81.37:8080/app/index',
		method: 'GET',
		data: {
			'aToken': aToken,
			'rToken':rToken,
		},
		header: {
			'Content-Type': 'application/x-www-form-urlencoded',
			
		},
	      success: (res) => {
	        this.posts = res.data;
	      },
	      fail: (err) => {
	        console.error('Failed to fetch posts:', err);
	      }
	    });
		
		},
		methods: {
			

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
