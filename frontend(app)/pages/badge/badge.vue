<template>
	<view class="container">
		<!-- 用户头像和用户名 -->
		<view class="user-info">
			<image class="avatar" src="../../static/common/nothing.png" />
			
		</view>
		<!-- 徽章盒子 -->
		<view class="medal-boxes">
			<view v-for="(category, index) in categories" :key="index" class="medal-category">
				<view class="category-name">
					<text>{{ category.name }}</text>
				</view>
				<view class="medal-box">
					<view v-for="medal in category.medals" :key="medal.id" class="medal-item">
						<image class="medal-image" :src="medal.src" />
					</view>
				</view>
			</view>
		</view>
	</view>
</template>

<script>
	export default {
		data() {
			return {
				categories: [{
						name: 'Running Mileage',
						medals: [{
								id: 1,
								src: '../../static/images/badge01.jpg'
							},
							{
								id: 2,
								src: '../../static/images/badge02.jpg'
							},
							{
								id: 3,
								src: '../../static/images/badge03.jpg'
							},
							// ... more medals
						]
					},
					{
						name: 'Vehicle Mileage',
						medals: [{
								id: 4,
								src: '../../static/images/badge04.jpg'
							},
							{
								id: 5,
								src: '../../static/images/badge05.jpg'
							},
							{
								id: 9,
								src: '../../static/images/badge05.jpg'
							},
							{
								id: 10,
								src: '../../static/images/badge05.jpg'
							},
							{
								id: 11,
								src: '../../static/images/badge05.jpg'
							},
							// ... more medals
						]
					},
					{
						name: 'Participation Honor',
						medals: [{
								id: 6,
								src: '../../static/images/badge06.jpg'
							},
							{
								id: 7,
								src: '../../static/images/badge07.jpg'
							},
							{
								id: 8,
								src: '../../static/images/badge08.jpg'
							},
							// ... more medals
						]
					},
					// ... more categories
				]
			};
		}
	};
</script>

<style>
	.container {
		display: flex;
		flex-direction: column;
		align-items: center;
		/* background: linear-gradient(120deg, #a6c0fe, #f68084); */
		/* 使用多个蓝色渐变层创建背景 */
		/* background: linear-gradient(120deg, #e0f7fa, #ffffff),
			linear-gradient(120deg, #e0f7fa, #89cff0),
			linear-gradient(to right, #ffffff, #ffffff);
		background-repeat: repeat; */
		background: linear-gradient(120deg, #b3e0ef, #ffffff),
			linear-gradient(120deg, #e0f7fa, #89cff0),
			linear-gradient(to right, #ffffff, #ffffff);
		background-repeat: repeat;
		padding: 20px;
		/* 添加一些内边距 */
	}

	.user-info {
		display: flex;
		flex-direction: column;
		align-items: center;
		margin-bottom: 20px;
		/* 调整外边距 */
	}

	.avatar {
		width: 120px;
		height: 120px;
		border-radius: 50%;
		margin-bottom: 15px;
		border: 4px solid #fff;
		/* 添加白色边框以突出头像 */
		box-shadow: 0 4px 8px rgba(0, 0, 0, 0.2);
		/* 添加阴影效果 */
	}

	.username {
		font-size: 20px;
		color: #333;
		margin-top: 5px;
		/* 调整用户名与头像的间距 */
	}

	.category-name {
		margin-bottom: 5px;
		/* 根据需要调整 */
		text-align: center;
		/* 如果需要文字居中显示 */
	}

	.category-name text {
		font-weight: bold;
		/* 粗体显示类别名称 */
		font-size: 18px;
		color: #333;
		/* 文字颜色 */
	}

	.medal-boxes {
		display: flex;
		flex-direction: column;
		/* 修改为列布局 */
		align-items: center;
		gap: 20px;
		/* 徽章盒子之间的间距 */
		width: 100%;
		/* 占满宽度 */
	}

	.medal-box {
		display: flex;
		flex-direction: row;
		flex-wrap: wrap;
		/* 修改为行布局，徽章横着排列 */
		justify-content: space-around;
		/* border: 2px solid #007bff; */
		/* 添加蓝色边框 */
		padding: 10px;
		margin: auto;
		/* 内部徽章与边框的距离 */
		background: #fff;
		border-radius: 10px;
		box-shadow: 0 4px 8px rgba(0, 0, 0, 0.1);
		/* 边框圆角 */
		width: 100%;
		/* 盒子宽度 */
		max-width: 300px;
		/* 最大宽度限制 */
		overflow-x: auto;
		/* 如果徽章超出盒子宽度，启用水平滚动 */
	}

	.medal-item {
		display: flex;
		justify-content: center;
		/* 水平居中徽章 */
		margin-bottom: 10px;
		/* 徽章之间的间距 */
	}

	.medal-image {
		width: 100px;
		height: 100px;
		border-radius: 20px;
		/* 徽章的圆角 */
		/* 移除了边框，因为外层的medal-box已经有了蓝色边框 */
	}

	/* 可以为medal-box添加一些阴影，增加层次感 */
	.medal-box {
		/* ... 其他属性 ... */
		/* box-shadow: 0 2px 10px rgba(0, 123, 255, 0.2); */
		/* 蓝色阴影 */
	}
</style>