package rules

func NewChinaSitesRule(tag string) *Rule {
	return &Rule{
		Tag:       tag,
		Condition: chinaSitesConds,
	}
}

const (
	anySubDomain = "^(.*\\.)?"
	dotAm        = "\\.am$"
	dotCc        = "\\.cc$"
	dotCn        = "\\.cn$"
	dotCom       = "\\.com$"
	dotInfo      = "\\.info$"
	dotIo        = "\\.io$"
	dotLa        = "\\.la$"
	dotMe        = "\\.me$"
	dotNet       = "\\.net$"
	dotOrg       = "\\.org$"
	dotTv        = "\\.tv$"
)

var (
	chinaSitesConds Condition
)

func init() {
	regexpDomains := []string{
		dotCn,
		"\\.xn--fiqs8s$", /* .中国 */

		anySubDomain + "10010" + dotCom,
		anySubDomain + "100offer" + dotCom,
		anySubDomain + "115" + dotCom,
		anySubDomain + "123juzi" + dotCom,
		anySubDomain + "123juzi" + dotNet,
		anySubDomain + "123u" + dotCom,
		anySubDomain + "126" + dotCom,
		anySubDomain + "126" + dotNet,
		anySubDomain + "127" + dotNet,
		anySubDomain + "163" + dotCom,
		anySubDomain + "17173" + dotCom,
		anySubDomain + "17cdn" + dotCom,
		anySubDomain + "188" + dotCom,
		anySubDomain + "1905" + dotCom,
		anySubDomain + "21cn" + dotCom,
		anySubDomain + "2288" + dotOrg,
		anySubDomain + "2345" + dotCom,
		anySubDomain + "263" + dotNet,
		anySubDomain + "2cto" + dotCom,
		anySubDomain + "3322" + dotOrg,
		anySubDomain + "35" + dotCom,
		anySubDomain + "360doc" + dotCom,
		anySubDomain + "360buy" + dotCom,
		anySubDomain + "360buyimg" + dotCom,
		anySubDomain + "360safe" + dotCom,
		anySubDomain + "36kr" + dotCom,
		anySubDomain + "39" + dotNet,
		anySubDomain + "3dmgame" + dotCom,
		anySubDomain + "3conline" + dotCom,
		anySubDomain + "4399" + dotCom,
		anySubDomain + "500d" + dotMe,
		anySubDomain + "50bang" + dotOrg,
		anySubDomain + "51" + dotLa,
		anySubDomain + "51credit" + dotCom,
		anySubDomain + "51cto" + dotCom,
		anySubDomain + "51job" + dotCom,
		anySubDomain + "51jobcdn" + dotCom,
		anySubDomain + "51wendang" + dotCom,
		anySubDomain + "55" + dotCom,
		anySubDomain + "51yes" + dotCom,
		anySubDomain + "55bbs" + dotCom,
		anySubDomain + "58" + dotCom,
		anySubDomain + "6rooms" + dotCom,
		anySubDomain + "71" + dotAm,
		anySubDomain + "7k7k" + dotCom,
		anySubDomain + "900" + dotLa,
		anySubDomain + "9718" + dotCom,
		anySubDomain + "9xu" + dotCom,
		anySubDomain + "abchina" + dotCom,
		anySubDomain + "acfun" + dotTv,
		anySubDomain + "acgvideo" + dotCom,
		anySubDomain + "agrantsem" + dotCom,
		anySubDomain + "aicdn" + dotCom,
		anySubDomain + "aixifan" + dotCom,
		anySubDomain + "alibaba" + dotCom,
		anySubDomain + "alicdn" + dotCom,
		anySubDomain + "aliimg.com" + dotCom,
		anySubDomain + "alipay" + dotCom,
		anySubDomain + "alipayobjects" + dotCom,
		anySubDomain + "aliyun" + dotCom,
		anySubDomain + "aliyuncdn" + dotCom,
		anySubDomain + "aliyuncs" + dotCom,
		anySubDomain + "allyes" + dotCom,
		anySubDomain + "amap" + dotCom,
		anySubDomain + "anjuke" + dotCom,
		anySubDomain + "anquan" + dotOrg,
		anySubDomain + "appinn" + dotCom,
		anySubDomain + "babytree" + dotCom,
		anySubDomain + "babytreeimg" + dotCom,
		anySubDomain + "baidu" + dotCom,
		anySubDomain + "baiducontent" + dotCom,
		anySubDomain + "baidupcs" + dotCom,
		anySubDomain + "baidustatic" + dotCom,
		anySubDomain + "baifendian" + dotCom,
		anySubDomain + "baifubao" + dotCom,
		anySubDomain + "baihe" + dotCom,
		anySubDomain + "baike" + dotCom,
		anySubDomain + "baixing" + dotCom,
		anySubDomain + "baixing" + dotNet,
		anySubDomain + "bankcomm" + dotCom,
		anySubDomain + "bankofchina" + dotCom,
		anySubDomain + "bcy" + dotNet,
		anySubDomain + "bdimg" + dotCom,
		anySubDomain + "bdstatic" + dotCom,
		anySubDomain + "bilibili" + dotCom,
		"cn\\.bing" + dotCom,
		anySubDomain + "bitauto" + dotCom,
		anySubDomain + "bitautoimg" + dotCom,
		anySubDomain + "bobo" + dotCom,
		anySubDomain + "btcfans" + dotCom,
		anySubDomain + "caiyunapp" + dotCom,
		anySubDomain + "ccb" + dotCom,
		anySubDomain + "cctv" + dotCom,
		anySubDomain + "cctvpic" + dotCom,
		anySubDomain + "cdn20" + dotCom,
		anySubDomain + "cebbank" + dotCom,
		anySubDomain + "ch" + dotCom,
		anySubDomain + "chashebao" + dotCom,
		anySubDomain + "che168" + dotCom,
		anySubDomain + "china" + dotCom,
		anySubDomain + "chinacache" + dotCom,
		anySubDomain + "chinacache" + dotNet,
		anySubDomain + "chinahr" + dotCom,
		anySubDomain + "chinamobile" + dotCom,
		anySubDomain + "chinatranslation" + dotNet,
		anySubDomain + "chinaz" + dotCom,
		anySubDomain + "chouti" + dotCom,
		anySubDomain + "chuangxin" + dotCom,
		anySubDomain + "chuansong" + dotMe,
		anySubDomain + "clouddn" + dotCom,
		anySubDomain + "cloudxns" + dotCom,
		anySubDomain + "cmbchina" + dotCom,
		anySubDomain + "cnbeta" + dotCom,
		anySubDomain + "cnbetacdn" + dotCom,
		anySubDomain + "cnblogs" + dotCom,
		anySubDomain + "cnepub" + dotCom,
		anySubDomain + "cnzz" + dotCom,
		anySubDomain + "coding" + dotNet,
		anySubDomain + "cqvip" + dotCom,
		anySubDomain + "csbew" + dotCom,
		anySubDomain + "csdn" + dotNet,
		anySubDomain + "ctrip" + dotCom,
		anySubDomain + "cubead" + dotCom,
		anySubDomain + "dajie" + dotCom,
		anySubDomain + "dajieimg" + dotCom,
		anySubDomain + "dangdang" + dotCom,
		anySubDomain + "daocloud" + dotIo,
		anySubDomain + "daovoice" + dotIo,
		anySubDomain + "dbank" + dotCom,
		anySubDomain + "dedecms" + dotCom,
		anySubDomain + "diandian" + dotCom,
		anySubDomain + "dianping" + dotCom,
		anySubDomain + "diopic" + dotNet,
		anySubDomain + "docin" + dotCom,
		anySubDomain + "dockerone" + dotCom,
		anySubDomain + "dockone" + dotIo,
		anySubDomain + "donews" + dotCom,
		anySubDomain + "douban" + dotCom,
		anySubDomain + "doubanio" + dotCom,
		anySubDomain + "dpfile" + dotCom,
		anySubDomain + "duomai" + dotCom,
		anySubDomain + "duoshuo" + dotCom,
		anySubDomain + "duowan" + dotCom,
		anySubDomain + "dxpmedia" + dotCom,
		anySubDomain + "eastday" + dotCom,
		anySubDomain + "ecitic" + dotCom,
		anySubDomain + "emarbox" + dotCom,
		anySubDomain + "eoeandroid" + dotCom,
		anySubDomain + "etao" + dotCom,
		anySubDomain + "fanli" + dotCom,
		anySubDomain + "fengniao" + dotCom,
		anySubDomain + "fhldns" + dotCom,
		anySubDomain + "foxmail" + dotCom,
		anySubDomain + "geekpark" + dotNet,
		anySubDomain + "geetest" + dotCom,
		anySubDomain + "geilicdn" + dotCom,
		anySubDomain + "getui" + dotCom,
		anySubDomain + "growingio" + dotCom,
		anySubDomain + "gtags" + dotNet,
		anySubDomain + "gwdang" + dotCom,
		anySubDomain + "hao123" + dotCom,
		anySubDomain + "hao123img" + dotCom,
		anySubDomain + "haosou" + dotCom,
		anySubDomain + "hdslb" + dotCom,
		anySubDomain + "henha" + dotCom,
		anySubDomain + "hexun" + dotCom,
		anySubDomain + "hichina" + dotCom,
		anySubDomain + "huanqiu" + dotCom,
		anySubDomain + "hunantv" + dotCom,
		anySubDomain + "huochepiao" + dotCom,
		anySubDomain + "hupu" + dotCom,
		anySubDomain + "hupucdn" + dotCom,
		anySubDomain + "huxiu" + dotCom,
		anySubDomain + "iask" + dotCom,
		anySubDomain + "iciba" + dotCom,
		anySubDomain + "idqqimg" + dotCom,
		anySubDomain + "ifanr" + dotCom,
		anySubDomain + "ifanrusercontent" + dotCom,
		anySubDomain + "ifanrx" + dotCom,
		anySubDomain + "ifeng" + dotCom,
		anySubDomain + "ifengimg" + dotCom,
		anySubDomain + "ijinshan" + dotCom,
		anySubDomain + "imedao" + dotCom,
		anySubDomain + "imgo" + dotTv,
		anySubDomain + "imooc" + dotCom,
		anySubDomain + "infoq" + dotCom,
		anySubDomain + "infoqstatic" + dotCom,
		anySubDomain + "ip138" + dotCom,
		anySubDomain + "ipinyou" + dotCom,
		anySubDomain + "ipip" + dotNet,
		anySubDomain + "ip-cdn" + dotCom,
		anySubDomain + "iqiyi" + dotCom,
		anySubDomain + "it165" + dotNet,
		anySubDomain + "it168" + dotCom,
		anySubDomain + "it610" + dotCom,
		anySubDomain + "iteye" + dotCom,
		anySubDomain + "itjuzi" + dotCom,
		anySubDomain + "jandan" + dotNet,
		anySubDomain + "jd" + dotCom,
		anySubDomain + "jb51" + dotCom,
		anySubDomain + "jia" + dotCom,
		anySubDomain + "jianshu" + dotCom,
		anySubDomain + "jianshu" + dotIo,
		anySubDomain + "jiasuhui" + dotCom,
		anySubDomain + "jiathis" + dotCom,
		anySubDomain + "jiayuan" + dotCom,
		anySubDomain + "jikexueyuan" + dotCom,
		anySubDomain + "jisuanke" + dotCom,
		anySubDomain + "jmstatic" + dotCom,
		anySubDomain + "jstv" + dotCom,
		anySubDomain + "jumei" + dotCom,
		anySubDomain + "jyimg" + dotCom,
		anySubDomain + "kaixin001" + dotCom,
		anySubDomain + "kanimg" + dotCom,
		anySubDomain + "kankanews" + dotCom,
		anySubDomain + "kejet" + dotNet,
		anySubDomain + "kf5" + dotCom,
		anySubDomain + "kimiss" + dotCom,
		anySubDomain + "kouclo" + dotCom,
		anySubDomain + "koudai" + dotCom,
		anySubDomain + "koudai8" + dotCom,
		anySubDomain + "ku6" + dotCom,
		anySubDomain + "ku6cdn" + dotCom,
		anySubDomain + "ku6img" + dotCom,
		anySubDomain + "kuqin" + dotCom,
		anySubDomain + "lady8844" + dotCom,
		anySubDomain + "lagou" + dotCom,
		anySubDomain + "le" + dotCom,
		anySubDomain + "leanote" + dotCom,
		anySubDomain + "leiphone" + dotCom,
		anySubDomain + "leju" + dotCom,
		anySubDomain + "leturich" + dotOrg,
		anySubDomain + "letv" + dotCom,
		anySubDomain + "letvcdn" + dotCom,
		anySubDomain + "letvimg" + dotCom,
		anySubDomain + "liantu" + dotMe,
		anySubDomain + "liaoxuefeng" + dotCom,
		anySubDomain + "liba" + dotCom,
		anySubDomain + "libaclub" + dotCom,
		anySubDomain + "liepin" + dotCom,
		anySubDomain + "lietou" + dotCom,
		anySubDomain + "lightonus" + dotCom,
		anySubDomain + "linkvans" + dotCom,
		anySubDomain + "linuxidc" + dotCom,
		anySubDomain + "liuxiaoer" + dotCom,
		anySubDomain + "lofter" + dotCom,
		anySubDomain + "lu" + dotCom,
		anySubDomain + "lufax" + dotCom,
		anySubDomain + "lufaxcdn" + dotCom,
		anySubDomain + "lvmama" + dotCom,
		anySubDomain + "lxdns" + dotCom,
		anySubDomain + "lxway" + dotCom,
		anySubDomain + "ly" + dotCom,
		anySubDomain + "mayihr" + dotCom,
		anySubDomain + "mechina" + dotOrg,
		anySubDomain + "mediav" + dotCom,
		anySubDomain + "meiqia" + dotCom,
		anySubDomain + "meika360" + dotCom,
		anySubDomain + "meilishuo" + dotCom,
		anySubDomain + "meishij" + dotNet,
		anySubDomain + "meituan" + dotCom,
		anySubDomain + "meizu" + dotCom,
		anySubDomain + "mgtv" + dotCom,
		anySubDomain + "mi" + dotCom,
		anySubDomain + "miaopai" + dotCom,
		anySubDomain + "miaozhen" + dotCom,
		anySubDomain + "mmbang" + dotCom,
		anySubDomain + "mmbang" + dotInfo,
		anySubDomain + "mmstat" + dotCom,
		anySubDomain + "mogucdn" + dotCom,
		anySubDomain + "mogujie" + dotCom,
		anySubDomain + "mop" + dotCom,
		anySubDomain + "mukewang" + dotCom,
		anySubDomain + "mydrivers" + dotCom,
		anySubDomain + "myshow360" + dotNet,
		anySubDomain + "mzstatic" + dotCom,
		anySubDomain + "netease" + dotCom,
		anySubDomain + "newbandeng" + dotCom,
		anySubDomain + "ngacn" + dotCc,
		anySubDomain + "ntalker" + dotCom,
		anySubDomain + "nvsheng" + dotCom,
		anySubDomain + "oeeee" + dotCom,
		anySubDomain + "ol-img" + dotCom,
		anySubDomain + "oneapm" + dotCom,
		anySubDomain + "onlinedown" + dotNet,
		anySubDomain + "onlinesjtu" + dotCom,
		anySubDomain + "oschina" + dotNet,
		anySubDomain + "paipai" + dotCom,
		anySubDomain + "pchome" + dotNet,
		anySubDomain + "pingan" + dotCom,
		anySubDomain + "pingplusplus" + dotCom,
		anySubDomain + "pps" + dotTv,
		anySubDomain + "psbc" + dotCom,
		anySubDomain + "pubyun" + dotCom,
		anySubDomain + "qbox" + dotMe,
		anySubDomain + "qcloud" + dotCom,
		anySubDomain + "qhimg" + dotCom,
		anySubDomain + "qiaobutang" + dotCom,
		anySubDomain + "qidian" + dotCom,
		anySubDomain + "qingcloud" + dotCom,
		anySubDomain + "qingsongchou" + dotCom,
		anySubDomain + "qiniu" + dotCom,
		anySubDomain + "qiniucdn" + dotCom,
		anySubDomain + "qiniudn" + dotCom,
		anySubDomain + "qiniudns" + dotCom,
		anySubDomain + "qiyi" + dotCom,
		anySubDomain + "qiyipic" + dotCom,
		anySubDomain + "qtmojo" + dotCom,
		anySubDomain + "qq" + dotCom,
		anySubDomain + "qqmail" + dotCom,
		anySubDomain + "qunar" + dotCom,
		anySubDomain + "qunarzz" + dotCom,
		anySubDomain + "qzone" + dotCom,
		anySubDomain + "renren" + dotCom,
		anySubDomain + "ruby-china" + dotOrg,
		anySubDomain + "sandai" + dotNet,
		anySubDomain + "sanguosha" + dotCom,
		anySubDomain + "sanwen" + dotNet,
		anySubDomain + "segmentfault" + dotCom,
		anySubDomain + "sf-express" + dotCom,
		anySubDomain + "sharejs" + dotCom,
		anySubDomain + "shutcm" + dotCom,
		anySubDomain + "simei8" + dotCom,
		anySubDomain + "sina" + dotCom,
		anySubDomain + "sinaapp" + dotCom,
		anySubDomain + "sinaedge" + dotCom,
		anySubDomain + "sinaimg" + dotCom,
		anySubDomain + "sinajs" + dotCom,
		anySubDomain + "szzfgjj" + dotCom,
		anySubDomain + "smzdm" + dotCom,
		anySubDomain + "sohu" + dotCom,
		anySubDomain + "sogou" + dotCom,
		anySubDomain + "sogoucdn" + dotCom,
		anySubDomain + "soso" + dotCom,
		anySubDomain + "sspai" + dotCom,
		anySubDomain + "starbaby" + dotCc,
		anySubDomain + "starbaby" + dotCom,
		anySubDomain + "staticfile" + dotOrg,
		anySubDomain + "stockstar" + dotCom,
		anySubDomain + "suning" + dotCom,
		anySubDomain + "szfw" + dotOrg,
		anySubDomain + "t1y5" + dotCom,
		anySubDomain + "tanx" + dotCom,
		anySubDomain + "tao123" + dotCom,
		anySubDomain + "taobao" + dotCom,
		anySubDomain + "taobaocdn" + dotCom,
		anySubDomain + "tbcache" + dotCom,
		anySubDomain + "tencent" + dotCom,
		anySubDomain + "tenpay" + dotCom,
		anySubDomain + "tenxcloud" + dotCom,
		anySubDomain + "tiebaimg" + dotCom,
		anySubDomain + "tietuku" + dotCom,
		anySubDomain + "tiexue" + dotNet,
		anySubDomain + "tmall" + dotCom,
		anySubDomain + "tmcdn" + dotNet,
		anySubDomain + "topthink" + dotCom,
		anySubDomain + "tudou" + dotCom,
		anySubDomain + "tudouui" + dotCom,
		anySubDomain + "tuicool" + dotCom,
		anySubDomain + "tuniu" + dotCom,
		anySubDomain + "u17" + dotCom,
		anySubDomain + "useso" + dotCom,
		anySubDomain + "unionpay" + dotCom,
		anySubDomain + "unionpaysecure" + dotCom,
		anySubDomain + "upyun" + dotCom,
		anySubDomain + "upaiyun" + dotCom,
		anySubDomain + "v2ex" + dotCom,
		anySubDomain + "v5875" + dotCom,
		anySubDomain + "vamaker" + dotCom,
		anySubDomain + "vancl" + dotCom,
		anySubDomain + "vip" + dotCom,
		anySubDomain + "wallstreetcn" + dotCom,
		anySubDomain + "wandoujia" + dotCom,
		anySubDomain + "wdjimg" + dotCom,
		anySubDomain + "webterren" + dotCom,
		anySubDomain + "weibo" + dotCom,
		anySubDomain + "weicaifu" + dotCom,
		anySubDomain + "weidian" + dotCom,
		anySubDomain + "weiyun" + dotCom,
		anySubDomain + "wonnder" + dotCom,
		anySubDomain + "worktile" + dotCom,
		anySubDomain + "wooyun" + dotOrg,
		anySubDomain + "wrating" + dotCom,
		anySubDomain + "wscdns" + dotCom,
		anySubDomain + "wumii" + dotCom,
		anySubDomain + "xiachufang" + dotCom,
		anySubDomain + "xiami" + dotCom,
		anySubDomain + "xiaokaxiu" + dotCom,
		anySubDomain + "xiaomi" + dotCom,
		anySubDomain + "xitu" + dotCom,
		anySubDomain + "xinhuanet" + dotCom,
		anySubDomain + "xinshipu" + dotCom,
		anySubDomain + "xiu8" + dotCom,
		anySubDomain + "xnpic" + dotCom,
		anySubDomain + "xueqiu" + dotCom,
		anySubDomain + "xunlei" + dotCom,
		anySubDomain + "xywy" + dotCom,
		anySubDomain + "yaolan" + dotCom,
		anySubDomain + "yccdn" + dotCom,
		anySubDomain + "yeepay" + dotCom,
		anySubDomain + "yesky" + dotCom,
		anySubDomain + "yigao" + dotCom,
		anySubDomain + "yihaodian" + dotCom,
		anySubDomain + "yihaodianimg" + dotCom,
		anySubDomain + "yingjiesheng" + dotCom,
		anySubDomain + "yinxiang" + dotCom,
		anySubDomain + "yjbys" + dotCom,
		anySubDomain + "yhd" + dotCom,
		anySubDomain + "youboy" + dotCom,
		anySubDomain + "youku" + dotCom,
		anySubDomain + "yunba" + dotIo,
		anySubDomain + "yundaex" + dotCom,
		anySubDomain + "yunshipei" + dotCom,
		anySubDomain + "yupoo" + dotCom,
		anySubDomain + "yuzua" + dotCom,
		anySubDomain + "yy" + dotCom,
		anySubDomain + "yytcdn" + dotCom,
		anySubDomain + "zampda" + dotNet,
		anySubDomain + "zastatic" + dotCom,
		anySubDomain + "zbjimg" + dotCom,
		anySubDomain + "zhenai" + dotCom,
		anySubDomain + "zhanqi" + dotTv,
		anySubDomain + "zhaopin" + dotCom,
		anySubDomain + "zhihu" + dotCom,
		anySubDomain + "zhimg" + dotCom,
		anySubDomain + "zhiziyun" + dotCom,
		anySubDomain + "zjstv" + dotCom,
		anySubDomain + "zhubajie" + dotCom,
		anySubDomain + "zrblog" + dotNet,
		anySubDomain + "zuche" + dotCom,
		anySubDomain + "zuchecdn" + dotCom,
	}

	conds := make([]Condition, len(regexpDomains))
	for idx, pattern := range regexpDomains {
		matcher, err := NewRegexpDomainMatcher(pattern)
		if err != nil {
			panic(err)
		}
		conds[idx] = matcher
	}

	anyConds := AnyCondition(conds)
	chinaSitesConds = &anyConds
}
