<template>
	<view>
		<cu-custom :isBack="true">
			<block slot="content">{{nickname}}</block>
		</cu-custom>
		<view class="loading ffff" :class="{'hide-loading':show != 1,'show-loading':show==1}">
			<view class="loading-img bg-img" style="background-image: url('https://ws1.sinaimg.cn/large/007l6U2xgy1g2h0jgy1ofg302s04875o.jpg');"></view>
			<view class="loading-title">
				鱼鱼鱼,正在打扫战斗数据...
			</view>
		</view>


		<view class="page" :class="{'hide-page':show != 2,'show-page':show == 2}">
			<view class="page-header bg-img" :style="'background-image:url('+headimg+');'">
				<view class="page-header-cover" style="width: 100%;background-color: rgba(0,0,0,.8);">

					<view class="avatar radius bg-img" :style="'background-image: url('+headimg+');'">
					</view>
					<view class="page-header-r">
						<view class="people">
							<view class="nickname">{{nickname}}</view>
							<view class="lv cu-tag radius line-black text-white" style="margin-left: 10upx; border: 2upx solid #C9A512;">lv
								{{rolelv}}</view>
						</view>
						<view class="other-msg" style="margin-top: 20upx;">
							<view class="grade-me text-gray text-sm" style="margin-right: 10upx;">实力：{{gerenshili}}</view>
							<view class="grade-me text-gray text-sm" style="margin-right: 10upx;">团分：{{tuanfen}}</view>
							<view class="grade-group text-gray  text-sm">胜率：{{winCount}}/{{matchCount}}={{shenglv}}%</view>
						</view>
					</view>
				</view>

			</view>

			<view v-if="relimg.length > 0" class="march-title bg-white">
				<view class="title-left">
				</view>
				<view class="title">
					最近常用
				</view>
			</view>
			<view v-if="relimg.length > 0" class="recent">
				<view v-for="(item,index) in relimg" :key="index" class="avatar">
					<view class="avatar-img  bg-img" :style="'background-image: url('+item+');'">
					</view>
				</view>
			</view>






			<view class="march-title bg-white">
				<view class="title-left">
				</view>
				<view class="title">
					战斗数据
				</view>
				<view @tap="toRe" class="cuIcon-refresh text-lg text-gray" style="margin-left: auto;">
				</view>
			</view>

			<view v-for="(item,index) in matchList" :key="index" @tap="toMatch" :data-key="index" :data-match="item.MatchID" class="march-list">
				<view class="avatar">
					<view class="avatar-img  bg-img" :style="'background-image: url(http://300report.jumpw.com/static/images/'+item.Hero.IconFile+');'">
					</view>
				</view>
				<view class="mid">
					<view class="mid-h">
						<view class="result" :class="{'result-false':item.Result == 2,'result-success':item.Result == 1,'result-tp':item.Result == 3}">
							{{[item.Result == 1?'胜利':item.Result == 2?'失败':'逃跑']}}
						</view>
						<view class="result-xx">
							{{[item.killCount + ' / ' + item.deathCount + ' / '+item.assistCount]}}
						</view>
					</view>
					<view class="type text-gray">
						{{[item.MatchType == 1?'竞技场':'氪金人的战场']}}
					</view>
				</view>
				<view class="right">
					<view class="time text-bold text-sm">{{item.MatchDate}}</view>
					<view v-if="item.isMvp" class="mvp bg-img" style="background-image: url('https://ws1.sinaimg.cn/large/007l6U2xgy1g2i9rh0d1yj302c01g743.jpg');"></view>
				</view>
				<view class="bg-img" style="width: 30upx;height: 30upx;background-image: url('https://wx3.sinaimg.cn/square/007l6U2xgy1g2h0cwyup8j3023023we9.jpg');transform: rotate(180deg);"></view>

			</view>

			<view class="text-gray text-sm" style="height: 150upx;text-align: center;line-height: 150upx;margin-bottom: 90upx;">
				{{desc}}
			</view>
		</view>


	</view>
</template>

<script>
	export default {
		data() {
			return {
				nickname: '',
				show: 1, // 1 = loading  2 = OKpage  3 = falsePage
				rolelv: 0,
				roleID: 0,
				winCount: 0,
				matchCount: 0,
				shenglv: 0,
				headimg: '',
				gerenshili: 0,
				tuanfen: 0,

				relimg: [],

				matchList: [],
				pageIndex: 0,

				desc: '下拉加载更多...',

				noReach: true,

			}
		},
		onLoad(doc) {
			this.show = 1
			let requestPromise = myUrl => {
				return new Promise((resolve, reject) => {
					uni.request({
						url: myUrl,
						success: res => resolve(res),
						fail: () => reject()
					})
				})
			}

			requestPromise('https://300hero.5idown.cn/index.php?/api/getrole?name=' + doc.nickname).then(
				res => {
					if (res.data.Result != 'OK') {
						uni.reLaunch({
							url: '/pages/index/index?false=true'
						})
						return
					}
					// 设置一些页面全局信息
					this.nickname = res.data.Role.RoleName
					this.rolelv = res.data.Role.RoleLevel
					this.winCount = res.data.Role.WinCount
					this.matchCount = res.data.Role.MatchCount
					this.shenglv = parseInt((res.data.Role.WinCount / res.data.Role.MatchCount) * 100)
					this.roleID = res.data.Role.RoleID

					if (res.data.Rank == null) {

					} else {
						for (let i = 0; res.data.Rank.length != undefined && i < res.data.Rank.length; i++) {
							if (res.data.Rank[i]['ValueName'] == '场数') {
								let img = res.data.Rank[i]['Type'] - 15 < 100 ? '00' + (res.data.Rank[i]['Type'] - 15) : '0' + (res.data.Rank[
									i]['Type'] - 15)
								this.relimg.push('http://300report.jumpw.com/static/images/herohead/chara_' + img + '.png')
							}
							if (res.data.Rank[i]['Type'] == 4) {
								this.gerenshili = res.data.Rank[i]['Value']
							}
							if (res.data.Rank[i]['Type'] == 1) {
								this.tuanfen = res.data.Rank[i]['Value']
							}
						}

					}
					this.headimg = this.relimg.length > 0 ? this.relimg[0] :
						'http://300report.jumpw.com/static/images/herohead/chara_0053.png'

					if (doc.t == 1) {
						uni.setStorage({
							key: 'history-' + this.nickname,
							data: {
								lv: this.rolelv,
								nickname: this.nickname,
								shenglv: parseInt((this.winCount / this.matchCount) * 100) + '%',
								tuanfen: this.tuanfen,
								img: this.relimg.length == 0 ? 'http://300report.jumpw.com/static/images/herohead/chara_0053.png' : this.relimg[
									0]
							},
						})
					}

				}
			)


			// 获取列表，并给mList 然后 内部操作
			let shouldNum = 0
			let okNum = 0
			requestPromise('https://300hero.5idown.cn/index.php?/api/getlist?name=' + doc.nickname + '&index=' + this.pageIndex).then(
				res => {
					if (res.data.Result == 'OK') {
						// 将获取到的数据先存在mList中
						let mList = res.data.List
						shouldNum = mList.length
						return mList
					}
				}).catch(() => {
				console.log("数据加载失败")
			}).then(mList => {
				let per
				for (let i = 0; i < mList.length; i++) {
					// 获取详细信息并绑定。
					requestPromise('https://300hero.5idown.cn/index.php?/api/getmatch?id=' + mList[i].MatchID).catch(() => {
						console.log("超时fanhui1")
					}).then(
						matchRes => {
							// 信息确定获取到了
							let isFind = false
							let maxNum = 0
							let myNum = 0
							for (let j = 0; j < matchRes.data.Match.WinSide.length; j++) {
								if (maxNum < matchRes.data.Match.WinSide[j].KDA) {
									maxNum = matchRes.data.Match.WinSide[j].KDA
								}
								if (isFind) {
									continue
								} else {
									if (matchRes.data.Match.WinSide[j].HeroID == mList[i].Hero.ID) {
										isFind = true
										mList[i].killCount = matchRes.data.Match.WinSide[j].KillCount
										mList[i].deathCount = matchRes.data.Match.WinSide[j].DeathCount
										mList[i].assistCount = matchRes.data.Match.WinSide[j].AssistCount
										myNum = matchRes.data.Match.WinSide[j].KDA
									}
								}
							}
							for (let j = 0; j < matchRes.data.Match.LoseSide.length; j++) {
								if (maxNum < matchRes.data.Match.LoseSide[j].KDA) {
									maxNum = matchRes.data.Match.LoseSide[j].KDA
								}
								if (isFind) {
									continue
								} else {
									if (matchRes.data.Match.LoseSide[j].HeroID == mList[i].Hero.ID) {
										isFind = true
										mList[i].killCount = matchRes.data.Match.LoseSide[j].KillCount
										mList[i].deathCount = matchRes.data.Match.LoseSide[j].DeathCount
										mList[i].assistCount = matchRes.data.Match.LoseSide[j].AssistCount
										myNum = matchRes.data.Match.LoseSide[j].KDA
									}
								}
							}
							mList[i].MatchDate = mList[i].MatchDate.substr(5, 11)
							if (myNum >= maxNum) {
								mList[i].isMvp = true
							}
							okNum += 1
							if (okNum == shouldNum) {
								return {
									ok: true,
									data: mList
								}
							}
							return {
								ok: false
							}
						}
					).then(t => {
						if (t.ok) {
							if (shouldNum < 10) {
								this.desc = "在拉也没有了。思米马赛！"
								uni.stopPullDownRefresh()
								this.noReach = false
							}
							this.pageIndex++
							this.matchList = t.data
							this.show = 2
						}
					})
				}
			})

		},
		onReady: function() {

		},
		onReachBottom: function() {
			if (!this.noReach) {
				return
			}

			this.show = 1
			let requestPromise = myUrl => {
				return new Promise((resolve, reject) => {
					uni.request({
						url: myUrl,
						success: res => resolve(res),
						fail: () => reject()

					})
				})
			}
			let shouldNum = 0
			let okNum = 0

			requestPromise('https://300hero.5idown.cn/index.php?/api/getlist?name=' + this.nickname + '&index=' + (this.pageIndex * 10)).then(
				res => {
					if (res.data.Result == 'OK') {
						// 将获取到的数据先存在mList中
						let mList = res.data.List
						shouldNum = mList.length
						return mList
					}
				}).then(mList => {
				let per
				for (let i = 0; i < mList.length; i++) {
					// 获取详细信息并绑定。
					requestPromise('https://300hero.5idown.cn/index.php?/api/getmatch?id=' + mList[i].MatchID).catch(() => {
					
					}).then(
						matchRes => {
							// 信息确定获取到了
							let isFind = false
							let maxNum = 0
							let myNum = 0
							for (let j = 0; j < matchRes.data.Match.WinSide.length; j++) {
								if (maxNum < matchRes.data.Match.WinSide[j].KDA) {
									maxNum = matchRes.data.Match.WinSide[j].KDA
								}
								if (isFind) {
									continue
								} else {
									if (matchRes.data.Match.WinSide[j].HeroID == mList[i].Hero.ID) {
										isFind = true
					
										mList[i].killCount = matchRes.data.Match.WinSide[j].KillCount
										mList[i].deathCount = matchRes.data.Match.WinSide[j].DeathCount
										mList[i].assistCount = matchRes.data.Match.WinSide[j].AssistCount
										myNum = matchRes.data.Match.WinSide[j].KDA
									}
								}
							}
							for (let j = 0; j < matchRes.data.Match.LoseSide.length; j++) {
								if (maxNum < matchRes.data.Match.LoseSide[j].KDA) {
									maxNum = matchRes.data.Match.LoseSide[j].KDA
								}
								if (isFind) {
									continue
								} else {
									if (matchRes.data.Match.LoseSide[j].HeroID == mList[i].Hero.ID) {
										isFind = true
							
										mList[i].killCount = matchRes.data.Match.LoseSide[j].KillCount
										mList[i].deathCount = matchRes.data.Match.LoseSide[j].DeathCount
										mList[i].assistCount = matchRes.data.Match.LoseSide[j].AssistCount
										myNum = matchRes.data.Match.LoseSide[j].KDA
									}
								}
							}
							mList[i].MatchDate = mList[i].MatchDate.substr(5, 11)
							if (myNum >= maxNum) {
								mList[i].isMvp = true
							}
							okNum += 1
							if (okNum == shouldNum) {
								console.log(mList)
								return {
									ok: true,
									data: mList
								}
							}
							return {
								ok: false
							}
						}
					).then(t => {
						if (t.ok) {
							if (shouldNum < 10) {
								this.desc = "在拉也没有了。思米马赛！"
								uni.stopPullDownRefresh()
								this.noReach = false
							}

							this.pageIndex++
							this.matchList = this.matchList.concat(t.data)
							this.show = 2

						}
					})
				}
			})

		},
		methods: {
			toRe: function() {
				uni.redirectTo({
					url: '/pages/list/list?nickname=' + this.nickname
				})
			},
			toMatch: function(e){
				// let matchD = this.matchList[e.currentTarget.dataset.key].MatchData
				
				
				
				uni.navigateTo({
					url:'/pages/match/match?matchID='+e.currentTarget.dataset.match+'&roleid='+this.roleID
				})
			}
		}
	}
</script>

<style>
	.ffff {
		position: fixed;
		width: 100%;

	}

	.march-list,
	.recent {
		position: relative;
		display: flex;
		justify-content: flex-start;
		align-items: center;
		padding: 20upx 23upx;
		background-color: white;
		border-bottom: 2upx solid #F4F4F4;
		background-color: #FCFCFC;
	}

	.march-list .avatar,
	.recent .avatar {
		border: 2upx solid #C9A512;
		padding: 2upx;
		border-radius: 50%;
		margin-right: 20upx;
	}

	.march-list .avatar .avatar-img,
	.recent .avatar .avatar-img {
		width: 74upx;
		height: 74upx;
		border-radius: 50%;
	}

	.recent .avatar {
		margin-right: 20upx;
	}


	.mid-h {
		position: relative;
		display: flex;
		justify-content: flex-start;
		flex-direction: row;
		align-items: center;
	}

	.mid-h .result {
		font-size: 30upx;
		margin-right: 10upx;
	}

	.mid-h .result.result-false {
		color: #ff636d;
	}

	.mid-h .result.result-success {
		color: #00b1c2;
	}

	.mid-h .result.result-tp {
		color: #f69c00;
	}

	.mid-h .result-xx {
		color: #4E6E73;
		font-style: italic;
		font-weight: bold;
	}

	.mid .type {
		font-size: 28upx;
		color: #9BADB0;
	}

	.right {
		margin-left: auto;
		margin-right: 10upx;
	}

	.right .mvp {
		width: 63upx;
		height: 39upx;
		margin-left: auto;
	}

	.time {
		color: #77888b;
	}

	.more {
		color: #77888b;
	}


	.march-title {
		width: 100%;
		padding: 20upx;
		border-bottom: 2upx solid #F5F5F5;
		display: flex;
		justify-content: flex-start;
		align-items: center;
	}

	.march-title .title-left {
		width: 10upx;
		height: 20upx;
		border-radius: 5upx;
		background-color: #c9b563;
		margin-right: 20upx;
	}

	.march-title .title {
		color: #c9b563;
	}

	.page-header {}

	.page-header-cover {
		width: 100%;
		padding: 50upx 23upx;
		background-color: #333333;
		position: relative;
		display: flex;
		flex-direction: row;
		justify-content: flex-start;
		align-items: center;
	}

	.page-header-cover .avatar {
		width: 138upx;
		height: 138upx;
		margin-right: 20upx;

	}

	.page-header-r .people,
	.page-header-r .other-msg {
		position: relative;
		display: flex;
		justify-content: flex-start;
		align-items: center;
	}

	.page-header-r .people .nickname {
		color: #fcfcfc;
		font-size: 32upx;
	}


	.loading {
		padding-top: 50upx;
		transform: translateY(-250upx);
		transition: transform .3s;
		height: 250upx;
	}

	.loading-img {
		width: 100upx;
		height: 152upx;
		margin: 0 auto;
	}

	.loading-title {
		color: #c9b563;
		text-align: center;
	}

	.show-loading {
		transform: translate(0upx);
	}

	.hide-loading {
		transform: translateY(-250upx);
	}

	.page {
		position: relative;
		top: 0upx;
		opacity: 0;
		transition: opacity .3s;
	}

	.hide-page {
		opacity: 0;
	}

	.show-page {
		opacity: 1;
	}
</style>
