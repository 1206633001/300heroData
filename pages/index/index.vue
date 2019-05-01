<template>
	<view>
		<view class="head-img bg-img" style="background-image: url('https://ws1.sinaimg.cn/large/007l6U2xgy1g2gz8y3k4cj30u00jgn45.jpg');">

		</view>
		<view class="view-cover">
			<view class="search">
				<text class="text-zhan">战</text>
				<input class="text-df" confirm-type="go" @input="toInput" @confirm="toPage" type="text" value="" placeholder="召唤师,请输入你的名称" />
			</view>
		</view>

		<view class="history">
			<view class="title">
				最近查询
			</view>


			<view v-for="(item,index) in historyList" :key="index" class="tag-box" @tap="toPage" :data-nickname="item.nickname">
				<view class="avatar radius bg-img" :style="'background-image: url('+item.img+');'">
				</view>
				<view class="tag-box-r">
					<view class="people">
						<view class="nickname text-bold">{{item.nickname}}</view>
						<view class="cu-tag sm radius text-white" style="background-color: #c9b563;">lv {{item.lv}}</view>
					</view>
					<view class="desc">
						<view class="text-sm text-gray" style="margin-right: 10upx;">胜率：{{item.shenglv}}</view>
						<view class="text-sm text-gray">团分：{{item.tuanfen}}</view>
					</view>
				</view>
				<view class="tag-clear" @click.stop="removeHistory" :data-key="'history-'+item.nickname">
					<text class="cuIcon-close text-lg text-black"></text>
				</view>

			</view>

		</view>

		<view class="msgerror bg-red text-white" :class="{'show':isShow,'hide':!isShow}">
			<text class="cuIcon-warn text-white warn"></text>
			<view>
				查无此召唤师！
			</view>
			<text @tap="toInput" class="cuIcon-close msgclose text-black"></text>
		</view>


	</view>
</template>

<script>
	export default {
		data() {
			return {
				isShow: false,
				historyList: [],
			}
		},
		onLoad(doc) {
			if (doc.false) {
				this.isShow = true
			}
			uni.getStorageInfo({
				success: res => {
					let h = []
					for (let i = 0; i < res.keys.length; i++) {
						uni.getStorage({
							key: res.keys[i],
							success: function(d) {
								h.push(d.data)
							}
						})
					}
					this.historyList = h
				}
			})
		},
		onShow: function(){
			uni.getStorageInfo({
				success: res => {
					let h = []
					for (let i = 0; i < res.keys.length; i++) {
						uni.getStorage({
							key: res.keys[i],
							success: function(d) {
								h.push(d.data)
							}
						})
					}
					this.historyList = h
				}
			})
		},
		methods: {
			toInput: function() {
				this.isShow = false
			},
			toPage: function(e) {
				let nickname = e.detail.value || e.currentTarget.dataset.nickname
				if (nickname == "") {
					this.isShow = true
					return
				}
				uni.navigateTo({
					url: '/pages/list/list?nickname=' + nickname +'&t=1'
				})
			},
			removeHistory: function(e) {
				uni.removeStorage({
					key:e.currentTarget.dataset.key
				})
				uni.reLaunch({
					url:'/pages/index/index'
				})
			}
		}
	}
</script>

<style>
	.head-img {
		width: 100%;
		height: 485upx;
	}

	.view-cover {
		padding-left: 23upx;
		padding-right: 23upx;
		margin-top: -90upx;
	}

	.search {
		height: 90upx;
		border-radius: 45upx;
		padding-left: 45upx;
		display: flex;
		position: relative;
		border: 2upx solid #c9b563;
		justify-content: flex-start;
		align-items: center;
		background-color: white;
	}

	.search .text-zhan {
		color: #c9b563;
		font-size: 48upx;
		margin-right: 20upx;

	}

	.search input {
		color: #c9a512;
		width: 550upx;
	}

	.msgerror {
		width: 500upx;
		position: fixed;
		bottom: 100upx;
		height: 60upx;
		line-height: 60upx;
		background-color: #c9b563;
		display: flex;
		justify-content: flex-start;
		align-items: center;
		padding: 20upx;
		box-sizing: border-box;
		transition: transform .5s;

	}

	.msgerror.show {
		transform: translate(10upx);
	}

	.msgerror.hide {
		transform: translate(-510upx);
	}

	.msgerror .warn {
		margin-right: 10upx;
	}

	.msgerror .msgclose {
		margin-left: auto;
	}

	.history {
		margin-top: 30upx;
		padding: 10upx 23upx;
	}

	.history .title {
		color: gray;
	}

	.history .avatar {
		width: 78upx;
		height: 78upx;
		border: 2upx solid #c9b563;
		margin-right: 20upx;
	}

	.tag-box {
		padding: 20upx 0;
		position: relative;
		display: flex;
		justify-content: flex-start;
		align-items: center;
		border-bottom: 1upx solid gray;
	}

	.people,
	.desc {
		display: flex;
		position: relative;
		flex-direction: row;
		justify-content: flex-start;
		align-items: center;
	}

	.people .nickname {
		color: #c9b563;
		margin-right: 20upx;
	}

	.tag-clear {
		margin-left: auto;
	}
</style>
