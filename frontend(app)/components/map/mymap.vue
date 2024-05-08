<template>
	<view class="map" :style="{height:mapheight+'vh'}">
	
	<view :id="myid(post_id)" class="amap" style='width: 100%;
			height: 100%;'></view>
	</view>

</template>


<script module="ModuleInstance" lang="renderjs">
	export default {
		props: ['path', 'center', 'zoom', 'mapheight','post_id'],
		data() {
			return {
				map: null,
				layer: null,
				parkList: [],
				markerId: '',
				longitude: 0,
				latitude: 0,
				markers: [ //标记点用于在地图上显示标记的位置
					{
						id: 1,
						latitude: "", //纬度
						longitude: "", //经度
						iconPath: "/static/images/business_map/my.png", //图标路径
						width: "30", //图标的宽
						height: "30" //图标的高
					}
				]

			};
		},
		mounted() {
			if (window.AMap) {
				this.initAmap();
			} else {
				const script = document.createElement('script');
				script.src = "https://webapi.amap.com/maps?v=1.4.15&key=88a2ba19bc223378211e8c2f00a25622";
				script.onload = () => {
					this.initAmap();
				}
				document.head.appendChild(script);
			}



		},
		created() {

		},
		methods: {
			myid(id){
				return "mymap_"+id
			},
			initAmap() {
				this.map = new AMap.Map(`mymap_${this.post_id}`, {
					zoom: this.zoom,
					center: this.center,
					mapStyle: 'amap://styles/57994c871bb604a4c79184f5f65d8782',
					resizeEnable: true
				});
				this.map.on("complete", () => {
					this.createLabelsLayer();
				});
				// 创建路径的坐标数组
				// 创建 Polyline 对象并添加到地图上
				var polyline = new AMap.Polyline({
					path: this.path, // 设置路径数组
					strokeColor: "#416dff", // 设置线颜色为红色
					strokeOpacity: 0.8, // 设置线透明度为 0.8
					strokeWeight: 5, // 设置线宽为 5
					strokeStyle: "solid", // 设置线样式为虚线
					strokeDasharray: [10, 5], // 设置虚线样式，间隔为 10 像素，实线长度为 5 像素
					showDir: true
				});
				var marker = new AMap.Marker({
					position: this.path[this.path.length - 1], // 终点坐标
					map: this.map, // 所属的地图实例
					icon: new AMap.Icon({ // 使用自定义图标
						size: new AMap.Size(24, 30), // 图标尺寸
						image: '//webapi.amap.com/theme/v1.3/markers/n/mark_r.png', // 图标地址
						imageSize: new AMap.Size(24, 30) // 图标尺寸
					})
				});
				polyline.setMap(this.map); // 将 Polyline 添加到地图上


			},


			createLabelsLayer() {
				if (!this.map) return;
				let that = this;
				AMap.plugin('AMap.Geolocation', function() {
					var geolocation = new AMap.Geolocation({
						enableHighAccuracy: true,
						timeout: 10000,
						offset: [10, 20],
						zoomToAccuracy: true,
						position: 'RB'
					})
					geolocation.getCurrentPosition(function(status, result) {

						if (status == 'complete') {
							onComplete(result)
						} else {
							onError(result)
						}
					});

					function onComplete(data) {
						let marker = new AMap.Marker({
							position: new AMap.LngLat(data.position.lng, data.position.lat),
						});
						that.map.setCenter(new AMap.LngLat(data.position.lng, data.position.lat));
						that.map.setZoom(18);
						that.map.add(marker);
					}

					function onError(data) {
						console.log("定位出错", data);
					}
				})
			}
		},
	};
</script>

<style lang="scss">
	page {
		width: 100%;
		height: 80%;
	}

	.map {
		width: 100%;


	}

	#bmap {
		width: 100%;
		height: 100%;
	}
</style>