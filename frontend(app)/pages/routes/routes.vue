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
				<view class="map-floatA">共计3段轨迹</view>
				<view class="map-floatB">2022/09/28</view>
			</map>
		</view>
 
		<!--底部盒子-->
		<view class="bottomBox">
			<view class="oneTier">
				<view class="placeBox">
					<view class="twig">
						<view class="tag">起</view>
						<view class="taginfo">江门豪爵研发中心</view>
					</view>
					<view class="twig">
						<view class="tag">终</view>
						<view class="taginfo">席帽山森林公园</view>
					</view>
				</view>
				<view class="rideBox">
					<text class="ride-mileage">骑行里程/</text>
					<view class="ride-text">
						<text class="ride-textA">25.9</text>
						<text class="ride-textB">km</text>
					</view>
				</view>
			</view>
 
			<view class="num-storey">
				<view class="ride-numbox">
					<view class="ride-num numA">01：19：46</view>
					<view class="ride-name">骑行时长</view>
				</view>
				<view class="ride-numbox">
					<view class="ride-num">
						<text class="numA">78</text>
						<text class="numB">km/h</text>
					</view>
					<view class="ride-name">最高速度</view>
				</view>
				<view class="ride-numbox">
					<view class="ride-num">
						<text class="numA">59.3</text>
						<text class="numB">km/h</text>
					</view>
					<view class="ride-name">平均速度</view>
				</view>
			</view>
 
			
		</view>
	</view>
</template>
 
<script>
	export default {
		components: {},
		data() {
			return {
				infoArr: [{
					name: "全天轨迹",
					id: -1
				}, {
					name: "轨迹1",
					id: 1
				}, {
					name: "轨迹2",
					id: 2
				}],
				rSelect: [-1], //按钮选中状态数组
				latitude: 39.909,
				longitude: 116.397,
				customCalloutMarkerIds: [0, 1, 3, 4],
				markersinit: [{
					id: 0,
					polylineId: 1,
					longitude: 116.368904,
					latitude: 39.913423,
					iconPath: "/static/img/icon/green.png",
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
					longitude: 116.398258,
					latitude: 39.904600,
					iconPath: "/static/img/icon/red.png",
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
					longitude: 116.358258,
					latitude: 39.904600,
					iconPath: "/static/img/icon/green.png",
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
					longitude: 116.398258,
					latitude: 39.924600,
					iconPath: "/static/img/icon/red.png",
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
								longitude: 116.368904,
								latitude: 39.913423
							},
							{
								longitude: 116.382122,
								latitude: 39.901176
							},
							{
								longitude: 116.387271,
								latitude: 39.912501
							},
							{
								longitude: 116.398258,
								latitude: 39.904600
							},
						],
					},
					{
						name: 'Track 2',
						arrowLine: true,
						color: '#27bd09e6',
						width: 8,
						id: 2,
						points: [{
								longitude: 116.358258,
								latitude: 39.904600
							},
 
							{
								longitude: 116.388258,
								latitude: 39.914600
							},
							{
								longitude: 116.398258,
								latitude: 39.924600
							},
						],
					},
				],
				polyline: [], //展示值
			};
		},
		// mounted(){
		// 	this.locusBtn()
		// },
		onShow() {
			this.locusBtn()
		},
		methods: {
			tapInfo(e) {
				if (this.rSelect.indexOf(e) == -1) {
					this.rSelect.splice(0, this.rSelect.length); //清空
					this.rSelect.push(e);
				} else {
					this.rSelect.splice(this.rSelect.indexOf(e), 1); //取消
				}
				this.locusBtn();
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
		color: #9b9b9b;
		border-radius: 10rpx;
		background-color: #f7f7f7;
	}
 
	.selectlocusnum {
		padding: 14rpx 20rpx;
		margin: 4rpx 12rpx;
		text-align: center;
		font-size: 28rpx;
		color: white;
		border-radius: 10rpx;
		background-color: black;
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
		width: 10%;
		font-size: 14px;
		padding: 14rpx;
		color: white;
		text-align: center;
		background-color: #999999;
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
		font-size: 14px;
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
		width: 27%;
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