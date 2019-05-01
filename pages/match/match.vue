<template>
	<view>
		<cu-custom :isBack="true">
			<block slot="content">战斗详情</block>
		</cu-custom>
		<view class="loading ffff" :class="{'hide-loading':show != 1,'show-loading':show==1}">
			<view class="loading-img bg-img" style="background-image: url('https://ws1.sinaimg.cn/large/007l6U2xgy1g2h0jgy1ofg302s04875o.jpg');"></view>
			<view class="loading-title">
				鱼鱼鱼,正在打扫战斗数据...
			</view>
		</view>
		<view v-if="show==0" class="match-state-box bg-img" :class="{'false':match.myteamState == 2,'success':match.myteamState == 1}">
			<view class="time">
				{{match.MatchDate}}
			</view>
			<view class="slip">
				|
			</view>
			<view class="hs">
				{{match.UsedTime}}
			</view>
			<view class="slip">
				|
			</view>
			<view class="type">
				{{[match.MatchType == 1?'竞技场':'战场']}}
			</view>
		</view>
		<view v-if="show==0" class="match-box bg-img" :class="{'false':match.myteamState == 2,'success':match.myteamState == 1}">
			<view class="mymatch">
				<view>{{[match.myteamState == 1 ? match.winTa : match.lowTa]}}</view>
				<view>{{[match.myteamState == 1 ? match.winXiaobing : match.lowXiaobing]}}</view>
				<view>{{[match.myteamState == 1 ? match.winZhu : match.lowZhu]}}</view>
				<view>{{[match.myteamState == 1 ? match.winShadi : match.lowShadi]}}</view>
			</view>
			<view class="othermatch">
				<view>{{[match.myteamState == 2 ? match.winShadi : match.lowShadi]}}</view>
				<view>{{[match.myteamState == 2 ? match.winZhu : match.lowZhu]}}</view>
				<view>{{[match.myteamState == 2 ? match.winXiaobing : match.lowXiaobing]}}</view>
				<view>{{[match.myteamState == 2 ? match.winTa : match.lowTa]}}</view>
			</view>
		</view>

		<view v-if="show==0" class="tt-state bg-img" :class="{'false':match.myteamState == 2,'success':match.myteamState == 1}">
			<view class="team-state text-white">
				我方{{[match.myteamState == 1?'胜':'败']}}
			</view>
			<view class="qian">{{[match.myteamState == 1?match.winJingji:match.lowJingji]}}</view>
			<view class="sha">
				{{[match.myteamState == 1? match.winShadi+'/'+match.winDeath+'/'+match.winZhu:match.lowShadi+'/'+match.lowDeath+'/'+match.lowZhu]}}
			</view>
		</view>

		<block v-if="match.myteamState == 1 && show==0">


			<view v-for="(item,i) in match.WinSide" :key="i" data-type="win" :data-key="i" class="hero-item" hover-class="hero-item-hover"
			 @tap="toMasg">
				<view class="msg">
					<view class="avatar-box" :data-nickname="item.RoleName" @tap.stop="toList">
						<view class="avatar bg-img round" :style="'background-image: url('+item.HeroIMG+');'">
						</view>
						<view v-if="item.isMe" class="tag round bg-img"></view>
						<view class="lv round">{{item.RoleLevel}}</view>
					</view>
					<view class="mid">
						<view class="nickname" :data-nickname="item.RoleName" @tap.stop="toList" :style="item.isMe?'color:#dbc56c;':''">
							{{item.RoleName}}
						</view>
						<view class="zhuangbei-list">
							<view v-for="(sk,index) in item.Skill" :key="index" class="zhuangbei-item radius bg-img" :style="'background-image: url(http://300report.jumpw.com/static/images/'+sk.IconFile+');'">
							</view>
							<view class="block" style="width: 6upx;height: 36upx;border-radius: 3upx;background-color: gray;margin: 0 5upx;"></view>

							<view v-for="(sk,index) in item.Equip" :key="index" class="zhuangbei-item radius bg-img" :style="'background-image: url(http://300report.jumpw.com/static/images/'+sk.IconFile+');'">
							</view>
						</view>
					</view>
					<view class="qian">
						<view class="bg-img" style="background-image: url('https://ws1.sinaimg.cn/large/007l6U2xgy1g2ifgccuccj300u00u0rt.jpg');"></view>
						{{item.TotalMoney}}
					</view>
					<view class="sha">
						<view class="sha-sha">
							<view class="bg-img" style="background-image: url('https://ws1.sinaimg.cn/large/007l6U2xgy1g2iff1bgamj300u00u741.jpg');"></view>
							{{item.KillCount}}/{{item.DeathCount}}/{{item.AssistCount}}
						</view>
						<view v-if="item.isM" class="bg-img mvp"></view>
					</view>
				</view>
				<view class="ttag">
					<view v-if="item.isTao" class="tt bg-img" style="background-image: url('https://ws1.sinaimg.cn/large/007l6U2xgy1g2jv7q6ibfj300q00q0sh.jpg');"></view>
					<!-- 杀 -->
					<view v-if="item.isShadi" class="tt bg-img" style="background-image: url('https://ws1.sinaimg.cn/large/007l6U2xgy1g2jv6c0cixj300q00q0q4.jpg');"></view>
					<view v-if="item.isZhu" class="tt bg-img" style="background-image: url('https://ws1.sinaimg.cn/large/007l6U2xgy1g2jv77etd4j300q00q0q1.jpg');"></view>
					<view v-if="item.isXiaobing" class="tt bg-img" style="background-image: url('https://ws1.sinaimg.cn/large/007l6U2xgy1g2jv86ombmj300q00q0o9.jpg');"></view>
					<view v-if="item.isT" class="tt bg-img" style="background-image: url('https://ws1.sinaimg.cn/large/007l6U2xgy1g2jv8q4jfbj300q00q0et.jpg');"></view>
					<view v-if="item.isJingji" class="tt bg-img" style="background-image: url('https://ws1.sinaimg.cn/large/007l6U2xgy1g2jv96iizlj300q00q0od.jpg');"></view>
					<view class="more-msg cuIcon-unfold text-gray text-lg"></view>
				</view>
				<view class="more-msg-box" v-if="winShowIndex == i">
					<view class="msg-b" style="width: 230upx;">杀了多少兵：{{item.KillUnitCount}}</view>
					<view class="msg-b" style="width: 230upx;">拆了多少塔：{{item.TowerDestroy}}</view>
					<view class="msg-b" style="width: 230upx;">得了多少经验：{{item.RewardExp}}</view>
					<view class="msg-b" style="width: 230upx;">赢了多少把：{{item.WinCount}}/{{item.MatchCount}}</view>
					<view class="msg-b" style="width: 230upx;">第一次赢吗：{{[item.IsFirstWin == 1?'首胜':'不是']}}</view>
					<view class="msg-b" style="width: 230upx;">拿到多少金币：{{item.RewardMoney}}</view>
					<view class="msg-b" style="width: 690upx;">你的团队实力（团分）：{{item.ELO}}</view>
					<view class="msg-b" style="width: 690upx;">你在本场中的表现：{{item.KDA}}</view>
				</view>
			</view>
		</block>
		<block v-else-if="show==0">
			<view v-for="(item,i) in match.LoseSide" :key="i" data-type="low" :data-key="i" class="hero-item" hover-class="hero-item-hover"
			 @tap="toMasg">
				<view class="msg">
					<view class="avatar-box" :data-nickname="item.RoleName" @tap.stop="toList">
						<view class="avatar bg-img round" :style="'background-image: url('+item.HeroIMG+');'">
						</view>
						<view v-if="item.isMe" class="tag round bg-img"></view>
						<view class="lv round">{{item.RoleLevel}}</view>
					</view>
					<view class="mid">
						<view class="nickname" :data-nickname="item.RoleName" @tap.stop="toList" :style="item.isMe?'color:#dbc56c;':''">
							{{item.RoleName}}
						</view>
						<view class="zhuangbei-list">
							<view v-for="(sk,index) in item.Skill" :key="index" class="zhuangbei-item radius bg-img" :style="'background-image: url(http://300report.jumpw.com/static/images/'+sk.IconFile+');'">
							</view>
							<view class="block" style="width: 6upx;height: 36upx;border-radius: 3upx;background-color: gray;margin: 0 5upx;"></view>

							<view v-for="(sk,index) in item.Equip" :key="index" class="zhuangbei-item radius bg-img" :style="'background-image: url(http://300report.jumpw.com/static/images/'+sk.IconFile+');'">
							</view>
						</view>
					</view>
					<view class="qian">
						<view class="bg-img" style="background-image: url('https://ws1.sinaimg.cn/large/007l6U2xgy1g2ifgccuccj300u00u0rt.jpg');"></view>
						{{item.TotalMoney}}
					</view>
					<view class="sha">
						<view class="sha-sha">
							<view class="bg-img" style="background-image: url('https://ws1.sinaimg.cn/large/007l6U2xgy1g2iff1bgamj300u00u741.jpg');"></view>
							{{item.KillCount}}/{{item.DeathCount}}/{{item.AssistCount}}
						</view>
						<view v-if="item.isM" class="bg-img mvp"></view>
					</view>
				</view>
				<view class="ttag">
					<view v-if="item.isTao" class="tt bg-img" style="background-image: url('https://ws1.sinaimg.cn/large/007l6U2xgy1g2jv7q6ibfj300q00q0sh.jpg');"></view>
					<!-- 杀 -->
					<view v-if="item.isShadi" class="tt bg-img" style="background-image: url('https://ws1.sinaimg.cn/large/007l6U2xgy1g2jv6c0cixj300q00q0q4.jpg');"></view>
					<view v-if="item.isZhu" class="tt bg-img" style="background-image: url('https://ws1.sinaimg.cn/large/007l6U2xgy1g2jv77etd4j300q00q0q1.jpg');"></view>
					<view v-if="item.isXiaobing" class="tt bg-img" style="background-image: url('https://ws1.sinaimg.cn/large/007l6U2xgy1g2jv86ombmj300q00q0o9.jpg');"></view>
					<view v-if="item.isT" class="tt bg-img" style="background-image: url('https://ws1.sinaimg.cn/large/007l6U2xgy1g2jv8q4jfbj300q00q0et.jpg');"></view>
					<view v-if="item.isJingji" class="tt bg-img" style="background-image: url('https://ws1.sinaimg.cn/large/007l6U2xgy1g2jv96iizlj300q00q0od.jpg');"></view>
					<view class="more-msg cuIcon-unfold text-gray text-lg"></view>
				</view>
				<view class="more-msg-box" v-if="lowShowIndex == i">
					<view class="msg-b" style="width: 230upx;">杀了多少兵：{{item.KillUnitCount}}</view>
					<view class="msg-b" style="width: 230upx;">拆了多少塔：{{item.TowerDestroy}}</view>
					<view class="msg-b" style="width: 230upx;">得了多少经验：{{item.RewardExp}}</view>
					<view class="msg-b" style="width: 230upx;">赢了多少把：{{item.WinCount}}/{{item.MatchCount}}</view>
					<view class="msg-b" style="width: 230upx;">第一次赢吗：{{[item.IsFirstWin == 1?'首胜':'不是']}}</view>
					<view class="msg-b" style="width: 230upx;">拿到多少金币：{{item.RewardMoney}}</view>
					<view class="msg-b" style="width: 690upx;">你的团队实力（团分）：{{item.ELO}}</view>
					<view class="msg-b" style="width: 690upx;">你在本场中的表现：{{item.KDA}}</view>
				</view>
			</view>
		</block>



		<view v-if="show==0" class="tt-state bg-img" :class="{'false':match.myteamState == 1,'success':match.myteamState == 2}">
			<view class="team-state text-white">
				敌方{{[match.myteamState == 2?'胜':'败']}}
			</view>		
			<view class="qian">{{[match.myteamState == 2?match.winJingji:match.lowJingji]}}</view>
			<view class="sha">
				{{[match.myteamState == 2? match.winShadi+'/'+match.winDeath+'/'+match.winZhu:match.lowShadi+'/'+match.lowDeath+'/'+match.lowZhu]}}
			</view>
		</view>


		<block v-if="match.myteamState == 2 &&  show==0">


			<view v-for="(item,i) in match.WinSide" :key="i" data-type="win" :data-key="i" class="hero-item" hover-class="hero-item-hover"
			 @tap="toMasg">
				<view class="msg">
					<view class="avatar-box" :data-nickname="item.RoleName" @tap.stop="toList">
						<view class="avatar bg-img round" :style="'background-image: url('+item.HeroIMG+');'">
						</view>
						<view v-if="item.isMe" class="tag round bg-img"></view>
						<view class="lv round">{{item.RoleLevel}}</view>
					</view>
					<view class="mid">
						<view class="nickname" :data-nickname="item.RoleName" @tap.stop="toList" :style="item.isMe?'color:#dbc56c;':''">
							{{item.RoleName}}
						</view>
						<view class="zhuangbei-list">
							<view v-for="(sk,index) in item.Skill" :key="index" class="zhuangbei-item radius bg-img" :style="'background-image: url(http://300report.jumpw.com/static/images/'+sk.IconFile+');'">
							</view>
							<view class="block" style="width: 6upx;height: 36upx;border-radius: 3upx;background-color: gray;margin: 0 5upx;"></view>

							<view v-for="(sk,index) in item.Equip" :key="index" class="zhuangbei-item radius bg-img" :style="'background-image: url(http://300report.jumpw.com/static/images/'+sk.IconFile+');'">
							</view>
						</view>
					</view>
					<view class="qian">
						<view class="bg-img" style="background-image: url('https://ws1.sinaimg.cn/large/007l6U2xgy1g2ifgccuccj300u00u0rt.jpg');"></view>
						{{item.TotalMoney}}
					</view>
					<view class="sha">
						<view class="sha-sha">
							<view class="bg-img" style="background-image: url('https://ws1.sinaimg.cn/large/007l6U2xgy1g2iff1bgamj300u00u741.jpg');"></view>
							{{item.KillCount}}/{{item.DeathCount}}/{{item.AssistCount}}
						</view>
						<view v-if="item.isM" class="bg-img mvp"></view>
					</view>
				</view>
				<view class="ttag">
					<view v-if="item.isTao" class="tt bg-img" style="background-image: url('https://ws1.sinaimg.cn/large/007l6U2xgy1g2jv7q6ibfj300q00q0sh.jpg');"></view>
					<!-- 杀 -->
					<view v-if="item.isShadi" class="tt bg-img" style="background-image: url('https://ws1.sinaimg.cn/large/007l6U2xgy1g2jv6c0cixj300q00q0q4.jpg');"></view>
					<view v-if="item.isZhu" class="tt bg-img" style="background-image: url('https://ws1.sinaimg.cn/large/007l6U2xgy1g2jv77etd4j300q00q0q1.jpg');"></view>
					<view v-if="item.isXiaobing" class="tt bg-img" style="background-image: url('https://ws1.sinaimg.cn/large/007l6U2xgy1g2jv86ombmj300q00q0o9.jpg');"></view>
					<view v-if="item.isT" class="tt bg-img" style="background-image: url('https://ws1.sinaimg.cn/large/007l6U2xgy1g2jv8q4jfbj300q00q0et.jpg');"></view>
					<view v-if="item.isJingji" class="tt bg-img" style="background-image: url('https://ws1.sinaimg.cn/large/007l6U2xgy1g2jv96iizlj300q00q0od.jpg');"></view>
					<view class="more-msg cuIcon-unfold text-gray text-lg"></view>
				</view>
				<view class="more-msg-box" v-if="winShowIndex == i">
					<view class="msg-b" style="width: 230upx;">杀了多少兵：{{item.KillUnitCount}}</view>
					<view class="msg-b" style="width: 230upx;">拆了多少塔：{{item.TowerDestroy}}</view>
					<view class="msg-b" style="width: 230upx;">得了多少经验：{{item.RewardExp}}</view>
					<view class="msg-b" style="width: 230upx;">赢了多少把：{{item.WinCount}}/{{item.MatchCount}}</view>
					<view class="msg-b" style="width: 230upx;">第一次赢吗：{{[item.IsFirstWin == 1?'首胜':'不是']}}</view>
					<view class="msg-b" style="width: 230upx;">拿到多少金币：{{item.RewardMoney}}</view>
					<view class="msg-b" style="width: 690upx;">你的团队实力（团分）：{{item.ELO}}</view>
					<view class="msg-b" style="width: 690upx;">你在本场中的表现：{{item.KDA}}</view>
				</view>
			</view>
		</block>
		<block v-else-if="show==0">
			<view v-for="(item,i) in match.LoseSide" :key="i" data-type="low" :data-key="i" class="hero-item" hover-class="hero-item-hover"
			 @tap="toMasg">
				<view class="msg">
					<view class="avatar-box" :data-nickname="item.RoleName" @tap.stop="toList">
						<view class="avatar bg-img round" :style="'background-image: url('+item.HeroIMG+');'">
						</view>
						<view v-if="item.isMe" class="tag round bg-img"></view>
						<view class="lv round">{{item.RoleLevel}}</view>
					</view>
					<view class="mid">
						<view class="nickname" :data-nickname="item.RoleName" @tap.stop="toList" :style="item.isMe?'color:#dbc56c;':''">
							{{item.RoleName}}
						</view>
						<view class="zhuangbei-list">
							<view v-for="(sk,index) in item.Skill" :key="index" class="zhuangbei-item radius bg-img" :style="'background-image: url(http://300report.jumpw.com/static/images/'+sk.IconFile+');'">
							</view>
							<view class="block" style="width: 6upx;height: 36upx;border-radius: 3upx;background-color: gray;margin: 0 5upx;"></view>

							<view v-for="(sk,index) in item.Equip" :key="index" class="zhuangbei-item radius bg-img" :style="'background-image: url(http://300report.jumpw.com/static/images/'+sk.IconFile+');'">
							</view>
						</view>
					</view>
					<view class="qian">
						<view class="bg-img" style="background-image: url('https://ws1.sinaimg.cn/large/007l6U2xgy1g2ifgccuccj300u00u0rt.jpg');"></view>
						{{item.TotalMoney}}
					</view>
					<view class="sha">
						<view class="sha-sha">
							<view class="bg-img" style="background-image: url('https://ws1.sinaimg.cn/large/007l6U2xgy1g2iff1bgamj300u00u741.jpg');"></view>
							{{item.KillCount}}/{{item.DeathCount}}/{{item.AssistCount}}
						</view>
						<view v-if="item.isM" class="bg-img mvp"></view>
					</view>
				</view>
				<view class="ttag">
					<view v-if="item.isTao" class="tt bg-img" style="background-image: url('https://ws1.sinaimg.cn/large/007l6U2xgy1g2jv7q6ibfj300q00q0sh.jpg');"></view>
					<!-- 杀 -->
					<view v-if="item.isShadi" class="tt bg-img" style="background-image: url('https://ws1.sinaimg.cn/large/007l6U2xgy1g2jv6c0cixj300q00q0q4.jpg');"></view>
					<view v-if="item.isZhu" class="tt bg-img" style="background-image: url('https://ws1.sinaimg.cn/large/007l6U2xgy1g2jv77etd4j300q00q0q1.jpg');"></view>
					<view v-if="item.isXiaobing" class="tt bg-img" style="background-image: url('https://ws1.sinaimg.cn/large/007l6U2xgy1g2jv86ombmj300q00q0o9.jpg');"></view>
					<view v-if="item.isT" class="tt bg-img" style="background-image: url('https://ws1.sinaimg.cn/large/007l6U2xgy1g2jv8q4jfbj300q00q0et.jpg');"></view>
					<view v-if="item.isJingji" class="tt bg-img" style="background-image: url('https://ws1.sinaimg.cn/large/007l6U2xgy1g2jv96iizlj300q00q0od.jpg');"></view>
					<view class="more-msg cuIcon-unfold text-gray text-lg"></view>
				</view>
				<view class="more-msg-box" v-if="lowShowIndex == i">
					<view class="msg-b" style="width: 230upx;">杀了多少兵：{{item.KillUnitCount}}</view>
					<view class="msg-b" style="width: 230upx;">拆了多少塔：{{item.TowerDestroy}}</view>
					<view class="msg-b" style="width: 230upx;">得了多少经验：{{item.RewardExp}}</view>
					<view class="msg-b" style="width: 230upx;">赢了多少把：{{item.WinCount}}/{{item.MatchCount}}</view>
					<view class="msg-b" style="width: 230upx;">第一次赢吗：{{[item.IsFirstWin == 1?'首胜':'不是']}}</view>
					<view class="msg-b" style="width: 230upx;">拿到多少金币：{{item.RewardMoney}}</view>
					<view class="msg-b" style="width: 690upx;">你的团队实力（团分）：{{item.ELO}}</view>
					<view class="msg-b" style="width: 690upx;">你在本场中的表现：{{item.KDA}}</view>
				</view>
			</view>
		</block>

	</view>
</template>

<script>
	export default {
		data() {
			return {
				winShowIndex: -1,
				lowShowIndex: -1,
				match: {},
				show: 1
			}
		},
		onLoad: function(doc) {
			this.show = 1
			let mid = doc.matchID
			let uid = doc.roleid
			let requestPromise = myUrl => {
				return new Promise((resolve, reject) => {
					uni.request({
						url: myUrl,
						success: res => resolve(res),
						fail: () => reject()

					})
				})
			}
			requestPromise('https://300hero.5idown.cn/index.php?/api/getmatch?id=' + mid).then(
				res => {
					let d = res.data
					if (d.Result != 'OK') {
						// 查找失败
						return
					}
					d = d.Match
					let h = Math.floor(d.UsedTime / 3600)
					let m = Math.floor(d.UsedTime / 60)
					let s = d.UsedTime % 60
					d.UsedTime = h > 0 ? h + '时' + m + '分' + s + '秒' : m + '分' + s + '秒'
					// 塔杀 小兵 助攻 杀敌

					let winTa = 0
					let winXiaobing = 0
					let winZhu = 0
					let winShadi = 0
					let winJingji = 0
					let winDeath = 0

					let maxTa = 0
					let maxTaIndex = -1
					let maxXiaobing = 0
					let maxXiaobingIndex = -1
					let maxZhu = 0
					let maxZhuIndex = -1
					let maxShadi = 0
					let maxShadiIndex = -1
					let maxJingji = 0
					let maxJingjiIndex = -1
					let maxKDA = 0
					let maxKDAIndex = -1

					d.myteamState = 0
					// 先获取胜利方数据
					for (let i = 0; i < 7; i++) {
						d.WinSide[i].HeroIMG = 'http://300report.jumpw.com/static/images/' + d.WinSide[i].Hero.IconFile

						// 判断自己赢了还是输了
						if (d.WinSide[i].RoleID == uid) {
							d.myteamState = 1
							d.WinSide[i].isMe = true
						}
						if (d.WinSide[i].TowerDestroy > maxTa) {
							maxTa = d.WinSide[i].TowerDestroy
							maxTaIndex = i
						}
						if (d.WinSide[i].KillUnitCount > maxXiaobing) {
							maxXiaobing = d.WinSide[i].KillUnitCount
							maxXiaobingIndex = i
						}

						if (d.WinSide[i].AssistCount > maxZhu) {
							maxZhu = d.WinSide[i].AssistCount
							maxZhuIndex = i
						}
						if (d.WinSide[i].KillCount > maxShadi) {
							maxShadi = d.WinSide[i].KillCount
							maxShadiIndex = i
						}

						if (d.WinSide[i].TotalMoney > maxJingji) {
							maxJingji = d.WinSide[i].TotalMoney
							maxJingjiIndex = i
						}

						if (d.WinSide[i].KDA > maxKDA) {
							maxKDA = d.WinSide[i].KDA
							maxKDAIndex = i
						}
						d.WinSide[i].isTao = d.WinSide[i].Result == 3



						winTa = winTa + d.WinSide[i].TowerDestroy
						winXiaobing = winXiaobing + d.WinSide[i].KillUnitCount
						winZhu = winZhu + d.WinSide[i].AssistCount
						winShadi = winShadi + d.WinSide[i].KillCount
						winJingji = winJingji + d.WinSide[i].TotalMoney
						winDeath = winDeath + d.WinSide[i].DeathCount
					}
					if (maxTaIndex != -1) {
						d.WinSide[maxTaIndex].isT = true
					}
					if (maxXiaobingIndex != -1) {
						d.WinSide[maxXiaobingIndex].isXiaobing = true
					}
					if (maxZhuIndex != -1) {
						d.WinSide[maxZhuIndex].isZhu = true
					}
					if (maxShadiIndex != -1) {
						d.WinSide[maxShadiIndex].isShadi = true
					}
					if (maxJingjiIndex != -1) {
						d.WinSide[maxJingjiIndex].isJingji = true
					}
					if (maxKDAIndex != -1 && maxKDA != 0) {
						d.WinSide[maxKDAIndex].isM = true
					}


					d.winShadi = winShadi
					d.winZhu = winZhu
					d.winXiaobing = winXiaobing
					d.winTa = winTa
					d.winJingji = winJingji
					d.winDeath = winDeath


					let lowTa = 0
					let lowXiaobing = 0
					let lowZhu = 0
					let lowShadi = 0
					let lowJingji = 0
					let lowDeath = 0


					let lmaxTa = 0
					let lmaxTaIndex = -1
					let lmaxXiaobing = 0
					let lmaxXiaobingIndex = -1
					let lmaxZhu = 0
					let lmaxZhuIndex = -1
					let lmaxShadi = 0
					let lmaxShadiIndex = -1
					let lmaxJingji = 0
					let lmaxJingjiIndex = -1
					let lmaxKDA = 0
					let lmaxKDAIndex = -1

					for (let i = 0; i < 7; i++) {
						d.LoseSide[i].HeroIMG = 'http://300report.jumpw.com/static/images/' + d.LoseSide[i].Hero.IconFile
						// 判断自己赢了还是输了
						if (d.LoseSide[i].RoleID == uid) {
							d.myteamState = 2
							d.LoseSide[i].isMe = true
						}

						if (d.LoseSide[i].TowerDestroy > lmaxTa) {
							lmaxTa = d.LoseSide[i].TowerDestroy
							lmaxTaIndex = i
						}
						if (d.LoseSide[i].KillUnitCount > lmaxXiaobing) {
							lmaxXiaobing = d.LoseSide[i].KillUnitCount
							lmaxXiaobingIndex = i
						}

						if (d.LoseSide[i].AssistCount > lmaxZhu) {
							lmaxZhu = d.LoseSide[i].AssistCount
							lmaxZhuIndex = i
						}
						if (d.LoseSide[i].KillCount > lmaxShadi) {
							lmaxShadi = d.LoseSide[i].KillCount
							lmaxShadiIndex = i
						}

						if (d.LoseSide[i].TotalMoney > lmaxJingji) {
							lmaxJingji = d.LoseSide[i].TotalMoney
							lmaxJingjiIndex = i
						}

						if (d.LoseSide[i].KDA > lmaxKDA) {
							lmaxKDA = d.LoseSide[i].KDA
							lmaxKDAIndex = i
						}


						d.LoseSide[i].isTao = d.LoseSide[i].Result == 3

						lowTa = lowTa + d.LoseSide[i].TowerDestroy
						lowXiaobing = lowXiaobing + d.LoseSide[i].KillUnitCount
						lowZhu = lowZhu + d.LoseSide[i].AssistCount
						lowShadi = lowShadi + d.LoseSide[i].KillCount
						lowJingji = lowJingji + d.LoseSide[i].TotalMoney
						lowDeath = lowDeath + d.LoseSide[i].DeathCount
					}
					if (lmaxTaIndex != -1) {
						d.LoseSide[lmaxTaIndex].isT = true
					}
					if (lmaxXiaobingIndex != -1) {
						d.LoseSide[lmaxXiaobingIndex].isXiaobing = true
					}
					if (lmaxZhuIndex != -1) {
						d.LoseSide[lmaxZhuIndex].isZhu = true
					}
					if (lmaxShadiIndex != -1) {
						d.LoseSide[lmaxShadiIndex].isShadi = true
					}
					if (lmaxJingjiIndex != -1  && lmaxKDA != 0) {
						d.LoseSide[lmaxJingjiIndex].isJingji = true
					}





					d.lowShadi = lowShadi
					d.lowZhu = lowZhu
					d.lowXiaobing = lowXiaobing
					d.lowTa = lowTa
					d.lowJingji = lowJingji
					d.lowDeath = lowDeath

					this.match = d
				}
			).then(
				() => {
					this.show = 0
				}
			)
		},
		methods: {
			toMasg: function(e) {
				let t = e.currentTarget.dataset.type
				let k = e.currentTarget.dataset.key
				switch (t) {
					case 'win':
						if (k == this.winShowIndex){
							this.winShowIndex = -1
							return
						}
						this.winShowIndex = k
						this.lowShowIndex = -1
						break;
					case 'low':
						if (k == this.lowShowIndex){
							this.lowShowIndex = -1
							return
						}
						this.lowShowIndex = k
						this.winShowIndex = -1
						break
				}
			},
			toList:function(e){
				let t = e.currentTarget.dataset.nickname
				uni.navigateTo({
					url:'/pages/list/list?nickname='+t
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

	.match-state-box {
		height: 54upx;
		background-color: #EDF0F0;
		position: relative;
		display: flex;
		justify-content: flex-start;
		align-items: center;
		width: 100%;
		font-size: 22upx;
		font-weight: bold;
		color: #788A8D;
		padding-right: 50upx;
	}

	.match-state-box.false {
		background-image: url('https://ws1.sinaimg.cn/large/007l6U2xgy1g2iarz0hngj30u0026glh.jpg');
	}

	.match-state-box.success {
		background-image: url('https://ws1.sinaimg.cn/large/007l6U2xgy1g2jszc79tpj30u0025gli.jpg');
	}

	.match-state-box .time {
		margin-left: 115upx;
	}

	.match-state-box .slip,
	.match-state-box .hs,
	.match-state-box .type {
		margin-left: auto;
	}

	.match-box {
		height: 95upx;
		width: 100%;
		padding-top: 55upx;
		padding-left: 6upx;
		padding-right: 6upx;
		color: #686868;
		position: relative;
		display: flex;
		justify-content: flex-start;
		align-items: center;
	}

	.match-box.false {
		background-image: url('https://ws1.sinaimg.cn/large/007l6U2xgy1g2id1qki6ej30u003uaa8.jpg');
	}

	.match-box.success {
		background-image: url('https://ws1.sinaimg.cn/large/007l6U2xgy1g2jsj8flmjj30u003udg0.jpg');
	}

	.match-box .mymatch,
	.match-box .othermatch {
		width: 264upx;
		position: relative;
		display: flex;
		justify-content: flex-start;
		align-items: center;
	}

	.match-box .othermatch {
		margin-left: auto;
	}

	.match-box .mymatch view {
		width: 66upx;
		text-align: center;
		font-size: 22upx;
		font-weight: bold;
		color: #686868;
	}

	.match-box .othermatch view {
		width: 66upx;
		text-align: center;
		font-size: 22upx;
		font-weight: bold;
		color: #686868;
	}

	.tt-state {
		height: 58upx;
		width: 100%;
		font-size: 20upx;
		font-weight: bold;
		position: relative;
		line-height: 58upx;
	}

	.tt-state.success {
		color: #00C1D3;
		background-image: url('https://ws1.sinaimg.cn/large/007l6U2xgy1g2jr2a9flqj30u002bt8j.jpg');
	}

	.tt-state.false {
		color: #ff7578;
		background-image: url('https://ws1.sinaimg.cn/large/007l6U2xgy1g2jr2j3hauj30u002ba9w.jpg');
	}

	.tt-state .team-state {
		position: absolute;
		left: 20upx;
		font-size: 24upx;
		font-weight: normal;
	}

	.tt-state .qian {
		position: absolute;
		left: 540upx;
	}

	.tt-state .sha {
		position: absolute;
		left: 645upx;
	}

	.hero-item {
		padding: 30upx 0 0 0;
		border-bottom: 2upx solid #f4f4f4;
		background-color: #FCFCFC;
	}

	.hero-item .msg {
		display: flex;
		justify-content: flex-start;
		align-content: center;
		padding: 0 30upx;
	}

	.hero-item .msg .avatar-box {
		width: 78upx;
		height: 78upx;
		position: relative;
		margin-right: 20upx;
	}

	.hero-item .msg .avatar-box .avatar {
		width: 78upx;
		height: 78upx;
	}

	.hero-item .msg .avatar-box .tag {
		width: 26upx;
		height: 26upx;
		position: absolute;
		left: 0;
		top: 0;
		background-image: url('https://ws1.sinaimg.cn/large/007l6U2xgy1g2ieqhx91uj302s02sgle.jpg');
	}

	.hero-item .msg .avatar-box .lv {
		width: 30upx;
		height: 30upx;
		background-color: #606060;
		position: absolute;
		right: -15upx;
		bottom: 0;
		color: white;
		font-size: 20upx;
		font-weight: bold;
		font-style: italic;
		line-height: 30upx;
		text-align: center;
	}

	.hero-item .msg .mid .nickname {
		color: #33474b;
		font-size: 30upx;
	}

	.hero-item .msg .mid .nickname.me {
		color: #cdad2f;
	}

	.hero-item .msg .mid .zhuangbei-list {
		display: flex;
		position: relative;
		justify-content: flex-start;
		align-items: center;
		margin-top: 20upx;
	}

	.hero-item .msg .mid .zhuangbei-list .zhuangbei-item {
		width: 36upx;
		height: 36upx;
		margin-right: 3upx;
	}

	.hero-item .msg .qian {
		margin-left: auto;
		margin-right: 10upx;
	}

	.hero-item .msg .qian,
	.hero-item .msg .sha .sha-sha {
		line-height: 30upx;
		color: #95a9ad;
		font-size: 20upx;
		display: flex;
		font-weight: bold;
	}

	.hero-item .msg .sha .mvp {
		background-image: url('https://ws1.sinaimg.cn/large/007l6U2xgy1g2ih83ss4ij301600qgld.jpg');
		width: 63upx;
		height: 39upx;
		margin-left: auto;
		margin-top: 20upx;
	}

	.hero-item .msg .sha,
	.hero-item .msg .qian {
		margin-top: 10upx;
	}

	.hero-item .msg .qian view,
	.hero-item .msg .sha .sha-sha view {
		width: 30upx;
		height: 30upx;
		margin-right: 2upx;
	}

	.hero-item .ttag {
		margin-top: 15upx;
		margin-left: 130upx;
		padding-right: 30upx;
		display: flex;
		justify-content: flex-start;
		align-items: center;
		margin-bottom: 25upx;
	}

	.hero-item .ttag .tt {
		margin-right: 10upx;
		width: 26upx;
		height: 26upx;
	}

	.hero-item .ttag .more-msg {
		margin-left: auto;
	}

	.more-msg-box {
		background-color: #f6f6f6;
		width: 100%;
		display: flex;
		font-size: 24upx;
		color: #aeaeae;
		flex-wrap: wrap;
		height: 200upx;
		overflow: hidden;
		padding: 10upx 30upx;
	}




	.hero-item .more-msg-box .b {
		height: 45upx;
		padding: 0;
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
</style>
