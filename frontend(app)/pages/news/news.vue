<template>
	<view>

		<view>
		
		<input type="text"  style="position: fixed; top:8px;z-index: 999;display: flex; width: 90%; height: 2.8%; margin-left: 2.5%;
		  justify-content: center;
		  align-items: center;"  v-model='search_text'/>	
	<common-list :item="myitem" v-for="myitem in show_posts" :key="myitem.post_id">

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
	import { mapState } from 'vuex'
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
									Author:'Super girl',
									visibility:1,
									Pub_time:'2024-05-05',
									like:0,
									Photo:'http://120.46.81.37:8080/app/static/user_pic/4.jpg',
									title: 'Cool day',
									collect:0,
									content:'Cycling by the lake',
									comments_num:0,
									markers:[
										    {"longitude": 103.981308, "latitude": 30.768426, iconPath:'../../static/images/marker.png', width: 10,
				          height: 10}  
									],
									path:[{
									   points:[{"longitude": 103.976838, "latitude": 30.768159},
									    {"longitude": 103.976871, "latitude": 30.768094},
									    {"longitude": 103.976905, "latitude": 30.768044},
									    {"longitude": 103.976921, "latitude": 30.768023},
									    {"longitude": 103.976938, "latitude": 30.768001},
									    {"longitude": 103.97703, "latitude": 30.768073},
									    {"longitude": 103.977097, "latitude": 30.768109},
									    {"longitude": 103.977155, "latitude": 30.768166},
									    {"longitude": 103.977222, "latitude": 30.768195},
									    {"longitude": 103.977339, "latitude": 30.76831},
									    {"longitude": 103.977439, "latitude": 30.768382},
									    {"longitude": 103.977548, "latitude": 30.768439},
									    {"longitude": 103.977673, "latitude": 30.768511},
									    {"longitude": 103.977799, "latitude": 30.768597},
									    {"longitude": 103.977983, "latitude": 30.768726},
									    {"longitude": 103.978108, "latitude": 30.768827},
									    {"longitude": 103.97825, "latitude": 30.76892},
									    {"longitude": 103.9784, "latitude": 30.769042},
									    {"longitude": 103.978526, "latitude": 30.769143},
									    {"longitude": 103.978651, "latitude": 30.76925},
									    {"longitude": 103.978801, "latitude": 30.769372},
									    {"longitude": 103.978969, "latitude": 30.769523},
									    {"longitude": 103.979127, "latitude": 30.769652},
									    {"longitude": 103.979303, "latitude": 30.769439},
									    {"longitude": 103.979345, "latitude": 30.769381},
									    {"longitude": 103.979445, "latitude": 30.769252},
									    {"longitude": 103.979545, "latitude": 30.769108},
									    {"longitude": 103.97962, "latitude": 30.769008},
									    {"longitude": 103.979796, "latitude": 30.768843},
									    {"longitude": 103.97998, "latitude": 30.768606},
									    {"longitude": 103.980113, "latitude": 30.768448},
									    {"longitude": 103.980205, "latitude": 30.768276},
									    {"longitude": 103.980222, "latitude": 30.768204},
									    {"longitude": 103.980272, "latitude": 30.768211},
									    {"longitude": 103.980356, "latitude": 30.768247},
									    {"longitude": 103.980447, "latitude": 30.768326},
									    {"longitude": 103.980573, "latitude": 30.76839},
									    {"longitude": 103.980673, "latitude": 30.768441},
									    {"longitude": 103.980715, "latitude": 30.768484},
									    {"longitude": 103.980665, "latitude": 30.768577},
									    {"longitude": 103.98059, "latitude": 30.768721},
									    {"longitude": 103.980498, "latitude": 30.768879},
									    {"longitude": 103.980297, "latitude": 30.76908},
									    {"longitude": 103.980105, "latitude": 30.769302},
									    {"longitude": 103.98003, "latitude": 30.769431},
									    {"longitude": 103.979971, "latitude": 30.769525},
									    {"longitude": 103.979946, "latitude": 30.769661},
									    {"longitude": 103.979896, "latitude": 30.769733},
									    {"longitude": 103.979904, "latitude": 30.76979},
									    {"longitude": 103.980038, "latitude": 30.769776},
									    {"longitude": 103.980172, "latitude": 30.76969},
									    {"longitude": 103.980272, "latitude": 30.769546},
									    {"longitude": 103.980272, "latitude": 30.76941},
									    {"longitude": 103.980305, "latitude": 30.769302},
									    {"longitude": 103.980322, "latitude": 30.769209},
									    {"longitude": 103.980322, "latitude": 30.769159},
									    {"longitude": 103.980389, "latitude": 30.769116},
									    {"longitude": 103.980506, "latitude": 30.769015},
									    {"longitude": 103.980581, "latitude": 30.768979},
									    {"longitude": 103.980706, "latitude": 30.7689},
									    {"longitude": 103.980757, "latitude": 30.768871},
									    {"longitude": 103.980832, "latitude": 30.768814},
									    {"longitude": 103.98094, "latitude": 30.768771},
									    {"longitude": 103.981007, "latitude": 30.768713},
									    {"longitude": 103.981074, "latitude": 30.76867},
									    {"longitude": 103.981183, "latitude": 30.768584},
									    {"longitude": 103.981258, "latitude": 30.76852},
									    {"longitude": 103.981283, "latitude": 30.768469},
									    {"longitude": 103.981308, "latitude": 30.768426}  
									
									],color: "#2e7efd",
									width: 10,
									dottedLine: true,
									arrowLine: true,
									colorList: true,
									
									}],
															center: [
									    103.980322,
									    30.769209]
									},
									{
									user_id:7,
									post_id:1,
									Author:'Yodo',
									visibility:1,
									Pub_time:'2024-05-01',
									like:7,
									Photo:'http://120.46.81.37:8080/app/static/user_pic/1.jpg',
									title: 'Nice day to have a walk',
									collect:7,
									content:"It's really nice",
									comments_num:0,
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
				          height: 10}  
									],comments: [
								
									{userpic:'http://120.46.81.37:8080/app/static/user_pic/2.jpg',username:'Vero',data:'Good!',time:'2024-4-7'},
									
								],
								 post_pic:[]},
								{	user_id:6,
									post_id:2,
									Author:'Veroooo',
									visibility:0,
									Pub_time:'2024-04-13',
									like:6,
									title: 'Running on the playground',
									Photo:'http://120.46.81.37:8080/app/static/user_pic/2.jpg',
									collect:3,
									content:"It's so hot!",
									comments_num:1,
									path:[
										{
				
				points:[
				  {"longitude": 103.986193, "latitude": 30.757726},
				  {"longitude": 103.986179, "latitude": 30.757617},
				  {"longitude": 103.986179, "latitude": 30.757559},
				  {"longitude": 103.986206, "latitude": 30.757415},
				  {"longitude": 103.98619, "latitude": 30.757272},
				  {"longitude": 103.986195, "latitude": 30.757178},
				  {"longitude": 103.986202, "latitude": 30.757023},
				  {"longitude": 103.986202, "latitude": 30.75693},
				  {"longitude": 103.986211, "latitude": 30.756844},
				  {"longitude": 103.986247, "latitude": 30.756751},
				  {"longitude": 103.986292, "latitude": 30.756662},
				  {"longitude": 103.986337, "latitude": 30.756627},
				  {"longitude": 103.986408, "latitude": 30.756593},
				  {"longitude": 103.98644, "latitude": 30.756565},
				  {"longitude": 103.986505, "latitude": 30.756551},
				  {"longitude": 103.98661, "latitude": 30.756544},
				  {"longitude": 103.986666, "latitude": 30.756551},
				  {"longitude": 103.986812, "latitude": 30.756579},
				  {"longitude": 103.986812, "latitude": 30.756607},
				  {"longitude": 103.98686, "latitude": 30.756641},
				  {"longitude": 103.986892, "latitude": 30.756662},
				  {"longitude": 103.986933, "latitude": 30.756711},
				  {"longitude": 103.986933, "latitude": 30.756746},
				  {"longitude": 103.986957, "latitude": 30.756794},
				  {"longitude": 103.986965, "latitude": 30.756822},
				  {"longitude": 103.986973, "latitude": 30.75687},
				  {"longitude": 103.986973, "latitude": 30.756926},
				  {"longitude": 103.986965, "latitude": 30.757023},
				  {"longitude": 103.986981, "latitude": 30.757134},
				  {"longitude": 103.986989, "latitude": 30.757204},
				  {"longitude": 103.986981, "latitude": 30.757273},
				  {"longitude": 103.986981, "latitude": 30.757349},
				  {"longitude": 103.986981, "latitude": 30.757488},
				  {"longitude": 103.986981, "latitude": 30.757564},
				  {"longitude": 103.986941, "latitude": 30.757655},
				  {"longitude": 103.986949, "latitude": 30.757703},
				  {"longitude": 103.986941, "latitude": 30.757745},
				  {"longitude": 103.986892, "latitude": 30.7578},
				  {"longitude": 103.98682, "latitude": 30.757869},
				  {"longitude": 103.986771, "latitude": 30.757918},
				  {"longitude": 103.98669, "latitude": 30.757918},
				  {"longitude": 103.986602, "latitude": 30.757945},
				  {"longitude": 103.986529, "latitude": 30.757931},
				  {"longitude": 103.986481, "latitude": 30.757931},
				  {"longitude": 103.9864, "latitude": 30.757918},
				  {"longitude": 103.986351, "latitude": 30.757883},
				  {"longitude": 103.986303, "latitude": 30.757841},
				  {"longitude": 103.986254, "latitude": 30.7578},
				  {"longitude": 103.986206, "latitude": 30.757723},
				  {"longitude": 103.986166, "latitude": 30.757523}
				],color: "#2e7efd",
									width: 10,
									dottedLine: true,
									arrowLine: true,
									colorList: true,
									
				
										}
				  ],	markers:[
										    {"longitude": 103.986166, "latitude": 30.757523, iconPath:'../../static/images/marker.png', width: 10,
				          height: 10}  
									],
									center:[
				    103.9864,
				    30.757418
				  ],
				  post_pic:['http://120.46.81.37:8080/app/static/post_pic/3.jpg'],
				  comments: [
				  					{userpic:'http://120.46.81.37:8080/app/static/user_pic/1.jpg',username:'Yodo',data:'Running is cool!',time:'2024-4-14'},
				  				
				  					
				  				],
								
														
									
								
								},	{
									user_id:2,
									post_id:3,
									name:'Peace of summer',
									visibility:1,
									Pub_time:'2024-04-06',
									like:5,
									title: 'I like lovely swan',
									Photo:'http://120.46.81.37:8080/app/static/user_pic/3.jpg',
									post_pic:['http://120.46.81.37:8080/app/static/post_pic/1.jpg','http://120.46.81.37:8080/app/static/post_pic/2.jpg'],
									collect:8,
									content:'Swan cygnets are absolutely adorable! Their fluffy down feathers and graceful movements make them a delightful sight to behold. Watching them glide across the water with their parents is truly a heartwarming experience',
									comments_num:2,
									path:[
				  {points:[
				  {"longitude": 103.984276, "latitude": 30.768216},
				  {"longitude": 103.984082, "latitude": 30.768406},
				  {"longitude": 103.983949, "latitude": 30.76852},
				  {"longitude": 103.983872, "latitude": 30.768615},
				  {"longitude": 103.983684, "latitude": 30.768767},
				  {"longitude": 103.983529, "latitude": 30.768891},
				  {"longitude": 103.983462, "latitude": 30.768995},
				  {"longitude": 103.983308, "latitude": 30.769176},
				  {"longitude": 103.983263, "latitude": 30.769185},
				  {"longitude": 103.98312, "latitude": 30.7691},
				  {"longitude": 103.982998, "latitude": 30.768967},
				  {"longitude": 103.982832, "latitude": 30.768853},
				  {"longitude": 103.982511, "latitude": 30.768625},
				  {"longitude": 103.982368, "latitude": 30.768511},
				  {"longitude": 103.982312, "latitude": 30.768492},
				  {"longitude": 103.982279, "latitude": 30.768492},
				  {"longitude": 103.982235, "latitude": 30.768596},
				  {"longitude": 103.982191, "latitude": 30.768663},
				  {"longitude": 103.982158, "latitude": 30.768653},
				  {"longitude": 103.981947, "latitude": 30.768587},
				  {"longitude": 103.981925, "latitude": 30.76853},
				  {"longitude": 103.981815, "latitude": 30.768463},
				  {"longitude": 103.981793, "latitude": 30.768397},
				  {"longitude": 103.981815, "latitude": 30.768302}
				],color: "#2e7efd",
									width: 10,
									dottedLine: true,
									arrowLine: true,
									colorList: true,
									}
				],markers:[
										    {"longitude": 103.981815, "latitude": 30.768302, iconPath:'../../static/images/marker.png', width: 10,
				          height: 10}  
									],
									 center: [
				    103.982191,
				    30.768663
				  ],comments: [
									{userpic:'http://120.46.81.37:8080/app/static/user_pic/1.jpg',username:'Yodo',data:'hahaha',time:'2024-4-8'},
									{userpic:'http://120.46.81.37:8080/app/static/user_pic/2.jpg',username:'Vero',data:'Beautiful swan!',time:'2024-4-7'},
									
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
		url: 'http://120.46.81.37:8080/app/friend_blog',
		method: 'GET',
header: {
				    'Content-Type': 'application/x-www-form-urlencoded',
					'aToken': this.aToken,
					'rToken':this.rToken,
				},
	      success: (res) => {
	        this.posts = res.data;
	      },
	      fail: (err) => {
	        console.error('Failed to fetch posts:', err);
	      }
	    });
		
		},
		computed:{
		show_posts(){
			return this.posts.filter(x=>{
				return x.title.includes(this.search_text)||x.name.includes(this.search_text)
			})
		},
		...mapState({
				loginStatus:state=>state.loginStatus,
				aToken: state => state.aToken,
				rToken: state=>state.rToken
				
			}),
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
	    background-image: url("http://120.46.81.37:8080/app/static/images/landing.jpeg");
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
