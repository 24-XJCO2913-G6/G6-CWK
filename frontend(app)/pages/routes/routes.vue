	<template>
		<view>
			<!--顶部轨迹-->
			<scroll-view class="scroll-view_H" scroll-x="true" scroll-left="0">
				<view class="topbox">
					<view class="locusnum" :class="{'selectlocusnum' : rSelect.indexOf(value.id)!=-1}"
						v-for="(value,index) in infoArr" :key="index" @tap="tapInfo(value.id)">
						{{value.name}}
					</view>
				</view>
			</scroll-view>

			<!--轨迹地图-->
			<view class="mapbox">
				<map class="mapstyle" :latitude="latitude" :longitude="longitude" :markers="markers" :polyline="polyline">
					<cover-view slot="callout">
						<block v-for="(item, index) in customCalloutMarkerIds" :key="index">
							<cover-view class="customCallout" :marker-id="item">
								<cover-view class="content">
									{{markersinit[index].stationName}}
								</cover-view>
							</cover-view>
						</block>
					</cover-view>
					<view class="map-floatA">All 2 trajectories</view>
					<view class="map-floatB">2024/05/03</view>
				</map>
			</view>

			<!--底部盒子-->
			<view class="bottomBox">
				<view class="oneTier">
					<view class="placeBox">
						<view class="twig">
							<view class="tag">Start</view>
							<view class="taginfo">{{ startPlace }}</view>
						</view>
						<view class="twig">
							<view class="tag">End</view>
							<view class="taginfo">{{ endPlace }}</view>
						</view>
					</view>
					<view class="rideBox">
						<text class="ride-mileage">  Cycling mileage</text>
						<view class="ride-text">
							<text class="ride-textA">{{ currentMileage }}</text> <!-- 绑定当前里程数 -->
							<text class="ride-textB">km</text>
						</view>
					</view>
				</view>


				<view class="num-storey">
					<view class="ride-numbox">
						<view class="ride-num numA">{{ currentTrackStats.duration }}</view>
						<view class="ride-name">Duration of ride</view>
					</view>
					<view class="ride-numbox">
						<view class="ride-num">
							<text class="numA">{{ currentTrackStats.avgSpeed }}</text>
						</view>
						<view class="ride-name">Avg speed</view>
					</view>
				</view>

				<!-- <view v-for="track in [1, 2]" :key="track" class="download-btn">
					<button @click="downloadTrack(track)">下载轨迹{{ track }}</button>
				</view> -->

				<!-- <button v-if="currentSelectedTrackId" @click="downloadTrack">下载轨迹</button> -->
				<!-- <view class="download-btn">
					<button v-if="currentSelectedTrackId" @click="downloadTrack" style="background-color: #4ba3c7; color: white;">Download Track</button>
				</view> -->
				<view class="download-btn">
					<button v-if="currentSelectedTrackId > 0 " @click="downloadTrack"
						style="background-color: #4ba3c7; color: white;">Download Track</button>
				</view>


			</view>
		</view>
	</template>

	<script>
		import { mapState } from 'vuex'
		export default {
			components: {},
			data() {
				return {
					infoArr: [{
						name: "All tracks",
						id: -1
					}, {
						name: "Track1",
						id: 1
					}, {
						name: "Track2",
						id: 2
					}],
					// 添加起点和终点的数据属性
					startPlace: '————————',
					endPlace: '————————',
					currentMileage: 3.7, // 初始化里程，可根据实际情况设置
					currentTrackStats: {
						duration: '01:19:46',
						avgSpeed: '14.3 km/h'
					}, // 初始化为一个空对象
					// 为每个轨迹添加统计信息
					trackStats: {
						'-1': { // 今天的轨迹
							duration: '00:8:46',
							avgSpeed: '14.3 km/h'
						},
						'1': { // 轨迹1的统计信息
							duration: '00:3:40',
							avgSpeed: '15 km/h'
						},
						'2': { // 轨迹2的统计信息
							duration: '0:5:06',
							avgSpeed: '14.0 km/h'
						}
					},
					rSelect: [-1], //按钮选中状态数组
					longitude: 103.978979,
					latitude: 30.767053,
					customCalloutMarkerIds: [0, 1, 3, 4],
					markersinit: [{
						id: 0,
						polylineId: 1,
						longitude: 103.977527,
						latitude: 30.767054,
						iconPath: "http://120.46.81.37:8080/app/static/img/icon/green.png",
						width: 38,
						height: 38,
						stationName: '1',
						customCallout: {
							anchorY: 22,
							anchorX: 0,
							display: 'ALWAYS',
						}
					}, {
						id: 1,
						polylineId: 1,
						longitude: 103.977527,
						latitude: 30.767054,
						iconPath: "http://120.46.81.37:8080/app/static/img/icon/red.png",
						width: 38,
						height: 38,
						stationName: '1',
						customCallout: {
							anchorY: 22,
							anchorX: 0,
							display: 'ALWAYS',
						}
					}, {
						id: 3,
						polylineId: 2,
						longitude: 103.977527,
						latitude: 30.767054,
						iconPath: "http://120.46.81.37:8080/app/static/img/icon/green.png",
						width: 38,
						height: 38,
						stationName: '2',
						customCallout: {
							anchorY: 22,
							anchorX: 0,
							display: 'ALWAYS',
						}
					}, {
						id: 4,
						polylineId: 2,
						longitude: 103.997555,
						latitude: 30.767054,
						iconPath: "http://120.46.81.37:8080/app/static/img/icon/red.png",
						width: 38,
						height: 38,
						stationName: '终',
						customCallout: {
							anchorY: 22,
							anchorX: 0,
							display: 'ALWAYS',
						}
					}],
					markers: [],
					polylineinit: [{ //初始值
							name: 'Track 1',
							arrowLine: true,
							color: '#27bd09e6',
							width: 8,
							id: 1,
							points: [{
									"longitude": 103.977527,
									"latitude": 30.767054
								},
								{
									"longitude": 103.977643,
									"latitude": 30.766825
								},
								{
									"longitude": 103.977727,
									"latitude": 30.766768
								},
								{
									"longitude": 103.977826,
									"latitude": 30.766625
								},
								{
									"longitude": 103.977926,
									"latitude": 30.766525
								},
								{
									"longitude": 103.978076,
									"latitude": 30.766325
								},
								{
									"longitude": 103.978375,
									"latitude": 30.765926
								},
								{
									"longitude": 103.978558,
									"latitude": 30.76574
								},
								{
									"longitude": 103.978741,
									"latitude": 30.765526
								},
								{
									"longitude": 103.978874,
									"latitude": 30.765297
								},
								{
									"longitude": 103.978973,
									"latitude": 30.765269
								},
								{
									"longitude": 103.97904,
									"latitude": 30.765255
								},
								{
									"longitude": 103.979106,
									"latitude": 30.765312
								},
								{
									"longitude": 103.979273,
									"latitude": 30.765355
								},
								{
									"longitude": 103.979406,
									"latitude": 30.765483
								},
								{
									"longitude": 103.979672,
									"latitude": 30.765583
								},
								{
									"longitude": 103.979788,
									"latitude": 30.765669
								},
								{
									"longitude": 103.979938,
									"latitude": 30.765769
								},
								{
									"longitude": 103.980154,
									"latitude": 30.765869
								},
								{
									"longitude": 103.980337,
									"latitude": 30.766012
								},
								{
									"longitude": 103.980486,
									"latitude": 30.766097
								},
								{
									"longitude": 103.980619,
									"latitude": 30.766169
								},
								{
									"longitude": 103.980669,
									"latitude": 30.766169
								},
								{
									"longitude": 103.980736,
									"latitude": 30.766026
								},
								{
									"longitude": 103.980785,
									"latitude": 30.76584
								},
								{
									"longitude": 103.980835,
									"latitude": 30.765697
								},
								{
									"longitude": 103.981051,
									"latitude": 30.765612
								},
								{
									"longitude": 103.981168,
									"latitude": 30.765569
								},
								{
									"longitude": 103.981317,
									"latitude": 30.765555
								},
								{
									"longitude": 103.981517,
									"latitude": 30.765555
								},
								{
									"longitude": 103.98175,
									"latitude": 30.765497
								},
								{
									"longitude": 103.981999,
									"latitude": 30.76534
								},
								{
									"longitude": 103.982165,
									"latitude": 30.765155
								},
								{
									"longitude": 103.982282,
									"latitude": 30.764969
								},
								{
									"longitude": 103.982365,
									"latitude": 30.764812
								}
							],
						},
						{
							name: 'Track 2',
							arrowLine: true,
							color: '#27bd09e6',
							width: 8,
							id: 2,
							points: [{
									"longitude": 103.976838,
									"latitude": 30.768159
								},
								{
									"longitude": 103.976871,
									"latitude": 30.768094
								},
								{
									"longitude": 103.976905,
									"latitude": 30.768044
								},
								{
									"longitude": 103.976921,
									"latitude": 30.768023
								},
								{
									"longitude": 103.976938,
									"latitude": 30.768001
								},
								{
									"longitude": 103.97703,
									"latitude": 30.768073
								},
								{
									"longitude": 103.977097,
									"latitude": 30.768109
								},
								{
									"longitude": 103.977155,
									"latitude": 30.768166
								},
								{
									"longitude": 103.977222,
									"latitude": 30.768195
								},
								{
									"longitude": 103.977339,
									"latitude": 30.76831
								},
								{
									"longitude": 103.977439,
									"latitude": 30.768382
								},
								{
									"longitude": 103.977548,
									"latitude": 30.768439
								},
								{
									"longitude": 103.977673,
									"latitude": 30.768511
								},
								{
									"longitude": 103.977799,
									"latitude": 30.768597
								},
								{
									"longitude": 103.977983,
									"latitude": 30.768726
								},
								{
									"longitude": 103.978108,
									"latitude": 30.768827
								},
								{
									"longitude": 103.97825,
									"latitude": 30.76892
								},
								{
									"longitude": 103.9784,
									"latitude": 30.769042
								},
								{
									"longitude": 103.978526,
									"latitude": 30.769143
								},
								{
									"longitude": 103.978651,
									"latitude": 30.76925
								},
								{
									"longitude": 103.978801,
									"latitude": 30.769372
								},
								{
									"longitude": 103.978969,
									"latitude": 30.769523
								},
								{
									"longitude": 103.979127,
									"latitude": 30.769652
								},
								{
									"longitude": 103.979303,
									"latitude": 30.769439
								},
								{
									"longitude": 103.979345,
									"latitude": 30.769381
								},
								{
									"longitude": 103.979445,
									"latitude": 30.769252
								},
								{
									"longitude": 103.979545,
									"latitude": 30.769108
								},
								{
									"longitude": 103.97962,
									"latitude": 30.769008
								},
								{
									"longitude": 103.979796,
									"latitude": 30.768843
								},
								{
									"longitude": 103.97998,
									"latitude": 30.768606
								},
								{
									"longitude": 103.980113,
									"latitude": 30.768448
								},
								{
									"longitude": 103.980205,
									"latitude": 30.768276
								},
								{
									"longitude": 103.980222,
									"latitude": 30.768204
								},
								{
									"longitude": 103.980272,
									"latitude": 30.768211
								},
								{
									"longitude": 103.980356,
									"latitude": 30.768247
								},
								{
									"longitude": 103.980447,
									"latitude": 30.768326
								},
								{
									"longitude": 103.980573,
									"latitude": 30.76839
								},
								{
									"longitude": 103.980673,
									"latitude": 30.768441
								},
								{
									"longitude": 103.980715,
									"latitude": 30.768484
								},
								{
									"longitude": 103.980665,
									"latitude": 30.768577
								},
								{
									"longitude": 103.98059,
									"latitude": 30.768721
								},
								{
									"longitude": 103.980498,
									"latitude": 30.768879
								},
								{
									"longitude": 103.980297,
									"latitude": 30.76908
								},
								{
									"longitude": 103.980105,
									"latitude": 30.769302
								},
								{
									"longitude": 103.98003,
									"latitude": 30.769431
								},
								{
									"longitude": 103.979971,
									"latitude": 30.769525
								},
								{
									"longitude": 103.979946,
									"latitude": 30.769661
								},
								{
									"longitude": 103.979896,
									"latitude": 30.769733
								},
								{
									"longitude": 103.979904,
									"latitude": 30.76979
								},
								{
									"longitude": 103.980038,
									"latitude": 30.769776
								},
								{
									"longitude": 103.980172,
									"latitude": 30.76969
								},
								{
									"longitude": 103.980272,
									"latitude": 30.769546
								},
								{
									"longitude": 103.980272,
									"latitude": 30.76941
								},
								{
									"longitude": 103.980305,
									"latitude": 30.769302
								},
								{
									"longitude": 103.980322,
									"latitude": 30.769209
								},
								{
									"longitude": 103.980322,
									"latitude": 30.769159
								},
								{
									"longitude": 103.980389,
									"latitude": 30.769116
								},
								{
									"longitude": 103.980506,
									"latitude": 30.769015
								},
								{
									"longitude": 103.980581,
									"latitude": 30.768979
								},
								{
									"longitude": 103.980706,
									"latitude": 30.7689
								},
								{
									"longitude": 103.980757,
									"latitude": 30.768871
								},
								{
									"longitude": 103.980832,
									"latitude": 30.768814
								},
								{
									"longitude": 103.98094,
									"latitude": 30.768771
								},
								{
									"longitude": 103.981007,
									"latitude": 30.768713
								},
								{
									"longitude": 103.981074,
									"latitude": 30.76867
								},
								{
									"longitude": 103.981183,
									"latitude": 30.768584
								},
								{
									"longitude": 103.981258,
									"latitude": 30.76852
								},
								{
									"longitude": 103.981283,
									"latitude": 30.768469
								},
								{
									"longitude": 103.981308,
									"latitude": 30.768426
								}
							],
						},
					],
					polyline: [], //展示值
				};
			},
			// mounted(){
			// 	this.locusBtn()
			// },
			computed: {
				computedCurrentTrackStats() {
					// 确保 rSelect[0] 存在，避免未定义的错误
					return (this.rSelect[0] && this.trackStats[this.rSelect[0]]) || this.currentTrackStats;
				},
			},
			onShow() {
				this.locusBtn()
			},
			methods: {
				fetchRoutes() {
					uni.request({
						url: 'http://120.46.81.37:8080/app/myroutes',
						method: 'GET',
						data: {
							aToken: this.aToken,
							rToken: this.rToken,
						},
						header: {
							'Content-Type': 'application/json;charset=utf-8',
						},
						success: (res) => {
							if (res.statusCode === 200) {
								try {
									const responseData = JSON.parse(res.data);
									this.infoArr = responseData.infoArr || this.infoArr;
									this.startPlace = responseData.startPlace || this.startPlace;
									this.endPlace = responseData.endPlace || this.endPlace;
									this.currentMileage = responseData.currentMileage || this.currentMileage;
									this.currentTrackStats = responseData.currentTrackStats || this
										.currentTrackStats;
									this.trackStats = responseData.trackStats ? {
										...this.trackStats,
										...responseData.trackStats
									} : this.trackStats;
									this.rSelect = responseData.rSelect || this.rSelect;
									this.longitude = responseData.longitude || this.longitude;
									this.latitude = responseData.latitude || this.latitude;
									this.customCalloutMarkerIds = responseData.customCalloutMarkerIds || this
										.customCalloutMarkerIds;
									this.markersinit = responseData.markersinit || this.markersinit;
									this.markers = responseData.markers || this.markers;
									this.polylineinit = responseData.polylineinit || this.polylineinit;
									this.polyline = responseData.polyline || this.polyline;
								} catch (error) {
									console.error('解析响应数据出错', error);
								}
							} else {
								console.error('请求失败，状态码：', res.statusCode);
							}
						},
						fail: (error) => {
							console.error('网络请求失败', error);
						}
					});
				},
				// 这个方法可能需要在组件加载时调用，以获取初始数据
				loadInitialData() {
					this.fetchRoutes();
				},
				downloadTrack() {
					// 检查是否有选中的轨迹
					if (!this.currentSelectedTrackId) {
						uni.showToast({
							title: '请先选择一个轨迹',
							icon: 'none',
						});
						return;
					}

					// 根据选中的轨迹ID获取轨迹数据
					const selectedTrack = this.polylineinit.find((item) => item.id === this.currentSelectedTrackId);

					// 如果没有找到轨迹数据，提示用户并返回
					if (!selectedTrack) {
						uni.showToast({
							title: '轨迹数据不存在',
							icon: 'none',
						});
						return;
					}

					// 将轨迹坐标转换为字符串
					const coordinates = selectedTrack.points.map(point => `${point.longitude},${point.latitude}`)
						.join('\n');

					// 创建一个Blob对象，表示下载的文件
					const blob = new Blob([coordinates], {
						type: 'text/plain;charset=utf-8'
					});

					// 创建一个链接元素，用于下载
					const downloadLink = document.createElement('a');
					downloadLink.href = URL.createObjectURL(blob);
					downloadLink.download = `轨迹${this.currentSelectedTrackId}.txt`;

					// 触发下载
					downloadLink.click();

					// 释放blob对象
					URL.revokeObjectURL(downloadLink.href);

					// 提示用户下载成功
					uni.showToast({
						title: '下载成功',
						icon: 'success',
						duration: 2000
					});
				},

				tapInfo(trackId) {
					// 根据 trackId 更新 rSelect 数组
					if (this.rSelect[0] === trackId) {
						// 如果点击的是当前已选中的轨迹，则取消选中
						this.rSelect.splice(0, 1);
					} else {
						// 否则，清空 rSelect 并添加新的轨迹 ID
						this.rSelect.splice(0, this.rSelect.length, trackId);
					}

					// 更新当前选中的轨迹ID
					this.currentSelectedTrackId = this.rSelect[0];

					// 调用 locusBtn 方法更新轨迹显示
					this.locusBtn();

					// 根据 trackId 更新页面上显示的数据
					if (trackId === -1) {
						// 今天的轨迹或所有轨迹
						this.startPlace = '————————';
						this.endPlace = '————————';
						this.currentMileage = 3.7;
					} else {
						// 根据实际轨迹ID设置起始点和终点
						// 这里需要根据实际数据设置
					}

					// 如果选择了 "All tracks"，则隐藏下载按钮
					if (trackId === -1) {
						this.currentSelectedTrackId = null;
					}
				},
				computed: {
					currentTrackStats() {
						// 返回当前选中轨迹的统计信息
						return this.trackStats[this.rSelect[0]];
					}
				},
				mounted() {
					this.loadInitialData();
				},

				locusBtn() {
					if (this.rSelect.length > 0) {
						if (this.rSelect[0] === -1) {
							//全天轨迹
							this.polyline = JSON.parse(JSON.stringify(this.polylineinit));
							this.markers = JSON.parse(JSON.stringify(this.markersinit));
						} else {
							//非全天轨迹
							let arr = JSON.parse(JSON.stringify(this.polylineinit));
							let obj = arr.find(item => {
								return item.id == this.rSelect[0]
							});
							this.polyline = [obj];

							let markerArr = JSON.parse(JSON.stringify(this.markersinit));
							this.markers.splice(0, this.markers.length); //清空
							for (var i = 0; i < markerArr.length; i++) {
								if (markerArr[i].polylineId === this.rSelect[0]) {
									this.markers.push(markerArr[i])
								}
							}
						}
					}
				}
			},
		};
	</script>

	<style lang="scss">
		page {
			background-color: white;
		}

		.download-btn {
			margin-top: 7px;
			font-weight: bold;

		}

		.scroll-view_H {
			white-space: nowrap;
			width: 100%;
		}

		.topbox {
			margin: 0px 0px 0px 6rpx;
			display: flex;
			align-items: center;
		}

		.locusnum {
			padding: 14rpx 20rpx;
			margin: 4rpx 12rpx;
			text-align: center;
			font-size: 28rpx;
			color: white;
			border-radius: 10rpx;
			background-color: darkgray;
		}

		.selectlocusnum {
			padding: 14rpx 20rpx;
			margin: 4rpx 12rpx;
			text-align: center;
			font-size: 28rpx;
			color: white;
			border-radius: 10rpx;
			background-color: #4ba3c7;
		}

		.mapbox {
			margin: 10rpx 20rpx;
		}

		.mapstyle {
			width: 100%;
			height: 360px;
			position: relative;
		}

		.map-floatA {
			width: 22%;
			position: absolute;
			bottom: 110rpx;
			left: 20rpx;
			background-color: #292929a6;
			color: white;
			font-size: 24rpx;
			padding: 10rpx 30rpx;
			border-radius: 10rpx;
			letter-spacing: 2rpx;
			text-align: center;
		}

		.map-floatB {
			width: 22%;
			position: absolute;
			bottom: 40rpx;
			left: 20rpx;
			background-color: #292929a6;
			color: white;
			font-size: 24rpx;
			padding: 10rpx 30rpx;
			border-radius: 10rpx;
			letter-spacing: 2rpx;
			text-align: center;
		}


		.bottomBox {
			margin: 10rpx 20rpx;
		}

		.oneTier {
			display: flex;
			justify-content: space-between;
		}

		.twig {
			width: 96%;
			// margin-bottom: 20rpx;
			display: flex;
			justify-content: flex-start;
		}

		.tag {
			width: 50%;
			font-size: 14px;
			padding: 14rpx;
			color: white;
			text-align: center;
			background-color: #4ba3c7;
			border-top-left-radius: 10rpx;
			border-bottom-left-radius: 10rpx;
		}

		.taginfo {
			width: 90%;
			font-size: 14px;
			padding: 14rpx 20rpx;
			text-align: left;
			color: #767676;
			background-color: #f7f7f7;
			border-top-right-radius: 10rpx;
			border-bottom-right-radius: 10rpx;
		}

		.placeBox {
			width: 66%;
			height: 140rpx;
			display: flex;
			flex-direction: column;
			justify-content: space-between;
		}

		.rideBox {
			width: 34%;
			height: 134rpx;
			padding: 5rpx 20rpx;
			background-color: #f7f7f7;
			border-radius: 10rpx;
		}

		.ride-mileage {
			font-size: 15px;
			color: #959595;
		}

		.ride-text {
			margin-top: 10rpx;
		}

		.ride-textA {
			font-size: 60rpx;
			font-weight: bold;
			font-style: italic;
		}

		.ride-textB {
			margin-left: 20rpx;
			color: #949494;
			font-style: italic;
		}

		.num-storey {
			display: flex;
			justify-content: space-between;
			align-items: center;
		}

		.ride-numbox {
			width: 43%;
			height: 120rpx;
			margin-top: 14rpx;
			padding: 5rpx 20rpx;
			background-color: #f7f7f7;
			border-radius: 10rpx;
			display: flex;
			flex-direction: column;
			justify-content: center;
			align-items: center;
		}

		.ride-num {
			margin-bottom: 10rpx;
		}

		.numA {
			font-size: 32rpx;
			font-weight: bold;
			font-style: italic;
		}

		.numB {
			font-size: 24rpx;
			margin-left: 20rpx;
			color: #949494;
			font-style: italic;
		}

		.ride-name {
			font-size: 14px;
			color: #959595;
		}

		.customCallout {
			padding: 2rpx 8rpx;
			text-align: center;
			color: #2A7BE2;
			font-size: 26rpx;
			font-weight: bold;
			background: #ffffffb8;
			border-radius: 30rpx;
		}
	</style>
